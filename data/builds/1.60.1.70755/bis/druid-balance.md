# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 39.7. Weights run: 5.6s. Verify run: 7.2s. 222 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.647 ± 0.004, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.443 per %), hit=0.122 ± 0.005 per rating point (10 rating = 1%, 1.219 per %), spell_haste=0.639 ± 0.059, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.836 ± 0.002, arcane_power=0.164 ± 0.001, spirit=0.187 ± 0.021, mp5=0.947 ± 0.041

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.45 DPS) | yes | Pristine Circlet (253949, -0.48 DPS) [crafted]; Trapper's Leather Hood (252505, -0.61 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.69 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.4 spell_power points (0.06 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.8 spell_power points (1.74 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.01 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.10 DPS) [crafted]; Slime-encrusted Pads (6461, -1.28 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (39.7 DPS) | yes | Pearl-clasped Cloak (5542, -0.13 DPS) [crafted]; Caretaker's Cape (20428, -0.22 DPS) [rep]; Black Whelp Cloak (7283, -0.28 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 11.8 spell_power points (1.90 DPS) | yes | Filigreed Pristine Gown (253901, -0.48 DPS) [crafted]; Totemic Leather Armor (252435, -0.62 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.69 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.1 spell_power points (0.98 DPS) | yes | Bright Bracers (3647, -0.56 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.66 DPS) [dungeon]; Owl Bracers (4796, -1.27 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.6 spell_power points (1.22 DPS) | yes | Gnoll Casting Gloves (892, -0.26 DPS) [world]; Windfelt Gloves (5630, -0.26 DPS) [quest]; Serpent Gloves (5970, -0.33 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.6 spell_power points (1.06 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.16 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.6 spell_power points (2.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.64 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.69 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.81 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (1.54 DPS) | yes | Stormrider's Leather Boots (252443, -0.06 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.38 DPS) [crafted]; Totemic Leather Boots (252442, -0.58 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.9 spell_power points (1.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.80 DPS) [world_drop]; Band of Purification (12996, -0.93 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.3 spell_power points (1.01 DPS) | yes | Lavishly Jeweled Ring (1156, -0.58 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.70 DPS) [world_drop]; Band of Purification (12996, -0.83 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.2 spell_power points (4.86 DPS) | yes | Emberstone Staff (5201, -1.00 DPS) [dungeon]; Buzzer Blade (2169, -1.32 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.33 DPS, sim-verified) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 3.2 spell_power points (0.52 DPS) | yes | Tear of Grief (5611, -0.09 DPS) [quest]; Seer's Fine Stein (7608, -0.18 DPS) [world_drop]; Dwarven Tome (279898, -0.33 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb

No-known-source sample (15 of 222, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde

### Band 30 (night-elf, 5232220113000000-00000000000000000000-0000000000000000)

Set DPS (verified): 80.2. Weights run: 5.8s. Verify run: 9.2s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.382 ± 0.011, crit=0.123 ± 0.002 per rating point (14 rating = 1%, 1.722 per %), hit=0.243 ± 0.009 per rating point (10 rating = 1%, 2.435 per %), spell_haste=0.845 ± 0.102, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.839 ± 0.001, arcane_power=0.161 ± 0.001, spirit=not significant (0.017 ± 0.010), mp5=0.533 ± 0.097

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (4.50 DPS) | yes | Holy Shroud (2721, -0.34 DPS) [world_drop]; Enchanter's Cowl (4322, -0.82 DPS) [crafted]; Silk Headband (7050, -1.12 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.3 spell_power points (0.86 DPS) | yes | Seal of the Expedition (271920, -0.14 DPS) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.27 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 spell_power points (4.66 DPS) | yes | Death Speaker Mantle (6685, -0.93 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -1.32 DPS) [crafted]; Magician's Mantle (12998, -1.50 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.35 DPS) [crafted]; Caretaker's Cape (19533, -0.35 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.0 spell_power points (5.23 DPS) | yes | Beguiler Robes (7728, -0.09 DPS) [dungeon]; Tree Bark Jacket (1486, -0.36 DPS) [dungeon]; Guardian Armor (4256, -0.43 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.6 spell_power points (3.97 DPS) | yes | Mindthrust Bracers (1974, -2.51 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.83 DPS) [world_drop]; Nightsky Wristbands (6407, -3.09 DPS) [world_drop] |
| hands | Gloves of Insight (9698) (or Serpent Gloves (5970)) | Return to Vahlarriel [quest] | 7.0 spell_power points (2.62 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Stormrider's Leather Gloves (252498, -0.18 DPS) [crafted]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.4 spell_power points (4.63 DPS) | yes | Highlander's Cloth Girdle (20099, -0.08 DPS) [rep]; Moss Cinch (6911, -0.13 DPS) [dungeon]; Belt of Arugal (6392, -0.80 DPS) [dungeon] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (5.25 DPS) | yes | Stormrider's Leather Pants (252502, -0.55 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -0.69 DPS) [dungeon]; Guardian Pants (5962, -1.10 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (3.65 DPS) | yes | Acidic Walkers (9454, -0.61 DPS) [dungeon]; Spidersilk Boots (4320, -0.61 DPS, sim-verified) [crafted]; Stormrider's Leather Boots (252443, -0.69 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 8.1 spell_power points (3.02 DPS) | yes | Minor Channeling Ring (1449, -0.86 DPS) [quest]; Electrocutioner Lagnut (9447, -1.85 DPS) [dungeon]; Snake Hoop (6750, -3.27 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.28 DPS, sim-verified) [quest]; Minor Channeling Ring (1449, -0.69 DPS) [quest]; Electrocutioner Lagnut (9447, -1.68 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.52 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.52 DPS) [world]; Glimmering Staff (249392, -4.03 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.3 spell_power points (3.48 DPS) | yes | Eye of Paleth (2943, -1.98 DPS) [quest]; Orb of Souls (249395, -2.31 DPS, sim-verified) [crafted]; Dwarven Tome (279898, -2.35 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Gloves of Insight; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (night-elf, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 162.9. Weights run: 6.0s. Verify run: 7.4s. 509 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.608 ± 0.018, crit=0.251 ± 0.010 per rating point (14 rating = 1%, 3.510 per %), hit=0.430 ± 0.024 per rating point (10 rating = 1%, 4.301 per %), spell_haste=2.620 ± 0.331, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.653 ± 0.002, arcane_power=0.347 ± 0.002, spirit=not significant (0.251 ± 0.148), mp5=1.074 ± 0.185

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.30 DPS) | yes | Big Voodoo Mask (8201, -1.38 DPS) [crafted]; Augural Shroud (2620, -1.55 DPS) [world]; Papal Fez (9431, -4.46 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.9 spell_power points (4.69 DPS) | yes | Triune Amulet (7722, -0.45 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.31 DPS) [dungeon]; Seal of the Expedition (271919, -3.01 DPS) [rep] |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.34 DPS, sim-verified) [dungeon]; Berylline Pads (4197, -0.64 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.21 DPS) [crafted]; Expeditionary's Cape (271922, -0.53 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -1.42 DPS) [crafted]; Dreamweave Vest (10021, -1.43 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -2.33 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 spell_power points (4.83 DPS) | yes | Guardian Leather Bracers (4260, -1.02 DPS) [crafted]; Arcane Runed Bracers (4744, -1.27 DPS) [quest]; Green Whelp Bracers (7386, -1.33 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -4.82 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -5.32 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.56 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.4 spell_power points (6.49 DPS) | yes | Skycaller's Leather Belt (252522, -0.70 DPS) [crafted]; Gilded Cord (254037, -0.81 DPS) [crafted]; Deathmage Sash (10771, -2.01 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.3 spell_power points (8.41 DPS) | yes | Warchief Kilt (7760, -2.42 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.49 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.54 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.8 spell_power points (10.18 DPS) | yes | Mender's Leather Shoes (252533, -4.64 DPS) [crafted]; Skycaller's Leather Shoes (252532, -4.72 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -5.33 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (5.15 DPS) | yes | Ring of Power Regulation (273030, -1.51 DPS) [dungeon]; Snake Hoop (6750, -1.58 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.99 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 12.2 spell_power points (4.83 DPS) | yes | Snake Hoop (6750, -1.26 DPS) [quest]; Ring of Power Regulation (273030, -1.28 DPS, sim-verified) [dungeon]; Ring of Forlorn Spirits (2043, -1.67 DPS) [quest] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.35 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -12.45 DPS) [dungeon]; Swinetusk Shank (6691, -15.85 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Green Silken Shoulders; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 509, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 5232220115101151-00000000000000000000-5400000000000000)

Set DPS (verified): 211.4. Weights run: 6.1s. Verify run: 10.5s. 665 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.576 ± 0.018, crit=0.283 ± 0.011 per rating point (14 rating = 1%, 3.960 per %), hit=0.534 ± 0.027 per rating point (10 rating = 1%, 5.341 per %), spell_haste=3.429 ± 0.334, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.648 ± 0.002, arcane_power=0.352 ± 0.002, spirit=not significant (-0.085 ± 0.153), mp5=1.562 ± 0.221

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.6 spell_power points (13.13 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -1.93 DPS) [crafted]; Dreamweave Circlet (10041, -2.02 DPS) [crafted]; Red Mageweave Headband (10033, -2.18 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.6 spell_power points (4.82 DPS) | yes | Mindburst Medallion (11196, -0.89 DPS) [quest]; Horizon Choker (13085, -1.47 DPS) [world_drop]; Triune Amulet (7722, -1.50 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 26.5 spell_power points (11.00 DPS) | yes | Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.42 DPS) [vendor]; Rotgrip Mantle (17732, -2.28 DPS, sim-verified) [dungeon]; Skycaller's Leather Shoulders (252537, -3.46 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.88 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -3.02 DPS) [crafted]; Nightfall Drape (12465, -3.51 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 30.5 spell_power points (12.66 DPS) | yes | Robe of the Magi (1716, -2.10 DPS) [world_drop]; Feathered Breastplate (8349, -2.39 DPS) [crafted]; Runecloth Tunic (13857, -2.98 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 15.8 spell_power points (6.54 DPS) | yes | Skycaller's Leather Bracers (252542, -0.72 DPS) [crafted]; Spidertank Oilrag (9448, -0.86 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.96 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Crackling Leather Gauntlets (220866, -4.21 DPS) [vendor]; Dreamweave Gloves (10019, -4.22 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -5.98 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 22.9 spell_power points (9.50 DPS) | yes | Satyrmane Sash (17755, -0.91 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -2.48 DPS) [dungeon]; Mender's Leather Waistguard (252477, -2.49 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 27.1 spell_power points (11.23 DPS) | yes | Big Voodoo Pants (8202, -2.61 DPS) [crafted]; Wizardweave Leggings (14132, -3.34 DPS) [crafted]; Red Mageweave Pants (10009, -4.53 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.96 DPS) | yes | Skycaller's Leather Boots (252471, -1.10 DPS) [crafted]; Stormcloth Boots (10039, -2.34 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -3.43 DPS, sim-verified) [vendor] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 18.2 spell_power points (7.57 DPS) | yes | Philanthropist's Ring (281635, -2.23 DPS) [quest]; Ring of Power Regulation (273030, -3.14 DPS) [dungeon]; Snake Hoop (6750, -3.89 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.39 DPS) | yes | Ring of Power Regulation (273030, -0.96 DPS) [dungeon]; Philanthropist's Ring (281635, -1.06 DPS, sim-verified) [quest]; Snake Hoop (6750, -1.71 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271905, -0.05 DPS) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271905, -0.55 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -3.39 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -5.09 DPS) [quest]; Radiant Staff (249453, -5.22 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Knight's Crackling Leather Leggings; finger1: Lorekeeper's Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength

No-known-source sample (15 of 665, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 336.0. Weights run: 5.0s. Verify run: 30.9s. 1443 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.228 ± 0.019, crit=0.426 ± 0.015 per rating point (14 rating = 1%, 5.966 per %), hit=0.634 ± 0.031 per rating point (10 rating = 1%, 6.344 per %), spell_haste=4.875 ± 0.383, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.580 ± 0.002, arcane_power=0.420 ± 0.002, spirit=not significant (0.200 ± 0.204), mp5=not significant (0.243 ± 0.240)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 31.9 spell_power points (15.39 DPS) | yes | Lieutenant Commander's Dragonhide Helm (227192, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Soothsayer's Headdress (17740, -3.68 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.3 spell_power points (6.89 DPS) | yes | Glowing Eye of Mordresh (10769, -2.63 DPS) [dungeon]; Mindburst Medallion (11196, -2.85 DPS) [quest]; Triune Amulet (7722, -3.77 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (336.0 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Argent Shoulders (19059, -8.08 DPS, sim-verified) [crafted] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | sim-verified (336.0 DPS) | yes | Spritecaster Cape (11623, -0.60 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -1.84 DPS) [rep]; Runecloth Cloak (13860, -3.37 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 29.3 spell_power points (14.16 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -11.55 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.2 spell_power points (12.17 DPS) | yes | Runecloth Cuffs (254123, -5.77 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.09 DPS) [crafted]; Runic Leather Bracers (15092, -6.25 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (336.0 DPS) | yes | Marshal's Dragonhide Gloves (231700, -1.67 DPS) [vendor]; Earth Warder's Gloves (21318, -3.88 DPS, sim-verified) [quest]; Knight-Lieutenant's Dragonhide Gloves (227193, -4.21 DPS) [pvp] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 21.3 spell_power points (10.29 DPS) | yes | Skycaller's Leather Waistguard (252476, -1.25 DPS) [crafted]; Stormpike Cloth Girdle (19094, -1.65 DPS, sim-verified) [rep]; Satyrmane Sash (17755, -2.44 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 48.1 spell_power points (23.19 DPS) | yes | Sentinel's Lizardhide Pants (237817, -0.90 DPS) [vendor]; Sentinel's Silk Leggings (237815, -2.22 DPS, sim-verified) [vendor]; Marshal's Dragonhide Pants (231697, -4.78 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.4 spell_power points (12.26 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.64 DPS) [vendor]; Venomspew Footpads (275606, -1.38 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.8 spell_power points (10.04 DPS) | yes | Deep Rooted Ring (19109, -3.15 DPS) [world]; Band of the Unicorn (7553, -3.76 DPS) [world_drop]; Don Julio's Band (19325, -4.09 DPS) [rep] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 15.0 spell_power points (7.23 DPS) | yes | Deep Rooted Ring (19109, -0.34 DPS) [world]; Band of the Unicorn (7553, -0.95 DPS) [world_drop]; Don Julio's Band (19325, -1.28 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (336.0 DPS) | yes | Ankh of Life (1713, -1.60 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.79 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.98 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (336.0 DPS) | yes | Evonice's Landin' Pilla (18951, -1.93 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.12 DPS) [quest]; Ankh of Life (1713, -3.03 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (336.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Sageclaw (20070, -0.48 DPS) [rep]; Ironbark Staff (20069, -10.33 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (14.00 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.50 DPS) [pvp]; Spirit of Aquementas (11904, -3.94 DPS, sim-verified) [quest]; Grand Marshal's Tome of Restoration (234590, -5.22 DPS) [pvp] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-verified (336.0 DPS) | yes | Idol of Swiftness (279250, -2.35 DPS, sim-verified) [crafted] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Cloak of Earth and Sky; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Highlander's Cloth Girdle; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Idol of the Moon

No-known-source sample (15 of 1443, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 609.6. Weights run: 4.7s. Verify run: 35.5s. 1443 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.201 ± 0.021, crit=0.587 ± 0.019 per rating point (14 rating = 1%, 8.222 per %), hit=0.887 ± 0.046 per rating point (10 rating = 1%, 8.867 per %), spell_haste=7.981 ± 0.495, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.358 ± 0.001, arcane_power=0.642 ± 0.001, spirit=not significant (-0.053 ± 0.071), mp5=not significant (0.231 ± 0.176)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 28.6 spell_power points (17.96 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, +0.00 DPS) [crafted]; Lieutenant Commander's Dragonhide Helm (227192, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.0 spell_power points (8.79 DPS) | yes | Glowing Eye of Mordresh (10769, -3.05 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.77 DPS) [dungeon]; Mindburst Medallion (11196, -4.27 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.8 spell_power points (17.47 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Lieutenant Commander's Dragonhide Spaulders (227190, -1.51 DPS) [vendor]; Argent Shoulders (19059, -1.78 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Earth and Sky (279270, -0.13 DPS) [crafted]; Stormpike Sage's Cloak (19086, -0.76 DPS) [rep]; Runecloth Cloak (13860, -2.89 DPS) [crafted] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | 27.6 spell_power points (17.33 DPS) | yes | Ironfeather Breastplate (15066, +0.00 DPS) [crafted]; Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.6 spell_power points (14.82 DPS) | yes | Runecloth Cuffs (254123, -6.65 DPS) [crafted]; Runic Leather Bracers (15092, -7.28 DPS) [crafted]; Skycaller's Leather Bracers (252542, -7.66 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Gloves (231700, -1.99 DPS) [vendor]; Earth Warder's Gloves (21318, -3.61 DPS, sim-verified) [quest]; Knight-Lieutenant's Dragonhide Gloves (227193, -4.75 DPS) [pvp] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 23.4 spell_power points (14.70 DPS) | yes | Stormpike Cloth Girdle (19094, -2.14 DPS) [rep]; Skycaller's Leather Waistguard (252476, -3.15 DPS) [crafted]; Satyrmane Sash (17755, -4.65 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) (or Sentinel's Lizardhide Pants (237817)) | Illiyana Moonblaze [vendor] | 50.5 spell_power points (31.66 DPS) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -0.38 DPS) [crafted]; Marshal's Dragonhide Pants (231697, -7.79 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.06 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.12 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.62 DPS) [vendor]; Omnicast Boots (11822, -1.00 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.4 spell_power points (12.18 DPS) | yes | Lorekeeper's Ring (19522, -2.81 DPS) [rep]; Deep Rooted Ring (19109, -3.39 DPS) [world]; Band of the Unicorn (7553, -4.02 DPS) [world_drop] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 17.1 spell_power points (10.72 DPS) | yes | Lorekeeper's Ring (19522, -1.36 DPS) [rep]; Deep Rooted Ring (19109, -1.93 DPS) [world]; Band of the Unicorn (7553, -2.57 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -4.14 DPS) [quest]; Enriched Seal (271903, -4.28 DPS) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.89 DPS) [quest]; Enriched Seal (271903, -3.04 DPS) [rep] |
| main_hand | Sageclaw (20070) | The League of Arathor [rep] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -1.33 DPS, sim-verified) [world_drop] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (18.20 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.87 DPS) [pvp]; Spirit of Aquementas (11904, -5.14 DPS, sim-verified) [quest]; Tome of Arcane Domination (19308, -5.28 DPS) [rep] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Swiftness (279250, -2.87 DPS, sim-verified) [crafted] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Fernpulse Jerkin; wrist: Dryad's Wrist Bindings; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Don Julio's Band; main_hand: Sageclaw; off_hand: Therazane's Touch; ranged: Idol of the Moon

No-known-source sample (15 of 1443, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 37.9. Weights run: 5.6s. Verify run: 8.7s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.591 ± 0.008, crit=0.059 ± 0.001 per rating point (14 rating = 1%, 0.823 per %), hit=0.225 ± 0.007 per rating point (10 rating = 1%, 2.252 per %), spell_haste=0.560 ± 0.079, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.802 ± 0.001, arcane_power=0.198 ± 0.001, spirit=not significant (-0.022 ± 0.007), mp5=1.296 ± 0.076

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (2.10 DPS) | yes | Wisdom's Leather Hood (252507, -0.69 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -0.70 DPS) [crafted]; Trapper's Leather Hood (252505, -1.00 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.3 spell_power points (2.41 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.68 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.48 DPS) [crafted]; Slime-encrusted Pads (6461, -1.50 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.11 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Battle Healer's Cloak (20427, -0.18 DPS) [rep] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 11.0 spell_power points (2.56 DPS) | yes | Wisdom's Leather Armor (252493, -0.69 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.70 DPS) [crafted]; Totemic Leather Armor (252435, -0.74 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.8 spell_power points (1.60 DPS) | yes | Owl Bracers (4796, -0.91 DPS) [vendor]; Bright Bracers (3647, -1.05 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.58 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.4 spell_power points (1.72 DPS) | yes | Serpent Gloves (5970, -0.09 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.32 DPS) [world]; Pristine Gloves (253913, -0.37 DPS) [crafted] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.4 spell_power points (1.49 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.14 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.23 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.5 spell_power points (3.16 DPS) | yes | Wisdom's Leather Pants (252503, -0.69 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.93 DPS) [crafted]; Totemic Leather Pants (252446, -0.98 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.4 spell_power points (2.19 DPS) | yes | Stormrider's Leather Boots (252443, -0.10 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.56 DPS) [crafted]; Totemic Leather Boots (252442, -0.79 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.6 spell_power points (1.77 DPS) | yes | Loop of Sacrifice (281673, -1.08 DPS) [quest]; Volcanic Rock Ring (12053, -1.36 DPS) [world_drop]; Furlfeather Ring (286745, -1.54 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.5 spell_power points (0.83 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; Furlfeather Ring (286745, -0.59 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Crescent Staff (6505, -0.20 DPS, sim-verified) [quest]; Staff of the Blessed Seer (2271, -0.28 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.74 DPS) [quest] |
| off_hand | Spellbinder Orb (15926) (or Seer's Fine Stein (7608)) | Wailing Caverns: Large Battered Chest [dungeon] | 1.8 spell_power points (0.41 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop]; Tork Wrench (11855, -0.14 DPS) [quest]; Ancestral Orb (15944, -0.28 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Spellbinder Orb

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 248005 Northshire Hammer

### Band 30 (tauren, 5232220113000000-00000000000000000000-0000000000000000)

Set DPS (verified): 80.0. Weights run: 5.8s. Verify run: 9.1s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.382 ± 0.011, crit=0.123 ± 0.002 per rating point (14 rating = 1%, 1.722 per %), hit=0.243 ± 0.009 per rating point (10 rating = 1%, 2.435 per %), spell_haste=0.845 ± 0.102, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.839 ± 0.001, arcane_power=0.161 ± 0.001, spirit=not significant (0.017 ± 0.010), mp5=0.533 ± 0.097

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (4.50 DPS) | yes | Holy Shroud (2721, -0.34 DPS) [world_drop]; Enchanter's Cowl (4322, -0.82 DPS) [crafted]; Silk Headband (7050, -1.12 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.3 spell_power points (0.86 DPS) | yes | Darkspear Warding Pendant (272075, -0.14 DPS) [vendor]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.27 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 spell_power points (4.66 DPS) | yes | Death Speaker Mantle (6685, -0.48 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -1.32 DPS) [crafted]; Magician's Mantle (12998, -1.50 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.03 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.37 DPS) [crafted]; Battle Healer's Cloak (19529, -0.37 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.0 spell_power points (5.23 DPS) | yes | Beguiler Robes (7728, -0.09 DPS) [dungeon]; Tree Bark Jacket (1486, -0.36 DPS) [dungeon]; Guardian Armor (4256, -0.43 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.6 spell_power points (3.97 DPS) | yes | Mindthrust Bracers (1974, -2.52 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.83 DPS) [world_drop]; Tabitha's Cuffs (251486, -2.85 DPS) [quest] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 9.8 spell_power points (3.66 DPS) | yes | Jutebraid Gloves (10654, -0.46 DPS, sim-verified) [quest]; Serpent Gloves (5970, -1.04 DPS) [dungeon]; Stormrider's Leather Gloves (252498, -1.21 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.4 spell_power points (4.63 DPS) | yes | Defiler's Cloth Girdle (20164, -0.08 DPS) [rep]; Moss Cinch (6911, -0.13 DPS) [dungeon]; Warsong Sash (16975, -0.51 DPS) [quest] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (5.25 DPS) | yes | Stormrider's Leather Pants (252502, -0.61 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.69 DPS) [dungeon]; Guardian Pants (5962, -1.10 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (3.65 DPS) | yes | Acidic Walkers (9454, -0.61 DPS) [dungeon]; Spidersilk Boots (4320, -0.61 DPS, sim-verified) [crafted]; Stormrider's Leather Boots (252443, -0.69 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 8.1 spell_power points (3.02 DPS) | yes | Solid Gold Band (279845, -1.49 DPS) [quest]; Electrocutioner Lagnut (9447, -1.85 DPS) [dungeon]; Snake Hoop (6750, -3.43 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.35 DPS, sim-verified) [quest]; Solid Gold Band (279845, -1.32 DPS) [quest]; Electrocutioner Lagnut (9447, -1.68 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.52 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.52 DPS) [world]; Glimmering Staff (249392, -4.00 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.3 spell_power points (3.48 DPS) | yes | Witch's Finger (16887, -2.34 DPS) [quest]; Orb of Souls (249395, -2.60 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -2.61 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 5232220115101141-00000000000000000000-0000000000000000)

Set DPS (verified): 160.7. Weights run: 6.0s. Verify run: 7.4s. 508 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.608 ± 0.018, crit=0.251 ± 0.010 per rating point (14 rating = 1%, 3.510 per %), hit=0.430 ± 0.024 per rating point (10 rating = 1%, 4.301 per %), spell_haste=2.620 ± 0.331, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.653 ± 0.002, arcane_power=0.347 ± 0.002, spirit=not significant (0.251 ± 0.148), mp5=1.074 ± 0.185

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.30 DPS) | yes | Big Voodoo Mask (8201, -1.38 DPS) [crafted]; Augural Shroud (2620, -1.55 DPS) [world]; Papal Fez (9431, -4.41 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.9 spell_power points (4.69 DPS) | yes | Triune Amulet (7722, -0.31 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.31 DPS) [dungeon]; Darkspear Warding Pendant (272074, -3.01 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.9 spell_power points (5.89 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Berylline Pads (4197, -0.72 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.21 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.53 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (160.7 DPS) | yes | Dreamweave Vest (10021, -1.17 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.42 DPS) [crafted]; Elemental Raiment (9434, -2.33 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 spell_power points (4.83 DPS) | yes | Guardian Leather Bracers (4260, -1.02 DPS) [crafted]; Radiant Silver Bracers (4545, -1.33 DPS) [quest]; Green Whelp Bracers (7386, -1.58 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -5.09 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -5.32 DPS) [crafted]; Skycaller's Leather Gloves (252529, -5.56 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.4 spell_power points (6.49 DPS) | yes | Skycaller's Leather Belt (252522, -0.70 DPS) [crafted]; Gilded Cord (254037, -0.81 DPS) [crafted]; Deathmage Sash (10771, -2.50 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.3 spell_power points (8.41 DPS) | yes | Warchief Kilt (7760, -2.42 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.54 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -3.17 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.8 spell_power points (10.18 DPS) | yes | Skycaller's Leather Shoes (252532, -4.30 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -4.64 DPS) [crafted]; Gilded Slippers (254001, -5.33 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (5.15 DPS) | yes | Ring of Power Regulation (273030, -1.51 DPS) [dungeon]; Snake Hoop (6750, -1.58 DPS) [quest]; Reedknot Ring (9622, -2.39 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 12.2 spell_power points (4.83 DPS) | yes | Snake Hoop (6750, -1.26 DPS) [quest]; Ring of Power Regulation (273030, -1.27 DPS, sim-verified) [dungeon]; Reedknot Ring (9622, -2.06 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.32 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -12.45 DPS) [dungeon]; Swinetusk Shank (6691, -15.85 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Talons of Wrath

No-known-source sample (15 of 508, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 50 (tauren, 5232220115101151-00000000000000000000-5400000000000000)

Set DPS (verified): 207.5. Weights run: 6.1s. Verify run: 10.0s. 664 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.576 ± 0.018, crit=0.283 ± 0.011 per rating point (14 rating = 1%, 3.960 per %), hit=0.534 ± 0.027 per rating point (10 rating = 1%, 5.341 per %), spell_haste=3.429 ± 0.334, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.648 ± 0.002, arcane_power=0.352 ± 0.002, spirit=not significant (-0.085 ± 0.153), mp5=1.562 ± 0.221

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.6 spell_power points (13.13 DPS) | yes | Red Mageweave Headband (10033, -1.74 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.93 DPS) [crafted]; Dreamweave Circlet (10041, -2.02 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.6 spell_power points (4.82 DPS) | yes | Mindburst Medallion (11196, -0.87 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.47 DPS) [world_drop]; Triune Amulet (7722, -1.50 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 26.5 spell_power points (11.00 DPS) | yes | Blood Guard's Crackling Leather Spaulders (220871, -1.42 DPS) [vendor]; Rotgrip Mantle (17732, -2.27 DPS, sim-verified) [dungeon]; Skycaller's Leather Shoulders (252537, -3.46 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (207.5 DPS) | yes | Deep Woodlands Cloak (19121, -0.11 DPS) [quest]; Runecloth Cloak (13860, -1.60 DPS) [crafted]; Big Voodoo Cloak (8216, -3.02 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 30.5 spell_power points (12.66 DPS) | yes | Robe of the Magi (1716, -2.10 DPS) [world_drop]; Feathered Breastplate (8349, -2.39 DPS) [crafted]; Runecloth Tunic (13857, -2.98 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 15.8 spell_power points (6.54 DPS) | yes | Skycaller's Leather Bracers (252542, -0.72 DPS) [crafted]; Spidertank Oilrag (9448, -0.86 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.96 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (207.5 DPS) | yes | First Sergeant's Crackling Leather Gauntlets (220867, -4.21 DPS) [vendor]; Dreamweave Gloves (10019, -4.22 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -6.03 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 22.9 spell_power points (9.50 DPS) | yes | Satyrmane Sash (17755, -1.17 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -2.48 DPS) [dungeon]; Mender's Leather Waistguard (252477, -2.49 DPS) [crafted] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | Lady Palanseer [vendor] | 27.1 spell_power points (11.23 DPS) | yes | Big Voodoo Pants (8202, -2.61 DPS) [crafted]; Wizardweave Leggings (14132, -3.34 DPS) [crafted]; Red Mageweave Pants (10009, -4.08 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.96 DPS) | yes | Skycaller's Leather Boots (252471, -1.10 DPS) [crafted]; Stormcloth Boots (10039, -2.34 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -4.25 DPS, sim-verified) [vendor] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 18.2 spell_power points (7.57 DPS) | yes | Philanthropist's Ring (281635, -2.23 DPS) [quest]; Darkspear Signet (272069, -3.03 DPS) [vendor]; Ring of Power Regulation (273030, -3.14 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.39 DPS) | yes | Darkspear Signet (272069, -0.86 DPS) [vendor]; Ring of Power Regulation (273030, -0.96 DPS) [dungeon]; Philanthropist's Ring (281635, -1.44 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (207.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.44 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (207.5 DPS) | yes | Uther's Strength (11302, -0.75 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.69 DPS) [quest]; Ankh of Life (1713, -1.80 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (207.5 DPS) | yes | Kindling Stave (11750, -3.57 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -5.09 DPS) [quest]; Radiant Staff (249453, -5.22 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Crackling Leather Leggings; finger1: Advisor's Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Darkspear Voodoo Seal

No-known-source sample (15 of 664, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60 (tauren, 5232220115101151-00000000000000000000-5553010000000000)

Set DPS (verified): 331.1. Weights run: 5.0s. Verify run: 32.2s. 1448 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.228 ± 0.019, crit=0.426 ± 0.015 per rating point (14 rating = 1%, 5.966 per %), hit=0.634 ± 0.031 per rating point (10 rating = 1%, 6.344 per %), spell_haste=4.875 ± 0.383, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.580 ± 0.002, arcane_power=0.420 ± 0.002, spirit=not significant (0.200 ± 0.204), mp5=not significant (0.243 ± 0.240)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 31.9 spell_power points (15.39 DPS) | yes | Champion's Dragonhide Helm (227186, +0.00 DPS) [pvp]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Soothsayer's Headdress (17740, -1.76 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.3 spell_power points (6.89 DPS) | yes | Glowing Eye of Mordresh (10769, -2.63 DPS) [dungeon]; Mindburst Medallion (11196, -2.85 DPS) [quest]; Triune Amulet (7722, -3.26 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (331.1 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Argent Shoulders (19059, -8.24 DPS, sim-verified) [crafted] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | sim-verified (331.1 DPS) | yes | Spritecaster Cape (11623, -0.60 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.81 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -1.84 DPS) [rep] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 29.3 spell_power points (14.16 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -11.69 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.2 spell_power points (12.17 DPS) | yes | Runecloth Cuffs (254123, -5.77 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.09 DPS) [crafted]; Runic Leather Bracers (15092, -6.25 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (331.1 DPS) | yes | General's Dragonhide Gloves (231677, -1.67 DPS) [pvp]; Earth Warder's Gloves (21318, -3.85 DPS, sim-verified) [quest]; Blood Guard's Dragonhide Gloves (227187, -4.21 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 21.3 spell_power points (10.29 DPS) | yes | Frostwolf Cloth Belt (19090, -0.51 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.25 DPS) [crafted]; Satyrmane Sash (17755, -2.44 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 48.1 spell_power points (23.19 DPS) | yes | General's Dragonhide Pants (231680, -4.78 DPS) [vendor]; Outrider's Silk Leggings (22747, -7.47 DPS, sim-verified) [rep]; Legionnaire's Dragonhide Pants (227185, -7.54 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.4 spell_power points (12.26 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.64 DPS) [pvp]; Venomspew Footpads (275606, -0.87 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.8 spell_power points (10.04 DPS) | yes | Deep Rooted Ring (19109, -3.15 DPS) [world]; Band of the Unicorn (7553, -3.76 DPS) [world_drop]; Don Julio's Band (19325, -4.09 DPS) [rep] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 15.0 spell_power points (7.23 DPS) | yes | Deep Rooted Ring (19109, -0.34 DPS) [world]; Band of the Unicorn (7553, -0.95 DPS) [world_drop]; Don Julio's Band (19325, -1.28 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (331.1 DPS) | yes | Rune of the Guard Captain (19120, -0.61 DPS) [quest]; Ankh of Life (1713, -1.60 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.79 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (331.1 DPS) | yes | Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Rune of the Guard Captain (19120, -0.75 DPS) [quest]; Ankh of Life (1713, -1.74 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (331.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Mindfang (20214, -0.48 DPS) [rep]; Ironbark Staff (20220, -10.38 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (14.00 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.50 DPS) [pvp]; Spirit of Aquementas (11904, -3.89 DPS, sim-verified) [quest]; High Warlord's Tome of Mending (234564, -5.22 DPS) [pvp] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-verified (331.1 DPS) | yes | Talons of Wrath (249441, +0.00 DPS) [crafted] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Cloak of Earth and Sky; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Defiler's Cloth Girdle; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Idol of the Moon

No-known-source sample (15 of 1448, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 604.1. Weights run: 4.7s. Verify run: 36.4s. 1448 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.222 ± 0.016, crit=0.377 ± 0.014 per rating point (14 rating = 1%, 5.277 per %), hit=0.558 ± 0.030 per rating point (10 rating = 1%, 5.582 per %), spell_haste=4.474 ± 0.382, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.382 ± 0.001, arcane_power=0.618 ± 0.001, spirit=not significant (0.165 ± 0.091), mp5=not significant (0.390 ± 0.159)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 31.3 spell_power points (18.27 DPS) | yes | Soothsayer's Headdress (17740, +0.00 DPS) [dungeon]; Champion's Dragonhide Helm (227186, +0.00 DPS) [pvp]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.2 spell_power points (8.30 DPS) | yes | Triune Amulet (7722, -2.96 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -3.18 DPS) [dungeon]; Mindburst Medallion (11196, -3.54 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (604.1 DPS) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -6.50 DPS, sim-verified) [crafted] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | sim-verified (604.1 DPS) | yes | Spritecaster Cape (11623, -0.57 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.93 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -1.93 DPS) [rep] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 28.3 spell_power points (16.50 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -8.80 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.9 spell_power points (14.55 DPS) | yes | Runecloth Cuffs (254123, -6.83 DPS) [crafted]; Skycaller's Leather Bracers (252542, -7.32 DPS) [crafted]; Runic Leather Bracers (15092, -7.41 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (604.1 DPS) | yes | General's Dragonhide Gloves (231677, -1.07 DPS) [pvp]; Earth Warder's Gloves (21318, -3.40 DPS, sim-verified) [quest]; Blood Guard's Dragonhide Gloves (227187, -4.05 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 20.6 spell_power points (12.02 DPS) | yes | Skycaller's Leather Waistguard (252476, -1.13 DPS) [crafted]; Satyrmane Sash (17755, -2.56 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -3.71 DPS, sim-verified) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 47.3 spell_power points (27.58 DPS) | yes | General's Dragonhide Pants (231680, -5.97 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -9.28 DPS) [vendor]; Outrider's Silk Leggings (22747, -11.44 DPS, sim-verified) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (14.67 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.74 DPS) [pvp]; Venomspew Footpads (275606, -1.07 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.5 spell_power points (11.99 DPS) | yes | Deep Rooted Ring (19109, -3.69 DPS) [world]; Band of the Unicorn (7553, -4.40 DPS) [world_drop]; Philanthropist's Ring (281635, -5.50 DPS) [quest] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 15.6 spell_power points (9.08 DPS) | yes | Deep Rooted Ring (19109, -0.78 DPS) [world]; Band of the Unicorn (7553, -1.49 DPS) [world_drop]; Philanthropist's Ring (281635, -2.60 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (604.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, -1.57 DPS) [quest]; Ankh of Life (1713, -1.78 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (604.1 DPS) | yes | Shard of the Splithooves (10659, -0.91 DPS) [quest]; Ankh of Life (1713, -1.13 DPS) [world_drop]; Uther's Strength (11302, -3.49 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (604.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Mindfang (20214, -0.74 DPS) [rep]; Ironbark Staff (20220, -9.20 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (16.92 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.65 DPS) [pvp]; Spirit of Aquementas (11904, -4.95 DPS, sim-verified) [quest]; Tome of Arcane Domination (19308, -5.06 DPS) [rep] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-verified (604.1 DPS) | yes | Idol of Swiftness (279250, -2.80 DPS, sim-verified) [crafted] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Cloak of Earth and Sky; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Defiler's Cloth Girdle; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Idol of the Moon

No-known-source sample (15 of 1448, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

