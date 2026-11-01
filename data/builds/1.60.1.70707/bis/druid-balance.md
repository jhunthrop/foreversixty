# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 34.4. Weights run: 2.6s. Verify run: 1.2s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.735 ± 0.007, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.444 per %), hit=0.095 ± 0.007 per rating point (10 rating = 1%, 0.952 per %), spell_haste=not significant (-0.033 ± 0.098), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.025 ± 0.000, arcane_power=0.975 ± 0.003, spirit=not significant (0.084 ± 0.041), mp5=0.356 ± 0.062

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 10.7 spell_power points (1.65 DPS) | yes | Totemic Leather Hood (252448, -0.42 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.73 DPS) [crafted]; Pristine Circlet (253949, -0.73 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.2 spell_power points (0.03 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 spell_power points (1.79 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.01 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.17 DPS) [crafted]; Slime-encrusted Pads (6461, -1.62 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.3 spell_power points (0.67 DPS) | yes | Black Whelp Cloak (7283, -0.21 DPS) [crafted]; Sanguine Cape (14376, -0.21 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.35 DPS, sim-verified) [crafted] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 12.7 spell_power points (1.95 DPS) | yes | Stormrider's Leather Armor (252492, -0.12 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.58 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.58 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 4.7 spell_power points (0.73 DPS) | yes | Bright Bracers (3647, -0.28 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.39 DPS) [dungeon]; Owl Bracers (4796, -0.78 DPS, sim-verified) [vendor] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.0 spell_power points (1.39 DPS) | yes | Stormrider's Leather Gloves (252498, -0.17 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS) [dungeon]; Windfelt Gloves (5630, -0.44 DPS) [quest] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 9.8 spell_power points (1.50 DPS) | yes | Pristine Sash (253925, -0.44 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.64 DPS, sim-verified) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 15.5 spell_power points (2.38 DPS) | yes | Stormrider's Leather Pants (252502, -0.11 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.57 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.73 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (1.53 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Pearly Boots (253897, -0.29 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.35 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (1.00 DPS) | yes | Volcanic Rock Ring (12053, -0.66 DPS) [world_drop]; Furlfeather Ring (286745, -0.84 DPS) [world]; The 1 Ring (8350, -0.87 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.4 spell_power points (0.68 DPS) | yes | Furlfeather Ring (286745, -0.52 DPS) [world]; The 1 Ring (8350, -0.55 DPS) [world]; Volcanic Rock Ring (12053, -0.74 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.23 DPS) | yes | Twisted Chanter's Staff (890, -0.10 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.55 DPS) [world]; Staff of Westfall (2042, -0.59 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde

### Band 30 (night-elf, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 62.1. Weights run: 2.8s. Verify run: 1.6s. 414 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.522 ± 0.008, crit=0.049 ± 0.002 per rating point (14 rating = 1%, 0.691 per %), hit=0.153 ± 0.009 per rating point (10 rating = 1%, 1.534 per %), spell_haste=not significant (-0.128 ± 0.117), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.322 ± 0.001, arcane_power=0.678 ± 0.002, spirit=not significant (0.234 ± 0.067), mp5=0.696 ± 0.083

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | sim-verified (62.1 DPS) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Enchanter's Cowl (4322, -0.16 DPS) [crafted]; Enduring Cap (3020, -0.42 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.1 spell_power points (0.66 DPS) | yes | Kaleidoscope Chain (13084, -0.02 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.07 DPS) [world_drop]; Pendant of Myzrael (4614, -0.36 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.7 spell_power points (2.87 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS) [dungeon]; Magician's Mantle (12998, -0.84 DPS) [world_drop]; Invoker's Mantle (215365, -0.86 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.05 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.01 DPS) [crafted]; Glowing Thresher Cape (6901, -0.03 DPS) [dungeon] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 17.1 spell_power points (3.59 DPS) | yes | Green Silk Armor (7065, -0.28 DPS) [crafted]; Stormrider's Leather Tunic (252510, -0.59 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.60 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.1 spell_power points (2.32 DPS) | yes | Mindthrust Bracers (1974, -1.34 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.45 DPS) [world_drop]; Nightsky Wristbands (6407, -1.52 DPS) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | sim-verified (62.1 DPS) | yes | Ephemeral Grips (273034, +0.00 DPS) [world]; Serpent Gloves (5970, -0.02 DPS) [dungeon]; Gloves of Insight (9698, -0.02 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.1 spell_power points (2.95 DPS) | yes | Moss Cinch (6911, -0.43 DPS) [dungeon]; Belt of Arugal (6392, -0.49 DPS) [dungeon]; Mender's Leather Belt (252523, -0.74 DPS) [crafted] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | sim-verified (62.1 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS) [dungeon]; Stormrider's Leather Pants (252502, +0.00 DPS) [crafted]; Pearly Leggings (253995, +0.00 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (62.1 DPS) | yes | Golden Slippers (254009, +0.00 DPS) [crafted]; Acidic Walkers (9454, -0.31 DPS) [dungeon]; Spidersilk Boots (4320, -0.52 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.6 spell_power points (1.80 DPS) | yes | Minor Channeling Ring (1449, -0.54 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Runemender's Seal (273031, -0.83 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.1 spell_power points (1.69 DPS) | yes | Minor Channeling Ring (1449, -0.43 DPS) [quest]; Electrocutioner Lagnut (9447, -0.72 DPS) [dungeon]; Runemender's Seal (273031, -0.72 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (62.1 DPS) | yes | Bloodscalp Channeling Staff (1998, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, +0.00 DPS) [dungeon]; Mechanic's Pipehammer (9604, +0.00 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Stormrider's Leather Gloves; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (night-elf, 5232221115100141-00000000000000000000-0000000000000000)

Set DPS (verified): 122.7. Weights run: 2.9s. Verify run: 1.8s. 549 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.344 ± 0.009, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.975 per %), hit=0.198 ± 0.013 per rating point (10 rating = 1%, 1.980 per %), spell_haste=not significant (0.415 ± 0.115), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.002, spirit=not significant (0.255 ± 0.098), mp5=0.515 ± 0.098

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.43 DPS) [world]; Big Voodoo Mask (8201, -0.59 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.21 DPS, sim-verified) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.8 spell_power points (2.46 DPS) | yes | Glowing Eye of Mordresh (10769, -0.01 DPS) [dungeon]; Necklace of Calisea (1714, -1.41 DPS) [dungeon]; Amberglow Talisman (10824, -1.82 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (3.04 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Inquisitor's Shawl (19507, -0.16 DPS) [dungeon]; Mantle of Doan (7712, -0.25 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 8.5 spell_power points (2.15 DPS) | yes | Long Silken Cloak (4326, -0.21 DPS) [crafted]; Guardian Cloak (5965, -0.21 DPS) [crafted]; Glowing Thresher Cape (6901, -0.88 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.92 DPS) [world_drop]; Robe of Power (7054, -1.13 DPS) [crafted]; Robe of the Magi (1716, -1.25 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.5 spell_power points (2.65 DPS) | yes | Guardian Leather Bracers (4260, -0.63 DPS) [crafted]; Green Whelp Bracers (7386, -0.86 DPS) [crafted]; Arcane Runed Bracers (4744, -1.16 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 28.9 spell_power points (7.26 DPS) | yes | Dreamweave Gloves (10019, -1.94 DPS) [crafted]; Black Mageweave Gloves (10003, -3.49 DPS) [crafted]; Red Mageweave Gloves (10018, -3.63 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 14.0 spell_power points (3.53 DPS) | yes | Skycaller's Leather Belt (252522, -0.24 DPS) [crafted]; Gilded Cord (254037, -0.44 DPS) [crafted]; Deathmage Sash (10771, -0.47 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.1 spell_power points (4.56 DPS) | yes | Gaze Dreamer Pants (6903, -0.96 DPS) [dungeon]; Dark Ritual Leggings (270031, -1.04 DPS) [quest]; Warchief Kilt (7760, -1.13 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.8 spell_power points (6.49 DPS) | yes | Skycaller's Leather Shoes (252532, -1.60 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -3.42 DPS) [crafted]; Gilded Slippers (254001, -3.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (2.95 DPS) | yes | Snake Hoop (6750, -1.01 DPS) [quest]; Ring of Power Regulation (273030, -1.05 DPS) [dungeon]; Reedknot Ring (9622, -1.19 DPS) [quest] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (2.01 DPS) | yes | Snake Hoop (6750, -0.07 DPS) [quest]; Ring of Power Regulation (273030, -0.11 DPS) [dungeon]; Reedknot Ring (9622, -0.25 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodscalp Channeling Staff (1998, -0.80 DPS) [world]; Rhahk'Zor's Hammer (5187, -1.11 DPS) [dungeon]; Windweaver Staff (7757, -16.93 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Papal Fez; neck: Triune Amulet; back: Icy Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 549, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 5232221115100151-00000000000000000000-5400000000000000)

Set DPS (verified): 149.1. Weights run: 3.1s. Verify run: 1.8s. 719 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.477 ± 0.013, crit=0.104 ± 0.004 per rating point (14 rating = 1%, 1.460 per %), hit=0.303 ± 0.020 per rating point (10 rating = 1%, 3.032 per %), spell_haste=not significant (0.348 ± 0.161), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.508 ± 0.001, arcane_power=0.492 ± 0.002, spirit=not significant (0.381 ± 0.152), mp5=not significant (0.609 ± 0.154)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 33.2 spell_power points (7.86 DPS) | yes | Dreamweave Circlet (10041, -0.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.10 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.47 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.8 spell_power points (2.56 DPS) | yes | Mindburst Medallion (11196, -0.01 DPS) [quest]; Triune Amulet (7722, -0.04 DPS) [dungeon]; Horizon Choker (13085, -0.62 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 27.6 spell_power points (6.53 DPS) | yes | Rotgrip Mantle (17732, -1.42 DPS) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.84 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.90 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.1 spell_power points (4.54 DPS) | yes | Runecloth Cloak (13860, -2.02 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -2.33 DPS) [crafted]; Icy Cloak (4327, -2.34 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 32.9 spell_power points (7.80 DPS) | yes | Flamestrider Robes (11747, -1.16 DPS) [dungeon]; Dreamweave Vest (10021, -1.25 DPS) [crafted]; Acumen Robes (17775, -1.40 DPS, sim-verified) [quest] |
| wrist | Skycaller's Leather Bracers (252542) | Leatherworking [crafted] | 15.2 spell_power points (3.61 DPS) | yes | Runic Leather Bracers (15092, -0.11 DPS) [crafted]; Nethergeld Cuffs (254061, -0.71 DPS) [crafted]; Nethershine Cuffs (254069, -0.85 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 29.1 spell_power points (6.89 DPS) | yes | Skycaller's Leather Gauntlets (252550, -1.54 DPS) [crafted]; Dreamweave Gloves (10019, -1.54 DPS) [crafted]; Runecloth Gloves (13863, -2.22 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.7 spell_power points (5.15 DPS) | yes | Satyrmane Sash (17755, -0.70 DPS) [dungeon]; Ghostweave Cord (254073, -0.84 DPS) [crafted]; Tooled Leather Belt (252544, -1.05 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 22.2 spell_power points (5.26 DPS) | yes | Knight's Crackling Leather Leggings (220864, +0.00 DPS, sim-verified) [vendor]; Jinxed Hoodoo Kilt (9474, -0.52 DPS) [dungeon]; Big Voodoo Pants (8202, -0.58 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.7 spell_power points (6.32 DPS) | yes | Skycaller's Leather Boots (252471, -1.52 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -1.86 DPS) [vendor]; Skycaller's Leather Shoes (252532, -2.00 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.08 DPS) | yes | Snake Hoop (6750, -1.09 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.18 DPS) [quest]; Ring of Power Regulation (273030, -1.23 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (2.93 DPS) | yes | Snake Hoop (6750, -0.95 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.04 DPS) [quest]; Ring of Power Regulation (273030, -1.08 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (149.1 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (149.1 DPS) | yes | Ankh of Life (1713, -0.34 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.52 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.70 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (149.1 DPS) | yes | Kindling Stave (11750, -0.22 DPS) [dungeon]; Spellshifter Rod (9527, -0.74 DPS) [quest]; Zum'rah's Vexing Cane (18082, -2.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Glowing Eye of Mordresh; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Skycaller's Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Senior Designer's Pantaloons; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar

No-known-source sample (15 of 719, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 5232221115100151-00000000000000000000-5533030000000000)

Set DPS (verified): 251.6. Weights run: 2.8s. Verify run: 4.5s. 1597 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.531 ± 0.017, crit=0.130 ± 0.005 per rating point (14 rating = 1%, 1.825 per %), hit=0.397 ± 0.027 per rating point (10 rating = 1%, 3.970 per %), spell_haste=0.999 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.002, spirit=not significant (0.576 ± 0.219), mp5=not significant (0.863 ± 0.230)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 42.6 spell_power points (11.81 DPS) | yes | Field Marshal's Dragonhide Helm (231695, -0.87 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -1.40 DPS) [vendor]; Soothsayer's Headdress (17740, -3.68 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.3 spell_power points (4.80 DPS) | yes | Triune Amulet (7722, -1.46 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.68 DPS) [dungeon]; Mindburst Medallion (11196, -1.97 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 30.2 spell_power points (8.38 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Pauldrons (231705, -0.79 DPS) [vendor]; Argent Shoulders (19059, -10.42 DPS, sim-verified) [crafted] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 23.1 spell_power points (6.39 DPS) | yes | Spritecaster Cape (11623, -0.67 DPS) [dungeon]; Runecloth Cloak (13860, -2.72 DPS) [crafted]; Icy Cloak (4327, -3.50 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 43.5 spell_power points (12.06 DPS) | yes | Field Marshal's Dragonhide Armor (231696, -0.93 DPS) [vendor]; Knight-Captain's Dragonhide Armor (227195, -1.95 DPS) [vendor]; Feathered Breastplate (8349, -10.68 DPS, sim-verified) [crafted] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (251.6 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, -5.20 DPS, sim-verified) [crafted] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-verified (251.6 DPS) | yes | Knight-Lieutenant's Dragonhide Gloves (227193, +0.00 DPS) [pvp]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Gloves of the Greatfather (17721, -8.30 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (251.6 DPS) | yes | Dreadmist Belt (16702, +0.00 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -4.86 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.4 spell_power points (13.69 DPS) | yes | Marshal's Dragonhide Pants (231697, -2.04 DPS) [vendor]; Luminary Kilt (11823, -3.37 DPS) [dungeon]; Ghostweave Pants (14144, -5.35 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.0 spell_power points (7.77 DPS) | yes | Knight-Lieutenant's Dragonhide Boots (227194, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Venomspew Footpads (275606, +0.00 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.2 spell_power points (6.98 DPS) | yes | Philanthropist's Ring (281635, -3.47 DPS) [quest]; Band of the Hierophant (13096, -3.75 DPS) [world_drop]; Eye of Adaegus (5266, -4.18 DPS) [world_drop] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.60 DPS) | yes | Philanthropist's Ring (281635, -0.10 DPS) [quest]; Band of the Hierophant (13096, -0.38 DPS) [world_drop]; Eye of Adaegus (5266, -0.81 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (251.6 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (251.6 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (251.6 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -12.64 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-verified (251.6 DPS) | yes | Talons of Wrath (249441, +0.00 DPS) [crafted] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Cloak of Earth and Sky; chest: Ironfeather Breastplate; wrist: Magister's Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd; ranged: Idol of the Moon

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 534.8. Weights run: 2.0s. Verify run: 3.4s. 1597 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.178 ± 0.011, crit=0.271 ± 0.011 per rating point (14 rating = 1%, 3.792 per %), hit=0.502 ± 0.030 per rating point (10 rating = 1%, 5.022 per %), spell_haste=2.953 ± 0.240, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.329 ± 0.000, arcane_power=0.671 ± 0.001, spirit=not significant (0.077 ± 0.052), mp5=not significant (0.218 ± 0.083)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 29.3 spell_power points (17.04 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.35 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.33 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.8 spell_power points (8.02 DPS) | yes | Triune Amulet (7722, -3.04 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -3.12 DPS) [dungeon]; Mindburst Medallion (11196, -3.68 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.5 spell_power points (14.86 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.42 DPS) [crafted]; Lieutenant Commander's Dragonhide Spaulders (227190, -2.43 DPS) [vendor] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 16.1 spell_power points (9.36 DPS) | yes | Spritecaster Cape (11623, -0.32 DPS) [dungeon]; Runecloth Cloak (13860, -3.30 DPS) [crafted]; Nightfall Drape (12465, -4.12 DPS) [dungeon] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | 27.4 spell_power points (15.95 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -3.00 DPS, sim-verified) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.8 spell_power points (7.43 DPS) | yes | Runic Leather Bracers (15092, -0.58 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.67 DPS) [crafted]; Nethershine Cuffs (254069, -1.41 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 27.3 spell_power points (15.88 DPS) | yes | Marshal's Dragonhide Gloves (231700, -1.48 DPS) [vendor]; Earth Warder's Gloves (21318, -3.17 DPS, sim-verified) [quest]; Knight-Lieutenant's Dragonhide Gloves (227193, -4.19 DPS) [pvp] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 18.1 spell_power points (10.55 DPS) | yes | Satyrmane Sash (17755, -1.37 DPS) [dungeon]; Ghostweave Cord (254073, -1.91 DPS) [crafted]; Mana-infused Cord (279267, -1.93 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 45.0 spell_power points (26.18 DPS) | yes | Marshal's Dragonhide Pants (231697, -6.36 DPS) [vendor]; Knight-Captain's Dragonhide Pants (227191, -9.63 DPS) [vendor]; Ghostweave Pants (14144, -12.82 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.5 spell_power points (14.28 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.07 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.85 DPS) [vendor]; Venomspew Footpads (275606, -1.12 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (11.47 DPS) | yes | Philanthropist's Ring (281635, -5.13 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.81 DPS) [quest]; Reedknot Ring (9622, -7.39 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.56 DPS) | yes | Philanthropist's Ring (281635, -1.23 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.91 DPS) [quest]; Reedknot Ring (9622, -3.49 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (534.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -2.09 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.18 DPS) [quest] |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (534.8 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -15.17 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-verified (534.8 DPS) | yes | Idol of the Dream (220606, +0.00 DPS) [dungeon] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Cloak of Earth and Sky; chest: Fernpulse Jerkin; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; main_hand: Hand of Edward the Odd; ranged: Idol of the Moon

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 32.5. Weights run: 2.6s. Verify run: 1.3s. 226 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.735 ± 0.007, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.444 per %), hit=0.095 ± 0.007 per rating point (10 rating = 1%, 0.952 per %), spell_haste=not significant (-0.033 ± 0.098), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.025 ± 0.000, arcane_power=0.975 ± 0.003, spirit=not significant (0.084 ± 0.041), mp5=0.356 ± 0.062

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 10.7 spell_power points (1.65 DPS) | yes | Totemic Leather Hood (252448, -0.41 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.73 DPS) [crafted]; Pristine Circlet (253949, -0.73 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.3 spell_power points (0.04 DPS) | yes | Scholarly Pendant (277203, -0.01 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 spell_power points (1.79 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.14 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.17 DPS) [crafted]; Slime-encrusted Pads (6461, -1.62 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.3 spell_power points (0.67 DPS) | yes | Pearl-clasped Cloak (5542, -0.02 DPS) [crafted]; Black Whelp Cloak (7283, -0.21 DPS) [crafted]; Sanguine Cape (14376, -0.21 DPS) [world_drop] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 12.7 spell_power points (1.95 DPS) | yes | Stormrider's Leather Armor (252492, -0.12 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.58 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.58 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 4.7 spell_power points (0.73 DPS) | yes | Tabitha's Cuffs (251486, -0.27 DPS) [quest]; Bright Bracers (3647, -0.28 DPS) [world_drop]; Owl Bracers (4796, -0.56 DPS, sim-verified) [vendor] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.0 spell_power points (1.39 DPS) | yes | Stormrider's Leather Gloves (252498, -0.17 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS) [dungeon]; Pristine Gloves (253913, -0.44 DPS) [crafted] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 9.8 spell_power points (1.50 DPS) | yes | Pristine Sash (253925, -0.44 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.63 DPS, sim-verified) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 15.5 spell_power points (2.38 DPS) | yes | Stormrider's Leather Pants (252502, -0.11 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.57 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.73 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (1.53 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Pearly Boots (253897, -0.29 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.35 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.4 spell_power points (0.68 DPS) | yes | Volcanic Rock Ring (12053, -0.34 DPS) [world_drop]; Furlfeather Ring (286745, -0.52 DPS) [world]; The 1 Ring (8350, -0.55 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 3.7 spell_power points (0.57 DPS) | yes | Volcanic Rock Ring (12053, -0.33 DPS, sim-verified) [world_drop]; Furlfeather Ring (286745, -0.41 DPS) [world]; The 1 Ring (8350, -0.44 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.23 DPS) | yes | Twisted Chanter's Staff (890, -0.10 DPS) [world_drop]; Crescent Staff (6505, -0.44 DPS) [quest]; Lesser Staff of the Spire (1300, -0.55 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 226, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (tauren, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 61.9. Weights run: 2.8s. Verify run: 1.4s. 414 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.522 ± 0.008, crit=0.049 ± 0.002 per rating point (14 rating = 1%, 0.691 per %), hit=0.153 ± 0.009 per rating point (10 rating = 1%, 1.534 per %), spell_haste=not significant (-0.128 ± 0.117), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.322 ± 0.001, arcane_power=0.678 ± 0.002, spirit=not significant (0.234 ± 0.067), mp5=0.696 ± 0.083

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | sim-verified (61.9 DPS) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Enchanter's Cowl (4322, -0.16 DPS) [crafted]; Enduring Cap (3020, -0.42 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.1 spell_power points (0.66 DPS) | yes | Kaleidoscope Chain (13084, -0.02 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.07 DPS) [world_drop]; Pendant of Myzrael (4614, -0.36 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.7 spell_power points (2.87 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS) [dungeon]; Magician's Mantle (12998, -0.84 DPS) [world_drop]; Invoker's Mantle (215365, -0.86 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | sim-verified (61.9 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.01 DPS) [crafted]; Glowing Thresher Cape (6901, -0.03 DPS) [dungeon] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 17.1 spell_power points (3.59 DPS) | yes | Green Silk Armor (7065, -0.28 DPS) [crafted]; Stormrider's Leather Tunic (252510, -0.59 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.60 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.1 spell_power points (2.32 DPS) | yes | Mindthrust Bracers (1974, -1.34 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.45 DPS) [world_drop]; Nightsky Wristbands (6407, -1.52 DPS) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.0 spell_power points (2.10 DPS) | yes | Ephemeral Grips (273034, -0.06 DPS) [world]; Jutebraid Gloves (10654, -0.30 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.62 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.1 spell_power points (2.95 DPS) | yes | Moss Cinch (6911, -0.43 DPS) [dungeon]; Belt of Arugal (6392, -0.49 DPS) [dungeon]; Warsong Sash (16975, -0.64 DPS) [quest] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (61.9 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS) [dungeon]; Pearly Leggings (253995, +0.00 DPS) [crafted]; Dark Ritual Leggings (270031, -0.01 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (61.9 DPS) | yes | Golden Slippers (254009, +0.00 DPS) [crafted]; Acidic Walkers (9454, -0.31 DPS) [dungeon]; Spidersilk Boots (4320, -0.52 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.6 spell_power points (1.80 DPS) | yes | Solid Gold Band (279845, -0.72 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Runemender's Seal (273031, -0.83 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.1 spell_power points (1.69 DPS) | yes | Solid Gold Band (279845, -0.61 DPS) [quest]; Electrocutioner Lagnut (9447, -0.72 DPS) [dungeon]; Runemender's Seal (273031, -0.72 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (61.9 DPS) | yes | Bloodscalp Channeling Staff (1998, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, +0.00 DPS) [dungeon]; Glimmering Staff (249392, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 5232221115100141-00000000000000000000-0000000000000000)

Set DPS (verified): 120.1. Weights run: 2.9s. Verify run: 1.7s. 547 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.344 ± 0.009, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.975 per %), hit=0.198 ± 0.013 per rating point (10 rating = 1%, 1.980 per %), spell_haste=not significant (0.415 ± 0.115), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.002, spirit=not significant (0.255 ± 0.098), mp5=0.515 ± 0.098

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (120.1 DPS) | yes | White Drakeskin Cap (4543, -0.15 DPS) [quest]; Augural Shroud (2620, -0.43 DPS) [world]; Spellpower Goggles Xtreme (10502, -1.83 DPS, sim-verified) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.8 spell_power points (2.46 DPS) | yes | Glowing Eye of Mordresh (10769, -0.01 DPS) [dungeon]; Necklace of Calisea (1714, -1.41 DPS) [dungeon]; Amberglow Talisman (10824, -1.82 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (3.04 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Inquisitor's Shawl (19507, -0.16 DPS) [dungeon]; Mantle of Doan (7712, -0.25 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 8.5 spell_power points (2.15 DPS) | yes | Long Silken Cloak (4326, -0.21 DPS) [crafted]; Guardian Cloak (5965, -0.21 DPS) [crafted]; Windsong Drape (15468, -0.63 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.3 spell_power points (6.38 DPS) | yes | Dreamweave Vest (10021, -0.17 DPS) [crafted]; Elemental Raiment (9434, -1.09 DPS) [world_drop]; Robe of Power (7054, -1.30 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.5 spell_power points (2.65 DPS) | yes | Green Whelp Bracers (7386, -0.86 DPS) [crafted]; Condor Bracers (15864, -0.89 DPS) [quest]; Guardian Leather Bracers (4260, -1.14 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 28.9 spell_power points (7.26 DPS) | yes | Dreamweave Gloves (10019, -0.98 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -3.49 DPS) [crafted]; Red Mageweave Gloves (10018, -3.63 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 14.0 spell_power points (3.53 DPS) | yes | Skycaller's Leather Belt (252522, -0.24 DPS) [crafted]; Gilded Cord (254037, -0.44 DPS) [crafted]; Deathmage Sash (10771, -0.47 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.1 spell_power points (4.56 DPS) | yes | Gaze Dreamer Pants (6903, -1.00 DPS, sim-verified) [dungeon]; Dark Ritual Leggings (270031, -1.04 DPS) [quest]; Warchief Kilt (7760, -1.13 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.8 spell_power points (6.49 DPS) | yes | Skycaller's Leather Shoes (252532, -1.81 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -3.42 DPS) [crafted]; Gilded Slippers (254001, -3.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (2.95 DPS) | yes | Ring of Power Regulation (273030, -1.05 DPS) [dungeon]; Reedknot Ring (9622, -1.19 DPS) [quest]; Solid Gold Band (279845, -1.62 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.7 spell_power points (1.94 DPS) | yes | Ring of Power Regulation (273030, -0.04 DPS) [dungeon]; Reedknot Ring (9622, -0.18 DPS) [quest]; Solid Gold Band (279845, -0.61 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodscalp Channeling Staff (1998, -0.80 DPS) [world]; Rhahk'Zor's Hammer (5187, -1.11 DPS) [dungeon]; Windweaver Staff (7757, -17.02 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Papal Fez; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 547, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 50 (tauren, 5232221115100151-00000000000000000000-5400000000000000)

Set DPS (verified): 147.4. Weights run: 3.1s. Verify run: 1.9s. 717 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.477 ± 0.013, crit=0.104 ± 0.004 per rating point (14 rating = 1%, 1.460 per %), hit=0.303 ± 0.020 per rating point (10 rating = 1%, 3.032 per %), spell_haste=not significant (0.348 ± 0.161), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.508 ± 0.001, arcane_power=0.492 ± 0.002, spirit=not significant (0.381 ± 0.152), mp5=not significant (0.609 ± 0.154)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 33.2 spell_power points (7.86 DPS) | yes | Dreamweave Circlet (10041, -0.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.10 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.47 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.8 spell_power points (2.56 DPS) | yes | Mindburst Medallion (11196, -0.01 DPS) [quest]; Triune Amulet (7722, -0.04 DPS) [dungeon]; Horizon Choker (13085, -0.62 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 27.6 spell_power points (6.53 DPS) | yes | Blood Guard's Crackling Leather Spaulders (220871, -1.84 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.90 DPS) [crafted]; Rotgrip Mantle (17732, -3.05 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.1 spell_power points (4.54 DPS) | yes | Deep Woodlands Cloak (19121, -0.68 DPS) [quest]; Runecloth Cloak (13860, -1.50 DPS) [crafted]; Big Voodoo Cloak (8216, -2.33 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 32.9 spell_power points (7.80 DPS) | yes | Flamestrider Robes (11747, -1.16 DPS) [dungeon]; Dreamweave Vest (10021, -1.25 DPS) [crafted]; Acumen Robes (17775, -2.26 DPS, sim-verified) [quest] |
| wrist | Skycaller's Leather Bracers (252542) | Leatherworking [crafted] | 15.2 spell_power points (3.61 DPS) | yes | Runic Leather Bracers (15092, -0.11 DPS) [crafted]; Nethergeld Cuffs (254061, -0.71 DPS) [crafted]; Nethershine Cuffs (254069, -0.85 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 29.1 spell_power points (6.89 DPS) | yes | Skycaller's Leather Gauntlets (252550, -1.54 DPS) [crafted]; Dreamweave Gloves (10019, -1.54 DPS) [crafted]; Runecloth Gloves (13863, -2.22 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.7 spell_power points (5.15 DPS) | yes | Ghostweave Cord (254073, -0.84 DPS) [crafted]; Tooled Leather Belt (252544, -1.05 DPS) [crafted]; Satyrmane Sash (17755, -1.57 DPS, sim-verified) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 22.2 spell_power points (5.26 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -0.22 DPS) [vendor]; Jinxed Hoodoo Kilt (9474, -0.52 DPS) [dungeon]; Big Voodoo Pants (8202, -0.58 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.7 spell_power points (6.32 DPS) | yes | Skycaller's Leather Boots (252471, -1.53 DPS, sim-verified) [crafted]; First Sergeant's Crackling Leather Boots (220863, -1.86 DPS) [vendor]; Skycaller's Leather Shoes (252532, -2.00 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.08 DPS) | yes | Snake Hoop (6750, -1.09 DPS) [quest]; Ring of Power Regulation (273030, -1.23 DPS) [dungeon]; Eye of Adaegus (5266, -1.32 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (2.93 DPS) | yes | Snake Hoop (6750, -0.95 DPS) [quest]; Ring of Power Regulation (273030, -1.08 DPS) [dungeon]; Eye of Adaegus (5266, -1.17 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (147.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.18 DPS) [quest]; Rune of the Guard Captain (19120, -0.58 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (147.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.14 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (147.4 DPS) | yes | Kindling Stave (11750, -0.22 DPS) [dungeon]; Spellshifter Rod (9527, -0.74 DPS) [quest]; Zum'rah's Vexing Cane (18082, -2.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Glowing Eye of Mordresh; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Skycaller's Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Senior Designer's Pantaloons; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar

No-known-source sample (15 of 717, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60 (tauren, 5232221115100151-00000000000000000000-5533030000000000)

Set DPS (verified): 249.4. Weights run: 2.8s. Verify run: 4.5s. 1597 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.531 ± 0.017, crit=0.130 ± 0.005 per rating point (14 rating = 1%, 1.825 per %), hit=0.397 ± 0.027 per rating point (10 rating = 1%, 3.970 per %), spell_haste=0.999 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.002, spirit=not significant (0.576 ± 0.219), mp5=not significant (0.863 ± 0.230)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 42.6 spell_power points (11.81 DPS) | yes | Warlord's Dragonhide Helm (231678, -0.87 DPS) [vendor]; Champion's Dragonhide Helm (227186, -1.40 DPS) [pvp]; Soothsayer's Headdress (17740, -3.70 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.3 spell_power points (4.80 DPS) | yes | Mindburst Medallion (11196, -1.45 DPS) [quest]; Triune Amulet (7722, -1.46 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.68 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 30.2 spell_power points (8.38 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Epaulets (16551, -0.01 DPS) [vendor]; Argent Shoulders (19059, -9.52 DPS, sim-verified) [crafted] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 23.1 spell_power points (6.39 DPS) | yes | Spritecaster Cape (11623, -0.67 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.74 DPS) [quest]; Runecloth Cloak (13860, -2.72 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 43.5 spell_power points (12.06 DPS) | yes | Warlord's Dragonhide Armor (231679, -0.93 DPS) [vendor]; Legionnaire's Dragonhide Armor (227189, -1.95 DPS) [vendor]; Feathered Breastplate (8349, -10.21 DPS, sim-verified) [crafted] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (249.4 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, -5.31 DPS, sim-verified) [crafted] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-verified (249.4 DPS) | yes | Blood Guard's Dragonhide Gloves (227187, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Gloves of the Greatfather (17721, -6.56 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (249.4 DPS) | yes | Dreadmist Belt (16702, +0.00 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -4.93 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.4 spell_power points (13.69 DPS) | yes | General's Dragonhide Pants (231680, -2.04 DPS) [vendor]; Luminary Kilt (11823, -3.37 DPS) [dungeon]; Ghostweave Pants (14144, -4.90 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.0 spell_power points (7.77 DPS) | yes | Blood Guard's Dragonhide Boots (227188, +0.00 DPS) [pvp]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Venomspew Footpads (275606, +0.00 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.2 spell_power points (6.98 DPS) | yes | Philanthropist's Ring (281635, -3.47 DPS) [quest]; Band of the Hierophant (13096, -3.75 DPS) [world_drop]; Eye of Adaegus (5266, -4.18 DPS) [world_drop] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.60 DPS) | yes | Philanthropist's Ring (281635, -0.10 DPS) [quest]; Band of the Hierophant (13096, -0.38 DPS) [world_drop]; Eye of Adaegus (5266, -0.81 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (249.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (249.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (249.4 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -12.19 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-verified (249.4 DPS) | yes | Idol of the Moon (23197, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Cloak of Earth and Sky; chest: Ironfeather Breastplate; wrist: Magister's Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 540.1. Weights run: 2.0s. Verify run: 4.3s. 1597 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.178 ± 0.011, crit=0.271 ± 0.011 per rating point (14 rating = 1%, 3.792 per %), hit=0.502 ± 0.030 per rating point (10 rating = 1%, 5.022 per %), spell_haste=2.953 ± 0.240, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.329 ± 0.000, arcane_power=0.671 ± 0.001, spirit=not significant (0.077 ± 0.052), mp5=not significant (0.218 ± 0.083)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 29.3 spell_power points (17.04 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.35 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -1.33 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.8 spell_power points (8.02 DPS) | yes | Glowing Eye of Mordresh (10769, -3.12 DPS) [dungeon]; Mindburst Medallion (11196, -3.68 DPS) [quest]; Triune Amulet (7722, -5.83 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (540.1 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Argent Shoulders (19059, -8.71 DPS, sim-verified) [crafted] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 16.1 spell_power points (9.36 DPS) | yes | Spritecaster Cape (11623, -0.32 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.45 DPS) [quest]; Runecloth Cloak (13860, -3.30 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (540.1 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -10.36 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.8 spell_power points (7.43 DPS) | yes | Runic Leather Bracers (15092, -0.58 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.67 DPS) [crafted]; Nethershine Cuffs (254069, -1.41 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 27.3 spell_power points (15.88 DPS) | yes | General's Dragonhide Gloves (231677, -1.48 DPS) [pvp]; Earth Warder's Gloves (21318, -3.17 DPS, sim-verified) [quest]; Blood Guard's Dragonhide Gloves (227187, -4.19 DPS) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 18.1 spell_power points (10.55 DPS) | yes | Satyrmane Sash (17755, -1.37 DPS) [dungeon]; Ghostweave Cord (254073, -1.91 DPS) [crafted]; Mana-infused Cord (279267, -1.93 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 45.0 spell_power points (26.18 DPS) | yes | General's Dragonhide Pants (231680, -6.36 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -9.63 DPS) [vendor]; Ghostweave Pants (14144, -13.52 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.5 spell_power points (14.28 DPS) | yes | General's Dragonhide Boots (231682, -0.07 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.85 DPS) [pvp]; Venomspew Footpads (275606, -1.12 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (11.47 DPS) | yes | Philanthropist's Ring (281635, -5.13 DPS) [quest]; Reedknot Ring (9622, -7.39 DPS) [quest]; Snake Hoop (6750, -7.46 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.56 DPS) | yes | Philanthropist's Ring (281635, -2.58 DPS, sim-verified) [quest]; Reedknot Ring (9622, -3.49 DPS) [quest]; Snake Hoop (6750, -3.56 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (540.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, -1.87 DPS) [quest]; Ankh of Life (1713, -2.09 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (540.1 DPS) | yes | Shard of the Splithooves (10659, -1.28 DPS) [quest]; Ankh of Life (1713, -1.51 DPS) [world_drop]; Uther's Strength (11302, -2.88 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (540.1 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -16.10 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-verified (540.1 DPS) | yes | Idol of the Dream (220606, +0.00 DPS) [dungeon] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Cloak of Earth and Sky; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; ranged: Idol of the Moon

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

