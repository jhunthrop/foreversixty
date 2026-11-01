# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 56.4. Weights run: 1.3s. Verify run: 0.8s. 186 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.088 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.416 per %), hit=0.090 ± 0.004 per rating point (10 rating = 1%, 0.902 per %), spell_haste=not significant (0.038 ± 0.047), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.795 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.8 spell_power points (2.28 DPS) | yes | Pristine Circlet (253949, -0.92 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.30 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (1.04 DPS) | yes | Black Whelp Cloak (7283, -0.50 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.52 DPS) [crafted]; Reinforced Linen Cape (2580, -0.78 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 11.1 spell_power points (2.90 DPS) | yes | Filigreed Shadow Gown (253907, -0.84 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.60 DPS) [crafted]; Gray Woolen Robe (2585, -1.86 DPS) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.26 DPS) | yes | Ivycloth Bracelets (9793, -0.29 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.83 DPS) | yes | Shadow Gloves (253919, -0.37 DPS) [crafted]; Gnoll Casting Gloves (892, -0.42 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.78 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 5.6 spell_power points (1.45 DPS) | yes | Pristine Sash (253925, -0.41 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.48 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.93 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 8.8 spell_power points (2.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.72 DPS) [crafted]; Silk-threaded Trousers (1929, -0.92 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.98 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.83 DPS) | yes | Shadow Boots (253895, -0.54 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.78 DPS) [crafted]; Pristine Boots (253889, -1.04 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.30 DPS) | yes | - |
| finger2 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | 1.0 spell_power points (0.26 DPS) | yes | Ring of the Shadow (1462, -0.29 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Witching Stave (1484) | Shadowfang Keep [dungeon] | 21.5 spell_power points (5.60 DPS) | yes | Shadowfang (1482, -6.64 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 85.5 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.75 DPS) [world_drop]; Deepblaze (279896, -3.78 DPS) [quest]; Sizzle Stick (8071, -4.14 DPS) [quest] |

**New at 20:** head: Shadow Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Furlfeather Ring; main_hand: Witching Stave; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 88.6. Weights run: 1.4s. Verify run: 1.4s. 323 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.021 ± 0.004, crit=0.052 ± 0.001 per rating point (14 rating = 1%, 0.722 per %), hit=0.122 ± 0.005 per rating point (10 rating = 1%, 1.217 per %), spell_haste=not significant (-0.208 ± 0.067), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.807 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 11.3 spell_power points (3.08 DPS) | yes | Holy Shroud (2721, -0.08 DPS) [world_drop]; Silk Headband (7050, -0.63 DPS) [crafted]; Shadow Circlet (253955, -0.66 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.1 spell_power points (0.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.01 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.2 spell_power points (2.50 DPS) | yes | Moonlit Amice (11884, -0.60 DPS) [quest]; Death Speaker Mantle (6685, -0.81 DPS) [dungeon]; Invoker's Mantle (215365, -1.33 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.36 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Black Whelp Cloak (7283, -0.54 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.54 DPS) | yes | Necrology Robes (2292, -0.44 DPS, sim-verified) [dungeon]; Shadow Gown (253967, -0.64 DPS) [crafted]; Green Silk Armor (7065, -1.02 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.45 DPS) | yes | Glowing Magical Bracelets (13106, -2.41 DPS) [world_drop]; Nightsky Wristbands (6407, -2.42 DPS) [world_drop]; Windsong Bangles (263336, -2.87 DPS, sim-verified) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.4 spell_power points (2.02 DPS) | yes | Hands of Darkness (7047, -0.04 DPS) [crafted]; Serpent Gloves (5970, -0.11 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.1 spell_power points (2.47 DPS) | yes | Invoker's Cord (215366, -0.53 DPS) [crafted]; Crimson Silk Belt (7055, -0.79 DPS) [crafted]; Ghamoo-ra's Bind (6908, -1.03 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.27 DPS) | yes | Shadow Leggings (253993, -0.59 DPS, sim-verified) [crafted]; Filigreed Shadow Leggings (253943, -0.82 DPS) [crafted]; Necromancer Leggings (2277, -1.01 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.90 DPS) | yes | Spidersilk Boots (4320, -0.97 DPS) [crafted]; Glinteye Slippers (273024, -0.99 DPS) [dungeon]; Gilded Slippers (254001, -1.18 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.1 spell_power points (1.66 DPS) | yes | Minor Channeling Ring (1449, -0.29 DPS) [quest]; Electrocutioner Lagnut (9447, -0.85 DPS) [dungeon]; Runemender's Seal (273031, -0.85 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.63 DPS) | yes | Minor Channeling Ring (1449, -0.47 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.82 DPS) [dungeon]; Runemender's Seal (273031, -0.82 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | sim-verified (88.6 DPS) | yes | Staff of the Shade (2549, -0.94 DPS, sim-verified) [dungeon]; Witching Stave (1484, -1.51 DPS) [dungeon]; Glimmering Staff (249392, -7.41 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 120.7 spell_power points (32.88 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.52 DPS) [crafted]; Gyromatic Icemaker (9489, -2.52 DPS) [dungeon] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Staff of Soran'ruk; ranged: Starfaller

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 136.0. Weights run: 1.5s. Verify run: 1.4s. 420 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.161 ± 0.005, crit=0.086 ± 0.001 per rating point (14 rating = 1%, 1.201 per %), hit=0.153 ± 0.006 per rating point (10 rating = 1%, 1.528 per %), spell_haste=not significant (0.220 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.830 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.24 DPS) | yes | Augural Shroud (2620, -2.49 DPS) [world]; Living Cowl (5608, -2.57 DPS, sim-verified) [world]; Filigreed Shadow Circlet (253981, -2.79 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.3 spell_power points (2.46 DPS) | yes | Triune Amulet (7722, -0.09 DPS) [dungeon]; Necklace of Calisea (1714, -2.13 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.18 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (3.10 DPS) | yes | Inquisitor's Shawl (19507, -0.40 DPS) [dungeon]; Green Silken Shoulders (7057, -0.41 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.55 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.08 DPS) | yes | Guardian Cloak (5965, -0.06 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.59 DPS) [quest]; Long Silken Cloak (4326, -1.62 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (6.82 DPS) | yes | Dreamweave Vest (10021, -1.04 DPS) [crafted]; Elemental Raiment (9434, -1.06 DPS, sim-verified) [world_drop]; Shadoweave Robe (10004, -1.65 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.67 DPS) | yes | Spidertank Oilrag (9448, -0.37 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.59 DPS) [quest]; Earthen Silk Cuffs (254019, -1.49 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (5.54 DPS) | yes | Shadoweave Gloves (10023, -1.11 DPS) [crafted]; Black Handwraps (254027, -1.75 DPS) [crafted]; Black Mageweave Gloves (10003, -2.13 DPS, sim-verified) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.86 DPS) | yes | Black Cord (254043, -0.51 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -1.04 DPS) [dungeon]; Deathmage Sash (10771, -1.06 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.4 spell_power points (5.18 DPS) | yes | Red Mageweave Pants (10009, -1.34 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -1.39 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.61 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.13 DPS) | yes | Gilded Slippers (254001, -4.72 DPS) [crafted]; Spidersilk Boots (4320, -4.86 DPS) [crafted]; Black Slippers (254007, -5.33 DPS, sim-verified) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.4 spell_power points (3.69 DPS) | yes | Ring of Forlorn Spirits (2043, -1.31 DPS) [quest]; Reedknot Ring (9622, -1.61 DPS) [quest]; Snake Hoop (6750, -1.67 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (3.21 DPS) | yes | Reedknot Ring (9622, -1.13 DPS) [quest]; Snake Hoop (6750, -1.19 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.59 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (136.0 DPS) | yes | Staff of Soran'ruk (15109, -2.21 DPS) [quest]; Staff of the Shade (2549, -2.95 DPS, sim-verified) [dungeon]; Witching Stave (1484, -3.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 133.5 spell_power points (39.66 DPS) | yes | Earthen Rod (9381, -3.18 DPS) [dungeon]; Twisted Nether Wand (249144, -3.88 DPS) [crafted]; Umbral Wand (5216, -4.00 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Star Belt; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 420, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 197.2. Weights run: 1.5s. Verify run: 1.6s. 531 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.298 ± 0.009, crit=0.132 ± 0.002 per rating point (14 rating = 1%, 1.847 per %), hit=0.208 ± 0.010 per rating point (10 rating = 1%, 2.075 per %), spell_haste=not significant (0.120 ± 0.139), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.844 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.57 DPS) | yes | Dreamweave Circlet (10041, -0.96 DPS) [crafted]; Shadoweave Mask (10025, -1.19 DPS) [crafted]; Red Mageweave Headband (10033, -1.46 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.4 spell_power points (2.98 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.51 DPS) [quest]; Horizon Choker (13085, -1.65 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.4 spell_power points (5.83 DPS) | yes | Netherpearl Shoulders (254055, +0.00 DPS, sim-verified) [crafted]; Shadoweave Shoulders (10028, -1.28 DPS) [crafted]; Black Mageweave Shoulders (10027, -1.80 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.8 spell_power points (5.01 DPS) | yes | Runecloth Cloak (13860, -1.82 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.16 DPS) [dungeon]; Big Voodoo Cloak (8216, -2.57 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.0 spell_power points (7.93 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.26 DPS) [world_drop]; Dreamweave Vest (10021, -1.36 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.1 spell_power points (4.15 DPS) | yes | Arcane Runed Bracers (4744, -1.29 DPS) [quest]; Spidertank Oilrag (9448, -1.29 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.44 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (6.10 DPS) | yes | Black Gloves (254101, -0.05 DPS) [crafted]; Shadoweave Gloves (10023, -1.06 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.11 DPS) [vendor] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.7 spell_power points (5.95 DPS) | yes | Satyrmane Sash (17755, -0.53 DPS, sim-verified) [dungeon]; Black Cord (254043, -1.44 DPS) [crafted]; Ghostweave Cord (254073, -1.50 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.5 spell_power points (8.11 DPS) | yes | Wizardweave Leggings (14132, -1.11 DPS, sim-verified) [crafted]; Shadoweave Pants (10002, -2.48 DPS) [crafted]; Red Mageweave Pants (10009, -2.52 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.62 DPS) | yes | Black Sandals (254113, -1.79 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -3.28 DPS) [crafted]; Black Mageweave Boots (10026, -3.47 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.3 spell_power points (4.23 DPS) | yes | Philanthropist's Ring (281635, -0.58 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.69 DPS) [quest]; Snake Hoop (6750, -1.85 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.13 DPS) | yes | Philanthropist's Ring (281635, -0.48 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.59 DPS) [quest]; Snake Hoop (6750, -1.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (197.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (197.2 DPS) | yes | Mark of the Chosen (17774, -2.07 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (197.2 DPS) | yes | Spire of Hakkar (10844, -1.09 DPS) [world]; Inventor's Focal Sword (17719, -1.14 DPS) [dungeon]; Spellshifter Rod (9527, -1.45 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 165.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.21 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 531, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 289.6. Weights run: 1.6s. Verify run: 2.4s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.121 ± 0.011, crit=0.189 ± 0.003 per rating point (14 rating = 1%, 2.649 per %), hit=0.357 ± 0.014 per rating point (10 rating = 1%, 3.572 per %), spell_haste=not significant (0.353 ± 0.188), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.843 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.83 DPS) | yes | Felcloth Hood (14111, +0.00 DPS, sim-verified) [crafted]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.38 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.2 spell_power points (4.32 DPS) | yes | Glowing Eye of Mordresh (10769, -1.71 DPS) [dungeon]; Triune Amulet (7722, -1.75 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -2.12 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (8.17 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Felcloth Shoulders (14112, -1.01 DPS) [crafted]; Mantle of the Timbermaw (19050, -1.24 DPS) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | 15.3 spell_power points (5.00 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.74 DPS) [crafted]; Nightfall Drape (12465, -2.06 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (15.40 DPS) | yes | Robe of Winter Night (14136, -3.28 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -3.98 DPS) [pvp]; Felcloth Robe (14106, -5.48 DPS) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.2 spell_power points (3.99 DPS) | yes | Netherpearl Cuffs (254067, +0.00 DPS, sim-verified) [crafted]; Arcane Runed Bracers (4744, -1.05 DPS) [quest]; Spidertank Oilrag (9448, -1.05 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (7.19 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.17 DPS) [pvp]; Dreamweave Gloves (10019, -1.15 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 17.1 spell_power points (5.59 DPS) | yes | Satyrmane Sash (17755, -0.62 DPS) [dungeon]; Ghostweave Cord (254073, -1.01 DPS) [crafted]; Star Belt (4329, -1.34 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (8.50 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Felcloth Pants (14107, -1.62 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.85 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.65 DPS) [crafted]; Omnicast Boots (11822, -0.83 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.8 spell_power points (6.16 DPS) | yes | Underworld Band (1980, -2.11 DPS) [world_drop]; Philanthropist's Ring (281635, -2.69 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.55 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.25 DPS) | yes | Underworld Band (1980, +0.00 DPS, sim-verified) [world_drop]; Philanthropist's Ring (281635, -0.78 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.63 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (289.6 DPS) | yes | Uther's Strength (11302, -0.89 DPS, sim-verified) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (289.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (289.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Inventor's Focal Sword (17719, -0.47 DPS) [dungeon]; Spire of Hakkar (10844, -0.69 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 240.8 spell_power points (78.73 DPS) | yes | Stormrager (16997, -2.30 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.92 DPS) [crafted]; Torch of Austen (13004, -21.94 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; back: Fel Cape; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; trinket2: Mark of the Chosen; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25520100135001050-0005200000000000000-0550004000000000)

Set DPS (verified): 569.8. Weights run: 1.7s. Verify run: 2.6s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.174 ± 0.019, crit=0.417 ± 0.010 per rating point (14 rating = 1%, 5.844 per %), hit=0.575 ± 0.034 per rating point (10 rating = 1%, 5.755 per %), spell_haste=not significant (0.525 ± 0.377), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.863 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 27.6 spell_power points (14.22 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -0.32 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (7.07 DPS) | yes | Glowing Eye of Mordresh (10769, -2.75 DPS) [dungeon]; Triune Amulet (7722, -2.95 DPS) [dungeon]; Mindburst Medallion (11196, -3.45 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.1 spell_power points (12.92 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.05 DPS) [crafted]; Felcloth Shoulders (14112, -1.38 DPS) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | 18.8 spell_power points (9.67 DPS) | yes | Spritecaster Cape (11623, -1.93 DPS) [dungeon]; Runecloth Cloak (13860, -4.32 DPS) [crafted]; Nightfall Drape (12465, -5.04 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (24.48 DPS) | yes | Robe of Winter Night (14136, -3.47 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -5.86 DPS) [pvp]; Felcloth Robe (14106, -8.50 DPS) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (6.56 DPS) | yes | Netherpearl Cuffs (254067, -0.16 DPS) [crafted]; Arcane Runed Bracers (4744, -1.92 DPS) [quest]; Spidertank Oilrag (9448, -1.92 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.4 spell_power points (11.54 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -0.22 DPS) [quest]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.38 DPS) [pvp] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.0 spell_power points (9.24 DPS) | yes | Satyrmane Sash (17755, -1.14 DPS) [dungeon]; Ghostweave Cord (254073, -2.04 DPS) [crafted]; Black Cord (254043, -2.31 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (13.38 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Felcloth Pants (14107, -0.76 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.35 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.98 DPS) [dungeon]; Venomspew Footpads (275606, -1.03 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (9.89 DPS) | yes | Underworld Band (1980, -3.23 DPS) [world_drop]; Philanthropist's Ring (281635, -4.30 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.77 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.69 DPS) | yes | Underworld Band (1980, +0.00 DPS, sim-verified) [world_drop]; Philanthropist's Ring (281635, -1.10 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.57 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (569.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (569.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (569.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Inventor's Focal Sword (17719, -1.07 DPS) [dungeon]; Spire of Hakkar (10844, -3.05 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.1 spell_power points (79.85 DPS) | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -16.71 DPS) [quest]; Greater Eternal Wand (249237, -21.89 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; legs: Ghostweave Pants; finger1: Maiden's Circle; trinket2: Mark of the Chosen; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 56.5. Weights run: 1.3s. Verify run: 0.9s. 180 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.088 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.416 per %), hit=0.090 ± 0.004 per rating point (10 rating = 1%, 0.902 per %), spell_haste=not significant (0.038 ± 0.047), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.795 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.8 spell_power points (2.28 DPS) | yes | Pristine Circlet (253949, -0.93 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.30 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (1.04 DPS) | yes | Mystic Shawl (3449, -0.52 DPS) [quest]; Pearl-clasped Cloak (5542, -0.52 DPS) [crafted]; Black Whelp Cloak (7283, -0.63 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 11.1 spell_power points (2.90 DPS) | yes | Filigreed Shadow Gown (253907, -1.08 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.60 DPS) [crafted]; Gray Woolen Robe (2585, -1.86 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.78 DPS) | yes | Owlbeard Bracers (16981, -0.52 DPS) [quest]; Windsong Bangles (263336, -0.59 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.83 DPS) | yes | Shadow Gloves (253919, -0.37 DPS) [crafted]; Gnoll Casting Gloves (892, -0.42 DPS, sim-verified) [world]; Blight Gloves (279877, -0.52 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 5.6 spell_power points (1.45 DPS) | yes | Pristine Sash (253925, -0.41 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.66 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.93 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 8.8 spell_power points (2.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.72 DPS) [crafted]; Silk-threaded Trousers (1929, -0.89 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.98 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.83 DPS) | yes | Shadow Boots (253895, -0.71 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.78 DPS) [crafted]; Pristine Boots (253889, -1.04 DPS) [crafted] |
| finger1 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | 1.0 spell_power points (0.26 DPS) | yes | Ring of the Shadow (1462, -0.29 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.78 DPS) | yes | Ring of the Shadow (1462, -1.13 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Witching Stave (1484) | Shadowfang Keep [dungeon] | 21.5 spell_power points (5.60 DPS) | yes | Shadowfang (1482, -6.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 85.5 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.75 DPS) [world_drop]; Sizzle Stick (8071, -4.14 DPS) [quest]; Greater Magic Wand (11288, -4.29 DPS) [crafted] |

**New at 20:** head: Shadow Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Furlfeather Ring; finger2: Sludge-Stained Band; main_hand: Witching Stave; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 180, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 88.2. Weights run: 1.4s. Verify run: 1.4s. 314 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.021 ± 0.004, crit=0.052 ± 0.001 per rating point (14 rating = 1%, 0.722 per %), hit=0.122 ± 0.005 per rating point (10 rating = 1%, 1.217 per %), spell_haste=not significant (-0.208 ± 0.067), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.807 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 11.3 spell_power points (3.08 DPS) | yes | Holy Shroud (2721, -0.08 DPS) [world_drop]; Silk Headband (7050, -0.63 DPS) [crafted]; Shadow Circlet (253955, -0.66 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.1 spell_power points (0.03 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.01 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.2 spell_power points (2.50 DPS) | yes | Invoker's Mantle (215365, -0.42 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.81 DPS) [dungeon]; Magician's Mantle (12998, -1.09 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.36 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Black Whelp Cloak (7283, -0.54 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.54 DPS) | yes | Necrology Robes (2292, -0.43 DPS) [dungeon]; Shadow Gown (253967, -0.64 DPS) [crafted]; Green Silk Armor (7065, -1.02 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.45 DPS) | yes | Tabitha's Cuffs (251486, -1.90 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -2.17 DPS) [quest]; Windsong Bangles (263336, -2.18 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.4 spell_power points (2.02 DPS) | yes | Hands of Darkness (7047, -0.04 DPS) [crafted]; Serpent Gloves (5970, -0.11 DPS) [dungeon]; Jutebraid Gloves (10654, -0.36 DPS) [quest] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (3.00 DPS) | yes | Belt of Arugal (6392, -0.53 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.82 DPS) [dungeon]; Invoker's Cord (215366, -1.06 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.27 DPS) | yes | Shadow Leggings (253993, -0.15 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.82 DPS) [crafted]; Necromancer Leggings (2277, -1.01 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.90 DPS) | yes | Spidersilk Boots (4320, -0.97 DPS) [crafted]; Glinteye Slippers (273024, -0.99 DPS) [dungeon]; Gilded Slippers (254001, -1.17 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.1 spell_power points (1.66 DPS) | yes | Solid Gold Band (279845, -0.57 DPS) [quest]; Electrocutioner Lagnut (9447, -0.85 DPS) [dungeon]; Runemender's Seal (273031, -0.85 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.63 DPS) | yes | Solid Gold Band (279845, -0.47 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.82 DPS) [dungeon]; Runemender's Seal (273031, -0.82 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | sim-verified (88.2 DPS) | yes | Staff of the Shade (2549, -0.92 DPS, sim-verified) [dungeon]; Witching Stave (1484, -1.51 DPS) [dungeon]; Glimmering Staff (249392, -7.41 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 120.7 spell_power points (32.88 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.52 DPS) [crafted]; Gyromatic Icemaker (9489, -2.52 DPS) [dungeon] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Staff of Soran'ruk; ranged: Starfaller

No-known-source sample (15 of 314, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 134.8. Weights run: 1.5s. Verify run: 1.4s. 411 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.161 ± 0.005, crit=0.086 ± 0.001 per rating point (14 rating = 1%, 1.201 per %), hit=0.153 ± 0.006 per rating point (10 rating = 1%, 1.528 per %), spell_haste=not significant (0.220 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.830 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.24 DPS) | yes | Augural Shroud (2620, -2.49 DPS) [world]; Living Cowl (5608, -2.55 DPS, sim-verified) [world]; Filigreed Shadow Circlet (253981, -2.79 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.3 spell_power points (2.46 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.13 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.18 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (3.10 DPS) | yes | Green Silken Shoulders (7057, -0.20 DPS) [crafted]; Inquisitor's Shawl (19507, -0.40 DPS) [dungeon]; Berylline Pads (4197, -0.55 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.08 DPS) | yes | Long Silken Cloak (4326, -0.06 DPS) [crafted]; Guardian Cloak (5965, -0.06 DPS) [crafted]; Hillman's Cloak (3719, -0.59 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (6.82 DPS) | yes | Elemental Raiment (9434, -0.75 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.04 DPS) [crafted]; Shadoweave Robe (10004, -1.65 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.67 DPS) | yes | Condor Bracers (15864, -0.79 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.10 DPS) [quest]; Earthen Silk Cuffs (254019, -1.49 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (5.54 DPS) | yes | Shadoweave Gloves (10023, -1.11 DPS) [crafted]; Black Mageweave Gloves (10003, -1.38 DPS, sim-verified) [crafted]; Black Handwraps (254027, -1.75 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.86 DPS) | yes | Black Cord (254043, -0.03 DPS) [crafted]; Warsong Sash (16975, -0.59 DPS) [quest]; Belt of Arugal (6392, -1.04 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.4 spell_power points (5.18 DPS) | yes | Red Mageweave Pants (10009, -0.76 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -1.39 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.61 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.13 DPS) | yes | Black Slippers (254007, -3.79 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -4.72 DPS) [crafted]; Spidersilk Boots (4320, -4.86 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.4 spell_power points (3.69 DPS) | yes | Reedknot Ring (9622, -1.61 DPS) [quest]; Snake Hoop (6750, -1.67 DPS) [quest]; Ring of Power Regulation (273030, -1.91 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (3.21 DPS) | yes | Snake Hoop (6750, -1.19 DPS) [quest]; Reedknot Ring (9622, -1.27 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -1.43 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (134.8 DPS) | yes | Staff of the Shade (2549, -1.81 DPS, sim-verified) [dungeon]; Staff of Soran'ruk (15109, -2.21 DPS) [quest]; Witching Stave (1484, -3.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 133.5 spell_power points (39.66 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -3.88 DPS) [crafted]; Umbral Wand (5216, -4.00 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 196.0. Weights run: 1.5s. Verify run: 1.6s. 522 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.298 ± 0.009, crit=0.132 ± 0.002 per rating point (14 rating = 1%, 1.847 per %), hit=0.208 ± 0.010 per rating point (10 rating = 1%, 2.075 per %), spell_haste=not significant (0.120 ± 0.139), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.844 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.57 DPS) | yes | Dreamweave Circlet (10041, -0.96 DPS) [crafted]; Shadoweave Mask (10025, -1.19 DPS) [crafted]; Red Mageweave Headband (10033, -1.76 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.4 spell_power points (2.98 DPS) | yes | Triune Amulet (7722, -0.44 DPS) [dungeon]; Mindburst Medallion (11196, -0.51 DPS) [quest]; Horizon Choker (13085, -1.65 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.4 spell_power points (5.83 DPS) | yes | Netherpearl Shoulders (254055, -0.52 DPS) [crafted]; Shadoweave Shoulders (10028, -1.28 DPS) [crafted]; Black Mageweave Shoulders (10027, -1.80 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.8 spell_power points (5.01 DPS) | yes | Deep Woodlands Cloak (19121, -0.69 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.40 DPS) [crafted]; Nightfall Drape (12465, -2.16 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.0 spell_power points (7.93 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.26 DPS) [world_drop]; Dreamweave Vest (10021, -1.36 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.1 spell_power points (4.15 DPS) | yes | Spidertank Oilrag (9448, -1.29 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.46 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -1.71 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (6.10 DPS) | yes | Black Gloves (254101, -0.05 DPS) [crafted]; Shadoweave Gloves (10023, -1.06 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.11 DPS) [vendor] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.7 spell_power points (5.95 DPS) | yes | Satyrmane Sash (17755, -0.75 DPS, sim-verified) [dungeon]; Black Cord (254043, -1.44 DPS) [crafted]; Ghostweave Cord (254073, -1.50 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.5 spell_power points (8.11 DPS) | yes | Wizardweave Leggings (14132, -1.89 DPS, sim-verified) [crafted]; Shadoweave Pants (10002, -2.48 DPS) [crafted]; Red Mageweave Pants (10009, -2.52 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.62 DPS) | yes | Black Sandals (254113, -1.97 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -3.28 DPS) [crafted]; Black Mageweave Boots (10026, -3.47 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.3 spell_power points (4.23 DPS) | yes | Philanthropist's Ring (281635, -0.58 DPS) [quest]; Snake Hoop (6750, -1.85 DPS) [quest]; Reedknot Ring (9622, -2.00 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.13 DPS) | yes | Philanthropist's Ring (281635, -1.02 DPS, sim-verified) [quest]; Snake Hoop (6750, -1.75 DPS) [quest]; Reedknot Ring (9622, -1.91 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (196.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (196.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (196.0 DPS) | yes | Spire of Hakkar (10844, -1.09 DPS) [world]; Inventor's Focal Sword (17719, -1.14 DPS) [dungeon]; Spellshifter Rod (9527, -1.45 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 165.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.21 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 522, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 290.6. Weights run: 1.6s. Verify run: 2.7s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.121 ± 0.011, crit=0.189 ± 0.003 per rating point (14 rating = 1%, 2.649 per %), hit=0.357 ± 0.014 per rating point (10 rating = 1%, 3.572 per %), spell_haste=not significant (0.353 ± 0.188), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.843 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | sim-verified (290.6 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.22 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.02 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.2 spell_power points (4.32 DPS) | yes | Glowing Eye of Mordresh (10769, -1.71 DPS) [dungeon]; Mindburst Medallion (11196, -2.12 DPS) [quest]; Triune Amulet (7722, -4.34 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (8.17 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.24 DPS) [crafted]; Felcloth Shoulders (14112, -1.76 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | 15.3 spell_power points (5.00 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Deep Woodlands Cloak (19121, -0.72 DPS) [quest]; Runecloth Cloak (13860, -1.74 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (15.40 DPS) | yes | Robe of Winter Night (14136, -3.24 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -3.98 DPS) [pvp]; Felcloth Robe (14106, -5.48 DPS) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.2 spell_power points (3.99 DPS) | yes | Spidertank Oilrag (9448, -1.05 DPS) [dungeon]; Earthenweave Cuffs (254131, -1.38 DPS) [crafted]; Netherpearl Cuffs (254067, -2.02 DPS, sim-verified) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (7.19 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.17 DPS) [pvp] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 17.1 spell_power points (5.59 DPS) | yes | Satyrmane Sash (17755, -0.62 DPS) [dungeon]; Ghostweave Cord (254073, -1.01 DPS) [crafted]; Star Belt (4329, -1.34 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (8.50 DPS) | yes | Felcloth Pants (14107, +0.00 DPS, sim-verified) [crafted]; Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.85 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.65 DPS) [crafted]; Omnicast Boots (11822, -0.83 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.8 spell_power points (6.16 DPS) | yes | Underworld Band (1980, -2.11 DPS) [world_drop]; Philanthropist's Ring (281635, -2.69 DPS) [quest]; Reedknot Ring (9622, -3.87 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.25 DPS) | yes | Underworld Band (1980, -0.20 DPS) [world_drop]; Philanthropist's Ring (281635, -0.78 DPS) [quest]; Reedknot Ring (9622, -1.96 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.01 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.69 DPS) [world]; Inventor's Focal Sword (17719, -2.52 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 240.8 spell_power points (78.73 DPS) | yes | Stormrager (16997, -2.94 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.92 DPS) [crafted]; Torch of Austen (13004, -21.94 DPS) [world_drop] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Fel Cape; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25520100135001050-0005200000000000000-0550004000000000)

Set DPS (verified): 572.7. Weights run: 1.7s. Verify run: 2.7s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.174 ± 0.019, crit=0.417 ± 0.010 per rating point (14 rating = 1%, 5.844 per %), hit=0.575 ± 0.034 per rating point (10 rating = 1%, 5.755 per %), spell_haste=not significant (0.525 ± 0.377), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.863 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 27.6 spell_power points (14.22 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -0.32 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (7.07 DPS) | yes | Triune Amulet (7722, -2.95 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.97 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.45 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.1 spell_power points (12.92 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.05 DPS) [crafted]; Felcloth Shoulders (14112, -1.38 DPS) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | 18.8 spell_power points (9.67 DPS) | yes | Spritecaster Cape (11623, -1.93 DPS) [dungeon]; Deep Woodlands Cloak (19121, -2.69 DPS) [quest]; Runecloth Cloak (13860, -4.32 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (24.48 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.86 DPS) [pvp]; Robe of Winter Night (14136, -7.37 DPS, sim-verified) [crafted]; Felcloth Robe (14106, -8.50 DPS) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (6.56 DPS) | yes | Netherpearl Cuffs (254067, -0.16 DPS) [crafted]; Spidertank Oilrag (9448, -1.92 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.33 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.4 spell_power points (11.54 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.38 DPS) [pvp]; Earth Warder's Gloves (21318, -3.60 DPS, sim-verified) [quest] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.0 spell_power points (9.24 DPS) | yes | Satyrmane Sash (17755, -1.14 DPS) [dungeon]; Ghostweave Cord (254073, -2.04 DPS) [crafted]; Black Cord (254043, -2.31 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (13.38 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Felcloth Pants (14107, -3.76 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.35 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.98 DPS) [dungeon]; Venomspew Footpads (275606, -1.03 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (9.89 DPS) | yes | Underworld Band (1980, -3.23 DPS) [world_drop]; Philanthropist's Ring (281635, -4.30 DPS) [quest]; Reedknot Ring (9622, -6.29 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.69 DPS) | yes | Underworld Band (1980, -0.03 DPS) [world_drop]; Philanthropist's Ring (281635, -1.10 DPS) [quest]; Reedknot Ring (9622, -3.09 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (572.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (572.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (572.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Inventor's Focal Sword (17719, -1.07 DPS) [dungeon]; Spire of Hakkar (10844, -3.05 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.1 spell_power points (79.85 DPS) | yes | Brilliant Wand (249385, -15.01 DPS) [crafted]; Stormrager (16997, -16.71 DPS) [quest]; Greater Eternal Wand (249237, -21.89 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; legs: Ghostweave Pants; finger1: Maiden's Circle; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

