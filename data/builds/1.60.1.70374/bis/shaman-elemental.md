# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 25.0. Weights run: 1.5s. Verify run: 2.1s. 217 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.656 ± 0.012, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.003 per %), hit=0.215 ± 0.009 per rating point (10 rating = 1%, 2.147 per %), spell_haste=-1.707 ± 0.106, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.32 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (0.91 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.58 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.00 DPS) [crafted]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Caretaker's Cape (20428, -0.08 DPS) [rep] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.03 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.05 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.05 DPS) [vendor]; Owl Bracers (4796, -1.41 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.6 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Windfelt Gloves (5630, -0.14 DPS) [quest]; Serpent Gloves (5970, -0.35 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.44 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.19 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -0.25 DPS) [quest]; Wisdom's Leather Pants (252503, -0.28 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (0.81 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.30 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.3 spell_power points (0.53 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; The 1 Ring (8350, -0.47 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.42 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; The 1 Ring (8350, -0.36 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Twisted Chanter's Staff (890, -0.12 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.34 DPS) [world]; Staff of Westfall (2042, -0.40 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (dwarf, 5530311300000000-000000000000000000-0000000000000000)

Set DPS (verified): 48.1. Weights run: 1.5s. Verify run: 2.1s. 383 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.635 ± 0.014, crit=0.088 ± 0.003 per rating point (14 rating = 1%, 1.235 per %), hit=0.259 ± 0.011 per rating point (10 rating = 1%, 2.590 per %), spell_haste=-2.472 ± 0.191, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.944 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.28 DPS) | yes | Holy Shroud (2721, -0.14 DPS) [world_drop]; Enduring Cap (3020, -0.23 DPS) [world_drop]; Totemic Leather Helm (252456, -0.68 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.12 DPS) | yes | Darkspear Warding Pendant (272075, -0.79 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.12 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (1.53 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Fairywing Mantle (9536, -0.31 DPS) [quest]; Magician's Mantle (12998, -0.41 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 5.1 spell_power points (0.53 DPS) | yes | Hillman's Cloak (3719, -0.01 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.47 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (1.79 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Guardian Armor (4256, -0.20 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.93 DPS) | yes | Nightsky Wristbands (6407, -0.54 DPS) [world_drop]; Technician's Bracers (270042, -0.54 DPS) [quest]; Glowing Magical Bracelets (13106, -0.95 DPS, sim-verified) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.5 spell_power points (0.78 DPS) | yes | Town Clerk's Mittens (270029, -0.01 DPS) [quest]; Shilly Mitts (9609, -0.06 DPS) [quest]; Gloves of Insight (9698, -0.06 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.8 spell_power points (1.43 DPS) | yes | Moss Cinch (6911, -0.19 DPS) [dungeon]; Belt of Arugal (6392, -0.30 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.38 DPS, sim-verified) [rep] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (1.46 DPS) | yes | Stormrider's Leather Pants (252502, -0.03 DPS) [crafted]; Guardian Pants (5962, -0.10 DPS) [crafted]; Dark Ritual Leggings (270031, -0.58 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.4 spell_power points (1.19 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.24 DPS) [crafted]; Acidic Walkers (9454, -1.32 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (0.95 DPS) | yes | Minor Channeling Ring (1449, -0.30 DPS) [quest]; Ring of Power Regulation (273030, -0.33 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.73 DPS) | yes | Ring of Power Regulation (273030, -0.10 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Minor Channeling Ring (1449, -0.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (48.1 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.66 DPS) [quest]; Manual Crowd Pummeler (9449, -7.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Spine of the Devourer

No-known-source sample (15 of 383, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots

### Band 40 (dwarf, 5530311300103050-010000000000000000-0000000000000000)

Set DPS (verified): 69.3. Weights run: 1.5s. Verify run: 2.7s. 597 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.893 ± 0.034, crit=0.249 ± 0.010 per rating point (14 rating = 1%, 3.492 per %), hit=0.431 ± 0.030 per rating point (10 rating = 1%, 4.309 per %), spell_haste=-2.478 ± 0.395, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.944 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.2 spell_power points (2.35 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.22 DPS) [crafted]; Augural Shroud (2620, -0.33 DPS) [world]; Big Voodoo Mask (8201, -1.79 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.1 spell_power points (1.44 DPS) | yes | Triune Amulet (7722, -0.62 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Scorn's Icy Choker (23169, -1.18 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.6 spell_power points (1.89 DPS) | yes | Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -1.13 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.5 spell_power points (1.06 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.06 DPS) [vendor]; Blackforge Cape (6424, -0.34 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.4 spell_power points (2.78 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of Power (7054, -0.27 DPS) [crafted]; Big Voodoo Robe (8200, -0.59 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.4 spell_power points (1.15 DPS) | yes | Turtle Scale Bracers (8198, -0.09 DPS) [crafted]; Arcane Runed Bracers (4744, -0.24 DPS) [quest]; Spidertank Oilrag (9448, -0.24 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.44 DPS) | yes | Red Mageweave Gloves (10018, -0.41 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.50 DPS) [crafted]; Dreamweave Gloves (10019, -1.33 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.4 spell_power points (2.07 DPS) | yes | Skycaller's Leather Belt (252522, -0.51 DPS) [crafted]; Gilded Cord (254037, -0.53 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.95 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.7 spell_power points (2.51 DPS) | yes | Crimson Silk Pantaloons (7062, -0.73 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.87 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.95 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.44 DPS) | yes | Skycaller's Leather Shoes (252532, -0.48 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.48 DPS) [crafted]; Mender's Leather Shoes (252533, -0.99 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.47 DPS) | yes | Ring of Forlorn Spirits (2043, -0.66 DPS) [quest]; Reedknot Ring (9622, -0.76 DPS) [quest]; Snake Hoop (6750, -2.36 DPS, sim-verified) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-verified (69.3 DPS) | yes | Ring of Forlorn Spirits (2043, -0.10 DPS) [quest]; Reedknot Ring (9622, -0.20 DPS) [quest]; Snake Hoop (6750, -0.69 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.18 DPS) [dungeon]; Spine of the Devourer (273040, -1.54 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 14.9 spell_power points (1.52 DPS) | yes | Thrash's Trash (276204, -0.09 DPS) [vendor]; Orb of Mystic Insight (249394, -0.26 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.22 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of Lorica

No-known-source sample (15 of 597, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 50 (dwarf, 5530311300103050-030000000000000000-5300000000000000)

Set DPS (verified): 96.4. Weights run: 1.5s. Verify run: 3.2s. 765 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.825 ± 0.041, crit=0.297 ± 0.012 per rating point (14 rating = 1%, 4.152 per %), hit=0.578 ± 0.040 per rating point (10 rating = 1%, 5.776 per %), spell_haste=-2.453 ± 0.478, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.936 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.5 spell_power points (3.38 DPS) | yes | Soothsayer's Headdress (17740, -0.01 DPS) [dungeon]; Dreamweave Circlet (10041, -0.60 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.81 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 spell_power points (1.30 DPS) | yes | Horizon Choker (13085, -0.20 DPS) [world_drop]; Mindburst Medallion (11196, -0.25 DPS) [quest]; Scorn's Icy Choker (23169, -0.91 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 31.5 spell_power points (3.00 DPS) | yes | Rockshard Pauldrons (9411, -0.84 DPS) [dungeon]; Rotgrip Mantle (17732, -0.97 DPS, sim-verified) [dungeon]; Skycaller's Leather Shoulder (252537, -1.06 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.9 spell_power points (1.81 DPS) | yes | Runecloth Cloak (13860, -0.32 DPS) [crafted]; Big Voodoo Cloak (8216, -0.62 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.71 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.5 spell_power points (3.38 DPS) | yes | Feathered Breastplate (8349, -0.79 DPS) [crafted]; Robe of the Magi (1716, -0.81 DPS) [world_drop]; Runecloth Tunic (13857, -0.90 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.2 spell_power points (1.74 DPS) | yes | Skycaller's Leather Bracers (252542, -0.24 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.24 DPS) [crafted]; Nethergeld Cuffs (254061, -0.52 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of the Greatfather (17721, -0.11 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.16 DPS) [crafted]; Raider Handguards (272102, -1.48 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.9 spell_power points (2.47 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.35 DPS) [dungeon]; Dawnspire Cord (12466, -0.40 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.9 spell_power points (2.28 DPS) | yes | Big Voodoo Pants (8202, -0.06 DPS) [crafted]; Turtle Scale Leggings (8185, -0.44 DPS) [crafted]; Wizardweave Leggings (14132, -0.47 DPS) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 24.1 spell_power points (2.29 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -0.01 DPS) [crafted]; Greaves of Withering Despair (22240, -0.46 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (1.35 DPS) | yes | Snake Hoop (6750, -0.38 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.58 DPS) [quest]; Band of the Unicorn (7553, -2.20 DPS, sim-verified) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.18 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.38 DPS) [quest]; Band of the Unicorn (7553, -1.56 DPS, sim-verified) [world_drop] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.29 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.02 DPS) [dungeon]; Hammer of the Northern Wind (810, -15.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; feet: Skycaller's Leather Boots; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spellshifter Rod

No-known-source sample (15 of 765, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60 (dwarf, 5530311300103050-030000000000000000-5533020000000000)

Set DPS (verified): 170.2. Weights run: 1.5s. Verify run: 10.6s. 1473 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.906 ± 0.037, crit=0.320 ± 0.013 per rating point (14 rating = 1%, 4.477 per %), hit=0.542 ± 0.041 per rating point (10 rating = 1%, 5.416 per %), spell_haste=not significant (-0.546 ± 0.466), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Black Dragonscale Helm (252605, -0.27 DPS) [crafted]; Magister's Crown (16686, -0.36 DPS) [dungeon]; Blue Dragonscale Helm (252604, -6.05 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 26.8 spell_power points (3.05 DPS) | yes | Beads of Ogre Mojo (22149, -0.33 DPS) [quest]; Orb of the Darkmoon (19426, -0.55 DPS) [quest]; Chains of the Lich (23125, -0.55 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -2.14 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.7 spell_power points (3.27 DPS) | yes | Hide of the Wild (18510, -0.64 DPS) [crafted]; Spritecaster Cape (11623, -1.05 DPS) [dungeon]; Runecloth Cloak (13860, -1.42 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Robes (16688, +0.00 DPS) [dungeon]; Chain of The Five Thunders (227024, +0.00 DPS) [vendor]; Vest of Elements (16666, -6.39 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.3 spell_power points (3.34 DPS) | yes | Runecloth Cuffs (254123, -1.05 DPS) [crafted]; Runic Leather Bracers (15092, -1.16 DPS) [crafted]; Bands of The Five Thunders (227017, -1.38 DPS) [vendor] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (170.2 DPS) | yes | Raider Handguards (272102, -0.46 DPS) [vendor]; Gloves of Undead Cleansing (23084, -0.62 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.4 spell_power points (5.74 DPS) | yes | Stormseeker's Girdle (272399, -1.44 DPS) [vendor]; Girdle of Insight (18504, -1.65 DPS) [crafted]; Belt of the Archmage (18405, -3.48 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.8 spell_power points (6.70 DPS) | yes | Sentinel's Silk Leggings (237815, -1.00 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.00 DPS) [vendor]; Red Dragonscale Leggings (252603, -1.12 DPS) [crafted] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 35.5 spell_power points (4.05 DPS) | yes | Skycaller's Leather Boots (252471, -1.20 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.20 DPS) [crafted]; Omnicast Boots (11822, -2.56 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 41.6 spell_power points (4.74 DPS) | yes | Maiden's Circle (13001, -1.96 DPS) [world_drop]; Deep Rooted Ring (19109, -2.34 DPS) [world]; Channeler's Ring (272406, -2.68 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 32.0 spell_power points (3.65 DPS) | yes | Deep Rooted Ring (19109, -1.24 DPS) [world]; Channeler's Ring (272406, -1.58 DPS) [vendor]; Maiden's Circle (13001, -1.63 DPS, sim-verified) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+8.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.91 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -1.04 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.45 DPS) [dungeon]; Hand of Edward the Odd (2243, -15.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Totem of Thunder

No-known-source sample (15 of 1473, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 5530311300103051-020000000000000000-0533520000000000)

Set DPS (verified): 492.7. Weights run: 1.4s. Verify run: 8.7s. 1473 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.341 ± 0.024, crit=0.324 ± 0.011 per rating point (14 rating = 1%, 4.538 per %), hit=0.637 ± 0.054 per rating point (10 rating = 1%, 6.366 per %), spell_haste=not significant (2.530 ± 0.677), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.780 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (492.7 DPS) | yes | Soothsayer's Headdress (17740, -1.18 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -1.62 DPS) [crafted]; Blue Dragonscale Helm (252604, -4.92 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.61 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.00 DPS) [quest]; Beads of Ogre Mojo (22149, -1.92 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.7 spell_power points (14.35 DPS) | yes | Darkspear Shoulderguards (272958, -4.48 DPS) [vendor]; Argent Shoulders (19059, -4.56 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.03 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.1 spell_power points (9.82 DPS) | yes | Hide of the Wild (18510, -3.01 DPS) [crafted]; Spritecaster Cape (11623, -3.54 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -4.34 DPS) [rep] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | 29.8 spell_power points (11.67 DPS) | yes | Fernpulse Jerkin (15786, -0.43 DPS) [quest]; Chain of The Five Thunders (227024, -0.58 DPS) [vendor]; Acumen Robes (17775, -1.57 DPS) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.7 spell_power points (9.68 DPS) | yes | Runecloth Cuffs (254123, -4.04 DPS) [crafted]; Runic Leather Bracers (15092, -4.43 DPS) [crafted]; Skycaller's Leather Bracers (252542, -4.83 DPS) [crafted] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | 24.4 spell_power points (9.54 DPS) | yes | Gloves of the Greatfather (17721, -0.14 DPS) [crafted]; Raider Handguards (272101, -0.47 DPS) [vendor]; Earth Warder's Gloves (21318, -0.93 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.9 spell_power points (15.22 DPS) | yes | Belt of the Archmage (18405, -5.34 DPS, sim-verified) [crafted]; Stormseeker's Girdle (272399, -5.75 DPS) [vendor]; Girdle of Insight (18504, -6.27 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 48.7 spell_power points (19.05 DPS) | yes | Sentinel's Silk Leggings (237815, -1.64 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.64 DPS) [vendor]; Pristine Scorpid Leggings (252606, -4.36 DPS) [crafted] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 26.5 spell_power points (10.36 DPS) | yes | Omnicast Boots (11822, -0.93 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.96 DPS) [crafted]; Venomspew Footpads (275606, -1.74 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 37.4 spell_power points (14.65 DPS) | yes | Maiden's Circle (13001, -6.67 DPS) [world_drop]; Blessed Band of Light (272407, -8.25 DPS) [vendor]; Deep Rooted Ring (19109, -8.62 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 25.8 spell_power points (10.08 DPS) | yes | Maiden's Circle (13001, -3.56 DPS, sim-verified) [world_drop]; Blessed Band of Light (272407, -3.67 DPS) [vendor]; Deep Rooted Ring (19109, -4.05 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+18.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.95 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.13 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.24 DPS) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.82 DPS) [world]; Hand of Edward the Odd (2243, -40.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, -2.87 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Storm Gauntlets; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Totem of Thunder

No-known-source sample (15 of 1473, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.7. Weights run: 1.5s. Verify run: 2.1s. 197 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.656 ± 0.012, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.003 per %), hit=0.215 ± 0.009 per rating point (10 rating = 1%, 2.147 per %), spell_haste=-1.707 ± 0.106, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.33 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (0.91 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.58 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Battle Healer's Cloak (20427, -0.08 DPS) [rep]; Pearl-clasped Cloak (5542, -0.40 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.03 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.25 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.3 spell_power points (0.27 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.02 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.6 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.14 DPS) [crafted]; Serpent Gloves (5970, -0.42 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.46 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.19 DPS) | yes | Stormrider's Leather Pants (252502, -0.19 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (0.81 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.30 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.42 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; The 1 Ring (8350, -0.36 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.9 spell_power points (0.33 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS) [quest]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; The 1 Ring (8350, -0.27 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.12 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.34 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon; 206389 Brimstone Icon; 206390 Ichthyan Icon; 206391 Syzygic Icon; 241089 Scarlet Dagger

### Band 30 (orc, 5530311300000000-000000000000000000-0000000000000000)

Set DPS (verified): 48.6. Weights run: 1.5s. Verify run: 2.0s. 363 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.635 ± 0.014, crit=0.088 ± 0.003 per rating point (14 rating = 1%, 1.235 per %), hit=0.259 ± 0.011 per rating point (10 rating = 1%, 2.590 per %), spell_haste=-2.472 ± 0.191, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.944 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.28 DPS) | yes | Holy Shroud (2721, -0.14 DPS) [world_drop]; Enduring Cap (3020, -0.23 DPS) [world_drop]; Totemic Leather Helm (252456, -0.82 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.12 DPS) | yes | Darkspear Warding Pendant (272075, -0.79 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.13 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (1.53 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Fairywing Mantle (9536, -0.31 DPS) [quest]; Magician's Mantle (12998, -0.41 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 5.1 spell_power points (0.53 DPS) | yes | Windsong Drape (15468, -0.01 DPS) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Hillman's Cloak (3719, -0.33 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (1.79 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Guardian Armor (4256, -0.20 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.93 DPS) | yes | Nightsky Wristbands (6407, -0.54 DPS) [world_drop]; Technician's Bracers (270042, -0.54 DPS) [quest]; Glowing Magical Bracelets (13106, -0.77 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.3 spell_power points (1.07 DPS) | yes | Jutebraid Gloves (10654, -0.11 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.34 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.8 spell_power points (1.43 DPS) | yes | Moss Cinch (6911, -0.19 DPS) [dungeon]; Warsong Sash (16975, -0.29 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.38 DPS, sim-verified) [rep] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (1.46 DPS) | yes | Stormrider's Leather Pants (252502, -0.03 DPS) [crafted]; Guardian Pants (5962, -0.10 DPS) [crafted]; Dark Ritual Leggings (270031, -0.42 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.4 spell_power points (1.19 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.24 DPS) [crafted]; Acidic Walkers (9454, -1.07 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (0.95 DPS) | yes | Ring of Power Regulation (273030, -0.33 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor]; Solid Gold Band (279845, -0.54 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.73 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Ring of Power Regulation (273030, -0.10 DPS) [dungeon]; Solid Gold Band (279845, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (48.6 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, -2.90 DPS) [dungeon]; Manual Crowd Pummeler (9449, -7.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Spine of the Devourer

No-known-source sample (15 of 363, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (orc, 5530311300103050-010000000000000000-0000000000000000)

Set DPS (verified): 67.9. Weights run: 1.5s. Verify run: 2.6s. 558 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.893 ± 0.034, crit=0.249 ± 0.010 per rating point (14 rating = 1%, 3.492 per %), hit=0.431 ± 0.030 per rating point (10 rating = 1%, 4.309 per %), spell_haste=-2.478 ± 0.395, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.944 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.2 spell_power points (2.35 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.22 DPS) [crafted]; Augural Shroud (2620, -0.33 DPS) [world]; Big Voodoo Mask (8201, -1.63 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.1 spell_power points (1.44 DPS) | yes | Triune Amulet (7722, -0.62 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Scorn's Icy Choker (23169, -1.56 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.6 spell_power points (1.89 DPS) | yes | Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -1.51 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.5 spell_power points (1.06 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.06 DPS) [vendor]; Blackforge Cape (6424, -0.34 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.4 spell_power points (2.78 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of Power (7054, -0.27 DPS) [crafted]; Big Voodoo Robe (8200, -0.59 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.4 spell_power points (1.15 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.02 DPS) [quest]; Turtle Scale Bracers (8198, -0.09 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.44 DPS) | yes | Red Mageweave Gloves (10018, -0.41 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.50 DPS) [crafted]; Dreamweave Gloves (10019, -1.38 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.4 spell_power points (2.07 DPS) | yes | Skycaller's Leather Belt (252522, -0.51 DPS) [crafted]; Gilded Cord (254037, -0.53 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.35 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.7 spell_power points (2.51 DPS) | yes | Crimson Silk Pantaloons (7062, -0.73 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.87 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.95 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.44 DPS) | yes | Skycaller's Leather Shoes (252532, -0.48 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.48 DPS) [crafted]; Mender's Leather Shoes (252533, -0.99 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.47 DPS) | yes | Advisor's Ring (19520, -0.56 DPS) [rep]; Reedknot Ring (9622, -0.76 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.5 spell_power points (1.06 DPS) | yes | Advisor's Ring (19520, -0.15 DPS) [rep]; Reedknot Ring (9622, -0.35 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (67.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.18 DPS) [dungeon]; Spine of the Devourer (273040, -1.54 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.42 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Mystic Insight (249394, -0.17 DPS) [crafted]; Anchorhold Buckler (15865, -0.37 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 558, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (orc, 5530311300103050-030000000000000000-5300000000000000)

Set DPS (verified): 97.8. Weights run: 1.5s. Verify run: 3.2s. 708 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.825 ± 0.041, crit=0.297 ± 0.012 per rating point (14 rating = 1%, 4.152 per %), hit=0.578 ± 0.040 per rating point (10 rating = 1%, 5.776 per %), spell_haste=-2.453 ± 0.478, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.936 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.5 spell_power points (3.38 DPS) | yes | Soothsayer's Headdress (17740, -0.01 DPS) [dungeon]; Dreamweave Circlet (10041, -0.60 DPS) [crafted]; Blood Guard's Pulsing Helmet (220848, -0.73 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 spell_power points (1.30 DPS) | yes | Horizon Choker (13085, -0.20 DPS) [world_drop]; Mindburst Medallion (11196, -0.25 DPS) [quest]; Scorn's Icy Choker (23169, -1.58 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 31.5 spell_power points (3.00 DPS) | yes | Rockshard Pauldrons (9411, -0.84 DPS) [dungeon]; Skycaller's Leather Shoulder (252537, -1.06 DPS) [crafted]; Rotgrip Mantle (17732, -1.72 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.4 spell_power points (1.85 DPS) | yes | Spritecaster Cape (11623, -0.05 DPS) [dungeon]; Runecloth Cloak (13860, -0.36 DPS) [crafted]; Big Voodoo Cloak (8216, -0.67 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.5 spell_power points (3.38 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.17 DPS) [vendor]; Feathered Breastplate (8349, -0.79 DPS) [crafted]; Robe of the Magi (1716, -0.81 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.2 spell_power points (1.74 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.24 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of the Greatfather (17721, -0.11 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.16 DPS) [crafted]; Raider Handguards (272102, -1.16 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.9 spell_power points (2.47 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.35 DPS) [dungeon]; Dawnspire Cord (12466, -0.40 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 33.7 spell_power points (3.21 DPS) | yes | Red Mageweave Pants (10009, -0.95 DPS, sim-verified) [crafted]; Big Voodoo Pants (8202, -0.99 DPS) [crafted]; Turtle Scale Leggings (8185, -1.37 DPS) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 24.1 spell_power points (2.29 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -0.01 DPS) [crafted]; Greaves of Withering Despair (22240, -0.46 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (1.35 DPS) | yes | Snake Hoop (6750, -0.38 DPS) [quest]; Reedknot Ring (9622, -0.68 DPS) [quest]; Band of the Unicorn (7553, -3.07 DPS, sim-verified) [world_drop] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.18 DPS) [quest]; Reedknot Ring (9622, -0.48 DPS) [quest]; Band of the Unicorn (7553, -1.38 DPS, sim-verified) [world_drop] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.19 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.11 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.29 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.02 DPS) [dungeon]; Hammer of the Northern Wind (810, -15.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; feet: Skycaller's Leather Boots; finger2: Advisor's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spellshifter Rod

No-known-source sample (15 of 708, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (orc, 5530311300103050-030000000000000000-5533020000000000)

Set DPS (verified): 167.8. Weights run: 1.5s. Verify run: 9.2s. 1396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.906 ± 0.037, crit=0.320 ± 0.013 per rating point (14 rating = 1%, 4.477 per %), hit=0.542 ± 0.041 per rating point (10 rating = 1%, 5.416 per %), spell_haste=not significant (-0.546 ± 0.466), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Coif of The Five Thunders (227002) | Saving the Best for Last [quest] | 48.1 spell_power points (5.49 DPS) | yes | Blue Dragonscale Helm (252604, -0.16 DPS) [crafted]; Warlord's Mail Helm (231663, -0.29 DPS) [pvp]; Living Crown (252561, -0.78 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 26.8 spell_power points (3.05 DPS) | yes | Beads of Ogre Mojo (22149, -0.33 DPS) [quest]; Orb of the Darkmoon (19426, -0.55 DPS) [quest]; Chains of the Lich (23125, -0.55 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.1 spell_power points (5.14 DPS) | yes | Darkspear Shoulderguards (272958, -0.52 DPS) [vendor]; Warlord's Mail Spaulders (231659, -0.58 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.98 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.7 spell_power points (3.27 DPS) | yes | Hide of the Wild (18510, -0.64 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.97 DPS) [quest]; Spritecaster Cape (11623, -1.05 DPS) [dungeon] |
| chest | Vest of The Five Thunders (227004) | Saving the Best for Last [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Vest of Elements (16666, -5.19 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.3 spell_power points (3.34 DPS) | yes | Runecloth Cuffs (254123, -1.05 DPS) [crafted]; Runic Leather Bracers (15092, -1.16 DPS) [crafted]; Bindings of The Five Thunders (227001, -1.29 DPS) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -0.05 DPS) [vendor]; Raider Handguards (272101, -2.02 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.4 spell_power points (5.74 DPS) | yes | Stormseeker's Girdle (272399, -1.44 DPS) [vendor]; Girdle of Insight (18504, -1.65 DPS) [crafted]; Belt of the Archmage (18405, -3.52 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.8 spell_power points (6.70 DPS) | yes | General's Mail Leggings (231664, -0.52 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -1.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -1.79 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 35.5 spell_power points (4.05 DPS) | yes | General's Mail Sabatons (231661, -0.01 DPS) [vendor]; General's Mail Boots (231667, -0.80 DPS) [pvp]; Omnicast Boots (11822, -2.88 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 41.6 spell_power points (4.74 DPS) | yes | Maiden's Circle (13001, -1.96 DPS) [world_drop]; Deep Rooted Ring (19109, -2.34 DPS) [world]; Channeler's Ring (272406, -2.68 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 32.0 spell_power points (3.65 DPS) | yes | Maiden's Circle (13001, -0.87 DPS) [world_drop]; Deep Rooted Ring (19109, -1.24 DPS) [world]; Channeler's Ring (272406, -1.58 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+8.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.91 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -1.04 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.45 DPS) [dungeon]; Hand of Edward the Odd (2243, -16.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Coif of The Five Thunders; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of The Five Thunders; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Totem of Thunder

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (orc, 5530311300103051-020000000000000000-0533520000000000)

Set DPS (verified): 482.3. Weights run: 1.4s. Verify run: 9.2s. 1396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.341 ± 0.024, crit=0.324 ± 0.011 per rating point (14 rating = 1%, 4.538 per %), hit=0.637 ± 0.054 per rating point (10 rating = 1%, 6.366 per %), spell_haste=not significant (2.530 ± 0.677), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.780 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 37.5 spell_power points (14.68 DPS) | yes | Warlord's Mail Helm (231663, -1.23 DPS) [pvp]; Coif of The Five Thunders (227002, -1.96 DPS, sim-verified) [quest]; Living Crown (252561, -2.49 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.61 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.00 DPS) [quest]; Beads of Ogre Mojo (22149, -1.92 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.7 spell_power points (14.35 DPS) | yes | Warlord's Mail Spaulders (231659, -2.22 DPS) [pvp]; Pauldrons of The Five Thunders (227003, -3.95 DPS, sim-verified) [quest]; Mantle of the Timbermaw (19050, -4.18 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.1 spell_power points (9.82 DPS) | yes | Hide of the Wild (18510, -3.01 DPS) [crafted]; Spritecaster Cape (11623, -3.54 DPS) [dungeon]; Deep Woodlands Cloak (19121, -3.92 DPS) [quest] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | 29.8 spell_power points (11.67 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -0.43 DPS) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.7 spell_power points (9.68 DPS) | yes | Runecloth Cuffs (254123, -4.04 DPS) [crafted]; Runic Leather Bracers (15092, -4.43 DPS) [crafted]; Skycaller's Leather Bracers (252542, -4.83 DPS) [crafted] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | 24.4 spell_power points (9.54 DPS) | yes | Gloves of the Greatfather (17721, +0.00 DPS, sim-verified) [crafted]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -0.28 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.9 spell_power points (15.22 DPS) | yes | Belt of the Archmage (18405, -3.90 DPS, sim-verified) [crafted]; Stormseeker's Girdle (272399, -5.75 DPS) [vendor]; Girdle of Insight (18504, -6.27 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 48.7 spell_power points (19.05 DPS) | yes | Sentinel's Silk Leggings (237815, -1.64 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.64 DPS) [vendor]; General's Mail Leggings (231664, -3.11 DPS) [pvp] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 26.5 spell_power points (10.36 DPS) | yes | General's Mail Sabatons (231661, -0.26 DPS) [vendor]; Omnicast Boots (11822, -0.93 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.96 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 37.4 spell_power points (14.65 DPS) | yes | Maiden's Circle (13001, -6.67 DPS) [world_drop]; Blessed Band of Light (272407, -8.25 DPS) [vendor]; Deep Rooted Ring (19109, -8.62 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 25.8 spell_power points (10.08 DPS) | yes | Maiden's Circle (13001, -2.59 DPS, sim-verified) [world_drop]; Blessed Band of Light (272407, -3.67 DPS) [vendor]; Deep Rooted Ring (19109, -4.05 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (482.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (482.3 DPS) | yes | Weakness Analyzer (272438, -2.92 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.13 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.24 DPS) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (482.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.82 DPS) [world]; Hand of Edward the Odd (2243, -38.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (482.3 DPS) | yes | Totem of the Storm (272432, -3.01 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Storm Gauntlets; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Totem of Thunder

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

