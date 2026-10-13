# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 30.2. Weights run: 3.7s. Verify run: 4.1s. 253 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.774 ± 0.007, crit=0.030 ± 0.002 per rating point (14 rating = 1%, 0.423 per %), hit=0.099 ± 0.006 per rating point (10 rating = 1%, 0.994 per %), spell_haste=-0.840 ± 0.044, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.001, arcane_power=0.967 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Wisdom's Leather Hood (252507, -0.40 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS) [crafted]; Trapper's Leather Hood (252505, -0.84 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 spell_power points (1.58 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -1.05 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.09 DPS, sim-verified) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.3 spell_power points (0.57 DPS) | yes | Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.17 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.53 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.07 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.39 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.9 spell_power points (0.51 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.10 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (1.07 DPS) | yes | Windfelt Gloves (5630, -0.23 DPS) [quest]; Pristine Gloves (253913, -0.23 DPS) [crafted]; Serpent Gloves (5970, -0.54 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.12 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.72 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.2 spell_power points (2.01 DPS) | yes | Dreamer's Leggings (270016, -0.35 DPS) [quest]; Wisdom's Leather Pants (252503, -0.47 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.65 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.1 spell_power points (1.34 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.29 DPS) [crafted]; Black Whelp Slippers (252424, -0.53 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (0.87 DPS) | yes | Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; The 1 Ring (8350, -0.76 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.66 DPS) | yes | Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; The 1 Ring (8350, -0.56 DPS) [world]; Lavishly Jeweled Ring (1156, -0.59 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.06 DPS) | yes | Lesser Staff of the Spire (1300, -0.44 DPS) [world]; Staff of Westfall (2042, -0.55 DPS) [quest]; Twisted Chanter's Staff (890, -1.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 253, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 3000 Brood Mother Carapace; 3018 Hide of Lupos

### Band 30 (night-elf, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 64.5. Weights run: 1.8s. Verify run: 2.3s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.549 ± 0.009, crit=0.037 ± 0.003 per rating point (14 rating = 1%, 0.512 per %), hit=0.125 ± 0.008 per rating point (10 rating = 1%, 1.254 per %), spell_haste=0.318 ± 0.055, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.335 ± 0.001, arcane_power=0.665 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.25 DPS) | yes | Holy Shroud (2721, -0.19 DPS) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted]; Enchanter's Cowl (4322, -1.48 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.93 DPS) | yes | Crystal Starfire Medallion (5003, -1.52 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.52 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.08 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.9 spell_power points (2.62 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS) [dungeon]; Fairywing Mantle (9536, -0.56 DPS) [quest]; Magician's Mantle (12998, -0.75 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hillman's Cloak (3719, -0.04 DPS) [crafted]; Darkspear Raider's Cloak (272078, -0.15 DPS) [vendor]; Vine Pruner's Cloak (279835, -1.23 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.1 spell_power points (3.03 DPS) | yes | Guardian Armor (4256, -0.31 DPS) [crafted]; Beguiler Robes (7728, -0.52 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.64 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.69 DPS) | yes | Nightsky Wristbands (6407, -1.07 DPS) [world_drop]; Technician's Bracers (270042, -1.07 DPS) [quest]; Glowing Magical Bracelets (13106, -2.05 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.0 spell_power points (1.89 DPS) | yes | Shilly Mitts (9609, -0.57 DPS) [quest]; Gloves of Insight (9698, -0.57 DPS) [quest]; Stormrider's Leather Gloves (252498, -1.04 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (2.50 DPS) | yes | Moss Cinch (6911, -0.24 DPS) [dungeon]; Belt of Arugal (6392, -0.50 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.93 DPS, sim-verified) [rep] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (2.63 DPS) | yes | Stormrider's Leather Pants (252502, -0.13 DPS) [crafted]; Guardian Pants (5962, -0.30 DPS) [crafted]; Abomination Skin Leggings (23173, -0.69 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.04 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -2.08 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.32 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.70 DPS) [dungeon]; Minor Channeling Ring (1449, -2.80 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.41 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; Minor Channeling Ring (1449, -1.28 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mechanic's Pipehammer (9604, -1.55 DPS) [quest]; Rhahk'Zor's Hammer (5187, -4.13 DPS) [dungeon]; Manual Crowd Pummeler (9449, -11.13 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2944 Cursed Eye of Paleth

