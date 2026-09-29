# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.6. Weights run: 1.4s. Verify run: 1.1s. 253 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.127, intellect=0.422 ± -0.128, crit=-0.462 ± -0.020, hit=-1.259 ± -0.070, spell_haste=-0.173 ± -0.159, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.000 ± -0.127

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -3.9) [crafted]; Flying Tiger Goggles (4368, -6.0) [crafted]; Lucky Fishing Hat (19972, -6.0) [quest] |
| neck | Sentinel's Medallion (20444) (or Tarnished Locket (279870)) | Silverwing Sentinels [rep] | 0.0 | yes | Tarnished Locket (279870, +0.0) [quest] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.7 | yes | Double-Stitched Woolen Shoulders (4314, -2.7) [crafted]; Slime-encrusted Pads (6461, -6.7) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, -0.7) [crafted]; Feyscale Cloak (6632, -1.0) [dungeon]; Black Whelp Cloak (7283, -1.0) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.1 | yes | Gray Woolen Robe (2585, -1.0) [crafted]; Green Woolen Robe (6243, -2.8) [crafted]; Green Woolen Vest (2582, -3.1) [crafted] |
| wrist | Tabitha's Cuffs (251486) | Quests [quest] | 2.5 | yes | Mindthrust Bracers (1974, -0.4) [dungeon]; Bright Bracers (3647, -0.8) [dungeon]; Seer's Cuffs (3645, -2.1) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -1.0) [world]; Pristine Gloves (253913, -1.7) [crafted]; Blight Gloves (279877, -4.0) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.7 | yes | Novice Arcanist's Sash (253885, -0.4) [crafted]; Keller's Girdle (2911, -2.3) [dungeon]; Novice Ardent's Sash (253887, -2.4) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.5 | yes | Abomination Skin Leggings (23173, +3.8) [dungeon]; Silk-threaded Trousers (1929, -1.5) [dungeon]; Colorful Kilt (10048, -3.5) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 | yes | Feather Padded Treads (285345, -3.7) [world]; Pristine Boots (253889, -4.4) [crafted]; Red Woolen Boots (4313, -4.7) [crafted] |
| finger1 | Minor Channeling Ring (1449) | Quests [quest] | 5.8 | yes | Sludge-Stained Band (286535, -2.8) [world]; Lavishly Jeweled Ring (1156, -3.3) [dungeon]; Black Pearl Ring (6332, -5.0) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -2.0) [world]; Lavishly Jeweled Ring (1156, -2.5) [dungeon]; Black Pearl Ring (6332, -4.2) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 4.2 | yes | Gnarled Necromancer's Staff (251534, +0.0) [quest]; Channeler's Staff (4437, -0.8) [world]; Lesser Staff of the Spire (1300, -1.7) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 3.0 | yes | Sizzle Stick (8071, +2.0) [quest]; Sable Wand (7607, +0.0) [quest]; Torchlight Wand (5240, -1.0) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 253, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.5. Weights run: 1.3s. Verify run: 1.2s. 475 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.261), intellect=not significant (-0.193 ± 0.270), crit=0.624 ± 0.029, hit=1.712 ± 0.125, spell_haste=-1.617 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.261)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | Razorfen Downs: Withered Battle Boar [dungeon] | 11.0 | yes | Silk Headband (7050, -2.0) [crafted]; Embalmed Shroud (7691, -3.0) [dungeon]; Filigreed Pristine Circlet (253975, -3.0) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Pendant of Myzrael (4614, -7.0) [dungeon]; Glowing Green Talisman (5002, -7.0) [dungeon]; Crystal Starfire Medallion (5003, -7.0) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 | yes | Moonlit Amice (11884, -2.0) [quest]; Invoker's Mantle (215365, -2.0) [crafted]; Death Speaker Mantle (6685, -3.0) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, -1.0) [crafted]; Prelacy Cape (7004, -1.0) [quest]; Caretaker's Cape (19533, -1.0) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, -4.0) [crafted]; Robes of Arcana (5770, -5.0) [crafted]; Death Speaker Robes (6682, -6.0) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Mindthrust Bracers (1974, -9.0) [dungeon]; Silver-lined Bracers (3224, -9.0) [world]; Seer's Cuffs (3645, -9.0) [dungeon] |
| hands | Serpent Gloves (5970) (or Shilly Mitts (9609)) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Shilly Mitts (9609, +0.0) [quest]; Gnoll Casting Gloves (892, -1.0) [world]; Gloves of Meditation (4318, -2.0) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.0 | yes | Belt of Arugal (6392, -2.0) [dungeon]; Ghamoo-ra's Bind (6908, -3.0) [dungeon]; Invoker's Cord (215366, -4.0) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -3.0) [dungeon]; Silk-threaded Trousers (1929, -5.0) [dungeon]; Pristine Leggings (253987, -5.0) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.0 | yes | Spidersilk Boots (4320, +0.0) [crafted]; Nimbus Boots (6998, -1.0) [quest]; Boots of the Enchanter (4325, -2.0) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -2.0) [quest]; Lorekeeper's Ring (20431, -2.0) [rep]; Electrocutioner Lagnut (9447, -4.0) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Minor Channeling Ring (1449, -1.0) [quest]; Electrocutioner Lagnut (9447, -3.0) [dungeon]; Sludge-Stained Band (286535, -3.0) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Notched Shortsword (727) (or Dragonmaw Shortsword (753), Gnarled Ash Staff (791), Small Hand Blade (816), Slicer Blade (820), Staff of Horrors (880), Black Metal Shortsword (886), Twisted Chanter's Staff (890), Venom Web Fang (899), Night Watch Shortsword (935), Rod of the Sleepwalker (1155), Redridge Machete (1219), Giant Tarantula Fang (1287), Lesser Staff of the Spire (1300), Hardened Root Staff (1317), Riverpaw Mystic Staff (1391), Foamspittle Staff (1405), Scimitar of Atun (1469), Riverside Staff (1473), Shadowfang (1482), Witching Stave (1484), Heavy Marauder Scimitar (1493), Gnarled Hermit's Staff (1539), Jeweled Dagger (1917), Defias Rapier (1925), Defias Mage Staff (1928), Staff of Conjuring (1933), Assassin's Blade (1935), Goblin Screwdriver (1936), Buzz Saw (1937), Blackwater Cutlass (1951), Bloodscalp Channeling Staff (1998), Twisted Sabre (2011), Cryptbone Staff (2013), Skeletal Longsword (2018), Hollowfang Blade (2020), Sword of the Night Sky (2035), Staff of Westfall (2042), Bluegill Kukri (2046), Dwarven Magestaff (2072), Solid Shortblade (2074), Magician Staff (2077), Northern Shortsword (2078), Long Crawler Limb (2088), Scrimshaw Dagger (2089), Carving Knife (2140), Buzzer Blade (2169), Craftsman's Dagger (2218), Ogremage Staff (2226), Brackclaw (2235), Blackfang (2236), Phytoblade (2263), Stonesplinter Dagger (2266), Staff of the Blessed Seer (2271), Kam's Walking Stick (2280), Staff of the Shade (2549), Evocator's Blade (2567), Curved Dagger (2632), Cross Dagger (2819), Copper Shortsword (2847), Bronze Shortsword (2850), Thornblade (2908), Claw of the Shadowmancer (2912), Prison Shank (2941), Icicle Rod (2950), Hook Dagger (3184), Acrobatic Staff (3185), Viking Sword (3186), Bloodstained Knife (3225), Nightbane Staff (3227), Flesh Piercer (3336), Lucine Longsword (3400), Doomspike (3413), Staff of the Friar (3415), Deadly Bronze Poniard (3490), Daryl's Shortsword (3572), Decapitating Sword (3740), Big Bronze Knife (3848), Hardened Iron Shortsword (3849), Jade Serpentblade (3850), Staff of Nobles (3902), Small Green Dagger (4302), Channeler's Staff (4437), Blackvenom Blade (4446), Naraxis' Fang (4449), Talon of Vultros (4454), Sturdy Quarterstaff (4566), War Knife (4571), Medicine Staff (4575), Enamelled Broadsword (4765), Feral Blade (4766), Ritual Blade (5112), Cruel Barb (5191), Thief's Blade (5192), Emberstone Staff (5201), Pearl-handled Dagger (5540), Staff of the Purifier (5613), Relic Blade (5627), Pale Skinner (5744), Wyvern Tailspike (5752), Balanced Fighting Stick (6215), Meteor Shard (6220), Odo's Ley Staff (6318), Spikelash Dagger (6333), Tail Spike (6448), Living Root (6631), Butcher's Slicer (6633), Thornspike (6681), Swinetusk Shank (6691), Bite of Serra'kis (6904), Copper Dagger (7166), Torturing Poker (7682), Electrocutioner Leg (9446), Hydrocane (9452), Toxic Revenger (9453), Gritroot Staff (9603), Darkwater Talwar (11121), Broad Bladed Knife (12247), Daring Dirk (12248), Staff of Soran'ruk (15109), Briarsteel Shortsword (15335), Curvewood Dagger (15396), Oakthrush Staff (15397), Brushwood Blade (18957), Lorekeeper's Staff (212580), Protector's Sword (212582), Sentinel's Blade (212583), Advisor's Gnarled Staff (212584), Legionnaire's Sword (212586), Scout's Blade (212587), Glimmering Staff (249392), Soulstaff (249393), Edward's Knife (251485), Gnarled Necromancer's Staff (251534), Skyseeker's Greatstaff (263937), Darkspear Insurgent's Spellblade (272086), Chol'aruk's Trophy Sword (276886), Quilboar Toothpick (276888), Plaguefang (279876), Webwood Slicer (282284), Bayne's Bite (283462), Gnarlpine War Staff (284166), Felweaver's Staff (284287), Chipped Spellstaff (285238), Explorer's Shortsword (285239), Alliance Outrunner's Sword (285346), Alliance Outrunner Staff (285350), Pristine Orcish Dagger (286730), Ukta's Conduit (286754), Helm Splitter (286755)) | Al'Aketh Honor Guard [world] | 0.0 | yes | Dragonmaw Shortsword (753, +0.0) [world]; Gnarled Ash Staff (791, +0.0) [dungeon]; Small Hand Blade (816, +0.0) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.0 | yes | Dwarven Tome (279898, -2.0) [quest]; Eye of Paleth (2943, -3.0) [quest]; Orb of Souls (249395, -3.0) [crafted] |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | Gnomeregan: Dark Iron Agent [dungeon] | 5.0 | yes | Sizzle Stick (8071, +0.0) [quest]; Greater Mystic Wand (217287, +0.0) [crafted]; Spellcrafter Wand (6677, -1.0) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Notched Shortsword; off_hand: Orb of Mystic Insight; ranged: Wand of Eventide

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 108.5. Weights run: 1.1s. Verify run: 1.1s. 641 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.358), intellect=not significant (0.698 ± 0.495), crit=1.096 ± 0.054, hit=3.563 ± 0.219, spell_haste=-2.263 ± 0.479, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.358)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -3.0) [world]; Enchanter's Cowl (4322, -8.0) [crafted]; Holy Shroud (2721, -10.0) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -6.3) [dungeon]; Darkspear Warding Pendant (272074, -6.3) [vendor]; Darkspear Warding Pendant (272075, -7.7) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Green Silken Shoulders (7057, -0.4) [crafted]; Bloodmage Mantle (7684, -0.8) [dungeon]; Berylline Pads (4197, -2.1) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.0) [crafted]; Darkspear Raider's Cloak (272077, -1.8) [vendor]; Icy Cloak (4327, -2.5) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 26.2 | yes | Dreamweave Vest (10021, -1.9) [crafted]; Robe of Power (7054, -3.8) [crafted]; Crimson Silk Vest (7058, -7.2) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Quests [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.0) [dungeon]; Condor Bracers (15864, -2.0) [quest]; Aurora Bracers (4043, -3.4) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 | yes | Red Mageweave Gloves (10018, -2.8) [crafted]; Black Mageweave Gloves (10003, -5.8) [crafted]; Gilded Handwraps (254021, -7.9) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.8 | yes | Deathmage Sash (10771, +0.7) [dungeon]; Gilded Cord (254037, -3.2) [crafted]; Highlander's Cloth Girdle (20099, -3.7) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Crimson Silk Pantaloons (7062, -5.3) [crafted]; Abomination Skin Leggings (23173, -7.8) [dungeon]; Stoneweaver Leggings (9407, -9.8) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -12.1) [crafted]; Acidic Walkers (9454, -13.4) [dungeon]; Spidersilk Boots (4320, -14.2) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -2.0) [quest]; Lorekeeper's Ring (19525, -2.0) [rep]; Minor Channeling Ring (1449, -2.6) [quest] |
| finger2 | Ring of Forlorn Spirits (2043) | Quests [quest] | 8.0 | yes | Reedknot Ring (9622, -1.0) [quest]; Minor Channeling Ring (1449, -1.6) [quest]; Sea Giant's Toe Ring (274746, -2.0) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Staff of Dar'Orahil (15106) | Quests [quest] | 43.3 | yes | Illusionary Rod (7713, -23.8) [dungeon]; Staff of Noh'Orahil (15105, -25.2) [quest]; Windweaver Staff (7757, -32.8) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fizzle's Zippy Lighter (6729) | Quests [quest] | 6.1 | yes | Burning Sliver (5249, -0.1) [quest]; Twisted Nether Wand (249144, -0.1) [crafted]; Wand of Eventide (5214, -1.1) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Dar'Orahil; ranged: Fizzle's Zippy Lighter

