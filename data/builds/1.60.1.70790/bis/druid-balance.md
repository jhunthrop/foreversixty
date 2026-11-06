# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 46.3. Weights run: 5.7s. Verify run: 7.8s. 228 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.523 ± 0.008, crit=0.054 ± 0.001 per rating point (14 rating = 1%, 0.763 per %), hit=0.206 ± 0.006 per rating point (10 rating = 1%, 2.059 per %), spell_haste=0.544 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.801 ± 0.001, arcane_power=0.199 ± 0.001, spirit=-0.092 ± 0.018, mp5=1.008 ± 0.075

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (2.70 DPS) | yes | Pristine Circlet (253949, -0.90 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.91 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.45 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.7 spell_power points (2.92 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.80 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.72 DPS) [crafted]; Slime-encrusted Pads (6461, -2.01 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (46.3 DPS) | yes | Pearl-clasped Cloak (5542, -0.29 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.30 DPS) [crafted]; Caretaker's Cape (20428, -0.30 DPS) [rep] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 10.6 spell_power points (3.19 DPS) | yes | Filigreed Pristine Gown (253901, -0.90 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.91 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.92 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 5.6 spell_power points (1.70 DPS) | yes | Bright Bracers (3647, -1.07 DPS) [world_drop]; Owl Bracers (4796, -1.11 DPS, sim-verified) [vendor]; Inscribed Leather Bracers (3205, -1.22 DPS) [dungeon] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.1 spell_power points (2.13 DPS) | yes | Serpent Gloves (5970, -0.03 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.33 DPS) [world]; Windfelt Gloves (5630, -0.46 DPS) [quest] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.1 spell_power points (1.83 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.16 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.30 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.1 spell_power points (3.95 DPS) | yes | Wisdom's Leather Pants (252503, -0.91 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -1.20 DPS) [crafted]; Totemic Leather Pants (252446, -1.22 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.1 spell_power points (2.73 DPS) | yes | Stormrider's Leather Boots (252443, -0.14 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.74 DPS) [crafted]; Totemic Leather Boots (252442, -0.93 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.11 DPS) | yes | Lavishly Jeweled Ring (1156, -1.17 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.64 DPS) [world_drop]; Furlfeather Ring (286745, -1.81 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.0 spell_power points (1.82 DPS) | yes | Lavishly Jeweled Ring (1156, -1.07 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -1.35 DPS) [world_drop]; Furlfeather Ring (286745, -1.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 29.0 spell_power points (8.73 DPS) | yes | Staff of the Blessed Seer (2271, -1.33 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.52 DPS) [dungeon]; Buzzer Blade (2169, -2.12 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.0 spell_power points (0.90 DPS) | yes | Mathystra Channeling Orb (281291, -0.12 DPS) [quest]; Tear of Grief (5611, -0.33 DPS, sim-verified) [quest]; Spellbinder Orb (15926, -0.43 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome

No-known-source sample (15 of 228, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 5232220113000000-00000000000000000000-0000000000000000)

Set DPS (verified): 81.8. Weights run: 6.3s. Verify run: 9.3s. 400 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.459 ± 0.010, crit=0.123 ± 0.003 per rating point (14 rating = 1%, 1.718 per %), hit=0.246 ± 0.009 per rating point (10 rating = 1%, 2.459 per %), spell_haste=0.795 ± 0.106, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.821 ± 0.001, arcane_power=0.179 ± 0.001, spirit=not significant (0.015 ± 0.010), mp5=0.732 ± 0.102

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (4.60 DPS) | yes | Holy Shroud (2721, -0.35 DPS) [world_drop]; Enchanter's Cowl (4322, -0.54 DPS) [crafted]; Silk Headband (7050, -1.15 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.8 spell_power points (1.06 DPS) | yes | Seal of the Expedition (271920, -0.18 DPS) [rep]; Kaleidoscope Chain (13084, -0.33 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.33 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.1 spell_power points (5.03 DPS) | yes | Death Speaker Mantle (6685, -0.89 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -1.47 DPS) [crafted]; Magician's Mantle (12998, -1.53 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.36 DPS) [crafted]; Caretaker's Cape (19533, -0.36 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.0 spell_power points (5.74 DPS) | yes | Beguiler Robes (7728, -0.13 DPS) [dungeon]; Guardian Armor (4256, -0.53 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.56 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.2 spell_power points (4.29 DPS) | yes | Mindthrust Bracers (1974, -2.41 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.88 DPS) [world_drop]; Nightsky Wristbands (6407, -3.22 DPS) [world_drop] |
| hands | Gloves of Insight (9698) (or Serpent Gloves (5970)) | Return to Vahlarriel [quest] | 7.0 spell_power points (2.68 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Stormrider's Leather Gloves (252498, -0.06 DPS) [crafted]; Truefaith Gloves (7049, -0.24 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.8 spell_power points (4.91 DPS) | yes | Highlander's Cloth Girdle (20099, -0.17 DPS) [rep]; Moss Cinch (6911, -0.31 DPS) [dungeon]; Belt of Arugal (6392, -0.91 DPS) [dungeon] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (5.37 DPS) | yes | Stormrider's Leather Pants (252502, -0.39 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -0.71 DPS) [dungeon]; Guardian Pants (5962, -0.89 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (3.94 DPS) | yes | Acidic Walkers (9454, -0.59 DPS) [dungeon]; Spidersilk Boots (4320, -0.69 DPS, sim-verified) [crafted]; Stormrider's Leather Boots (252443, -0.76 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 8.5 spell_power points (3.24 DPS) | yes | Minor Channeling Ring (1449, -0.98 DPS) [quest]; Electrocutioner Lagnut (9447, -2.05 DPS) [dungeon]; Snake Hoop (6750, -3.43 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.33 DPS, sim-verified) [quest]; Minor Channeling Ring (1449, -0.87 DPS) [quest]; Electrocutioner Lagnut (9447, -1.95 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.56 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.56 DPS) [world]; Glimmering Staff (249392, -3.96 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.8 spell_power points (3.74 DPS) | yes | Eye of Paleth (2943, -2.21 DPS) [quest]; Orb of Souls (249395, -2.51 DPS, sim-verified) [crafted]; Dwarven Tome (279898, -2.58 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Gloves of Insight; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 168.8. Weights run: 6.4s. Verify run: 9.1s. 526 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.675 ± 0.018, crit=0.259 ± 0.011 per rating point (14 rating = 1%, 3.621 per %), hit=0.457 ± 0.025 per rating point (10 rating = 1%, 4.566 per %), spell_haste=2.660 ± 0.338, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.653 ± 0.002, arcane_power=0.347 ± 0.002, spirit=not significant (0.280 ± 0.149), mp5=1.189 ± 0.184

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Big Voodoo Mask (8201, -1.01 DPS) [crafted]; Augural Shroud (2620, -1.29 DPS) [world]; Papal Fez (9431, -4.78 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.4 spell_power points (4.90 DPS) | yes | Triune Amulet (7722, -0.96 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.23 DPS) [quest]; Necklace of Calisea (1714, -2.26 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.8 spell_power points (6.23 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.80 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.27 DPS) [crafted]; Expeditionary's Cape (271922, -0.33 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -1.23 DPS) [crafted]; Dreamweave Vest (10021, -1.52 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -2.55 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.6 spell_power points (4.97 DPS) | yes | Guardian Leather Bracers (4260, -1.00 DPS) [crafted]; Green Whelp Bracers (7386, -1.27 DPS, sim-verified) [crafted]; Arcane Runed Bracers (4744, -1.41 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -4.08 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -5.05 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.32 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Leather Belt (252522, -0.60 DPS) [crafted]; Gilded Cord (254037, -0.64 DPS) [crafted]; Deathmage Sash (10771, -2.73 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.1 spell_power points (8.73 DPS) | yes | Crimson Silk Pantaloons (7062, -2.11 DPS, sim-verified) [crafted]; Warchief Kilt (7760, -2.42 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.50 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.0 spell_power points (10.26 DPS) | yes | Skycaller's Leather Shoes (252532, -3.04 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -4.46 DPS) [crafted]; Gilded Slippers (254001, -5.18 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (5.28 DPS) | yes | Ring of Power Regulation (273030, -1.50 DPS) [dungeon]; Snake Hoop (6750, -1.58 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.12 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 12.6 spell_power points (4.97 DPS) | yes | Snake Hoop (6750, -1.26 DPS) [quest]; Ring of Power Regulation (273030, -1.28 DPS, sim-verified) [dungeon]; Ring of Forlorn Spirits (2043, -1.80 DPS) [quest] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.63 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -13.00 DPS) [dungeon]; Glimmering Staff (249392, -16.49 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 526, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 5232220115101151-00000000000000000000-5400000000000000)

Set DPS (verified): 220.8. Weights run: 6.4s. Verify run: 14.6s. 690 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.014, crit=0.136 ± 0.007 per rating point (14 rating = 1%, 1.904 per %), hit=0.277 ± 0.018 per rating point (10 rating = 1%, 2.769 per %), spell_haste=not significant (0.644 ± 0.268), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.726 ± 0.002, arcane_power=0.274 ± 0.001, spirit=not significant (0.423 ± 0.122), mp5=0.796 ± 0.130

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 34.7 spell_power points (10.78 DPS) | yes | Dreamweave Circlet (10041, -0.96 DPS) [crafted]; Red Mageweave Headband (10033, -1.43 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -2.40 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -0.15 DPS) [dungeon]; Horizon Choker (13085, -0.61 DPS) [world_drop]; Mindburst Medallion (11196, -0.76 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 29.5 spell_power points (9.15 DPS) | yes | Kentic Amice (11624, -1.51 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -2.02 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -2.78 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.38 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -3.06 DPS) [crafted]; Expeditionary's Cape (271921, -3.10 DPS) [rep] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 34.7 spell_power points (10.77 DPS) | yes | Acumen Robes (17775, -1.43 DPS) [quest]; Flamestrider Robes (11747, -1.78 DPS) [dungeon]; Dreamweave Vest (10021, -1.80 DPS) [crafted] |
| wrist | Skycaller's Leather Bracers (252542) | Leatherworking [crafted] | 16.0 spell_power points (4.97 DPS) | yes | Runic Leather Bracers (15092, -0.14 DPS) [crafted]; Nethergeld Cuffs (254061, -0.93 DPS) [crafted]; Mender's Leather Bracers (252543, -1.11 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Gauntlets (252550, -2.40 DPS) [crafted]; Dreamweave Gloves (10019, -2.51 DPS) [crafted]; Bloodfire Talons (12464, -3.84 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Satyrmane Sash (17755, -0.41 DPS) [dungeon]; Ghostweave Cord (254073, -0.69 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -0.95 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 28.6 spell_power points (8.86 DPS) | yes | Knight's Crackling Leather Leggings (220864, -1.96 DPS) [vendor]; Jinxed Hoodoo Kilt (9474, -2.16 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -3.98 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.0 spell_power points (8.37 DPS) | yes | Skycaller's Leather Boots (252471, -1.67 DPS, sim-verified) [crafted]; Stormcloth Boots (10039, -2.28 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -2.37 DPS) [vendor] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 15.2 spell_power points (4.71 DPS) | yes | Philanthropist's Ring (281635, -0.75 DPS) [quest]; Brainlash (6440, -1.47 DPS) [dungeon]; Cyclopean Band (11824, -6.40 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Philanthropist's Ring (281635, -0.07 DPS) [quest]; Cyclopean Band (11824, -0.50 DPS, sim-verified) [dungeon]; Brainlash (6440, -0.79 DPS) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Enriched Seal (271903, -1.98 DPS, sim-verified) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.55 DPS) [quest]; Enriched Seal (271903, -1.43 DPS, sim-verified) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.78 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -2.97 DPS) [world_drop]; Spire of Hakkar (10844, -3.33 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Skycaller's Leather Bracers; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Lorekeeper's Ring; finger2: Band of the Unicorn; trinket1: Thunderbrew's Boot Flask; trinket2: Uther's Strength

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 369.9. Weights run: 5.2s. Verify run: 43.5s. 1755 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.368 ± 0.022, crit=0.452 ± 0.017 per rating point (14 rating = 1%, 6.323 per %), hit=0.759 ± 0.037 per rating point (10 rating = 1%, 7.590 per %), spell_haste=4.771 ± 0.456, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.538 ± 0.002, arcane_power=0.462 ± 0.002, spirit=not significant (0.042 ± 0.224), mp5=not significant (0.610 ± 0.280)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Dragonhide Helm (227192, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -2.14 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.77 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burial Shawl (18681, +0.00 DPS) [dungeon]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -6.76 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.44 DPS) [world]; Spritecaster Cape (11623, -4.86 DPS) [dungeon]; Cloak of Earth and Sky (279270, -5.28 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -9.24 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.2 spell_power points (12.17 DPS) | yes | Runecloth Cuffs (254123, -5.09 DPS) [crafted]; Runic Leather Bracers (15092, -5.57 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.01 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, -0.44 DPS) [dungeon]; Marshal's Dragonhide Gloves (231700, -1.26 DPS) [vendor]; Earth Warder's Gloves (21318, -3.56 DPS) [quest] |
| waist | Elunite Cord (272401) | Pix Xizzix [vendor] | 24.8 spell_power points (11.98 DPS) | yes | Ban'thok Sash (11662, -0.58 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -1.11 DPS) [rep]; Stormpike Cloth Girdle (19094, -1.52 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.9 spell_power points (24.57 DPS) | yes | Sentinel's Silk Leggings (237815, -1.26 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.26 DPS) [vendor]; Marshal's Dragonhide Pants (231697, -5.81 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.4 spell_power points (11.78 DPS) | yes | Knight-Lieutenant's Dragonhide Boots (227194, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.06 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.04 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.69 DPS) [dungeon]; Naglering (11669, -9.64 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (369.9 DPS) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, -0.29 DPS) [dungeon]; Lorekeeper's Ring (19522, -0.72 DPS) [rep] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.49 DPS) [crafted]; Uther's Strength (11302, -2.89 DPS) [world_drop]; Shard of the Splithooves (10659, -4.02 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.89 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -7.32 DPS) [crafted]; Uther's Strength (11302, -7.72 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Sageclaw (20070, -0.43 DPS) [rep]; Ironbark Staff (20069, -8.88 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.99 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Spirit of Aquementas (11904, -3.89 DPS, sim-verified) [quest]; Tome of the Lost (22253, -4.06 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -1.60 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Elunite Cord; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 725.6. Weights run: 5.0s. Verify run: 43.1s. 1755 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.273 ± 0.024, crit=0.745 ± 0.023 per rating point (14 rating = 1%, 10.435 per %), hit=1.231 ± 0.058 per rating point (10 rating = 1%, 12.309 per %), spell_haste=9.707 ± 0.599, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.352 ± 0.001, arcane_power=0.648 ± 0.001, spirit=not significant (0.083 ± 0.069), mp5=not significant (0.097 ± 0.191)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Helm (227192, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helmet (16451, -1.22 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Emberfury Talisman (12929, -2.44 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.0 spell_power points (19.46 DPS) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Lieutenant Commander's Dragonhide Spaulders (227190, -1.18 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.72 DPS, sim-verified) [world]; Cloak of Earth and Sky (279270, -9.00 DPS) [crafted]; Spritecaster Cape (11623, -9.02 DPS) [dungeon] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Breastplate (16452, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.8 spell_power points (15.56 DPS) | yes | Runecloth Cuffs (254123, -6.93 DPS) [crafted]; Runic Leather Bracers (15092, -7.56 DPS) [crafted]; Skycaller's Leather Bracers (252542, -7.81 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Dragonhide Gloves (231700, -0.99 DPS) [vendor]; Earth Warder's Gloves (21318, -3.46 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-verified (725.6 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Elunite Cord (272401, -2.21 DPS) [vendor]; Stormpike Cloth Girdle (19094, -3.36 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) (or Sentinel's Lizardhide Pants (237817)) | Illiyana Moonblaze [vendor] | 55.6 spell_power points (34.93 DPS) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -1.41 DPS) [crafted]; Marshal's Dragonhide Pants (231697, -8.49 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) | PvP reward · Alliance [vendor] | 26.0 spell_power points (16.36 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.14 DPS) [pvp]; Earthen Silk Slippers (254013, -0.91 DPS) [crafted]; Knight-Lieutenant's Dragonhide Boots (227194, -1.24 DPS) [vendor] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.97 DPS) [dungeon]; Maiden's Circle (13001, -1.61 DPS) [world_drop]; Naglering (11669, -14.19 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.82 DPS) [dungeon]; Maiden's Circle (13001, -1.47 DPS) [world_drop]; Songstone of Ironforge (12543, -1.57 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.58 DPS) [crafted]; Uther's Strength (11302, -3.77 DPS) [world_drop]; Ankh of Life (1713, -6.91 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.49 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -6.86 DPS) [crafted]; Uther's Strength (11302, -10.05 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.03 DPS) [pvp]; Staff of Jordan (873, -22.51 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -24.66 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -2.91 DPS, sim-verified) [vendor] |

**New at 60:** head: Mask of the Unforgiven; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Fernpulse Jerkin; wrist: Dryad's Wrist Bindings; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 43.9. Weights run: 5.7s. Verify run: 7.7s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.488 ± 0.003, crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.322 per %), hit=0.079 ± 0.003 per rating point (10 rating = 1%, 0.794 per %), spell_haste=0.429 ± 0.035, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.841 ± 0.002, arcane_power=0.159 ± 0.001, spirit=0.105 ± 0.013, mp5=0.694 ± 0.029

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.96 DPS) | yes | Pristine Circlet (253949, -0.65 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.89 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.11 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.3 spell_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.4 spell_power points (2.05 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.03 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.17 DPS) [crafted]; Slime-encrusted Pads (6461, -1.59 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (43.9 DPS) | yes | Pearl-clasped Cloak (5542, -0.21 DPS) [crafted]; Battle Healer's Cloak (20427, -0.26 DPS) [rep]; Black Whelp Cloak (7283, -0.31 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 10.8 spell_power points (2.34 DPS) | yes | Filigreed Pristine Gown (253901, -0.65 DPS) [crafted]; Totemic Leather Armor (252435, -0.72 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.89 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 4.5 spell_power points (0.98 DPS) | yes | Owl Bracers (4796, -0.45 DPS) [vendor]; Owlbeard Bracers (16981, -0.55 DPS) [quest]; Tabitha's Cuffs (251486, -1.44 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.52 DPS) | yes | Stormrider's Leather Gloves (252498, -0.01 DPS) [crafted]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Pristine Gloves (253913, -0.33 DPS) [crafted] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.0 spell_power points (1.30 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.22 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.3 spell_power points (2.91 DPS) | yes | Filigreed Pristine Leggings (253937, -0.87 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.89 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.96 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.0 spell_power points (1.95 DPS) | yes | Stormrider's Leather Boots (252443, -0.11 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.55 DPS) [crafted]; Totemic Leather Boots (252442, -0.64 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.4 spell_power points (1.39 DPS) | yes | Loop of Sacrifice (281673, -0.86 DPS) [quest]; Volcanic Rock Ring (12053, -1.07 DPS) [world_drop]; Furlfeather Ring (286745, -1.17 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.9 spell_power points (0.64 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Volcanic Rock Ring (12053, -0.32 DPS) [world_drop]; Furlfeather Ring (286745, -0.42 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 29.5 spell_power points (6.42 DPS) | yes | Crescent Staff (6505, -0.02 DPS) [quest]; Staff of the Blessed Seer (2271, -0.10 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.53 DPS) [quest] |
| off_hand | Seer's Fine Stein (7608) | World drop [world_drop] | 1.6 spell_power points (0.34 DPS) | yes | Spellbinder Orb (15926, -0.02 DPS) [dungeon]; Tork Wrench (11855, -0.13 DPS) [quest]; Veildust Medicine Bag (15866, -0.21 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Seer's Fine Stein

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance

### Band 30 (tauren, 5232220113000000-00000000000000000000-0000000000000000)

Set DPS (verified): 81.5. Weights run: 6.3s. Verify run: 9.4s. 401 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.459 ± 0.010, crit=0.123 ± 0.003 per rating point (14 rating = 1%, 1.718 per %), hit=0.246 ± 0.009 per rating point (10 rating = 1%, 2.459 per %), spell_haste=0.795 ± 0.106, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.821 ± 0.001, arcane_power=0.179 ± 0.001, spirit=not significant (0.015 ± 0.010), mp5=0.732 ± 0.102

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (4.60 DPS) | yes | Holy Shroud (2721, -0.35 DPS) [world_drop]; Enchanter's Cowl (4322, -0.54 DPS) [crafted]; Silk Headband (7050, -1.15 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.8 spell_power points (1.06 DPS) | yes | Darkspear Warding Pendant (272075, -0.18 DPS) [vendor]; Kaleidoscope Chain (13084, -0.33 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.33 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.1 spell_power points (5.03 DPS) | yes | Death Speaker Mantle (6685, -0.45 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -1.47 DPS) [crafted]; Magician's Mantle (12998, -1.53 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.02 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.38 DPS) [crafted]; Battle Healer's Cloak (19529, -0.38 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.0 spell_power points (5.74 DPS) | yes | Beguiler Robes (7728, -0.13 DPS) [dungeon]; Guardian Armor (4256, -0.53 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.56 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.2 spell_power points (4.29 DPS) | yes | Mindthrust Bracers (1974, -2.44 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.88 DPS) [world_drop]; Tabitha's Cuffs (251486, -3.14 DPS) [quest] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 9.9 spell_power points (3.80 DPS) | yes | Jutebraid Gloves (10654, -0.41 DPS, sim-verified) [quest]; Serpent Gloves (5970, -1.12 DPS) [dungeon]; Stormrider's Leather Gloves (252498, -1.18 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.8 spell_power points (4.91 DPS) | yes | Defiler's Cloth Girdle (20164, -0.17 DPS) [rep]; Moss Cinch (6911, -0.31 DPS) [dungeon]; Warsong Sash (16975, -0.70 DPS) [quest] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (5.37 DPS) | yes | Stormrider's Leather Pants (252502, -0.45 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.71 DPS) [dungeon]; Guardian Pants (5962, -0.89 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (3.94 DPS) | yes | Acidic Walkers (9454, -0.59 DPS) [dungeon]; Spidersilk Boots (4320, -0.71 DPS, sim-verified) [crafted]; Stormrider's Leather Boots (252443, -0.76 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 8.5 spell_power points (3.24 DPS) | yes | Solid Gold Band (279845, -1.68 DPS) [quest]; Electrocutioner Lagnut (9447, -2.05 DPS) [dungeon]; Snake Hoop (6750, -3.59 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.40 DPS, sim-verified) [quest]; Solid Gold Band (279845, -1.58 DPS) [quest]; Electrocutioner Lagnut (9447, -1.95 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.56 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.56 DPS) [world]; Glimmering Staff (249392, -3.96 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.8 spell_power points (3.74 DPS) | yes | Witch's Finger (16887, -2.33 DPS) [quest]; Satyr's Rod (15962, -2.67 DPS) [world_drop]; Orb of Souls (249395, -2.78 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 167.0. Weights run: 6.4s. Verify run: 9.3s. 525 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.675 ± 0.018, crit=0.259 ± 0.011 per rating point (14 rating = 1%, 3.621 per %), hit=0.457 ± 0.025 per rating point (10 rating = 1%, 4.566 per %), spell_haste=2.660 ± 0.338, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.653 ± 0.002, arcane_power=0.347 ± 0.002, spirit=not significant (0.280 ± 0.149), mp5=1.189 ± 0.184

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Big Voodoo Mask (8201, -1.01 DPS) [crafted]; Augural Shroud (2620, -1.29 DPS) [world]; Papal Fez (9431, -4.06 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.4 spell_power points (4.90 DPS) | yes | Triune Amulet (7722, -0.96 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.23 DPS) [quest]; Necklace of Calisea (1714, -2.26 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.8 spell_power points (6.23 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.80 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.27 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.33 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -1.23 DPS) [crafted]; Dreamweave Vest (10021, -1.34 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -2.55 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.6 spell_power points (4.97 DPS) | yes | Guardian Leather Bracers (4260, -1.00 DPS) [crafted]; Radiant Silver Bracers (4545, -1.25 DPS) [quest]; Green Whelp Bracers (7386, -1.31 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -4.66 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -5.05 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.32 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Leather Belt (252522, -0.60 DPS) [crafted]; Gilded Cord (254037, -0.64 DPS) [crafted]; Deathmage Sash (10771, -2.59 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.1 spell_power points (8.73 DPS) | yes | Warchief Kilt (7760, -2.42 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.50 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.91 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.0 spell_power points (10.26 DPS) | yes | Skycaller's Leather Shoes (252532, -3.74 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -4.46 DPS) [crafted]; Gilded Slippers (254001, -5.18 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (5.28 DPS) | yes | Ring of Power Regulation (273030, -1.50 DPS) [dungeon]; Snake Hoop (6750, -1.58 DPS) [quest]; Darkspear Signet (272070, -2.47 DPS) [vendor] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 12.6 spell_power points (4.97 DPS) | yes | Snake Hoop (6750, -1.26 DPS) [quest]; Ring of Power Regulation (273030, -1.27 DPS, sim-verified) [dungeon]; Darkspear Signet (272070, -2.15 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.49 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -13.00 DPS) [dungeon]; Glimmering Staff (249392, -16.49 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 5232220115101151-00000000000000000000-5400000000000000)

Set DPS (verified): 217.4. Weights run: 6.4s. Verify run: 14.7s. 689 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.014, crit=0.136 ± 0.007 per rating point (14 rating = 1%, 1.904 per %), hit=0.277 ± 0.018 per rating point (10 rating = 1%, 2.769 per %), spell_haste=not significant (0.644 ± 0.268), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.726 ± 0.002, arcane_power=0.274 ± 0.001, spirit=not significant (0.423 ± 0.122), mp5=0.796 ± 0.130

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 34.7 spell_power points (10.78 DPS) | yes | Red Mageweave Headband (10033, -1.43 DPS) [crafted]; Dreamweave Circlet (10041, -1.51 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme Plus (15999, -2.40 DPS) [crafted] |
| neck | Mindburst Medallion (11196) | Broken Alliances [quest] | 11.4 spell_power points (3.55 DPS) | yes | Glowing Eye of Mordresh (10769, -0.00 DPS) [dungeon]; Triune Amulet (7722, -0.15 DPS) [dungeon]; Horizon Choker (13085, -0.61 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 29.5 spell_power points (9.15 DPS) | yes | Rotgrip Mantle (17732, -2.02 DPS) [dungeon]; Kentic Amice (11624, -2.15 DPS, sim-verified) [dungeon]; Skycaller's Leather Shoulders (252537, -2.78 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.89 DPS) [quest]; Runecloth Cloak (13860, -2.00 DPS) [crafted]; Big Voodoo Cloak (8216, -3.06 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 34.7 spell_power points (10.77 DPS) | yes | Acumen Robes (17775, -1.43 DPS) [quest]; Flamestrider Robes (11747, -1.78 DPS) [dungeon]; Dreamweave Vest (10021, -1.80 DPS) [crafted] |
| wrist | Skycaller's Leather Bracers (252542) | Leatherworking [crafted] | 16.0 spell_power points (4.97 DPS) | yes | Runic Leather Bracers (15092, -0.14 DPS) [crafted]; Nethergeld Cuffs (254061, -0.93 DPS) [crafted]; Mender's Leather Bracers (252543, -1.11 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Gauntlets (252550, -2.40 DPS) [crafted]; Dreamweave Gloves (10019, -2.51 DPS) [crafted]; Bloodfire Talons (12464, -3.63 DPS, sim-verified) [dungeon] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, -0.55 DPS) [dungeon]; Satyrmane Sash (17755, -0.97 DPS) [dungeon]; Ghostweave Cord (254073, -1.24 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 28.6 spell_power points (8.86 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -1.96 DPS) [vendor]; Jinxed Hoodoo Kilt (9474, -2.16 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -5.11 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.0 spell_power points (8.37 DPS) | yes | Skycaller's Leather Boots (252471, -2.13 DPS, sim-verified) [crafted]; Stormcloth Boots (10039, -2.28 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -2.37 DPS) [vendor] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 15.2 spell_power points (4.71 DPS) | yes | Philanthropist's Ring (281635, -0.75 DPS) [quest]; Brainlash (6440, -1.47 DPS) [dungeon]; Cyclopean Band (11824, -6.69 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | sim-verified (217.4 DPS) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.07 DPS) [quest]; Brainlash (6440, -0.79 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.55 DPS) [quest]; Rune of the Guard Captain (19120, -0.73 DPS, sim-verified) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -3.33 DPS) [world]; Kindling Stave (11750, -3.39 DPS) [dungeon]; Glowing Brightwood Staff (812, -6.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Mindburst Medallion; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Skycaller's Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger1: Advisor's Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength

No-known-source sample (15 of 689, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 367.4. Weights run: 5.2s. Verify run: 43.1s. 1759 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.368 ± 0.022, crit=0.452 ± 0.017 per rating point (14 rating = 1%, 6.323 per %), hit=0.759 ± 0.037 per rating point (10 rating = 1%, 7.590 per %), spell_haste=4.771 ± 0.456, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.538 ± 0.002, arcane_power=0.462 ± 0.002, spirit=not significant (0.042 ± 0.224), mp5=not significant (0.610 ± 0.280)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Champion's Dragonhide Helm (227186, +0.00 DPS) [pvp]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -2.14 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.77 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burial Shawl (18681, +0.00 DPS) [dungeon]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -5.38 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.44 DPS) [world]; Spritecaster Cape (11623, -4.86 DPS) [dungeon]; Cloak of Earth and Sky (279270, -5.28 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -8.93 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.2 spell_power points (12.17 DPS) | yes | Runecloth Cuffs (254123, -5.09 DPS) [crafted]; Runic Leather Bracers (15092, -5.57 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.01 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, -0.44 DPS) [dungeon]; General's Dragonhide Gloves (231677, -1.26 DPS) [pvp]; Earth Warder's Gloves (21318, -3.56 DPS) [quest] |
| waist | Elunite Cord (272401) | Pix Xizzix [vendor] | 24.8 spell_power points (11.98 DPS) | yes | Ban'thok Sash (11662, -0.58 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -1.11 DPS) [rep]; Frostwolf Cloth Belt (19090, -1.52 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.9 spell_power points (24.57 DPS) | yes | Skyshroud Leggings (13170, -5.16 DPS, sim-verified) [dungeon]; General's Dragonhide Pants (231680, -5.81 DPS) [vendor]; Outrider's Silk Leggings (22747, -7.49 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Dragonhide Boots (227188, +0.00 DPS) [pvp]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.86 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.04 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.69 DPS) [dungeon]; Naglering (11669, -9.49 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (367.4 DPS) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, -0.29 DPS) [dungeon]; Advisor's Ring (19518, -0.72 DPS) [rep] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.49 DPS) [crafted]; Uther's Strength (11302, -2.89 DPS) [world_drop]; Rune of the Guard Captain (19120, -3.22 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.86 DPS) [vendor]; Frozen Heart of the Mountain (249469, -7.32 DPS) [crafted]; Shard of the Splithooves (10659, -8.73 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Mindfang (20214, -0.43 DPS) [rep]; Ironbark Staff (20220, -8.47 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (13.99 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Spirit of Aquementas (11904, -3.84 DPS, sim-verified) [quest]; Tome of the Lost (22253, -4.06 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -1.60 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Elunite Cord; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 717.2. Weights run: 5.0s. Verify run: 42.5s. 1759 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.273 ± 0.024, crit=0.745 ± 0.023 per rating point (14 rating = 1%, 10.435 per %), hit=1.231 ± 0.058 per rating point (10 rating = 1%, 12.309 per %), spell_haste=9.707 ± 0.599, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.352 ± 0.001, arcane_power=0.648 ± 0.001, spirit=not significant (0.083 ± 0.069), mp5=not significant (0.097 ± 0.191)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 35.1 spell_power points (22.02 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Champion's Dragonhide Helm (227186, +0.00 DPS) [pvp]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Emberfury Talisman (12929, -2.44 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.0 spell_power points (19.46 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Champion's Dragonhide Spaulders (227184, -1.18 DPS) [pvp]; Argent Shoulders (19059, -4.34 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.07 DPS, sim-verified) [world]; Cloak of Earth and Sky (279270, -9.00 DPS) [crafted]; Spritecaster Cape (11623, -9.02 DPS) [dungeon] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Hauberk (16549, +0.00 DPS) [vendor]; Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.8 spell_power points (15.56 DPS) | yes | Runecloth Cuffs (254123, -6.93 DPS) [crafted]; Runic Leather Bracers (15092, -7.56 DPS) [crafted]; Skycaller's Leather Bracers (252542, -7.81 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Dragonhide Gloves (231677, -0.99 DPS) [pvp]; Earth Warder's Gloves (21318, -3.46 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-verified (717.2 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Elunite Cord (272401, -2.21 DPS) [vendor]; Frostwolf Cloth Belt (19090, -3.36 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 53.3 spell_power points (33.51 DPS) | yes | General's Dragonhide Pants (231680, -7.07 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -10.43 DPS) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -12.93 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Crackling Leather Boots (220863) | PvP reward · Horde [vendor] | 26.0 spell_power points (16.36 DPS) | yes | General's Dragonhide Boots (231682, -0.14 DPS) [pvp]; Earthen Silk Slippers (254013, -0.91 DPS) [crafted]; Blood Guard's Dragonhide Boots (227188, -1.24 DPS) [pvp] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.97 DPS) [dungeon]; Maiden's Circle (13001, -1.61 DPS) [world_drop]; Naglering (11669, -13.81 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.82 DPS) [dungeon]; Maiden's Circle (13001, -1.47 DPS) [world_drop]; Eye of Orgrimmar (12545, -1.57 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.58 DPS) [crafted]; Rune of the Guard Captain (19120, -2.13 DPS) [quest]; Uther's Strength (11302, -3.77 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.39 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -6.86 DPS) [crafted]; Rune of the Guard Captain (19120, -8.41 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.03 DPS) [pvp]; Whiteout Staff (19101, -21.73 DPS) [rep]; Staff of Jordan (873, -21.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -2.87 DPS, sim-verified) [vendor] |

**New at 60:** head: Mask of the Unforgiven; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Fernpulse Jerkin; wrist: Dryad's Wrist Bindings; waist: Defiler's Cloth Girdle; legs: Ironfeather Leggings; feet: First Sergeant's Crackling Leather Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