### Band 40 (night-elf, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 83.0. Weights run: 1.7s. Verify run: 2.1s. 539 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.394 ± 0.007, crit=0.073 ± 0.005 per rating point (14 rating = 1%, 1.023 per %), hit=0.176 ± 0.011 per rating point (10 rating = 1%, 1.764 per %), spell_haste=not significant (-0.311 ± 0.103), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.28 DPS) | yes | Augural Shroud (2620, -1.24 DPS) [world]; Big Voodoo Mask (8201, -1.32 DPS) [crafted]; Living Cowl (5608, -1.63 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.2 spell_power points (2.07 DPS) | yes | Necklace of Calisea (1714, -1.51 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.51 DPS) [vendor]; Scorn's Icy Choker (23169, -1.58 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (2.56 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.0 spell_power points (1.63 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.40 DPS) [quest] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.11 DPS) [world_drop]; Robe of Power (7054, -0.57 DPS) [crafted]; Robe of the Magi (1716, -1.49 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Leather Bracers (4260, -0.13 DPS) [crafted]; Condor Bracers (15864, -0.41 DPS) [quest]; Arcane Runed Bracers (4744, -1.27 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Dreamweave Gloves (10019, -1.20 DPS, sim-verified) [crafted]; Prospector Gloves (4980, -1.36 DPS) [quest]; Black Mageweave Gloves (10003, -1.84 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.6 spell_power points (3.18 DPS) | yes | Star Belt (4329, -0.53 DPS) [crafted]; Deathmage Sash (10771, -0.54 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.65 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.7 spell_power points (3.82 DPS) | yes | Crimson Silk Pantaloons (7062, -1.14 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.30 DPS) [crafted]; Dark Ritual Leggings (270031, -2.43 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Skycaller's Leather Shoes (252532, -1.50 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.70 DPS) [crafted]; Gilded Slippers (254001, -2.90 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (2.44 DPS) | yes | Ring of Forlorn Spirits (2043, -0.81 DPS) [quest]; Reedknot Ring (9622, -1.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.22 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.84 DPS) | yes | Reedknot Ring (9622, -0.41 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.61 DPS) [vendor]; Ring of Forlorn Spirits (2043, -1.43 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.16 DPS) [dungeon]; Hand of Righteousness (7721, -3.29 DPS) [dungeon] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.85 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Lorica (11262, -0.83 DPS) [quest]; Orb of Mystic Insight (249394, -0.95 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 539, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2622 Nimar's Tribal Headdress

