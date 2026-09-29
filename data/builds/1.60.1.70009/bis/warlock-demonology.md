# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 36.4. Weights run: 1.4s. Verify run: 1.1s. 253 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.185, intellect=-0.954 ± 0.213, crit=0.695 ± 0.032, hit=2.196 ± 0.120, spell_haste=2.238 ± 0.209, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.185, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.83 DPS) [crafted]; Lucky Fishing Hat (19972, -0.83 DPS) [quest]; Flying Tiger Goggles (4368, -2.66 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) (or Tarnished Locket (279870)) | Silverwing Sentinels [rep] | 0.0 | yes | Tarnished Locket (279870, -0.03 DPS, sim-verified) [quest] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.22 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.69 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.14 DPS) [crafted]; Caretaker's Cape (20428, -0.14 DPS) [rep]; Feyscale Cloak (6632, -0.23 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.14 DPS) [crafted]; Bloody Apron (6226, -0.14 DPS) [dungeon]; Green Woolen Vest (2582, -1.28 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) (or Silver-lined Bracers (3224), Seer's Cuffs (3645), Bright Bracers (3647), Green Linen Bracers (4308), Timberland Armguards (5315), Willow Bracers (6543), Shimmering Bracers (6563), Tabitha's Cuffs (251486), Cloudy Stormsewn Cuffs (277011), Azure Stormsewn Cuffs (277035), Al'aketh Cuffs (277075), Fellicent's Bindings (283471)) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Seer's Cuffs (3645, +0.00 DPS) [world_drop]; Bright Bracers (3647, +0.00 DPS) [world_drop]; Silver-lined Bracers (3224, -0.18 DPS, sim-verified) [world] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.26 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.69 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.0 | yes | Novice Ardent's Sash (253887, -0.28 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.55 DPS) [world]; Novice Arcanist's Sash (253885, -0.76 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted]; Colorful Kilt (10048, -0.55 DPS) [crafted]; Silk-threaded Trousers (1929, -0.76 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.42 DPS) [crafted]; Feather Padded Treads (285345, -0.50 DPS, sim-verified) [world]; Pristine Boots (253889, -0.55 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.69 DPS) [dungeon]; Ring of the Shadow (1462, -0.69 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.60 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.69 DPS) [dungeon]; Ring of the Shadow (1462, -0.69 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Notched Shortsword (727) (or Small Hand Blade (816), Slicer Blade (820), Staff of Horrors (880), Twisted Chanter's Staff (890), Venom Web Fang (899), Night Watch Shortsword (935), Redridge Machete (1219), Giant Tarantula Fang (1287), Lesser Staff of the Spire (1300), Riverpaw Mystic Staff (1391), Foamspittle Staff (1405), Scimitar of Atun (1469), Riverside Staff (1473), Shadowfang (1482), Witching Stave (1484), Gnarled Hermit's Staff (1539), Jeweled Dagger (1917), Defias Rapier (1925), Defias Mage Staff (1928), Staff of Conjuring (1933), Assassin's Blade (1935), Goblin Screwdriver (1936), Buzz Saw (1937), Blackwater Cutlass (1951), Hollowfang Blade (2020), Staff of Westfall (2042), Bluegill Kukri (2046), Solid Shortblade (2074), Northern Shortsword (2078), Long Crawler Limb (2088), Scrimshaw Dagger (2089), Carving Knife (2140), Buzzer Blade (2169), Craftsman's Dagger (2218), Brackclaw (2235), Blackfang (2236), Stonesplinter Dagger (2266), Staff of the Blessed Seer (2271), Evocator's Blade (2567), Curved Dagger (2632), Copper Shortsword (2847), Bronze Shortsword (2850), Thornblade (2908), Hook Dagger (3184), Bloodstained Knife (3225), Nightbane Staff (3227), Deadly Bronze Poniard (3490), Daryl's Shortsword (3572), Decapitating Sword (3740), Big Bronze Knife (3848), Staff of Nobles (3902), Small Green Dagger (4302), Channeler's Staff (4437), Sturdy Quarterstaff (4566), War Knife (4571), Medicine Staff (4575), Enamelled Broadsword (4765), Feral Blade (4766), Ritual Blade (5112), Cruel Barb (5191), Thief's Blade (5192), Emberstone Staff (5201), Pearl-handled Dagger (5540), Pale Skinner (5744), Balanced Fighting Stick (6215), Spikelash Dagger (6333), Tail Spike (6448), Living Root (6631), Butcher's Slicer (6633), Copper Dagger (7166), Gritroot Staff (9603), Briarsteel Shortsword (15335), Curvewood Dagger (15396), Oakthrush Staff (15397), Brushwood Blade (18957), Edward's Knife (251485), Gnarled Necromancer's Staff (251534), Skyseeker's Greatstaff (263937), Chol'aruk's Trophy Sword (276886), Quilboar Toothpick (276888), Plaguefang (279876), Webwood Slicer (282284), Bayne's Bite (283462), Gnarlpine War Staff (284166), Felweaver's Staff (284287), Chipped Spellstaff (285238), Explorer's Shortsword (285239), Pristine Orcish Dagger (286730), Ukta's Conduit (286754), Helm Splitter (286755)) | World drop [world_drop] | 0.0 | yes | Small Hand Blade (816, +0.00 DPS, sim-verified) [world]; Slicer Blade (820, +0.00 DPS) [world]; Staff of Horrors (880, +0.00 DPS) [world] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Pulsating Hydra Heart (5183, -0.69 DPS) [world]; Tear of Grief (5611, -0.69 DPS) [quest]; Grayson's Torch (1172, -1.13 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 3.0 | yes | Sable Wand (7607, +0.00 DPS) [quest]; Torchlight Wand (5240, -0.14 DPS) [quest]; Sizzle Stick (8071, -0.64 DPS, sim-verified) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Notched Shortsword; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 253, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 54.6. Weights run: 1.4s. Verify run: 1.2s. 476 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.134, intellect=not significant (0.398 ± 0.149), crit=0.570 ± 0.027, hit=1.837 ± 0.085, spell_haste=0.913 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.134, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Silk Headband (7050, -0.36 DPS) [crafted]; Embalmed Shroud (7691, -0.55 DPS) [dungeon]; Enchanter's Cowl (4322, -1.14 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 | yes | Crystal Starfire Medallion (5003, -1.42 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.45 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.71 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 | yes | Fairywing Mantle (9536, -0.55 DPS) [quest]; Invoker's Mantle (215365, -0.65 DPS) [crafted]; Death Speaker Mantle (6685, -0.65 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Prelacy Cape (7004, -0.18 DPS) [quest]; Repairman's Cape (9605, -0.38 DPS, sim-verified) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Death Speaker Robes (6682, -0.30 DPS) [dungeon]; Pristine Gown (253961, -0.58 DPS) [crafted]; Green Silk Armor (7065, -0.68 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -1.20 DPS) [quest]; Mindthrust Bracers (1974, -1.28 DPS) [dungeon]; Nightsky Wristbands (6407, -2.13 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Shilly Mitts (9609, +0.00 DPS) [quest]; Truefaith Gloves (7049, -0.15 DPS) [crafted]; Town Clerk's Mittens (270029, -0.56 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.2 | yes | Belt of Arugal (6392, -0.41 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.58 DPS) [crafted]; Crimson Silk Belt (7055, -0.62 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Pristine Leggings (253987, -0.40 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted]; Abomination Skin Leggings (23173, -0.66 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.8 | yes | Acidic Walkers (9454, -0.29 DPS) [dungeon]; Nimbus Boots (6998, -0.69 DPS) [quest]; Spidersilk Boots (4320, -1.71 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.22 DPS) [quest]; Lorekeeper's Ring (20431, -0.36 DPS) [rep]; Electrocutioner Lagnut (9447, -0.73 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.55 DPS) [dungeon]; Sludge-Stained Band (286535, -0.55 DPS) [world]; Minor Channeling Ring (1449, -1.73 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.4 | yes | Twisted Chanter's Staff (890, +0.14 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.07 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | World drop [world_drop] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Spellcrafter Wand (6677, -0.18 DPS) [quest]; Sizzle Stick (8071, -1.88 DPS, sim-verified) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 476, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 105.5. Weights run: 1.3s. Verify run: 1.0s. 643 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.155, intellect=not significant (0.238 ± 0.169), crit=0.439 ± 0.023, hit=1.419 ± 0.085, spell_haste=0.810 ± 0.162, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.155, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -1.68 DPS, sim-verified) [world]; Holy Shroud (2721, -2.56 DPS) [world_drop]; Silk Headband (7050, -3.07 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 | yes | Darkspear Warding Pendant (272074, -1.73 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.85 DPS) [vendor]; Necklace of Calisea (1714, -1.99 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 | yes | Green Silken Shoulders (7057, -0.13 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.27 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.05 DPS) [crafted]; Caretaker's Cape (19532, -0.30 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.4 | yes | Dreamweave Vest (10021, -0.94 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.68 DPS) [crafted]; Crimson Silk Vest (7058, -2.32 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, -0.35 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.51 DPS) [quest]; Earthen Silk Cuffs (254019, -1.28 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | Black Mageweave Gloves (10003, -0.84 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.43 DPS) [crafted]; Gilded Handwraps (254021, -2.38 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.0 | yes | Star Belt (4329, -0.37 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -0.83 DPS) [rep]; Deathmage Sash (10771, -1.12 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.9 | yes | Gaze Dreamer Pants (6903, -0.90 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.48 DPS) [crafted]; Abomination Skin Leggings (23173, -1.52 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -2.48 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.11 DPS) [crafted]; Acidic Walkers (9454, -4.38 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.51 DPS) [quest]; Lorekeeper's Ring (19525, -0.51 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.77 DPS) [vendor] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 | yes | Sea Giant's Toe Ring (274746, -0.51 DPS) [vendor]; Reedknot Ring (9622, -0.53 DPS, sim-verified) [quest]; Minor Channeling Ring (1449, -0.65 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 16.8 | yes | Illusionary Rod (7713, -0.63 DPS, sim-verified) [dungeon]; Staff of Noh'Orahil (15105, -2.49 DPS) [quest]; Windweaver Staff (7757, -3.39 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Burning Sliver (5249) (or Twisted Nether Wand (249144)) | Crushridge Warmongers [quest] | 6.0 | yes | Twisted Nether Wand (249144, -0.25 DPS, sim-verified) [crafted]; Wand of Eventide (5214, -0.26 DPS) [world_drop]; Sizzle Stick (8071, -0.26 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Burning Sliver

