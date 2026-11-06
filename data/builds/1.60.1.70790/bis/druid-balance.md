# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 46.3. Weights run: 3.3s. Verify run: 4.8s. 228 eligible items had no known source.

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

No-known-source sample (15 of 228, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 5232220113000000-00000000000000000000-0000000000000000)

Set DPS (verified): 82.2. Weights run: 3.4s. Verify run: 5.4s. 400 eligible items had no known source.

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
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 8.8 spell_power points (3.36 DPS) | yes | Snake Hoop (6750, -0.13 DPS) [quest]; Minor Channeling Ring (1449, -1.08 DPS) [quest]; Electrocutioner Lagnut (9447, -2.21 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.7 spell_power points (3.32 DPS) | yes | Snake Hoop (6750, -0.09 DPS) [quest]; Minor Channeling Ring (1449, -1.04 DPS) [quest]; Electrocutioner Lagnut (9447, -2.17 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.53 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.53 DPS) [world]; Glimmering Staff (249392, -4.22 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.9 spell_power points (3.80 DPS) | yes | Orb of Souls (249395, -2.27 DPS) [crafted]; Eye of Paleth (2943, -2.63 DPS, sim-verified) [quest]; Dwarven Tome (279898, -2.65 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Gloves of Insight; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 169.4. Weights run: 3.7s. Verify run: 6.9s. 526 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.579 ± 0.017, crit=0.233 ± 0.010 per rating point (14 rating = 1%, 3.267 per %), hit=0.448 ± 0.023 per rating point (10 rating = 1%, 4.482 per %), spell_haste=1.506 ± 0.203, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.618 ± 0.002, arcane_power=0.382 ± 0.002, spirit=0.571 ± 0.133, mp5=1.174 ± 0.171

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Whitemane's Chapeau (7720, -0.55 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.32 DPS) [dungeon]; Papal Fez (9431, -3.64 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -0.51 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.46 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.38 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Mantle of Doan (7712, -0.89 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.62 DPS) [crafted]; Guardian Cloak (5965, -0.62 DPS) [crafted]; Caretaker's Cape (19532, -0.64 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -1.15 DPS) [crafted]; Dreamweave Vest (10021, -1.65 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -2.98 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.5 spell_power points (5.10 DPS) | yes | Green Whelp Bracers (7386, -0.95 DPS) [crafted]; Guardian Leather Bracers (4260, -1.24 DPS) [crafted]; Arcane Runed Bracers (4744, -1.43 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.98 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -5.45 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.69 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.3 spell_power points (6.64 DPS) | yes | Skycaller's Leather Belt (252522, -0.23 DPS) [crafted]; Deathmage Sash (10771, -0.26 DPS) [dungeon]; Gilded Cord (254037, -1.77 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.9 spell_power points (8.53 DPS) | yes | Stoneweaver Leggings (9407, -1.01 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.55 DPS) [dungeon]; Warchief Kilt (7760, -2.87 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.0 spell_power points (11.40 DPS) | yes | Skycaller's Leather Shoes (252532, -3.69 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -4.86 DPS) [crafted]; Gilded Slippers (254001, -5.97 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (5.25 DPS) | yes | Ring of Power Regulation (273030, -1.37 DPS) [dungeon]; Snake Hoop (6750, -1.63 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.99 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 12.5 spell_power points (5.10 DPS) | yes | Ring of Power Regulation (273030, -1.27 DPS, sim-verified) [dungeon]; Snake Hoop (6750, -1.48 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.84 DPS) [quest] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.87 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -13.71 DPS) [dungeon]; Swinetusk Shank (6691, -16.99 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 526, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 5232220115101151-00000000000000000000-5400000000000000)

Set DPS (verified): 221.0. Weights run: 3.6s. Verify run: 10.1s. 690 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.613 ± 0.012, crit=0.129 ± 0.006 per rating point (14 rating = 1%, 1.810 per %), hit=0.283 ± 0.015 per rating point (10 rating = 1%, 2.831 per %), spell_haste=0.493 ± 0.078, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.737 ± 0.002, arcane_power=0.263 ± 0.001, spirit=0.484 ± 0.073, mp5=0.804 ± 0.086

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 36.1 spell_power points (10.82 DPS) | yes | Dreamweave Circlet (10041, -0.94 DPS) [crafted]; Red Mageweave Headband (10033, -1.44 DPS) [crafted]; Knight-Lieutenant's Crackling Leather Helmet (220872, -2.64 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -0.15 DPS) [dungeon]; Horizon Choker (13085, -0.42 DPS) [world_drop]; Mindburst Medallion (11196, -0.76 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 31.1 spell_power points (9.34 DPS) | yes | Kentic Amice (11624, -1.51 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -2.13 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -2.91 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.38 DPS, sim-verified) [crafted]; Expeditionary's Cape (271921, -2.87 DPS) [rep]; Big Voodoo Cloak (8216, -3.02 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 36.8 spell_power points (11.03 DPS) | yes | Acumen Robes (17775, -1.65 DPS) [quest]; Dreamweave Vest (10021, -1.94 DPS) [crafted]; Flamestrider Robes (11747, -1.97 DPS) [dungeon] |
| wrist | Skycaller's Leather Bracers (252542) | Leatherworking [crafted] | 16.7 spell_power points (5.01 DPS) | yes | Runic Leather Bracers (15092, -0.17 DPS) [crafted]; Nethergeld Cuffs (254061, -0.90 DPS) [crafted]; Mender's Leather Bracers (252543, -1.05 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Gauntlets (252550, -2.08 DPS) [crafted]; Dreamweave Gloves (10019, -2.26 DPS) [crafted]; Bloodfire Talons (12464, -3.84 DPS, sim-verified) [dungeon] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, -0.53 DPS) [dungeon]; Satyrmane Sash (17755, -0.97 DPS) [dungeon]; Tooled Leather Belt (252544, -1.12 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 29.1 spell_power points (8.74 DPS) | yes | Jinxed Hoodoo Kilt (9474, -1.74 DPS) [dungeon]; Knight's Crackling Leather Leggings (220864, -1.91 DPS) [vendor]; Senior Designer's Pantaloons (11841, -3.98 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.4 spell_power points (8.22 DPS) | yes | Skycaller's Leather Boots (252471, -1.67 DPS, sim-verified) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -2.23 DPS) [vendor]; Stormcloth Boots (10039, -2.28 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.2 spell_power points (4.57 DPS) | yes | Philanthropist's Ring (281635, -0.65 DPS) [quest]; Band of the Unicorn (7553, -0.67 DPS) [world_drop]; Brainlash (6440, -1.08 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 15.2 spell_power points (4.57 DPS) | yes | Band of the Unicorn (7553, -0.67 DPS) [world_drop]; Brainlash (6440, -1.08 DPS) [dungeon]; Philanthropist's Ring (281635, -1.30 DPS, sim-verified) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Enriched Seal (271903, -1.98 DPS, sim-verified) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.06 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest]; Enriched Seal (271903, -1.43 DPS, sim-verified) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.78 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -2.76 DPS) [world_drop]; Spire of Hakkar (10844, -3.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Skycaller's Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Lorekeeper's Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Uther's Strength

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 371.3. Weights run: 3.2s. Verify run: 34.4s. 1755 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.353 ± 0.021, crit=0.467 ± 0.018 per rating point (14 rating = 1%, 6.539 per %), hit=0.745 ± 0.038 per rating point (10 rating = 1%, 7.453 per %), spell_haste=3.642 ± 0.354, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.480 ± 0.002, arcane_power=0.520 ± 0.002, spirit=not significant (0.293 ± 0.191), mp5=not significant (0.600 ± 0.242)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 35.4 spell_power points (16.72 DPS) | yes | Lieutenant Commander's Dragonhide Helm (227192, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -0.13 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -1.21 DPS) [dungeon]; Triune Amulet (7722, -2.58 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Spaulders (227190, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -8.45 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Earth and Sky (279270, -3.38 DPS) [crafted]; Crystalline Threaded Cape (20697, -3.86 DPS, sim-verified) [world]; Spritecaster Cape (11623, -3.96 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 33.4 spell_power points (15.77 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Alanna's Embrace (13314, -10.05 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.9 spell_power points (12.67 DPS) | yes | Runecloth Cuffs (254123, -5.82 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.10 DPS) [crafted]; Runic Leather Bracers (15092, -6.29 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Dragonhide Gloves (231700, -0.11 DPS) [vendor]; Knight-Lieutenant's Dragonhide Gloves (227193, -2.88 DPS) [pvp] |
| waist | Elunite Cord (272401) | Pix Xizzix [vendor] | 24.5 spell_power points (11.54 DPS) | yes | Highlander's Cloth Girdle (20047, -0.85 DPS) [rep]; Stormpike Cloth Girdle (19094, -1.38 DPS) [rep]; Ban'thok Sash (11662, -1.67 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.9 spell_power points (24.00 DPS) | yes | Sentinel's Silk Leggings (237815, -1.08 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.08 DPS) [vendor]; Marshal's Dragonhide Pants (231697, -4.47 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.1 spell_power points (12.29 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.17 DPS) [vendor]; Venomspew Footpads (275606, -0.81 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.28 DPS) [quest]; Don Mauricio's Band of Domination (22433, -2.21 DPS) [dungeon]; Naglering (11669, -10.48 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.06 DPS) [dungeon]; Lorekeeper's Ring (19522, -1.60 DPS) [rep] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.92 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -7.21 DPS) [crafted]; Uther's Strength (11302, -7.55 DPS) [world_drop] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.50 DPS) [crafted]; Uther's Strength (11302, -2.83 DPS) [world_drop]; Second Wind (11819, -6.59 DPS, sim-verified) [dungeon] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Sageclaw (20070, -1.40 DPS) [rep]; Ironbark Staff (20069, -9.20 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.68 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Spirit of Aquementas (11904, -3.92 DPS, sim-verified) [quest]; Tome of the Lost (22253, -4.02 DPS) [dungeon] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -2.37 DPS, sim-verified) [vendor] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Elunite Cord; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Burst of Knowledge; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Idol of the Moon

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 725.5. Weights run: 3.1s. Verify run: 32.1s. 1755 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.268 ± 0.023, crit=0.734 ± 0.024 per rating point (14 rating = 1%, 10.273 per %), hit=1.234 ± 0.057 per rating point (10 rating = 1%, 12.341 per %), spell_haste=7.244 ± 0.471, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.348 ± 0.001, arcane_power=0.652 ± 0.001, spirit=not significant (0.019 ± 0.045), mp5=not significant (0.114 ± 0.179)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Helm (227192, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helmet (16451, -1.59 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Emberfury Talisman (12929, -2.71 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.8 spell_power points (19.33 DPS) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Lieutenant Commander's Dragonhide Spaulders (227190, -1.50 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.72 DPS, sim-verified) [world]; Spritecaster Cape (11623, -9.28 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -9.55 DPS) [dungeon] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Breastplate (16452, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.3 spell_power points (15.25 DPS) | yes | Runecloth Cuffs (254123, -6.66 DPS) [crafted]; Runic Leather Bracers (15092, -7.28 DPS) [crafted]; Blackmist Armguards (12966, -7.50 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Dragonhide Gloves (231700, -1.33 DPS) [vendor]; Earth Warder's Gloves (21318, -3.44 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Elunite Cord (272401, -2.16 DPS) [vendor]; Stormpike Cloth Girdle (19094, -3.27 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) (or Sentinel's Lizardhide Pants (237817)) | Illiyana Moonblaze [vendor] | 55.2 spell_power points (34.70 DPS) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -1.34 DPS) [crafted]; Marshal's Dragonhide Pants (231697, -8.80 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) | PvP reward · Alliance [vendor] | 26.0 spell_power points (16.35 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.66 DPS) [pvp]; Earthen Silk Slippers (254013, -1.19 DPS) [crafted]; Knight-Lieutenant's Dragonhide Boots (227194, -1.42 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.84 DPS) [dungeon]; Maiden's Circle (13001, -1.65 DPS) [world_drop]; Naglering (11669, -14.07 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.76 DPS) [dungeon]; Maiden's Circle (13001, -1.57 DPS) [world_drop]; Songstone of Ironforge (12543, -1.59 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.56 DPS) [crafted]; Uther's Strength (11302, -3.77 DPS) [world_drop]; Shard of the Splithooves (10659, -7.11 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.49 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -6.84 DPS) [crafted]; Uther's Strength (11302, -10.05 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.05 DPS) [pvp]; Staff of Jordan (873, -22.51 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -24.55 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -2.91 DPS, sim-verified) [vendor] |

**New at 60:** head: Mask of the Unforgiven; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Fernpulse Jerkin; wrist: Dryad's Wrist Bindings; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 44.0. Weights run: 3.3s. Verify run: 5.3s. 224 eligible items had no known source.

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

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance

### Band 30 (tauren, 5232220113000000-00000000000000000000-0000000000000000)

Set DPS (verified): 81.8. Weights run: 3.4s. Verify run: 5.4s. 401 eligible items had no known source.

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
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 8.8 spell_power points (3.36 DPS) | yes | Snake Hoop (6750, -0.13 DPS) [quest]; Solid Gold Band (279845, -1.83 DPS) [quest]; Darkspear Signet (272071, -1.99 DPS) [vendor] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.7 spell_power points (3.32 DPS) | yes | Snake Hoop (6750, -0.51 DPS, sim-verified) [quest]; Solid Gold Band (279845, -1.79 DPS) [quest]; Darkspear Signet (272071, -1.95 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.53 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.53 DPS) [world]; Glimmering Staff (249392, -4.11 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.9 spell_power points (3.80 DPS) | yes | Witch's Finger (16887, -2.30 DPS) [quest]; Satyr's Rod (15962, -2.68 DPS) [world_drop]; Orb of Souls (249395, -3.06 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 167.3. Weights run: 3.7s. Verify run: 6.7s. 525 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.634 ± 0.011, crit=0.100 ± 0.005 per rating point (14 rating = 1%, 1.404 per %), hit=0.197 ± 0.011 per rating point (10 rating = 1%, 1.971 per %), spell_haste=not significant (0.383 ± 0.104), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.709 ± 0.002, arcane_power=0.291 ± 0.001, spirit=0.342 ± 0.076, mp5=0.711 ± 0.070

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Big Voodoo Mask (8201, -0.94 DPS) [crafted]; Whitemane's Chapeau (7720, -1.07 DPS) [dungeon]; Papal Fez (9431, -4.03 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.1 spell_power points (3.62 DPS) | yes | Triune Amulet (7722, -0.50 DPS) [dungeon]; Necklace of Calisea (1714, -1.57 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.72 DPS) [quest] |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodmage Mantle (7684, -0.08 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.34 DPS, sim-verified) [dungeon]; Mantle of Doan (7712, -0.43 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.04 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.25 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.95 DPS) [crafted]; Elemental Raiment (9434, -1.95 DPS) [world_drop]; Dreamweave Vest (10021, -1.96 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Leather Bracers (4260, -0.40 DPS) [crafted]; Radiant Silver Bracers (4545, -0.62 DPS) [quest]; Green Whelp Bracers (7386, -1.99 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, -4.12 DPS) [crafted]; Skycaller's Leather Gloves (252529, -4.31 DPS) [crafted]; Dreamweave Gloves (10019, -4.35 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.5 spell_power points (4.96 DPS) | yes | Skycaller's Leather Belt (252522, -0.41 DPS) [crafted]; Gilded Cord (254037, -0.42 DPS) [crafted]; Deathmage Sash (10771, -1.91 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.6 spell_power points (6.48 DPS) | yes | Crimson Silk Pantaloons (7062, -1.61 DPS) [crafted]; Stoneweaver Leggings (9407, -1.63 DPS) [dungeon]; Warchief Kilt (7760, -2.93 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.4 spell_power points (7.91 DPS) | yes | Mender's Leather Shoes (252533, -3.47 DPS) [crafted]; Gilded Slippers (254001, -4.07 DPS) [crafted]; Skycaller's Leather Shoes (252532, -4.44 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (3.95 DPS) | yes | Snake Hoop (6750, -1.20 DPS) [quest]; Ring of Power Regulation (273030, -1.51 DPS) [dungeon]; Reedknot Ring (9622, -1.85 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 11.1 spell_power points (3.34 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Ring of Power Regulation (273030, -0.90 DPS) [dungeon]; Reedknot Ring (9622, -1.24 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -8.68 DPS) [dungeon]; Illusionary Rod (7713, -9.05 DPS, sim-verified) [dungeon]; Swinetusk Shank (6691, -11.29 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Green Silken Shoulders; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 5232220115101151-00000000000000000000-5400000000000000)

Set DPS (verified): 217.2. Weights run: 3.6s. Verify run: 10.4s. 689 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.613 ± 0.012, crit=0.129 ± 0.006 per rating point (14 rating = 1%, 1.810 per %), hit=0.283 ± 0.015 per rating point (10 rating = 1%, 2.831 per %), spell_haste=0.493 ± 0.078, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.737 ± 0.002, arcane_power=0.263 ± 0.001, spirit=0.484 ± 0.073, mp5=0.804 ± 0.086

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 36.1 spell_power points (10.82 DPS) | yes | Red Mageweave Headband (10033, -1.44 DPS) [crafted]; Dreamweave Circlet (10041, -1.51 DPS, sim-verified) [crafted]; Blood Guard's Crackling Leather Helmet (220873, -2.64 DPS) [vendor] |
| neck | Mindburst Medallion (11196) | Broken Alliances [quest] | 12.1 spell_power points (3.63 DPS) | yes | Glowing Eye of Mordresh (10769, -0.06 DPS) [dungeon]; Triune Amulet (7722, -0.21 DPS) [dungeon]; Horizon Choker (13085, -0.47 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 31.1 spell_power points (9.34 DPS) | yes | Rotgrip Mantle (17732, -2.13 DPS) [dungeon]; Kentic Amice (11624, -2.15 DPS, sim-verified) [dungeon]; Skycaller's Leather Shoulders (252537, -2.91 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.92 DPS) [quest]; Runecloth Cloak (13860, -2.00 DPS) [crafted]; Darkspear Raider's Cloak (272076, -2.87 DPS) [vendor] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 36.8 spell_power points (11.03 DPS) | yes | Acumen Robes (17775, -1.65 DPS) [quest]; Dreamweave Vest (10021, -1.94 DPS) [crafted]; Flamestrider Robes (11747, -1.97 DPS) [dungeon] |
| wrist | Skycaller's Leather Bracers (252542) | Leatherworking [crafted] | 16.7 spell_power points (5.01 DPS) | yes | Runic Leather Bracers (15092, -0.17 DPS) [crafted]; Nethergeld Cuffs (254061, -0.90 DPS) [crafted]; Mender's Leather Bracers (252543, -1.05 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Gauntlets (252550, -2.08 DPS) [crafted]; Dreamweave Gloves (10019, -2.26 DPS) [crafted]; Bloodfire Talons (12464, -3.63 DPS, sim-verified) [dungeon] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, -0.53 DPS) [dungeon]; Satyrmane Sash (17755, -0.97 DPS) [dungeon]; Tooled Leather Belt (252544, -1.12 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 29.1 spell_power points (8.74 DPS) | yes | Jinxed Hoodoo Kilt (9474, -1.74 DPS) [dungeon]; Stone Guard's Crackling Leather Leggings (220865, -1.91 DPS) [vendor]; Senior Designer's Pantaloons (11841, -5.11 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.4 spell_power points (8.22 DPS) | yes | Skycaller's Leather Boots (252471, -2.13 DPS, sim-verified) [crafted]; First Sergeant's Crackling Leather Boots (220863, -2.23 DPS) [vendor]; Stormcloth Boots (10039, -2.28 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.2 spell_power points (4.57 DPS) | yes | Philanthropist's Ring (281635, -0.65 DPS) [quest]; Band of the Unicorn (7553, -0.67 DPS) [world_drop]; Brainlash (6440, -1.08 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 15.2 spell_power points (4.57 DPS) | yes | Band of the Unicorn (7553, -0.67 DPS) [world_drop]; Brainlash (6440, -1.08 DPS) [dungeon]; Philanthropist's Ring (281635, -1.20 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.06 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest]; Rune of the Guard Captain (19120, -0.73 DPS, sim-verified) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -3.48 DPS) [world]; Kindling Stave (11750, -3.66 DPS) [dungeon]; Glowing Brightwood Staff (812, -6.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Mindburst Medallion; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Skycaller's Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength

No-known-source sample (15 of 689, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 367.9. Weights run: 3.2s. Verify run: 34.6s. 1759 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.353 ± 0.021, crit=0.467 ± 0.018 per rating point (14 rating = 1%, 6.539 per %), hit=0.745 ± 0.038 per rating point (10 rating = 1%, 7.453 per %), spell_haste=3.642 ± 0.354, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.480 ± 0.002, arcane_power=0.520 ± 0.002, spirit=not significant (0.293 ± 0.191), mp5=not significant (0.600 ± 0.242)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Dragonhide Helm (227186, +0.00 DPS) [pvp]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -0.13 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -1.21 DPS) [dungeon]; Triune Amulet (7722, -2.58 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Dragonhide Spaulders (227184, +0.00 DPS) [pvp]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -6.15 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.33 DPS, sim-verified) [world]; Cloak of Earth and Sky (279270, -3.38 DPS) [crafted]; Spritecaster Cape (11623, -3.96 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 33.4 spell_power points (15.77 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Alanna's Embrace (13314, -9.38 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.9 spell_power points (12.67 DPS) | yes | Runecloth Cuffs (254123, -5.82 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.10 DPS) [crafted]; Runic Leather Bracers (15092, -6.29 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Dragonhide Gloves (231677, -0.11 DPS) [pvp]; Blood Guard's Dragonhide Gloves (227187, -2.88 DPS) [vendor] |
| waist | Elunite Cord (272401) | Pix Xizzix [vendor] | 24.5 spell_power points (11.54 DPS) | yes | Ban'thok Sash (11662, -0.53 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -0.85 DPS) [rep]; Frostwolf Cloth Belt (19090, -1.38 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.9 spell_power points (24.00 DPS) | yes | General's Dragonhide Pants (231680, -4.47 DPS) [vendor]; Outrider's Silk Leggings (22747, -6.19 DPS, sim-verified) [rep]; Skyshroud Leggings (13170, -6.63 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.1 spell_power points (12.29 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.17 DPS) [pvp]; Venomspew Footpads (275606, -0.81 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.28 DPS) [quest]; Don Mauricio's Band of Domination (22433, -2.21 DPS) [dungeon]; Naglering (11669, -10.09 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.06 DPS) [dungeon]; Advisor's Ring (19518, -1.60 DPS) [rep] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.50 DPS) [crafted]; Uther's Strength (11302, -2.83 DPS) [world_drop]; Rune of the Guard Captain (19120, -3.20 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.83 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -7.21 DPS) [crafted]; Uther's Strength (11302, -7.55 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Mindfang (20214, -1.40 DPS) [rep]; Ironbark Staff (20220, -10.23 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (13.68 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Spirit of Aquementas (11904, -3.83 DPS, sim-verified) [quest]; Tome of the Lost (22253, -4.02 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -1.37 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Elunite Cord; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 715.7. Weights run: 3.1s. Verify run: 31.9s. 1759 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.268 ± 0.023, crit=0.734 ± 0.024 per rating point (14 rating = 1%, 10.273 per %), hit=1.234 ± 0.057 per rating point (10 rating = 1%, 12.341 per %), spell_haste=7.244 ± 0.471, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.348 ± 0.001, arcane_power=0.652 ± 0.001, spirit=not significant (0.019 ± 0.045), mp5=not significant (0.114 ± 0.179)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 35.0 spell_power points (21.96 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Champion's Dragonhide Helm (227186, +0.00 DPS) [pvp]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Emberfury Talisman (12929, -2.71 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.8 spell_power points (19.33 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Champion's Dragonhide Spaulders (227184, -1.50 DPS) [pvp]; Argent Shoulders (19059, -4.34 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.07 DPS, sim-verified) [world]; Spritecaster Cape (11623, -9.28 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -9.55 DPS) [dungeon] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Hauberk (16549, +0.00 DPS) [vendor]; Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.3 spell_power points (15.25 DPS) | yes | Runecloth Cuffs (254123, -6.66 DPS) [crafted]; Runic Leather Bracers (15092, -7.28 DPS) [crafted]; Blackmist Armguards (12966, -7.50 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Dragonhide Gloves (231677, -1.33 DPS) [pvp]; Earth Warder's Gloves (21318, -3.44 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Elunite Cord (272401, -2.16 DPS) [vendor]; Frostwolf Cloth Belt (19090, -3.27 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 53.1 spell_power points (33.36 DPS) | yes | General's Dragonhide Pants (231680, -7.47 DPS) [vendor]; Skyshroud Leggings (13170, -10.65 DPS) [dungeon]; Stone Guard's Crackling Leather Leggings (220865, -12.93 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Crackling Leather Boots (220863) | PvP reward · Horde [vendor] | 26.0 spell_power points (16.35 DPS) | yes | General's Dragonhide Boots (231682, -0.66 DPS) [pvp]; Earthen Silk Slippers (254013, -1.19 DPS) [crafted]; Blood Guard's Dragonhide Boots (227188, -1.42 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.84 DPS) [dungeon]; Maiden's Circle (13001, -1.65 DPS) [world_drop]; Naglering (11669, -14.51 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.76 DPS) [dungeon]; Maiden's Circle (13001, -1.57 DPS) [world_drop]; Eye of Orgrimmar (12545, -1.59 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.56 DPS) [crafted]; Rune of the Guard Captain (19120, -2.11 DPS) [quest]; Uther's Strength (11302, -3.77 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.39 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -6.84 DPS) [crafted]; Rune of the Guard Captain (19120, -8.39 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.05 DPS) [pvp]; Whiteout Staff (19101, -21.53 DPS) [rep]; Staff of Jordan (873, -21.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -2.87 DPS, sim-verified) [vendor] |

**New at 60:** head: Mask of the Unforgiven; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Fernpulse Jerkin; wrist: Dryad's Wrist Bindings; waist: Defiler's Cloth Girdle; legs: Ironfeather Leggings; feet: First Sergeant's Crackling Leather Boots; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

