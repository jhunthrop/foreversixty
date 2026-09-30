# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.7. Weights run: 0.9s. Verify run: 1.2s. 169 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.290 ± 0.038, hit=0.591 ± 0.033, melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.37 DPS) | yes | Flying Tiger Goggles (4368, -0.45 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 attack_power points (0.28 DPS) | yes | Erudite's Amulet (277204, -0.11 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.23 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.28 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.09 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.2 attack_power points (0.51 DPS) | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.18 DPS) [crafted]; Dark Leather Tunic (2317, -0.23 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.23 DPS) | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.09 DPS) [crafted]; Forest Leather Gloves (3058, -0.09 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.82 DPS) | yes | Deviate Scale Belt (6468, -0.52 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.63 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 attack_power points (0.41 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 attack_power points (0.37 DPS) | yes | Footpads of the Fang (10411, -0.09 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.28 DPS, sim-verified) [dungeon] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 6.4 attack_power points (0.29 DPS) | yes | Protector's Band (20439, -0.11 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.28 DPS) | yes | Protector's Band (20439, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; The 1 Ring (8350, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (10.35 DPS) | yes | Assassin's Blade (1935, -0.32 DPS) [dungeon]; Evocator's Blade (2567, -0.65 DPS) [dungeon]; Buzzer Blade (2169, -1.48 DPS) [dungeon] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | 222.7 attack_power points (10.09 DPS) | yes | Assassin's Blade (1935, -0.02 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, -0.01 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Pyrewood Signet Ring; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Edward's Knife; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 169, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6255 Fishing Pole (JEFFTEST); 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 35.2. Weights run: 0.9s. Verify run: 1.1s. 300 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.421 ± 0.047, hit=0.238 ± 0.024, melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.1 attack_power points (0.47 DPS) | yes | Brawler's Leather Hood (252504, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.65 DPS) | yes | Sentinel's Medallion (19541, -0.30 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.37 DPS) [rep]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 attack_power points (0.52 DPS) | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.32 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.46 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.18 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.74 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.23 DPS) [quest]; Green Leather Armor (4255, -0.37 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.46 DPS) | yes | Jurassic Wristguards (6198, -0.18 DPS) [world]; Insignia Bracers (6410, -0.18 DPS) [world_drop]; Unearthed Bands (9428, -0.52 DPS, sim-verified) [dungeon] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.74 DPS) | yes | Pilferer's Gloves (7358, -0.40 DPS, sim-verified) [crafted]; Wolfclaw Gloves (1978, -0.46 DPS) [dungeon]; Toughened Leather Gloves (4253, -0.46 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.12 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.69 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.21 DPS) | yes | Troll's Bane Leggings (13114, -0.55 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.60 DPS) [crafted]; Petrolspill Leggings (9509, -0.60 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.56 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.1 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.09 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.09 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (14.97 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.02 DPS) [vendor]; Torturing Poker (7682, -1.18 DPS) [dungeon]; Thornspike (6681, -1.57 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (14.88 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -0.85 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -14.83 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Silver Star (3463, -0.20 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.23 DPS) [world_drop]; Crystalpine Stinger (13037, -0.23 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (night-elf, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 67.1. Weights run: 1.0s. Verify run: 1.3s. 429 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.663 ± 0.081, hit=0.583 ± 0.056, melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.2 attack_power points (0.65 DPS) | yes | Nightscape Headband (8176, -0.09 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (65.7 DPS) | yes | Sentinel's Medallion (19540, -0.14 DPS) [rep]; Sentinel's Medallion (19541, -0.29 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.84 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.15 DPS) | yes | Forest Tracker Epaulets (2278, -0.42 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.1 attack_power points (0.50 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Imperial Cloak (6432, -0.10 DPS) [world_drop]; Parachute Cloak (10518, -0.10 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (66.3 DPS) | yes | Raptorbane Armor (3566, -0.06 DPS) [quest]; Nightscape Tunic (8175, -0.10 DPS) [crafted]; Quillward Harness (10583, -1.42 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.60 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 29.3 attack_power points (1.45 DPS) | yes | Heavy Earthen Gloves (7359, -0.86 DPS, sim-verified) [crafted]; Skulker's Leather Gloves (252525, -0.95 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.95 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.49 DPS) | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.36 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest]; Petrolspill Leggings (9509, -0.59 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.2 attack_power points (0.65 DPS) | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.39 DPS, sim-verified) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Insurgent's Band (272066, -0.40 DPS) [vendor]; Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.54 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Insurgent's Band (272066, -0.48 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.54 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.50 DPS) | yes | Black Menace (6831, -2.93 DPS) [quest]; Coldrage Dagger (10761, -2.93 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272085, -3.44 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 432.8 attack_power points (21.44 DPS) | yes | Black Menace (6831, -3.40 DPS, sim-verified) [quest]; Satyr's Rod (15962, -21.39 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Swiftwind (13038, -0.34 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.84 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 429, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 50 (night-elf, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 96.8. Weights run: 1.0s. Verify run: 1.3s. 557 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.218 ± 0.061, crit=6.499 ± 0.352, hit=0.529 ± 0.085, melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 112.3 attack_power points (6.08 DPS) | yes | Eye of Theradras (17715, -1.85 DPS, sim-verified) [dungeon]; Ebon Mask (19984, -3.78 DPS) [quest]; Embrace of the Lycan (9479, -4.35 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (92.9 DPS) | yes | Sentinel's Medallion (19539, -0.07 DPS) [rep]; Ghostshard Talisman (7731, -0.10 DPS) [dungeon]; Zealous Shadowshard Pendant (17772, -1.58 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 104.3 attack_power points (5.64 DPS) | yes | Sunburn Spaulders (274751, -0.39 DPS, sim-verified) [vendor]; Phytoskin Spaulders (17749, -4.59 DPS) [dungeon]; Forest Tracker Epaulets (2278, -4.92 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (92.8 DPS) | yes | Blackveil Cape (11626, -0.07 DPS) [dungeon]; Duskbat Drape (19982, -0.07 DPS) [quest]; Blisterbane Wrap (12552, -1.44 DPS, sim-verified) [dungeon] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 114.3 attack_power points (6.19 DPS) | yes | Fungus Shroud Armor (17742, -2.82 DPS, sim-verified) [dungeon]; Blazewind Breastplate (11193, -4.67 DPS) [quest]; Quillward Harness (10583, -4.93 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.08 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.42 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 111.0 attack_power points (6.01 DPS) | yes | Sergeant Major's Leather Gauntlets (220856, -0.39 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.08 DPS) [crafted]; Shadowskin Gloves (18238, -1.08 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 111.0 attack_power points (6.01 DPS) | yes | Highlander's Lizardhide Girdle (20103, -1.08 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.30 DPS, sim-verified) [rep]; Girdle of Beastial Fury (11686, -4.38 DPS) [dungeon] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (92.4 DPS) | yes | Stormshroud Pants (15057, -1.06 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.78 DPS) [dungeon]; Basilisk Hide Pants (1718, -4.80 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.4 attack_power points (1.32 DPS) | yes | Sandstalker Ankleguards (12470, -0.20 DPS) [dungeon]; Sergeant Major's Leather Boots (220860, -0.34 DPS) [vendor]; Whisperwalk Boots (20255, -1.81 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 25.3 attack_power points (1.37 DPS) | yes | Masons Fraternity Ring (9533, -0.45 DPS) [quest]; Insurgent's Band (272065, -0.56 DPS) [vendor]; Mark of Kern (2262, -2.11 DPS, sim-verified) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | sim-verified (92.7 DPS) | yes | Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Insurgent's Band (272065, -0.27 DPS) [vendor]; Mark of Kern (2262, -1.30 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowblade (2163) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -2.27 DPS) [quest]; Charstone Dirk (17710, -2.27 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (27.69 DPS) | yes | Thermotastic Egg Timer (9644, -27.50 DPS) [quest]; Satyr's Rod (15962, -27.63 DPS) [world_drop]; Barman Shanker (12791, -46.86 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, -0.17 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.28 DPS, sim-verified) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; trinket1: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 557, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 60 (night-elf, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 164.7. Weights run: 1.0s. Verify run: 1.2s. 1223 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=8.118 ± 0.504, hit=not significant (0.458 ± 0.134), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ragefury Eyepatch (11735) (or Bloodvine Lens (19998)) | Blackrock Depths: Guzzler [dungeon] | 227.3 attack_power points (12.08 DPS) | yes | Bloodvine Lens (19998, -0.47 DPS, sim-verified) [crafted]; Outlaw's Collar (279253, -1.70 DPS) [crafted]; Duskwraith Helmet (239560, -2.48 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (164.7 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Mark of Fordring (15411, -5.50 DPS) [quest]; Imperial Jewel (11933, -5.62 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 145.5 attack_power points (7.74 DPS) | yes | Lieutenant Commander's Leather Shoulders (23313, -0.28 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.28 DPS) [vendor]; Knight-Lieutenant's Leather Shoulders (220852, -1.79 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 113.7 attack_power points (6.04 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Puissant Cape (18541, -3.89 DPS) [world]; Cloak of the Honor Guard (20073, -3.92 DPS) [rep] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-verified (164.7 DPS) | yes | Duskwraith Breastplate (239562, -2.72 DPS) [vendor]; Tunic of Undead Slaying (23089, -2.75 DPS, sim-verified) [world]; Dawn Armor (252483, -3.48 DPS) [crafted] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (164.7 DPS) | yes | Wristwraps of Undead Slaying (23093, +0.00 DPS, sim-verified) [world]; Duskwraith Bracers (239555, -0.09 DPS) [vendor]; Dragonspur Wraps (20615, -0.20 DPS) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 159.6 attack_power points (8.48 DPS) | yes | Marshal's Leather Handgrips (16454, -1.18 DPS) [vendor]; Marshal's Leather Handgrips (231544, -1.18 DPS) [vendor]; Devilsaur Gauntlets (15063, -4.34 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 157.2 attack_power points (8.36 DPS) | yes | Highlander's Leather Girdle (20115, -1.25 DPS) [rep]; Belt of the Archmage (18405, -2.31 DPS) [crafted]; Highlander's Leather Girdle (20045, -3.45 DPS, sim-verified) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 259.2 attack_power points (13.78 DPS) | yes | Knight-Captain's Leather Legguards (16419, -1.69 DPS) [pvp]; Sentinel's Silk Leggings (237815, -1.69 DPS) [vendor]; Stormshroud Pants (15057, -2.26 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 159.6 attack_power points (8.48 DPS) | yes | Darkmantle Footpads (226831, -6.80 DPS) [vendor]; Highlander's Leather Boots (20052, -6.88 DPS) [rep]; Pads of the Dread Wolf (13210, -8.00 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (164.7 DPS) | yes | Band of the Penitent (13217, -1.09 DPS) [quest]; Ring of Entropy (18543, -1.09 DPS) [world]; Naglering (11669, -3.66 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (164.7 DPS) | yes | Band of the Penitent (13217, -0.24 DPS) [quest]; Ring of Entropy (18543, -0.24 DPS) [world]; Naglering (11669, -2.89 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (164.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Earthstrike (21180) | Champion's Battlegear [quest] | sim-verified (164.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.01 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (164.7 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [vendor]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; The Lobotomizer (19324, -9.33 DPS, sim-verified) [rep] |
| off_hand | Emerald Dragonfang (20578) | Ysondre [world] | 749.2 attack_power points (39.82 DPS) | yes | Core Hound Tooth (18805, +0.00 DPS, sim-verified) [world_drop]; Distracting Dagger (18392, -9.48 DPS) [dungeon]; Tome of Knowledge (13385, -39.32 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (164.7 DPS) | yes | Bow of Searing Arrows (2825, -1.36 DPS, sim-verified) [world_drop]; Polished Ironwood Crossbow (20599, -4.77 DPS) [world]; Riphook (12653, -4.87 DPS) [dungeon] |

**New at 60:** head: Ragefury Eyepatch; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Stormshroud Armor; wrist: Bracers of the Eclipse; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Sentinel's Leather Pants; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Earthstrike; main_hand: Shadowsong's Sorrow; off_hand: Emerald Dragonfang; ranged: The Purifier

No-known-source sample (15 of 1223, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.4. Weights run: 0.9s. Verify run: 1.2s. 173 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.290 ± 0.038, hit=0.591 ± 0.033, melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.37 DPS) | yes | Flying Tiger Goggles (4368, -0.44 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 attack_power points (0.28 DPS) | yes | Erudite's Amulet (277204, -0.11 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.23 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.27 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.09 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 attack_power points (0.32 DPS) | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.27 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.23 DPS) | yes | Wolf Bracers (4794, -0.05 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.09 DPS) [crafted]; Forest Leather Gloves (3058, -0.09 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.82 DPS) | yes | Deviate Scale Belt (6468, -0.54 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.63 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 attack_power points (0.41 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 attack_power points (0.37 DPS) | yes | Footpads of the Fang (10411, -0.09 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.27 DPS, sim-verified) [dungeon] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 6.4 attack_power points (0.29 DPS) | yes | Legionnaire's Band (20429, -0.11 DPS) [rep]; Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.28 DPS) | yes | Legionnaire's Band (20429, +0.00 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (10.35 DPS) | yes | Assassin's Blade (1935, -0.32 DPS) [dungeon]; Evocator's Blade (2567, -0.65 DPS) [dungeon]; Buzzer Blade (2169, -1.48 DPS) [dungeon] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | 222.7 attack_power points (10.09 DPS) | yes | Assassin's Blade (1935, -0.15 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, -0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Pyrewood Signet Ring; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Edward's Knife; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 173, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6255 Fishing Pole (JEFFTEST); 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (troll, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 34.7. Weights run: 0.9s. Verify run: 1.2s. 304 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.421 ± 0.047, hit=0.238 ± 0.024, melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.1 attack_power points (0.47 DPS) | yes | Brawler's Leather Hood (252504, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.10 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.65 DPS) | yes | Scout's Medallion (19537, -0.31 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.37 DPS) [rep]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 attack_power points (0.52 DPS) | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.31 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.46 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 attack_power points (0.66 DPS) | yes | Panther Armor (6670, -0.16 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.28 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.28 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.46 DPS) | yes | Jurassic Wristguards (6198, -0.18 DPS) [world]; Insignia Bracers (6410, -0.18 DPS) [world_drop]; Unearthed Bands (9428, -0.52 DPS, sim-verified) [dungeon] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.74 DPS) | yes | Pilferer's Gloves (7358, -0.41 DPS, sim-verified) [crafted]; Braced Handguards (6784, -0.42 DPS) [quest]; Wolfclaw Gloves (1978, -0.46 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.12 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.37 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.21 DPS) | yes | Troll's Bane Leggings (13114, -0.55 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.60 DPS) [crafted]; Petrolspill Leggings (9509, -0.62 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.56 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Vorrel's Boots (7751, -0.18 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.1 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.09 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.10 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (14.97 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.02 DPS) [vendor]; Torturing Poker (7682, -1.18 DPS) [dungeon]; Thornspike (6681, -1.57 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (14.88 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -0.85 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -14.83 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Silver Star (3463, -0.21 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.23 DPS) [world_drop]; Crystalpine Stinger (13037, -0.23 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 304, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 66.1. Weights run: 1.0s. Verify run: 1.3s. 433 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.663 ± 0.081, hit=0.583 ± 0.056, melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.2 attack_power points (0.65 DPS) | yes | Nightscape Headband (8176, -0.08 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (64.6 DPS) | yes | Scout's Medallion (19536, -0.14 DPS) [rep]; Scout's Medallion (19537, -0.29 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.83 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.15 DPS) | yes | Forest Tracker Epaulets (2278, -0.41 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.1 attack_power points (0.50 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Wildhunter Cloak (16658, -0.01 DPS) [quest]; Imperial Cloak (6432, -0.10 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (65.2 DPS) | yes | Nightscape Tunic (8175, -0.10 DPS) [crafted]; Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Quillward Harness (10583, -1.44 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.59 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 29.3 attack_power points (1.45 DPS) | yes | Heavy Earthen Gloves (7359, -0.78 DPS, sim-verified) [crafted]; Skulker's Leather Gloves (252525, -0.95 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.95 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.49 DPS) | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.35 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest]; Petrolspill Leggings (9509, -0.59 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.2 attack_power points (0.65 DPS) | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.39 DPS, sim-verified) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Insurgent's Band (272066, -0.40 DPS) [vendor]; Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.54 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Insurgent's Band (272066, -0.47 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.54 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.50 DPS) | yes | Coldrage Dagger (10761, -2.93 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272085, -3.44 DPS) [vendor]; Hypnotic Blade (7714, -4.92 DPS) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 432.8 attack_power points (21.44 DPS) | yes | Coldrage Dagger (10761, -2.98 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -21.39 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Swiftwind (13038, -0.34 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.83 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 433, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 99.2. Weights run: 1.0s. Verify run: 1.4s. 562 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.218 ± 0.061, crit=6.499 ± 0.352, hit=0.529 ± 0.085, melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Leather Headband (220851) | Lady Palanseer [vendor] | 112.3 attack_power points (6.08 DPS) | yes | Eye of Theradras (17715, -1.98 DPS, sim-verified) [dungeon]; Ebon Mask (19984, -3.78 DPS) [quest]; Embrace of the Lycan (9479, -4.35 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (96.2 DPS) | yes | Scout's Medallion (19535, -0.07 DPS) [rep]; Ghostshard Talisman (7731, -0.10 DPS) [dungeon]; Zealous Shadowshard Pendant (17772, -1.57 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Leather Shoulders (220853) | Lady Palanseer [vendor] | 104.3 attack_power points (5.64 DPS) | yes | Sunburn Spaulders (274751, -0.59 DPS, sim-verified) [vendor]; Phytoskin Spaulders (17749, -4.59 DPS) [dungeon]; Forest Tracker Epaulets (2278, -4.92 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (96.0 DPS) | yes | Blackveil Cape (11626, -0.07 DPS) [dungeon]; Duskbat Drape (19982, -0.07 DPS) [quest]; Blisterbane Wrap (12552, -1.44 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 114.3 attack_power points (6.19 DPS) | yes | Fungus Shroud Armor (17742, -3.00 DPS, sim-verified) [dungeon]; Blazewind Breastplate (11193, -4.67 DPS) [quest]; Quillward Harness (10583, -4.93 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.08 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.42 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 111.0 attack_power points (6.01 DPS) | yes | First Sergeant's Leather Gauntlets (220857, -0.39 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.08 DPS) [crafted]; Shadowskin Gloves (18238, -1.08 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 111.0 attack_power points (6.01 DPS) | yes | Defiler's Lizardhide Girdle (20174, -1.08 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.30 DPS, sim-verified) [rep]; Girdle of Beastial Fury (11686, -4.38 DPS) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | sim-verified (96.3 DPS) | yes | Stormshroud Pants (15057, -1.66 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.78 DPS) [dungeon]; Basilisk Hide Pants (1718, -4.80 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.4 attack_power points (1.32 DPS) | yes | Sandstalker Ankleguards (12470, -0.20 DPS) [dungeon]; First Sergeant's Leather Boots (220861, -0.34 DPS) [vendor]; Whisperwalk Boots (20255, -1.92 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 25.3 attack_power points (1.37 DPS) | yes | Mark of Kern (2262, -0.29 DPS) [dungeon]; Assault Band (13095, -0.29 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.45 DPS) [quest] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.30 DPS) | yes | Assault Band (13095, -0.22 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.38 DPS) [quest]; Mark of Kern (2262, -1.56 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Smoking Heart of the Mountain (11811, -0.93 DPS, sim-verified) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -2.27 DPS) [quest]; Charstone Dirk (17710, -2.27 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (27.69 DPS) | yes | Thermotastic Egg Timer (9644, -27.50 DPS) [quest]; Satyr's Rod (15962, -27.63 DPS) [world_drop]; Barman Shanker (12791, -45.54 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, -0.17 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.38 DPS, sim-verified) [world_drop] |

**New at 50:** head: Blood Guard's Leather Headband; neck: Skibi's Pendant; shoulder: Blood Guard's Leather Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 562, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 161.3. Weights run: 1.0s. Verify run: 1.3s. 1228 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=8.118 ± 0.504, hit=not significant (0.458 ± 0.134), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ragefury Eyepatch (11735) (or Bloodvine Lens (19998)) | Blackrock Depths: Guzzler [dungeon] | 227.3 attack_power points (12.08 DPS) | yes | Bloodvine Lens (19998, -0.47 DPS, sim-verified) [crafted]; Outlaw's Collar (279253, -1.70 DPS) [crafted]; Duskwraith Helmet (239560, -2.48 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (161.3 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Mark of Fordring (15411, -5.50 DPS) [quest]; Imperial Jewel (11933, -5.62 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 145.5 attack_power points (7.74 DPS) | yes | Champion's Leather Shoulders (227056, -0.28 DPS) [vendor]; Champion's Leather Shoulders (23258, -0.28 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, -1.45 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 113.7 attack_power points (6.04 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Puissant Cape (18541, -3.89 DPS) [world]; Deathguard's Cloak (20068, -3.92 DPS) [rep] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-verified (161.3 DPS) | yes | Tunic of Undead Slaying (23089, -2.59 DPS, sim-verified) [world]; Duskwraith Breastplate (239562, -2.72 DPS) [vendor]; Dawn Armor (252483, -3.48 DPS) [crafted] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (161.3 DPS) | yes | Wristwraps of Undead Slaying (23093, +0.00 DPS, sim-verified) [world]; Duskwraith Bracers (239555, -0.09 DPS) [vendor]; Dragonspur Wraps (20615, -0.20 DPS) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 159.6 attack_power points (8.48 DPS) | yes | General's Leather Mitts (16560, -1.18 DPS) [vendor]; General's Leather Mitts (231555, -1.18 DPS) [vendor]; Devilsaur Gauntlets (15063, -4.43 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 157.2 attack_power points (8.36 DPS) | yes | Defiler's Leather Girdle (20193, -1.25 DPS) [rep]; Belt of the Archmage (18405, -2.31 DPS) [crafted]; Defiler's Leather Girdle (20190, -3.50 DPS, sim-verified) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 259.2 attack_power points (13.78 DPS) | yes | Legionnaire's Leather Leggings (16508, -1.69 DPS) [pvp]; Sentinel's Silk Leggings (237815, -1.69 DPS) [vendor]; Stormshroud Pants (15057, -2.33 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 159.6 attack_power points (8.48 DPS) | yes | Darkmantle Footpads (226831, -6.80 DPS) [vendor]; Defiler's Leather Boots (20186, -6.88 DPS) [rep]; Pads of the Dread Wolf (13210, -8.15 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (161.3 DPS) | yes | Band of the Penitent (13217, -1.09 DPS) [quest]; Ring of Entropy (18543, -1.09 DPS) [world]; Naglering (11669, -3.73 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (161.3 DPS) | yes | Band of the Penitent (13217, -0.24 DPS) [quest]; Ring of Entropy (18543, -0.24 DPS) [world]; Naglering (11669, -2.96 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (161.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (161.3 DPS) | yes | Earthstrike (21180, +0.00 DPS, sim-verified) [quest]; Counterattack Lodestone (18537, -1.23 DPS) [dungeon]; Hand of Justice (11815, -1.34 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (161.3 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [vendor]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; The Lobotomizer (19324, -9.76 DPS, sim-verified) [rep] |
| off_hand | Emerald Dragonfang (20578) | Ysondre [world] | 749.2 attack_power points (39.82 DPS) | yes | Core Hound Tooth (18805, +0.00 DPS, sim-verified) [world_drop]; Distracting Dagger (18392, -9.48 DPS) [dungeon]; Tome of Knowledge (13385, -39.32 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (161.3 DPS) | yes | Bow of Searing Arrows (2825, -1.15 DPS, sim-verified) [world_drop]; Polished Ironwood Crossbow (20599, -4.77 DPS) [world]; Riphook (12653, -4.87 DPS) [dungeon] |

**New at 60:** head: Ragefury Eyepatch; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Stormshroud Armor; wrist: Bracers of the Eclipse; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Sentinel's Leather Pants; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Emerald Dragonfang; ranged: The Purifier

No-known-source sample (15 of 1228, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

