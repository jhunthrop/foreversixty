# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.5. Weights run: 2.8s. Verify run: 1.6s. 227 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.577 ± 0.009, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.031 per %), hit=0.188 ± 0.007 per rating point (10 rating = 1%, 1.884 per %), spell_haste=-1.782 ± 0.126, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.872 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.12 DPS) [crafted]; Totemic Leather Hood (252448, -0.31 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Double-Stitched Woolen Shoulders (4314, -0.28 DPS) [crafted]; Magician's Mantle (12998, -0.35 DPS, sim-verified) [world_drop] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Whelp Cloak (7283, -0.06 DPS) [crafted]; Sanguine Cape (14376, -0.12 DPS) [world_drop]; Heavy Woolen Cloak (4311, -0.29 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.01 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.23 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.05 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.10 DPS) [dungeon]; Owl Bracers (4796, -1.03 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.3 spell_power points (0.61 DPS) | yes | Gnoll Casting Gloves (892, -0.11 DPS) [world]; Windfelt Gloves (5630, -0.13 DPS) [quest]; Serpent Gloves (5970, -0.17 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.43 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamer's Leggings (270016, -0.01 DPS) [quest]; Filigreed Pristine Leggings (253937, -0.08 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.43 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.3 spell_power points (0.78 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.28 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.51 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.46 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.5 spell_power points (0.29 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.24 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Twisted Chanter's Staff (890, -0.19 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.38 DPS) [world]; Staff of Westfall (2042, -0.43 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Reinforced Woolen Shoulders; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 227, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (dwarf, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 48.2. Weights run: 3.1s. Verify run: 1.7s. 427 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.580 ± 0.013, crit=0.090 ± 0.002 per rating point (14 rating = 1%, 1.261 per %), hit=0.259 ± 0.009 per rating point (10 rating = 1%, 2.585 per %), spell_haste=-2.286 ± 0.162, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.819 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (1.31 DPS) | yes | Enchanter's Cowl (4322, -0.02 DPS) [crafted]; Holy Shroud (2721, -0.11 DPS) [world_drop]; Enduring Cap (3020, -0.30 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.5 spell_power points (0.38 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.2 spell_power points (1.55 DPS) | yes | Death Speaker Mantle (6685, -0.20 DPS) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Magician's Mantle (12998, -0.44 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.54 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Prelacy Cape (7004, -0.11 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.5 spell_power points (1.80 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Guardian Armor (4256, -0.19 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Nightsky Wristbands (6407, -0.60 DPS) [world_drop]; Technician's Bracers (270042, -0.60 DPS) [quest]; Glowing Magical Bracelets (13106, -1.01 DPS, sim-verified) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.3 spell_power points (0.80 DPS) | yes | Town Clerk's Mittens (270029, -0.03 DPS) [quest]; Shilly Mitts (9609, -0.03 DPS) [quest]; Gloves of Insight (9698, -0.03 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.5 spell_power points (1.47 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Crimson Silk Belt (7055, -0.37 DPS) [crafted]; Moss Cinch (6911, -1.23 DPS, sim-verified) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.09 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.73 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.1 spell_power points (1.21 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.24 DPS) [crafted]; Acidic Walkers (9454, -1.31 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.9 spell_power points (0.97 DPS) | yes | Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon]; Minor Channeling Ring (1449, -1.94 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Minor Channeling Ring (1449, -0.57 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.74 DPS) [quest]; Manual Crowd Pummeler (9449, -7.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 427, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (dwarf, 5510331300001350-010000000000000000-0000000000000000)