No-known-source sample (15 of 641, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 178.9. Weights run: 1.3s. Verify run: 1.2s. 834 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.222, intellect=not significant (0.052 ± 0.247), crit=0.455 ± 0.023, hit=2.181 ± 0.152, spell_haste=-9.055 ± 0.409, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Dreamweave Circlet (10041, -5.5) [crafted]; Spellpower Goggles Xtreme (10502, -6.0) [crafted]; Red Mageweave Headband (10033, -7.0) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 | yes | Mindburst Medallion (11196, -1.0) [quest]; Horizon Choker (13085, -6.6) [world]; Darkspear Warding Pendant (272073, -6.8) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.9 | yes | Black Mageweave Shoulders (10027, -3.5) [crafted]; Bloodmage Mantle (7684, -4.5) [dungeon]; Nethergeld Shoulders (254049, -4.5) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 | yes | Runecloth Cloak (13860, -4.9) [crafted]; Nightfall Drape (12465, -5.3) [world]; Icy Cloak (4327, -7.3) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 22.3 | yes | Acumen Robes (17775, -2.3) [quest]; Dreamweave Vest (10021, -3.8) [crafted]; Brightcloth Robe (14100, -4.3) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Quests [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.0) [dungeon]; Nethergeld Cuffs (254061, -1.6) [crafted]; Condor Bracers (15864, -2.0) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 | yes | Black Mageweave Gloves (10003, -3.2) [crafted]; Brightcloth Gloves (14101, -5.2) [crafted]; Runecloth Gloves (13863, -5.7) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.5 | yes | Highlander's Cloth Girdle (20097, +1.1) [rep]; Highlander's Cloth Girdle (20098, -0.3) [rep]; Ghostweave Cord (254073, -0.5) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 | yes | Red Mageweave Pants (10009, -4.4) [crafted]; Senior Designer's Pantaloons (11841, -6.6) [dungeon]; Gaze Dreamer Pants (6903, -7.0) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Black Mageweave Boots (10026, -12.6) [crafted]; Earthenweave Boots (254093, -15.0) [crafted]; Southsea Mojo Boots (20641, -15.4) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.8 | yes | Ring of Forlorn Spirits (2043, -13.8) [quest]; Reedknot Ring (9622, -14.8) [quest]; Sea Giant's Toe Ring (274746, -15.8) [vendor] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Lorekeeper's Ring (19524, -3.0) [rep]; Ring of Forlorn Spirits (2043, -4.0) [quest]; Reedknot Ring (9622, -5.0) [quest] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Abyss Shard (20534, +12.0) [quest]; Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Abyss Shard (20534, +6.0) [quest]; Thunderbrew's Boot Flask (744, -6.0) [quest]; Tidal Charm (1404, -6.0) [vendor] |
| main_hand | Soul Harvester (20536) | Quests [quest] | 22.6 | yes | Staff of Dar'Orahil (15106, -0.3) [quest]; Kindling Stave (11750, -15.6) [dungeon]; Illusionary Rod (7713, -16.0) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Dreambough Wand (249234) | Enchanting [crafted] | 7.0 | yes | Lesser Eternal Wand (249232, +1.0) [crafted]; Burning Sliver (5249, -1.0) [quest]; Twisted Nether Wand (249144, -1.0) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Smoking Heart of the Mountain; trinket2: Uther's Strength; main_hand: Soul Harvester; ranged: Dreambough Wand

No-known-source sample (15 of 834, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 427.2. Weights run: 1.3s. Verify run: 1.2s. 1200 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.837), intellect=not significant (-1.254 ± 0.970), crit=1.999 ± 0.106, hit=6.401 ± 0.516, spell_haste=not significant (0.241 ± 0.963), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.837)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Quests [quest] | 153.0 | yes | Bloodvine Goggles (19999, +3.0) [crafted]; Doomcaller's Circlet (21337, -28.0) [quest]; Deathmist Mask (226909, -71.0) [quest] |
| neck | Onyxia Tooth Pendant (18404) | Quests [quest] | 92.0 | yes | Fury of the Forgotten Swarm (21809, +0.0) [raid]; Gem of Trapped Innocents (23057, -21.0) [raid]; Prestor's Talisman of Connivery (19377, -28.0) [raid] |
| shoulder | Plagueheart Shoulderpads (22507) | Quests [quest] | 100.0 | yes | Doomcaller's Mantle (21335, -8.0) [quest]; Mantle of the Timbermaw (19050, -55.0) [crafted]; Champion's Dreadweave Spaulders (23256, -60.0) [vendor] |
| back | Earthweave Cloak (21187) | Quests [quest] | 64.0 | yes | Chromatic Cloak (18509, -36.0) [crafted]; Shroud of Unspoken Names (21418, -46.0) [quest]; Spritecaster Cape (11623, -50.0) [dungeon] |
| chest | Plagueheart Robe (22504) | Quests [quest] | 143.0 | yes | Bloodvine Vest (19682, +12.0) [crafted]; Zandalar Demoniac's Robe (20033, -52.0) [quest]; Doomcaller's Robes (21334, -74.0) [quest] |
| wrist | Rockfury Bracers (21186) | Quests [quest] | 91.0 | yes | Burrower Bracers (21611, -63.0) [raid]; Plagueheart Bindings (22511, -68.0) [quest]; Dryad's Wrist Bindings (19595, -69.0) [rep] |
| hands | Dark Storm Gauntlets (21585) | Ahn'Qiraj [raid] | 101.0 | yes | Deathmist Wraps (22077, -24.0) [quest]; Deathmist Wraps (226911, -24.0) [quest]; Gloves of Spell Mastery (14146, -36.0) [crafted] |
| waist | Plagueheart Belt (22510) | Quests [quest] | 62.0 | yes | Belt of the Archmage (18405, -14.0) [crafted]; Highlander's Cloth Girdle (20047, -20.0) [rep]; Highlander's Cloth Girdle (20097, -25.0) [rep] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 101.0 | yes | Deathmist Leggings (226910, -22.0) [quest]; Plagueheart Leggings (22505, -36.0) [quest]; Magister's Leggings (16687, -36.4) [dungeon] |
| feet | Plagueheart Sandals (22508) | Quests [quest] | 60.0 | yes | Bloodvine Boots (19684, +23.0) [crafted]; Boots of Epiphany (21600, -26.0) [raid]; Doomcaller's Footwraps (21338, -32.0) [quest] |
| finger1 | Seal of the Damned (23025) | Naxxramas [raid] | 113.0 | yes | Ring of Unspoken Names (21417, -7.0) [quest]; Band of the Inevitable (23031, -13.0) [raid]; Don Julio's Band (19325, -21.0) [rep] |
| finger2 | Ring of the Fallen God (21709) | Quests [quest] | 101.0 | yes | Ring of Unspoken Names (21417, +5.0) [quest]; Band of the Inevitable (23031, -1.0) [raid]; Don Julio's Band (19325, -9.0) [rep] |
| trinket1 | The Restrained Essence of Sapphiron (23046) | Naxxramas [raid] | 40.0 | yes | Neltharion's Tear (19379, +132.0) [raid]; Drake Fang Talisman (19406, +88.0) [raid]; Kiss of the Spider (22954, +52.0) [raid] |
| trinket2 | Eye of the Dead (23047) | Naxxramas [raid] | 0.0 | yes | Neltharion's Tear (19379, +172.0) [raid]; Drake Fang Talisman (19406, +128.0) [raid]; Kiss of the Spider (22954, +92.0) [raid] |
| main_hand | Atiesh, Greatstaff of the Guardian (22589) | Quests [quest] | 278.0 | yes | Atiesh, Greatstaff of the Guardian (22630, -72.1) [quest]; Blessed Qiraji Acolyte Staff (21273, -122.0) [quest]; High Warlord's War Staff (234549, -141.1) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Dreambough Wand (249234, -1.0) [crafted]; Burning Sliver (5249, -2.0) [quest]; Twisted Nether Wand (249144, -2.0) [crafted] |

