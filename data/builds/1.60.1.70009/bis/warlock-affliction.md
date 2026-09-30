# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.3. Weights run: 0.8s. Verify run: 0.8s. 128 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.102, intellect=not significant (-0.303 ± 0.109), crit=0.392 ± 0.018, hit=1.206 ± 0.062, spell_haste=not significant (-0.036 ± 0.114), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.102

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.96 DPS) [crafted]; Lucky Fishing Hat (19972, -0.96 DPS) [quest]; Flying Tiger Goggles (4368, -2.75 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) (or Magician's Mantle (12998)) | Tailoring [crafted] | 5.0 | yes | Magician's Mantle (12998, +0.00 DPS, sim-verified) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.16 DPS) [crafted]; Slime-encrusted Pads (6461, -0.80 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.16 DPS) [crafted]; Caretaker's Cape (20428, -0.16 DPS) [rep]; Feyscale Cloak (6632, -0.44 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.16 DPS) [crafted]; Bloody Apron (6226, -0.16 DPS) [dungeon]; Green Woolen Vest (2582, -1.76 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Silver-lined Bracers (3224, -0.16 DPS) [world]; Seer's Cuffs (3645, -0.16 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.43 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.48 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.80 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (38.3 DPS) | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.59 DPS, sim-verified) [crafted]; Lesser Belt of the Spire (1299, -0.64 DPS) [world] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted]; Colorful Kilt (10048, -0.64 DPS) [crafted]; Silk-threaded Trousers (1929, -1.31 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.48 DPS) [crafted]; Pristine Boots (253889, -0.64 DPS) [crafted]; Feather Padded Treads (285345, -1.06 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.32 DPS) [world]; Lavishly Jeweled Ring (1156, -0.80 DPS) [dungeon]; Ring of the Shadow (1462, -0.80 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.80 DPS) [dungeon]; Ring of the Shadow (1462, -0.80 DPS) [world]; Sludge-Stained Band (286535, -0.81 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Staff of Westfall (2042), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), Keen Skinning Knife (269716), Riverpaw Mystic Staff (1391), Defias Mage Staff (1928), Skyseeker's Greatstaff (263937), Oakthrush Staff (15397), Felweaver's Staff (284287), Apprentice's Spellstaff (248008), Makeshift Stormcaller (281274), Blackfang (2236), Staff of the Blessed Seer (2271), Nightbane Staff (3227), Deadly Bronze Poniard (3490), Shadowfang (1482), Assassin's Blade (1935), Bluegill Kukri (2046), Bronze Shortsword (2850), Decapitating Sword (3740), Cruel Barb (5191), Spikelash Dagger (6333), Edward's Knife (251485), Chol'aruk's Trophy Sword (276886), Plaguefang (279876), Staff of Horrors (880), Buzz Saw (1937), Pearl-handled Dagger (5540), Living Root (6631), Butcher's Slicer (6633), Thief's Blade (5192), Emberstone Staff (5201), Militant Shortsword (15211), Scimitar of Atun (1469), Gnarled Hermit's Staff (1539), Buzzer Blade (2169), Brackclaw (2235), Tail Spike (6448), Night Watch Shortsword (935), Thornblade (2908), Hook Dagger (3184), Big Bronze Knife (3848), Staff of Nobles (3902), Ironpatch Blade (12976), Venom Web Fang (899), Riverside Staff (1473), Blackwater Cutlass (1951), Solid Shortblade (2074), Scrimshaw Dagger (2089), Medicine Staff (4575), Goblin Screwdriver (1936), Hollowfang Blade (2020), Northern Shortsword (2078), Slicer Blade (820), Foamspittle Staff (1405), Daryl's Shortsword (3572), War Knife (4571), Ritual Blade (5112), Redridge Machete (1219), Defias Rapier (1925), Raider Shortsword (15210), Daggerfang's Daggerfang (281257), Giant Tarantula Fang (1287), Staff of Conjuring (1933), Long Crawler Limb (2088), Curved Dagger (2632), Enamelled Broadsword (4765), Pale Skinner (5744), Briarsteel Shortsword (15335), Curvewood Dagger (15396), Craftsman's Dagger (2218), Stonesplinter Dagger (2266), Sturdy Quarterstaff (4566), Feral Blade (4766), Balanced Fighting Stick (6215), Balanced Quarterstaff (257343), Quickblade's Dagger (257346), Gnarlpine War Staff (284166), Jeweled Dagger (1917), Small Green Dagger (4302), Militia Sword (248006), Militia Shortblade (248007), Tim's Lost Rib (282064), Bayne's Bite (283462), Pristine Orcish Dagger (286730), Ukta's Conduit (286754), Small Hand Blade (816), Carving Knife (2140), Bloodstained Knife (3225), Copper Dagger (7166), Hoof-Holstered Blade (277988), Helm Splitter (286755), Notched Shortsword (727), Copper Shortsword (2847), Gritroot Staff (9603), Brushwood Blade (18957), Webwood Slicer (282284), Chipped Spellstaff (285238), Explorer's Shortsword (285239)) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 | yes | Staff of Westfall (2042, +0.00 DPS) [quest]; Gnarled Necromancer's Staff (251534, +0.00 DPS) [quest]; Twisted Chanter's Staff (890, -0.45 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Pulsating Hydra Heart (5183, -0.80 DPS) [world]; Tear of Grief (5611, -0.80 DPS) [quest]; Grayson's Torch (1172, -1.35 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 142.1 | yes | Skycaller (12984, -1.81 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.50 DPS) [dungeon]; Deepblaze (279896, -4.26 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 128, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10047 Simple Kilt; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 57.3. Weights run: 0.8s. Verify run: 0.8s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.164, intellect=0.928 ± 0.177, crit=0.391 ± 0.020, hit=1.447 ± 0.093, spell_haste=0.972 ± 0.185, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.164

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.3 | yes | Holy Shroud (2721, -0.90 DPS) [world_drop]; Shadow Hood (4323, -1.06 DPS) [crafted]; Nightsky Cowl (4039, -1.57 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.6 | yes | Darkspear Warding Pendant (272075, -1.80 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.85 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.4 | yes | Fairywing Mantle (9536, -0.63 DPS) [quest]; Death Speaker Mantle (6685, -0.70 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.84 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.4 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.15 DPS) [quest]; Resilient Cape (14400, -0.39 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.1 | yes | Death Speaker Robes (6682, -0.70 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -1.58 DPS) [crafted]; Tree Bark Jacket (1486, -1.69 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.72 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.72 DPS) [quest]; Glowing Magical Bracelets (13106, -2.11 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | sim-verified (55.8 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.07 DPS) [dungeon]; Serpent Gloves (5970, -0.16 DPS) [dungeon]; Town Clerk's Mittens (270029, -1.32 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.8 | yes | Belt of Arugal (6392, -0.42 DPS) [dungeon]; Invoker's Cord (215366, -0.45 DPS) [crafted]; Crimson Silk Belt (7055, -1.27 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (56.0 DPS) | yes | Gaze Dreamer Pants (6903, -0.31 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Abomination Skin Leggings (23173, -1.52 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.5 | yes | Spidersilk Boots (4320, -0.58 DPS) [crafted]; Frothing Slippers (254003, -1.46 DPS) [crafted]; Acidic Walkers (9454, -2.17 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Black Widow Band (6199, -0.10 DPS) [world]; Snake Hoop (6750, -0.10 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.21 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 | yes | Snake Hoop (6750, -0.07 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.18 DPS) [vendor]; Black Widow Band (6199, -1.06 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (54.2 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, -0.32 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.2 | yes | Twisted Chanter's Staff (890, -0.01 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.19 DPS) [quest]; Channeler's Staff (4437, -0.58 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 161.8 | yes | Starfaller (13063, -0.75 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.79 DPS) [crafted]; Gravestone Scepter (7001, -4.84 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 110.0. Weights run: 0.7s. Verify run: 0.7s. 295 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.127, intellect=0.992 ± 0.157, crit=0.387 ± 0.020, hit=1.104 ± 0.078, spell_haste=1.008 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.127

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -1.42 DPS) [crafted]; Augural Shroud (2620, -1.62 DPS, sim-verified) [world]; Craftsman's Monocle (4393, -1.71 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.0 | yes | Darkspear Warding Pendant (272074, -1.68 DPS) [vendor]; Necklace of Calisea (1714, -1.79 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.24 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.9 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.55 DPS) [dungeon]; Berylline Pads (4197, -0.83 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.0 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Cloak of Rot (4462, -0.85 DPS) [world] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.0 | yes | Robe of Power (7054, -0.57 DPS) [crafted]; Dreamweave Vest (10021, -0.87 DPS, sim-verified) [crafted]; Crimson Silk Vest (7058, -1.69 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Windchaser Cuffs (14429, -0.02 DPS) [world_drop]; Spidertank Oilrag (9448, -0.15 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.30 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 | yes | Stormcloth Gloves (10011, -1.70 DPS) [crafted]; Red Mageweave Gloves (10018, -1.87 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.95 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (108.5 DPS) | yes | Gilded Cord (254037, -0.57 DPS) [crafted]; Highlander's Cloth Girdle (20099, -1.12 DPS) [rep]; Deathmage Sash (10771, -1.73 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.9 | yes | Crimson Silk Pantaloons (7062, -1.47 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -2.69 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.10 DPS) [dungeon]; Spidersilk Boots (4320, -3.65 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.0 | yes | Ring of Forlorn Spirits (2043, -2.23 DPS) [quest]; Reedknot Ring (9622, -2.51 DPS) [quest]; Minor Channeling Ring (1449, -2.51 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Ring of Forlorn Spirits (2043, -0.13 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.56 DPS) [quest]; Lorekeeper's Ring (19525, -0.56 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 22.0 | yes | Windweaver Staff (7757, -0.99 DPS, sim-verified) [dungeon]; Illusionary Rod (7713, -2.96 DPS) [dungeon]; Staff of Jordan (873, -3.09 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (108.6 DPS) | yes | Umbral Wand (5216, -0.01 DPS) [world_drop]; Burning Sliver (5249, -1.31 DPS) [quest]; Jaina's Firestarter (13064, -1.86 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Twisted Nether Wand

No-known-source sample (15 of 295, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 187.3. Weights run: 0.7s. Verify run: 0.8s. 376 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.227, intellect=not significant (0.199 ± 0.264), crit=0.457 ± 0.026, hit=2.450 ± 0.155, spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Red Mageweave Headband (10033, -1.13 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.18 DPS) [vendor]; Dreamweave Circlet (10041, -2.08 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.2 | yes | Mindburst Medallion (11196, -0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.52 DPS) [world_drop]; Darkspear Warding Pendant (272073, -1.80 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.6 | yes | Black Mageweave Shoulders (10027, -1.34 DPS) [crafted]; Bloodmage Mantle (7684, -1.62 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.72 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 | yes | Runecloth Cloak (13860, -1.33 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.74 DPS) [dungeon]; Icy Cloak (4327, -2.30 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 | yes | Acumen Robes (17775, -0.45 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.61 DPS) [world_drop]; Knight's Dreadweave Vest (220886, -0.73 DPS) [vendor] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 | yes | Sergeant Major's Dreadweave Gloves (220890, -1.12 DPS) [vendor]; Runecloth Gloves (13863, -1.40 DPS) [crafted]; Black Mageweave Gloves (10003, -1.50 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 16.4 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.45 DPS) [rep]; Ghostweave Cord (254073, -0.67 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | sim-verified (178.2 DPS) | yes | Red Mageweave Pants (10009, -0.73 DPS) [crafted]; Senior Designer's Pantaloons (11841, -1.57 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -2.29 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (180.5 DPS) | yes | Gilded Sandals (254107, -3.14 DPS) [crafted]; Black Mageweave Boots (10026, -3.25 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.63 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 24.5 | yes | Lorekeeper's Ring (19523, -3.50 DPS) [rep]; Philanthropist's Ring (281635, -3.73 DPS) [quest]; Lorekeeper's Ring (19524, -4.35 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Lorekeeper's Ring (19523, +0.00 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.51 DPS) [quest]; Lorekeeper's Ring (19524, -1.12 DPS) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (176.1 DPS) | yes | Frozen Heart of the Mountain (249469, -1.10 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -1.68 DPS) [quest]; Tidal Charm (1404, -1.68 DPS) [vendor] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (176.1 DPS) | yes | Frozen Heart of the Mountain (249469, -3.12 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -3.36 DPS) [quest]; Tidal Charm (1404, -3.36 DPS) [vendor] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | sim-verified (176.1 DPS) | yes | Staff of Dar'Orahil (15106, -0.28 DPS) [quest]; Shortsword of Vengeance (754, -1.43 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -5.30 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (180.6 DPS) | yes | Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted]; Pyric Caduceus (11748, -4.72 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Highlander's Cloth Girdle; legs: Wizardweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Abyss Shard; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 376, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 382.5. Weights run: 0.8s. Verify run: 0.8s. 777 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.837), intellect=not significant (-1.254 ± 0.970), crit=1.999 ± 0.106, hit=6.401 ± 0.516, spell_haste=not significant (0.241 ± 0.963), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.837)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 218.0 | yes | Deathmist Mask (226909, -19.72 DPS) [quest]; Deathmist Mask (22074, -20.01 DPS) [quest]; Bloodvine Goggles (19999, -27.12 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (366.1 DPS) | yes | Blazefury Medallion (17111, -1.58 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -5.23 DPS) [quest]; Orb of the Darkmoon (19426, -6.09 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 132.0 | yes | Rugged Mantle of the Timbermaw (227808, -6.35 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -11.89 DPS) [vendor]; Mantle of the Timbermaw (19050, -12.62 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 80.0 | yes | Howler's Furs (272414, -2.32 DPS) [vendor]; Stalwart Cloak (272415, -2.32 DPS) [vendor]; Earthweave Cloak (21187, -7.36 DPS, sim-verified) [quest] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | sim-verified (382.5 DPS) | yes | Heretic Garb (240146, -1.45 DPS) [vendor]; Earthpower Vest (21183, -9.86 DPS) [quest]; Bloodvine Vest (19682, -15.96 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 91.0 | yes | Heretic Bindings (240145, +0.00 DPS, sim-verified) [vendor]; Heretic Wristguards (240152, -1.89 DPS) [vendor]; Dryad's Wrist Bindings (19595, -10.01 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 101.0 | yes | Deathmist Wraps (226911, -3.48 DPS) [quest]; Gloves of Spell Mastery (14146, -5.23 DPS) [crafted]; Deathmist Wraps (22077, -10.25 DPS, sim-verified) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 160.0 | yes | Heretic Waistguard (240151, -9.98 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -10.29 DPS) [vendor]; Belt of the Archmage (18405, -16.24 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 175.0 | yes | Bloodvine Leggings (19683, -10.73 DPS) [crafted]; Sentinel's Silk Leggings (237815, -12.62 DPS) [vendor]; Heretic Pants (240149, -17.00 DPS, sim-verified) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 100.0 | yes | Sergeant Major's Dreadweave Boots (220891, -4.06 DPS) [vendor]; Argent Elite Boots (227816, -5.22 DPS) [vendor]; Bloodvine Boots (19684, -6.95 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (361.5 DPS) | yes | Don Julio's Band (19325, -0.00 DPS) [rep]; Band of Earthen Might (21182, -0.00 DPS) [quest]; Signet Ring of the Bronze Dragonflight (234028, -0.29 DPS) [vendor] |
| finger2 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (366.1 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Blessed Band of Light (272407, +0.00 DPS) [vendor]; Don Julio's Band (19325, -5.70 DPS, sim-verified) [rep] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (361.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.16 DPS) [world_drop]; Abyss Shard (20534, -6.18 DPS, sim-verified) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (361.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.32 DPS) [world_drop]; Abyss Shard (20534, -3.69 DPS, sim-verified) [quest] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (366.1 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -2.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 535.2 | yes | Wand of Biting Cold (19108, -0.33 DPS, sim-verified) [quest]; Stormrager (16997, -14.94 DPS) [quest]; Brilliant Wand (249385, -16.74 DPS) [crafted] |

**New at 60:** head: Heretic Cowl; neck: Beads of Ogre Might; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Rockfury Bracers; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Wrath of Cenarius; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Staff of Dar'Orahil; ranged: Torch of Light

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.6. Weights run: 0.8s. Verify run: 0.8s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.102, intellect=not significant (-0.303 ± 0.109), crit=0.392 ± 0.018, hit=1.206 ± 0.062, spell_haste=not significant (-0.036 ± 0.114), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.102

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.96 DPS) [crafted]; Lucky Fishing Hat (19972, -0.96 DPS) [quest]; Flying Tiger Goggles (4368, -2.73 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) (or Magician's Mantle (12998)) | Tailoring [crafted] | 5.0 | yes | Magician's Mantle (12998, +0.00 DPS, sim-verified) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.16 DPS) [crafted]; Slime-encrusted Pads (6461, -0.80 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.16 DPS) [crafted]; Battle Healer's Cloak (20427, -0.16 DPS) [rep]; Feyscale Cloak (6632, -0.32 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.16 DPS) [crafted]; Bloody Apron (6226, -0.16 DPS) [dungeon]; Green Woolen Vest (2582, -2.12 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) (or Windsong Bangles (263336)) | Earthen Arise [quest] | 1.0 | yes | Mindthrust Bracers (1974, -0.16 DPS) [dungeon]; Silver-lined Bracers (3224, -0.16 DPS) [world]; Windsong Bangles (263336, -0.52 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.26 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.48 DPS) [quest]; Pristine Gloves (253913, -0.48 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (36.6 DPS) | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.64 DPS, sim-verified) [crafted]; Lesser Belt of the Spire (1299, -0.64 DPS) [world] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted]; Colorful Kilt (10048, -0.64 DPS) [crafted]; Silk-threaded Trousers (1929, -2.00 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.48 DPS) [crafted]; Pristine Boots (253889, -0.64 DPS) [crafted]; Feather Padded Treads (285345, -1.58 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.80 DPS) [dungeon]; Ring of the Shadow (1462, -0.80 DPS) [world]; Ring of Scorn (3235, -0.80 DPS) [quest] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, -0.40 DPS, sim-verified) [dungeon]; Ring of the Shadow (1462, -0.48 DPS) [world]; Ring of Scorn (3235, -0.48 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Clear Crystal Rod (16894), Lesser Staff of the Spire (1300), Keen Skinning Knife (269716), Riverpaw Mystic Staff (1391), Staff of Orgrimmar (15444), Defias Mage Staff (1928), Skyseeker's Greatstaff (263937), Felweaver's Staff (284287), Apprentice's Spellstaff (248008), Makeshift Stormcaller (281274), Blackfang (2236), Staff of the Blessed Seer (2271), Nightbane Staff (3227), Deadly Bronze Poniard (3490), Shadowfang (1482), Assassin's Blade (1935), Bluegill Kukri (2046), Bronze Shortsword (2850), Decapitating Sword (3740), Cruel Barb (5191), Spikelash Dagger (6333), Wingblade (6504), Crescent Staff (6505), Edward's Knife (251485), Chol'aruk's Trophy Sword (276886), Plaguefang (279876), Staff of Horrors (880), Buzz Saw (1937), Pearl-handled Dagger (5540), Living Root (6631), Butcher's Slicer (6633), Ceranium Rod (3452), Thief's Blade (5192), Emberstone Staff (5201), Militant Shortsword (15211), Scimitar of Atun (1469), Gnarled Hermit's Staff (1539), Buzzer Blade (2169), Brackclaw (2235), Tail Spike (6448), Claystone Shortsword (16891), Night Watch Shortsword (935), Hook Dagger (3184), Big Bronze Knife (3848), Staff of Nobles (3902), Harpy Skinner (5279), Wind Rider Staff (5306), Elegant Shortsword (5321), Ironpatch Blade (12976), Venom Web Fang (899), Riverside Staff (1473), Blackwater Cutlass (1951), Medicine Staff (4575), Goblin Screwdriver (1936), Hollowfang Blade (2020), Northern Shortsword (2078), Serrated Knife (3581), Cursed Felblade (14145), Chanting Blade (14151), Kris of Orgrimmar (15443), Slicer Blade (820), Foamspittle Staff (1405), War Knife (4571), Ritual Blade (5112), Redridge Machete (1219), Defias Rapier (1925), Raider Shortsword (15210), Daggerfang's Daggerfang (281257), Giant Tarantula Fang (1287), Staff of Conjuring (1933), Long Crawler Limb (2088), Cauldron Stirrer (5340), Curved Dagger (2632), Enamelled Broadsword (4765), Pale Skinner (5744), Stonesplinter Dagger (2266), Darkwood Staff (3446), Sturdy Quarterstaff (4566), Feral Blade (4766), Balanced Quarterstaff (257343), Quickblade's Dagger (257346), Gnarlpine War Staff (284166), Jeweled Dagger (1917), Small Green Dagger (4302), Compact Fighting Knife (4974), Militia Sword (248006), Militia Shortblade (248007), Tim's Lost Rib (282064), Bayne's Bite (283462), Pristine Orcish Dagger (286730), Ukta's Conduit (286754), Small Hand Blade (816), Carving Knife (2140), Bloodstained Knife (3225), Jagged Dagger (4947), Copper Dagger (7166), Hoof-Holstered Blade (277988), Helm Splitter (286755), Notched Shortsword (727), Copper Shortsword (2847), Webwood Slicer (282284), Chipped Spellstaff (285238), Explorer's Shortsword (285239)) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS) [quest]; Quilboar Toothpick (276888, +0.00 DPS) [quest]; Twisted Chanter's Staff (890, -0.61 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Nightglow Concoction (3451, -0.80 DPS) [quest]; Pulsating Hydra Heart (5183, -0.80 DPS) [world]; Grayson's Torch (1172, -1.00 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 142.1 | yes | Skycaller (12984, -1.58 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.50 DPS) [dungeon]; Deepblaze (279896, -4.26 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance; 209620 Insignia of the Horde; 241089 Scarlet Dagger

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 55.3. Weights run: 0.8s. Verify run: 0.8s. 214 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.164, intellect=0.928 ± 0.177, crit=0.391 ± 0.020, hit=1.447 ± 0.093, spell_haste=0.972 ± 0.185, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.164

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.3 | yes | Holy Shroud (2721, -0.90 DPS) [world_drop]; Shadow Hood (4323, -1.06 DPS) [crafted]; Nightsky Cowl (4039, -1.64 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.6 | yes | Crystal Starfire Medallion (5003, -1.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.08 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.4 | yes | Death Speaker Mantle (6685, -0.62 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.63 DPS) [quest]; Magician's Mantle (12998, -0.84 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.4 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.39 DPS) [world_drop]; Hillman's Cloak (3719, -0.51 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.1 | yes | Death Speaker Robes (6682, -0.65 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -1.58 DPS) [crafted]; Tree Bark Jacket (1486, -1.69 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.72 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.72 DPS) [quest]; Glowing Magical Bracelets (13106, -2.36 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | sim-verified (54.0 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.07 DPS) [dungeon]; Serpent Gloves (5970, -0.16 DPS) [dungeon]; Jutebraid Gloves (10654, -0.80 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.8 | yes | Belt of Arugal (6392, -0.42 DPS) [dungeon]; Invoker's Cord (215366, -0.45 DPS) [crafted]; Crimson Silk Belt (7055, -1.14 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (54.6 DPS) | yes | Gaze Dreamer Pants (6903, -0.31 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Abomination Skin Leggings (23173, -1.39 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.5 | yes | Spidersilk Boots (4320, -0.58 DPS) [crafted]; Frothing Slippers (254003, -1.46 DPS) [crafted]; Acidic Walkers (9454, -2.34 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Snake Hoop (6750, -0.10 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.21 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.5 | yes | Snake Hoop (6750, -0.01 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (53.1 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -0.40 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.2 | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.19 DPS) [quest]; Channeler's Staff (4437, -0.58 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 161.8 | yes | Starfaller (13063, -0.68 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.79 DPS) [crafted]; Gravestone Scepter (7001, -4.84 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 214, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 107.8. Weights run: 0.7s. Verify run: 0.7s. 291 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.127, intellect=0.992 ± 0.157, crit=0.387 ± 0.020, hit=1.104 ± 0.078, spell_haste=1.008 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.127

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -1.42 DPS) [crafted]; Craftsman's Monocle (4393, -1.71 DPS) [crafted]; Augural Shroud (2620, -2.70 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.0 | yes | Darkspear Warding Pendant (272074, -1.68 DPS) [vendor]; Necklace of Calisea (1714, -1.80 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.24 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.9 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.55 DPS) [dungeon]; Berylline Pads (4197, -0.83 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.0 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Cloak of Rot (4462, -0.85 DPS) [world] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.0 | yes | Robe of Power (7054, -0.57 DPS) [crafted]; Dreamweave Vest (10021, -0.87 DPS, sim-verified) [crafted]; Crimson Silk Vest (7058, -1.69 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (105.0 DPS) | yes | Windchaser Cuffs (14429, -0.02 DPS) [world_drop]; Aurora Bracers (4043, -0.30 DPS) [world_drop]; Radiant Silver Bracers (4545, -1.10 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 | yes | Stormcloth Gloves (10011, -1.70 DPS) [crafted]; Red Mageweave Gloves (10018, -1.84 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.95 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (105.1 DPS) | yes | Gilded Cord (254037, -0.57 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.12 DPS) [rep]; Deathmage Sash (10771, -1.17 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.9 | yes | Crimson Silk Pantaloons (7062, -1.10 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -2.62 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.10 DPS) [dungeon]; Spidersilk Boots (4320, -3.65 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.0 | yes | Reedknot Ring (9622, -2.51 DPS) [quest]; Ogremind Ring (1993, -2.52 DPS) [world_drop]; Voodoo Band (1996, -2.52 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Reedknot Ring (9622, -0.49 DPS, sim-verified) [quest]; Advisor's Ring (19521, -0.56 DPS) [rep]; Ogremind Ring (1993, -0.58 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 22.0 | yes | Windweaver Staff (7757, -0.79 DPS, sim-verified) [dungeon]; Illusionary Rod (7713, -2.96 DPS) [dungeon]; Staff of Jordan (873, -3.09 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (105.4 DPS) | yes | Umbral Wand (5216, -0.01 DPS) [world_drop]; Jaina's Firestarter (13064, -1.50 DPS, sim-verified) [world_drop]; Necrotic Wand (7708, -1.63 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Twisted Nether Wand

No-known-source sample (15 of 291, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10047 Simple Kilt; 10049 Diabolist's Blade

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 185.1. Weights run: 0.7s. Verify run: 0.8s. 372 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.227, intellect=not significant (0.199 ± 0.264), crit=0.457 ± 0.026, hit=2.450 ± 0.155, spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Red Mageweave Headband (10033, -1.13 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.18 DPS) [vendor]; Dreamweave Circlet (10041, -2.34 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.2 | yes | Mindburst Medallion (11196, -0.60 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.52 DPS) [world_drop]; Darkspear Warding Pendant (272073, -1.80 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.6 | yes | Black Mageweave Shoulders (10027, -1.34 DPS) [crafted]; Bloodmage Mantle (7684, -1.62 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.34 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 | yes | Deep Woodlands Cloak (19121, -0.83 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.29 DPS) [crafted]; Nightfall Drape (12465, -1.74 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 | yes | Elemental Raiment (9434, -0.61 DPS) [world_drop]; Stone Guard's Dreadweave Vest (220904, -0.73 DPS) [vendor]; Acumen Robes (17775, -1.37 DPS, sim-verified) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest]; Bloodband Bracers (11469, -0.62 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 | yes | Black Mageweave Gloves (10003, -0.90 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.12 DPS) [vendor]; Runecloth Gloves (13863, -1.40 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | sim-verified (174.6 DPS) | yes | Defiler's Cloth Girdle (20166, -0.33 DPS) [rep]; Ghostweave Cord (254073, -0.56 DPS) [crafted]; Defiler's Cloth Girdle (20165, -2.11 DPS, sim-verified) [rep] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | sim-verified (174.7 DPS) | yes | Red Mageweave Pants (10009, -0.73 DPS) [crafted]; Senior Designer's Pantaloons (11841, -1.57 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -2.21 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (176.7 DPS) | yes | Gilded Sandals (254107, -3.14 DPS) [crafted]; Black Mageweave Boots (10026, -3.25 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.26 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 24.5 | yes | Advisor's Ring (19519, -3.50 DPS) [rep]; Philanthropist's Ring (281635, -3.73 DPS) [quest]; Advisor's Ring (19520, -4.35 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Advisor's Ring (19519, -0.16 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.51 DPS) [quest]; Advisor's Ring (19520, -1.12 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (172.3 DPS) | yes | Uther's Strength (11302, -1.68 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.14 DPS, sim-verified) [crafted]; Tidal Charm (1404, -3.36 DPS) [vendor] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (172.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Tidal Charm (1404, -4.81 DPS) [vendor] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | sim-verified (172.3 DPS) | yes | Staff of Dar'Orahil (15106, -0.28 DPS) [quest]; Shortsword of Vengeance (754, -2.01 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -5.30 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (177.5 DPS) | yes | Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted]; Pyric Caduceus (11748, -5.01 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Rune of the Guard Captain; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 372, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 369.4. Weights run: 0.8s. Verify run: 0.7s. 774 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.837), intellect=not significant (-1.254 ± 0.970), crit=1.999 ± 0.106, hit=6.401 ± 0.516, spell_haste=not significant (0.241 ± 0.963), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.837)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 218.0 | yes | Deathmist Mask (226909, -19.72 DPS) [quest]; Deathmist Mask (22074, -20.01 DPS) [quest]; Bloodvine Goggles (19999, -24.41 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (356.3 DPS) | yes | Blazefury Medallion (17111, -1.24 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -5.23 DPS) [quest]; Orb of the Darkmoon (19426, -6.09 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 132.0 | yes | Rugged Mantle of the Timbermaw (227808, -7.72 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -11.89 DPS) [vendor]; Mantle of the Timbermaw (19050, -12.62 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 80.0 | yes | Howler's Furs (272414, -2.32 DPS) [vendor]; Stalwart Cloak (272415, -2.32 DPS) [vendor]; Earthweave Cloak (21187, -7.38 DPS, sim-verified) [quest] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | sim-verified (369.4 DPS) | yes | Heretic Garb (240146, -1.45 DPS) [vendor]; Earthpower Vest (21183, -9.86 DPS) [quest]; Bloodvine Vest (19682, -13.35 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 91.0 | yes | Heretic Bindings (240145, +0.00 DPS, sim-verified) [vendor]; Heretic Wristguards (240152, -1.89 DPS) [vendor]; Dryad's Wrist Bindings (19595, -10.01 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 101.0 | yes | Deathmist Wraps (226911, -3.48 DPS) [quest]; Gloves of Spell Mastery (14146, -5.23 DPS) [crafted]; Deathmist Wraps (22077, -10.14 DPS, sim-verified) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 160.0 | yes | Heretic Waistguard (240151, -9.77 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -10.29 DPS) [vendor]; Belt of the Archmage (18405, -16.24 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 175.0 | yes | Bloodvine Leggings (19683, -10.73 DPS) [crafted]; Sentinel's Silk Leggings (237815, -12.62 DPS) [vendor]; Heretic Pants (240149, -15.28 DPS, sim-verified) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 100.0 | yes | First Sergeant's Dreadweave Boots (220909, -4.06 DPS) [vendor]; Argent Elite Boots (227816, -5.22 DPS) [vendor]; Bloodvine Boots (19684, -7.50 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (351.9 DPS) | yes | Don Julio's Band (19325, -0.00 DPS) [rep]; Band of Earthen Might (21182, -0.00 DPS) [quest]; Signet Ring of the Bronze Dragonflight (234028, -0.29 DPS) [vendor] |
| finger2 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (356.0 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Blessed Band of Light (272407, +0.00 DPS) [vendor]; Don Julio's Band (19325, -6.24 DPS, sim-verified) [rep] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (347.9 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -3.56 DPS, sim-verified) [vendor] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (351.9 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Shortsword of Vengeance (754) | World drop [world_drop] | sim-verified (356.3 DPS) | yes | Sword of Zeal (6622, +0.00 DPS, sim-verified) [world_drop]; Staff of Dar'Orahil (15106, +0.00 DPS) [quest]; High Warlord's War Staff (234549, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 535.2 | yes | Wand of Biting Cold (19108, -1.52 DPS, sim-verified) [quest]; Stormrager (16997, -14.94 DPS) [quest]; Brilliant Wand (249385, -16.74 DPS) [crafted] |

**New at 60:** head: Heretic Cowl; neck: Beads of Ogre Might; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Rockfury Bracers; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Wrath of Cenarius; trinket1: Serenity Field; trinket2: Abyss Shard; main_hand: Shortsword of Vengeance; ranged: Torch of Light

No-known-source sample (15 of 774, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