No-known-source sample (15 of 643, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 124.8. Weights run: 1.3s. Verify run: 1.0s. 804 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.555), intellect=not significant (-2.913 ± 0.807), crit=1.725 ± 0.106, hit=4.994 ± 0.369, spell_haste=not significant (0.288 ± 0.740), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.840 ± 0.555), fire_power=0.161 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 38.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.03 DPS) [crafted]; Eye of Theradras (17715, -1.29 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.36 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -0.64 DPS) [world_drop]; Pendant of Myzrael (4614, -0.64 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 32.1 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -1.76 DPS) [dungeon]; Black Mageweave Shoulders (10027, -2.04 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Runecloth Cloak (13860, -0.46 DPS) [crafted]; Icy Cloak (4327, -0.64 DPS) [crafted]; Nightfall Drape (12465, -1.57 DPS, sim-verified) [world] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 36.1 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -1.30 DPS) [world_drop]; Acumen Robes (17775, -1.58 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.54 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.18 DPS) [quest]; Nethergeld Cuffs (254061, -0.18 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.46 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.46 DPS) [vendor]; Black Mageweave Gloves (10003, -1.12 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 33.1 | yes | Satyrmane Sash (17755, +0.87 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -1.76 DPS) [rep]; Ghostweave Cord (254073, -1.76 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 36.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -1.58 DPS) [crafted]; Red Mageweave Pants (10009, -2.04 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 57.9 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -3.13 DPS) [crafted]; Black Mageweave Boots (10026, -4.32 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 49.9 | yes | Ring of Forlorn Spirits (2043, -3.86 DPS) [quest]; Reedknot Ring (9622, -3.95 DPS) [quest]; Sea Giant's Toe Ring (274746, -4.05 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Ring of Forlorn Spirits (2043, -0.37 DPS) [quest]; Reedknot Ring (9622, -0.46 DPS) [quest]; Lorekeeper's Ring (19524, -1.22 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 44.9 | yes | Abyss Shard (20534, -3.03 DPS) [quest]; Uther's Strength (11302, -3.59 DPS) [world]; Thunderbrew's Boot Flask (744, -4.14 DPS) [quest] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Abyss Shard (20534, +1.10 DPS) [quest]; Uther's Strength (11302, +0.55 DPS) [world]; Thunderbrew's Boot Flask (744, +0.00 DPS) [quest] |
| main_hand | Staff of Dar'Orahil (15106) (or Soul Harvester (20536)) | The Completed Orb of Dar'Orahil [quest] | 49.9 | yes | Soul Harvester (20536, -0.35 DPS, sim-verified) [quest]; Illusionary Rod (7713, -2.37 DPS) [dungeon]; Kindling Stave (11750, -2.37 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Burning Sliver (5249, -0.18 DPS) [quest]; Twisted Nether Wand (249144, -0.18 DPS) [crafted]; Dreambough Wand (249234, -0.49 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; ranged: Lesser Eternal Wand

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 324.3. Weights run: 1.3s. Verify run: 1.1s. 1166 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.620, intellect=4.547 ± -0.743, crit=-1.774 ± -0.098, hit=-5.572 ± -0.332, spell_haste=-2.242 ± -0.730, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.083 ± -0.620, fire_power=-0.082 ± -0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | 146.7 | yes | Field Marshal's Coronal (17578, +1.05 DPS) [vendor]; Doomcaller's Circlet (21337, +0.86 DPS) [quest]; Magister's Crown (16686, -9.58 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 0.0 | yes | Pendant of Forgotten Names (22947, -0.94 DPS, sim-verified) [raid]; Amulet of the Dawn (22657, -14.03 DPS) [quest]; Charm of the Shifting Sands (21504, -15.06 DPS) [quest] |
| shoulder | Field Marshal's Dreadweave Shoulders (17580) | Captain Dirgehammer [vendor] | 102.3 | yes | Warlord's Dreadweave Mantle (17590, -0.00 DPS) [vendor]; Field Marshal's Dreadweave Shoulders (231583, -0.00 DPS) [pvp]; Magister's Mantle (16689, -4.56 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 59.5 | yes | Cloak of Clarity (21583, +0.93 DPS) [raid]; Shroud of Unspoken Names (21418, +0.10 DPS) [quest]; Darkspear Raider's Cloak (272076, -7.95 DPS, sim-verified) [vendor] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 151.0 | yes | Warlord's Dreadweave Robe (17592, +1.88 DPS) [vendor]; Field Marshal's Dreadweave Robe (231582, +1.88 DPS) [pvp]; Field Marshal's Dreadweave Robe (17581, -9.17 DPS, sim-verified) [vendor] |
| wrist | Burrower Bracers (21611) | Ahn'Qiraj: Ouro [raid] | 87.1 | yes | Marshal's Dreadweave Cuffs (17582, +0.14 DPS) [pvp]; General's Dreadweave Bracers (17587, +0.14 DPS) [pvp]; Plagueheart Bindings (22511, -1.54 DPS, sim-verified) [quest] |
| hands | Dark Storm Gauntlets (21585) | Ahn'Qiraj: C'Thun [raid] | 105.2 | yes | Gloves of the Messiah (21619, +3.58 DPS) [raid]; Plagueheart Gloves (22509, +0.36 DPS) [quest]; Raider Handwraps (272098, -8.54 DPS, sim-verified) [vendor] |
| waist | Marshal's Dreadweave Sash (17585) (or General's Dreadweave Belt (17589)) | Rank 16 [pvp] | 104.6 | yes | Devout Belt (16696, +1.83 DPS) [world]; Magister's Belt (16685, +0.40 DPS) [dungeon]; General's Dreadweave Belt (17589, +0.00 DPS, sim-verified) [pvp] |
| legs | Plagueheart Leggings (22505) | Plagueheart Leggings [quest] | 150.7 | yes | Marshal's Dreadweave Leggings (17579, +5.17 DPS) [vendor]; General's Dreadweave Pants (17593, +5.17 DPS) [vendor]; Doomcaller's Trousers (21336, -1.08 DPS, sim-verified) [quest] |
| feet | Boots of Epiphany (21600) | Ahn'Qiraj: Emperor Vek'lor [raid] | 120.4 | yes | Bloodvine Boots (19684, +5.42 DPS) [crafted]; Doomcaller's Footwraps (21338, +3.72 DPS) [quest]; Plagueheart Sandals (22508, +0.14 DPS, sim-verified) [quest] |
| finger1 | Band of Sulfuras (19138) | Molten Core: Ragnaros [raid] | 104.6 | yes | Signet Ring of the Bronze Dragonflight (21210, +6.75 DPS) [quest]; Cauterizing Band (19140, +6.63 DPS) [world]; Elemental Focus Band (20682, +6.17 DPS) [world] |
| finger2 | Band of Forced Concentration (19403) | Blackwing Lair: Ebonroc [raid] | 75.6 | yes | Signet Ring of the Bronze Dragonflight (21210, +1.26 DPS) [quest]; Cauterizing Band (19140, +1.14 DPS) [world]; Elemental Focus Band (20682, -1.69 DPS, sim-verified) [world] |
| trinket1 | Eye of the Dead (23047) | Naxxramas: Sapphiron [raid] | 0.0 | yes | Uther's Strength (11302, -1.14 DPS) [raid]; Abyss Shard (20534, -2.27 DPS) [quest]; Neltharion's Tear (19379, -8.33 DPS) [raid] |
| trinket2 | The Restrained Essence of Sapphiron (23046) | Naxxramas: Sapphiron [raid] | 40.0 | yes | Uther's Strength (11302, +6.44 DPS) [raid]; Abyss Shard (20534, +5.30 DPS) [quest]; Neltharion's Tear (19379, -0.76 DPS) [raid] |
| main_hand | Atiesh, Greatstaff of the Guardian (22589) | Atiesh, Greatstaff of the Guardian [quest] | 295.5 | yes | High Warlord's War Staff (234549, +20.81 DPS) [pvp]; Atiesh, Greatstaff of the Guardian (22631, +9.12 DPS) [quest]; Atiesh, Greatstaff of the Guardian (22630, +2.58 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Ledger from Tanaris [quest] | 23.2 | yes | Flash Wand (5248, +0.95 DPS) [quest]; Stormrager (16997, +0.09 DPS) [quest]; Cairnstone Sliver (9654, -6.44 DPS, sim-verified) [quest] |

**New at 60:** head: Plagueheart Circlet; neck: Blazefury Medallion; shoulder: Field Marshal's Dreadweave Shoulders; back: Hide of the Wild; chest: Plagueheart Robe; wrist: Burrower Bracers; hands: Dark Storm Gauntlets; waist: Marshal's Dreadweave Sash; legs: Plagueheart Leggings; feet: Boots of Epiphany; finger1: Band of Sulfuras; finger2: Band of Forced Concentration; trinket1: Eye of the Dead; trinket2: The Restrained Essence of Sapphiron; main_hand: Atiesh, Greatstaff of the Guardian; ranged: Charged Lightning Rod

No-known-source sample (15 of 1166, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 34.7. Weights run: 1.4s. Verify run: 1.1s. 250 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.185, intellect=-0.954 ± 0.213, crit=0.695 ± 0.032, hit=2.196 ± 0.120, spell_haste=2.238 ± 0.209, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.185, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.83 DPS) [crafted]; Lucky Fishing Hat (19972, -0.83 DPS) [quest]; Flying Tiger Goggles (4368, -2.18 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) (or Tarnished Locket (279870)) | Warsong Outriders [rep] | 0.0 | yes | Tarnished Locket (279870, -0.03 DPS, sim-verified) [quest] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Slime-encrusted Pads (6461, -0.69 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.94 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.14 DPS) [crafted]; Battle Healer's Cloak (20427, -0.14 DPS) [rep]; Feyscale Cloak (6632, -0.26 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.14 DPS) [crafted]; Bloody Apron (6226, -0.14 DPS) [dungeon]; Green Woolen Vest (2582, -2.34 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.0 | yes | Mindthrust Bracers (1974, +0.29 DPS, sim-verified) [dungeon]; Silver-lined Bracers (3224, -0.14 DPS) [world]; Seer's Cuffs (3645, -0.14 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.42 DPS) [quest]; Pristine Gloves (253913, -0.42 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.0 | yes | Novice Ardent's Sash (253887, -0.28 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.55 DPS) [world]; Novice Arcanist's Sash (253885, -0.74 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted]; Colorful Kilt (10048, -0.55 DPS) [crafted]; Silk-threaded Trousers (1929, -2.32 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.42 DPS) [crafted]; Pristine Boots (253889, -0.55 DPS) [crafted]; Feather Padded Treads (285345, -1.18 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.69 DPS) [dungeon]; Ring of the Shadow (1462, -0.69 DPS) [world]; Ring of Scorn (3235, -0.69 DPS) [quest] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, +0.09 DPS, sim-verified) [dungeon]; Ring of the Shadow (1462, -0.42 DPS) [world]; Ring of Scorn (3235, -0.42 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Notched Shortsword (727) (or Small Hand Blade (816), Slicer Blade (820), Staff of Horrors (880), Twisted Chanter's Staff (890), Venom Web Fang (899), Night Watch Shortsword (935), Redridge Machete (1219), Giant Tarantula Fang (1287), Lesser Staff of the Spire (1300), Riverpaw Mystic Staff (1391), Foamspittle Staff (1405), Scimitar of Atun (1469), Riverside Staff (1473), Shadowfang (1482), Witching Stave (1484), Gnarled Hermit's Staff (1539), Jeweled Dagger (1917), Defias Rapier (1925), Defias Mage Staff (1928), Staff of Conjuring (1933), Assassin's Blade (1935), Goblin Screwdriver (1936), Buzz Saw (1937), Blackwater Cutlass (1951), Hollowfang Blade (2020), Bluegill Kukri (2046), Northern Shortsword (2078), Long Crawler Limb (2088), Carving Knife (2140), Buzzer Blade (2169), Brackclaw (2235), Blackfang (2236), Stonesplinter Dagger (2266), Staff of the Blessed Seer (2271), Evocator's Blade (2567), Curved Dagger (2632), Copper Shortsword (2847), Bronze Shortsword (2850), Hook Dagger (3184), Bloodstained Knife (3225), Nightbane Staff (3227), Darkwood Staff (3446), Ceranium Rod (3452), Deadly Bronze Poniard (3490), Serrated Knife (3581), Decapitating Sword (3740), Big Bronze Knife (3848), Staff of Nobles (3902), Small Green Dagger (4302), Channeler's Staff (4437), Sturdy Quarterstaff (4566), War Knife (4571), Medicine Staff (4575), Enamelled Broadsword (4765), Feral Blade (4766), Jagged Dagger (4947), Compact Fighting Knife (4974), Ritual Blade (5112), Cruel Barb (5191), Thief's Blade (5192), Emberstone Staff (5201), Harpy Skinner (5279), Wind Rider Staff (5306), Elegant Shortsword (5321), Cauldron Stirrer (5340), Pearl-handled Dagger (5540), Pale Skinner (5744), Spikelash Dagger (6333), Tail Spike (6448), Wingblade (6504), Crescent Staff (6505), Living Root (6631), Butcher's Slicer (6633), Copper Dagger (7166), Cursed Felblade (14145), Chanting Blade (14151), Kris of Orgrimmar (15443), Staff of Orgrimmar (15444), Claystone Shortsword (16891), Clear Crystal Rod (16894), Edward's Knife (251485), Gnarled Necromancer's Staff (251534), Skyseeker's Greatstaff (263937), Chol'aruk's Trophy Sword (276886), Quilboar Toothpick (276888), Plaguefang (279876), Webwood Slicer (282284), Bayne's Bite (283462), Gnarlpine War Staff (284166), Felweaver's Staff (284287), Chipped Spellstaff (285238), Explorer's Shortsword (285239), Pristine Orcish Dagger (286730), Ukta's Conduit (286754), Helm Splitter (286755)) | World drop [world_drop] | 0.0 | yes | Small Hand Blade (816, +0.00 DPS, sim-verified) [world]; Slicer Blade (820, +0.00 DPS) [world]; Staff of Horrors (880, +0.00 DPS) [world] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Nightglow Concoction (3451, -0.69 DPS) [quest]; Pulsating Hydra Heart (5183, -0.69 DPS) [world]; Grayson's Torch (1172, -0.96 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 3.0 | yes | Torchlight Wand (5240, -0.14 DPS) [quest]; Greater Magic Wand (11288, -0.14 DPS) [crafted]; Sizzle Stick (8071, -0.57 DPS, sim-verified) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Notched Shortsword; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 250, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 52.5. Weights run: 1.4s. Verify run: 1.2s. 474 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.134, intellect=not significant (0.398 ± 0.149), crit=0.570 ± 0.027, hit=1.837 ± 0.085, spell_haste=0.913 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.134, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Silk Headband (7050, -0.36 DPS) [crafted]; Embalmed Shroud (7691, -0.55 DPS) [dungeon]; Enchanter's Cowl (4322, -0.68 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 | yes | Crystal Starfire Medallion (5003, -1.42 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.60 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.71 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 | yes | Death Speaker Mantle (6685, -0.47 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.55 DPS) [quest]; Invoker's Mantle (215365, -0.65 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Battle Healer's Cloak (19529, -0.18 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.2 | yes | Tree Bark Jacket (1486, +0.29 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.51 DPS) [dungeon]; Pristine Gown (253961, -0.80 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -1.20 DPS) [quest]; Mindthrust Bracers (1974, -1.28 DPS) [dungeon]; Nightsky Wristbands (6407, -1.59 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.0 | yes | Serpent Gloves (5970, -0.06 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.33 DPS) [crafted]; Gnoll Casting Gloves (892, -0.36 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.2 | yes | Warsong Sash (16975, -0.22 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.36 DPS) [dungeon]; Invoker's Cord (215366, -0.58 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 | yes | Gaze Dreamer Pants (6903, +0.31 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.44 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.69 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.8 | yes | Acidic Walkers (9454, -0.29 DPS) [dungeon]; Boots of the Enchanter (4325, -0.87 DPS) [crafted]; Spidersilk Boots (4320, -1.76 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.36 DPS) [rep]; Electrocutioner Lagnut (9447, -0.73 DPS) [dungeon]; Sludge-Stained Band (286535, -0.73 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.55 DPS) [world]; Black Widow Band (6199, -0.58 DPS) [world]; Electrocutioner Lagnut (9447, -1.96 DPS, sim-verified) [dungeon] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Darkspear Voodoo Seal (272059, +0.00 DPS) [vendor] |
| trinket2 | Relentless Raider's Seal (272062) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS, sim-verified) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Darkspear Voodoo Seal (272059, +0.00 DPS) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.4 | yes | Twisted Chanter's Staff (890, -0.04 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.07 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | World drop [world_drop] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Cookie's Stirring Rod (5198, -0.36 DPS) [dungeon]; Sizzle Stick (8071, -1.80 DPS, sim-verified) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Defiler's Talisman; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 474, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 104.9. Weights run: 1.3s. Verify run: 1.1s. 640 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.155, intellect=not significant (0.238 ± 0.169), crit=0.439 ± 0.023, hit=1.419 ± 0.085, spell_haste=0.810 ± 0.162, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.155, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -2.32 DPS, sim-verified) [world]; Holy Shroud (2721, -2.56 DPS) [world_drop]; Silk Headband (7050, -3.07 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 | yes | Necklace of Calisea (1714, -1.72 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -1.73 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.85 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 | yes | Green Silken Shoulders (7057, -0.21 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.27 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.05 DPS) [crafted]; Battle Healer's Cloak (19528, -0.30 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.4 | yes | Dreamweave Vest (10021, -0.61 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.68 DPS) [crafted]; Crimson Silk Vest (7058, -2.32 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Condor Bracers (15864, -0.63 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.79 DPS) [quest]; Earthen Silk Cuffs (254019, -1.28 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | Black Mageweave Gloves (10003, -1.20 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.43 DPS) [crafted]; Gilded Handwraps (254021, -2.38 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.0 | yes | Defiler's Cloth Girdle (20164, -0.83 DPS) [rep]; Star Belt (4329, -0.84 DPS, sim-verified) [crafted]; Warsong Sash (16975, -1.01 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.9 | yes | Gaze Dreamer Pants (6903, -0.81 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.48 DPS) [crafted]; Abomination Skin Leggings (23173, -1.52 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -2.32 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.11 DPS) [crafted]; Acidic Walkers (9454, -4.38 DPS) [dungeon] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.51 DPS) [rep]; Advisor's Ring (20426, -1.02 DPS) [rep]; Reedknot Ring (9622, -2.18 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Sludge-Stained Band (286535, -0.77 DPS) [world]; Reedknot Ring (9622, -1.37 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 16.8 | yes | Illusionary Rod (7713, -1.20 DPS, sim-verified) [dungeon]; Staff of Noh'Orahil (15105, -2.49 DPS) [quest]; Windweaver Staff (7757, -3.39 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 6.0 | yes | Sizzle Stick (8071, -0.26 DPS) [quest]; Greater Mystic Wand (217287, -0.26 DPS) [crafted]; Wand of Eventide (5214, -0.47 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Dar'Orahil; ranged: Twisted Nether Wand

No-known-source sample (15 of 640, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 122.4. Weights run: 1.3s. Verify run: 1.0s. 801 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.555), intellect=not significant (-2.913 ± 0.807), crit=1.725 ± 0.106, hit=4.994 ± 0.369, spell_haste=not significant (0.288 ± 0.740), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.840 ± 0.555), fire_power=0.161 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 38.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.03 DPS) [crafted]; Eye of Theradras (17715, -1.29 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.45 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -0.64 DPS) [world_drop]; Choker of the High Shaman (4112, -0.64 DPS) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 32.1 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -1.76 DPS) [dungeon]; Black Mageweave Shoulders (10027, -2.04 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Deep Woodlands Cloak (19121, -0.28 DPS, sim-verified) [quest]; Nightfall Drape (12465, -0.46 DPS) [world]; Runecloth Cloak (13860, -0.46 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 36.1 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -1.30 DPS) [world_drop]; Acumen Robes (17775, -1.58 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, -0.18 DPS) [crafted]; Bloodband Bracers (11469, -0.37 DPS) [quest]; Condor Bracers (15864, -0.63 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.46 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.46 DPS) [vendor]; Black Mageweave Gloves (10003, -0.94 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 33.1 | yes | Satyrmane Sash (17755, +0.75 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -1.76 DPS) [rep]; Ghostweave Cord (254073, -1.76 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 36.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -1.58 DPS) [crafted]; Red Mageweave Pants (10009, -2.04 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 57.9 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -3.13 DPS) [crafted]; Black Mageweave Boots (10026, -4.32 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 49.9 | yes | Reedknot Ring (9622, -3.95 DPS) [quest]; Sea Giant's Toe Ring (274746, -4.05 DPS) [vendor]; Electrocutioner Lagnut (9447, -4.32 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Reedknot Ring (9622, -0.46 DPS) [quest]; Advisor's Ring (19521, -0.46 DPS) [rep]; Advisor's Ring (19520, -0.75 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 44.9 | yes | Rune of the Guard Captain (19120, -0.92 DPS) [quest]; Abyss Shard (20534, -3.03 DPS) [quest]; Uther's Strength (11302, -3.59 DPS) [world] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +3.22 DPS) [quest]; Abyss Shard (20534, +1.10 DPS) [quest]; Uther's Strength (11302, +0.55 DPS) [world] |
| main_hand | Staff of Dar'Orahil (15106) (or Soul Harvester (20536)) | The Completed Orb of Dar'Orahil [quest] | 49.9 | yes | Soul Harvester (20536, -0.14 DPS, sim-verified) [quest]; Illusionary Rod (7713, -2.37 DPS) [dungeon]; Kindling Stave (11750, -2.37 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Twisted Nether Wand (249144, -0.18 DPS) [crafted]; Wand of Eventide (5214, -0.28 DPS) [world_drop]; Dreambough Wand (249234, -0.36 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; ranged: Lesser Eternal Wand

No-known-source sample (15 of 801, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 317.3. Weights run: 1.3s. Verify run: 1.1s. 1163 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.620, intellect=4.547 ± -0.743, crit=-1.774 ± -0.098, hit=-5.572 ± -0.332, spell_haste=-2.242 ± -0.730, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.083 ± -0.620, fire_power=-0.082 ± -0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | 146.7 | yes | Field Marshal's Coronal (17578, +1.05 DPS) [vendor]; Doomcaller's Circlet (21337, +0.86 DPS) [quest]; Magister's Crown (16686, -8.88 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 0.0 | yes | Pendant of Forgotten Names (22947, +0.26 DPS, sim-verified) [raid]; Amulet of the Dawn (22657, -14.03 DPS) [quest]; Charm of the Shifting Sands (21504, -15.06 DPS) [quest] |
| shoulder | Field Marshal's Dreadweave Shoulders (17580) | Captain Dirgehammer [vendor] | 102.3 | yes | Warlord's Dreadweave Mantle (17590, -0.00 DPS) [vendor]; Field Marshal's Dreadweave Shoulders (231583, -0.00 DPS) [pvp]; Magister's Mantle (16689, -5.82 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 59.5 | yes | Cloak of Clarity (21583, +0.93 DPS) [raid]; Shroud of Unspoken Names (21418, +0.10 DPS) [quest]; Darkspear Raider's Cloak (272076, -7.07 DPS, sim-verified) [vendor] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 151.0 | yes | Warlord's Dreadweave Robe (17592, +1.88 DPS) [vendor]; Field Marshal's Dreadweave Robe (231582, +1.88 DPS) [pvp]; Field Marshal's Dreadweave Robe (17581, -7.68 DPS, sim-verified) [vendor] |
| wrist | Burrower Bracers (21611) | Ahn'Qiraj: Ouro [raid] | 87.1 | yes | Marshal's Dreadweave Cuffs (17582, +0.14 DPS) [pvp]; General's Dreadweave Bracers (17587, +0.14 DPS) [pvp]; Plagueheart Bindings (22511, -1.12 DPS, sim-verified) [quest] |
| hands | Dark Storm Gauntlets (21585) | Ahn'Qiraj: C'Thun [raid] | 105.2 | yes | Gloves of the Messiah (21619, +3.58 DPS) [raid]; Plagueheart Gloves (22509, +0.36 DPS) [quest]; Raider Handwraps (272098, -8.45 DPS, sim-verified) [vendor] |
| waist | Marshal's Dreadweave Sash (17585) (or General's Dreadweave Belt (17589)) | Rank 16 [pvp] | 104.6 | yes | Devout Belt (16696, +1.83 DPS) [world]; Magister's Belt (16685, +0.40 DPS) [dungeon]; General's Dreadweave Belt (17589, +0.00 DPS, sim-verified) [pvp] |
| legs | Plagueheart Leggings (22505) | Plagueheart Leggings [quest] | 150.7 | yes | Marshal's Dreadweave Leggings (17579, +5.17 DPS) [vendor]; General's Dreadweave Pants (17593, +5.17 DPS) [vendor]; Doomcaller's Trousers (21336, -0.63 DPS, sim-verified) [quest] |
| feet | Boots of Epiphany (21600) | Ahn'Qiraj: Emperor Vek'lor [raid] | 120.4 | yes | Bloodvine Boots (19684, +5.42 DPS) [crafted]; Doomcaller's Footwraps (21338, +3.72 DPS) [quest]; Plagueheart Sandals (22508, -0.78 DPS, sim-verified) [quest] |
| finger1 | Band of Sulfuras (19138) | Molten Core: Ragnaros [raid] | 104.6 | yes | Signet Ring of the Bronze Dragonflight (21210, +6.75 DPS) [quest]; Cauterizing Band (19140, +6.63 DPS) [world]; Elemental Focus Band (20682, +6.17 DPS) [world] |
| finger2 | Band of Forced Concentration (19403) | Blackwing Lair: Ebonroc [raid] | 75.6 | yes | Signet Ring of the Bronze Dragonflight (21210, +1.26 DPS) [quest]; Cauterizing Band (19140, +1.14 DPS) [world]; Elemental Focus Band (20682, -1.32 DPS, sim-verified) [world] |
| trinket1 | Eye of the Dead (23047) | Naxxramas: Sapphiron [raid] | 0.0 | yes | Uther's Strength (11302, -1.14 DPS) [raid]; Abyss Shard (20534, -2.27 DPS) [quest]; Neltharion's Tear (19379, -8.33 DPS) [raid] |
| trinket2 | The Restrained Essence of Sapphiron (23046) | Naxxramas: Sapphiron [raid] | 40.0 | yes | Uther's Strength (11302, +6.44 DPS) [raid]; Abyss Shard (20534, +5.30 DPS) [quest]; Neltharion's Tear (19379, -0.76 DPS) [raid] |
| main_hand | Atiesh, Greatstaff of the Guardian (22589) | Atiesh, Greatstaff of the Guardian [quest] | 295.5 | yes | High Warlord's War Staff (234549, +20.81 DPS) [pvp]; Atiesh, Greatstaff of the Guardian (22631, +9.12 DPS) [quest]; Atiesh, Greatstaff of the Guardian (22630, +2.58 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Nature's Breath (19118) | Dark Vessels [quest] | 27.3 | yes | Charged Lightning Rod (11860, +2.17 DPS, sim-verified) [quest]; Flash Wand (5248, +1.72 DPS) [quest]; Stormrager (16997, +0.86 DPS) [quest] |

**New at 60:** head: Plagueheart Circlet; neck: Blazefury Medallion; shoulder: Field Marshal's Dreadweave Shoulders; back: Hide of the Wild; chest: Plagueheart Robe; wrist: Burrower Bracers; hands: Dark Storm Gauntlets; waist: Marshal's Dreadweave Sash; legs: Plagueheart Leggings; feet: Boots of Epiphany; finger1: Band of Sulfuras; finger2: Band of Forced Concentration; trinket1: Eye of the Dead; trinket2: The Restrained Essence of Sapphiron; main_hand: Atiesh, Greatstaff of the Guardian; ranged: Nature's Breath

No-known-source sample (15 of 1163, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