**New at 60:** head: Plagueheart Circlet; neck: Onyxia Tooth Pendant; shoulder: Plagueheart Shoulderpads; back: Earthweave Cloak; chest: Plagueheart Robe; wrist: Rockfury Bracers; hands: Dark Storm Gauntlets; waist: Plagueheart Belt; legs: Bloodvine Leggings; feet: Plagueheart Sandals; finger1: Seal of the Damned; finger2: Ring of the Fallen God; trinket1: The Restrained Essence of Sapphiron; trinket2: Eye of the Dead; main_hand: Atiesh, Greatstaff of the Guardian; ranged: Lesser Eternal Wand

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.5. Weights run: 1.4s. Verify run: 1.1s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.127, intellect=0.422 ± -0.128, crit=-0.462 ± -0.020, hit=-1.259 ± -0.070, spell_haste=-0.173 ± -0.159, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.000 ± -0.127

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -3.9) [crafted]; Flying Tiger Goggles (4368, -6.0) [crafted]; Lucky Fishing Hat (19972, -6.0) [quest] |
| neck | Scout's Medallion (20442) (or Tarnished Locket (279870)) | Warsong Outriders [rep] | 0.0 | yes | Tarnished Locket (279870, +0.0) [quest] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.7 | yes | Double-Stitched Woolen Shoulders (4314, -2.7) [crafted]; Slime-encrusted Pads (6461, -6.7) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, -0.7) [crafted]; Feyscale Cloak (6632, -1.0) [dungeon]; Black Whelp Cloak (7283, -1.0) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.1 | yes | Gray Woolen Robe (2585, -1.0) [crafted]; Green Woolen Robe (6243, -2.8) [crafted]; Green Woolen Vest (2582, -3.1) [crafted] |
| wrist | Tabitha's Cuffs (251486) | Quests [quest] | 2.5 | yes | Mindthrust Bracers (1974, -0.4) [dungeon]; Featherbead Bracers (15452, -0.4) [quest]; Owlbeard Bracers (16981, -0.7) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -1.0) [world]; Pristine Gloves (253913, -1.7) [crafted]; Apothecary Gloves (10919, -3.0) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.7 | yes | Novice Arcanist's Sash (253885, -0.4) [crafted]; Keller's Girdle (2911, -2.3) [dungeon]; Novice Ardent's Sash (253887, -2.4) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.5 | yes | Abomination Skin Leggings (23173, +3.8) [dungeon]; Silk-threaded Trousers (1929, -1.5) [dungeon]; Colorful Kilt (10048, -3.5) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 | yes | Feather Padded Treads (285345, -3.7) [world]; Pristine Boots (253889, -4.4) [crafted]; Red Woolen Boots (4313, -4.7) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -2.5) [dungeon]; Black Pearl Ring (6332, -4.2) [world]; The 1 Ring (8350, -4.6) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, -0.5) [dungeon]; Black Pearl Ring (6332, -2.2) [world]; The 1 Ring (8350, -2.6) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 4.2 | yes | Gnarled Necromancer's Staff (251534, +0.0) [quest]; Channeler's Staff (4437, -0.8) [world]; Lesser Staff of the Spire (1300, -1.7) [world] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Quests [quest] | 5.0 | yes | Cookie's Stirring Rod (5198, -2.0) [dungeon]; Torchlight Wand (5240, -3.0) [quest]; Greater Magic Wand (11288, -3.0) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 58.9. Weights run: 1.3s. Verify run: 1.3s. 470 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.261), intellect=not significant (-0.193 ± 0.270), crit=0.624 ± 0.029, hit=1.712 ± 0.125, spell_haste=-1.617 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.261)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | Razorfen Downs: Withered Battle Boar [dungeon] | 11.0 | yes | Silk Headband (7050, -2.0) [crafted]; Embalmed Shroud (7691, -3.0) [dungeon]; Filigreed Pristine Circlet (253975, -3.0) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Pendant of Myzrael (4614, -7.0) [dungeon]; Glowing Green Talisman (5002, -7.0) [dungeon]; Crystal Starfire Medallion (5003, -7.0) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 | yes | Chestnut Mantle (17695, -1.0) [quest]; Invoker's Mantle (215365, -2.0) [crafted]; Death Speaker Mantle (6685, -3.0) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.0) [quest]; Heavy Woolen Cloak (4311, -1.0) [crafted]; Battle Healer's Cloak (19529, -1.0) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, -4.0) [crafted]; High Robe of the Adjudicator (3461, -5.0) [quest]; Robes of Arcana (5770, -5.0) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Owlbeard Bracers (16981, -8.0) [quest]; Mindthrust Bracers (1974, -9.0) [dungeon]; Silver-lined Bracers (3224, -9.0) [world] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -1.0) [world]; Jutebraid Gloves (10654, -1.0) [quest]; Gloves of Meditation (4318, -2.0) [crafted] |
| waist | Warsong Sash (16975) (or Defiler's Cloth Girdle (20164)) | Quests [quest] | 11.0 | yes | Defiler's Cloth Girdle (20164, +0.0) [rep]; Belt of Arugal (6392, -2.0) [dungeon]; Ghamoo-ra's Bind (6908, -3.0) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -3.0) [dungeon]; Silk-threaded Trousers (1929, -5.0) [dungeon]; Pristine Leggings (253987, -5.0) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.0 | yes | Spidersilk Boots (4320, +0.0) [crafted]; Boots of the Enchanter (4325, -2.0) [crafted]; Acidic Walkers (9454, -2.0) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -2.0) [rep]; Electrocutioner Lagnut (9447, -4.0) [dungeon]; Sludge-Stained Band (286535, -4.0) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -3.0) [dungeon]; Sludge-Stained Band (286535, -3.0) [world]; Sacred Band (6669, -4.0) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Notched Shortsword (727) (or Dragonmaw Shortsword (753), Gnarled Ash Staff (791), Small Hand Blade (816), Slicer Blade (820), Staff of Horrors (880), Black Metal Shortsword (886), Twisted Chanter's Staff (890), Venom Web Fang (899), Night Watch Shortsword (935), Rod of the Sleepwalker (1155), Redridge Machete (1219), Giant Tarantula Fang (1287), Lesser Staff of the Spire (1300), Riverpaw Mystic Staff (1391), Foamspittle Staff (1405), Scimitar of Atun (1469), Riverside Staff (1473), Shadowfang (1482), Witching Stave (1484), Heavy Marauder Scimitar (1493), Gnarled Hermit's Staff (1539), Jeweled Dagger (1917), Defias Rapier (1925), Defias Mage Staff (1928), Staff of Conjuring (1933), Assassin's Blade (1935), Goblin Screwdriver (1936), Buzz Saw (1937), Blackwater Cutlass (1951), Bloodscalp Channeling Staff (1998), Twisted Sabre (2011), Cryptbone Staff (2013), Skeletal Longsword (2018), Hollowfang Blade (2020), Sword of the Night Sky (2035), Bluegill Kukri (2046), Dwarven Magestaff (2072), Magician Staff (2077), Northern Shortsword (2078), Long Crawler Limb (2088), Carving Knife (2140), Buzzer Blade (2169), Ogremage Staff (2226), Brackclaw (2235), Blackfang (2236), Stonesplinter Dagger (2266), Staff of the Blessed Seer (2271), Kam's Walking Stick (2280), Staff of the Shade (2549), Evocator's Blade (2567), Curved Dagger (2632), Cross Dagger (2819), Copper Shortsword (2847), Bronze Shortsword (2850), Claw of the Shadowmancer (2912), Prison Shank (2941), Hook Dagger (3184), Acrobatic Staff (3185), Viking Sword (3186), Bloodstained Knife (3225), Nightbane Staff (3227), Flesh Piercer (3336), Doomspike (3413), Staff of the Friar (3415), Darkwood Staff (3446), Ceranium Rod (3452), Talonstrike (3462), Deadly Bronze Poniard (3490), Serrated Knife (3581), Decapitating Sword (3740), Big Bronze Knife (3848), Hardened Iron Shortsword (3849), Jade Serpentblade (3850), Staff of Nobles (3902), Small Green Dagger (4302), Channeler's Staff (4437), Blackvenom Blade (4446), Naraxis' Fang (4449), Talon of Vultros (4454), Sturdy Quarterstaff (4566), War Knife (4571), Medicine Staff (4575), Enamelled Broadsword (4765), Feral Blade (4766), Jagged Dagger (4947), Compact Fighting Knife (4974), Ritual Blade (5112), Cruel Barb (5191), Thief's Blade (5192), Emberstone Staff (5201), Harpy Skinner (5279), Wind Rider Staff (5306), Elegant Shortsword (5321), Cauldron Stirrer (5340), Pearl-handled Dagger (5540), Pale Skinner (5744), Wyvern Tailspike (5752), Meteor Shard (6220), Odo's Ley Staff (6318), Spikelash Dagger (6333), Tail Spike (6448), Wingblade (6504), Crescent Staff (6505), Living Root (6631), Butcher's Slicer (6633), Thornspike (6681), Swinetusk Shank (6691), Bite of Serra'kis (6904), Copper Dagger (7166), Torturing Poker (7682), Electrocutioner Leg (9446), Hydrocane (9452), Toxic Revenger (9453), Darkwater Talwar (11121), Broad Bladed Knife (12247), Daring Dirk (12248), Cursed Felblade (14145), Chanting Blade (14151), Staff of Soran'ruk (15109), Kris of Orgrimmar (15443), Staff of Orgrimmar (15444), Polished Walking Staff (16889), Slatemetal Cutlass (16890), Claystone Shortsword (16891), Clear Crystal Rod (16894), Lorekeeper's Staff (212580), Protector's Sword (212582), Sentinel's Blade (212583), Advisor's Gnarled Staff (212584), Legionnaire's Sword (212586), Scout's Blade (212587), Glimmering Staff (249392), Soulstaff (249393), Edward's Knife (251485), Gnarled Necromancer's Staff (251534), Skyseeker's Greatstaff (263937), Darkspear Insurgent's Spellblade (272086), Chol'aruk's Trophy Sword (276886), Quilboar Toothpick (276888), Plaguefang (279876), Webwood Slicer (282284), Bayne's Bite (283462), Gnarlpine War Staff (284166), Felweaver's Staff (284287), Chipped Spellstaff (285238), Explorer's Shortsword (285239), Alliance Outrunner's Sword (285346), Alliance Outrunner Staff (285350), Pristine Orcish Dagger (286730), Ukta's Conduit (286754), Helm Splitter (286755)) | Al'Aketh Honor Guard [world] | 0.0 | yes | Dragonmaw Shortsword (753, +0.0) [world]; Gnarled Ash Staff (791, +0.0) [dungeon]; Small Hand Blade (816, +0.0) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.0 | yes | Dwarven Tome (279898, -2.0) [quest]; Orb of Souls (249395, -3.0) [crafted]; Alliance Outrunner Healing Rod (285348, -3.0) [world] |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | Gnomeregan: Dark Iron Agent [dungeon] | 5.0 | yes | Sizzle Stick (8071, +0.0) [quest]; Greater Mystic Wand (217287, +0.0) [crafted]; Cookie's Stirring Rod (5198, -2.0) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Notched Shortsword; off_hand: Orb of Mystic Insight; ranged: Wand of Eventide