Set DPS (verified): 70.2. Weights run: 3.4s. Verify run: 2.1s. 676 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.919 ± 0.026, crit=0.249 ± 0.008 per rating point (14 rating = 1%, 3.483 per %), hit=0.390 ± 0.018 per rating point (10 rating = 1%, 3.903 per %), spell_haste=-2.318 ± 0.289, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.695 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.6 spell_power points (2.49 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.28 DPS) [crafted]; Augural Shroud (2620, -0.36 DPS) [world]; Big Voodoo Mask (8201, -1.85 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.4 spell_power points (1.51 DPS) | yes | Necklace of Calisea (1714, -0.83 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.93 DPS) [world]; Triune Amulet (7722, -1.21 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.9 spell_power points (2.00 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -1.19 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.6 spell_power points (1.12 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.34 DPS) [dungeon]; Icy Cloak (4327, -0.38 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.5 spell_power points (2.90 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of Power (7054, -0.26 DPS) [crafted]; Big Voodoo Robe (8200, -0.59 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.5 spell_power points (1.21 DPS) | yes | Turtle Scale Bracers (8198, -0.10 DPS) [crafted]; Arcane Runed Bracers (4744, -0.26 DPS) [quest]; Spidertank Oilrag (9448, -0.26 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.53 DPS) | yes | Red Mageweave Gloves (10018, -0.40 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.50 DPS) [crafted]; Dreamweave Gloves (10019, -1.36 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.8 spell_power points (2.19 DPS) | yes | Gilded Cord (254037, -0.57 DPS) [crafted]; Prefect's Belt (250559, -0.77 DPS) [crafted]; Skycaller's Leather Belt (252522, -1.24 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.0 spell_power points (2.64 DPS) | yes | Crimson Silk Pantaloons (7062, -0.74 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -1.00 DPS) [crafted]; Guardian Pants (5962, -1.02 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.53 DPS) | yes | Skycaller's Mail Boots (252563, -0.48 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.94 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.01 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.54 DPS) | yes | Ring of Forlorn Spirits (2043, -0.69 DPS) [quest]; Reedknot Ring (9622, -0.80 DPS) [quest]; Minor Channeling Ring (1449, -0.82 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.6 spell_power points (1.12 DPS) | yes | Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.38 DPS) [quest]; Minor Channeling Ring (1449, -0.40 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (70.2 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Spine of the Devourer (273040, -1.62 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 15.2 spell_power points (1.60 DPS) | yes | Thrash's Trash (276204, -0.13 DPS) [vendor]; Orb of Mystic Insight (249394, -0.28 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.36 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica

No-known-source sample (15 of 676, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 5510331300001350-030000000000000000-5030000000000000)

Set DPS (verified): 115.3. Weights run: 3.4s. Verify run: 2.4s. 879 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.757 ± 0.028, crit=0.294 ± 0.009 per rating point (14 rating = 1%, 4.112 per %), hit=0.460 ± 0.023 per rating point (10 rating = 1%, 4.604 per %), spell_haste=not significant (-0.681 ± 0.309), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.698 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 34.3 spell_power points (3.88 DPS) | yes | Dreamweave Circlet (10041, -0.65 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.83 DPS) [crafted]; Red Mageweave Headband (10033, -1.39 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.1 spell_power points (1.48 DPS) | yes | Horizon Choker (13085, -0.28 DPS) [world_drop]; Mindburst Medallion (11196, -0.28 DPS) [quest]; Triune Amulet (7722, -0.57 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 30.1 spell_power points (3.41 DPS) | yes | Rockshard Pauldrons (9411, -0.94 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -1.17 DPS) [crafted]; Rotgrip Mantle (17732, -1.71 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.5 spell_power points (2.10 DPS) | yes | Big Voodoo Cloak (8216, -0.76 DPS) [crafted]; Long Silken Cloak (4326, -0.99 DPS) [crafted]; Runecloth Cloak (13860, -1.24 DPS, sim-verified) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (115.3 DPS) | yes | Robe of the Magi (1716, -0.00 DPS) [world_drop]; Runecloth Tunic (13857, -0.14 DPS) [crafted]; Acumen Robes (17775, -1.48 DPS, sim-verified) [quest] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.6 spell_power points (1.99 DPS) | yes | Skycaller's Leather Bracers (252542, -0.26 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.26 DPS) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.71 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.13 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.13 DPS) [crafted]; Dreamweave Gloves (10019, -0.34 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.1 spell_power points (2.83 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.40 DPS) [dungeon]; Dawnspire Cord (12466, -0.53 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.1 spell_power points (2.61 DPS) | yes | Big Voodoo Pants (8202, -0.06 DPS) [crafted]; Wizardweave Leggings (14132, -0.46 DPS) [crafted]; Turtle Scale Leggings (8185, -0.51 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.71 DPS) | yes | Skycaller's Leather Boots (252471, -0.08 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.08 DPS) [crafted]; Greaves of Withering Despair (22240, -0.61 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (1.56 DPS) | yes | Snake Hoop (6750, -0.45 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.65 DPS) [quest]; Reedknot Ring (9622, -0.77 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.47 DPS) | yes | Snake Hoop (6750, -0.36 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.57 DPS) [quest]; Reedknot Ring (9622, -0.68 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -1.21 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.06 DPS) [quest]; Kindling Stave (11750, -0.32 DPS) [dungeon]; Hammer of the Northern Wind (810, -24.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 5510331300001350-030000000000000000-5033520000000000)

Set DPS (verified): 131.3. Weights run: 3.4s. Verify run: 7.3s. 1826 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.830 ± 0.034, crit=0.353 ± 0.010 per rating point (14 rating = 1%, 4.946 per %), hit=0.556 ± 0.026 per rating point (10 rating = 1%, 5.563 per %), spell_haste=not significant (-1.163 ± 0.375), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.709 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (131.3 DPS) | yes | Black Dragonscale Helm (252605, -0.41 DPS) [crafted]; Magister's Crown (16686, -0.47 DPS) [dungeon]; Blue Dragonscale Helm (252604, -4.43 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.3 spell_power points (2.39 DPS) | yes | Glowing Eye of Mordresh (10769, -0.78 DPS) [dungeon]; Horizon Choker (13085, -1.02 DPS) [world_drop]; Mindburst Medallion (11196, -1.10 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (131.3 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS) [dungeon]; Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Ironfeather Shoulders (15067, -4.73 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.0 spell_power points (2.23 DPS) | yes | Big Voodoo Cloak (8216, -0.77 DPS) [crafted]; Long Silken Cloak (4326, -1.04 DPS) [crafted]; Runecloth Cloak (13860, -1.76 DPS, sim-verified) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (131.3 DPS) | yes | Vest of Elements (16666, +0.00 DPS) [dungeon]; Acumen Robes (17775, +0.00 DPS) [quest]; Ironfeather Breastplate (15066, -4.63 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (131.3 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -5.92 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (131.3 DPS) | yes | Storm Gauntlets (12632, +0.00 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, -5.61 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (131.3 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -6.49 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 57.9 spell_power points (6.81 DPS) | yes | Red Dragonscale Leggings (252603, -1.21 DPS, sim-verified) [crafted]; Luminary Kilt (11823, -2.27 DPS) [dungeon]; Pristine Scorpid Leggings (252606, -2.34 DPS) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 30.0 spell_power points (3.52 DPS) | yes | Skycaller's Leather Boots (252471, -0.69 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.69 DPS) [crafted]; Earthen Silk Slippers (254013, -0.70 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.8 spell_power points (2.80 DPS) | yes | Band of the Unicorn (7553, -1.27 DPS) [world_drop]; Snake Hoop (6750, -1.61 DPS) [quest]; Glowing Crystal Ring (18402, -1.63 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (1.66 DPS) | yes | Band of the Unicorn (7553, -0.14 DPS) [world_drop]; Snake Hoop (6750, -0.47 DPS) [quest]; Glowing Crystal Ring (18402, -0.49 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (131.3 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (131.3 DPS) | yes | Mark of the Chosen (17774, -2.40 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (131.3 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -3.14 DPS) [world_drop] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (2.47 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Truesilver Conduit (249455, +0.00 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1826, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (dwarf, 5510331300001351-055000000000000000-0530020000000000)

Set DPS (verified): 463.3. Weights run: 2.0s. Verify run: 5.8s. 1826 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.448 ± 0.021, crit=0.360 ± 0.008 per rating point (14 rating = 1%, 5.040 per %), hit=0.623 ± 0.038 per rating point (10 rating = 1%, 6.234 per %), spell_haste=4.234 ± 0.509, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.670 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soothsayer's Headdress (17740, -1.47 DPS) [dungeon]; Red Mageweave Headband (10033, -2.24 DPS) [crafted]; Blue Dragonscale Helm (252604, -8.53 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.5 spell_power points (7.23 DPS) | yes | Glowing Eye of Mordresh (10769, -2.59 DPS) [dungeon]; Mindburst Medallion (11196, -3.42 DPS) [quest]; Triune Amulet (7722, -3.72 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Rotgrip Mantle (17732, -1.27 DPS) [dungeon]; Mantle of the Timbermaw (19050, -9.16 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.7 spell_power points (7.32 DPS) | yes | Runecloth Cloak (13860, -1.80 DPS) [crafted]; Big Voodoo Cloak (8216, -3.36 DPS) [crafted]; Nightfall Drape (12465, -3.37 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fernpulse Jerkin (15786, +0.00 DPS) [quest]; Acumen Robes (17775, +0.00 DPS) [quest]; Vest of Elements (16666, -8.75 DPS, sim-verified) [dungeon] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 15.5 spell_power points (6.79 DPS) | yes | Runic Leather Bracers (15092, -0.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.03 DPS) [crafted]; Skycaller's Mail Bracers (252571, -1.03 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (463.3 DPS) | yes | Earth Warder's Gloves (21318, -0.88 DPS) [quest]; Skycaller's Leather Gauntlets (252550, -1.74 DPS) [crafted]; Storm Gauntlets (12632, -4.78 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 21.4 spell_power points (9.38 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -1.27 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.80 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.1 spell_power points (22.41 DPS) | yes | Red Dragonscale Leggings (252603, -7.35 DPS) [crafted]; Luminary Kilt (11823, -8.83 DPS) [dungeon]; Pristine Scorpid Leggings (252606, -11.35 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.4 spell_power points (11.13 DPS) | yes | Earthen Silk Slippers (254013, -0.60 DPS) [crafted]; Venomspew Footpads (275606, -1.48 DPS) [crafted]; Bileblister Boots (275612, -1.48 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.1 spell_power points (9.27 DPS) | yes | Philanthropist's Ring (281635, -3.90 DPS) [quest]; Snake Hoop (6750, -5.66 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.76 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.70 DPS) | yes | Philanthropist's Ring (281635, -0.33 DPS) [quest]; Snake Hoop (6750, -2.09 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.19 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -6.22 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -13.35 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.21 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Orb of the Forgotten Seer (7685, -3.63 DPS, sim-verified) [dungeon] |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Totem of Thunder

No-known-source sample (15 of 1826, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.0. Weights run: 2.8s. Verify run: 1.5s. 216 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.577 ± 0.009, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.031 per %), hit=0.188 ± 0.007 per rating point (10 rating = 1%, 1.884 per %), spell_haste=-1.782 ± 0.126, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.872 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.12 DPS) [crafted]; Totemic Leather Hood (252448, -0.31 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.2 spell_power points (0.85 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.52 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.78 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Sanguine Cape (14376, -0.14 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.16 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.01 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.23 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.25 DPS) | yes | Mindthrust Bracers (1974, -0.01 DPS) [dungeon]; Owl Bracers (4796, -0.01 DPS) [vendor]; Featherbead Bracers (15452, -0.01 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.3 spell_power points (0.61 DPS) | yes | Gnoll Casting Gloves (892, -0.11 DPS) [world]; Pristine Gloves (253913, -0.13 DPS) [crafted]; Serpent Gloves (5970, -0.19 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.43 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.08 DPS) [crafted]; Totemic Leather Pants (252446, -0.10 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.43 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.3 spell_power points (0.78 DPS) | yes | Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.28 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.56 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.5 spell_power points (0.29 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.9 spell_power points (0.24 DPS) | yes | The 1 Ring (8350, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.47 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Twisted Chanter's Staff (890, -0.19 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.38 DPS) [world]; Gnarled Necromancer's Staff (251534, -0.38 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection; 21568 Rune of Duty

### Band 30 (orc, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 48.5. Weights run: 3.1s. Verify run: 1.7s. 420 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.580 ± 0.013, crit=0.090 ± 0.002 per rating point (14 rating = 1%, 1.261 per %), hit=0.259 ± 0.009 per rating point (10 rating = 1%, 2.585 per %), spell_haste=-2.286 ± 0.162, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.819 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.09 DPS) [world_drop]; Enduring Cap (3020, -0.27 DPS) [world_drop]; Totemic Leather Helm (252456, -0.78 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.5 spell_power points (0.38 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.2 spell_power points (1.55 DPS) | yes | Death Speaker Mantle (6685, -0.20 DPS) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Magician's Mantle (12998, -0.44 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.54 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.14 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.11 DPS) [dungeon]; Guardian Armor (4256, -0.13 DPS) [crafted]; Green Silk Armor (7065, -0.54 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Nightsky Wristbands (6407, -0.60 DPS) [world_drop]; Technician's Bracers (270042, -0.60 DPS) [quest]; Glowing Magical Bracelets (13106, -0.98 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.2 spell_power points (1.11 DPS) | yes | Jutebraid Gloves (10654, -0.14 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.31 DPS) [crafted]; Serpent Gloves (5970, -0.34 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.5 spell_power points (1.47 DPS) | yes | Warsong Sash (16975, -0.27 DPS) [quest]; Belt of Arugal (6392, -0.30 DPS) [dungeon]; Moss Cinch (6911, -0.70 DPS, sim-verified) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.09 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.84 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.1 spell_power points (1.21 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.24 DPS) [crafted]; Acidic Walkers (9454, -1.35 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.9 spell_power points (0.97 DPS) | yes | Solid Gold Band (279845, -0.53 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.65 DPS) | yes | Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Solid Gold Band (279845, -0.44 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, -3.05 DPS) [dungeon]; Manual Crowd Pummeler (9449, -7.08 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 420, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (orc, 5510331300001350-010000000000000000-0000000000000000)

Set DPS (verified): 79.2. Weights run: 3.4s. Verify run: 2.1s. 658 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.919 ± 0.026, crit=0.249 ± 0.008 per rating point (14 rating = 1%, 3.483 per %), hit=0.390 ± 0.018 per rating point (10 rating = 1%, 3.903 per %), spell_haste=-2.318 ± 0.289, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.695 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.6 spell_power points (2.49 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.28 DPS) [crafted]; Augural Shroud (2620, -0.36 DPS) [world]; Big Voodoo Mask (8201, -2.15 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.4 spell_power points (1.51 DPS) | yes | Necklace of Calisea (1714, -0.83 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.93 DPS) [world]; Triune Amulet (7722, -1.38 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.9 spell_power points (2.00 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -1.05 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.6 spell_power points (1.12 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.34 DPS) [dungeon]; Icy Cloak (4327, -0.38 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.5 spell_power points (2.90 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of Power (7054, -0.26 DPS) [crafted]; Big Voodoo Robe (8200, -0.59 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.5 spell_power points (1.21 DPS) | yes | Radiant Silver Bracers (4545, -0.02 DPS) [quest]; Turtle Scale Bracers (8198, -0.10 DPS) [crafted]; Spidertank Oilrag (9448, -0.26 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.53 DPS) | yes | Red Mageweave Gloves (10018, -0.40 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.50 DPS) [crafted]; Dreamweave Gloves (10019, -1.23 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.8 spell_power points (2.19 DPS) | yes | Gilded Cord (254037, -0.57 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.72 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -0.77 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.0 spell_power points (2.64 DPS) | yes | Crimson Silk Pantaloons (7062, -0.88 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -1.00 DPS) [crafted]; Guardian Pants (5962, -1.02 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.53 DPS) | yes | Skycaller's Leather Shoes (252532, -0.48 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.48 DPS) [crafted]; Mender's Leather Shoes (252533, -1.01 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.54 DPS) | yes | Reedknot Ring (9622, -0.80 DPS) [quest]; Ogremind Ring (1993, -0.86 DPS) [world_drop]; Voodoo Band (1996, -0.86 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.6 spell_power points (1.12 DPS) | yes | Ogremind Ring (1993, -0.44 DPS) [world_drop]; Voodoo Band (1996, -0.44 DPS) [world]; Reedknot Ring (9622, -0.96 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Spine of the Devourer (273040, -1.62 DPS) [world] |
| off_hand | Thrash's Trash (276204) | Friz Frazzlespark [vendor] | sim-verified (79.2 DPS) | yes | Orb of Mystic Insight (249394, -0.16 DPS) [crafted]; Anchorhold Buckler (15865, -0.37 DPS) [quest]; Orb of the Forgotten Seer (7685, -10.66 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Thrash's Trash

No-known-source sample (15 of 658, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (orc, 5510331300001350-030000000000000000-5030000000000000)

Set DPS (verified): 113.9. Weights run: 3.4s. Verify run: 2.4s. 849 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.757 ± 0.028, crit=0.294 ± 0.009 per rating point (14 rating = 1%, 4.112 per %), hit=0.460 ± 0.023 per rating point (10 rating = 1%, 4.604 per %), spell_haste=not significant (-0.681 ± 0.309), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.698 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 34.3 spell_power points (3.88 DPS) | yes | Dreamweave Circlet (10041, -0.65 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.83 DPS) [crafted]; Red Mageweave Headband (10033, -1.08 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.1 spell_power points (1.48 DPS) | yes | Horizon Choker (13085, -0.28 DPS) [world_drop]; Mindburst Medallion (11196, -0.28 DPS) [quest]; Triune Amulet (7722, -0.57 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 30.1 spell_power points (3.41 DPS) | yes | Rockshard Pauldrons (9411, -0.94 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -1.17 DPS) [crafted]; Rotgrip Mantle (17732, -2.17 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.8 spell_power points (2.13 DPS) | yes | Spritecaster Cape (11623, -0.03 DPS) [dungeon]; Runecloth Cloak (13860, -0.42 DPS) [crafted]; Big Voodoo Cloak (8216, -0.79 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 34.1 spell_power points (3.86 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.28 DPS) [vendor]; Feathered Breastplate (8349, -0.85 DPS) [crafted]; Robe of the Magi (1716, -0.86 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.6 spell_power points (1.99 DPS) | yes | Skycaller's Leather Bracers (252542, -0.26 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.26 DPS) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.71 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.13 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.13 DPS) [crafted]; Dreamweave Gloves (10019, -0.34 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.1 spell_power points (2.83 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.40 DPS) [dungeon]; Dawnspire Cord (12466, -0.53 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | PvP rank 10 · Stone Guard · Horde [vendor] | 31.7 spell_power points (3.58 DPS) | yes | Big Voodoo Pants (8202, -1.03 DPS) [crafted]; Wizardweave Leggings (14132, -1.43 DPS) [crafted]; Red Mageweave Pants (10009, -2.13 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.71 DPS) | yes | Skycaller's Leather Boots (252471, -0.08 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.08 DPS) [crafted]; Greaves of Withering Despair (22240, -0.61 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (1.56 DPS) | yes | Snake Hoop (6750, -0.45 DPS) [quest]; Reedknot Ring (9622, -0.77 DPS) [quest]; Ring of Power Regulation (273030, -0.88 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.47 DPS) | yes | Snake Hoop (6750, -0.36 DPS) [quest]; Reedknot Ring (9622, -0.68 DPS) [quest]; Ring of Power Regulation (273030, -0.79 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (113.9 DPS) | yes | Frozen Heart of the Mountain (249469, -0.21 DPS) [crafted]; Rune of the Guard Captain (19120, -0.31 DPS) [quest] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (113.9 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (113.9 DPS) | yes | Spellshifter Rod (9527, -0.06 DPS) [quest]; Kindling Stave (11750, -0.32 DPS) [dungeon]; Hammer of the Northern Wind (810, -24.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Mark of the Chosen; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 849, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (orc, 5510331300001350-030000000000000000-5033520000000000)

Set DPS (verified): 130.5. Weights run: 3.4s. Verify run: 7.3s. 1782 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.830 ± 0.034, crit=0.353 ± 0.010 per rating point (14 rating = 1%, 4.946 per %), hit=0.556 ± 0.026 per rating point (10 rating = 1%, 5.563 per %), spell_haste=not significant (-1.163 ± 0.375), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.709 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (130.5 DPS) | yes | Warlord's Mail Helm (231663, +0.00 DPS) [pvp]; Champion's Mail Helm (227162, -0.20 DPS) [pvp]; Blue Dragonscale Helm (252604, -4.28 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.3 spell_power points (2.39 DPS) | yes | Glowing Eye of Mordresh (10769, -0.78 DPS) [dungeon]; Horizon Choker (13085, -1.02 DPS) [world_drop]; Mindburst Medallion (11196, -1.10 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (130.5 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Ironfeather Shoulders (15067, -4.78 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (130.5 DPS) | yes | Deep Woodlands Cloak (19121, +0.00 DPS) [quest]; Runecloth Cloak (13860, -0.39 DPS) [crafted]; Big Voodoo Cloak (8216, -0.77 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (130.5 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -5.17 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (130.5 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -5.03 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (130.5 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, +0.00 DPS) [vendor]; Gloves of the Greatfather (17721, -5.26 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (130.5 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -5.81 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 57.9 spell_power points (6.81 DPS) | yes | General's Mail Leggings (231664, -0.59 DPS) [pvp]; Red Dragonscale Leggings (252603, -1.33 DPS) [crafted]; Legionnaire's Mail Leggings (227161, -2.10 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 30.0 spell_power points (3.52 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; General's Mail Boots (231667, -0.32 DPS) [pvp]; Skycaller's Leather Boots (252471, -1.17 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.8 spell_power points (2.80 DPS) | yes | Band of the Unicorn (7553, -1.27 DPS) [world_drop]; Snake Hoop (6750, -1.61 DPS) [quest]; Glowing Crystal Ring (18402, -1.63 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (1.66 DPS) | yes | Band of the Unicorn (7553, -0.14 DPS) [world_drop]; Snake Hoop (6750, -0.47 DPS) [quest]; Glowing Crystal Ring (18402, -0.49 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (130.5 DPS) | yes | Rune of the Guard Captain (19120, -0.25 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (130.5 DPS) | yes | Rune of the Guard Captain (19120, -0.13 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (130.5 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -3.14 DPS) [world_drop] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (2.47 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Truesilver Conduit (249455, +0.00 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; back: Spritecaster Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Frozen Heart of the Mountain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1782, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (orc, 5510331300001351-055000000000000000-0530020000000000)

Set DPS (verified): 456.6. Weights run: 2.0s. Verify run: 5.8s. 1782 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.448 ± 0.021, crit=0.360 ± 0.008 per rating point (14 rating = 1%, 5.040 per %), hit=0.623 ± 0.038 per rating point (10 rating = 1%, 6.234 per %), spell_haste=4.234 ± 0.509, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.670 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (456.6 DPS) | yes | Warlord's Mail Helm (231663, +0.00 DPS) [pvp]; Champion's Mail Helm (227162, -0.38 DPS) [pvp]; Blue Dragonscale Helm (252604, -8.17 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.5 spell_power points (7.23 DPS) | yes | Glowing Eye of Mordresh (10769, -2.59 DPS) [dungeon]; Mindburst Medallion (11196, -3.42 DPS) [quest]; Triune Amulet (7722, -3.72 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (456.6 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -8.96 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.7 spell_power points (7.32 DPS) | yes | Deep Woodlands Cloak (19121, -0.29 DPS) [quest]; Runecloth Cloak (13860, -1.80 DPS) [crafted]; Big Voodoo Cloak (8216, -3.36 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (456.6 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Vest of Elements (16666, -11.79 DPS, sim-verified) [dungeon] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 15.5 spell_power points (6.79 DPS) | yes | Runic Leather Bracers (15092, -0.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.03 DPS) [crafted]; Skycaller's Mail Bracers (252571, -1.03 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (456.6 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, +0.00 DPS) [vendor]; Storm Gauntlets (12632, -2.74 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 21.4 spell_power points (9.38 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -1.27 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.80 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.1 spell_power points (22.41 DPS) | yes | General's Mail Leggings (231664, -3.21 DPS) [pvp]; Red Dragonscale Leggings (252603, -7.35 DPS) [crafted]; Pristine Scorpid Leggings (252606, -10.88 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.4 spell_power points (11.13 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Earthen Silk Slippers (254013, -0.60 DPS) [crafted]; Venomspew Footpads (275606, -1.48 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.1 spell_power points (9.27 DPS) | yes | Philanthropist's Ring (281635, -3.90 DPS) [quest]; Snake Hoop (6750, -5.66 DPS) [quest]; Reedknot Ring (9622, -6.20 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.70 DPS) | yes | Philanthropist's Ring (281635, -0.33 DPS) [quest]; Snake Hoop (6750, -2.09 DPS) [quest]; Reedknot Ring (9622, -2.63 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (456.6 DPS) | yes | Frozen Heart of the Mountain (249469, -0.17 DPS) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (456.6 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -13.35 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.21 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Orb of the Forgotten Seer (7685, -3.57 DPS, sim-verified) [dungeon] |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (456.6 DPS) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Spritecaster Cape; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Totem of Thunder

No-known-source sample (15 of 1782, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