### Band 50 (night-elf, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 136.6. Weights run: 1.8s. Verify run: 3.0s. 677 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.467 ± 0.012, crit=0.114 ± 0.009 per rating point (14 rating = 1%, 1.595 per %), hit=0.277 ± 0.019 per rating point (10 rating = 1%, 2.765 per %), spell_haste=not significant (0.166 ± 0.189), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.379 ± 0.001, arcane_power=0.621 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.0 spell_power points (5.66 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -0.82 DPS) [crafted]; Red Mageweave Headband (10033, -1.36 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.7 spell_power points (2.03 DPS) | yes | Mindburst Medallion (11196, -0.36 DPS) [quest]; Horizon Choker (13085, -0.79 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.75 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.3 spell_power points (4.59 DPS) | yes | Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.93 DPS) [vendor]; Skycaller's Leather Shoulder (252537, -1.35 DPS) [crafted]; Rotgrip Mantle (17732, -2.76 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (3.17 DPS) | yes | Big Voodoo Cloak (8216, -1.43 DPS) [crafted]; Nightfall Drape (12465, -1.47 DPS) [dungeon]; Runecloth Cloak (13860, -1.96 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.3 spell_power points (5.35 DPS) | yes | Feathered Breastplate (8349, -0.88 DPS) [crafted]; Dreamweave Vest (10021, -1.16 DPS) [crafted]; Robe of the Magi (1716, -2.80 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.7 spell_power points (2.77 DPS) | yes | Skycaller's Leather Bracers (252542, -0.26 DPS) [crafted]; Nethergeld Cuffs (254061, -0.83 DPS) [crafted]; Mender's Leather Bracers (252543, -1.02 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 25.1 spell_power points (4.73 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.92 DPS) [crafted]; Dreamweave Gloves (10019, -0.98 DPS) [crafted]; Gloves of the Greatfather (17721, -2.35 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.6 spell_power points (4.08 DPS) | yes | Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.08 DPS) [rep]; Mender's Leather Waistguard (252477, -1.13 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 21.0 spell_power points (3.97 DPS) | yes | Red Mageweave Pants (10009, -0.27 DPS) [crafted]; Wizardweave Leggings (14132, -0.38 DPS) [crafted]; Big Voodoo Pants (8202, -1.48 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.53 DPS) | yes | Skycaller's Leather Boots (252471, -0.73 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -1.05 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.46 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.45 DPS) | yes | Lorekeeper's Ring (19523, -0.19 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.94 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.3 spell_power points (2.33 DPS) | yes | Lorekeeper's Ring (19523, +0.00 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.82 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (136.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (136.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (136.6 DPS) | yes | Spellshifter Rod (9527, -0.16 DPS) [quest]; Spire of Hakkar (10844, -0.68 DPS) [world]; Hammer of the Northern Wind (810, -24.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Knight's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave

