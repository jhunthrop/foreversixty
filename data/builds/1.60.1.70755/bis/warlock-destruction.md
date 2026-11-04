# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 54.9. Weights run: 1.6s. Verify run: 0.8s. 186 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.045 ± 0.004, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.422 per %), hit=0.082 ± 0.004 per rating point (10 rating = 1%, 0.817 per %), spell_haste=0.311 ± 0.054, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.789 ± 0.001, fire_power=0.211 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.7 spell_power points (2.19 DPS) | yes | Pristine Circlet (253949, -0.88 DPS, sim-verified) [crafted]; Flame Circlet (253953, -1.60 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.26 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (1.01 DPS) | yes | Pearl-clasped Cloak (5542, -0.50 DPS) [crafted]; Black Whelp Cloak (7283, -0.61 DPS, sim-verified) [crafted]; Reinforced Linen Cape (2580, -0.76 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 11.1 spell_power points (2.78 DPS) | yes | Filigreed Shadow Gown (253907, -0.85 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.52 DPS) [crafted]; Gray Woolen Robe (2585, -1.78 DPS) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.25 DPS) | yes | Ivycloth Bracelets (9793, -0.29 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.76 DPS) | yes | Shadow Gloves (253919, -0.37 DPS) [crafted]; Gnoll Casting Gloves (892, -0.42 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.76 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 5.5 spell_power points (1.39 DPS) | yes | Pristine Sash (253925, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.60 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.89 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 8.7 spell_power points (2.19 DPS) | yes | Filigreed Pristine Leggings (253937, -0.68 DPS) [crafted]; Colorful Kilt (10048, -0.93 DPS) [crafted]; Silk-threaded Trousers (1929, -1.08 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.76 DPS) | yes | Shadow Boots (253895, -0.68 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.76 DPS) [crafted]; Pristine Boots (253889, -1.01 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.26 DPS) | yes | - |
| finger2 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | 1.0 spell_power points (0.25 DPS) | yes | Ring of the Shadow (1462, -0.29 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Witching Stave (1484) | Shadowfang Keep [dungeon] | 21.3 spell_power points (5.37 DPS) | yes | Shadowfang (1482, -6.50 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 88.6 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.75 DPS) [world_drop]; Deepblaze (279896, -3.78 DPS) [quest]; Sizzle Stick (8071, -4.16 DPS) [quest] |

**New at 20:** head: Shadow Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Furlfeather Ring; main_hand: Witching Stave; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 89.7. Weights run: 1.6s. Verify run: 1.2s. 323 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.147 ± 0.006, crit=0.076 ± 0.002 per rating point (14 rating = 1%, 1.067 per %), hit=0.125 ± 0.006 per rating point (10 rating = 1%, 1.250 per %), spell_haste=not significant (0.179 ± 0.094), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.787 ± 0.001, fire_power=0.213 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 11.0 spell_power points (3.03 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.65 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.9 spell_power points (0.24 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (2.84 DPS) | yes | Invoker's Mantle (215365, -0.58 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.74 DPS) [dungeon]; Moonlit Amice (11884, -0.91 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.37 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Black Whelp Cloak (7283, -0.55 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.57 DPS) | yes | Necrology Robes (2292, -0.34 DPS) [dungeon]; Shadow Gown (253967, -0.48 DPS) [crafted]; Green Silk Armor (7065, -0.57 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.47 DPS) | yes | Phoenix Bindings (210781, -1.77 DPS, sim-verified) [crafted]; Glowing Magical Bracelets (13106, -2.15 DPS) [world_drop]; Windsong Bangles (263336, -2.20 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.1 spell_power points (2.23 DPS) | yes | Hands of Darkness (7047, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.58 DPS) [world] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.4 spell_power points (2.59 DPS) | yes | Invoker's Cord (215366, -0.47 DPS) [crafted]; Crimson Silk Belt (7055, -0.66 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.69 DPS, sim-verified) [dungeon] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 12.1 spell_power points (3.31 DPS) | yes | Gaze Dreamer Pants (6903, -0.01 DPS) [dungeon]; Filigreed Shadow Leggings (253943, -0.69 DPS) [crafted]; Necromancer Leggings (2277, -0.70 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.3 spell_power points (3.10 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.08 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.17 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.7 spell_power points (1.85 DPS) | yes | Minor Channeling Ring (1449, -0.40 DPS) [quest]; Electrocutioner Lagnut (9447, -1.03 DPS) [dungeon]; Runemender's Seal (273031, -1.03 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.65 DPS) | yes | Minor Channeling Ring (1449, -0.19 DPS) [quest]; Electrocutioner Lagnut (9447, -0.82 DPS) [dungeon]; Runemender's Seal (273031, -0.82 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | 34.0 spell_power points (9.34 DPS) | yes | Staff of the Shade (2549, -1.19 DPS, sim-verified) [dungeon]; Witching Stave (1484, -3.30 DPS) [dungeon]; Glimmering Staff (249392, -8.90 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 120.2 spell_power points (33.02 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.64 DPS) [crafted]; Gyromatic Icemaker (9489, -2.66 DPS) [dungeon] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Belt of Arugal; legs: Shadow Leggings; feet: Black Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Staff of Soran'ruk; ranged: Starfaller

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 141.7. Weights run: 1.6s. Verify run: 1.6s. 420 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.473 ± 0.013, crit=0.111 ± 0.003 per rating point (14 rating = 1%, 1.555 per %), hit=0.236 ± 0.014 per rating point (10 rating = 1%, 2.358 per %), spell_haste=not significant (0.408 ± 0.172), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.466 ± 0.001, fire_power=0.534 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.14 DPS) | yes | Augural Shroud (2620, -1.54 DPS) [world]; Living Cowl (5608, -2.34 DPS) [world]; Papal Fez (9431, -2.41 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.8 spell_power points (3.15 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.18 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.32 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.3 spell_power points (3.88 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.03 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.4 spell_power points (2.45 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.40 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.98 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.8 spell_power points (7.26 DPS) | yes | Dreamweave Vest (10021, -0.62 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.12 DPS) [world_drop]; Robe of Power (7054, -1.51 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.63 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.58 DPS) [quest]; Phoenix Bindings (210781, -0.60 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.9 spell_power points (5.81 DPS) | yes | Black Mageweave Gloves (10003, -1.43 DPS) [crafted]; Red Mageweave Gloves (10018, -1.83 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.51 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | sim-verified (141.7 DPS) | yes | Gilded Cord (254037, -0.36 DPS) [crafted]; Fiery Cord (254041, -0.51 DPS) [crafted]; Deathmage Sash (10771, -1.48 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.7 spell_power points (5.75 DPS) | yes | Crimson Silk Pantaloons (7062, -1.15 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.24 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.60 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.02 DPS) | yes | Fiery Slippers (254005, -4.02 DPS) [crafted]; Black Slippers (254007, -4.28 DPS) [crafted]; Gilded Slippers (254001, -4.81 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (3.61 DPS) | yes | Snake Hoop (6750, -1.17 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.28 DPS) [quest]; Reedknot Ring (9622, -1.57 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 8.9 spell_power points (2.60 DPS) | yes | Snake Hoop (6750, -0.15 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.26 DPS) [quest]; Reedknot Ring (9622, -0.55 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Noh'Orahil (15105, -2.85 DPS) [quest]; Searing Golden Blade (12260, -2.89 DPS) [crafted]; Staff of Dar'Orahil (15106, -3.35 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 137.5 spell_power points (40.20 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.45 DPS) [crafted]; Umbral Wand (5216, -4.54 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Underworld Band; ranged: Jaina's Firestarter

No-known-source sample (15 of 420, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 237.8. Weights run: 1.6s. Verify run: 1.2s. 531 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.323 ± 0.015, crit=0.160 ± 0.004 per rating point (14 rating = 1%, 2.243 per %), hit=0.337 ± 0.018 per rating point (10 rating = 1%, 3.366 per %), spell_haste=not significant (-0.179 ± 0.222), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.403 ± 0.000, fire_power=0.597 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.06 DPS) | yes | Dreamweave Circlet (10041, -1.03 DPS) [crafted]; Red Mageweave Headband (10033, -1.57 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.23 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.6 spell_power points (3.57 DPS) | yes | Triune Amulet (7722, -0.59 DPS) [dungeon]; Mindburst Medallion (11196, -0.61 DPS) [quest]; Horizon Choker (13085, -1.89 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.8 spell_power points (7.01 DPS) | yes | Black Mageweave Shoulders (10027, -1.99 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.11 DPS) [vendor]; Netherflame Shoulders (254053, -2.27 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (5.94 DPS) | yes | Cindercloth Cloak (14044, -2.08 DPS) [crafted]; Nightfall Drape (12465, -2.58 DPS) [dungeon]; Runecloth Cloak (13860, -3.35 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.5 spell_power points (9.48 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.66 DPS) [world_drop]; Dreamweave Vest (10021, -1.70 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.0 spell_power points (3.73 DPS) | yes | Nethergeld Cuffs (254061, -0.28 DPS) [crafted]; Arcane Runed Bracers (4744, -0.38 DPS) [quest]; Spidertank Oilrag (9448, -0.38 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (7.18 DPS) | yes | Black Mageweave Gloves (10003, -1.60 DPS) [crafted]; Runecloth Gloves (13863, -1.63 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.91 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.2 spell_power points (6.42 DPS) | yes | Fiery Waistcord (254085, -1.11 DPS) [crafted]; Ghostweave Cord (254073, -1.20 DPS) [crafted]; Star Belt (4329, -1.58 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (7.08 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.33 DPS) [vendor]; Red Mageweave Pants (10009, -0.42 DPS) [crafted]; Felcloth Pants (14107, -1.73 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.94 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -3.62 DPS) [vendor]; Gilded Sandals (254107, -3.76 DPS) [crafted]; Fiery Sandals (254111, -3.96 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.84 DPS) | yes | Ring of Forlorn Spirits (2043, -1.86 DPS) [quest]; Snake Hoop (6750, -2.01 DPS) [quest]; Underworld Band (1980, -2.14 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.6 spell_power points (4.33 DPS) | yes | Snake Hoop (6750, -1.49 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.51 DPS, sim-verified) [quest]; Underworld Band (1980, -1.62 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (237.8 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (237.8 DPS) | yes | Mark of the Chosen (17774, -2.40 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (237.8 DPS) | yes | Inventor's Focal Sword (17719, -1.44 DPS) [dungeon]; Spire of Hakkar (10844, -1.53 DPS) [world]; Spellshifter Rod (9527, -1.75 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 148.7 spell_power points (55.39 DPS) | yes | Noxious Shooter (17745, -1.80 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -6.66 DPS) [crafted]; Wand of Allistarj (13065, -6.97 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 531, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 359.5. Weights run: 1.7s. Verify run: 2.3s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.248 ± 0.019, crit=0.212 ± 0.006 per rating point (14 rating = 1%, 2.965 per %), hit=0.450 ± 0.024 per rating point (10 rating = 1%, 4.500 per %), spell_haste=not significant (-0.753 ± 0.317), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.397 ± 0.001, fire_power=0.603 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.35 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -1.28 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.5 spell_power points (6.09 DPS) | yes | Glowing Eye of Mordresh (10769, -2.31 DPS) [dungeon]; Triune Amulet (7722, -2.72 DPS) [dungeon]; Mindburst Medallion (11196, -2.94 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.51 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.76 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.86 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (6.51 DPS) | yes | Mageflame Cloak (13007, -1.18 DPS) [world_drop]; Runecloth Cloak (13860, -1.89 DPS) [crafted]; Cindercloth Cloak (14044, -2.38 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (20.27 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.32 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.68 DPS) [pvp]; Acumen Robes (17775, -10.20 DPS) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.5 spell_power points (5.67 DPS) | yes | Netherflame Cuffs (254065, -1.64 DPS) [crafted]; Arcane Runed Bracers (4744, -1.88 DPS) [quest]; Spidertank Oilrag (9448, -1.88 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (9.25 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.00 DPS) [pvp]; Dreamweave Gloves (10019, -1.26 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (6.93 DPS) | yes | Wisdom of the Timbermaw (19047, -0.84 DPS) [crafted]; Ghostweave Cord (254073, -1.04 DPS) [crafted]; Fiery Waistcord (254085, -1.17 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (10.93 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Leggings (17567, -2.12 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.09 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.43 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.54 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (8.30 DPS) | yes | Philanthropist's Ring (281635, -3.57 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.93 DPS) [quest]; Snake Hoop (6750, -5.25 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.46 DPS) | yes | Philanthropist's Ring (281635, -0.74 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.10 DPS) [quest]; Snake Hoop (6750, -2.42 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (359.5 DPS) | yes | - |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (359.5 DPS) | yes | Elemental Mage Staff (944, +0.00 DPS) [world_drop]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Inventor's Focal Sword (17719, -1.25 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 188.6 spell_power points (79.29 DPS) | yes | Brilliant Wand (249385, -15.19 DPS) [crafted]; Stormrager (16997, -16.07 DPS) [quest]; Greater Eternal Wand (249237, -22.17 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25500000020000000-0000000000000000000-2053345103101351)

Set DPS (verified): 667.9. Weights run: 1.8s. Verify run: 2.5s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.136 ± 0.025, crit=0.335 ± 0.009 per rating point (14 rating = 1%, 4.696 per %), hit=0.742 ± 0.039 per rating point (10 rating = 1%, 7.420 per %), spell_haste=not significant (0.381 ± 0.463), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.324 ± 0.001, fire_power=0.676 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.62 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.85 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.4 spell_power points (8.22 DPS) | yes | Glowing Eye of Mordresh (10769, -3.24 DPS) [dungeon]; Triune Amulet (7722, -3.30 DPS) [dungeon]; Mindburst Medallion (11196, -4.03 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (15.38 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.94 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -4.02 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.8 spell_power points (9.12 DPS) | yes | Mageflame Cloak (13007, -0.39 DPS) [world_drop]; Runecloth Cloak (13860, -2.91 DPS) [crafted]; Cindercloth Cloak (14044, -3.04 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.2 spell_power points (29.06 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.36 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.00 DPS) [pvp]; Robe of the Magi (1716, -13.79 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.4 spell_power points (7.61 DPS) | yes | Netherflame Cuffs (254065, -1.62 DPS) [crafted]; Arcane Runed Bracers (4744, -2.07 DPS) [quest]; Spidertank Oilrag (9448, -2.07 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.54 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS, sim-verified) [crafted]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.28 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.4 spell_power points (9.45 DPS) | yes | Fiery Waistcord (254085, -0.80 DPS) [crafted]; Ghostweave Cord (254073, -0.84 DPS) [crafted]; Star Belt (4329, -1.45 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (16.00 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -4.91 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.77 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.23 DPS) [crafted]; Omnicast Boots (11822, -1.46 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.0 spell_power points (11.66 DPS) | yes | Philanthropist's Ring (281635, -5.09 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.74 DPS) [quest]; Blackstone Ring (17713, -7.10 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (8.00 DPS) | yes | Philanthropist's Ring (281635, -1.43 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.08 DPS) [quest]; Blackstone Ring (17713, -3.43 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (667.9 DPS) | yes | Uther's Strength (11302, -0.42 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (667.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (667.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -2.76 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -3.94 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 130.7 spell_power points (80.46 DPS) | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -17.35 DPS) [quest]; Greater Eternal Wand (249237, -21.59 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 55.1. Weights run: 1.6s. Verify run: 0.8s. 180 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.045 ± 0.004, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.422 per %), hit=0.082 ± 0.004 per rating point (10 rating = 1%, 0.817 per %), spell_haste=0.311 ± 0.054, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.789 ± 0.001, fire_power=0.211 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.7 spell_power points (2.19 DPS) | yes | Pristine Circlet (253949, -0.88 DPS, sim-verified) [crafted]; Flame Circlet (253953, -1.60 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.26 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.33 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (1.01 DPS) | yes | Mystic Shawl (3449, -0.50 DPS) [quest]; Pearl-clasped Cloak (5542, -0.50 DPS) [crafted]; Black Whelp Cloak (7283, -0.77 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 11.1 spell_power points (2.78 DPS) | yes | Filigreed Shadow Gown (253907, -1.01 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.52 DPS) [crafted]; Gray Woolen Robe (2585, -1.78 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.76 DPS) | yes | Owlbeard Bracers (16981, -0.50 DPS) [quest]; Windsong Bangles (263336, -0.57 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.76 DPS) | yes | Shadow Gloves (253919, -0.37 DPS) [crafted]; Gnoll Casting Gloves (892, -0.43 DPS, sim-verified) [world]; Blight Gloves (279877, -0.50 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 5.5 spell_power points (1.39 DPS) | yes | Pristine Sash (253925, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.68 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.89 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 8.7 spell_power points (2.19 DPS) | yes | Filigreed Pristine Leggings (253937, -0.68 DPS) [crafted]; Colorful Kilt (10048, -0.93 DPS) [crafted]; Silk-threaded Trousers (1929, -1.05 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.76 DPS) | yes | Red Woolen Boots (4313, -0.76 DPS) [crafted]; Shadow Boots (253895, -0.76 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.01 DPS) [crafted] |
| finger1 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | 1.0 spell_power points (0.25 DPS) | yes | Ring of the Shadow (1462, -0.29 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.76 DPS) | yes | Ring of the Shadow (1462, -1.15 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Witching Stave (1484) | Shadowfang Keep [dungeon] | 21.3 spell_power points (5.37 DPS) | yes | Shadowfang (1482, -6.68 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 88.6 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.75 DPS) [world_drop]; Sizzle Stick (8071, -4.16 DPS) [quest]; Greater Magic Wand (11288, -4.30 DPS) [crafted] |

**New at 20:** head: Shadow Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Furlfeather Ring; finger2: Sludge-Stained Band; main_hand: Witching Stave; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 180, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 89.2. Weights run: 1.6s. Verify run: 1.2s. 314 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.147 ± 0.006, crit=0.076 ± 0.002 per rating point (14 rating = 1%, 1.067 per %), hit=0.125 ± 0.006 per rating point (10 rating = 1%, 1.250 per %), spell_haste=not significant (0.179 ± 0.094), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.787 ± 0.001, fire_power=0.213 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 11.0 spell_power points (3.03 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.65 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.9 spell_power points (0.24 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (2.84 DPS) | yes | Death Speaker Mantle (6685, -0.74 DPS) [dungeon]; Invoker's Mantle (215365, -0.90 DPS, sim-verified) [crafted]; Magician's Mantle (12998, -1.10 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.37 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Black Whelp Cloak (7283, -0.55 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.57 DPS) | yes | Necrology Robes (2292, -0.42 DPS, sim-verified) [dungeon]; Shadow Gown (253967, -0.48 DPS) [crafted]; Green Silk Armor (7065, -0.57 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.47 DPS) | yes | Phoenix Bindings (210781, -1.71 DPS) [crafted]; Owlbeard Bracers (16981, -2.12 DPS) [quest]; Tabitha's Cuffs (251486, -2.20 DPS, sim-verified) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.1 spell_power points (2.23 DPS) | yes | Hands of Darkness (7047, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS) [dungeon]; Jutebraid Gloves (10654, -0.38 DPS) [quest] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (3.02 DPS) | yes | Belt of Arugal (6392, -0.43 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.82 DPS) [dungeon]; Invoker's Cord (215366, -0.90 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 12.1 spell_power points (3.31 DPS) | yes | Gaze Dreamer Pants (6903, -0.01 DPS) [dungeon]; Filigreed Shadow Leggings (253943, -0.69 DPS) [crafted]; Necromancer Leggings (2277, -0.70 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.3 spell_power points (3.10 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.05 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.17 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.7 spell_power points (1.85 DPS) | yes | Solid Gold Band (279845, -0.75 DPS) [quest]; Electrocutioner Lagnut (9447, -1.03 DPS) [dungeon]; Runemender's Seal (273031, -1.03 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.65 DPS) | yes | Solid Gold Band (279845, -0.77 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.82 DPS) [dungeon]; Runemender's Seal (273031, -0.82 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | 34.0 spell_power points (9.34 DPS) | yes | Staff of the Shade (2549, -1.20 DPS, sim-verified) [dungeon]; Witching Stave (1484, -3.30 DPS) [dungeon]; Glimmering Staff (249392, -8.90 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 120.2 spell_power points (33.02 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.64 DPS) [crafted]; Gyromatic Icemaker (9489, -2.66 DPS) [dungeon] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Warsong Sash; legs: Shadow Leggings; feet: Black Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Staff of Soran'ruk; ranged: Starfaller

No-known-source sample (15 of 314, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 141.4. Weights run: 1.6s. Verify run: 1.6s. 411 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.473 ± 0.013, crit=0.111 ± 0.003 per rating point (14 rating = 1%, 1.555 per %), hit=0.236 ± 0.014 per rating point (10 rating = 1%, 2.358 per %), spell_haste=not significant (0.408 ± 0.172), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.466 ± 0.001, fire_power=0.534 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.14 DPS) | yes | Augural Shroud (2620, -1.54 DPS) [world]; Papal Fez (9431, -2.09 DPS, sim-verified) [dungeon]; Living Cowl (5608, -2.34 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.8 spell_power points (3.15 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.18 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.32 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.3 spell_power points (3.88 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Inquisitor's Shawl (19507, -0.03 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.4 spell_power points (2.45 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.40 DPS) [crafted]; Hillman's Cloak (3719, -0.98 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.8 spell_power points (7.26 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.12 DPS) [world_drop]; Robe of Power (7054, -1.51 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.63 DPS) | yes | Radiant Silver Bracers (4545, -0.36 DPS) [quest]; Condor Bracers (15864, -0.58 DPS) [quest]; Phoenix Bindings (210781, -0.60 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.9 spell_power points (5.81 DPS) | yes | Black Mageweave Gloves (10003, -1.43 DPS) [crafted]; Red Mageweave Gloves (10018, -2.19 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.51 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | sim-verified (141.4 DPS) | yes | Gilded Cord (254037, -0.36 DPS) [crafted]; Fiery Cord (254041, -0.51 DPS) [crafted]; Deathmage Sash (10771, -2.27 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.7 spell_power points (5.75 DPS) | yes | Crimson Silk Pantaloons (7062, -1.94 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.24 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.60 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.02 DPS) | yes | Fiery Slippers (254005, -4.02 DPS) [crafted]; Gilded Slippers (254001, -4.22 DPS, sim-verified) [crafted]; Black Slippers (254007, -4.28 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (3.61 DPS) | yes | Snake Hoop (6750, -1.17 DPS) [quest]; Reedknot Ring (9622, -1.57 DPS) [quest]; Ring of Power Regulation (273030, -1.86 DPS) [dungeon] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 8.9 spell_power points (2.60 DPS) | yes | Snake Hoop (6750, -0.15 DPS) [quest]; Reedknot Ring (9622, -0.55 DPS) [quest]; Ring of Power Regulation (273030, -0.85 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Dar'Orahil (15106, -2.23 DPS, sim-verified) [quest]; Staff of Noh'Orahil (15105, -2.85 DPS) [quest]; Searing Golden Blade (12260, -2.89 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 137.5 spell_power points (40.20 DPS) | yes | Earthen Rod (9381, -3.74 DPS) [dungeon]; Twisted Nether Wand (249144, -4.45 DPS) [crafted]; Umbral Wand (5216, -4.54 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Underworld Band; ranged: Jaina's Firestarter

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 236.5. Weights run: 1.6s. Verify run: 1.6s. 522 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.323 ± 0.015, crit=0.160 ± 0.004 per rating point (14 rating = 1%, 2.243 per %), hit=0.337 ± 0.018 per rating point (10 rating = 1%, 3.366 per %), spell_haste=not significant (-0.179 ± 0.222), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.403 ± 0.000, fire_power=0.597 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.06 DPS) | yes | Dreamweave Circlet (10041, -1.03 DPS) [crafted]; Red Mageweave Headband (10033, -1.58 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.23 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.6 spell_power points (3.57 DPS) | yes | Triune Amulet (7722, -0.59 DPS) [dungeon]; Mindburst Medallion (11196, -0.61 DPS) [quest]; Horizon Choker (13085, -1.89 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.8 spell_power points (7.01 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -2.11 DPS) [vendor]; Netherflame Shoulders (254053, -2.27 DPS) [crafted]; Black Mageweave Shoulders (10027, -2.41 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (5.94 DPS) | yes | Deep Woodlands Cloak (19121, -1.01 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.62 DPS) [crafted]; Cindercloth Cloak (14044, -2.08 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.5 spell_power points (9.48 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.66 DPS) [world_drop]; Dreamweave Vest (10021, -1.70 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.0 spell_power points (3.73 DPS) | yes | Nethergeld Cuffs (254061, -0.28 DPS) [crafted]; Spidertank Oilrag (9448, -0.38 DPS) [dungeon]; Bloodband Bracers (11469, -0.79 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (7.18 DPS) | yes | Black Mageweave Gloves (10003, -1.60 DPS) [crafted]; Runecloth Gloves (13863, -1.63 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.26 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.2 spell_power points (6.42 DPS) | yes | Fiery Waistcord (254085, -1.11 DPS) [crafted]; Ghostweave Cord (254073, -1.20 DPS) [crafted]; Star Belt (4329, -1.58 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (7.08 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.33 DPS) [vendor]; Red Mageweave Pants (10009, -0.42 DPS) [crafted]; Felcloth Pants (14107, -1.73 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.94 DPS) | yes | Fiery Sandals (254111, -3.31 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.62 DPS) [vendor]; Gilded Sandals (254107, -3.76 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.84 DPS) | yes | Snake Hoop (6750, -2.01 DPS) [quest]; Underworld Band (1980, -2.14 DPS) [world_drop]; Reedknot Ring (9622, -2.23 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.6 spell_power points (4.33 DPS) | yes | Snake Hoop (6750, -1.61 DPS, sim-verified) [quest]; Underworld Band (1980, -1.62 DPS) [world_drop]; Reedknot Ring (9622, -1.72 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (236.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (236.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (236.5 DPS) | yes | Inventor's Focal Sword (17719, -0.96 DPS, sim-verified) [dungeon]; Spire of Hakkar (10844, -1.53 DPS) [world]; Spellshifter Rod (9527, -1.75 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 148.7 spell_power points (55.39 DPS) | yes | Noxious Shooter (17745, -1.28 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -6.66 DPS) [crafted]; Wand of Allistarj (13065, -6.97 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 522, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 357.0. Weights run: 1.7s. Verify run: 2.4s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.248 ± 0.019, crit=0.212 ± 0.006 per rating point (14 rating = 1%, 2.965 per %), hit=0.450 ± 0.024 per rating point (10 rating = 1%, 4.500 per %), spell_haste=not significant (-0.753 ± 0.317), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.397 ± 0.001, fire_power=0.603 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.35 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -2.19 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.5 spell_power points (6.09 DPS) | yes | Glowing Eye of Mordresh (10769, -2.56 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -2.72 DPS) [dungeon]; Mindburst Medallion (11196, -2.94 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.51 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.76 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -2.86 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (6.51 DPS) | yes | Mageflame Cloak (13007, -1.18 DPS) [world_drop]; Runecloth Cloak (13860, -1.89 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.21 DPS, sim-verified) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (20.27 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.32 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.68 DPS) [pvp]; Acumen Robes (17775, -10.76 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.5 spell_power points (5.67 DPS) | yes | Netherflame Cuffs (254065, -1.64 DPS) [crafted]; Spidertank Oilrag (9448, -1.88 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.99 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (9.25 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.00 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (6.93 DPS) | yes | Ghostweave Cord (254073, -1.04 DPS) [crafted]; Fiery Waistcord (254085, -1.17 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.39 DPS, sim-verified) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (10.93 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -6.05 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.09 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.84 DPS) [crafted]; Omnicast Boots (11822, -1.40 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (8.30 DPS) | yes | Philanthropist's Ring (281635, -3.57 DPS) [quest]; Snake Hoop (6750, -5.25 DPS) [quest]; Reedknot Ring (9622, -5.35 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.46 DPS) | yes | Philanthropist's Ring (281635, -0.74 DPS) [quest]; Snake Hoop (6750, -2.42 DPS) [quest]; Reedknot Ring (9622, -2.52 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (357.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (357.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (357.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.36 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.61 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 188.6 spell_power points (79.29 DPS) | yes | Brilliant Wand (249385, -15.19 DPS) [crafted]; Stormrager (16997, -16.07 DPS) [quest]; Greater Eternal Wand (249237, -22.17 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25500000020000000-0000000000000000000-2053345103101351)

Set DPS (verified): 669.8. Weights run: 1.8s. Verify run: 2.9s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.136 ± 0.025, crit=0.335 ± 0.009 per rating point (14 rating = 1%, 4.696 per %), hit=0.742 ± 0.039 per rating point (10 rating = 1%, 7.420 per %), spell_haste=not significant (0.381 ± 0.463), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.324 ± 0.001, fire_power=0.676 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.62 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.85 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.4 spell_power points (8.22 DPS) | yes | Glowing Eye of Mordresh (10769, -3.24 DPS) [dungeon]; Triune Amulet (7722, -3.30 DPS) [dungeon]; Mindburst Medallion (11196, -4.03 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (15.38 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS, sim-verified) [crafted]; Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -4.02 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.8 spell_power points (9.12 DPS) | yes | Mageflame Cloak (13007, +0.00 DPS, sim-verified) [world_drop]; Deep Woodlands Cloak (19121, -0.98 DPS) [quest]; Runecloth Cloak (13860, -2.91 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.2 spell_power points (29.06 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.36 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.00 DPS) [pvp]; Robe of the Magi (1716, -12.50 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.4 spell_power points (7.61 DPS) | yes | Netherflame Cuffs (254065, -1.62 DPS) [crafted]; Spidertank Oilrag (9448, -2.07 DPS) [dungeon]; Phoenix Bindings (210781, -2.20 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.54 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS, sim-verified) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.28 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.4 spell_power points (9.45 DPS) | yes | Fiery Waistcord (254085, +0.00 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -0.84 DPS) [crafted]; Star Belt (4329, -1.45 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (16.00 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -5.15 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.77 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.23 DPS) [crafted]; Omnicast Boots (11822, -1.46 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.0 spell_power points (11.66 DPS) | yes | Philanthropist's Ring (281635, -5.09 DPS) [quest]; Blackstone Ring (17713, -7.10 DPS) [dungeon]; Reedknot Ring (9622, -7.36 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (8.00 DPS) | yes | Philanthropist's Ring (281635, +0.00 DPS, sim-verified) [quest]; Blackstone Ring (17713, -3.43 DPS) [dungeon]; Reedknot Ring (9622, -3.69 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (669.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (669.8 DPS) | yes | Uther's Strength (11302, -0.42 DPS) [world_drop] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (669.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -2.93 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.94 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 130.7 spell_power points (80.46 DPS) | yes | Brilliant Wand (249385, -14.75 DPS) [crafted]; Stormrager (16997, -17.35 DPS) [quest]; Greater Eternal Wand (249237, -21.59 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

