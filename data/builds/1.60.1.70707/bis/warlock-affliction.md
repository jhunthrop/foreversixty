# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 57.2. Weights run: 1.7s. Verify run: 1.2s. 161 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.068 ± 0.003, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.393 per %), hit=0.105 ± 0.004 per rating point (10 rating = 1%, 1.049 per %), spell_haste=not significant (0.043 ± 0.037), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.780 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.68 DPS) | yes | Red Winter Hat (21524, -3.77 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.40 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.28 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (1.12 DPS) | yes | Black Whelp Cloak (7283, -0.52 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Reinforced Linen Cape (2580, -0.84 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.40 DPS) | yes | Green Woolen Vest (2582, -0.28 DPS) [crafted]; Bloody Apron (6226, -0.28 DPS) [dungeon]; Gray Woolen Robe (2585, -1.78 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.28 DPS) | yes | Ivycloth Bracelets (9793, -0.31 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.96 DPS) | yes | Gnoll Casting Gloves (892, -0.44 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.84 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.56 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.36 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.28 DPS) [crafted]; Heavy Woolen Pants (4316, -0.84 DPS) [crafted]; Silk-threaded Trousers (1929, -1.67 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.96 DPS) | yes | Red Woolen Boots (4313, -1.00 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.12 DPS) [crafted]; Soft-soled Linen Boots (4312, -1.40 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.40 DPS) | yes | Ring of the Shadow (1462, -1.69 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.84 DPS) | yes | Ring of the Shadow (1462, -1.05 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 89 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.68 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (1.40 DPS) | yes | Bouquet of Red Roses (22206, -1.63 DPS, sim-verified) [dungeon] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 87.9 spell_power points (24.64 DPS) | yes | Cookie's Stirring Rod (5198, -0.63 DPS, sim-verified) [dungeon]; Sable Wand (7607, -1.57 DPS) [quest]; Skycaller (12984, -3.07 DPS) [world_drop] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 161, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 100.9. Weights run: 1.6s. Verify run: 1.1s. 298 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.016 ± 0.004, crit=0.047 ± 0.001 per rating point (14 rating = 1%, 0.659 per %), hit=0.152 ± 0.005 per rating point (10 rating = 1%, 1.518 per %), spell_haste=-0.292 ± 0.060, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.792 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (3.24 DPS) | yes | Filigreed Pristine Circlet (253975, -0.88 DPS) [crafted]; Enchanter's Cowl (4322, -1.43 DPS) [crafted]; Silk Headband (7050, -4.54 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.1 spell_power points (0.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.01 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.1 spell_power points (2.69 DPS) | yes | Moonlit Amice (11884, -0.63 DPS) [quest]; Death Speaker Mantle (6685, -0.87 DPS) [dungeon]; Invoker's Mantle (215365, -1.04 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.47 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.29 DPS) [crafted]; Prelacy Cape (7004, -0.29 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.83 DPS) | yes | Green Silk Armor (7065, -0.81 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.12 DPS) [dungeon]; Mechbuilder's Overalls (9508, -1.42 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.65 DPS) | yes | Glowing Magical Bracelets (13106, -2.61 DPS) [world_drop]; Nightsky Wristbands (6407, -2.62 DPS) [world_drop]; Windsong Bangles (263336, -2.93 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (2.06 DPS) | yes | Gnoll Casting Gloves (892, -0.29 DPS) [world]; Serpent Gloves (5970, -0.47 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.57 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.0 spell_power points (2.66 DPS) | yes | Invoker's Cord (215366, -0.58 DPS) [crafted]; Crimson Silk Belt (7055, -0.86 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.96 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.53 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.47 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.74 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.1 spell_power points (2.09 DPS) | yes | Glinteye Slippers (273024, -0.03 DPS) [dungeon]; Nimbus Boots (6998, -0.33 DPS) [quest]; Spidersilk Boots (4320, -2.74 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.1 spell_power points (1.79 DPS) | yes | Minor Channeling Ring (1449, -0.31 DPS) [quest]; Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Runemender's Seal (273031, -0.91 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.77 DPS) | yes | Minor Channeling Ring (1449, -0.41 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon]; Runemender's Seal (273031, -0.88 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (10.59 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -4.71 DPS) [quest]; Glimmering Staff (249392, -10.54 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 111.7 spell_power points (32.88 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.40 DPS) [crafted]; Gyromatic Icemaker (9489, -2.52 DPS) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 131.6. Weights run: 1.8s. Verify run: 1.3s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.171 ± 0.005, crit=0.083 ± 0.001 per rating point (14 rating = 1%, 1.166 per %), hit=0.201 ± 0.007 per rating point (10 rating = 1%, 2.006 per %), spell_haste=0.392 ± 0.079, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.815 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.57 DPS) | yes | Augural Shroud (2620, -2.60 DPS) [world]; Living Cowl (5608, -2.68 DPS, sim-verified) [world]; Holy Shroud (2721, -3.13 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.4 spell_power points (2.62 DPS) | yes | Triune Amulet (7722, -1.74 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.24 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.30 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (3.30 DPS) | yes | Green Silken Shoulders (7057, -0.21 DPS) [crafted]; Inquisitor's Shawl (19507, -0.41 DPS) [dungeon]; Berylline Pads (4197, -0.57 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.19 DPS) | yes | Guardian Cloak (5965, -0.05 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.63 DPS) [quest]; Long Silken Cloak (4326, -0.81 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (7.21 DPS) | yes | Dreamweave Vest (10021, -1.09 DPS) [crafted]; Elemental Raiment (9434, -1.49 DPS, sim-verified) [world_drop]; Robe of Power (7054, -2.18 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.82 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.63 DPS) [quest]; Earthen Silk Cuffs (254019, -1.56 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (5.85 DPS) | yes | Black Mageweave Gloves (10003, -1.70 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.87 DPS) [crafted]; Gilded Handwraps (254021, -2.97 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (4.07 DPS) | yes | Belt of Arugal (6392, -1.09 DPS) [dungeon]; Gilded Cord (254037, -1.14 DPS) [crafted]; Deathmage Sash (10771, -2.00 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.0 spell_power points (5.02 DPS) | yes | Gaze Dreamer Pants (6903, -1.14 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.82 DPS) [crafted]; Stoneweaver Leggings (9407, -2.40 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.51 DPS) | yes | Gilded Slippers (254001, -4.14 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.11 DPS) [crafted]; Glinteye Slippers (273024, -5.32 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.9 spell_power points (3.40 DPS) | yes | Reedknot Ring (9622, -1.21 DPS) [quest]; Snake Hoop (6750, -1.25 DPS) [quest]; Ring of Power Regulation (273030, -1.52 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (2.50 DPS) | yes | Reedknot Ring (9622, -0.34 DPS, sim-verified) [quest]; Snake Hoop (6750, -0.36 DPS) [quest]; Ring of Power Regulation (273030, -0.63 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (131.6 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.11 DPS) [dungeon]; Spine of the Devourer (273040, -2.93 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.38 DPS) | yes | Beacon of Hope (9393, -1.88 DPS) [dungeon]; Orb of Lorica (11262, -1.97 DPS) [quest]; Orb of Mystic Insight (249394, -2.43 DPS, sim-verified) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 126.8 spell_power points (39.70 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -3.82 DPS) [crafted]; Umbral Wand (5216, -4.03 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 208.9. Weights run: 1.8s. Verify run: 1.5s. 507 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.272 ± 0.008, crit=0.126 ± 0.002 per rating point (14 rating = 1%, 1.768 per %), hit=0.260 ± 0.009 per rating point (10 rating = 1%, 2.599 per %), spell_haste=not significant (0.230 ± 0.129), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.828 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.08 DPS) | yes | Dreamweave Circlet (10041, -1.10 DPS) [crafted]; Red Mageweave Headband (10033, -1.76 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.02 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.2 spell_power points (3.09 DPS) | yes | Mindburst Medallion (11196, -0.52 DPS) [quest]; Triune Amulet (7722, -0.96 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.81 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.9 spell_power points (6.02 DPS) | yes | Black Mageweave Shoulders (10027, -1.29 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.91 DPS) [vendor]; Bloodmage Mantle (7684, -2.17 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (5.26 DPS) | yes | Runecloth Cloak (13860, -1.75 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.23 DPS) [dungeon]; Big Voodoo Cloak (8216, -2.75 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.4 spell_power points (8.22 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.16 DPS) [world_drop]; Dreamweave Vest (10021, -1.34 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.03 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Bloodband Bracers (11469, -0.52 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (6.42 DPS) | yes | Black Mageweave Gloves (10003, -1.38 DPS) [crafted]; Runecloth Gloves (13863, -1.56 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.17 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.7 spell_power points (5.63 DPS) | yes | Ghostweave Cord (254073, -0.92 DPS) [crafted]; Star Belt (4329, -1.25 DPS) [crafted]; Gilded Waistcord (254081, -1.44 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.39 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.66 DPS) [vendor]; Red Mageweave Pants (10009, -0.73 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -1.71 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.07 DPS) | yes | Gilded Sandals (254107, -1.13 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.68 DPS) [vendor]; Black Mageweave Boots (10026, -3.73 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.37 DPS) | yes | Ring of Forlorn Spirits (2043, -1.68 DPS) [quest]; Snake Hoop (6750, -1.90 DPS) [quest]; Reedknot Ring (9622, -2.02 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.4 spell_power points (3.82 DPS) | yes | Snake Hoop (6750, -1.35 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.38 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.47 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (208.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (208.9 DPS) | yes | Mark of the Chosen (17774, -2.45 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (208.9 DPS) | yes | Spire of Hakkar (10844, -0.43 DPS) [world]; Spellshifter Rod (9527, -1.01 DPS) [quest]; Hanzo Sword (8190, -29.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 156.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -4.06 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 315.2. Weights run: 1.9s. Verify run: 3.9s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.124 ± 0.010, crit=0.178 ± 0.003 per rating point (14 rating = 1%, 2.499 per %), hit=0.439 ± 0.014 per rating point (10 rating = 1%, 4.391 per %), spell_haste=not significant (0.051 ± 0.166), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.829 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.47 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.44 DPS) [pvp]; Dreamweave Circlet (10041, -2.06 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.2 spell_power points (4.65 DPS) | yes | Glowing Eye of Mordresh (10769, -1.84 DPS) [dungeon]; Triune Amulet (7722, -2.06 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -2.28 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (8.77 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.65 DPS, sim-verified) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.12 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.7 spell_power points (5.17 DPS) | yes | Nightfall Drape (12465, -2.02 DPS) [dungeon]; Runecloth Cloak (13860, -2.40 DPS, sim-verified) [crafted]; Icy Cloak (4327, -2.72 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (16.53 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.26 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -6.89 DPS) [pvp]; Robe of the Magi (1716, -10.16 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (315.2 DPS) | yes | Arcane Runed Bracers (4744, +0.00 DPS) [quest]; Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Runecloth Cuffs (254123, -6.94 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (315.2 DPS) | yes | Knight-Lieutenant's Dreadweave Handwraps (227100, +0.00 DPS) [pvp]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -4.79 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (315.2 DPS) | yes | Star Belt (4329, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -6.53 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.12 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.95 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.42 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.70 DPS) [crafted]; Omnicast Boots (11822, -0.88 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.9 spell_power points (6.62 DPS) | yes | Philanthropist's Ring (281635, -2.89 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.81 DPS) [quest]; Reedknot Ring (9622, -4.17 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.56 DPS) | yes | Philanthropist's Ring (281635, -1.21 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.75 DPS) [quest]; Reedknot Ring (9622, -2.11 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (315.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (315.2 DPS) | yes | Mark of the Chosen (17774, -4.74 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (315.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.18 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -34.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 224.8 spell_power points (78.87 DPS) | yes | Brilliant Wand (249385, -2.29 DPS, sim-verified) [crafted]; Stormrager (16997, -15.96 DPS) [quest]; Torch of Austen (13004, -22.08 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25520100135001050-0005200000000000000-0550004000000000)