No-known-source sample (15 of 677, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2622 Nimar's Tribal Headdress

### Band 60 (night-elf, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 242.4. Weights run: 1.8s. Verify run: 12.2s. 1319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.662 ± 0.020, crit=0.145 ± 0.011 per rating point (14 rating = 1%, 2.032 per %), hit=0.371 ± 0.027 per rating point (10 rating = 1%, 3.709 per %), spell_haste=not significant (0.188 ± 0.372), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.377 ± 0.001, arcane_power=0.623 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cowl (226773) | Saving the Best for Last [quest] | 39.9 spell_power points (8.01 DPS) | yes | Field Marshal's Dragonhide Helm (231695, -0.34 DPS) [vendor]; Living Crown (252561, -0.61 DPS) [crafted]; Lieutenant Commander's Dragonhide Helm (227192, -1.00 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.6 spell_power points (4.74 DPS) | yes | Orb of the Darkmoon (19426, -0.32 DPS) [quest]; Beads of Ogre Mojo (22149, -0.53 DPS) [quest]; Chains of the Lich (23125, -2.65 DPS, sim-verified) [dungeon] |
| shoulder | Feralheart Spaulders (226778) | Anthion's Parting Words [quest] | sim-verified (242.4 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.16 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -5.11 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.0 spell_power points (5.02 DPS) | yes | Hide of the Wild (18510, -0.88 DPS) [crafted]; Spritecaster Cape (11623, -1.41 DPS) [dungeon]; Runecloth Cloak (13860, -2.15 DPS) [crafted] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | 37.9 spell_power points (7.61 DPS) | yes | Field Marshal's Dragonhide Armor (231696, -0.19 DPS) [vendor]; Acumen Robes (17775, -1.14 DPS) [quest]; Knight-Captain's Dragonhide Armor (227195, -6.24 DPS, sim-verified) [vendor] |
| wrist | Feralheart Wraps (226775) | Mokvar [vendor] | sim-verified (242.4 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 27.6 spell_power points (5.54 DPS) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Gloves of Undead Cleansing (23084, -0.60 DPS) [quest]; Raider Handwraps (272097, -8.48 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.3 spell_power points (8.68 DPS) | yes | Elunite Cord (272401, -2.28 DPS) [vendor]; Girdle of Insight (18504, -2.62 DPS) [crafted]; Belt of the Archmage (18405, -4.61 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.9 spell_power points (10.42 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.86 DPS) [vendor]; Sentinel's Silk Leggings (22752, -2.28 DPS) [rep] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 32.6 spell_power points (6.54 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.47 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -1.00 DPS) [vendor]; Omnicast Boots (11822, -12.53 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 37.7 spell_power points (7.56 DPS) | yes | Deep Rooted Ring (19109, -3.82 DPS) [world]; Lorekeeper's Ring (19522, -4.75 DPS) [rep]; Blessed Band of Light (272407, -4.81 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 29.3 spell_power points (5.87 DPS) | yes | Lorekeeper's Ring (19522, -3.07 DPS) [rep]; Blessed Band of Light (272407, -3.12 DPS) [vendor]; Deep Rooted Ring (19109, -3.61 DPS, sim-verified) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (242.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (242.4 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -1.61 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.14 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (242.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.26 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.44 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Cowl; neck: Amulet of the Dawn; shoulder: Feralheart Spaulders; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Feralheart Wraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spellshifter Rod

No-known-source sample (15 of 1319, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2622 Nimar's Tribal Headdress

### Band 60, raid preset (night-elf, 4132220115501051-05000000000000000000-5053000000000000)

Set DPS (verified): 533.4. Weights run: 1.6s. Verify run: 10.2s. 1319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.238 ± 0.015, crit=0.284 ± 0.016 per rating point (14 rating = 1%, 3.969 per %), hit=0.654 ± 0.032 per rating point (10 rating = 1%, 6.540 per %), spell_haste=3.731 ± 0.148, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.194 ± 0.000, arcane_power=0.806 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (533.4 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.01 DPS) [vendor]; Feralheart Cowl (226773, -7.98 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.25 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.82 DPS) [quest]; Beads of Ogre Mojo (22149, -2.86 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.5 spell_power points (16.09 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, -1.97 DPS) [pvp]; Feralheart Spaulders (226778, -2.86 DPS) [quest]; Argent Shoulders (19059, -4.45 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.4 spell_power points (11.39 DPS) | yes | Hide of the Wild (18510, -3.76 DPS) [crafted]; Spritecaster Cape (11623, -4.20 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -4.87 DPS) [rep] |
| chest | Knight-Captain's Dragonhide Armor (227195) | Captain Dirgehammer [vendor] | sim-verified (533.4 DPS) | yes | Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -0.44 DPS) [vendor]; Feralheart Vest (226776, -12.47 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.9 spell_power points (11.14 DPS) | yes | Runecloth Cuffs (254123, -4.90 DPS) [crafted]; Runic Leather Bracers (15092, -5.37 DPS) [crafted]; Skycaller's Leather Bracers (252542, -5.70 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (11.18 DPS) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -0.93 DPS) [quest]; Feralheart Hands (226777, -5.93 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 36.8 spell_power points (17.14 DPS) | yes | Belt of the Archmage (18405, -4.08 DPS, sim-verified) [crafted]; Elunite Cord (272401, -7.02 DPS) [vendor]; Girdle of Insight (18504, -7.59 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 46.3 spell_power points (21.55 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.83 DPS) [vendor]; Sentinel's Silk Leggings (237815, -1.83 DPS) [vendor]; Marshal's Dragonhide Pants (231697, -5.57 DPS) [vendor] |
| feet | Knight-Lieutenant's Dragonhide Boots (227194) | Captain Dirgehammer [vendor] | sim-verified (533.4 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, +0.00 DPS) [crafted]; Feralheart Galoshes (226774, -10.69 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 36.7 spell_power points (17.09 DPS) | yes | Blessed Band of Light (272407, -9.39 DPS) [vendor]; Deep Rooted Ring (19109, -10.39 DPS) [world]; Lorekeeper's Ring (19522, -10.57 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 24.6 spell_power points (11.47 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Deep Rooted Ring (19109, -4.77 DPS) [world]; Lorekeeper's Ring (19522, -4.95 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (533.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (533.4 DPS) | yes | Weakness Analyzer (272438, -3.07 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.73 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.78 DPS) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (533.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.40 DPS) [world]; Hand of Edward the Odd (2243, -40.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Knight-Captain's Dragonhide Armor; wrist: Dryad's Wrist Bindings; hands: Gloves of the Greatfather; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Knight-Lieutenant's Dragonhide Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field

No-known-source sample (15 of 1319, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2622 Nimar's Tribal Headdress

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 29.0. Weights run: 3.7s. Verify run: 2.5s. 243 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.774 ± 0.007, crit=0.030 ± 0.002 per rating point (14 rating = 1%, 0.423 per %), hit=0.099 ± 0.006 per rating point (10 rating = 1%, 0.994 per %), spell_haste=-0.840 ± 0.044, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.001, arcane_power=0.967 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Trapper's Leather Hood (252505, -0.38 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.40 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 spell_power points (1.58 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.04 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.05 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.3 spell_power points (0.57 DPS) | yes | Heavy Woolen Cloak (4311, -0.04 DPS) [crafted]; Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.17 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.07 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.39 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.9 spell_power points (0.51 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.10 DPS) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (1.07 DPS) | yes | Serpent Gloves (5970, -0.14 DPS) [dungeon]; Pristine Gloves (253913, -0.23 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.26 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.12 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.71 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.2 spell_power points (2.01 DPS) | yes | Stormrider's Leather Pants (252502, -0.07 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.47 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.60 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Boots (252444, -0.26 DPS) [crafted]; Spidersilk Boots (4320, -0.31 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.50 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.15 DPS) [quest]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; The 1 Ring (8350, -0.56 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.6 spell_power points (0.61 DPS) | yes | Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; Loop of Sacrifice (281673, -0.33 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.51 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.06 DPS) | yes | Twisted Chanter's Staff (890, -0.03 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.44 DPS) [world]; Gnarled Necromancer's Staff (251534, -0.83 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 243, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 3000 Brood Mother Carapace

### Band 30 (tauren, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 52.6. Weights run: 1.8s. Verify run: 2.3s. 416 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.549 ± 0.009, crit=0.037 ± 0.003 per rating point (14 rating = 1%, 0.512 per %), hit=0.125 ± 0.008 per rating point (10 rating = 1%, 1.254 per %), spell_haste=0.318 ± 0.055, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.335 ± 0.001, arcane_power=0.665 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.09 DPS) [world_drop]; Silk Headband (7050, -0.47 DPS) [crafted]; Totemic Leather Helm (252456, -0.60 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.93 DPS) | yes | Crystal Starfire Medallion (5003, -1.52 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.52 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.79 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.9 spell_power points (2.62 DPS) | yes | Mantle of Woe (7750, -0.44 DPS) [quest]; Fairywing Mantle (9536, -0.56 DPS) [quest]; Death Speaker Mantle (6685, -1.27 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.94 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.11 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.1 spell_power points (3.03 DPS) | yes | Guardian Armor (4256, -0.31 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.46 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.52 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.69 DPS) | yes | Nightsky Wristbands (6407, -1.07 DPS) [world_drop]; Technician's Bracers (270042, -1.07 DPS) [quest]; Glowing Magical Bracelets (13106, -1.36 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.1 spell_power points (1.90 DPS) | yes | Stormrider's Leather Gloves (252498, -0.55 DPS) [crafted]; Serpent Gloves (5970, -0.58 DPS) [dungeon]; Jutebraid Gloves (10654, -0.94 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (2.50 DPS) | yes | Defiler's Cloth Girdle (20164, -0.12 DPS) [rep]; Moss Cinch (6911, -0.24 DPS) [dungeon]; Warsong Sash (16975, -0.43 DPS) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormrider's Leather Pants (252502, -0.02 DPS) [crafted]; Guardian Pants (5962, -0.19 DPS) [crafted]; Dark Ritual Leggings (270031, -0.58 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.04 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -2.73 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.32 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.70 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.75 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.13 DPS) | yes | Snake Hoop (6750, -0.41 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.56 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, -0.37 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -0.47 DPS) [quest]; Manual Crowd Pummeler (9449, -1.68 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

### Band 40 (tauren, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 82.7. Weights run: 1.7s. Verify run: 2.3s. 525 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.394 ± 0.007, crit=0.073 ± 0.005 per rating point (14 rating = 1%, 1.023 per %), hit=0.176 ± 0.011 per rating point (10 rating = 1%, 1.764 per %), spell_haste=not significant (-0.311 ± 0.103), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.28 DPS) | yes | Augural Shroud (2620, -1.24 DPS) [world]; Big Voodoo Mask (8201, -1.32 DPS) [crafted]; Living Cowl (5608, -1.63 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.2 spell_power points (2.07 DPS) | yes | Necklace of Calisea (1714, -1.51 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.51 DPS) [vendor]; Scorn's Icy Choker (23169, -1.51 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (2.56 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.0 spell_power points (1.63 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Battle Healer's Cloak (19528, -0.40 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (82.7 DPS) | yes | Elemental Raiment (9434, -0.11 DPS) [world_drop]; Zealot's Robe (17043, -0.53 DPS) [quest]; Robe of the Magi (1716, -1.70 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.84 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.38 DPS) [quest]; Guardian Leather Bracers (4260, -0.75 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Dreamweave Gloves (10019, -1.03 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.84 DPS) [crafted]; Red Mageweave Gloves (10018, -1.85 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.6 spell_power points (3.18 DPS) | yes | Star Belt (4329, -0.53 DPS) [crafted]; Deathmage Sash (10771, -0.54 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.65 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.7 spell_power points (3.82 DPS) | yes | Crimson Silk Pantaloons (7062, -1.14 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.30 DPS) [crafted]; Dark Ritual Leggings (270031, -2.25 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Skycaller's Leather Shoes (252532, -1.62 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.70 DPS) [crafted]; Gilded Slippers (254001, -2.90 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (2.44 DPS) | yes | Reedknot Ring (9622, -1.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.22 DPS) [vendor]; Electrocutioner Lagnut (9447, -1.83 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.84 DPS) | yes | Sea Giant's Toe Ring (274746, -0.61 DPS) [vendor]; Electrocutioner Lagnut (9447, -1.22 DPS) [dungeon]; Reedknot Ring (9622, -1.74 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.16 DPS) [dungeon]; Skullbreaker (17039, -0.49 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.85 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Mystic Insight (249394, -0.95 DPS) [crafted]; Omega Orb (7749, -1.02 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

### Band 50 (tauren, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 135.0. Weights run: 1.8s. Verify run: 2.6s. 659 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.467 ± 0.012, crit=0.114 ± 0.009 per rating point (14 rating = 1%, 1.595 per %), hit=0.277 ± 0.019 per rating point (10 rating = 1%, 2.765 per %), spell_haste=not significant (0.166 ± 0.189), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.379 ± 0.001, arcane_power=0.621 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.0 spell_power points (5.66 DPS) | yes | Red Mageweave Headband (10033, -0.31 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -0.82 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.7 spell_power points (2.03 DPS) | yes | Mindburst Medallion (11196, -0.36 DPS) [quest]; Horizon Choker (13085, -0.79 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.49 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.3 spell_power points (4.59 DPS) | yes | Blood Guard's Crackling Leather Spaulders (220871, -0.93 DPS) [vendor]; Skycaller's Leather Shoulder (252537, -1.35 DPS) [crafted]; Rotgrip Mantle (17732, -2.61 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (3.17 DPS) | yes | Deep Woodlands Cloak (19121, -0.11 DPS) [quest]; Runecloth Cloak (13860, -0.77 DPS) [crafted]; Big Voodoo Cloak (8216, -1.43 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.3 spell_power points (5.35 DPS) | yes | Feathered Breastplate (8349, -0.88 DPS) [crafted]; Dreamweave Vest (10021, -1.16 DPS) [crafted]; Robe of the Magi (1716, -1.97 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.7 spell_power points (2.77 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.26 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 25.1 spell_power points (4.73 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.92 DPS) [crafted]; Dreamweave Gloves (10019, -0.98 DPS) [crafted]; Gloves of the Greatfather (17721, -2.17 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.6 spell_power points (4.08 DPS) | yes | Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.08 DPS) [rep]; Mender's Leather Waistguard (252477, -1.13 DPS) [crafted] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | Lady Palanseer [vendor] | 21.0 spell_power points (3.97 DPS) | yes | Red Mageweave Pants (10009, -0.27 DPS) [crafted]; Wizardweave Leggings (14132, -0.38 DPS) [crafted]; Big Voodoo Pants (8202, -1.61 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.53 DPS) | yes | Skycaller's Leather Boots (252471, -0.73 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -1.05 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.46 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.45 DPS) | yes | Advisor's Ring (19519, -0.19 DPS) [rep]; Reedknot Ring (9622, -1.13 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.3 spell_power points (2.33 DPS) | yes | Advisor's Ring (19519, -0.06 DPS) [rep]; Reedknot Ring (9622, -1.01 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (135.0 DPS) | yes | Rune of the Guard Captain (19120, -0.10 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (135.0 DPS) | yes | Rune of the Guard Captain (19120, -1.84 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (135.0 DPS) | yes | Spellshifter Rod (9527, -0.16 DPS) [quest]; Spire of Hakkar (10844, -0.68 DPS) [world]; Hammer of the Northern Wind (810, -23.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave

No-known-source sample (15 of 659, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

### Band 60 (tauren, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 238.9. Weights run: 1.8s. Verify run: 8.1s. 1315 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.662 ± 0.020, crit=0.145 ± 0.011 per rating point (14 rating = 1%, 2.032 per %), hit=0.371 ± 0.027 per rating point (10 rating = 1%, 3.709 per %), spell_haste=not significant (0.188 ± 0.372), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.377 ± 0.001, arcane_power=0.623 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cowl (226773) | Saving the Best for Last [quest] | 39.9 spell_power points (8.01 DPS) | yes | Warlord's Dragonhide Helm (231678, -0.34 DPS) [vendor]; Champion's Dragonhide Helm (227186, -1.00 DPS) [pvp]; Living Crown (252561, -3.41 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.6 spell_power points (4.74 DPS) | yes | Orb of the Darkmoon (19426, -0.32 DPS) [quest]; Beads of Ogre Mojo (22149, -0.53 DPS) [quest]; Chains of the Lich (23125, -3.06 DPS, sim-verified) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 39.0 spell_power points (7.82 DPS) | yes | Warlord's Dragonhide Spaulders (231681, -0.94 DPS) [vendor]; Darkspear Shoulderpads (272103, -1.82 DPS) [vendor]; Feralheart Spaulders (226778, -3.54 DPS, sim-verified) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.0 spell_power points (5.02 DPS) | yes | Hide of the Wild (18510, -0.88 DPS) [crafted]; Spritecaster Cape (11623, -1.41 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.41 DPS) [quest] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | 37.9 spell_power points (7.61 DPS) | yes | Warlord's Dragonhide Armor (231679, -0.19 DPS) [vendor]; Legionnaire's Dragonhide Armor (227189, -0.73 DPS) [vendor]; Acumen Robes (17775, -11.15 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.3 spell_power points (5.48 DPS) | yes | Runecloth Cuffs (254123, -1.94 DPS) [crafted]; Runic Leather Bracers (15092, -2.14 DPS) [crafted]; Feralheart Wraps (226775, -2.41 DPS) [vendor] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 27.6 spell_power points (5.54 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Gloves of Undead Cleansing (23084, -0.60 DPS) [quest]; Raider Handwraps (272097, -8.53 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.3 spell_power points (8.68 DPS) | yes | Elunite Cord (272401, -2.28 DPS) [vendor]; Girdle of Insight (18504, -2.62 DPS) [crafted]; Belt of the Archmage (18405, -5.21 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.9 spell_power points (10.42 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.86 DPS) [vendor]; Outrider's Silk Leggings (22747, -2.28 DPS) [rep]; Sentinel's Silk Leggings (237815, -3.80 DPS, sim-verified) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 32.6 spell_power points (6.54 DPS) | yes | General's Dragonhide Boots (231682, -0.47 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -1.00 DPS) [pvp]; Omnicast Boots (11822, -14.21 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 37.7 spell_power points (7.56 DPS) | yes | Deep Rooted Ring (19109, -3.82 DPS) [world]; Advisor's Ring (19518, -4.75 DPS) [rep]; Blessed Band of Light (272407, -4.81 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 29.3 spell_power points (5.87 DPS) | yes | Advisor's Ring (19518, -3.07 DPS) [rep]; Blessed Band of Light (272407, -3.12 DPS) [vendor]; Deep Rooted Ring (19109, -3.89 DPS, sim-verified) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (238.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (238.9 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -1.61 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.14 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (238.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.26 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.44 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spellshifter Rod

No-known-source sample (15 of 1315, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

### Band 60, raid preset (tauren, 4132220115501051-05000000000000000000-5053000000000000)

Set DPS (verified): 527.1. Weights run: 1.6s. Verify run: 9.2s. 1315 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.238 ± 0.015, crit=0.284 ± 0.016 per rating point (14 rating = 1%, 3.969 per %), hit=0.654 ± 0.032 per rating point (10 rating = 1%, 6.540 per %), spell_haste=3.731 ± 0.148, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.194 ± 0.000, arcane_power=0.806 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+8.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.01 DPS) [pvp]; Feralheart Cowl (226773, -8.11 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.25 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.82 DPS) [quest]; Beads of Ogre Mojo (22149, -2.86 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.5 spell_power points (16.09 DPS) | yes | Warlord's Dragonhide Spaulders (231681, -1.97 DPS) [vendor]; Feralheart Spaulders (226778, -2.86 DPS) [quest]; Argent Shoulders (19059, -4.45 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.4 spell_power points (11.39 DPS) | yes | Hide of the Wild (18510, -3.76 DPS) [crafted]; Spritecaster Cape (11623, -4.20 DPS) [dungeon]; Deep Woodlands Cloak (19121, -4.80 DPS) [quest] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | 30.3 spell_power points (14.11 DPS) | yes | Fernpulse Jerkin (15786, +0.00 DPS) [quest]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Legionnaire's Dragonhide Armor (227189, -0.59 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.9 spell_power points (11.14 DPS) | yes | Runecloth Cuffs (254123, -4.90 DPS) [crafted]; Runic Leather Bracers (15092, -5.37 DPS) [crafted]; Skycaller's Leather Bracers (252542, -5.70 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (11.18 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -0.93 DPS) [quest]; Feralheart Hands (226777, -4.63 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 36.8 spell_power points (17.14 DPS) | yes | Belt of the Archmage (18405, -5.10 DPS, sim-verified) [crafted]; Elunite Cord (272401, -7.02 DPS) [vendor]; Girdle of Insight (18504, -7.59 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 46.3 spell_power points (21.55 DPS) | yes | Sentinel's Silk Leggings (237815, -1.83 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.83 DPS) [vendor]; General's Dragonhide Pants (231680, -5.57 DPS) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 25.8 spell_power points (12.03 DPS) | yes | General's Dragonhide Boots (231682, -0.69 DPS) [pvp]; Earthen Silk Slippers (254013, -0.85 DPS) [crafted]; Blood Guard's Dragonhide Boots (227188, -1.13 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 36.7 spell_power points (17.09 DPS) | yes | Elemental Focus Band (20682, -5.62 DPS) [world]; Deep Rooted Ring (19109, -10.39 DPS) [world]; Advisor's Ring (19518, -10.57 DPS) [rep] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Deep Rooted Ring (19109, -1.00 DPS) [world]; Advisor's Ring (19518, -1.18 DPS) [rep]; Elemental Focus Band (20682, -6.13 DPS, sim-verified) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+24.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.07 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.73 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.78 DPS) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.40 DPS) [world]; Hand of Edward the Odd (2243, -40.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; hands: Gloves of the Greatfather; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket1: Talisman of Ascendance; trinket2: Serenity Field

No-known-source sample (15 of 1315, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