No-known-source sample (15 of 470, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 105.8. Weights run: 1.1s. Verify run: 1.1s. 636 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.358), intellect=not significant (0.698 ± 0.495), crit=1.096 ± 0.054, hit=3.563 ± 0.219, spell_haste=-2.263 ± 0.479, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.358)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -3.0) [world]; Enchanter's Cowl (4322, -8.0) [crafted]; Holy Shroud (2721, -10.0) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -6.3) [dungeon]; Darkspear Warding Pendant (272074, -6.3) [vendor]; Darkspear Warding Pendant (272075, -7.7) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Green Silken Shoulders (7057, -0.4) [crafted]; Bloodmage Mantle (7684, -0.8) [dungeon]; Berylline Pads (4197, -2.1) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.0) [crafted]; Darkspear Raider's Cloak (272077, -1.8) [vendor]; Icy Cloak (4327, -2.5) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 26.2 | yes | Dreamweave Vest (10021, -1.9) [crafted]; Robe of Power (7054, -3.8) [crafted]; Crimson Silk Vest (7058, -7.2) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Quests [quest] | 9.6 | yes | Spidertank Oilrag (9448, -0.6) [dungeon]; Condor Bracers (15864, -2.6) [quest]; Aurora Bracers (4043, -4.0) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 | yes | Red Mageweave Gloves (10018, -2.8) [crafted]; Black Mageweave Gloves (10003, -5.8) [crafted]; Gilded Handwraps (254021, -7.9) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.8 | yes | Deathmage Sash (10771, +0.7) [dungeon]; Gilded Cord (254037, -3.2) [crafted]; Defiler's Cloth Girdle (20164, -3.7) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Crimson Silk Pantaloons (7062, -5.3) [crafted]; Abomination Skin Leggings (23173, -7.8) [dungeon]; Stoneweaver Leggings (9407, -9.8) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -12.1) [crafted]; Acidic Walkers (9454, -13.4) [dungeon]; Spidersilk Boots (4320, -14.2) [crafted] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Reedknot Ring (9622, -2.0) [quest]; Advisor's Ring (19521, -2.0) [rep]; Advisor's Ring (20426, -4.0) [rep] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Reedknot Ring (9622, +1.0) [quest]; Ogremind Ring (1993, -1.1) [world]; Voodoo Band (1996, -1.1) [world] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Staff of Dar'Orahil (15106) | Quests [quest] | 43.3 | yes | Illusionary Rod (7713, -23.8) [dungeon]; Staff of Noh'Orahil (15105, -25.2) [quest]; Windweaver Staff (7757, -32.8) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fizzle's Zippy Lighter (6729) | Quests [quest] | 6.1 | yes | Twisted Nether Wand (249144, -0.1) [crafted]; Wand of Eventide (5214, -1.1) [dungeon]; Sizzle Stick (8071, -1.1) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Dar'Orahil; ranged: Fizzle's Zippy Lighter

