# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 46.3. Weights run: 3.8s. Verify run: 5.9s. 251 eligible items had no known source.

Race sweep: high-order-skyborne 46.4 ± 0.1, night-elf 46.3 ± 0.1 (dps); high-order-skyborne beats the guide race night-elf by 0.0 (0.09%) ± 0.1, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.523 ± 0.009, crit=0.052 ± 0.001 per rating point (14 rating = 1%, 0.729 per %), hit=0.201 ± 0.006 per rating point (10 rating = 1%, 2.012 per %), spell_haste=0.521 ± 0.064, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.792 ± 0.001, arcane_power=0.208 ± 0.001, spirit=-0.078 ± 0.016, mp5=0.794 ± 0.081

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (2.72 DPS) | yes | Pristine Circlet (253949, -0.91 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.91 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.45 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.7 spell_power points (2.93 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.80 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.72 DPS) [crafted]; Slime-encrusted Pads (6461, -2.21 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.29 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.30 DPS) [crafted]; Caretaker's Cape (20428, -0.30 DPS) [rep] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 10.6 spell_power points (3.20 DPS) | yes | Filigreed Pristine Gown (253901, -0.91 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.91 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.92 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 5.0 spell_power points (1.51 DPS) | yes | Bright Bracers (3647, -0.88 DPS) [world_drop]; Inscribed Leather Bracers (3205, -1.03 DPS) [dungeon]; Owl Bracers (4796, -1.11 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.1 spell_power points (2.14 DPS) | yes | Serpent Gloves (5970, -0.03 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.33 DPS) [world]; Windfelt Gloves (5630, -0.46 DPS) [quest] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.1 spell_power points (1.84 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.16 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.30 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.1 spell_power points (3.96 DPS) | yes | Wisdom's Leather Pants (252503, -0.91 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -1.21 DPS) [crafted]; Totemic Leather Pants (252446, -1.22 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.1 spell_power points (2.74 DPS) | yes | Stormrider's Leather Boots (252443, -0.14 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.75 DPS) [crafted]; Totemic Leather Boots (252442, -0.93 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.6 spell_power points (1.99 DPS) | yes | Lavishly Jeweled Ring (1156, -1.04 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.51 DPS) [world_drop]; Furlfeather Ring (286745, -1.69 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.0 spell_power points (1.82 DPS) | yes | Lavishly Jeweled Ring (1156, -1.07 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -1.35 DPS) [world_drop]; Furlfeather Ring (286745, -1.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 29.0 spell_power points (8.76 DPS) | yes | Staff of the Blessed Seer (2271, -1.33 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.52 DPS) [dungeon]; Buzzer Blade (2169, -2.13 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.0 spell_power points (0.91 DPS) | yes | Mathystra Channeling Orb (281291, -0.12 DPS) [quest]; Tear of Grief (5611, -0.33 DPS, sim-verified) [quest]; Spellbinder Orb (15926, -0.43 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome

No-known-source sample (15 of 251, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (night-elf, 5232220113000000-00000000000000000000-0000000000000000)

Set DPS (verified): 82.3. Weights run: 3.9s. Verify run: 5.8s. 424 eligible items had no known source.

Race sweep: night-elf 82.3 ± 0.1, high-order-skyborne 81.9 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.521 ± 0.003, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.523 per %), hit=0.091 ± 0.003 per rating point (10 rating = 1%, 0.908 per %), spell_haste=0.404 ± 0.024, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.867 ± 0.002, arcane_power=0.133 ± 0.001, spirit=not significant (0.035 ± 0.014), mp5=0.466 ± 0.026

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (3.32 DPS) | yes | Holy Shroud (2721, -0.22 DPS) [world_drop]; Enchanter's Cowl (4322, -0.67 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.83 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.1 spell_power points (0.87 DPS) | yes | Seal of the Expedition (271920, -0.14 DPS) [rep]; Kaleidoscope Chain (13084, -0.25 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.26 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.7 spell_power points (3.79 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.11 DPS) [world_drop]; Invoker's Mantle (215365, -1.13 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Expeditionary's Cape (271923, -0.20 DPS) [rep]; Prelacy Cape (7004, -0.22 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.8 spell_power points (4.37 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.42 DPS) [dungeon]; Guardian Armor (4256, -0.43 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 spell_power points (2.88 DPS) | yes | Mindthrust Bracers (1974, -1.77 DPS) [dungeon]; Nightsky Wristbands (6407, -1.98 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.96 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | 7.3 spell_power points (2.03 DPS) | yes | Serpent Gloves (5970, -0.09 DPS) [dungeon]; Gloves of Insight (9698, -0.09 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.20 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (3.67 DPS) | yes | Highlander's Cloth Girdle (20099, -0.20 DPS) [rep]; Moss Cinch (6911, -0.35 DPS) [dungeon]; Belt of Arugal (6392, -0.70 DPS) [dungeon] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (3.88 DPS) | yes | Stormrider's Leather Pants (252502, -0.20 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.47 DPS) [dungeon]; Guardian Pants (5962, -0.51 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.99 DPS) | yes | Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -0.47 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.60 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.9 spell_power points (2.74 DPS) | yes | Ring of Power Regulation (273030, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.98 DPS) [world]; Minor Channeling Ring (1449, -1.07 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.9 spell_power points (2.20 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Runemender's Seal (273031, -0.43 DPS) [world]; Minor Channeling Ring (1449, -0.52 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.15 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.15 DPS) [world]; Glimmering Staff (249392, -3.83 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.1 spell_power points (2.80 DPS) | yes | Eye of Paleth (2943, -1.70 DPS) [quest]; Satyr's Rod (15962, -1.92 DPS) [world_drop]; Orb of Souls (249395, -2.21 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 169.4. Weights run: 4.0s. Verify run: 8.3s. 607 eligible items had no known source.

Race sweep: night-elf 169.4 ± 0.2, high-order-skyborne 169.0 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.550 ± 0.017, crit=0.241 ± 0.011 per rating point (14 rating = 1%, 3.369 per %), hit=0.457 ± 0.023 per rating point (10 rating = 1%, 4.570 per %), spell_haste=1.748 ± 0.204, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.619 ± 0.002, arcane_power=0.381 ± 0.002, spirit=not significant (0.184 ± 0.138), mp5=not significant (0.490 ± 0.135)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Papal Fez (9431, -0.81 DPS) [dungeon]; Corpseshroud (10574, -1.45 DPS) [dungeon]; Miner's Hat of the Deep (9429, -1.67 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.4 spell_power points (4.64 DPS) | yes | Triune Amulet (7722, -0.86 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.98 DPS) [quest]; Necklace of Calisea (1714, -2.55 DPS) [dungeon] |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodmage Mantle (7684, -0.04 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.33 DPS, sim-verified) [dungeon]; Berylline Pads (4197, -0.63 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.26 DPS) [crafted]; Caretaker's Cape (19532, -0.74 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (10.68 DPS) | yes | Dreamweave Vest (10021, -1.01 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.69 DPS) [crafted]; Elemental Raiment (9434, -2.13 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.5 spell_power points (4.26 DPS) | yes | Guardian Leather Bracers (4260, -0.48 DPS) [crafted]; Arcane Runed Bracers (4744, -0.60 DPS) [quest]; Green Whelp Bracers (7386, -1.17 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -4.69 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -5.57 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.80 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.2 spell_power points (6.60 DPS) | yes | Skycaller's Leather Belt (252522, -0.88 DPS) [crafted]; Star Belt (4329, -1.00 DPS) [crafted]; Deathmage Sash (10771, -2.39 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.6 spell_power points (8.39 DPS) | yes | Crimson Silk Pantaloons (7062, -2.39 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -2.67 DPS) [crafted]; Dark Ritual Leggings (270031, -2.69 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.3 spell_power points (10.30 DPS) | yes | Skycaller's Leather Shoes (252532, -4.20 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -4.95 DPS) [crafted]; Gilded Slippers (254001, -5.58 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (5.19 DPS) | yes | Ring of Forlorn Spirits (2043, -1.93 DPS) [quest]; Runemender's Seal (273031, -2.10 DPS) [world]; Snake Hoop (6750, -4.66 DPS, sim-verified) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.92 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.01 DPS) [quest]; Runemender's Seal (273031, -1.17 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.06 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -12.52 DPS) [dungeon]; Swinetusk Shank (6691, -15.86 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Green Silken Shoulders; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 5232220115101151-00000000000000000000-0540000000000000)

Set DPS (verified): 249.9. Weights run: 4.0s. Verify run: 10.0s. 815 eligible items had no known source.

Race sweep: night-elf 249.9 ± 0.3, high-order-skyborne 249.0 ± 0.3 (dps); the guide race night-elf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.690 ± 0.020, crit=0.311 ± 0.012 per rating point (14 rating = 1%, 4.360 per %), hit=0.654 ± 0.031 per rating point (10 rating = 1%, 6.542 per %), spell_haste=1.878 ± 0.243, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.647 ± 0.002, arcane_power=0.353 ± 0.002, spirit=not significant (0.243 ± 0.149), mp5=0.816 ± 0.156

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 35.3 spell_power points (15.81 DPS) | yes | Dreamweave Circlet (10041, -2.00 DPS) [crafted]; Red Mageweave Headband (10033, -2.37 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.71 DPS) [crafted] |
| neck | Talvash's Enhancing Necklace (7673) | Restoring the Necklace [quest] | 12.7 spell_power points (5.68 DPS) | yes | Mindburst Medallion (11196, -0.60 DPS) [quest]; Gemshard Heart (17707, -0.66 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.32 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 30.7 spell_power points (13.77 DPS) | yes | Rotgrip Mantle (17732, -1.94 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -2.37 DPS) [vendor]; Kentic Amice (11624, -2.83 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -1.40 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -3.76 DPS) [crafted]; Expeditionary's Cape (271921, -3.91 DPS) [rep] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 31.7 spell_power points (14.21 DPS) | yes | Forest's Embrace (22272, -1.47 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -1.85 DPS) [crafted]; Robes of Insight (940, -1.88 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 16.9 spell_power points (7.57 DPS) | yes | Skycaller's Leather Bracers (252542, -0.38 DPS) [crafted]; Aristocratic Cuffs (12546, -1.14 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.72 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.59 DPS) [crafted]; Ogreseer Fists (11665, -3.65 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -5.30 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 25.4 spell_power points (11.39 DPS) | yes | Skycaller's Leather Waistguard (252476, -0.52 DPS) [crafted]; Satyrmane Sash (17755, -2.03 DPS) [dungeon]; Dawnspire Cord (12466, -2.84 DPS) [dungeon] |
| legs | Knight's Crackling Leather Leggings (220864) | PvP rank 6 · Knight · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Brusslehide Leggings (17751, -0.88 DPS) [dungeon]; Red Mageweave Pants (10009, -3.37 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.7 spell_power points (11.51 DPS) | yes | Sergeant Major's Crackling Leather Boots (220862, -0.57 DPS) [vendor]; Skycaller's Leather Boots (252471, -1.40 DPS) [crafted]; Skycaller's Leather Shoes (252532, -2.98 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 19.1 spell_power points (8.56 DPS) | yes | Cyclopean Band (11824, -1.93 DPS) [dungeon]; Philanthropist's Ring (281635, -2.54 DPS) [quest]; Band of the Unicorn (7553, -2.74 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 15.3 spell_power points (6.84 DPS) | yes | Philanthropist's Ring (281635, -0.81 DPS) [quest]; Band of the Unicorn (7553, -1.01 DPS) [world_drop]; Cyclopean Band (11824, -1.44 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -1.33 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.55 DPS) [quest] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.52 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.74 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -6.86 DPS, sim-verified) [world_drop] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Enthralled Sphere (11625, -0.48 DPS) [dungeon]; Orb of Lorica (11262, -0.49 DPS) [quest]; Truesilver Conduit (249455, -1.11 DPS, sim-verified) [crafted] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Talvash's Enhancing Necklace; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Ban'thok Sash; legs: Knight's Crackling Leather Leggings; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Enriched Seal; main_hand: Blade of Eternal Darkness; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 815, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 373.3. Weights run: 3.4s. Verify run: 31.0s. 1969 eligible items had no known source.

Race sweep: night-elf 373.3 ± 0.4, high-order-skyborne 371.7 ± 0.4 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.371 ± 0.020, crit=0.452 ± 0.017 per rating point (14 rating = 1%, 6.330 per %), hit=0.713 ± 0.037 per rating point (10 rating = 1%, 7.134 per %), spell_haste=3.631 ± 0.328, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.477 ± 0.002, arcane_power=0.523 ± 0.002, spirit=not significant (0.361 ± 0.176), mp5=not significant (0.309 ± 0.218)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 37.3 spell_power points (17.74 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.03 DPS) [vendor]; Living Crown (252561, -0.28 DPS) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 19.1 spell_power points (9.08 DPS) | yes | Lady Maye's Pendant (14558, -0.50 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -1.33 DPS) [world_drop]; Archlight Talisman (15856, -1.61 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Spaulders (227190, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -7.66 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.46 DPS, sim-verified) [world]; Spritecaster Cape (11623, -3.66 DPS) [dungeon]; Shroud of the Exile (15421, -3.74 DPS) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 35.6 spell_power points (16.91 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Feathered Breastplate (8349, -12.24 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 27.5 spell_power points (13.07 DPS) | yes | Runecloth Cuffs (254123, -6.08 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.22 DPS) [crafted]; Runic Leather Bracers (15092, -6.55 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Hands of Power (13253, -2.25 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dragonhide Gloves (227193, -2.67 DPS) [pvp] |
| waist | Elunite Cord (272401) | Pix Xizzix [vendor] | 24.9 spell_power points (11.84 DPS) | yes | Highlander's Cloth Girdle (20047, -1.12 DPS) [rep]; Crystallized Girdle (12606, -1.28 DPS) [dungeon]; Ban'thok Sash (11662, -3.26 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.0 spell_power points (24.24 DPS) | yes | Sentinel's Silk Leggings (237815, -1.26 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.26 DPS) [vendor]; Marshal's Dragonhide Pants (231697, -4.22 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.5 spell_power points (12.61 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.18 DPS) [vendor]; Venomspew Footpads (275606, -0.78 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.30 DPS) [vendor]; The Postmaster's Seal (13392, -2.05 DPS) [dungeon]; Naglering (11669, -11.94 DPS, sim-verified) [dungeon] |
| finger2 | Songstone of Ironforge (12543) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -1.02 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -1.55 DPS) [dungeon]; Blessed Band of Light (272407, -3.46 DPS, sim-verified) [vendor] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, -4.16 DPS, sim-verified) [dungeon]; Smolderweb's Eye (13213, -5.16 DPS) [dungeon]; Serenity Field (272439, -5.23 DPS) [vendor] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -0.58 DPS) [dungeon]; Ironbark Staff (20069, -9.65 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.78 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Spirit of Aquementas (11904, -3.99 DPS, sim-verified) [quest]; Tome of the Lost (22253, -3.99 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, +0.00 DPS) [world_drop] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Elunite Cord; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Songstone of Ironforge; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch

No-known-source sample (15 of 1969, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 730.8. Weights run: 3.2s. Verify run: 29.2s. 1969 eligible items had no known source.

Race sweep: night-elf 730.8 ± 0.7, high-order-skyborne 729.9 ± 0.7 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.422 ± 0.018, crit=0.406 ± 0.016 per rating point (14 rating = 1%, 5.677 per %), hit=0.646 ± 0.034 per rating point (10 rating = 1%, 6.455 per %), spell_haste=2.716 ± 0.298, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.436 ± 0.001, arcane_power=0.564 ± 0.002, spirit=not significant (0.111 ± 0.146), mp5=not significant (0.626 ± 0.228)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 35.3 spell_power points (19.92 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.03 DPS) [vendor]; Living Crown (252561, -0.67 DPS) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Archlight Talisman (15856, -1.47 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -2.33 DPS) [world_drop]; Lady Maye's Pendant (14558, -5.22 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.2 spell_power points (15.88 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Burial Shawl (18681, -0.80 DPS) [dungeon]; Lieutenant Commander's Dragonhide Spaulders (227190, -1.10 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -4.87 DPS) [dungeon]; Shroud of the Exile (15421, -4.98 DPS) [quest]; Crystalline Threaded Cape (20697, -6.41 DPS, sim-verified) [world] |
| chest | Alanna's Embrace (13314) | Scholomance: Ras Frostwhisper [dungeon] | 31.4 spell_power points (17.73 DPS) | yes | Fernpulse Jerkin (15786, +0.00 DPS) [quest]; Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.2 spell_power points (14.75 DPS) | yes | Runecloth Cuffs (254123, -6.16 DPS) [crafted]; Runic Leather Bracers (15092, -6.73 DPS) [crafted]; Skycaller's Leather Bracers (252542, -7.13 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Gloves (231700, -0.23 DPS) [vendor]; Hands of Power (13253, -0.38 DPS) [dungeon]; Knight-Lieutenant's Dragonhide Gloves (227193, -3.21 DPS) [pvp] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormpike Cloth Girdle (19094, -0.26 DPS) [rep]; Highlander's Cloth Girdle (20047, -0.26 DPS) [rep]; Elunite Cord (272401, -3.13 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -1.32 DPS, sim-verified) [crafted]; Marshal's Dragonhide Pants (231697, -4.44 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.1 spell_power points (14.14 DPS) | yes | Knight-Lieutenant's Dragonhide Boots (227194, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.16 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.38 DPS) [quest]; Don Mauricio's Band of Domination (22433, -2.79 DPS) [dungeon]; Naglering (11669, -10.31 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Don Mauricio's Band of Domination (22433, -0.38 DPS) [dungeon]; Lorekeeper's Ring (19522, -0.47 DPS) [rep]; Songstone of Ironforge (12543, -1.98 DPS, sim-verified) [quest] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -3.14 DPS, sim-verified) [dungeon]; Serenity Field (272439, -6.20 DPS) [vendor]; Uther's Strength (11302, -6.77 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.04 DPS) [pvp]; Trindlehaven Staff (13161, -12.06 DPS) [dungeon]; Headmaster's Charge (13937, -13.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Brutality (23198, +0.00 DPS) [dungeon] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Alanna's Embrace; wrist: Dryad's Wrist Bindings; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1969, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 44.0. Weights run: 3.8s. Verify run: 6.4s. 247 eligible items had no known source.

Race sweep: windshaper-skyborne 44.8 ± 0.0, tauren 44.0 ± 0.0 (dps); windshaper-skyborne beats the guide race tauren by 0.9 (2.01%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.523 ± 0.009, crit=0.052 ± 0.001 per rating point (14 rating = 1%, 0.729 per %), hit=0.201 ± 0.006 per rating point (10 rating = 1%, 2.012 per %), spell_haste=0.521 ± 0.064, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.792 ± 0.001, arcane_power=0.208 ± 0.001, spirit=-0.078 ± 0.016, mp5=0.794 ± 0.081

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (2.72 DPS) | yes | Pristine Circlet (253949, -0.91 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.91 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.45 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.7 spell_power points (2.93 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.65 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.72 DPS) [crafted]; Slime-encrusted Pads (6461, -2.21 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.13 DPS) [crafted]; Black Whelp Cloak (7283, -0.30 DPS) [crafted]; Battle Healer's Cloak (20427, -0.30 DPS) [rep] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 10.6 spell_power points (3.20 DPS) | yes | Filigreed Pristine Gown (253901, -0.91 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.91 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.92 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 5.0 spell_power points (1.51 DPS) | yes | Tabitha's Cuffs (251486, -0.67 DPS, sim-verified) [quest]; Owl Bracers (4796, -0.72 DPS) [vendor]; Bright Bracers (3647, -0.88 DPS) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.1 spell_power points (2.14 DPS) | yes | Serpent Gloves (5970, -0.03 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.33 DPS) [world]; Pristine Gloves (253913, -0.46 DPS) [crafted] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.1 spell_power points (1.84 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.16 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.30 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.1 spell_power points (3.96 DPS) | yes | Wisdom's Leather Pants (252503, -0.91 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -1.21 DPS) [crafted]; Totemic Leather Pants (252446, -1.22 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.1 spell_power points (2.74 DPS) | yes | Stormrider's Leather Boots (252443, -0.14 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.75 DPS) [crafted]; Totemic Leather Boots (252442, -0.93 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.6 spell_power points (1.99 DPS) | yes | Loop of Sacrifice (281673, -1.20 DPS) [quest]; Volcanic Rock Ring (12053, -1.51 DPS) [world_drop]; Furlfeather Ring (286745, -1.69 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.1 spell_power points (0.95 DPS) | yes | Loop of Sacrifice (281673, -0.16 DPS) [quest]; Volcanic Rock Ring (12053, -0.47 DPS) [world_drop]; Furlfeather Ring (286745, -0.65 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (44.0 DPS) | yes | Crescent Staff (6505, -0.27 DPS, sim-verified) [quest]; Staff of the Blessed Seer (2271, -0.32 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.92 DPS) [quest] |
| off_hand | Spellbinder Orb (15926) (or Seer's Fine Stein (7608)) | Wailing Caverns: Large Battered Chest [dungeon] | 1.6 spell_power points (0.47 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop]; Tork Wrench (11855, -0.16 DPS) [quest]; Ancestral Orb (15944, -0.32 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Spellbinder Orb

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 8350 The 1 Ring; 14389 Durability Shoulderpads

### Band 30 (tauren, 5232220113000000-00000000000000000000-0000000000000000)

Set DPS (verified): 82.0. Weights run: 3.9s. Verify run: 6.2s. 425 eligible items had no known source.

Race sweep: windshaper-skyborne 83.1 ± 0.1, tauren 82.0 ± 0.1 (dps); windshaper-skyborne beats the guide race tauren by 1.1 (1.38%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.521 ± 0.003, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.523 per %), hit=0.091 ± 0.003 per rating point (10 rating = 1%, 0.908 per %), spell_haste=0.404 ± 0.024, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.867 ± 0.002, arcane_power=0.133 ± 0.001, spirit=not significant (0.035 ± 0.014), mp5=0.466 ± 0.026

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (3.32 DPS) | yes | Holy Shroud (2721, -0.22 DPS) [world_drop]; Enchanter's Cowl (4322, -0.59 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.83 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.1 spell_power points (0.87 DPS) | yes | Kaleidoscope Chain (13084, -0.25 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.26 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.42 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.7 spell_power points (3.79 DPS) | yes | Death Speaker Mantle (6685, -0.93 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.11 DPS) [world_drop]; Invoker's Mantle (215365, -1.13 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.04 DPS) [crafted]; Darkspear Raider's Cloak (272078, -0.24 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.8 spell_power points (4.37 DPS) | yes | Mechbuilder's Overalls (9508, -0.42 DPS) [dungeon]; Guardian Armor (4256, -0.43 DPS) [crafted]; Beguiler Robes (7728, -0.48 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 spell_power points (2.88 DPS) | yes | Mindthrust Bracers (1974, -1.77 DPS) [dungeon]; Nightsky Wristbands (6407, -1.98 DPS) [world_drop]; Glowing Magical Bracelets (13106, -3.13 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.0 spell_power points (2.78 DPS) | yes | Jutebraid Gloves (10654, -0.82 DPS, sim-verified) [quest]; Stormrider's Leather Gloves (252498, -0.82 DPS) [crafted]; Serpent Gloves (5970, -0.84 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (3.67 DPS) | yes | Moss Cinch (6911, -0.35 DPS) [dungeon]; Defiler's Cloth Girdle (20164, -0.43 DPS, sim-verified) [rep]; Warsong Sash (16975, -0.63 DPS) [quest] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (3.88 DPS) | yes | Gaze Dreamer Pants (6903, -0.47 DPS) [dungeon]; Guardian Pants (5962, -0.51 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.57 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.99 DPS) | yes | Spidersilk Boots (4320, -0.47 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.60 DPS) [crafted]; Acidic Walkers (9454, -0.74 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.9 spell_power points (2.74 DPS) | yes | Ring of Power Regulation (273030, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.98 DPS) [world]; Solid Gold Band (279845, -1.58 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.9 spell_power points (2.20 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Runemender's Seal (273031, -0.43 DPS) [world]; Solid Gold Band (279845, -1.04 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.15 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.15 DPS) [world]; Glimmering Staff (249392, -3.91 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.1 spell_power points (2.80 DPS) | yes | Witch's Finger (16887, -1.65 DPS) [quest]; Satyr's Rod (15962, -1.92 DPS) [world_drop]; Orb of Souls (249395, -2.97 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 167.3. Weights run: 4.0s. Verify run: 7.8s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 169.0 ± 0.2, tauren 167.3 ± 0.2 (dps); windshaper-skyborne beats the guide race tauren by 1.7 (1.02%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.550 ± 0.017, crit=0.241 ± 0.011 per rating point (14 rating = 1%, 3.369 per %), hit=0.457 ± 0.023 per rating point (10 rating = 1%, 4.570 per %), spell_haste=1.748 ± 0.204, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.619 ± 0.002, arcane_power=0.381 ± 0.002, spirit=not significant (0.184 ± 0.138), mp5=not significant (0.490 ± 0.135)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Papal Fez (9431, -0.81 DPS) [dungeon]; Corpseshroud (10574, -1.45 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.4 spell_power points (4.64 DPS) | yes | Triune Amulet (7722, -0.86 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.98 DPS) [quest]; Necklace of Calisea (1714, -2.55 DPS) [dungeon] |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Bloodmage Mantle (7684, -0.04 DPS) [dungeon]; Berylline Pads (4197, -0.63 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.26 DPS) [crafted]; Battle Healer's Cloak (19528, -0.74 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (10.68 DPS) | yes | Dreamweave Vest (10021, -0.28 DPS) [crafted]; Robe of Power (7054, -1.69 DPS) [crafted]; Elemental Raiment (9434, -2.13 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.5 spell_power points (4.26 DPS) | yes | Green Whelp Bracers (7386, -0.21 DPS) [crafted]; Guardian Leather Bracers (4260, -0.48 DPS) [crafted]; Radiant Silver Bracers (4545, -0.84 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.54 DPS) [crafted]; Red Mageweave Gloves (10018, -5.57 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.80 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.2 spell_power points (6.60 DPS) | yes | Deathmage Sash (10771, -0.39 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.88 DPS) [crafted]; Star Belt (4329, -1.00 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.6 spell_power points (8.39 DPS) | yes | Crimson Silk Pantaloons (7062, -2.22 DPS) [crafted]; Stormrider's Leather Pants (252502, -2.67 DPS) [crafted]; Dark Ritual Leggings (270031, -2.69 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.3 spell_power points (10.30 DPS) | yes | Skycaller's Leather Shoes (252532, -3.06 DPS) [crafted]; Mender's Leather Shoes (252533, -4.95 DPS) [crafted]; Gilded Slippers (254001, -5.58 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (5.19 DPS) | yes | Snake Hoop (6750, -0.66 DPS) [quest]; Runemender's Seal (273031, -2.10 DPS) [world]; Ring of Power Regulation (273030, -2.15 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Runemender's Seal (273031, -1.17 DPS) [world]; Ring of Power Regulation (273030, -1.22 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -11.29 DPS) [dungeon]; Hypnotic Blade (7714, -12.52 DPS) [dungeon]; Swinetusk Shank (6691, -15.86 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Green Silken Shoulders; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 5232220115101151-00000000000000000000-0540000000000000)

Set DPS (verified): 247.5. Weights run: 4.0s. Verify run: 10.1s. 814 eligible items had no known source.

Race sweep: windshaper-skyborne 249.9 ± 0.3, tauren 247.5 ± 0.3 (dps); windshaper-skyborne beats the guide race tauren by 2.4 (0.96%) ± 0.4, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.690 ± 0.020, crit=0.311 ± 0.012 per rating point (14 rating = 1%, 4.360 per %), hit=0.654 ± 0.031 per rating point (10 rating = 1%, 6.542 per %), spell_haste=1.878 ± 0.243, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.647 ± 0.002, arcane_power=0.353 ± 0.002, spirit=not significant (0.243 ± 0.149), mp5=0.816 ± 0.156

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 35.3 spell_power points (15.81 DPS) | yes | Red Mageweave Headband (10033, -1.85 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -2.00 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.71 DPS) [crafted] |
| neck | Jarkal's Enhancing Necklace (7888) | Necklace Recovery, Take 3 [quest] | 12.7 spell_power points (5.68 DPS) | yes | Mindburst Medallion (11196, -0.60 DPS) [quest]; Gemshard Heart (17707, -0.66 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.73 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 30.7 spell_power points (13.77 DPS) | yes | Rotgrip Mantle (17732, -2.22 DPS, sim-verified) [dungeon]; Blood Guard's Crackling Leather Spaulders (220871, -2.37 DPS) [vendor]; Kentic Amice (11624, -2.83 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.62 DPS) [quest]; Runecloth Cloak (13860, -2.28 DPS) [crafted]; Big Voodoo Cloak (8216, -3.76 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 31.7 spell_power points (14.21 DPS) | yes | Forest's Embrace (22272, -1.65 DPS) [quest]; Dreamweave Vest (10021, -1.85 DPS) [crafted]; Robes of Insight (940, -1.88 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 16.9 spell_power points (7.57 DPS) | yes | Skycaller's Leather Bracers (252542, -0.38 DPS) [crafted]; Aristocratic Cuffs (12546, -1.14 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.72 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.59 DPS) [crafted]; Ogreseer Fists (11665, -3.65 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -6.18 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 25.4 spell_power points (11.39 DPS) | yes | Skycaller's Leather Waistguard (252476, -1.42 DPS, sim-verified) [crafted]; Satyrmane Sash (17755, -2.03 DPS) [dungeon]; Dawnspire Cord (12466, -2.84 DPS) [dungeon] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Brusslehide Leggings (17751, -0.88 DPS) [dungeon]; Red Mageweave Pants (10009, -3.37 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.7 spell_power points (11.51 DPS) | yes | Skycaller's Leather Boots (252471, -1.40 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -2.43 DPS, sim-verified) [vendor]; Skycaller's Leather Shoes (252532, -2.98 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 19.1 spell_power points (8.56 DPS) | yes | Cyclopean Band (11824, -1.93 DPS) [dungeon]; Philanthropist's Ring (281635, -2.54 DPS) [quest]; Band of the Unicorn (7553, -2.74 DPS) [world_drop] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 15.3 spell_power points (6.84 DPS) | yes | Cyclopean Band (11824, -0.21 DPS) [dungeon]; Philanthropist's Ring (281635, -0.81 DPS) [quest]; Band of the Unicorn (7553, -1.01 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.59 DPS) [quest]; Darkspear Voodoo Seal (272061, -0.81 DPS) [vendor]; Ankh of Life (1713, -1.33 DPS) [world_drop] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.64 DPS) [quest]; Darkspear Voodoo Seal (272061, -0.86 DPS) [vendor]; Ankh of Life (1713, -2.35 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -7.09 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.5 spell_power points (6.95 DPS) | yes | Flask of Forest Mojo (19115, -0.45 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.68 DPS) [dungeon]; Enthralled Sphere (11625, -1.16 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Jarkal's Enhancing Necklace; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Ban'thok Sash; legs: Stone Guard's Crackling Leather Leggings; finger1: Brainlash; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit

No-known-source sample (15 of 814, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 365.1. Weights run: 3.4s. Verify run: 30.0s. 1973 eligible items had no known source.

Race sweep: windshaper-skyborne 368.9 ± 0.4, tauren 365.1 ± 0.4 (dps); windshaper-skyborne beats the guide race tauren by 3.8 (1.04%) ± 0.5, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.371 ± 0.020, crit=0.452 ± 0.017 per rating point (14 rating = 1%, 6.330 per %), hit=0.713 ± 0.037 per rating point (10 rating = 1%, 7.134 per %), spell_haste=3.631 ± 0.328, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.477 ± 0.002, arcane_power=0.523 ± 0.002, spirit=not significant (0.361 ± 0.176), mp5=not significant (0.309 ± 0.218)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 37.3 spell_power points (17.74 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.03 DPS) [pvp]; Living Crown (252561, -2.21 DPS, sim-verified) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 19.1 spell_power points (9.08 DPS) | yes | Lady Maye's Pendant (14558, -0.50 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -1.33 DPS) [world_drop]; Archlight Talisman (15856, -1.61 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Dragonhide Spaulders (227184, +0.00 DPS) [pvp]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -7.70 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.91 DPS, sim-verified) [world]; Spritecaster Cape (11623, -3.66 DPS) [dungeon]; Shroud of the Exile (15421, -3.74 DPS) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 35.6 spell_power points (16.91 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Feathered Breastplate (8349, -12.09 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.5 spell_power points (13.07 DPS) | yes | Runecloth Cuffs (254123, -6.08 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.22 DPS) [crafted]; Runic Leather Bracers (15092, -6.55 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Hands of Power (13253, -0.93 DPS, sim-verified) [dungeon]; Blood Guard's Dragonhide Gloves (227187, -2.67 DPS) [vendor] |
| waist | Elunite Cord (272401) | Pix Xizzix [vendor] | 24.9 spell_power points (11.84 DPS) | yes | Ban'thok Sash (11662, -0.98 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -1.12 DPS) [rep]; Crystallized Girdle (12606, -1.28 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.0 spell_power points (24.24 DPS) | yes | General's Dragonhide Pants (231680, -4.22 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -6.88 DPS) [vendor]; Outrider's Silk Leggings (22747, -7.61 DPS, sim-verified) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.5 spell_power points (12.61 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.18 DPS) [pvp]; Venomspew Footpads (275606, -0.78 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.30 DPS) [vendor]; The Postmaster's Seal (13392, -2.05 DPS) [dungeon]; Naglering (11669, -9.77 DPS, sim-verified) [dungeon] |
| finger2 | Eye of Orgrimmar (12545) | The Princess Saved? [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.27 DPS) [vendor]; The Postmaster's Seal (13392, -1.02 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -1.55 DPS) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Smolderweb's Eye (13213, -5.16 DPS) [dungeon]; Serenity Field (272439, -5.23 DPS) [vendor] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -0.58 DPS) [dungeon]; Ironbark Staff (20220, -8.93 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (13.78 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Spirit of Aquementas (11904, -3.87 DPS, sim-verified) [quest]; Tome of the Lost (22253, -3.99 DPS) [dungeon] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Brutality (23198, -2.33 DPS, sim-verified) [dungeon] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Elunite Cord; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Eye of Orgrimmar; trinket1: Second Wind; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Idol of the Moon

No-known-source sample (15 of 1973, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 721.1. Weights run: 3.2s. Verify run: 30.6s. 1973 eligible items had no known source.

Race sweep: windshaper-skyborne 727.3 ± 0.7, tauren 721.1 ± 0.7 (dps); windshaper-skyborne beats the guide race tauren by 6.2 (0.86%) ± 1.0, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.248 ± 0.024, crit=0.756 ± 0.024 per rating point (14 rating = 1%, 10.585 per %), hit=1.116 ± 0.058 per rating point (10 rating = 1%, 11.159 per %), spell_haste=7.206 ± 0.451, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.351 ± 0.001, arcane_power=0.649 ± 0.001, spirit=not significant (0.050 ± 0.060), mp5=not significant (-0.038 ± 0.208)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.0 spell_power points (20.54 DPS) | yes | Champion's Dragonhide Helm (227186, +0.00 DPS) [pvp]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -0.05 DPS) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 22.1 spell_power points (13.79 DPS) | yes | Star of Mystaria (12103, -3.46 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.00 DPS) [world_drop]; Archlight Talisman (15856, -4.77 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.8 spell_power points (19.19 DPS) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Champion's Dragonhide Spaulders (227184, -1.31 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.22 DPS, sim-verified) [world]; Spritecaster Cape (11623, -8.32 DPS) [dungeon]; Shroud of the Exile (15421, -8.91 DPS) [quest] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Hauberk (16549, +0.00 DPS) [vendor]; Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.3 spell_power points (15.16 DPS) | yes | Runecloth Cuffs (254123, -6.76 DPS) [crafted]; Runic Leather Bracers (15092, -7.38 DPS) [crafted]; Skycaller's Leather Bracers (252542, -7.69 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, -1.17 DPS) [dungeon]; General's Dragonhide Gloves (231677, -1.33 DPS) [pvp]; Earth Warder's Gloves (21318, -3.43 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 26.1 spell_power points (16.24 DPS) | yes | Ban'thok Sash (11662, -0.27 DPS) [dungeon]; Elunite Cord (272401, -2.56 DPS) [vendor]; Frostwolf Cloth Belt (19090, -3.48 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 53.1 spell_power points (33.04 DPS) | yes | General's Dragonhide Pants (231680, -7.16 DPS) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -10.37 DPS, sim-verified) [vendor]; Legionnaire's Dragonhide Pants (227185, -10.49 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.42 DPS) [pvp]; First Sergeant's Crackling Leather Boots (220863, -1.54 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.12 DPS) [vendor]; Maiden's Circle (13001, -1.06 DPS) [world_drop]; Naglering (11669, -11.13 DPS, sim-verified) [dungeon] |
| finger2 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.02 DPS) [vendor]; Maiden's Circle (13001, -0.96 DPS) [world_drop]; Eye of Orgrimmar (12545, -1.15 DPS) [quest] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+6.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -4.26 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.96 DPS) [crafted]; Second Wind (11819, -6.86 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.12 DPS) [pvp]; Headmaster's Charge (13937, -10.00 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -19.11 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talons of Wrath (249441, +0.00 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Fernpulse Jerkin; wrist: Dryad's Wrist Bindings; waist: Defiler's Cloth Girdle; legs: Ironfeather Leggings; finger1: Don Julio's Band; finger2: Don Mauricio's Band of Domination; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1973, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

