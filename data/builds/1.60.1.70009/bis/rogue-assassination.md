# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.6. Weights run: 0.8s. Verify run: 1.1s. 169 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.067 ± 0.020, crit=2.221 ± 0.111, hit=0.606 ± 0.034, melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 attack_power points (0.40 DPS) | yes | Flying Tiger Goggles (4368, -0.66 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.4 attack_power points (0.30 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 attack_power points (0.25 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.42 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.09 DPS, sim-verified) [quest]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.7 attack_power points (0.55 DPS) | yes | Brawler's Leather Armor (252490, -0.01 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.20 DPS) [crafted]; Dark Leather Tunic (2317, -0.25 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 attack_power points (0.25 DPS) | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.09 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 31.1 attack_power points (1.45 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -1.15 DPS) [dungeon]; Forest Leather Gloves (3058, -1.25 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.84 DPS) | yes | Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.64 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.6 attack_power points (0.45 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.5 attack_power points (0.40 DPS) | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.15 DPS) [crafted]; Blackened Defias Boots (10402, -0.37 DPS, sim-verified) [dungeon] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 6.7 attack_power points (0.31 DPS) | yes | Protector's Band (20439, -0.11 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The 1 Ring (8350, -0.26 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 attack_power points (0.30 DPS) | yes | Protector's Band (20439, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.61 DPS) | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Barrens Basher (274744, -1.20 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 attack_power points (10.71 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.3 attack_power points (0.20 DPS) | yes | Fine Longbow (11304, -0.07 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Pyrewood Signet Ring; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 169, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6255 Fishing Pole (JEFFTEST); 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 47.7. Weights run: 0.9s. Verify run: 1.1s. 300 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.016, crit=3.306 ± 0.128, hit=0.769 ± 0.040, melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 attack_power points (0.52 DPS) | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.16 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.66 DPS) | yes | Sentinel's Medallion (19541, -0.20 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.35 DPS) [rep]; Kaleidoscope Chain (13084, -0.45 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 attack_power points (0.57 DPS) | yes | Dark Leather Shoulders (4252, -0.21 DPS) [crafted]; Insignia Mantle (4721, -0.21 DPS) [world_drop]; Mantle of Thieves (2264, -0.34 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.47 DPS) | yes | Tigerstrike Mantle (13108, +0.00 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.11 DPS) [world_drop]; Cloak of Night (4447, -0.16 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.75 DPS) | yes | Dusky Leather Armor (7374, +0.00 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.19 DPS) [quest]; Green Leather Armor (4255, -0.34 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.47 DPS) | yes | Jurassic Wristguards (6198, -0.16 DPS) [world]; Insignia Bracers (6410, -0.16 DPS) [world_drop]; Unearthed Bands (9428, -0.55 DPS, sim-verified) [dungeon] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | sim-verified (38.1 DPS) | yes | Pilferer's Gloves (7358, -0.34 DPS) [crafted]; Wolfclaw Gloves (1978, -0.44 DPS) [dungeon]; Fletcher's Gloves (7348, -0.49 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.13 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.67 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.23 DPS) | yes | Petrolspill Leggings (9509, -0.43 DPS, sim-verified) [dungeon]; Troll's Bane Leggings (13114, -0.50 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.55 DPS) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.15 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.15 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 attack_power points (0.47 DPS) | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Pyrewood Signet Ring (277210, -0.11 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, +0.00 DPS, sim-verified) [quest]; Pyrewood Signet Ring (277210, -0.07 DPS) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.21 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -0.12 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.20 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | sim-verified (47.2 DPS) | yes | Swinetusk Shank (6691, -9.65 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.08 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Silver Star (3463, -0.14 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop]; Crystalpine Stinger (13037, -0.22 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Royal Diplomatic Scepter; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (night-elf, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 36.3. Weights run: 0.9s. Verify run: 1.2s. 429 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.105 ± 0.014, crit=2.769 ± 0.073, hit=0.555 ± 0.022, melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 14.4 attack_power points (0.48 DPS) | yes | Nightscape Headband (8176, -0.05 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.07 DPS) [crafted]; Hawkeye's Helm (14591, -0.07 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (35.5 DPS) | yes | Sentinel's Medallion (19540, -0.06 DPS) [rep]; Sentinel's Medallion (19541, -0.17 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.48 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.81 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.40 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 11.0 attack_power points (0.37 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Imperial Cloak (6432, -0.07 DPS) [world_drop]; Parachute Cloak (10518, -0.07 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (35.8 DPS) | yes | Nightscape Tunic (8175, -0.07 DPS) [crafted]; Raptorbane Armor (3566, -0.09 DPS) [quest]; Quillward Harness (10583, -0.80 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Cultist's Armguards (270032, -0.35 DPS, sim-verified) [quest]; Imperial Leather Bracers (4061, -0.37 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 58.8 attack_power points (1.97 DPS) | yes | Shadowskin Gloves (18238, -0.67 DPS) [crafted]; Fletcher's Gloves (7348, -0.69 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -1.43 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.00 DPS) | yes | Highlander's Leather Girdle (20117, -0.20 DPS) [rep]; Highlander's Chain Girdle (20090, -0.21 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.87 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.20 DPS) [quest]; Petrolspill Leggings (9509, -0.35 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.4 attack_power points (0.48 DPS) | yes | Imperial Leather Boots (6431, +0.00 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.07 DPS) [crafted]; Worn Running Boots (9398, -0.07 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Insurgent's Band (272066, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Insurgent's Band (272066, -0.28 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (15.89 DPS) | yes | Vanquisher's Sword (10823, -1.09 DPS) [quest]; Dazzling Longsword (869, -1.13 DPS) [world_drop]; Southsea Lamp (9359, -1.29 DPS) [world_drop] |
| off_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 attack_power points (15.41 DPS) | yes | Vanquisher's Sword (10823, -0.89 DPS, sim-verified) [quest]; Satyr's Rod (15962, -15.37 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.17 DPS) [vendor]; Swiftwind (13038, -0.21 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.48 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Ardent Custodian; ranged: The Silencer

No-known-source sample (15 of 429, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 50 (night-elf, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 55.1. Weights run: 1.0s. Verify run: 1.2s. 557 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.107 ± 0.016, crit=3.272 ± 0.082, hit=0.678 ± 0.027, melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 68.6 attack_power points (2.31 DPS) | yes | Ebon Mask (19984, -0.99 DPS) [quest]; Embrace of the Lycan (9479, -1.23 DPS) [dungeon]; Eye of Theradras (17715, -1.40 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (52.6 DPS) | yes | Ghostshard Talisman (7731, -0.01 DPS) [dungeon]; Sentinel's Medallion (19539, -0.04 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.90 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 60.6 attack_power points (2.04 DPS) | yes | Sunburn Spaulders (274751, -0.53 DPS, sim-verified) [vendor]; Phytoskin Spaulders (17749, -1.45 DPS) [dungeon]; Forest Tracker Epaulets (2278, -1.63 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (52.5 DPS) | yes | Blackveil Cape (11626, -0.04 DPS) [dungeon]; Duskbat Drape (19982, -0.04 DPS) [quest]; Blisterbane Wrap (12552, -0.82 DPS, sim-verified) [dungeon] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 70.6 attack_power points (2.38 DPS) | yes | Blazewind Breastplate (11193, -1.52 DPS) [quest]; Quillward Harness (10583, -1.67 DPS) [dungeon]; Fungus Shroud Armor (17742, -1.96 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.26 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.8 attack_power points (2.22 DPS) | yes | Sergeant Major's Leather Gauntlets (220856, -0.23 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -0.67 DPS) [crafted]; Shadowskin Gloves (18238, -0.67 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 65.8 attack_power points (2.22 DPS) | yes | Highlander's Lizardhide Girdle (20103, -0.67 DPS) [rep]; Highlander's Cloth Girdle (20097, -0.78 DPS, sim-verified) [rep]; Girdle of Beastial Fury (11686, -1.21 DPS) [dungeon] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (52.5 DPS) | yes | Stormshroud Pants (15057, -0.81 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -1.50 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.60 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 attack_power points (0.75 DPS) | yes | Sandstalker Ankleguards (12470, -0.11 DPS) [dungeon]; Sergeant Major's Leather Boots (220860, -0.14 DPS) [vendor]; Whisperwalk Boots (20255, -1.09 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 26.8 attack_power points (0.90 DPS) | yes | Masons Fraternity Ring (9533, -0.38 DPS) [quest]; Insurgent's Band (272065, -0.40 DPS) [vendor]; Mark of Kern (2262, -1.56 DPS, sim-verified) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | sim-verified (52.4 DPS) | yes | Masons Fraternity Ring (9533, -0.15 DPS) [quest]; Insurgent's Band (272065, -0.17 DPS) [vendor]; Mark of Kern (2262, -0.78 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Inventor's Focal Sword (17719, +0.00 DPS) [dungeon]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hammer of the Northern Wind (810, -1.68 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (18.39 DPS) | yes | Claw of Celebras (17738, -2.12 DPS) [dungeon]; Inventor's Focal Sword (17719, -12.43 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -18.28 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, -0.05 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.05 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.77 DPS, sim-verified) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; trinket1: Frozen Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 557, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 60 (night-elf, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 250.8. Weights run: 1.0s. Verify run: 1.4s. 1223 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=0.863 ± 0.035, melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 130.4 attack_power points (4.32 DPS) | yes | Bloodvine Lens (19998, -0.49 DPS) [crafted]; Outlaw's Collar (279253, -0.66 DPS) [crafted]; Ragefury Eyepatch (11735, -14.12 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 attack_power points (0.00 DPS) | yes | Blazefury Medallion (17111, -1.16 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -1.63 DPS) [quest]; Imperial Jewel (11933, -1.65 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 88.3 attack_power points (2.93 DPS) | yes | Lieutenant Commander's Leather Shoulders (23313, +0.00 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [vendor]; Knight-Lieutenant's Leather Shoulders (220852, -2.01 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 57.7 attack_power points (1.91 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Puissant Cape (18541, -0.56 DPS) [world]; Cloak of the Honor Guard (20073, -0.60 DPS) [rep] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Stormshroud Armor (15056, -0.21 DPS) [crafted]; Dawn Armor (252483, -0.45 DPS) [crafted]; Tunic of Undead Slaying (23089, -15.14 DPS, sim-verified) [world] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Bracers of the Eclipse (18375, -0.05 DPS) [dungeon]; Dragonspur Wraps (20615, -0.16 DPS) [world]; Wristwraps of Undead Slaying (23093, -7.93 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 105.9 attack_power points (3.51 DPS) | yes | Marshal's Leather Handgrips (16454, -0.85 DPS) [vendor]; Marshal's Leather Handgrips (231544, -0.85 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.71 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 103.7 attack_power points (3.44 DPS) | yes | Highlander's Leather Girdle (20115, -0.86 DPS) [rep]; Belt of the Archmage (18405, -1.52 DPS) [crafted]; Highlander's Leather Girdle (20045, -6.67 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | sim-verified (250.8 DPS) | yes | Stormshroud Pants (15057, -0.53 DPS) [crafted]; Knight-Captain's Leather Legguards (16419, -0.53 DPS) [pvp]; Sentinel's Leather Pants (237818, -8.26 DPS, sim-verified) [vendor] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 105.9 attack_power points (3.51 DPS) | yes | Darkmantle Footpads (226831, -2.41 DPS) [vendor]; Duskwraith Treads (239553, -2.51 DPS) [vendor]; Pads of the Dread Wolf (13210, -11.95 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Penitent (13217, -0.82 DPS) [quest]; Ring of Entropy (18543, -0.82 DPS) [world]; Naglering (11669, -6.88 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Penitent (13217, -0.29 DPS) [quest]; Ring of Entropy (18543, -0.29 DPS) [world]; Naglering (11669, -6.06 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Earthstrike (21180) | Champion's Battlegear [quest] | 0.0 attack_power points (0.00 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Heroism (19287, -6.94 DPS, sim-verified) [quest] |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Misplaced Servo Arm (23221, -136.80 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 attack_power points (0.00 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -147.81 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 attack_power points (0.00 DPS) | yes | Polished Ironwood Crossbow (20599, -1.12 DPS) [world]; Riphook (12653, -1.18 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.82 DPS, sim-verified) [world_drop] |

**New at 60:** head: Duskwraith Helmet; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Earthstrike; main_hand: Ebon Hand; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1223, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

## Horde

### Band 20 (troll, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 34.9. Weights run: 0.8s. Verify run: 1.1s. 173 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.067 ± 0.020, crit=2.221 ± 0.111, hit=0.606 ± 0.034, melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 attack_power points (0.40 DPS) | yes | Flying Tiger Goggles (4368, -0.63 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.4 attack_power points (0.30 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 attack_power points (0.25 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.39 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.08 DPS, sim-verified) [quest]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.5 attack_power points (0.35 DPS) | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.33 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 attack_power points (0.25 DPS) | yes | Wolf Bracers (4794, -0.08 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 31.1 attack_power points (1.45 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -1.15 DPS) [dungeon]; Forest Leather Gloves (3058, -1.25 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.84 DPS) | yes | Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.64 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.6 attack_power points (0.45 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.5 attack_power points (0.40 DPS) | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.15 DPS) [crafted]; Blackened Defias Boots (10402, -0.35 DPS, sim-verified) [dungeon] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 6.7 attack_power points (0.31 DPS) | yes | Legionnaire's Band (20429, -0.11 DPS) [rep]; Bounty Hunter's Ring (5351, -0.16 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 attack_power points (0.30 DPS) | yes | Legionnaire's Band (20429, +0.00 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.61 DPS) | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Wingblade (6504, -1.12 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 attack_power points (10.71 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.3 attack_power points (0.20 DPS) | yes | Fine Longbow (11304, -0.05 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Pyrewood Signet Ring; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 173, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6255 Fishing Pole (JEFFTEST); 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (troll, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 47.3. Weights run: 0.9s. Verify run: 1.2s. 304 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.016, crit=3.306 ± 0.128, hit=0.769 ± 0.040, melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 attack_power points (0.52 DPS) | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.16 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.66 DPS) | yes | Scout's Medallion (19537, -0.20 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.35 DPS) [rep]; Kaleidoscope Chain (13084, -0.45 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 attack_power points (0.57 DPS) | yes | Dark Leather Shoulders (4252, -0.21 DPS) [crafted]; Insignia Mantle (4721, -0.21 DPS) [world_drop]; Mantle of Thieves (2264, -0.34 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.47 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Tigerstrike Mantle (13108, -0.06 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.11 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 15.4 attack_power points (0.72 DPS) | yes | Panther Armor (6670, -0.22 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.31 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.31 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.47 DPS) | yes | Jurassic Wristguards (6198, -0.16 DPS) [world]; Insignia Bracers (6410, -0.16 DPS) [world_drop]; Unearthed Bands (9428, -0.55 DPS, sim-verified) [dungeon] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | sim-verified (37.6 DPS) | yes | Pilferer's Gloves (7358, -0.34 DPS) [crafted]; Braced Handguards (6784, -0.39 DPS) [quest]; Fletcher's Gloves (7348, -0.47 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.13 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.36 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.23 DPS) | yes | Petrolspill Leggings (9509, -0.45 DPS, sim-verified) [dungeon]; Troll's Bane Leggings (13114, -0.50 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.55 DPS) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.15 DPS) [world_drop]; Vorrel's Boots (7751, -0.15 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 attack_power points (0.47 DPS) | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Pyrewood Signet Ring (277210, -0.11 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.00 DPS, sim-verified) [quest]; Pyrewood Signet Ring (277210, -0.07 DPS) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.21 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -0.12 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.20 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | sim-verified (46.9 DPS) | yes | Swinetusk Shank (6691, -9.73 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.08 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Silver Star (3463, -0.14 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop]; Crystalpine Stinger (13037, -0.22 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Royal Diplomatic Scepter; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 304, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 36.1. Weights run: 0.9s. Verify run: 1.2s. 433 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.105 ± 0.014, crit=2.769 ± 0.073, hit=0.555 ± 0.022, melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 14.4 attack_power points (0.48 DPS) | yes | Nightscape Headband (8176, -0.06 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.07 DPS) [crafted]; Hawkeye's Helm (14591, -0.07 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (35.3 DPS) | yes | Scout's Medallion (19536, -0.06 DPS) [rep]; Scout's Medallion (19537, -0.17 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.48 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.81 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.40 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 11.0 attack_power points (0.37 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Wildhunter Cloak (16658, -0.04 DPS) [quest]; Imperial Cloak (6432, -0.07 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (35.6 DPS) | yes | Nightscape Tunic (8175, -0.07 DPS) [crafted]; Dusky Leather Armor (7374, -0.11 DPS) [crafted]; Quillward Harness (10583, -0.80 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Cultist's Armguards (270032, -0.34 DPS, sim-verified) [quest]; Imperial Leather Bracers (4061, -0.37 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 58.8 attack_power points (1.97 DPS) | yes | Shadowskin Gloves (18238, -0.67 DPS) [crafted]; Fletcher's Gloves (7348, -0.68 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -1.43 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.00 DPS) | yes | Defiler's Leather Girdle (20191, -0.20 DPS) [rep]; Defiler's Chain Girdle (20152, -0.21 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.87 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.20 DPS) [quest]; Petrolspill Leggings (9509, -0.35 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.4 attack_power points (0.48 DPS) | yes | Imperial Leather Boots (6431, +0.00 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.07 DPS) [crafted]; Worn Running Boots (9398, -0.07 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Insurgent's Band (272066, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Insurgent's Band (272066, -0.27 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (15.89 DPS) | yes | Vanquisher's Sword (10823, -1.09 DPS) [quest]; Dazzling Longsword (869, -1.13 DPS) [world_drop]; Southsea Lamp (9359, -1.29 DPS) [world_drop] |
| off_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 attack_power points (15.41 DPS) | yes | Vanquisher's Sword (10823, -0.86 DPS, sim-verified) [quest]; Satyr's Rod (15962, -15.37 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.17 DPS) [vendor]; Swiftwind (13038, -0.21 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.48 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Ardent Custodian; ranged: The Silencer

No-known-source sample (15 of 433, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 57.3. Weights run: 1.0s. Verify run: 1.2s. 562 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.107 ± 0.016, crit=3.272 ± 0.082, hit=0.678 ± 0.027, melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Leather Headband (220851) | Lady Palanseer [vendor] | 68.6 attack_power points (2.31 DPS) | yes | Ebon Mask (19984, -0.99 DPS) [quest]; Embrace of the Lycan (9479, -1.23 DPS) [dungeon]; Eye of Theradras (17715, -1.24 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (55.5 DPS) | yes | Ghostshard Talisman (7731, -0.01 DPS) [dungeon]; Scout's Medallion (19535, -0.04 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.93 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Leather Shoulders (220853) | Lady Palanseer [vendor] | 60.6 attack_power points (2.04 DPS) | yes | Sunburn Spaulders (274751, -0.36 DPS, sim-verified) [vendor]; Phytoskin Spaulders (17749, -1.45 DPS) [dungeon]; Forest Tracker Epaulets (2278, -1.63 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (55.4 DPS) | yes | Blackveil Cape (11626, -0.04 DPS) [dungeon]; Duskbat Drape (19982, -0.04 DPS) [quest]; Blisterbane Wrap (12552, -0.85 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 70.6 attack_power points (2.38 DPS) | yes | Blazewind Breastplate (11193, -1.52 DPS) [quest]; Quillward Harness (10583, -1.67 DPS) [dungeon]; Fungus Shroud Armor (17742, -1.80 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.26 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.8 attack_power points (2.22 DPS) | yes | First Sergeant's Leather Gauntlets (220857, -0.23 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -0.67 DPS) [crafted]; Shadowskin Gloves (18238, -0.67 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 65.8 attack_power points (2.22 DPS) | yes | Defiler's Lizardhide Girdle (20174, -0.67 DPS) [rep]; Defiler's Cloth Girdle (20165, -0.78 DPS, sim-verified) [rep]; Girdle of Beastial Fury (11686, -1.21 DPS) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | sim-verified (55.5 DPS) | yes | Stormshroud Pants (15057, -0.95 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -1.50 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.60 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 attack_power points (0.75 DPS) | yes | Sandstalker Ankleguards (12470, -0.11 DPS) [dungeon]; First Sergeant's Leather Boots (220861, -0.14 DPS) [vendor]; Whisperwalk Boots (20255, -1.13 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 26.8 attack_power points (0.90 DPS) | yes | Mark of Kern (2262, -0.23 DPS) [dungeon]; Assault Band (13095, -0.23 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.38 DPS) [quest] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (0.81 DPS) | yes | Assault Band (13095, -0.13 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.29 DPS) [quest]; Mark of Kern (2262, -0.94 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Smoking Heart of the Mountain (11811, -0.65 DPS, sim-verified) [crafted] |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Inventor's Focal Sword (17719, +0.00 DPS) [dungeon]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hammer of the Northern Wind (810, -1.87 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (18.39 DPS) | yes | Claw of Celebras (17738, -2.12 DPS) [dungeon]; White Bone Shredder (11863, -3.45 DPS) [quest]; Inventor's Focal Sword (17719, -12.29 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, -0.05 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.05 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.79 DPS, sim-verified) [world_drop] |

**New at 50:** head: Blood Guard's Leather Headband; neck: Skibi's Pendant; shoulder: Blood Guard's Leather Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 562, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 249.8. Weights run: 1.0s. Verify run: 1.4s. 1228 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=0.863 ± 0.035, melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 130.4 attack_power points (4.32 DPS) | yes | Bloodvine Lens (19998, -0.49 DPS) [crafted]; Outlaw's Collar (279253, -0.66 DPS) [crafted]; Ragefury Eyepatch (11735, -15.55 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 attack_power points (0.00 DPS) | yes | Blazefury Medallion (17111, -0.71 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -1.63 DPS) [quest]; Imperial Jewel (11933, -1.65 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 88.3 attack_power points (2.93 DPS) | yes | Champion's Leather Shoulders (23258, +0.00 DPS) [vendor]; Champion's Leather Shoulders (227056, +0.00 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, -1.38 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 57.7 attack_power points (1.91 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Puissant Cape (18541, -0.56 DPS) [world]; Deathguard's Cloak (20068, -0.60 DPS) [rep] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Stormshroud Armor (15056, -0.21 DPS) [crafted]; Dawn Armor (252483, -0.45 DPS) [crafted]; Tunic of Undead Slaying (23089, -16.70 DPS, sim-verified) [world] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Bracers of the Eclipse (18375, -0.05 DPS) [dungeon]; Dragonspur Wraps (20615, -0.16 DPS) [world]; Wristwraps of Undead Slaying (23093, -9.48 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 105.9 attack_power points (3.51 DPS) | yes | General's Leather Mitts (16560, -0.85 DPS) [vendor]; General's Leather Mitts (231555, -0.85 DPS) [vendor]; Devilsaur Gauntlets (15063, -9.33 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 103.7 attack_power points (3.44 DPS) | yes | Defiler's Leather Girdle (20193, -0.86 DPS) [rep]; Belt of the Archmage (18405, -1.52 DPS) [crafted]; Defiler's Leather Girdle (20190, -8.33 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | sim-verified (249.8 DPS) | yes | Stormshroud Pants (15057, -0.53 DPS) [crafted]; Legionnaire's Leather Leggings (16508, -0.53 DPS) [pvp]; Sentinel's Leather Pants (237818, -9.48 DPS, sim-verified) [vendor] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 105.9 attack_power points (3.51 DPS) | yes | Darkmantle Footpads (226831, -2.41 DPS) [vendor]; Duskwraith Treads (239553, -2.51 DPS) [vendor]; Pads of the Dread Wolf (13210, -13.67 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Penitent (13217, -0.82 DPS) [quest]; Ring of Entropy (18543, -0.82 DPS) [world]; Naglering (11669, -8.45 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Penitent (13217, -0.29 DPS) [quest]; Ring of Entropy (18543, -0.29 DPS) [world]; Naglering (11669, -7.65 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Earthstrike (21180) | Champion's Battlegear [quest] | 0.0 attack_power points (0.00 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, -1.05 DPS, sim-verified) [quest] |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Misplaced Servo Arm (23221, -134.74 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 attack_power points (0.00 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -145.33 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 attack_power points (0.00 DPS) | yes | Polished Ironwood Crossbow (20599, -1.12 DPS) [world]; Riphook (12653, -1.18 DPS) [dungeon]; Bow of Searing Arrows (2825, -2.09 DPS, sim-verified) [world_drop] |

**New at 60:** head: Duskwraith Helmet; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Earthstrike; main_hand: Ebon Hand; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1228, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

