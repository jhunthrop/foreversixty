# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 35.9. Weights run: 0.8s. Verify run: 0.8s. 128 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.103, intellect=not significant (-0.337 ± 0.116), crit=0.435 ± 0.021, hit=1.467 ± 0.069, spell_haste=0.633 ± 0.114, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.103, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.81 DPS) [crafted]; Lucky Fishing Hat (19972, -0.81 DPS) [quest]; Flying Tiger Goggles (4368, -2.76 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) (or Magician's Mantle (12998)) | Tailoring [crafted] | 5.0 | yes | Magician's Mantle (12998, +0.00 DPS, sim-verified) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.13 DPS) [crafted]; Slime-encrusted Pads (6461, -0.67 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Caretaker's Cape (20428, -0.13 DPS) [rep]; Feyscale Cloak (6632, -0.33 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.13 DPS) [crafted]; Bloody Apron (6226, -0.13 DPS) [dungeon]; Green Woolen Vest (2582, -1.55 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Silver-lined Bracers (3224, -0.13 DPS) [world]; Seer's Cuffs (3645, -0.13 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.36 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.40 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.67 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (35.9 DPS) | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.53 DPS, sim-verified) [crafted]; Lesser Belt of the Spire (1299, -0.54 DPS) [world] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Colorful Kilt (10048, -0.54 DPS) [crafted]; Silk-threaded Trousers (1929, -1.04 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.40 DPS) [crafted]; Pristine Boots (253889, -0.54 DPS) [crafted]; Feather Padded Treads (285345, -0.81 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.27 DPS) [world]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Ring of the Shadow (1462, -0.67 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.56 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Ring of the Shadow (1462, -0.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Staff of Westfall (2042), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), Keen Skinning Knife (269716), Riverpaw Mystic Staff (1391), Defias Mage Staff (1928), Skyseeker's Greatstaff (263937), Oakthrush Staff (15397), Felweaver's Staff (284287), Apprentice's Spellstaff (248008), Makeshift Stormcaller (281274), Blackfang (2236), Staff of the Blessed Seer (2271), Nightbane Staff (3227), Deadly Bronze Poniard (3490), Shadowfang (1482), Assassin's Blade (1935), Bluegill Kukri (2046), Bronze Shortsword (2850), Decapitating Sword (3740), Cruel Barb (5191), Spikelash Dagger (6333), Edward's Knife (251485), Chol'aruk's Trophy Sword (276886), Plaguefang (279876), Staff of Horrors (880), Buzz Saw (1937), Pearl-handled Dagger (5540), Living Root (6631), Butcher's Slicer (6633), Thief's Blade (5192), Emberstone Staff (5201), Militant Shortsword (15211), Scimitar of Atun (1469), Gnarled Hermit's Staff (1539), Buzzer Blade (2169), Brackclaw (2235), Tail Spike (6448), Night Watch Shortsword (935), Thornblade (2908), Hook Dagger (3184), Big Bronze Knife (3848), Staff of Nobles (3902), Ironpatch Blade (12976), Venom Web Fang (899), Riverside Staff (1473), Blackwater Cutlass (1951), Solid Shortblade (2074), Scrimshaw Dagger (2089), Medicine Staff (4575), Goblin Screwdriver (1936), Hollowfang Blade (2020), Northern Shortsword (2078), Slicer Blade (820), Foamspittle Staff (1405), Daryl's Shortsword (3572), War Knife (4571), Ritual Blade (5112), Redridge Machete (1219), Defias Rapier (1925), Raider Shortsword (15210), Daggerfang's Daggerfang (281257), Giant Tarantula Fang (1287), Staff of Conjuring (1933), Long Crawler Limb (2088), Curved Dagger (2632), Enamelled Broadsword (4765), Pale Skinner (5744), Briarsteel Shortsword (15335), Curvewood Dagger (15396), Craftsman's Dagger (2218), Stonesplinter Dagger (2266), Sturdy Quarterstaff (4566), Feral Blade (4766), Balanced Fighting Stick (6215), Balanced Quarterstaff (257343), Quickblade's Dagger (257346), Gnarlpine War Staff (284166), Jeweled Dagger (1917), Small Green Dagger (4302), Militia Sword (248006), Militia Shortblade (248007), Tim's Lost Rib (282064), Bayne's Bite (283462), Pristine Orcish Dagger (286730), Ukta's Conduit (286754), Small Hand Blade (816), Carving Knife (2140), Bloodstained Knife (3225), Copper Dagger (7166), Hoof-Holstered Blade (277988), Helm Splitter (286755), Notched Shortsword (727), Copper Shortsword (2847), Gritroot Staff (9603), Brushwood Blade (18957), Webwood Slicer (282284), Chipped Spellstaff (285238), Explorer's Shortsword (285239)) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 | yes | Staff of Westfall (2042, +0.00 DPS) [quest]; Gnarled Necromancer's Staff (251534, +0.00 DPS) [quest]; Twisted Chanter's Staff (890, -0.84 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Pulsating Hydra Heart (5183, -0.67 DPS) [world]; Tear of Grief (5611, -0.67 DPS) [quest]; Grayson's Torch (1172, -1.18 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 169.3 | yes | Skycaller (12984, -1.67 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 128, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10047 Simple Kilt; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 54.7. Weights run: 0.9s. Verify run: 0.7s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.288), intellect=not significant (0.400 ± 0.325), crit=0.845 ± 0.045, hit=2.738 ± 0.165, spell_haste=1.692 ± 0.302, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.288), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Silk Headband (7050, -0.18 DPS) [crafted]; Embalmed Shroud (7691, -0.27 DPS) [dungeon]; Enchanter's Cowl (4322, -0.80 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 | yes | Crystal Starfire Medallion (5003, -0.71 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.60 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 | yes | Fairywing Mantle (9536, -0.27 DPS) [quest]; Invoker's Mantle (215365, -0.33 DPS) [crafted]; Death Speaker Mantle (6685, -0.66 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Prelacy Cape (7004, -0.09 DPS) [quest]; Repairman's Cape (9605, -0.37 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.2 | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.26 DPS) [dungeon]; Pristine Gown (253961, -0.40 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.60 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.60 DPS) [quest]; Glowing Magical Bracelets (13106, -2.10 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (54.6 DPS) | yes | Shilly Mitts (9609, +0.00 DPS) [quest]; Truefaith Gloves (7049, -0.07 DPS) [crafted]; Town Clerk's Mittens (270029, -0.66 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.2 | yes | Invoker's Cord (215366, -0.29 DPS) [crafted]; Crimson Silk Belt (7055, -0.31 DPS) [crafted]; Belt of Arugal (6392, -0.47 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | sim-verified (54.6 DPS) | yes | Pristine Leggings (253987, -0.20 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.33 DPS) [crafted]; Abomination Skin Leggings (23173, -0.64 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.8 | yes | Acidic Walkers (9454, -0.15 DPS) [dungeon]; Nimbus Boots (6998, -0.35 DPS) [quest]; Spidersilk Boots (4320, -1.65 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.11 DPS) [quest]; Lorekeeper's Ring (20431, -0.18 DPS) [rep]; Electrocutioner Lagnut (9447, -0.37 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.27 DPS) [dungeon]; Sludge-Stained Band (286535, -0.27 DPS) [world]; Minor Channeling Ring (1449, -1.62 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (53.6 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, +0.00 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.4 | yes | Twisted Chanter's Staff (890, -0.01 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.04 DPS) [quest]; Channeler's Staff (4437, -0.11 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 365.7 | yes | Starfaller (13063, -0.64 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 105.0. Weights run: 0.8s. Verify run: 0.7s. 295 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.308), intellect=not significant (0.667 ± 0.396), crit=0.836 ± 0.047, hit=3.138 ± 0.219, spell_haste=1.901 ± 0.341, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.308), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -0.91 DPS) [crafted]; Holy Shroud (2721, -1.09 DPS) [world_drop]; Augural Shroud (2620, -2.22 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Darkspear Warding Pendant (272074, -0.69 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.84 DPS) [vendor]; Necklace of Calisea (1714, -1.51 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.7 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.22 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.22 DPS) [vendor]; Icy Cloak (4327, -0.25 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 | yes | Robe of Power (7054, -0.44 DPS) [crafted]; Elemental Raiment (9434, -0.54 DPS) [world_drop]; Dreamweave Vest (10021, -0.80 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.22 DPS) [quest]; Windchaser Cuffs (14429, -0.33 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 | yes | Black Mageweave Gloves (10003, -0.62 DPS) [crafted]; Gilded Handwraps (254021, -0.87 DPS) [crafted]; Red Mageweave Gloves (10018, -1.70 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (105.0 DPS) | yes | Gilded Cord (254037, -0.36 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.40 DPS) [rep]; Deathmage Sash (10771, -1.39 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Abomination Skin Leggings (23173, -0.84 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.05 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.60 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -1.49 DPS) [dungeon]; Spidersilk Boots (4320, -1.56 DPS) [crafted]; Gilded Slippers (254001, -2.66 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 | yes | Ring of Forlorn Spirits (2043, -0.65 DPS) [quest]; Reedknot Ring (9622, -0.76 DPS) [quest]; Minor Channeling Ring (1449, -0.84 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.22 DPS) [quest]; Lorekeeper's Ring (19525, -0.22 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.48 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 38.7 | yes | Illusionary Rod (7713, -1.04 DPS, sim-verified) [dungeon]; Staff of Noh'Orahil (15105, -2.65 DPS) [quest]; Windweaver Staff (7757, -3.13 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 365.6 | yes | Umbral Wand (5216, -0.16 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -5.16 DPS) [crafted]; Ember Wand (5215, -6.15 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 295, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 136.3. Weights run: 0.8s. Verify run: 0.7s. 376 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.208, intellect=not significant (-0.672 ± 0.236), crit=0.535 ± 0.036, hit=1.703 ± 0.134, spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Dreamweave Circlet (10041, -1.68 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -3.21 DPS, sim-verified) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.60 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.96 DPS) [world_drop]; Pendant of Myzrael (4614, -1.96 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 15.5 | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.54 DPS) [crafted]; Bloodmage Mantle (7684, -1.82 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Runecloth Cloak (13860, -1.40 DPS) [crafted]; Nightfall Drape (12465, -1.91 DPS, sim-verified) [dungeon]; Icy Cloak (4327, -1.96 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 | yes | Elemental Raiment (9434, -0.14 DPS, sim-verified) [world_drop]; Knight's Dreadweave Vest (220886, -0.70 DPS) [vendor]; Acumen Robes (17775, -0.84 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, -0.15 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest]; Nethergeld Cuffs (254061, -0.56 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Black Mageweave Gloves (10003, -0.76 DPS, sim-verified) [crafted]; Brightcloth Gloves (14101, -1.40 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.40 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 16.5 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.70 DPS) [rep]; Ghostweave Cord (254073, -0.70 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 19.5 | yes | Wizardweave Leggings (14132, +0.00 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -1.54 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.10 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (134.7 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -3.04 DPS, sim-verified) [vendor]; Black Mageweave Boots (10026, -3.64 DPS) [crafted]; Gilded Sandals (254107, -3.64 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 17.0 | yes | Lorekeeper's Ring (19523, -1.41 DPS) [rep]; Philanthropist's Ring (281635, -1.97 DPS) [quest]; Lorekeeper's Ring (19524, -2.25 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Lorekeeper's Ring (19523, -0.31 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.84 DPS) [quest]; Lorekeeper's Ring (19524, -1.12 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (131.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.32 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -3.36 DPS) [quest] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (131.3 DPS) | yes | Soul Harvester (20536, +0.00 DPS) [quest]; Shortsword of Vengeance (754, -1.54 DPS, sim-verified) [world_drop]; Illusionary Rod (7713, -2.67 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (133.4 DPS) | yes | Pyric Caduceus (11748, -1.67 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Guardian Talisman; trinket2: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 376, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 275.3. Weights run: 0.8s. Verify run: 0.7s. 777 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.620, intellect=4.547 ± -0.743, crit=-1.774 ± -0.098, hit=-5.572 ± -0.332, spell_haste=-2.242 ± -0.730, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.083 ± -0.620, fire_power=-0.082 ± -0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 147.4 | yes | Field Marshal's Coronal (17578, +0.00 DPS) [vendor]; Field Marshal's Coronal (231584, +0.00 DPS) [vendor]; Gnomish Turban of Psychic Might (21517, -3.12 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (274.8 DPS) | yes | Horizon Choker (13085, +0.00 DPS) [world_drop]; Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Blazefury Medallion (17111, -5.43 DPS, sim-verified) [world] |
| shoulder | Darkspear Shoulderpads (272103) | Creeg Bothunk [vendor] | 134.8 | yes | Field Marshal's Dreadweave Shoulders (17580, +0.00 DPS) [vendor]; Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [vendor]; Magister's Mantle (16689, -1.83 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | 72.8 | yes | Hide of the Wild (18510, +0.00 DPS) [crafted]; Arcanoweave Cloak (272411, +0.00 DPS) [vendor]; Darkspear Raider's Cloak (272076, -0.51 DPS, sim-verified) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 142.3 | yes | Field Marshal's Dreadweave Robe (17581, +0.00 DPS) [vendor]; Field Marshal's Dreadweave Robe (231582, +0.00 DPS) [vendor]; Magister's Robes (16688, -15.84 DPS, sim-verified) [dungeon] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 85.1 | yes | Deathmist Bracers (22071, +0.00 DPS) [quest]; Marshal's Dreadweave Cuffs (17582, -0.24 DPS) [pvp]; Magiskull Cuffs (13107, -11.55 DPS, sim-verified) [world_drop] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 130.8 | yes | Mooncloth Gloves (18409, +0.00 DPS) [crafted]; Heretic Gloves (240140, +0.00 DPS) [vendor]; Raider Handwraps (272098, -1.31 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 125.0 | yes | Marshal's Dreadweave Sash (17585, +0.00 DPS) [pvp]; Heretic Belt (240144, +0.00 DPS) [vendor]; Magister's Belt (16685, -8.36 DPS, sim-verified) [dungeon] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 132.3 | yes | Marshal's Dreadweave Leggings (17579, +0.00 DPS) [vendor]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (22752, -10.34 DPS, sim-verified) [rep] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 99.7 | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (17583, +0.00 DPS) [vendor]; Bloodvine Boots (19684, -7.29 DPS, sim-verified) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (274.8 DPS) | yes | Cauterizing Band (19140, +0.00 DPS) [world_drop]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Wrath of Cenarius (21190, -3.52 DPS, sim-verified) [quest] |
| finger2 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | sim-verified (274.8 DPS) | yes | Cauterizing Band (19140, +0.00 DPS) [world_drop]; Signet Ring of the Bronze Dragonflight (21210, +0.00 DPS) [quest]; Wrath of Cenarius (21190, -5.50 DPS, sim-verified) [quest] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (274.8 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, -3.26 DPS, sim-verified) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (274.8 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, -5.52 DPS, sim-verified) [quest] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (274.8 DPS) | yes | Grand Marshal's Stave (18873, -6.99 DPS) [vendor]; Grand Marshal's Stave (234571, -8.88 DPS) [vendor]; Shortsword of Vengeance (754, -10.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 40.8 | yes | Cairnstone Sliver (9654, +0.00 DPS) [quest]; Charged Lightning Rod (11860, +0.00 DPS) [quest]; Jaina's Firestarter (13064, -8.36 DPS, sim-verified) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Amulet of the Dawn; shoulder: Darkspear Shoulderpads; back: Darkspear Raider's Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Weakness Analyzer; trinket2: Serenity Field; main_hand: Crackling Staff; ranged: Brilliant Wand

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 35.1. Weights run: 0.8s. Verify run: 0.8s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.103, intellect=not significant (-0.337 ± 0.116), crit=0.435 ± 0.021, hit=1.467 ± 0.069, spell_haste=0.633 ± 0.114, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.103, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.81 DPS) [crafted]; Lucky Fishing Hat (19972, -0.81 DPS) [quest]; Flying Tiger Goggles (4368, -2.53 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | sim-verified (34.2 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.13 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.37 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.67 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Battle Healer's Cloak (20427, -0.13 DPS) [rep]; Feyscale Cloak (6632, -0.25 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.13 DPS) [crafted]; Bloody Apron (6226, -0.13 DPS) [dungeon]; Green Woolen Vest (2582, -2.09 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) (or Windsong Bangles (263336)) | Earthen Arise [quest] | 1.0 | yes | Mindthrust Bracers (1974, -0.13 DPS) [dungeon]; Silver-lined Bracers (3224, -0.13 DPS) [world]; Windsong Bangles (263336, -0.47 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.28 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.40 DPS) [quest]; Pristine Gloves (253913, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (34.7 DPS) | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.54 DPS) [world]; Novice Arcanist's Sash (253885, -0.84 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Colorful Kilt (10048, -0.54 DPS) [crafted]; Silk-threaded Trousers (1929, -2.22 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.40 DPS) [crafted]; Pristine Boots (253889, -0.54 DPS) [crafted]; Feather Padded Treads (285345, -1.12 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Ring of the Shadow (1462, -0.67 DPS) [world]; Ring of Scorn (3235, -0.67 DPS) [quest] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, -0.15 DPS, sim-verified) [dungeon]; Ring of the Shadow (1462, -0.40 DPS) [world]; Ring of Scorn (3235, -0.40 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Clear Crystal Rod (16894), Lesser Staff of the Spire (1300), Keen Skinning Knife (269716), Riverpaw Mystic Staff (1391), Staff of Orgrimmar (15444), Defias Mage Staff (1928), Skyseeker's Greatstaff (263937), Felweaver's Staff (284287), Apprentice's Spellstaff (248008), Makeshift Stormcaller (281274), Blackfang (2236), Staff of the Blessed Seer (2271), Nightbane Staff (3227), Deadly Bronze Poniard (3490), Shadowfang (1482), Assassin's Blade (1935), Bluegill Kukri (2046), Bronze Shortsword (2850), Decapitating Sword (3740), Cruel Barb (5191), Spikelash Dagger (6333), Wingblade (6504), Crescent Staff (6505), Edward's Knife (251485), Chol'aruk's Trophy Sword (276886), Plaguefang (279876), Staff of Horrors (880), Buzz Saw (1937), Pearl-handled Dagger (5540), Living Root (6631), Butcher's Slicer (6633), Ceranium Rod (3452), Thief's Blade (5192), Emberstone Staff (5201), Militant Shortsword (15211), Scimitar of Atun (1469), Gnarled Hermit's Staff (1539), Buzzer Blade (2169), Brackclaw (2235), Tail Spike (6448), Claystone Shortsword (16891), Night Watch Shortsword (935), Hook Dagger (3184), Big Bronze Knife (3848), Staff of Nobles (3902), Harpy Skinner (5279), Wind Rider Staff (5306), Elegant Shortsword (5321), Ironpatch Blade (12976), Venom Web Fang (899), Riverside Staff (1473), Blackwater Cutlass (1951), Medicine Staff (4575), Goblin Screwdriver (1936), Hollowfang Blade (2020), Northern Shortsword (2078), Serrated Knife (3581), Cursed Felblade (14145), Chanting Blade (14151), Kris of Orgrimmar (15443), Slicer Blade (820), Foamspittle Staff (1405), War Knife (4571), Ritual Blade (5112), Redridge Machete (1219), Defias Rapier (1925), Raider Shortsword (15210), Daggerfang's Daggerfang (281257), Giant Tarantula Fang (1287), Staff of Conjuring (1933), Long Crawler Limb (2088), Cauldron Stirrer (5340), Curved Dagger (2632), Enamelled Broadsword (4765), Pale Skinner (5744), Stonesplinter Dagger (2266), Darkwood Staff (3446), Sturdy Quarterstaff (4566), Feral Blade (4766), Balanced Quarterstaff (257343), Quickblade's Dagger (257346), Gnarlpine War Staff (284166), Jeweled Dagger (1917), Small Green Dagger (4302), Compact Fighting Knife (4974), Militia Sword (248006), Militia Shortblade (248007), Tim's Lost Rib (282064), Bayne's Bite (283462), Pristine Orcish Dagger (286730), Ukta's Conduit (286754), Small Hand Blade (816), Carving Knife (2140), Bloodstained Knife (3225), Jagged Dagger (4947), Copper Dagger (7166), Hoof-Holstered Blade (277988), Helm Splitter (286755), Notched Shortsword (727), Copper Shortsword (2847), Webwood Slicer (282284), Chipped Spellstaff (285238), Explorer's Shortsword (285239)) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS) [quest]; Quilboar Toothpick (276888, +0.00 DPS) [quest]; Twisted Chanter's Staff (890, -0.27 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Nightglow Concoction (3451, -0.67 DPS) [quest]; Pulsating Hydra Heart (5183, -0.67 DPS) [world]; Grayson's Torch (1172, -0.97 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 169.3 | yes | Skycaller (12984, -1.60 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance; 209620 Insignia of the Horde; 241089 Scarlet Dagger

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 52.8. Weights run: 0.9s. Verify run: 0.7s. 214 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.288), intellect=not significant (0.400 ± 0.325), crit=0.845 ± 0.045, hit=2.738 ± 0.165, spell_haste=1.692 ± 0.302, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.288), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Silk Headband (7050, -0.18 DPS) [crafted]; Embalmed Shroud (7691, -0.27 DPS) [dungeon]; Enchanter's Cowl (4322, -0.51 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 | yes | Crystal Starfire Medallion (5003, -0.71 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.47 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 | yes | Fairywing Mantle (9536, -0.27 DPS) [quest]; Invoker's Mantle (215365, -0.33 DPS) [crafted]; Death Speaker Mantle (6685, -0.57 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, -0.02 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Battle Healer's Cloak (19529, -0.09 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.2 | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.26 DPS) [dungeon]; Pristine Gown (253961, -0.40 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.60 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.60 DPS) [quest]; Glowing Magical Bracelets (13106, -1.62 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.0 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.16 DPS) [crafted]; Gnoll Casting Gloves (892, -0.18 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.2 | yes | Warsong Sash (16975, -0.14 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Invoker's Cord (215366, -0.29 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.35 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.8 | yes | Acidic Walkers (9454, -0.15 DPS) [dungeon]; Boots of the Enchanter (4325, -0.44 DPS) [crafted]; Spidersilk Boots (4320, -1.43 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.18 DPS) [rep]; Electrocutioner Lagnut (9447, -0.37 DPS) [dungeon]; Sludge-Stained Band (286535, -0.37 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.27 DPS) [world]; Black Widow Band (6199, -0.29 DPS) [world]; Electrocutioner Lagnut (9447, -1.83 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.4 | yes | Gnarled Necromancer's Staff (251534, -0.04 DPS) [quest]; Channeler's Staff (4437, -0.11 DPS) [world]; Twisted Chanter's Staff (890, -0.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 365.7 | yes | Starfaller (13063, -0.34 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Relentless Raider's Seal; trinket2: Rune of Perfection; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 214, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 102.6. Weights run: 0.8s. Verify run: 0.7s. 291 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.308), intellect=not significant (0.667 ± 0.396), crit=0.836 ± 0.047, hit=3.138 ± 0.219, spell_haste=1.901 ± 0.341, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.308), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -0.91 DPS) [crafted]; Holy Shroud (2721, -1.09 DPS) [world_drop]; Augural Shroud (2620, -1.29 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Darkspear Warding Pendant (272074, -0.69 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.84 DPS) [vendor]; Necklace of Calisea (1714, -1.25 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.7 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.22 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.22 DPS) [vendor]; Icy Cloak (4327, -0.25 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 | yes | Robe of Power (7054, -0.44 DPS) [crafted]; Elemental Raiment (9434, -0.54 DPS) [world_drop]; Dreamweave Vest (10021, -0.78 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (101.6 DPS) | yes | Condor Bracers (15864, -0.22 DPS) [quest]; Windchaser Cuffs (14429, -0.33 DPS) [world_drop]; Radiant Silver Bracers (4545, -1.07 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 | yes | Black Mageweave Gloves (10003, -0.62 DPS) [crafted]; Gilded Handwraps (254021, -0.87 DPS) [crafted]; Red Mageweave Gloves (10018, -0.98 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (102.5 DPS) | yes | Gilded Cord (254037, -0.36 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.40 DPS) [rep]; Deathmage Sash (10771, -2.00 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Crimson Silk Pantaloons (7062, -0.64 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.84 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.05 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -1.49 DPS) [dungeon]; Gilded Slippers (254001, -1.55 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 | yes | Reedknot Ring (9622, -0.76 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.87 DPS) [vendor]; Ogremind Ring (1993, -1.02 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Reedknot Ring (9622, +0.00 DPS, sim-verified) [quest]; Advisor's Ring (19521, -0.22 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 38.7 | yes | Illusionary Rod (7713, -0.82 DPS, sim-verified) [dungeon]; Staff of Noh'Orahil (15105, -2.65 DPS) [quest]; Windweaver Staff (7757, -3.13 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 365.6 | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -5.16 DPS) [crafted]; Ember Wand (5215, -6.15 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 291, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10047 Simple Kilt; 10049 Diabolist's Blade

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 136.3. Weights run: 0.8s. Verify run: 0.7s. 372 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.208, intellect=not significant (-0.672 ± 0.236), crit=0.535 ± 0.036, hit=1.703 ± 0.134, spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Dreamweave Circlet (10041, -1.68 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.74 DPS, sim-verified) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.35 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.96 DPS) [world_drop]; Choker of the High Shaman (4112, -1.96 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (131.6 DPS) | yes | Black Mageweave Shoulders (10027, -0.84 DPS) [crafted]; Bloodmage Mantle (7684, -1.12 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.49 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Deep Woodlands Cloak (19121, +0.00 DPS, sim-verified) [quest]; Nightfall Drape (12465, -1.40 DPS) [dungeon]; Runecloth Cloak (13860, -1.40 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 | yes | Elemental Raiment (9434, -0.66 DPS, sim-verified) [world_drop]; Stone Guard's Dreadweave Vest (220904, -0.70 DPS) [vendor]; Acumen Robes (17775, -0.84 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, -0.56 DPS) [crafted]; Condor Bracers (15864, -0.95 DPS, sim-verified) [quest]; Bloodband Bracers (11469, -1.12 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Black Mageweave Gloves (10003, -0.98 DPS, sim-verified) [crafted]; Brightcloth Gloves (14101, -1.40 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.40 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | sim-verified (131.5 DPS) | yes | Defiler's Cloth Girdle (20166, +0.00 DPS) [rep]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Defiler's Cloth Girdle (20165, -1.40 DPS, sim-verified) [rep] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 19.5 | yes | Wizardweave Leggings (14132, +0.00 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -1.54 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.10 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (132.7 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.54 DPS, sim-verified) [vendor]; Black Mageweave Boots (10026, -3.64 DPS) [crafted]; Gilded Sandals (254107, -3.64 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 17.0 | yes | Advisor's Ring (19519, -1.41 DPS) [rep]; Philanthropist's Ring (281635, -1.97 DPS) [quest]; Advisor's Ring (19520, -2.25 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Advisor's Ring (19519, -0.38 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.84 DPS) [quest]; Advisor's Ring (19520, -1.12 DPS) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (131.0 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.63 DPS, sim-verified) [crafted]; Tidal Charm (1404, -1.68 DPS) [vendor] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (131.0 DPS) | yes | Rune of the Guard Captain (19120, -0.02 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.66 DPS, sim-verified) [crafted]; Tidal Charm (1404, -3.36 DPS) [vendor] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (131.0 DPS) | yes | Soul Harvester (20536, +0.00 DPS) [quest]; Shortsword of Vengeance (754, -1.53 DPS, sim-verified) [world_drop]; Illusionary Rod (7713, -2.67 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 187.3 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.51 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Stone Guard's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Abyss Shard; ranged: Pyric Caduceus

No-known-source sample (15 of 372, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 263.9. Weights run: 0.8s. Verify run: 0.7s. 774 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.620, intellect=4.547 ± -0.743, crit=-1.774 ± -0.098, hit=-5.572 ± -0.332, spell_haste=-2.242 ± -0.730, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.083 ± -0.620, fire_power=-0.082 ± -0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 147.4 | yes | Warlord's Dreadweave Hood (17591, +0.00 DPS) [vendor]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [vendor]; Gnomish Turban of Psychic Might (21517, -1.82 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (264.5 DPS) | yes | Horizon Choker (13085, +0.00 DPS) [world_drop]; Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Blazefury Medallion (17111, -4.26 DPS, sim-verified) [world] |
| shoulder | Darkspear Shoulderpads (272103) | Creeg Bothunk [vendor] | 134.8 | yes | Warlord's Dreadweave Mantle (17590, +0.00 DPS) [vendor]; Warlord's Dreadweave Mantle (231592, +0.00 DPS) [vendor]; Magister's Mantle (16689, -1.85 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | 72.8 | yes | Hide of the Wild (18510, +0.00 DPS) [crafted]; Deep Woodlands Cloak (19121, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272076, -0.65 DPS, sim-verified) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 142.3 | yes | Warlord's Dreadweave Robe (17592, +0.00 DPS) [vendor]; Warlord's Dreadweave Robe (231591, +0.00 DPS) [vendor]; Magister's Robes (16688, -17.79 DPS, sim-verified) [dungeon] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 85.1 | yes | Deathmist Bracers (22071, +0.00 DPS) [quest]; General's Dreadweave Bracers (17587, -0.24 DPS) [pvp]; Magiskull Cuffs (13107, -9.69 DPS, sim-verified) [world_drop] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 130.8 | yes | Mooncloth Gloves (18409, +0.00 DPS) [crafted]; Heretic Gloves (240140, +0.00 DPS) [vendor]; Raider Handwraps (272098, -0.36 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 125.0 | yes | General's Dreadweave Belt (17589, +0.00 DPS) [pvp]; Heretic Belt (240144, +0.00 DPS) [vendor]; Magister's Belt (16685, -8.73 DPS, sim-verified) [dungeon] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 132.3 | yes | General's Dreadweave Pants (17593, +0.00 DPS) [vendor]; General's Dreadweave Pants (231588, +0.00 DPS) [vendor]; Outrider's Silk Leggings (22747, -11.20 DPS, sim-verified) [rep] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 99.7 | yes | General's Dreadweave Boots (17586, +0.00 DPS) [vendor]; General's Dreadweave Boots (231593, +0.00 DPS) [vendor]; Bloodvine Boots (19684, -5.27 DPS, sim-verified) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (264.5 DPS) | yes | Cauterizing Band (19140, +0.00 DPS) [world_drop]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Wrath of Cenarius (21190, -3.27 DPS, sim-verified) [quest] |
| finger2 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | sim-verified (264.5 DPS) | yes | Cauterizing Band (19140, +0.00 DPS) [world_drop]; Signet Ring of the Bronze Dragonflight (21210, +0.00 DPS) [quest]; Wrath of Cenarius (21190, -5.76 DPS, sim-verified) [quest] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (264.5 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, -3.11 DPS, sim-verified) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (264.5 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, -5.07 DPS, sim-verified) [quest] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-verified (264.5 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Shortsword of Vengeance (754, -3.85 DPS, sim-verified) [world_drop]; High Warlord's War Staff (234549, -13.61 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 40.8 | yes | Charged Lightning Rod (11860, +0.00 DPS) [quest]; Nature's Breath (19118, +0.00 DPS) [quest]; Jaina's Firestarter (13064, -7.77 DPS, sim-verified) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Amulet of the Dawn; shoulder: Darkspear Shoulderpads; back: Darkspear Raider's Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Weakness Analyzer; trinket2: Serenity Field; main_hand: Staff of Hale Magefire; ranged: Brilliant Wand

No-known-source sample (15 of 774, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

