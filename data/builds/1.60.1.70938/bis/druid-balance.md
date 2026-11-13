# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 46.3. Weights run: 4.2s. Verify run: 6.1s. 251 eligible items had no known source.

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

Set DPS (verified): 82.2. Weights run: 4.3s. Verify run: 7.0s. 424 eligible items had no known source.

Race sweep: night-elf 82.2 ± 0.1, high-order-skyborne 81.8 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.489 ± 0.010, crit=0.115 ± 0.002 per rating point (14 rating = 1%, 1.610 per %), hit=0.242 ± 0.009 per rating point (10 rating = 1%, 2.419 per %), spell_haste=0.649 ± 0.081, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.818 ± 0.001, arcane_power=0.182 ± 0.001, spirit=not significant (-0.008 ± 0.010), mp5=0.894 ± 0.098

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (4.59 DPS) | yes | Holy Shroud (2721, -0.39 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.43 DPS) [crafted]; Silk Headband (7050, -1.15 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.9 spell_power points (1.12 DPS) | yes | Seal of the Expedition (271920, -0.19 DPS) [rep]; Crystal Starfire Medallion (5003, -0.37 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.37 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.4 spell_power points (5.13 DPS) | yes | Death Speaker Mantle (6685, -0.76 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -1.51 DPS) [crafted]; Magician's Mantle (12998, -1.53 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.38 DPS) [crafted]; Caretaker's Cape (19533, -0.38 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.4 spell_power points (5.87 DPS) | yes | Beguiler Robes (7728, -0.44 DPS, sim-verified) [dungeon]; Guardian Armor (4256, -0.56 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.57 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.7 spell_power points (4.47 DPS) | yes | Mindthrust Bracers (1974, -2.64 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.97 DPS) [world_drop]; Nightsky Wristbands (6407, -3.35 DPS) [world_drop] |
| hands | Gloves of Insight (9698) (or Serpent Gloves (5970)) | Return to Vahlarriel [quest] | 7.0 spell_power points (2.68 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Stormrider's Leather Gloves (252498, -0.02 DPS) [crafted]; Truefaith Gloves (7049, -0.20 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.9 spell_power points (4.95 DPS) | yes | Highlander's Cloth Girdle (20099, -0.18 DPS) [rep]; Moss Cinch (6911, -0.36 DPS) [dungeon]; Belt of Arugal (6392, -0.94 DPS) [dungeon] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (5.36 DPS) | yes | Stormrider's Leather Pants (252502, -0.44 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -0.77 DPS) [dungeon]; Guardian Pants (5962, -0.80 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.4 spell_power points (3.99 DPS) | yes | Spidersilk Boots (4320, -0.58 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.58 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.76 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 8.8 spell_power points (3.36 DPS) | yes | Snake Hoop (6750, -0.13 DPS) [quest]; Minor Channeling Ring (1449, -1.08 DPS) [quest]; Runemender's Seal (273031, -1.09 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.7 spell_power points (3.32 DPS) | yes | Snake Hoop (6750, -0.09 DPS) [quest]; Minor Channeling Ring (1449, -1.04 DPS) [quest]; Runemender's Seal (273031, -1.05 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.53 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.53 DPS) [world]; Glimmering Staff (249392, -4.22 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.9 spell_power points (3.80 DPS) | yes | Orb of Souls (249395, -2.27 DPS) [crafted]; Eye of Paleth (2943, -2.63 DPS, sim-verified) [quest]; Dwarven Tome (279898, -2.65 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Gloves of Insight; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 165.9. Weights run: 4.6s. Verify run: 8.6s. 607 eligible items had no known source.

Race sweep: night-elf 165.9 ± 0.2, high-order-skyborne 165.3 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.527 ± 0.016, crit=0.226 ± 0.010 per rating point (14 rating = 1%, 3.158 per %), hit=0.451 ± 0.023 per rating point (10 rating = 1%, 4.505 per %), spell_haste=1.212 ± 0.196, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.623 ± 0.002, arcane_power=0.377 ± 0.002, spirit=not significant (0.308 ± 0.128), mp5=1.358 ± 0.177

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.59 DPS) | yes | Miner's Hat of the Deep (9429, -0.02 DPS) [dungeon]; Papal Fez (9431, -0.52 DPS) [dungeon]; Corpseshroud (10574, -1.63 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.2 spell_power points (4.59 DPS) | yes | Triune Amulet (7722, -0.43 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.91 DPS) [quest]; Necklace of Calisea (1714, -2.20 DPS) [dungeon] |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Bloodmage Mantle (7684, -0.02 DPS) [dungeon]; Mantle of Doan (7712, -0.48 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Icy Cloak (4327, +0.00 DPS) [crafted]; Guardian Cloak (5965, +0.00 DPS) [crafted]; Caretaker's Cape (19532, -0.45 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Vest (10021, +0.00 DPS) [crafted]; Robe of Power (7054, -1.60 DPS) [crafted]; Elemental Raiment (9434, -2.33 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.1 spell_power points (5.35 DPS) | yes | Green Whelp Bracers (7386, -1.33 DPS) [crafted]; Guardian Leather Bracers (4260, -1.60 DPS) [crafted]; Arcane Runed Bracers (4744, -1.67 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.26 DPS) [crafted]; Red Mageweave Gloves (10018, -5.71 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.92 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.1 spell_power points (6.59 DPS) | yes | Deathmage Sash (10771, -0.49 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.70 DPS) [crafted]; Star Belt (4329, -0.77 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.3 spell_power points (8.31 DPS) | yes | Warchief Kilt (7760, -1.98 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.21 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.24 DPS) [crafted] |
| feet | Skycaller's Leather Shoes (252532) | Leatherworking [crafted] | 18.2 spell_power points (7.46 DPS) | yes | Mender's Leather Shoes (252533, -1.79 DPS) [crafted]; Gilded Slippers (254001, -2.58 DPS) [crafted]; Acidic Walkers (9454, -3.18 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Power Regulation (273030, -1.05 DPS) [dungeon]; Snake Hoop (6750, -1.64 DPS) [quest]; Runemender's Seal (273031, -1.89 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Power Regulation (273030, -1.23 DPS) [dungeon]; Snake Hoop (6750, -1.82 DPS) [quest]; Runemender's Seal (273031, -2.07 DPS) [world] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -11.41 DPS) [dungeon]; Hypnotic Blade (7714, -12.78 DPS) [dungeon]; Swinetusk Shank (6691, -16.01 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Green Silken Shoulders; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Skycaller's Leather Shoes; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 5232220115101151-00000000000000000000-0540000000000000)

Set DPS (verified): 247.8. Weights run: 4.5s. Verify run: 8.0s. 817 eligible items had no known source.

Race sweep: night-elf 247.8 ± 0.3, high-order-skyborne 246.9 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.632 ± 0.019, crit=0.305 ± 0.012 per rating point (14 rating = 1%, 4.269 per %), hit=0.629 ± 0.030 per rating point (10 rating = 1%, 6.287 per %), spell_haste=1.696 ± 0.226, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.642 ± 0.002, arcane_power=0.358 ± 0.002, spirit=not significant (0.174 ± 0.125), mp5=0.902 ± 0.145

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 33.9 spell_power points (15.26 DPS) | yes | Red Mageweave Headband (10033, -1.54 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -2.01 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.10 DPS) [crafted] |
| neck | Talvash's Enhancing Necklace (7673) | Restoring the Necklace [quest] | 12.6 spell_power points (5.69 DPS) | yes | Gemshard Heart (17707, -0.88 DPS) [dungeon]; Mindburst Medallion (11196, -0.89 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.33 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 29.0 spell_power points (13.08 DPS) | yes | Rotgrip Mantle (17732, -1.62 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.99 DPS) [vendor]; Kentic Amice (11624, -2.60 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -3.18 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -3.67 DPS) [crafted]; Expeditionary's Cape (271921, -4.11 DPS) [rep] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 29.5 spell_power points (13.29 DPS) | yes | Robe of the Magi (1716, -1.28 DPS) [world_drop]; Robes of Insight (940, -1.50 DPS) [world_drop]; Forest's Embrace (22272, -2.01 DPS, sim-verified) [quest] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 16.3 spell_power points (7.35 DPS) | yes | Skycaller's Leather Bracers (252542, -0.46 DPS) [crafted]; Aristocratic Cuffs (12546, -1.28 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.81 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.91 DPS) [crafted]; Ogreseer Fists (11665, -4.22 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -5.91 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 24.6 spell_power points (11.09 DPS) | yes | Skycaller's Leather Waistguard (252476, -1.63 DPS, sim-verified) [crafted]; Satyrmane Sash (17755, -1.93 DPS) [dungeon]; Dawnspire Cord (12466, -2.97 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 29.3 spell_power points (13.21 DPS) | yes | Knight's Crackling Leather Leggings (220864, -0.20 DPS) [vendor]; Brusslehide Leggings (17751, -0.85 DPS) [dungeon]; Red Mageweave Pants (10009, -3.49 DPS) [crafted] |
| feet | Sergeant Major's Crackling Leather Boots (220862) | PvP rank 5 · Sergeant Major · Alliance [vendor] | 23.6 spell_power points (10.63 DPS) | yes | Skycaller's Leather Boots (252471, -1.91 DPS, sim-verified) [crafted]; Skycaller's Leather Shoes (252532, -2.39 DPS) [crafted]; Stormcloth Boots (10039, -2.40 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 18.2 spell_power points (8.20 DPS) | yes | Cyclopean Band (11824, -1.83 DPS) [dungeon]; Philanthropist's Ring (281635, -2.27 DPS) [quest]; Band of the Unicorn (7553, -2.34 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 15.6 spell_power points (7.03 DPS) | yes | Philanthropist's Ring (281635, -1.10 DPS) [quest]; Band of the Unicorn (7553, -1.18 DPS) [world_drop]; Cyclopean Band (11824, -1.42 DPS, sim-verified) [dungeon] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -1.25 DPS) [quest]; Ankh of Life (1713, -1.42 DPS, sim-verified) [world_drop] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.56 DPS, sim-verified) [crafted]; Ankh of Life (1713, -1.76 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.92 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Moonshadow Stave (22458, +0.00 DPS) [quest]; Staff of Jordan (873, -8.50 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.1 spell_power points (6.78 DPS) | yes | Orb of the Forgotten Seer (7685, -0.48 DPS) [dungeon]; Orb of Lorica (11262, -1.23 DPS) [quest]; Enthralled Sphere (11625, -1.28 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Talvash's Enhancing Necklace; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit

No-known-source sample (15 of 817, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 376.1. Weights run: 3.9s. Verify run: 33.4s. 1988 eligible items had no known source.

Race sweep: night-elf 376.1 ± 0.3, high-order-skyborne 374.8 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.359 ± 0.020, crit=0.445 ± 0.017 per rating point (14 rating = 1%, 6.231 per %), hit=0.759 ± 0.037 per rating point (10 rating = 1%, 7.594 per %), spell_haste=3.436 ± 0.325, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.493 ± 0.002, arcane_power=0.507 ± 0.002, spirit=not significant (0.243 ± 0.186), mp5=not significant (0.877 ± 0.225)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 43.7 spell_power points (20.98 DPS) | yes | Crimson Felt Hat (18727, -1.83 DPS, sim-verified) [dungeon]; Field Marshal's Dragonhide Helm (231695, -2.70 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -3.77 DPS) [vendor] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 19.5 spell_power points (9.34 DPS) | yes | Lady Maye's Pendant (14558, -0.60 DPS, sim-verified) [world_drop]; Archlight Talisman (15856, -1.86 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -2.19 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Spaulders (227190, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -7.27 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.41 DPS) [world]; Spritecaster Cape (11623, -4.25 DPS) [dungeon]; Shroud of the Exile (15421, -4.40 DPS) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 32.1 spell_power points (15.40 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Alanna's Embrace (13314, -9.71 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.6 spell_power points (12.74 DPS) | yes | Runecloth Cuffs (254123, -5.75 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.16 DPS) [crafted]; Runic Leather Bracers (15092, -6.23 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Gloves (231700, -0.32 DPS) [vendor]; Hands of Power (13253, -1.11 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dragonhide Gloves (227193, -3.05 DPS) [pvp] |
| waist | Elunite Cord (272401) | Pix Xizzix [vendor] | 24.6 spell_power points (11.81 DPS) | yes | Ban'thok Sash (11662, -0.69 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -1.07 DPS) [rep]; Stormpike Cloth Girdle (19094, -1.45 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.7 spell_power points (24.31 DPS) | yes | Sentinel's Silk Leggings (237815, -1.27 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.27 DPS) [vendor]; Marshal's Dragonhide Pants (231697, -4.80 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.3 spell_power points (11.66 DPS) | yes | Knight-Lieutenant's Dragonhide Boots (227194, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.18 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.17 DPS) [vendor]; Lorekeeper's Ring (19522, -2.14 DPS) [rep]; Naglering (11669, -10.52 DPS, sim-verified) [dungeon] |
| finger2 | Songstone of Ironforge (12543) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.71 DPS, sim-verified) [vendor]; Lorekeeper's Ring (19522, -1.44 DPS) [rep]; Don Mauricio's Band of Domination (22433, -1.57 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, -2.70 DPS, sim-verified) [dungeon]; Smolderweb's Eye (13213, -4.99 DPS) [dungeon]; Burst of Knowledge (11832, -5.27 DPS) [dungeon] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Blade of Eternal Darkness (17780, -6.98 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.91 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Spirit of Aquementas (11904, -4.00 DPS, sim-verified) [quest]; Tome of the Lost (22253, -4.07 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -2.20 DPS, sim-verified) [world_drop] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Elunite Cord; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Songstone of Ironforge; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch

No-known-source sample (15 of 1988, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 737.9. Weights run: 3.5s. Verify run: 32.8s. 1988 eligible items had no known source.

Race sweep: night-elf 737.9 ± 0.7, high-order-skyborne 737.4 ± 0.7 (dps); the guide race night-elf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.422 ± 0.018, crit=0.406 ± 0.016 per rating point (14 rating = 1%, 5.677 per %), hit=0.646 ± 0.034 per rating point (10 rating = 1%, 6.455 per %), spell_haste=2.716 ± 0.298, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.436 ± 0.001, arcane_power=0.564 ± 0.002, spirit=not significant (0.111 ± 0.146), mp5=not significant (0.626 ± 0.228)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 40.9 spell_power points (23.06 DPS) | yes | Field Marshal's Dragonhide Helm (231695, -1.69 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -3.17 DPS) [vendor]; Crimson Felt Hat (18727, -8.60 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Archlight Talisman (15856, -1.47 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -2.33 DPS) [world_drop]; Lady Maye's Pendant (14558, -5.08 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.2 spell_power points (15.88 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Lieutenant Commander's Dragonhide Spaulders (227190, -1.10 DPS) [vendor]; Burial Shawl (18681, -3.20 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.34 DPS) [world]; Spritecaster Cape (11623, -4.87 DPS) [dungeon]; Cloak of Flames (3475, -10.86 DPS, sim-verified) [world_drop] |
| chest | Alanna's Embrace (13314) | Scholomance: Ras Frostwhisper [dungeon] | 31.4 spell_power points (17.73 DPS) | yes | Fernpulse Jerkin (15786, +0.00 DPS) [quest]; Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.2 spell_power points (14.75 DPS) | yes | Runecloth Cuffs (254123, -6.16 DPS) [crafted]; Runic Leather Bracers (15092, -6.73 DPS) [crafted]; Skycaller's Leather Bracers (252542, -7.13 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Gloves (231700, -0.23 DPS) [vendor]; Hands of Power (13253, -0.38 DPS) [dungeon]; Knight-Lieutenant's Dragonhide Gloves (227193, -3.21 DPS) [pvp] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormpike Cloth Girdle (19094, -0.26 DPS) [rep]; Highlander's Cloth Girdle (20047, -0.26 DPS) [rep]; Elunite Cord (272401, -3.77 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -1.97 DPS, sim-verified) [crafted]; Marshal's Dragonhide Pants (231697, -4.44 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.1 spell_power points (14.14 DPS) | yes | Knight-Lieutenant's Dragonhide Boots (227194, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.18 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.38 DPS) [quest]; Don Mauricio's Band of Domination (22433, -2.79 DPS) [dungeon]; Naglering (11669, -12.63 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Don Mauricio's Band of Domination (22433, -0.38 DPS) [dungeon]; Lorekeeper's Ring (19522, -0.47 DPS) [rep]; Songstone of Ironforge (12543, -2.00 DPS, sim-verified) [quest] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Second Wind (11819, -9.71 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -3.52 DPS, sim-verified) [dungeon]; Serenity Field (272439, -6.20 DPS) [vendor]; Uther's Strength (11302, -6.77 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.04 DPS) [pvp]; Headmaster's Charge (13937, -10.83 DPS, sim-verified) [dungeon]; Trindlehaven Staff (13161, -12.06 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Brutality (23198, +0.00 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Alanna's Embrace; wrist: Dryad's Wrist Bindings; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1988, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 44.0. Weights run: 4.2s. Verify run: 6.7s. 247 eligible items had no known source.

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

Set DPS (verified): 81.8. Weights run: 4.3s. Verify run: 7.1s. 425 eligible items had no known source.

Race sweep: windshaper-skyborne 82.9 ± 0.1, tauren 81.8 ± 0.1 (dps); windshaper-skyborne beats the guide race tauren by 1.1 (1.28%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.489 ± 0.010, crit=0.115 ± 0.002 per rating point (14 rating = 1%, 1.610 per %), hit=0.242 ± 0.009 per rating point (10 rating = 1%, 2.419 per %), spell_haste=0.649 ± 0.081, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.818 ± 0.001, arcane_power=0.182 ± 0.001, spirit=not significant (-0.008 ± 0.010), mp5=0.894 ± 0.098

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (4.59 DPS) | yes | Enchanter's Cowl (4322, -0.43 DPS) [crafted]; Holy Shroud (2721, -0.45 DPS, sim-verified) [world_drop]; Silk Headband (7050, -1.15 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.9 spell_power points (1.12 DPS) | yes | Crystal Starfire Medallion (5003, -0.37 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.37 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.49 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.4 spell_power points (5.13 DPS) | yes | Death Speaker Mantle (6685, -1.07 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -1.51 DPS) [crafted]; Magician's Mantle (12998, -1.53 DPS) [world_drop] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.38 DPS) [crafted]; Battle Healer's Cloak (19529, -0.38 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.4 spell_power points (5.87 DPS) | yes | Beguiler Robes (7728, -0.55 DPS, sim-verified) [dungeon]; Guardian Armor (4256, -0.56 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.57 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.7 spell_power points (4.47 DPS) | yes | Mindthrust Bracers (1974, -2.58 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.97 DPS) [world_drop]; Tabitha's Cuffs (251486, -3.32 DPS) [quest] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.0 spell_power points (3.82 DPS) | yes | Jutebraid Gloves (10654, -0.69 DPS, sim-verified) [quest]; Serpent Gloves (5970, -1.14 DPS) [dungeon]; Stormrider's Leather Gloves (252498, -1.16 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.9 spell_power points (4.95 DPS) | yes | Moss Cinch (6911, -0.36 DPS) [dungeon]; Warsong Sash (16975, -0.74 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.76 DPS, sim-verified) [rep] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (5.36 DPS) | yes | Stormrider's Leather Pants (252502, -0.58 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -0.77 DPS) [dungeon]; Guardian Pants (5962, -0.80 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.4 spell_power points (3.99 DPS) | yes | Acidic Walkers (9454, -0.58 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.76 DPS) [crafted]; Spidersilk Boots (4320, -1.14 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 8.8 spell_power points (3.36 DPS) | yes | Snake Hoop (6750, -0.13 DPS) [quest]; Runemender's Seal (273031, -1.09 DPS) [world]; Solid Gold Band (279845, -1.83 DPS) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.7 spell_power points (3.32 DPS) | yes | Snake Hoop (6750, -0.51 DPS, sim-verified) [quest]; Runemender's Seal (273031, -1.05 DPS) [world]; Solid Gold Band (279845, -1.79 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.53 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.53 DPS) [world]; Glimmering Staff (249392, -4.11 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.9 spell_power points (3.80 DPS) | yes | Witch's Finger (16887, -2.30 DPS) [quest]; Satyr's Rod (15962, -2.68 DPS) [world_drop]; Orb of Souls (249395, -3.06 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 163.4. Weights run: 4.6s. Verify run: 8.6s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 165.0 ± 0.2, tauren 163.4 ± 0.2 (dps); windshaper-skyborne beats the guide race tauren by 1.6 (1.00%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.527 ± 0.016, crit=0.226 ± 0.010 per rating point (14 rating = 1%, 3.158 per %), hit=0.451 ± 0.023 per rating point (10 rating = 1%, 4.505 per %), spell_haste=1.212 ± 0.196, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.623 ± 0.002, arcane_power=0.377 ± 0.002, spirit=not significant (0.308 ± 0.128), mp5=1.358 ± 0.177

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.59 DPS) | yes | Papal Fez (9431, -0.52 DPS) [dungeon]; Corpseshroud (10574, -1.63 DPS) [dungeon]; Miner's Hat of the Deep (9429, -2.10 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.2 spell_power points (4.59 DPS) | yes | Triune Amulet (7722, -0.36 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.91 DPS) [quest]; Necklace of Calisea (1714, -2.20 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.8 spell_power points (5.66 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Bloodmage Mantle (7684, -0.04 DPS) [dungeon]; Mantle of Doan (7712, -0.50 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.09 DPS) [crafted]; Guardian Cloak (5965, -0.09 DPS) [crafted]; Battle Healer's Cloak (19528, -0.53 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -1.60 DPS) [crafted]; Dreamweave Vest (10021, -1.65 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -2.33 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.1 spell_power points (5.35 DPS) | yes | Green Whelp Bracers (7386, -0.89 DPS, sim-verified) [crafted]; Guardian Leather Bracers (4260, -1.60 DPS) [crafted]; Radiant Silver Bracers (4545, -1.99 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -4.55 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -5.71 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.92 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.1 spell_power points (6.59 DPS) | yes | Skycaller's Leather Belt (252522, -0.70 DPS) [crafted]; Star Belt (4329, -0.77 DPS) [crafted]; Deathmage Sash (10771, -1.47 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.3 spell_power points (8.31 DPS) | yes | Stoneweaver Leggings (9407, -2.21 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.24 DPS) [crafted]; Warchief Kilt (7760, -2.97 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Shoes (252532) | Leatherworking [crafted] | 18.2 spell_power points (7.46 DPS) | yes | Mender's Leather Shoes (252533, -2.08 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -2.58 DPS) [crafted]; Acidic Walkers (9454, -3.18 DPS) [dungeon] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 13.1 spell_power points (5.35 DPS) | yes | Ring of Power Regulation (273030, -1.23 DPS) [dungeon]; Snake Hoop (6750, -1.82 DPS) [quest]; Darkspear Signet (272070, -2.01 DPS) [vendor] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.6 spell_power points (5.17 DPS) | yes | Snake Hoop (6750, -1.64 DPS) [quest]; Darkspear Signet (272070, -1.83 DPS) [vendor]; Ring of Power Regulation (273030, -1.84 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.97 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -12.78 DPS) [dungeon]; Swinetusk Shank (6691, -16.01 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Skycaller's Leather Shoes; finger1: Advisor's Ring; finger2: Philanthropist's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 5232220115101151-00000000000000000000-0540000000000000)

Set DPS (verified): 243.8. Weights run: 4.5s. Verify run: 7.5s. 816 eligible items had no known source.

Race sweep: windshaper-skyborne 246.5 ± 0.3, tauren 243.8 ± 0.3 (dps); windshaper-skyborne beats the guide race tauren by 2.7 (1.09%) ± 0.4, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.632 ± 0.019, crit=0.305 ± 0.012 per rating point (14 rating = 1%, 4.269 per %), hit=0.629 ± 0.030 per rating point (10 rating = 1%, 6.287 per %), spell_haste=1.696 ± 0.226, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.642 ± 0.002, arcane_power=0.358 ± 0.002, spirit=not significant (0.174 ± 0.125), mp5=0.902 ± 0.145

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 33.9 spell_power points (15.26 DPS) | yes | Red Mageweave Headband (10033, -1.16 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -2.01 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.10 DPS) [crafted] |
| neck | Jarkal's Enhancing Necklace (7888) | Necklace Recovery, Take 3 [quest] | 12.6 spell_power points (5.69 DPS) | yes | Gemshard Heart (17707, -0.88 DPS) [dungeon]; Mindburst Medallion (11196, -0.89 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.69 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 29.0 spell_power points (13.08 DPS) | yes | Blood Guard's Crackling Leather Spaulders (220871, -1.99 DPS) [vendor]; Rotgrip Mantle (17732, -2.50 DPS, sim-verified) [dungeon]; Kentic Amice (11624, -2.60 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.52 DPS) [quest]; Runecloth Cloak (13860, -2.15 DPS) [crafted]; Big Voodoo Cloak (8216, -3.67 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 29.5 spell_power points (13.29 DPS) | yes | Robe of the Magi (1716, -1.28 DPS) [world_drop]; Robes of Insight (940, -1.50 DPS) [world_drop]; Forest's Embrace (22272, -1.89 DPS, sim-verified) [quest] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 16.3 spell_power points (7.35 DPS) | yes | Skycaller's Leather Bracers (252542, -0.46 DPS) [crafted]; Aristocratic Cuffs (12546, -1.28 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.81 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.91 DPS) [crafted]; Ogreseer Fists (11665, -4.22 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -4.42 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 24.6 spell_power points (11.09 DPS) | yes | Skycaller's Leather Waistguard (252476, -0.46 DPS) [crafted]; Satyrmane Sash (17755, -1.93 DPS) [dungeon]; Dawnspire Cord (12466, -2.97 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 29.3 spell_power points (13.21 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -0.51 DPS, sim-verified) [vendor]; Brusslehide Leggings (17751, -0.85 DPS) [dungeon]; Red Mageweave Pants (10009, -3.49 DPS) [crafted] |
| feet | First Sergeant's Crackling Leather Boots (220863) | PvP rank 5 · First Sergeant · Horde [vendor] | 23.6 spell_power points (10.63 DPS) | yes | Skycaller's Leather Boots (252471, -1.29 DPS, sim-verified) [crafted]; Skycaller's Leather Shoes (252532, -2.39 DPS) [crafted]; Stormcloth Boots (10039, -2.40 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 18.2 spell_power points (8.20 DPS) | yes | Cyclopean Band (11824, -1.83 DPS) [dungeon]; Philanthropist's Ring (281635, -2.27 DPS) [quest]; Band of the Unicorn (7553, -2.34 DPS) [world_drop] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 15.6 spell_power points (7.03 DPS) | yes | Philanthropist's Ring (281635, -1.10 DPS) [quest]; Band of the Unicorn (7553, -1.18 DPS) [world_drop]; Cyclopean Band (11824, -1.48 DPS, sim-verified) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.15 DPS) [crafted]; Rune of the Guard Captain (19120, -0.72 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.61 DPS, sim-verified) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (-2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Moonshadow Stave (22458, +0.00 DPS) [quest]; Staff of Jordan (873, -6.92 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.1 spell_power points (6.78 DPS) | yes | Flask of Forest Mojo (19115, -0.45 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.48 DPS) [dungeon]; Enthralled Sphere (11625, -1.28 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Jarkal's Enhancing Necklace; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: First Sergeant's Crackling Leather Boots; finger1: Brainlash; finger2: Advisor's Ring; trinket1: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit

No-known-source sample (15 of 816, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 366.9. Weights run: 3.9s. Verify run: 32.8s. 1992 eligible items had no known source.

Race sweep: windshaper-skyborne 369.8 ± 0.3, tauren 366.9 ± 0.3 (dps); windshaper-skyborne beats the guide race tauren by 2.9 (0.79%) ± 0.5, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.359 ± 0.020, crit=0.445 ± 0.017 per rating point (14 rating = 1%, 6.231 per %), hit=0.759 ± 0.037 per rating point (10 rating = 1%, 7.594 per %), spell_haste=3.436 ± 0.325, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.493 ± 0.002, arcane_power=0.507 ± 0.002, spirit=not significant (0.243 ± 0.186), mp5=not significant (0.877 ± 0.225)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 43.7 spell_power points (20.98 DPS) | yes | Warlord's Dragonhide Helm (231678, -2.70 DPS) [vendor]; Crimson Felt Hat (18727, -3.42 DPS, sim-verified) [dungeon]; Champion's Dragonhide Helm (227186, -3.77 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 19.5 spell_power points (9.34 DPS) | yes | Lady Maye's Pendant (14558, -0.62 DPS, sim-verified) [world_drop]; Archlight Talisman (15856, -1.86 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -2.19 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.9 spell_power points (13.38 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Champion's Dragonhide Spaulders (227184, -0.37 DPS) [pvp]; Argent Shoulders (19059, -0.58 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -1.98 DPS, sim-verified) [world]; Spritecaster Cape (11623, -4.25 DPS) [dungeon]; Shroud of the Exile (15421, -4.40 DPS) [quest] |
| chest | The Postmaster's Tunic (13388) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -10.01 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.6 spell_power points (12.74 DPS) | yes | Runecloth Cuffs (254123, -5.75 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.16 DPS) [crafted]; Runic Leather Bracers (15092, -6.23 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Gloves (231677, -0.32 DPS) [pvp]; Hands of Power (13253, -0.36 DPS) [dungeon]; Blood Guard's Dragonhide Gloves (227187, -3.05 DPS) [vendor] |
| waist | Elunite Cord (272401) | Pix Xizzix [vendor] | 24.6 spell_power points (11.81 DPS) | yes | Ban'thok Sash (11662, -0.69 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -1.07 DPS) [rep]; Frostwolf Cloth Belt (19090, -1.45 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.7 spell_power points (24.31 DPS) | yes | General's Dragonhide Pants (231680, -4.80 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -7.39 DPS) [vendor]; Outrider's Silk Leggings (22747, -7.56 DPS, sim-verified) [rep] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blood Guard's Dragonhide Boots (227188, +0.00 DPS) [pvp]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Omnicast Boots (11822, -6.51 DPS, sim-verified) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Advisor's Ring (19518, +0.00 DPS) [rep]; Blessed Band of Light (272407, +0.00 DPS) [vendor]; Maiden's Circle (13001, -9.28 DPS, sim-verified) [world_drop] |
| finger2 | Eye of Orgrimmar (12545) | The Princess Saved? [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Advisor's Ring (19518, -1.44 DPS) [rep]; Blessed Band of Light (272407, -2.30 DPS, sim-verified) [vendor] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -4.95 DPS, sim-verified) [dungeon]; Smolderweb's Eye (13213, -4.99 DPS) [dungeon]; Serenity Field (272439, -5.28 DPS) [vendor] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -1.08 DPS) [dungeon]; Ironbark Staff (20220, -10.34 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (13.91 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Spirit of Aquementas (11904, -4.02 DPS, sim-verified) [quest]; Tome of the Lost (22253, -4.07 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -1.99 DPS, sim-verified) [world_drop] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: The Postmaster's Tunic; wrist: Dryad's Wrist Bindings; waist: Elunite Cord; legs: Ironfeather Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Eye of Orgrimmar; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch

No-known-source sample (15 of 1992, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 722.9. Weights run: 3.5s. Verify run: 31.5s. 1992 eligible items had no known source.

Race sweep: windshaper-skyborne 730.6 ± 0.6, tauren 722.9 ± 0.6 (dps); windshaper-skyborne beats the guide race tauren by 7.6 (1.06%) ± 0.9, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.272 ± 0.024, crit=0.723 ± 0.024 per rating point (14 rating = 1%, 10.121 per %), hit=1.240 ± 0.057 per rating point (10 rating = 1%, 12.405 per %), spell_haste=6.982 ± 0.455, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.350 ± 0.001, arcane_power=0.650 ± 0.001, spirit=not significant (0.007 ± 0.067), mp5=not significant (0.155 ± 0.215)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 45.5 spell_power points (28.57 DPS) | yes | Warlord's Dragonhide Helm (231678, -4.33 DPS) [vendor]; Champion's Dragonhide Helm (227186, -5.92 DPS) [pvp]; Mask of the Unforgiven (13404, -8.86 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.6 spell_power points (14.81 DPS) | yes | Star of Mystaria (12103, -3.98 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -4.66 DPS) [world_drop]; Emberfury Talisman (12929, -4.68 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.7 spell_power points (19.25 DPS) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Champion's Dragonhide Spaulders (227184, -1.56 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.74 DPS, sim-verified) [world]; Spritecaster Cape (11623, -9.36 DPS) [dungeon]; Shroud of the Exile (15421, -9.93 DPS) [quest] |
| chest | Tunic of the Crescent Moon (22409) | Stratholme: Baron Rivendare [dungeon] | 28.5 spell_power points (17.87 DPS) | yes | Alanna's Embrace (13314, +0.00 DPS) [dungeon]; Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.2 spell_power points (15.21 DPS) | yes | Runecloth Cuffs (254123, -6.60 DPS) [crafted]; Runic Leather Bracers (15092, -7.23 DPS) [crafted]; Blackmist Armguards (12966, -7.42 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Gloves (231677, -1.37 DPS) [pvp]; Hands of Power (13253, -2.91 DPS, sim-verified) [dungeon]; Earth Warder's Gloves (21318, -3.45 DPS) [quest] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.1 spell_power points (17.03 DPS) | yes | Defiler's Cloth Girdle (20163, -0.86 DPS) [rep]; Elunite Cord (272401, -2.89 DPS) [vendor]; Frostwolf Cloth Belt (19090, -4.02 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 53.0 spell_power points (33.29 DPS) | yes | General's Dragonhide Pants (231680, -7.55 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -10.81 DPS) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -12.85 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Crackling Leather Boots (220863) | PvP rank 5 · First Sergeant · Horde [vendor] | 26.1 spell_power points (16.40 DPS) | yes | General's Dragonhide Boots (231682, -0.77 DPS) [pvp]; Omnicast Boots (11822, -1.27 DPS, sim-verified) [dungeon]; Blood Guard's Dragonhide Boots (227188, -1.49 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.88 DPS) [dungeon]; Maiden's Circle (13001, -1.62 DPS) [world_drop]; Naglering (11669, -12.64 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.84 DPS) [dungeon]; Maiden's Circle (13001, -1.58 DPS) [world_drop]; Eye of Orgrimmar (12545, -1.61 DPS) [quest] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of the Beast (13968, -1.54 DPS, sim-verified) [quest]; Smolderweb's Eye (13213, -3.51 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.29 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (-5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.03 DPS) [pvp]; Headmaster's Charge (13937, -11.04 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -17.71 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Brutality (23198, -2.88 DPS, sim-verified) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Tunic of the Crescent Moon; wrist: Dryad's Wrist Bindings; legs: Ironfeather Leggings; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Second Wind; main_hand: Ironbark Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1992, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