No-known-source sample (15 of 636, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 173.6. Weights run: 1.3s. Verify run: 1.1s. 829 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.222, intellect=not significant (0.052 ± 0.247), crit=0.455 ± 0.023, hit=2.181 ± 0.152, spell_haste=-9.055 ± 0.409, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Dreamweave Circlet (10041, -5.5) [crafted]; Spellpower Goggles Xtreme (10502, -6.0) [crafted]; Red Mageweave Headband (10033, -7.0) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 | yes | Mindburst Medallion (11196, -1.0) [quest]; Horizon Choker (13085, -6.6) [world]; Darkspear Warding Pendant (272073, -6.8) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.9 | yes | Black Mageweave Shoulders (10027, -3.5) [crafted]; Bloodmage Mantle (7684, -4.5) [dungeon]; Nethergeld Shoulders (254049, -4.5) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 | yes | Deep Woodlands Cloak (19121, -1.8) [quest]; Runecloth Cloak (13860, -4.9) [crafted]; Nightfall Drape (12465, -5.3) [world] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 22.3 | yes | Acumen Robes (17775, -2.3) [quest]; Dreamweave Vest (10021, -3.8) [crafted]; Brightcloth Robe (14100, -4.3) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, -1.6) [crafted]; Condor Bracers (15864, -2.0) [quest]; Bloodband Bracers (11469, -3.5) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 | yes | Black Mageweave Gloves (10003, -3.2) [crafted]; Brightcloth Gloves (14101, -5.2) [crafted]; Runecloth Gloves (13863, -5.7) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 15.6 | yes | Satyrmane Sash (17755, -1.1) [dungeon]; Defiler's Cloth Girdle (20166, -1.4) [rep]; Ghostweave Cord (254073, -1.6) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 | yes | Red Mageweave Pants (10009, -4.4) [crafted]; Senior Designer's Pantaloons (11841, -6.6) [dungeon]; Gaze Dreamer Pants (6903, -7.0) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Black Mageweave Boots (10026, -12.6) [crafted]; Earthenweave Boots (254093, -15.0) [crafted]; Southsea Mojo Boots (20641, -15.4) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.8 | yes | Reedknot Ring (9622, -14.8) [quest]; Sea Giant's Toe Ring (274746, -15.8) [vendor]; Electrocutioner Lagnut (9447, -18.8) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Advisor's Ring (19520, -3.0) [rep]; Reedknot Ring (9622, -5.0) [quest]; Advisor's Ring (19521, -5.0) [rep] |
| trinket1 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Rune of the Guard Captain (19120, +9.3) [quest]; Abyss Shard (20534, +6.0) [quest]; Tidal Charm (1404, -6.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of the Guard Captain (19120, +15.3) [quest]; Abyss Shard (20534, +12.0) [quest]; Tidal Charm (1404, +0.0) [vendor] |
| main_hand | Soul Harvester (20536) | Quests [quest] | 22.6 | yes | Staff of Dar'Orahil (15106, -0.3) [quest]; Kindling Stave (11750, -15.6) [dungeon]; Illusionary Rod (7713, -16.0) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Dreambough Wand (249234, -1.0) [crafted]; Twisted Nether Wand (249144, -2.0) [crafted]; Charged Lightning Rod (11860, -2.8) [quest] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Wizardweave Leggings; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Uther's Strength; trinket2: Darkspear Voodoo Seal; main_hand: Soul Harvester; ranged: Lesser Eternal Wand

No-known-source sample (15 of 829, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 426.0. Weights run: 1.3s. Verify run: 1.2s. 1195 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.837), intellect=not significant (-1.254 ± 0.970), crit=1.999 ± 0.106, hit=6.401 ± 0.516, spell_haste=not significant (0.241 ± 0.963), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.837)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Quests [quest] | 153.0 | yes | Bloodvine Goggles (19999, +3.0) [crafted]; Doomcaller's Circlet (21337, -28.0) [quest]; Deathmist Mask (226909, -71.0) [quest] |
| neck | Onyxia Tooth Pendant (18404) | Quests [quest] | 92.0 | yes | Fury of the Forgotten Swarm (21809, +0.0) [raid]; Gem of Trapped Innocents (23057, -21.0) [raid]; Prestor's Talisman of Connivery (19377, -28.0) [raid] |
| shoulder | Plagueheart Shoulderpads (22507) | Quests [quest] | 100.0 | yes | Doomcaller's Mantle (21335, -8.0) [quest]; Mantle of the Timbermaw (19050, -55.0) [crafted]; Champion's Dreadweave Spaulders (23256, -60.0) [vendor] |
| back | Earthweave Cloak (21187) | Quests [quest] | 64.0 | yes | Chromatic Cloak (18509, -36.0) [crafted]; Shroud of Unspoken Names (21418, -46.0) [quest]; Spritecaster Cape (11623, -50.0) [dungeon] |
| chest | Plagueheart Robe (22504) | Quests [quest] | 143.0 | yes | Bloodvine Vest (19682, +12.0) [crafted]; Zandalar Demoniac's Robe (20033, -52.0) [quest]; Doomcaller's Robes (21334, -74.0) [quest] |
| wrist | Rockfury Bracers (21186) | Quests [quest] | 91.0 | yes | Burrower Bracers (21611, -63.0) [raid]; Plagueheart Bindings (22511, -68.0) [quest]; Dryad's Wrist Bindings (19595, -69.0) [rep] |
| hands | Dark Storm Gauntlets (21585) | Ahn'Qiraj [raid] | 101.0 | yes | Deathmist Wraps (22077, -24.0) [quest]; Deathmist Wraps (226911, -24.0) [quest]; Gloves of Spell Mastery (14146, -36.0) [crafted] |
| waist | Plagueheart Belt (22510) | Quests [quest] | 62.0 | yes | Belt of the Archmage (18405, -14.0) [crafted]; Defiler's Cloth Girdle (20163, -20.0) [rep]; Defiler's Cloth Girdle (20165, -25.0) [rep] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 101.0 | yes | Deathmist Leggings (226910, -22.0) [quest]; Plagueheart Leggings (22505, -36.0) [quest]; Magister's Leggings (16687, -36.4) [dungeon] |
| feet | Plagueheart Sandals (22508) | Quests [quest] | 60.0 | yes | Bloodvine Boots (19684, +23.0) [crafted]; Boots of Epiphany (21600, -26.0) [raid]; Doomcaller's Footwraps (21338, -32.0) [quest] |
| finger1 | Seal of the Damned (23025) | Naxxramas [raid] | 113.0 | yes | Ring of Unspoken Names (21417, -7.0) [quest]; Band of the Inevitable (23031, -13.0) [raid]; Don Julio's Band (19325, -21.0) [rep] |
| finger2 | Ring of the Fallen God (21709) | Quests [quest] | 101.0 | yes | Ring of Unspoken Names (21417, +5.0) [quest]; Band of the Inevitable (23031, -1.0) [raid]; Don Julio's Band (19325, -9.0) [rep] |
| trinket1 | The Restrained Essence of Sapphiron (23046) | Naxxramas [raid] | 40.0 | yes | Neltharion's Tear (19379, +132.0) [raid]; Drake Fang Talisman (19406, +88.0) [raid]; Kiss of the Spider (22954, +52.0) [raid] |
| trinket2 | Eye of the Dead (23047) | Naxxramas [raid] | 0.0 | yes | Neltharion's Tear (19379, +172.0) [raid]; Drake Fang Talisman (19406, +128.0) [raid]; Kiss of the Spider (22954, +92.0) [raid] |
| main_hand | Atiesh, Greatstaff of the Guardian (22589) | Quests [quest] | 278.0 | yes | Atiesh, Greatstaff of the Guardian (22630, -72.1) [quest]; Blessed Qiraji Acolyte Staff (21273, -122.0) [quest]; High Warlord's War Staff (234549, -141.1) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Dreambough Wand (249234) | Enchanting [crafted] | 7.0 | yes | Lesser Eternal Wand (249232, +1.0) [crafted]; Twisted Nether Wand (249144, -1.0) [crafted]; Wand of Eventide (5214, -2.0) [dungeon] |

**New at 60:** head: Plagueheart Circlet; neck: Onyxia Tooth Pendant; shoulder: Plagueheart Shoulderpads; back: Earthweave Cloak; chest: Plagueheart Robe; wrist: Rockfury Bracers; hands: Dark Storm Gauntlets; waist: Plagueheart Belt; legs: Bloodvine Leggings; feet: Plagueheart Sandals; finger1: Seal of the Damned; finger2: Ring of the Fallen God; trinket1: The Restrained Essence of Sapphiron; trinket2: Eye of the Dead; main_hand: Atiesh, Greatstaff of the Guardian; ranged: Dreambough Wand

No-known-source sample (15 of 1195, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

