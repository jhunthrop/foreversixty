# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.5. Weights run: 2.9s. Verify run: 1.6s. 227 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.577 ± 0.009, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.031 per %), hit=0.188 ± 0.007 per rating point (10 rating = 1%, 1.884 per %), spell_haste=not significant (-0.026 ± 0.112), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.872 ± 0.003

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

Set DPS (verified): 48.6. Weights run: 3.2s. Verify run: 1.8s. 427 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.630 ± 0.012, crit=0.091 ± 0.002 per rating point (14 rating = 1%, 1.270 per %), hit=0.247 ± 0.008 per rating point (10 rating = 1%, 2.467 per %), spell_haste=-2.228 ± 0.170, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.806 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.3 spell_power points (1.35 DPS) | yes | Holy Shroud (2721, -0.14 DPS) [world_drop]; Enduring Cap (3020, -0.24 DPS) [world_drop]; Totemic Leather Helm (252456, -0.81 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.8 spell_power points (0.42 DPS) | yes | Crystal Starfire Medallion (5003, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Magician's Mantle (12998, -0.44 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.55 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Prelacy Cape (7004, -0.11 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.2 spell_power points (1.89 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Guardian Armor (4256, -0.21 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.99 DPS) | yes | Nightsky Wristbands (6407, -0.57 DPS) [world_drop]; Technician's Bracers (270042, -0.57 DPS) [quest]; Glowing Magical Bracelets (13106, -1.08 DPS, sim-verified) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.5 spell_power points (0.83 DPS) | yes | Town Clerk's Mittens (270029, -0.01 DPS) [quest]; Shilly Mitts (9609, -0.06 DPS) [quest]; Gloves of Insight (9698, -0.06 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.8 spell_power points (1.52 DPS) | yes | Belt of Arugal (6392, -0.32 DPS) [dungeon]; Crimson Silk Belt (7055, -0.37 DPS) [crafted]; Moss Cinch (6911, -1.20 DPS, sim-verified) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.08 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.20 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.54 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.4 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.21 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Acidic Walkers (9454, -1.22 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (1.01 DPS) | yes | Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; Azora's Will (4999, -0.66 DPS) [world_drop]; Minor Channeling Ring (1449, -1.73 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Azora's Will (4999, -0.31 DPS) [world_drop]; Minor Channeling Ring (1449, -0.49 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.76 DPS) [quest]; Manual Crowd Pummeler (9449, -7.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 427, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (dwarf, 5510331300001350-010000000000000000-0000000000000000)

Set DPS (verified): 71.7. Weights run: 3.6s. Verify run: 2.2s. 676 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.936 ± 0.025, crit=0.247 ± 0.008 per rating point (14 rating = 1%, 3.459 per %), hit=0.431 ± 0.019 per rating point (10 rating = 1%, 4.305 per %), spell_haste=-2.455 ± 0.273, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.692 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.9 spell_power points (2.58 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.31 DPS) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Big Voodoo Mask (8201, -1.89 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.5 spell_power points (1.56 DPS) | yes | Necklace of Calisea (1714, -0.86 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.96 DPS) [world]; Triune Amulet (7722, -1.56 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.2 spell_power points (2.07 DPS) | yes | Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest]; Green Silken Shoulders (7057, -1.25 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.7 spell_power points (1.15 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.34 DPS) [dungeon]; Icy Cloak (4327, -0.40 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.6 spell_power points (2.98 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of Power (7054, -0.26 DPS) [crafted]; Big Voodoo Robe (8200, -0.60 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.6 spell_power points (1.25 DPS) | yes | Turtle Scale Bracers (8198, -0.10 DPS) [crafted]; Arcane Runed Bracers (4744, -0.28 DPS) [quest]; Spidertank Oilrag (9448, -0.28 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.59 DPS) | yes | Red Mageweave Gloves (10018, -0.39 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.49 DPS) [crafted]; Dreamweave Gloves (10019, -1.41 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.0 spell_power points (2.27 DPS) | yes | Gilded Cord (254037, -0.60 DPS) [crafted]; Prefect's Belt (250559, -0.81 DPS) [crafted]; Skycaller's Leather Belt (252522, -1.46 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.2 spell_power points (2.72 DPS) | yes | Stormrider's Leather Pants (252502, -1.04 DPS) [crafted]; Guardian Pants (5962, -1.05 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.17 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.59 DPS) | yes | Skycaller's Mail Boots (252563, -0.48 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.92 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.02 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (1.58 DPS) | yes | Ring of Forlorn Spirits (2043, -0.72 DPS) [quest]; Reedknot Ring (9622, -0.83 DPS) [quest]; Minor Channeling Ring (1449, -0.84 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.7 spell_power points (1.15 DPS) | yes | Ring of Forlorn Spirits (2043, -0.29 DPS) [quest]; Reedknot Ring (9622, -0.40 DPS) [quest]; Minor Channeling Ring (1449, -0.41 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (71.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.20 DPS) [dungeon]; Spine of the Devourer (273040, -1.67 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 15.4 spell_power points (1.66 DPS) | yes | Thrash's Trash (276204, -0.15 DPS) [vendor]; Orb of Mystic Insight (249394, -0.30 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.54 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica

No-known-source sample (15 of 676, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 5510331300001350-030000000000000000-5030000000000000)

Set DPS (verified): 115.7. Weights run: 3.6s. Verify run: 2.4s. 879 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.801 ± 0.028, crit=0.298 ± 0.009 per rating point (14 rating = 1%, 4.174 per %), hit=0.491 ± 0.023 per rating point (10 rating = 1%, 4.913 per %), spell_haste=not significant (-1.235 ± 0.318), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.691 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.0 spell_power points (4.04 DPS) | yes | Soothsayer's Headdress (17740, -0.00 DPS) [dungeon]; Dreamweave Circlet (10041, -0.69 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.92 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (1.55 DPS) | yes | Horizon Choker (13085, -0.25 DPS) [world_drop]; Mindburst Medallion (11196, -0.30 DPS) [quest]; Triune Amulet (7722, -0.62 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 31.0 spell_power points (3.58 DPS) | yes | Rockshard Pauldrons (9411, -0.99 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -1.25 DPS) [crafted]; Rotgrip Mantle (17732, -1.91 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.8 spell_power points (2.17 DPS) | yes | Big Voodoo Cloak (8216, -0.76 DPS) [crafted]; Long Silken Cloak (4326, -1.01 DPS) [crafted]; Runecloth Cloak (13860, -1.32 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.0 spell_power points (4.04 DPS) | yes | Feathered Breastplate (8349, -0.92 DPS) [crafted]; Robe of the Magi (1716, -0.95 DPS) [world_drop]; Runecloth Tunic (13857, -1.06 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.0 spell_power points (2.08 DPS) | yes | Skycaller's Leather Bracers (252542, -0.28 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.28 DPS) [crafted]; Nethergeld Cuffs (254061, -0.62 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.09 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.09 DPS) [crafted]; Dreamweave Gloves (10019, -0.32 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.6 spell_power points (2.95 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.42 DPS) [dungeon]; Dawnspire Cord (12466, -0.51 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.6 spell_power points (2.72 DPS) | yes | Big Voodoo Pants (8202, -0.07 DPS) [crafted]; Turtle Scale Leggings (8185, -0.53 DPS) [crafted]; Wizardweave Leggings (14132, -0.53 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Skycaller's Leather Boots (252471, -0.02 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.02 DPS) [crafted]; Greaves of Withering Despair (22240, -0.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.61 DPS) | yes | Snake Hoop (6750, -0.46 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.69 DPS) [quest]; Reedknot Ring (9622, -0.81 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.50 DPS) | yes | Snake Hoop (6750, -0.35 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.58 DPS) [quest]; Reedknot Ring (9622, -0.69 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (115.7 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (115.7 DPS) | yes | Mark of the Chosen (17774, -2.14 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (115.7 DPS) | yes | Spellshifter Rod (9527, -0.09 DPS) [quest]; Kindling Stave (11750, -0.42 DPS) [dungeon]; Hammer of the Northern Wind (810, -25.29 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 5510331300001350-030000000000000000-5033520000000000)

Set DPS (verified): 134.3. Weights run: 3.5s. Verify run: 7.5s. 1826 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.867 ± 0.034, crit=0.361 ± 0.010 per rating point (14 rating = 1%, 5.052 per %), hit=0.635 ± 0.027 per rating point (10 rating = 1%, 6.346 per %), spell_haste=-1.861 ± 0.401, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.701 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (134.3 DPS) | yes | Black Dragonscale Helm (252605, -0.31 DPS) [crafted]; Magister's Crown (16686, -0.43 DPS) [dungeon]; Blue Dragonscale Helm (252604, -4.73 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.7 spell_power points (2.47 DPS) | yes | Glowing Eye of Mordresh (10769, -0.80 DPS) [dungeon]; Horizon Choker (13085, -1.02 DPS) [world_drop]; Opaline Medallion (16623, -1.12 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (134.3 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS) [dungeon]; Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Ironfeather Shoulders (15067, -4.43 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.2 spell_power points (2.29 DPS) | yes | Runecloth Cloak (13860, -0.39 DPS) [crafted]; Big Voodoo Cloak (8216, -0.76 DPS) [crafted]; Long Silken Cloak (4326, -1.06 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (134.3 DPS) | yes | Vest of Elements (16666, +0.00 DPS) [dungeon]; Acumen Robes (17775, +0.00 DPS) [quest]; Ironfeather Breastplate (15066, -4.25 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (134.3 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -6.42 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (134.3 DPS) | yes | Gloves of the Greatfather (17721, +0.00 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, +0.00 DPS) [crafted]; Storm Gauntlets (12632, -6.90 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (134.3 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -6.10 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.7 spell_power points (7.00 DPS) | yes | Red Dragonscale Leggings (252603, -1.30 DPS) [crafted]; Luminary Kilt (11823, -2.31 DPS) [dungeon]; Pristine Scorpid Leggings (252606, -2.46 DPS) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 30.4 spell_power points (3.63 DPS) | yes | Skycaller's Mail Sabatons (252577, -0.70 DPS) [crafted]; Earthen Silk Slippers (254013, -0.76 DPS) [crafted]; Skycaller's Leather Boots (252471, -1.29 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.1 spell_power points (2.87 DPS) | yes | Band of the Unicorn (7553, -1.32 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.63 DPS) [quest]; Snake Hoop (6750, -1.64 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.71 DPS) | yes | Band of the Unicorn (7553, -0.16 DPS) [world_drop]; Glowing Crystal Ring (18402, -0.47 DPS) [quest]; Snake Hoop (6750, -0.48 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (134.3 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (134.3 DPS) | yes | Mark of the Chosen (17774, -2.41 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (134.3 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -3.12 DPS) [world_drop] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (2.51 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Truesilver Conduit (249455, +0.00 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1826, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (dwarf, 5510331300001351-055000000000000000-0530020000000000)

Set DPS (verified): 462.4. Weights run: 2.1s. Verify run: 5.9s. 1826 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.500 ± 0.019, crit=0.369 ± 0.008 per rating point (14 rating = 1%, 5.167 per %), hit=0.683 ± 0.038 per rating point (10 rating = 1%, 6.831 per %), spell_haste=not significant (1.762 ± 0.464), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.645 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (462.4 DPS) | yes | Soothsayer's Headdress (17740, -1.55 DPS) [dungeon]; Red Mageweave Headband (10033, -2.22 DPS) [crafted]; Blue Dragonscale Helm (252604, -7.13 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.0 spell_power points (7.54 DPS) | yes | Glowing Eye of Mordresh (10769, -2.66 DPS) [dungeon]; Mindburst Medallion (11196, -3.55 DPS) [quest]; Triune Amulet (7722, -3.99 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (462.4 DPS) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Rotgrip Mantle (17732, -1.33 DPS) [dungeon]; Mantle of the Timbermaw (19050, -6.94 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.0 spell_power points (7.54 DPS) | yes | Runecloth Cloak (13860, -2.16 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -3.33 DPS) [crafted]; Nightfall Drape (12465, -3.55 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (462.4 DPS) | yes | Fernpulse Jerkin (15786, +0.00 DPS) [quest]; Acumen Robes (17775, +0.00 DPS) [quest]; Vest of Elements (16666, -6.46 DPS, sim-verified) [dungeon] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 16.0 spell_power points (7.10 DPS) | yes | Runic Leather Bracers (15092, -0.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.11 DPS) [crafted]; Skycaller's Mail Bracers (252571, -1.11 DPS) [crafted] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | 24.8 spell_power points (11.02 DPS) | yes | Gloves of the Greatfather (17721, +0.00 DPS, sim-verified) [crafted]; Earth Warder's Gloves (21318, -1.26 DPS) [quest]; Skycaller's Leather Gauntlets (252550, -1.92 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 22.0 spell_power points (9.76 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -1.33 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.77 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.2 spell_power points (23.15 DPS) | yes | Pristine Scorpid Leggings (252606, -6.90 DPS, sim-verified) [crafted]; Red Dragonscale Leggings (252603, -7.17 DPS) [crafted]; Luminary Kilt (11823, -8.95 DPS) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.0 spell_power points (11.54 DPS) | yes | Earthen Silk Slippers (254013, -0.89 DPS) [crafted]; Venomspew Footpads (275606, -1.77 DPS) [crafted]; Bileblister Boots (275612, -1.77 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.5 spell_power points (9.54 DPS) | yes | Philanthropist's Ring (281635, -3.99 DPS) [quest]; Snake Hoop (6750, -5.77 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.99 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.77 DPS) | yes | Philanthropist's Ring (281635, -0.22 DPS) [quest]; Snake Hoop (6750, -2.00 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.22 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (462.4 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (462.4 DPS) | yes | Mark of the Chosen (17774, -6.13 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (462.4 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -13.53 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.32 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Orb of the Forgotten Seer (7685, -3.58 DPS, sim-verified) [dungeon] |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (462.4 DPS) | yes | Totem of the Storm (272432, -2.06 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; hands: Storm Gauntlets; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Totem of Thunder

No-known-source sample (15 of 1826, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.0. Weights run: 2.9s. Verify run: 1.6s. 216 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.577 ± 0.009, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.031 per %), hit=0.188 ± 0.007 per rating point (10 rating = 1%, 1.884 per %), spell_haste=not significant (-0.026 ± 0.112), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.872 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.12 DPS) [crafted]; Totemic Leather Hood (252448, -0.31 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.2 spell_power points (0.85 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.52 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.75 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.02 DPS) [crafted]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Sanguine Cape (14376, -0.14 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.01 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.23 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.25 DPS) | yes | Mindthrust Bracers (1974, -0.01 DPS) [dungeon]; Owl Bracers (4796, -0.01 DPS) [vendor]; Featherbead Bracers (15452, -0.01 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.3 spell_power points (0.61 DPS) | yes | Gnoll Casting Gloves (892, -0.11 DPS) [world]; Pristine Gloves (253913, -0.13 DPS) [crafted]; Serpent Gloves (5970, -0.27 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.43 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.08 DPS) [crafted]; Totemic Leather Pants (252446, -0.10 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.43 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.3 spell_power points (0.78 DPS) | yes | Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.28 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.49 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.5 spell_power points (0.29 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.9 spell_power points (0.24 DPS) | yes | The 1 Ring (8350, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.39 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Twisted Chanter's Staff (890, -0.19 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.26 DPS, sim-verified) [quest]; Lesser Staff of the Spire (1300, -0.38 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection; 21568 Rune of Duty

### Band 30 (orc, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 48.6. Weights run: 3.2s. Verify run: 1.7s. 420 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.630 ± 0.012, crit=0.091 ± 0.002 per rating point (14 rating = 1%, 1.270 per %), hit=0.247 ± 0.008 per rating point (10 rating = 1%, 2.467 per %), spell_haste=-2.228 ± 0.170, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.806 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.3 spell_power points (1.35 DPS) | yes | Holy Shroud (2721, -0.14 DPS) [world_drop]; Enduring Cap (3020, -0.24 DPS) [world_drop]; Totemic Leather Helm (252456, -0.94 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.8 spell_power points (0.42 DPS) | yes | Crystal Starfire Medallion (5003, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Magician's Mantle (12998, -0.44 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.55 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.12 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.2 spell_power points (1.89 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Guardian Armor (4256, -0.21 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.99 DPS) | yes | Nightsky Wristbands (6407, -0.57 DPS) [world_drop]; Technician's Bracers (270042, -0.57 DPS) [quest]; Glowing Magical Bracelets (13106, -0.66 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.3 spell_power points (1.13 DPS) | yes | Jutebraid Gloves (10654, -0.12 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.30 DPS) [crafted]; Serpent Gloves (5970, -0.36 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.8 spell_power points (1.52 DPS) | yes | Warsong Sash (16975, -0.31 DPS) [quest]; Belt of Arugal (6392, -0.32 DPS) [dungeon]; Moss Cinch (6911, -0.64 DPS, sim-verified) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (48.6 DPS) | yes | Guardian Pants (5962, -0.08 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.20 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.49 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.4 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.21 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Acidic Walkers (9454, -1.28 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (1.01 DPS) | yes | Solid Gold Band (279845, -0.57 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; Azora's Will (4999, -0.66 DPS) [world_drop] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.66 DPS) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Azora's Will (4999, -0.31 DPS) [world_drop]; Solid Gold Band (279845, -0.71 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, -3.08 DPS) [dungeon]; Manual Crowd Pummeler (9449, -7.27 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 420, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (orc, 5510331300001350-010000000000000000-0000000000000000)

Set DPS (verified): 81.1. Weights run: 3.6s. Verify run: 2.2s. 658 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.936 ± 0.025, crit=0.247 ± 0.008 per rating point (14 rating = 1%, 3.459 per %), hit=0.431 ± 0.019 per rating point (10 rating = 1%, 4.305 per %), spell_haste=-2.455 ± 0.273, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.692 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.9 spell_power points (2.58 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.31 DPS) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Big Voodoo Mask (8201, -1.64 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.5 spell_power points (1.56 DPS) | yes | Necklace of Calisea (1714, -0.86 DPS) [dungeon]; Triune Amulet (7722, -0.92 DPS, sim-verified) [dungeon]; Archmagister's Faceted Pendant (273049, -0.96 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.2 spell_power points (2.07 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.7 spell_power points (1.15 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.34 DPS) [dungeon]; Icy Cloak (4327, -0.40 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.6 spell_power points (2.98 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of Power (7054, -0.26 DPS) [crafted]; Big Voodoo Robe (8200, -0.60 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.6 spell_power points (1.25 DPS) | yes | Radiant Silver Bracers (4545, -0.01 DPS) [quest]; Turtle Scale Bracers (8198, -0.10 DPS) [crafted]; Spidertank Oilrag (9448, -0.28 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.59 DPS) | yes | Red Mageweave Gloves (10018, -0.39 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.49 DPS) [crafted]; Dreamweave Gloves (10019, -1.08 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.0 spell_power points (2.27 DPS) | yes | Skycaller's Leather Belt (252522, -0.59 DPS) [crafted]; Gilded Cord (254037, -0.60 DPS) [crafted]; Prefect's Belt (250559, -0.81 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.2 spell_power points (2.72 DPS) | yes | Crimson Silk Pantaloons (7062, -0.67 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -1.04 DPS) [crafted]; Guardian Pants (5962, -1.05 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.59 DPS) | yes | Skycaller's Leather Shoes (252532, -0.48 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.48 DPS) [crafted]; Mender's Leather Shoes (252533, -1.02 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (1.58 DPS) | yes | Reedknot Ring (9622, -0.83 DPS) [quest]; Ogremind Ring (1993, -0.88 DPS) [world_drop]; Voodoo Band (1996, -0.88 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.7 spell_power points (1.15 DPS) | yes | Reedknot Ring (9622, -0.40 DPS) [quest]; Ogremind Ring (1993, -0.45 DPS) [world_drop]; Voodoo Band (1996, -0.45 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.20 DPS) [dungeon]; Spine of the Devourer (273040, -1.67 DPS) [world] |
| off_hand | Thrash's Trash (276204) | Friz Frazzlespark [vendor] | sim-verified (81.1 DPS) | yes | Orb of Mystic Insight (249394, -0.15 DPS) [crafted]; Anchorhold Buckler (15865, -0.37 DPS) [quest]; Orb of the Forgotten Seer (7685, -11.45 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Thrash's Trash

No-known-source sample (15 of 658, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (orc, 5510331300001350-030000000000000000-5030000000000000)

Set DPS (verified): 117.3. Weights run: 3.6s. Verify run: 2.5s. 849 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.801 ± 0.028, crit=0.298 ± 0.009 per rating point (14 rating = 1%, 4.174 per %), hit=0.491 ± 0.023 per rating point (10 rating = 1%, 4.913 per %), spell_haste=not significant (-1.235 ± 0.318), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.691 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.0 spell_power points (4.04 DPS) | yes | Soothsayer's Headdress (17740, -0.00 DPS) [dungeon]; Dreamweave Circlet (10041, -0.69 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.92 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (1.55 DPS) | yes | Horizon Choker (13085, -0.25 DPS) [world_drop]; Mindburst Medallion (11196, -0.30 DPS) [quest]; Triune Amulet (7722, -0.62 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 31.0 spell_power points (3.58 DPS) | yes | Rockshard Pauldrons (9411, -0.99 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -1.25 DPS) [crafted]; Rotgrip Mantle (17732, -2.42 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.2 spell_power points (2.21 DPS) | yes | Spritecaster Cape (11623, -0.05 DPS) [dungeon]; Runecloth Cloak (13860, -0.44 DPS) [crafted]; Big Voodoo Cloak (8216, -0.81 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.0 spell_power points (4.04 DPS) | yes | Feathered Breastplate (8349, -0.92 DPS) [crafted]; Robe of the Magi (1716, -0.95 DPS) [world_drop]; Stone Guard's Pulsing Breastplate (220844, -1.56 DPS, sim-verified) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.0 spell_power points (2.08 DPS) | yes | Skycaller's Leather Bracers (252542, -0.28 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.28 DPS) [crafted]; Nethergeld Cuffs (254061, -0.62 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.09 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.09 DPS) [crafted]; Dreamweave Gloves (10019, -0.32 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.6 spell_power points (2.95 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.42 DPS) [dungeon]; Dawnspire Cord (12466, -0.51 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | PvP rank 10 · Stone Guard · Horde [vendor] | 32.5 spell_power points (3.75 DPS) | yes | Big Voodoo Pants (8202, -1.10 DPS) [crafted]; Turtle Scale Leggings (8185, -1.56 DPS) [crafted]; Red Mageweave Pants (10009, -2.59 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Skycaller's Leather Boots (252471, -0.02 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.02 DPS) [crafted]; Greaves of Withering Despair (22240, -0.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.61 DPS) | yes | Snake Hoop (6750, -0.46 DPS) [quest]; Reedknot Ring (9622, -0.81 DPS) [quest]; Ring of Power Regulation (273030, -0.92 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.50 DPS) | yes | Snake Hoop (6750, -0.35 DPS) [quest]; Reedknot Ring (9622, -0.69 DPS) [quest]; Ring of Power Regulation (273030, -0.81 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (117.3 DPS) | yes | Rune of the Guard Captain (19120, -0.11 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (117.3 DPS) | yes | Rune of the Guard Captain (19120, -1.36 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (117.3 DPS) | yes | Spellshifter Rod (9527, -0.09 DPS) [quest]; Kindling Stave (11750, -0.42 DPS) [dungeon]; Hammer of the Northern Wind (810, -25.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 849, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (orc, 5510331300001350-030000000000000000-5033520000000000)

Set DPS (verified): 134.1. Weights run: 3.5s. Verify run: 7.5s. 1782 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.867 ± 0.034, crit=0.361 ± 0.010 per rating point (14 rating = 1%, 5.052 per %), hit=0.635 ± 0.027 per rating point (10 rating = 1%, 6.346 per %), spell_haste=-1.861 ± 0.401, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.701 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (134.1 DPS) | yes | Warlord's Mail Helm (231663, +0.00 DPS) [pvp]; Champion's Mail Helm (227162, -0.20 DPS) [pvp]; Blue Dragonscale Helm (252604, -6.07 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.7 spell_power points (2.47 DPS) | yes | Glowing Eye of Mordresh (10769, -0.80 DPS) [dungeon]; Horizon Choker (13085, -1.02 DPS) [world_drop]; Mindburst Medallion (11196, -1.13 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (134.1 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Ironfeather Shoulders (15067, -4.99 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.8 spell_power points (2.36 DPS) | yes | Spritecaster Cape (11623, -0.07 DPS) [dungeon]; Runecloth Cloak (13860, -0.46 DPS) [crafted]; Big Voodoo Cloak (8216, -0.84 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (134.1 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -5.77 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (134.1 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -7.00 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (134.1 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, +0.00 DPS) [vendor]; Storm Gauntlets (12632, -6.94 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (134.1 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -7.31 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.7 spell_power points (7.00 DPS) | yes | General's Mail Leggings (231664, -0.57 DPS) [pvp]; Red Dragonscale Leggings (252603, -1.48 DPS, sim-verified) [crafted]; Legionnaire's Mail Leggings (227161, -2.13 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 30.4 spell_power points (3.63 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; General's Mail Boots (231667, -0.30 DPS) [pvp]; Skycaller's Leather Boots (252471, -1.71 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.1 spell_power points (2.87 DPS) | yes | Band of the Unicorn (7553, -1.32 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.63 DPS) [quest]; Snake Hoop (6750, -1.64 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.71 DPS) | yes | Band of the Unicorn (7553, -0.16 DPS) [world_drop]; Glowing Crystal Ring (18402, -0.47 DPS) [quest]; Snake Hoop (6750, -0.48 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (134.1 DPS) | yes | Rune of the Guard Captain (19120, -0.15 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (134.1 DPS) | yes | Rune of the Guard Captain (19120, -1.46 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (134.1 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -3.12 DPS) [world_drop] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (2.51 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Truesilver Conduit (249455, -1.59 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1782, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (orc, 5510331300001351-055000000000000000-0530020000000000)

Set DPS (verified): 461.4. Weights run: 2.1s. Verify run: 5.9s. 1782 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.500 ± 0.019, crit=0.369 ± 0.008 per rating point (14 rating = 1%, 5.167 per %), hit=0.683 ± 0.038 per rating point (10 rating = 1%, 6.831 per %), spell_haste=not significant (1.762 ± 0.464), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.645 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (461.4 DPS) | yes | Warlord's Mail Helm (231663, +0.00 DPS) [pvp]; Champion's Mail Helm (227162, -0.37 DPS) [pvp]; Blue Dragonscale Helm (252604, -7.52 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.0 spell_power points (7.54 DPS) | yes | Glowing Eye of Mordresh (10769, -2.66 DPS) [dungeon]; Mindburst Medallion (11196, -3.55 DPS) [quest]; Triune Amulet (7722, -3.99 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (461.4 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -7.21 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.0 spell_power points (7.54 DPS) | yes | Deep Woodlands Cloak (19121, -0.22 DPS) [quest]; Runecloth Cloak (13860, -1.77 DPS) [crafted]; Big Voodoo Cloak (8216, -3.33 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (461.4 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Vest of Elements (16666, -8.25 DPS, sim-verified) [dungeon] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 16.0 spell_power points (7.10 DPS) | yes | Runic Leather Bracers (15092, -0.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.11 DPS) [crafted]; Skycaller's Mail Bracers (252571, -1.11 DPS) [crafted] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | 24.8 spell_power points (11.02 DPS) | yes | Gloves of the Greatfather (17721, +0.00 DPS, sim-verified) [crafted]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, +0.00 DPS) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 22.0 spell_power points (9.76 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -1.33 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.77 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.2 spell_power points (23.15 DPS) | yes | General's Mail Leggings (231664, -3.11 DPS) [pvp]; Red Dragonscale Leggings (252603, -7.17 DPS) [crafted]; Pristine Scorpid Leggings (252606, -7.83 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.0 spell_power points (11.54 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Earthen Silk Slippers (254013, -0.89 DPS) [crafted]; Venomspew Footpads (275606, -1.77 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.5 spell_power points (9.54 DPS) | yes | Philanthropist's Ring (281635, -3.99 DPS) [quest]; Snake Hoop (6750, -5.77 DPS) [quest]; Reedknot Ring (9622, -6.43 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.77 DPS) | yes | Philanthropist's Ring (281635, -0.22 DPS) [quest]; Snake Hoop (6750, -2.00 DPS) [quest]; Reedknot Ring (9622, -2.66 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (461.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (461.4 DPS) | yes | Frozen Heart of the Mountain (249469, -2.17 DPS, sim-verified) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (461.4 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -13.53 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.32 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Orb of the Forgotten Seer (7685, -3.59 DPS, sim-verified) [dungeon] |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (461.4 DPS) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Spritecaster Cape; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; hands: Storm Gauntlets; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; trinket1: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Totem of Thunder

No-known-source sample (15 of 1782, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

