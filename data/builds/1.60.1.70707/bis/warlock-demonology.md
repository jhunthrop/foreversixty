# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 44.1. Weights run: 1.7s. Verify run: 1.1s. 161 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.998 ± 0.027, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.385 per %), hit=0.083 ± 0.004 per rating point (10 rating = 1%, 0.834 per %), spell_haste=not significant (-0.103 ± 0.064), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.724 ± 0.002, fire_power=0.276 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.25 DPS) | yes | Flame Circlet (253953, -0.62 DPS) [crafted]; Shadow Goggles (4373, -2.54 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.0 spell_power points (2.92 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.38 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -2.08 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (44.1 DPS) | yes | Sanguine Cape (14376, -0.00 DPS) [world_drop]; Black Whelp Cloak (7283, -0.21 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.67 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 10.0 spell_power points (2.09 DPS) | yes | Filigreed Flame Gown (253905, -0.52 DPS) [crafted]; Mystic's Wrap (14369, -0.63 DPS) [world_drop]; Gray Woolen Robe (2585, -1.63 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 5.0 spell_power points (1.04 DPS) | yes | Bright Bracers (3647, -0.36 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -0.63 DPS) [world_drop]; Repurposed Hair Band (281256, -0.63 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.46 DPS) | yes | Pristine Gloves (253913, -0.00 DPS) [crafted]; Gnoll Casting Gloves (892, -0.21 DPS) [world]; Tomb Robber's Gloves (280096, -0.21 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 8.0 spell_power points (1.67 DPS) | yes | Novice Arcanist's Sash (253885, -0.21 DPS) [crafted]; Flame Sash (253929, -0.43 DPS) [crafted]; Keller's Girdle (2911, -1.76 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 12.0 spell_power points (2.50 DPS) | yes | Silk-threaded Trousers (1929, -1.04 DPS) [dungeon]; Darkweave Breeches (12987, -1.04 DPS) [world_drop]; Filigreed Flame Leggings (253941, -2.11 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.0 spell_power points (2.30 DPS) | yes | Pristine Boots (253889, -0.41 DPS, sim-verified) [crafted]; Flame Boots (253893, -1.32 DPS) [crafted]; Red Woolen Boots (4313, -1.46 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (1.46 DPS) | yes | Volcanic Rock Ring (12053, -0.84 DPS) [world_drop]; The 1 Ring (8350, -1.25 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.0 spell_power points (1.25 DPS) | yes | Volcanic Rock Ring (12053, -0.22 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -1.04 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 10.0 spell_power points (2.08 DPS) | yes | Lesser Staff of the Spire (1300, -0.83 DPS) [world]; Witching Stave (1484, -1.04 DPS) [dungeon]; Staff of Westfall (2042, -1.04 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 116.3 spell_power points (24.28 DPS) | yes | Cookie's Stirring Rod (5198, -0.49 DPS, sim-verified) [dungeon]; Sable Wand (7607, -1.43 DPS) [quest]; Skycaller (12984, -2.72 DPS) [world_drop] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 161, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 81.6. Weights run: 1.8s. Verify run: 1.1s. 298 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.033 ± 0.005, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.582 per %), hit=0.113 ± 0.005 per rating point (10 rating = 1%, 1.134 per %), spell_haste=not significant (0.286 ± 0.076), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.734 ± 0.002, fire_power=0.266 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.54 DPS) | yes | Filigreed Pristine Circlet (253975, -0.69 DPS) [crafted]; Enchanter's Cowl (4322, -1.08 DPS) [crafted]; Silk Headband (7050, -3.75 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.2 spell_power points (0.05 DPS) | yes | Kaleidoscope Chain (13084, -0.02 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.34 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.3 spell_power points (2.15 DPS) | yes | Moonlit Amice (11884, -0.53 DPS) [quest]; Death Speaker Mantle (6685, -0.68 DPS) [dungeon]; Invoker's Mantle (215365, -0.70 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.15 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Prelacy Cape (7004, -0.23 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.00 DPS) | yes | Green Silk Armor (7065, -0.60 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.83 DPS) [dungeon]; Mechbuilder's Overalls (9508, -1.06 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.08 DPS) | yes | Phoenix Bindings (210781, -1.66 DPS, sim-verified) [crafted]; Windsong Bangles (263336, -1.85 DPS) [quest]; Glowing Magical Bracelets (13106, -2.02 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.62 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.44 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.1 spell_power points (2.10 DPS) | yes | Invoker's Cord (215366, -0.45 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.57 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.66 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (81.6 DPS) | yes | Silk-threaded Trousers (1929, -0.05 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.24 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.82 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.2 spell_power points (1.67 DPS) | yes | Glinteye Slippers (273024, -0.05 DPS) [dungeon]; Nimbus Boots (6998, -0.29 DPS) [quest]; Spidersilk Boots (4320, -1.85 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.2 spell_power points (1.42 DPS) | yes | Minor Channeling Ring (1449, -0.25 DPS) [quest]; Electrocutioner Lagnut (9447, -0.73 DPS) [dungeon]; Runemender's Seal (273031, -0.73 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.39 DPS) | yes | Minor Channeling Ring (1449, -0.36 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.69 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.32 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.70 DPS) [quest]; Staff of Soran'ruk (15109, -6.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 142.4 spell_power points (32.89 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Gyromatic Icemaker (9489, -2.53 DPS) [dungeon]; Greater Mystic Wand (217287, -2.73 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 121.8. Weights run: 1.8s. Verify run: 1.2s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.208 ± 0.007, crit=0.060 ± 0.002 per rating point (14 rating = 1%, 0.839 per %), hit=0.176 ± 0.007 per rating point (10 rating = 1%, 1.761 per %), spell_haste=0.767 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.764 ± 0.002, fire_power=0.236 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.65 DPS) | yes | Augural Shroud (2620, -1.80 DPS, sim-verified) [world]; Living Cowl (5608, -2.15 DPS) [world]; Papal Fez (9431, -2.55 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.7 spell_power points (2.33 DPS) | yes | Triune Amulet (7722, -0.53 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.94 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.00 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.9 spell_power points (2.93 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Inquisitor's Shawl (19507, -0.32 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.0 spell_power points (1.90 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.01 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.55 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 spell_power points (6.26 DPS) | yes | Elemental Raiment (9434, -0.60 DPS) [world_drop]; Dreamweave Vest (10021, -0.91 DPS) [crafted]; Robe of Power (7054, -1.82 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.42 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.54 DPS) [quest]; Earthen Silk Cuffs (254019, -1.35 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.07 DPS) | yes | Black Mageweave Gloves (10003, -0.95 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.55 DPS) [crafted]; Gilded Handwraps (254021, -2.53 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.50 DPS) | yes | Gilded Cord (254037, -0.90 DPS) [crafted]; Belt of Arugal (6392, -0.91 DPS) [dungeon]; Deathmage Sash (10771, -1.09 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.5 spell_power points (4.44 DPS) | yes | Gaze Dreamer Pants (6903, -0.46 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.56 DPS) [crafted]; Stoneweaver Leggings (9407, -2.11 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.46 DPS) | yes | Gilded Slippers (254001, -2.95 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.35 DPS) [crafted]; Glinteye Slippers (273024, -4.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (2.97 DPS) | yes | Snake Hoop (6750, -1.08 DPS) [quest]; Reedknot Ring (9622, -1.09 DPS) [quest]; Ring of Power Regulation (273030, -1.36 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (2.15 DPS) | yes | Snake Hoop (6750, -0.26 DPS) [quest]; Reedknot Ring (9622, -0.27 DPS) [quest]; Ring of Power Regulation (273030, -0.54 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.11 DPS) [dungeon]; Spine of the Devourer (273040, -2.60 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.77 DPS) | yes | Orb of Lorica (11262, -1.60 DPS) [quest]; Beacon of Hope (9393, -1.62 DPS) [dungeon]; Orb of Mystic Insight (249394, -1.77 DPS, sim-verified) [crafted] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | sim-verified (121.8 DPS) | yes | Umbral Wand (5216, -0.73 DPS) [dungeon]; Twisted Nether Wand (249144, -0.78 DPS) [crafted]; Jaina's Firestarter (13064, -1.60 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Earthen Rod

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 195.5. Weights run: 1.8s. Verify run: 1.4s. 507 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.266 ± 0.012, crit=0.091 ± 0.003 per rating point (14 rating = 1%, 1.268 per %), hit=0.368 ± 0.013 per rating point (10 rating = 1%, 3.680 per %), spell_haste=not significant (0.194 ± 0.161), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.700 ± 0.001, fire_power=0.300 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.04 DPS) | yes | Dreamweave Circlet (10041, -0.99 DPS) [crafted]; Red Mageweave Headband (10033, -1.43 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.79 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.1 spell_power points (2.72 DPS) | yes | Mindburst Medallion (11196, -0.46 DPS) [quest]; Triune Amulet (7722, -0.78 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.61 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.8 spell_power points (5.30 DPS) | yes | Black Mageweave Shoulders (10027, -1.02 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.83 DPS) [vendor]; Bloodmage Mantle (7684, -1.91 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (4.65 DPS) | yes | Runecloth Cloak (13860, -1.89 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.97 DPS) [dungeon]; Big Voodoo Cloak (8216, -2.44 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.3 spell_power points (7.25 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.99 DPS) [world_drop]; Dreamweave Vest (10021, -1.17 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.68 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.04 DPS) [crafted]; Bloodband Bracers (11469, -0.48 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (5.68 DPS) | yes | Black Mageweave Gloves (10003, -1.21 DPS) [crafted]; Runecloth Gloves (13863, -1.39 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.13 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.7 spell_power points (4.96 DPS) | yes | Ghostweave Cord (254073, -0.79 DPS) [crafted]; Star Belt (4329, -1.09 DPS) [crafted]; Gilded Waistcord (254081, -1.27 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.66 DPS) | yes | Red Mageweave Pants (10009, -0.54 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -0.76 DPS) [vendor]; Senior Designer's Pantaloons (11841, -1.53 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.15 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -3.06 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -3.16 DPS) [crafted]; Black Mageweave Boots (10026, -3.32 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.87 DPS) | yes | Ring of Forlorn Spirits (2043, -1.49 DPS) [quest]; Snake Hoop (6750, -1.69 DPS) [quest]; Reedknot Ring (9622, -1.79 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.3 spell_power points (3.38 DPS) | yes | Ring of Forlorn Spirits (2043, -0.74 DPS, sim-verified) [quest]; Snake Hoop (6750, -1.19 DPS) [quest]; Reedknot Ring (9622, -1.29 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (195.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (195.5 DPS) | yes | Mark of the Chosen (17774, -3.72 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (195.5 DPS) | yes | Spire of Hakkar (10844, -0.36 DPS) [world]; Spellshifter Rod (9527, -0.92 DPS) [quest]; Hanzo Sword (8190, -24.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 176.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.37 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 343.4. Weights run: 2.0s. Verify run: 2.4s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.379 ± 0.027, crit=0.131 ± 0.005 per rating point (14 rating = 1%, 1.837 per %), hit=0.554 ± 0.025 per rating point (10 rating = 1%, 5.538 per %), spell_haste=not significant (1.104 ± 0.324), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.746 ± 0.002, fire_power=0.254 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.23 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -2.57 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.8 spell_power points (5.98 DPS) | yes | Glowing Eye of Mordresh (10769, -2.18 DPS) [dungeon]; Mindburst Medallion (11196, -2.85 DPS) [quest]; Triune Amulet (7722, -2.95 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.48 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Rotgrip Mantle (17732, -1.96 DPS) [dungeon]; Mantle of the Timbermaw (19050, -2.28 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.3 spell_power points (6.17 DPS) | yes | Nightfall Drape (12465, -2.76 DPS) [dungeon]; Big Voodoo Cloak (8216, -2.98 DPS) [crafted]; Runecloth Cloak (13860, -3.27 DPS, sim-verified) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.4 spell_power points (18.73 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.15 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -6.38 DPS) [pvp]; Acumen Robes (17775, -9.98 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 14.8 spell_power points (5.60 DPS) | yes | Nethergeld Cuffs (254061, -1.95 DPS) [crafted]; Dreadmist Bracers (16703, -2.13 DPS) [dungeon]; Arcane Runed Bracers (4744, -2.19 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.34 DPS) | yes | Knight-Lieutenant's Dreadweave Handwraps (227100, +0.00 DPS) [pvp]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Dreamweave Gloves (10019, -1.87 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (343.4 DPS) | yes | Magister's Belt (16685, -0.51 DPS) [dungeon]; Dreadmist Belt (16702, -0.76 DPS) [dungeon]; Satyrmane Sash (17755, -3.77 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.85 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -5.15 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.5 spell_power points (9.30 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.21 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.7 spell_power points (7.83 DPS) | yes | Philanthropist's Ring (281635, -3.32 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.79 DPS) [quest]; Snake Hoop (6750, -4.84 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.93 DPS) | yes | Philanthropist's Ring (281635, -0.42 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.90 DPS) [quest]; Snake Hoop (6750, -1.94 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -4.76 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.70 DPS) [quest]; Teebu's Blazing Longsword (1728, -32.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 208.5 spell_power points (79.04 DPS) | yes | Brilliant Wand (249385, -2.04 DPS, sim-verified) [crafted]; Stormrager (16997, -15.63 DPS) [quest]; Torch of Austen (13004, -22.25 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Wisdom of the Timbermaw; legs: Ghostweave Pants; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 05000000000000000-2035113011220001351-0550005000000000)

Set DPS (verified): 675.1. Weights run: 2.1s. Verify run: 4.1s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.205 ± 0.039, crit=0.387 ± 0.013 per rating point (14 rating = 1%, 5.414 per %), hit=0.781 ± 0.050 per rating point (10 rating = 1%, 7.813 per %), spell_haste=4.520 ± 0.672, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.781 ± 0.002, fire_power=0.219 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (675.1 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -11.59 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.0 spell_power points (7.94 DPS) | yes | Glowing Eye of Mordresh (10769, -3.06 DPS) [dungeon]; Triune Amulet (7722, -3.42 DPS) [dungeon]; Mindburst Medallion (11196, -3.85 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.1 spell_power points (14.18 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.83 DPS) [pvp]; Argent Shoulders (19059, -4.12 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 spell_power points (8.61 DPS) | yes | Nightfall Drape (12465, -3.52 DPS) [dungeon]; Runecloth Cloak (13860, -3.98 DPS, sim-verified) [crafted]; Long Silken Cloak (4326, -4.64 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.8 spell_power points (27.05 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.18 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.60 DPS) [pvp]; Robe of the Magi (1716, -17.43 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (675.1 DPS) | yes | Arcane Runed Bracers (4744, +0.00 DPS) [quest]; Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Runecloth Cuffs (254123, -10.86 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (675.1 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.03 DPS) [pvp]; Earth Warder's Gloves (21318, -10.04 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (675.1 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -12.57 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (14.70 DPS) | yes | Knight's Dreadweave Leggings (220888, +0.00 DPS) [vendor]; Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.57 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.87 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -1.09 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.4 spell_power points (10.99 DPS) | yes | Philanthropist's Ring (281635, -4.75 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.46 DPS) [quest]; Snake Hoop (6750, -7.02 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.35 DPS) | yes | Philanthropist's Ring (281635, -1.12 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.83 DPS) [quest]; Snake Hoop (6750, -3.38 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (675.1 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (675.1 DPS) | yes | Mark of the Chosen (17774, -7.45 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (675.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.26 DPS) [world]; Teebu's Blazing Longsword (1728, -54.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 141.8 spell_power points (80.16 DPS) | yes | Brilliant Wand (249385, -14.68 DPS) [crafted]; Stormrager (16997, -16.89 DPS) [quest]; Greater Eternal Wand (249237, -21.74 DPS) [crafted] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Void; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 41.8. Weights run: 1.7s. Verify run: 1.1s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.998 ± 0.027, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.385 per %), hit=0.083 ± 0.004 per rating point (10 rating = 1%, 0.834 per %), spell_haste=not significant (-0.103 ± 0.064), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.724 ± 0.002, fire_power=0.276 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.25 DPS) | yes | Flame Circlet (253953, -0.62 DPS) [crafted]; Shadow Goggles (4373, -2.38 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.0 spell_power points (2.92 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.27 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -2.08 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (41.8 DPS) | yes | Sanguine Cape (14376, -0.00 DPS) [world_drop]; Black Whelp Cloak (7283, -0.21 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.74 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 10.0 spell_power points (2.09 DPS) | yes | Filigreed Flame Gown (253905, -0.52 DPS) [crafted]; Mystic's Wrap (14369, -0.63 DPS) [world_drop]; Gray Woolen Robe (2585, -1.59 DPS, sim-verified) [crafted] |
| wrist | Featherbead Bracers (15452) (or Mindthrust Bracers (1974)) | Returning the Lost Satchel [quest] | 5.0 spell_power points (1.04 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.21 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.42 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.46 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.00 DPS) [quest]; Gnoll Casting Gloves (892, -0.21 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 8.0 spell_power points (1.67 DPS) | yes | Novice Arcanist's Sash (253885, -0.21 DPS) [crafted]; Flame Sash (253929, -0.43 DPS) [crafted]; Keller's Girdle (2911, -1.69 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 12.0 spell_power points (2.50 DPS) | yes | Silk-threaded Trousers (1929, -1.04 DPS) [dungeon]; Darkweave Breeches (12987, -1.04 DPS) [world_drop]; Filigreed Flame Leggings (253941, -2.07 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.0 spell_power points (2.30 DPS) | yes | Pristine Boots (253889, -0.33 DPS, sim-verified) [crafted]; Flame Boots (253893, -1.32 DPS) [crafted]; Red Woolen Boots (4313, -1.46 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.0 spell_power points (1.25 DPS) | yes | Volcanic Rock Ring (12053, -0.63 DPS) [world_drop]; The 1 Ring (8350, -1.04 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 5.0 spell_power points (1.04 DPS) | yes | Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; The 1 Ring (8350, -0.83 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 10.0 spell_power points (2.08 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.83 DPS) [world]; Clear Crystal Rod (16894, -1.04 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 116.3 spell_power points (24.28 DPS) | yes | Cookie's Stirring Rod (5198, -0.47 DPS, sim-verified) [dungeon]; Skycaller (12984, -2.72 DPS) [world_drop]; Greater Magic Wand (11288, -6.36 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Featherbead Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 79.4. Weights run: 1.8s. Verify run: 1.0s. 289 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.033 ± 0.005, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.582 per %), hit=0.113 ± 0.005 per rating point (10 rating = 1%, 1.134 per %), spell_haste=not significant (0.286 ± 0.076), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.734 ± 0.002, fire_power=0.266 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.54 DPS) | yes | Filigreed Pristine Circlet (253975, -0.69 DPS) [crafted]; Enchanter's Cowl (4322, -1.08 DPS) [crafted]; Silk Headband (7050, -3.76 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.2 spell_power points (0.05 DPS) | yes | Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.02 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.3 spell_power points (2.15 DPS) | yes | Invoker's Mantle (215365, -0.49 DPS) [crafted]; Chestnut Mantle (17695, -0.56 DPS, sim-verified) [quest]; Death Speaker Mantle (6685, -0.68 DPS) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.15 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Black Whelp Cloak (7283, -0.46 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.00 DPS) | yes | Green Silk Armor (7065, -0.38 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.83 DPS) [dungeon]; Mechbuilder's Overalls (9508, -1.06 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.08 DPS) | yes | Phoenix Bindings (210781, -1.32 DPS, sim-verified) [crafted]; Tabitha's Cuffs (251486, -1.39 DPS) [quest]; Owlbeard Bracers (16981, -1.83 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.62 DPS) | yes | Jutebraid Gloves (10654, -0.19 DPS) [quest]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.44 DPS) [crafted] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.54 DPS) | yes | Belt of Arugal (6392, -0.44 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.69 DPS) [dungeon]; Invoker's Cord (215366, -0.89 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.77 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.15 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.34 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.2 spell_power points (1.67 DPS) | yes | Glinteye Slippers (273024, -0.05 DPS) [dungeon]; Acidic Walkers (9454, -0.45 DPS) [dungeon]; Spidersilk Boots (4320, -1.97 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.2 spell_power points (1.42 DPS) | yes | Solid Gold Band (279845, -0.50 DPS) [quest]; Electrocutioner Lagnut (9447, -0.73 DPS) [dungeon]; Runemender's Seal (273031, -0.73 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.39 DPS) | yes | Solid Gold Band (279845, -0.46 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.69 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.32 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -6.23 DPS) [quest]; Glimmering Staff (249392, -8.23 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 142.4 spell_power points (32.89 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -2.53 DPS) [dungeon]; Greater Mystic Wand (217287, -2.73 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 289, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 120.6. Weights run: 1.8s. Verify run: 1.2s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.208 ± 0.007, crit=0.060 ± 0.002 per rating point (14 rating = 1%, 0.839 per %), hit=0.176 ± 0.007 per rating point (10 rating = 1%, 1.761 per %), spell_haste=0.767 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.764 ± 0.002, fire_power=0.236 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.65 DPS) | yes | Living Cowl (5608, -2.15 DPS) [world]; Augural Shroud (2620, -2.24 DPS, sim-verified) [world]; Papal Fez (9431, -2.55 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.7 spell_power points (2.33 DPS) | yes | Triune Amulet (7722, -0.90 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.94 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.00 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.9 spell_power points (2.93 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Inquisitor's Shawl (19507, -0.32 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.0 spell_power points (1.90 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.01 DPS) [crafted]; Hillman's Cloak (3719, -0.55 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 spell_power points (6.26 DPS) | yes | Elemental Raiment (9434, -0.60 DPS) [world_drop]; Dreamweave Vest (10021, -0.91 DPS) [crafted]; Robe of Power (7054, -1.82 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.42 DPS) | yes | Condor Bracers (15864, -0.50 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.90 DPS) [quest]; Earthen Silk Cuffs (254019, -1.35 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.07 DPS) | yes | Black Mageweave Gloves (10003, -0.71 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.55 DPS) [crafted]; Gilded Handwraps (254021, -2.53 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.50 DPS) | yes | Deathmage Sash (10771, -0.78 DPS) [dungeon]; Warsong Sash (16975, -0.80 DPS, sim-verified) [quest]; Gilded Cord (254037, -0.90 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.5 spell_power points (4.44 DPS) | yes | Gaze Dreamer Pants (6903, -1.21 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.56 DPS) [crafted]; Stoneweaver Leggings (9407, -2.11 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.46 DPS) | yes | Gilded Slippers (254001, -2.77 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.35 DPS) [crafted]; Glinteye Slippers (273024, -4.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (2.97 DPS) | yes | Reedknot Ring (9622, -1.09 DPS) [quest]; Ring of Power Regulation (273030, -1.36 DPS) [dungeon]; Solid Gold Band (279845, -1.90 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.0 spell_power points (1.90 DPS) | yes | Reedknot Ring (9622, -0.01 DPS) [quest]; Ring of Power Regulation (273030, -0.28 DPS) [dungeon]; Solid Gold Band (279845, -0.82 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.11 DPS) [dungeon]; Spine of the Devourer (273040, -2.60 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.77 DPS) | yes | Omega Orb (7749, -0.79 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -1.55 DPS) [crafted]; Beacon of Hope (9393, -1.62 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | sim-verified (120.6 DPS) | yes | Umbral Wand (5216, -0.73 DPS) [dungeon]; Twisted Nether Wand (249144, -0.78 DPS) [crafted]; Jaina's Firestarter (13064, -1.55 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Earthen Rod

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 193.6. Weights run: 1.8s. Verify run: 1.4s. 498 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.266 ± 0.012, crit=0.091 ± 0.003 per rating point (14 rating = 1%, 1.268 per %), hit=0.368 ± 0.013 per rating point (10 rating = 1%, 3.680 per %), spell_haste=not significant (0.194 ± 0.161), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.700 ± 0.001, fire_power=0.300 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.04 DPS) | yes | Dreamweave Circlet (10041, -0.99 DPS) [crafted]; Red Mageweave Headband (10033, -1.56 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.79 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.1 spell_power points (2.72 DPS) | yes | Mindburst Medallion (11196, -0.46 DPS) [quest]; Triune Amulet (7722, -1.08 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.61 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.8 spell_power points (5.30 DPS) | yes | Black Mageweave Shoulders (10027, -1.08 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.83 DPS) [vendor]; Bloodmage Mantle (7684, -1.91 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (4.65 DPS) | yes | Deep Woodlands Cloak (19121, -0.74 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.33 DPS) [crafted]; Nightfall Drape (12465, -1.97 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.3 spell_power points (7.25 DPS) | yes | Robe of the Magi (1716, -0.22 DPS) [world_drop]; Elemental Raiment (9434, -0.99 DPS) [world_drop]; Dreamweave Vest (10021, -1.17 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.68 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -0.48 DPS) [quest]; Condor Bracers (15864, -0.60 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (5.68 DPS) | yes | Black Mageweave Gloves (10003, -1.21 DPS) [crafted]; Runecloth Gloves (13863, -1.39 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.55 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.7 spell_power points (4.96 DPS) | yes | Ghostweave Cord (254073, -0.79 DPS) [crafted]; Star Belt (4329, -1.09 DPS) [crafted]; Gilded Waistcord (254081, -1.27 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.66 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.76 DPS) [vendor]; Red Mageweave Pants (10009, -1.21 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -1.53 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.15 DPS) | yes | Gilded Sandals (254107, -3.16 DPS) [crafted]; Black Mageweave Boots (10026, -3.32 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.67 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.87 DPS) | yes | Snake Hoop (6750, -1.69 DPS) [quest]; Reedknot Ring (9622, -1.79 DPS) [quest]; Ring of Power Regulation (273030, -2.09 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.3 spell_power points (3.38 DPS) | yes | Snake Hoop (6750, -1.22 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.29 DPS) [quest]; Ring of Power Regulation (273030, -1.59 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (193.6 DPS) | yes | Rune of the Guard Captain (19120, -0.22 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (193.6 DPS) | yes | Rune of the Guard Captain (19120, -2.40 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (193.6 DPS) | yes | Spire of Hakkar (10844, -0.36 DPS) [world]; Spellshifter Rod (9527, -0.92 DPS) [quest]; Hanzo Sword (8190, -24.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 176.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.37 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 498, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 336.3. Weights run: 2.0s. Verify run: 2.4s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.379 ± 0.027, crit=0.131 ± 0.005 per rating point (14 rating = 1%, 1.837 per %), hit=0.554 ± 0.025 per rating point (10 rating = 1%, 5.538 per %), spell_haste=not significant (1.104 ± 0.324), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.746 ± 0.002, fire_power=0.254 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.23 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -1.39 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.8 spell_power points (5.98 DPS) | yes | Glowing Eye of Mordresh (10769, -2.18 DPS) [dungeon]; Mindburst Medallion (11196, -2.85 DPS) [quest]; Triune Amulet (7722, -2.95 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.48 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.58 DPS, sim-verified) [crafted]; Rotgrip Mantle (17732, -1.96 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.3 spell_power points (6.17 DPS) | yes | Deep Woodlands Cloak (19121, -0.33 DPS) [quest]; Runecloth Cloak (13860, -1.61 DPS) [crafted]; Nightfall Drape (12465, -2.76 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.4 spell_power points (18.73 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.15 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -6.38 DPS) [pvp]; Acumen Robes (17775, -10.20 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 14.8 spell_power points (5.60 DPS) | yes | Nethergeld Cuffs (254061, -1.95 DPS) [crafted]; Dreadmist Bracers (16703, -2.13 DPS) [dungeon]; Spidertank Oilrag (9448, -2.19 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.34 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Blood Guard's Dreadweave Handwraps (227099, +0.00 DPS) [pvp]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (336.3 DPS) | yes | Magister's Belt (16685, -0.51 DPS) [dungeon]; Dreadmist Belt (16702, -0.76 DPS) [dungeon]; Satyrmane Sash (17755, -4.32 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.85 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.73 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.5 spell_power points (9.30 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.21 DPS) [crafted]; Blood Guard's Dreadweave Walkers (227098, -0.61 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.7 spell_power points (7.83 DPS) | yes | Philanthropist's Ring (281635, -3.32 DPS) [quest]; Snake Hoop (6750, -4.84 DPS) [quest]; Reedknot Ring (9622, -5.17 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.93 DPS) | yes | Philanthropist's Ring (281635, -0.42 DPS) [quest]; Snake Hoop (6750, -1.94 DPS) [quest]; Reedknot Ring (9622, -2.27 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.70 DPS) [quest]; Teebu's Blazing Longsword (1728, -31.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 208.5 spell_power points (79.04 DPS) | yes | Brilliant Wand (249385, -15.04 DPS) [crafted]; Stormrager (16997, -15.63 DPS) [quest]; Torch of Austen (13004, -22.25 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Wisdom of the Timbermaw; legs: Ghostweave Pants; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 05000000000000000-2035113011220001351-0550005000000000)

Set DPS (verified): 660.0. Weights run: 2.1s. Verify run: 2.7s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.205 ± 0.039, crit=0.387 ± 0.013 per rating point (14 rating = 1%, 5.414 per %), hit=0.781 ± 0.050 per rating point (10 rating = 1%, 7.813 per %), spell_haste=4.520 ± 0.672, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.781 ± 0.002, fire_power=0.219 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (15.26 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -2.21 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.0 spell_power points (7.94 DPS) | yes | Glowing Eye of Mordresh (10769, -3.06 DPS) [dungeon]; Triune Amulet (7722, -3.42 DPS) [dungeon]; Mindburst Medallion (11196, -3.85 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.1 spell_power points (14.18 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.04 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -2.83 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 spell_power points (8.61 DPS) | yes | Deep Woodlands Cloak (19121, -0.78 DPS) [quest]; Runecloth Cloak (13860, -2.60 DPS) [crafted]; Nightfall Drape (12465, -3.52 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.8 spell_power points (27.05 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.18 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.60 DPS) [pvp]; Robe of the Magi (1716, -15.23 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.0 spell_power points (7.38 DPS) | yes | Spidertank Oilrag (9448, -2.29 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.61 DPS) [crafted]; Earthenweave Cuffs (254131, -2.85 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (660.0 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.03 DPS) [pvp]; Earth Warder's Gloves (21318, -8.69 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.0 spell_power points (9.07 DPS) | yes | Ghostweave Cord (254073, -1.16 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.23 DPS) [crafted]; Star Belt (4329, -1.72 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (14.70 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS) [vendor]; Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.57 DPS) | yes | Omnicast Boots (11822, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.13 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.4 spell_power points (10.99 DPS) | yes | Philanthropist's Ring (281635, -4.75 DPS) [quest]; Snake Hoop (6750, -7.02 DPS) [quest]; Reedknot Ring (9622, -7.03 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.35 DPS) | yes | Philanthropist's Ring (281635, -1.12 DPS) [quest]; Snake Hoop (6750, -3.38 DPS) [quest]; Reedknot Ring (9622, -3.39 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.30 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.88 DPS) [quest]; Mark of the Chosen (17774, -3.92 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.26 DPS) [world]; Teebu's Blazing Longsword (1728, -48.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 141.8 spell_power points (80.16 DPS) | yes | Brilliant Wand (249385, -14.68 DPS) [crafted]; Stormrager (16997, -16.89 DPS) [quest]; Greater Eternal Wand (249237, -21.74 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