Set DPS (verified): 599.4. Weights run: 2.0s. Verify run: 3.1s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.140 ± 0.019, crit=0.400 ± 0.010 per rating point (14 rating = 1%, 5.597 per %), hit=0.678 ± 0.033 per rating point (10 rating = 1%, 6.780 per %), spell_haste=not significant (0.492 ± 0.364), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.870 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (14.70 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.50 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.4 spell_power points (7.29 DPS) | yes | Glowing Eye of Mordresh (10769, -2.87 DPS) [dungeon]; Triune Amulet (7722, -2.94 DPS) [dungeon]; Mindburst Medallion (11196, -3.57 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (13.61 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.04 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.8 spell_power points (8.08 DPS) | yes | Nightfall Drape (12465, -3.18 DPS) [dungeon]; Runecloth Cloak (13860, -3.41 DPS, sim-verified) [crafted]; Icy Cloak (4327, -4.27 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.3 spell_power points (25.72 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.48 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.59 DPS) [pvp]; Robe of the Magi (1716, -14.28 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.4 spell_power points (6.75 DPS) | yes | Spidertank Oilrag (9448, -1.85 DPS) [dungeon]; Earthenweave Cuffs (254131, -2.39 DPS) [crafted]; Arcane Runed Bracers (4744, -3.30 DPS, sim-verified) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (11.97 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS, sim-verified) [crafted]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.24 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.4 spell_power points (8.38 DPS) | yes | Star Belt (4329, -1.31 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.33 DPS) [crafted]; Ghostweave Cord (254073, -3.03 DPS, sim-verified) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (14.15 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -3.37 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.06 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.09 DPS) [crafted]; Omnicast Boots (11822, -1.26 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.0 spell_power points (10.33 DPS) | yes | Philanthropist's Ring (281635, -4.51 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.98 DPS) [quest]; Reedknot Ring (9622, -6.52 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.08 DPS) | yes | Philanthropist's Ring (281635, -1.25 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.72 DPS) [quest]; Reedknot Ring (9622, -3.27 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (599.4 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (599.4 DPS) | yes | Mark of the Chosen (17774, -6.76 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (599.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.18 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -46.31 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 147.0 spell_power points (80.03 DPS) | yes | Brilliant Wand (249385, -15.02 DPS) [crafted]; Stormrager (16997, -16.96 DPS) [quest]; Greater Eternal Wand (249237, -21.80 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 54.1. Weights run: 1.7s. Verify run: 1.2s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.068 ± 0.003, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.393 per %), hit=0.105 ± 0.004 per rating point (10 rating = 1%, 1.049 per %), spell_haste=not significant (0.043 ± 0.037), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.780 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.68 DPS) | yes | Red Winter Hat (21524, -3.66 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.40 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.28 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.42 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (1.12 DPS) | yes | Black Whelp Cloak (7283, -0.34 DPS, sim-verified) [crafted]; Mystic Shawl (3449, -0.56 DPS) [quest]; Pearl-clasped Cloak (5542, -0.56 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.40 DPS) | yes | Green Woolen Vest (2582, -0.28 DPS) [crafted]; Bloody Apron (6226, -0.28 DPS) [dungeon]; Gray Woolen Robe (2585, -1.71 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.84 DPS) | yes | Owlbeard Bracers (16981, -0.56 DPS) [quest]; Windsong Bangles (263336, -0.61 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.96 DPS) | yes | Gnoll Casting Gloves (892, -0.42 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.84 DPS) [quest]; Pristine Gloves (253913, -0.84 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.56 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.16 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.28 DPS) [crafted]; Embroidered Leggings (270011, -0.56 DPS) [quest]; Silk-threaded Trousers (1929, -2.07 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.96 DPS) | yes | Red Woolen Boots (4313, -1.03 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.12 DPS) [crafted]; Soft-soled Linen Boots (4312, -1.40 DPS) [crafted] |
| finger1 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.84 DPS) | yes | - |
| finger2 | - | - |  |  |  |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 98 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Seer's Fine Stein (7608) (or Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), Antipodean Rod (2879), Pulsating Hydra Heart (5183), and 7 more) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Tork Wrench (11855, -0.15 DPS, sim-verified) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 87.9 spell_power points (24.64 DPS) | yes | Cookie's Stirring Rod (5198, -0.59 DPS, sim-verified) [dungeon]; Skycaller (12984, -3.07 DPS) [world_drop]; Greater Magic Wand (11288, -6.58 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Seer's Fine Stein; ranged: Sizzle Stick

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 99.5. Weights run: 1.6s. Verify run: 1.1s. 289 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.016 ± 0.004, crit=0.047 ± 0.001 per rating point (14 rating = 1%, 0.659 per %), hit=0.152 ± 0.005 per rating point (10 rating = 1%, 1.518 per %), spell_haste=-0.292 ± 0.060, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.792 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (3.24 DPS) | yes | Filigreed Pristine Circlet (253975, -0.88 DPS) [crafted]; Enchanter's Cowl (4322, -1.43 DPS) [crafted]; Silk Headband (7050, -4.77 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.1 spell_power points (0.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.01 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.1 spell_power points (2.69 DPS) | yes | Invoker's Mantle (215365, -0.61 DPS) [crafted]; Death Speaker Mantle (6685, -0.87 DPS) [dungeon]; Chestnut Mantle (17695, -1.18 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.47 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.29 DPS) [crafted]; Black Whelp Cloak (7283, -0.59 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.83 DPS) | yes | Green Silk Armor (7065, -0.76 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.12 DPS) [dungeon]; Mechbuilder's Overalls (9508, -1.42 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.65 DPS) | yes | Tabitha's Cuffs (251486, -2.21 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -2.35 DPS) [quest]; Windsong Bangles (263336, -2.35 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.06 DPS) | yes | Gnoll Casting Gloves (892, -0.29 DPS) [world]; Truefaith Gloves (7049, -0.57 DPS) [crafted]; Jutebraid Gloves (10654, -0.58 DPS, sim-verified) [quest] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (3.24 DPS) | yes | Belt of Arugal (6392, -0.57 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.88 DPS) [dungeon]; Invoker's Cord (215366, -1.15 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.53 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.47 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.74 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.1 spell_power points (2.09 DPS) | yes | Glinteye Slippers (273024, -0.03 DPS) [dungeon]; Acidic Walkers (9454, -0.58 DPS) [dungeon]; Spidersilk Boots (4320, -2.57 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.1 spell_power points (1.79 DPS) | yes | Solid Gold Band (279845, -0.61 DPS) [quest]; Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Runemender's Seal (273031, -0.91 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.77 DPS) | yes | Solid Gold Band (279845, -0.59 DPS) [quest]; Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon]; Runemender's Seal (273031, -0.88 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (10.59 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -10.54 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -10.55 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 111.7 spell_power points (32.88 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.40 DPS) [crafted]; Gyromatic Icemaker (9489, -2.52 DPS) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 289, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 130.9. Weights run: 1.8s. Verify run: 1.4s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.171 ± 0.005, crit=0.083 ± 0.001 per rating point (14 rating = 1%, 1.166 per %), hit=0.201 ± 0.007 per rating point (10 rating = 1%, 2.006 per %), spell_haste=0.392 ± 0.079, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.815 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.57 DPS) | yes | Augural Shroud (2620, -2.60 DPS) [world]; Living Cowl (5608, -2.64 DPS, sim-verified) [world]; Holy Shroud (2721, -3.13 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.4 spell_power points (2.62 DPS) | yes | Triune Amulet (7722, -0.89 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.24 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.30 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (3.30 DPS) | yes | Inquisitor's Shawl (19507, -0.41 DPS) [dungeon]; Green Silken Shoulders (7057, -0.45 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.57 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.19 DPS) | yes | Long Silken Cloak (4326, -0.05 DPS) [crafted]; Guardian Cloak (5965, -0.05 DPS) [crafted]; Hillman's Cloak (3719, -0.63 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (7.21 DPS) | yes | Elemental Raiment (9434, -0.74 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.09 DPS) [crafted]; Robe of Power (7054, -2.18 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.82 DPS) | yes | Condor Bracers (15864, -0.67 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.14 DPS) [quest]; Earthen Silk Cuffs (254019, -1.56 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (5.85 DPS) | yes | Black Mageweave Gloves (10003, -1.36 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.87 DPS) [crafted]; Gilded Handwraps (254021, -2.97 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (4.07 DPS) | yes | Warsong Sash (16975, -0.69 DPS, sim-verified) [quest]; Deathmage Sash (10771, -1.08 DPS) [dungeon]; Belt of Arugal (6392, -1.09 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.0 spell_power points (5.02 DPS) | yes | Gaze Dreamer Pants (6903, -1.27 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.82 DPS) [crafted]; Stoneweaver Leggings (9407, -2.40 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.51 DPS) | yes | Gilded Slippers (254001, -3.30 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.11 DPS) [crafted]; Glinteye Slippers (273024, -5.32 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.9 spell_power points (3.40 DPS) | yes | Snake Hoop (6750, -1.25 DPS) [quest]; Ring of Power Regulation (273030, -1.52 DPS) [dungeon]; Solid Gold Band (279845, -2.14 DPS) [quest] |
| finger2 | Reedknot Ring (9622) | Jarl Needs a Blade [quest] | 7.0 spell_power points (2.19 DPS) | yes | Snake Hoop (6750, -0.05 DPS) [quest]; Ring of Power Regulation (273030, -0.31 DPS) [dungeon]; Solid Gold Band (279845, -0.94 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.11 DPS) [dungeon]; Spine of the Devourer (273040, -2.93 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.38 DPS) | yes | Omega Orb (7749, -1.54 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -1.87 DPS) [crafted]; Beacon of Hope (9393, -1.88 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | sim-verified (130.9 DPS) | yes | Twisted Nether Wand (249144, -0.65 DPS) [crafted]; Umbral Wand (5216, -0.86 DPS) [dungeon]; Jaina's Firestarter (13064, -1.39 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Reedknot Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Earthen Rod

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 210.6. Weights run: 1.8s. Verify run: 1.5s. 498 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.272 ± 0.008, crit=0.126 ± 0.002 per rating point (14 rating = 1%, 1.768 per %), hit=0.260 ± 0.009 per rating point (10 rating = 1%, 2.599 per %), spell_haste=not significant (0.230 ± 0.129), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.828 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.08 DPS) | yes | Red Mageweave Headband (10033, -1.08 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.10 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.02 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.2 spell_power points (3.09 DPS) | yes | Mindburst Medallion (11196, -0.52 DPS) [quest]; Triune Amulet (7722, -0.62 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.81 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.9 spell_power points (6.02 DPS) | yes | Black Mageweave Shoulders (10027, -0.71 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.91 DPS) [vendor]; Bloodmage Mantle (7684, -2.17 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (5.26 DPS) | yes | Deep Woodlands Cloak (19121, -1.01 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.50 DPS) [crafted]; Nightfall Drape (12465, -2.23 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.4 spell_power points (8.22 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.16 DPS) [world_drop]; Dreamweave Vest (10021, -1.34 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (210.6 DPS) | yes | Bloodband Bracers (11469, -0.49 DPS) [quest]; Condor Bracers (15864, -0.64 DPS) [quest]; Spidertank Oilrag (9448, -2.65 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (6.42 DPS) | yes | Black Mageweave Gloves (10003, -1.38 DPS) [crafted]; Runecloth Gloves (13863, -1.56 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.11 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.7 spell_power points (5.63 DPS) | yes | Ghostweave Cord (254073, +0.00 DPS, sim-verified) [crafted]; Star Belt (4329, -1.25 DPS) [crafted]; Gilded Waistcord (254081, -1.44 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.39 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.66 DPS) [vendor]; Red Mageweave Pants (10009, -0.76 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -1.71 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.07 DPS) | yes | Gilded Sandals (254107, -3.55 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.68 DPS) [vendor]; Black Mageweave Boots (10026, -3.73 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.37 DPS) | yes | Snake Hoop (6750, -1.90 DPS) [quest]; Reedknot Ring (9622, -2.02 DPS) [quest]; Ring of Power Regulation (273030, -2.36 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.4 spell_power points (3.82 DPS) | yes | Snake Hoop (6750, -1.41 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.47 DPS) [quest]; Ring of Power Regulation (273030, -1.80 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.41 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.42 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.23 DPS) [crafted]; Rune of the Guard Captain (19120, -2.93 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.43 DPS) [world]; Spellshifter Rod (9527, -1.01 DPS) [quest]; Hanzo Sword (8190, -28.01 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 156.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -4.06 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 498, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 308.0. Weights run: 1.9s. Verify run: 4.0s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.124 ± 0.010, crit=0.178 ± 0.003 per rating point (14 rating = 1%, 2.499 per %), hit=0.439 ± 0.014 per rating point (10 rating = 1%, 4.391 per %), spell_haste=not significant (0.051 ± 0.166), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.829 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.47 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.44 DPS) [pvp]; Dreamweave Circlet (10041, -1.67 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.2 spell_power points (4.65 DPS) | yes | Glowing Eye of Mordresh (10769, -1.84 DPS) [dungeon]; Mindburst Medallion (11196, -2.28 DPS) [quest]; Triune Amulet (7722, -3.59 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (8.77 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.36 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -3.12 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.7 spell_power points (5.17 DPS) | yes | Deep Woodlands Cloak (19121, -0.57 DPS) [quest]; Runecloth Cloak (13860, -1.67 DPS) [crafted]; Nightfall Drape (12465, -2.02 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (16.53 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.26 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -6.89 DPS) [pvp]; Robe of the Magi (1716, -9.30 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (308.0 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -5.65 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (308.0 DPS) | yes | Blood Guard's Dreadweave Handwraps (227099, +0.00 DPS) [pvp]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -4.89 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (308.0 DPS) | yes | Star Belt (4329, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -7.07 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.12 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -3.21 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.42 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.70 DPS) [crafted]; Omnicast Boots (11822, -0.88 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.9 spell_power points (6.62 DPS) | yes | Philanthropist's Ring (281635, -2.89 DPS) [quest]; Reedknot Ring (9622, -4.17 DPS) [quest]; Snake Hoop (6750, -4.30 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.56 DPS) | yes | Philanthropist's Ring (281635, -0.74 DPS, sim-verified) [quest]; Reedknot Ring (9622, -2.11 DPS) [quest]; Snake Hoop (6750, -2.24 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (308.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (308.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (308.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.18 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -33.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 224.8 spell_power points (78.87 DPS) | yes | Brilliant Wand (249385, -0.92 DPS, sim-verified) [crafted]; Stormrager (16997, -15.96 DPS) [quest]; Torch of Austen (13004, -22.08 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25520100135001050-0005200000000000000-0550004000000000)

Set DPS (verified): 590.4. Weights run: 2.0s. Verify run: 2.8s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.140 ± 0.019, crit=0.400 ± 0.010 per rating point (14 rating = 1%, 5.597 per %), hit=0.678 ± 0.033 per rating point (10 rating = 1%, 6.780 per %), spell_haste=not significant (0.492 ± 0.364), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.870 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (14.70 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.50 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.4 spell_power points (7.29 DPS) | yes | Glowing Eye of Mordresh (10769, -2.87 DPS) [dungeon]; Triune Amulet (7722, -2.94 DPS) [dungeon]; Mindburst Medallion (11196, -3.57 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (13.61 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.32 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -3.04 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.8 spell_power points (8.08 DPS) | yes | Runecloth Cloak (13860, -2.57 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.05 DPS, sim-verified) [quest]; Nightfall Drape (12465, -3.18 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.3 spell_power points (25.72 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.48 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.59 DPS) [pvp]; Robe of the Magi (1716, -14.62 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.4 spell_power points (6.75 DPS) | yes | Earthenweave Cuffs (254131, -2.39 DPS) [crafted]; Nethergeld Cuffs (254061, -2.41 DPS) [crafted]; Spidertank Oilrag (9448, -3.26 DPS, sim-verified) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (11.97 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS, sim-verified) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.24 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.4 spell_power points (8.38 DPS) | yes | Star Belt (4329, -1.31 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.33 DPS) [crafted]; Ghostweave Cord (254073, -2.31 DPS, sim-verified) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (14.15 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Stone Guard's Dreadweave Leggings (220906, -3.54 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.06 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.09 DPS) [crafted]; Omnicast Boots (11822, -1.26 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.0 spell_power points (10.33 DPS) | yes | Philanthropist's Ring (281635, -4.51 DPS) [quest]; Reedknot Ring (9622, -6.52 DPS) [quest]; Snake Hoop (6750, -6.68 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.08 DPS) | yes | Philanthropist's Ring (281635, -1.25 DPS) [quest]; Reedknot Ring (9622, -3.27 DPS) [quest]; Snake Hoop (6750, -3.43 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (590.4 DPS) | yes | Uther's Strength (11302, -0.06 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (590.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (590.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.18 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -47.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 147.0 spell_power points (80.03 DPS) | yes | Brilliant Wand (249385, -2.80 DPS, sim-verified) [crafted]; Stormrager (16997, -16.96 DPS) [quest]; Greater Eternal Wand (249237, -21.80 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

