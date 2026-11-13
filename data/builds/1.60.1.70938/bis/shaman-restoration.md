# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-2530000010000000)

Set DPS (verified): 73.1. Weights run: 7.2s. Verify run: 41.7s. 286 eligible items had no known source.

Race sweep: dwarf 73.1 ± 0.0 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.119, intellect=0.235 ± 0.014, spirit=0.025 ± 0.004, mp5=2.902 ± 0.076, crit=0.076 ± 0.007 per rating point (14 rating = 1%, 1.067 per %), spell_haste=not significant (0.038 ± 0.073)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.01 DPS) | yes | Wisdom's Leather Hood (252507, -0.06 DPS, sim-verified) [crafted]; Totemic Leather Hood (252448, -0.50 DPS) [crafted]; Trapper's Leather Hood (252505, -0.90 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.1 healing_power points (0.00 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | sim-verified (73.1 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.07 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.17 DPS) [crafted]; Slime-encrusted Pads (6461, -0.26 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.27 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.34 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 16.3 healing_power points (0.91 DPS) | yes | Filigreed Pristine Gown (253901, -0.11 DPS, sim-verified) [crafted]; Bloody Apron (6226, -0.18 DPS) [dungeon]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.9 healing_power points (0.55 DPS) | yes | Owl Bracers (4796, -0.49 DPS) [vendor]; Windsong Bangles (263336, -0.50 DPS) [quest]; Bright Bracers (3647, -0.50 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 11.7 healing_power points (0.66 DPS) | yes | Wisdom's Leather Gloves (252499, -0.06 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.32 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 11.9 healing_power points (0.67 DPS) | yes | Wisdom's Leather Belt (252433, -0.16 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.29 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.39 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 21.5 healing_power points (1.20 DPS) | yes | Filigreed Pristine Leggings (253937, -0.11 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.56 DPS) [crafted]; Totemic Leather Pants (252446, -0.76 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 12.2 healing_power points (0.68 DPS) | yes | Black Whelp Slippers (252424, -0.12 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.14 DPS) [crafted]; Spidersilk Boots (4320, -0.24 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 10.8 healing_power points (0.60 DPS) | yes | Furlfeather Ring (286745, -0.49 DPS) [world]; Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 healing_power points (0.31 DPS) | yes | Furlfeather Ring (286745, -0.17 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 49.3 healing_power points (2.76 DPS) | yes | Taskmaster Axe (5194, -0.18 DPS, sim-verified) [dungeon]; Living Root (6631, -0.24 DPS) [dungeon]; Evocator's Blade (2567, -1.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (dwarf, 0000000000000000-000000000000000000-5530001010411000)

Set DPS (verified): 133.8. Weights run: 7.7s. Verify run: 46.0s. 475 eligible items had no known source.

Race sweep: dwarf 133.9 ± 0.0 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.019, intellect=0.689 ± 0.004, spirit=0.457 ± 0.004, mp5=1.405 ± 0.004, crit=0.079 ± 0.003 per rating point (14 rating = 1%, 1.112 per %), spell_haste=-0.611 ± 0.064

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 35.7 healing_power points (7.43 DPS) | yes | Wisdom's Leather Helm (252515, -3.07 DPS) [crafted]; Filigreed Pristine Circlet (253975, -3.51 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -3.69 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.6 healing_power points (0.95 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.16 DPS, sim-verified) [world]; Seal of the Expedition (271920, -0.24 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.2 healing_power points (3.16 DPS) | yes | Death Speaker Mantle (6685, -0.24 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.83 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-verified (133.8 DPS) | yes | Prelacy Cape (7004, -0.26 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.87 DPS) [dungeon]; Repairman's Cape (9605, -0.93 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 27.6 healing_power points (5.73 DPS) | yes | Wisdom's Leather Tunic (252511, -0.56 DPS) [crafted]; Pristine Gown (253961, -0.62 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.38 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.2 healing_power points (2.75 DPS) | yes | Mindthrust Bracers (1974, -0.69 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.60 DPS) [world_drop]; Nightsky Wristbands (6407, -1.60 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.1 healing_power points (3.55 DPS) | yes | Pristine Gloves (253913, -0.83 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.90 DPS) [crafted]; Naga Battle Gloves (888, -0.95 DPS, sim-verified) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 25.9 healing_power points (5.38 DPS) | yes | Skycaller's Leather Belt (252522, -2.34 DPS, sim-verified) [crafted]; Pristine Sash (253925, -2.52 DPS) [crafted]; Belt of Arugal (6392, -2.60 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 29.1 healing_power points (6.05 DPS) | yes | Wisdom's Leather Leggings (252519, -0.02 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.65 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.07 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 26.7 healing_power points (5.54 DPS) | yes | Acidic Walkers (9454, -2.98 DPS) [dungeon]; Black Whelp Slippers (252424, -3.10 DPS) [crafted]; Wisdom's Leather Boots (252444, -3.24 DPS, sim-verified) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 15.9 healing_power points (3.30 DPS) | yes | Ring of Power Regulation (273030, -1.18 DPS) [dungeon]; Lorekeeper's Ring (19525, -1.26 DPS) [rep]; Snake Hoop (6750, -1.34 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 healing_power points (2.54 DPS) | yes | Ring of Power Regulation (273030, -0.14 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19525, -0.50 DPS) [rep]; Snake Hoop (6750, -0.57 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (133.8 DPS) | yes | Odo's Ley Staff (6318, -2.19 DPS) [dungeon]; Staff of Westfall (2042, -2.87 DPS) [quest]; Staff of the Friar (3415, -4.89 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.7 healing_power points (3.27 DPS) | yes | Heart of Agamaggan (6694, -0.53 DPS, sim-verified) [dungeon]; Eye of Paleth (2943, -0.57 DPS) [quest]; Orb of Mystic Insight (249394, -0.96 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 0000000000000000-000000000000000000-5530001011513051)

Set DPS (verified): 207.7. Weights run: 14.9s. Verify run: 168.7s. 804 eligible items had no known source.

Race sweep: dwarf 207.8 ± 0.1 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.025, intellect=1.186 ± 0.010, spirit=0.902 ± 0.009, mp5=2.385 ± 0.010, crit=0.146 ± 0.005 per rating point (14 rating = 1%, 2.044 per %), spell_haste=not significant (0.056 ± 0.087)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 50.3 healing_power points (8.54 DPS) | yes | Holy Shroud (2721, -2.01 DPS) [world_drop]; Whitemane's Chapeau (7720, -2.92 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -3.08 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.5 healing_power points (5.01 DPS) | yes | Triune Amulet (7722, -2.58 DPS) [dungeon]; Amberglow Talisman (10824, -3.48 DPS) [quest]; Necklace of Calisea (1714, -3.94 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.4 healing_power points (6.35 DPS) | yes | Sheepshear Mantle (13115, -2.21 DPS) [world_drop]; Mantle of Doan (7712, -2.73 DPS) [dungeon]; Earthen Silk Shoulders (254033, -5.01 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.19 DPS, sim-verified) [rep]; Prelacy Cape (7004, -1.04 DPS) [quest]; Glowing Thresher Cape (6901, -1.41 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.2 healing_power points (11.08 DPS) | yes | Dreamweave Vest (10021, -4.06 DPS) [crafted]; Death's Head Vestment (10581, -4.42 DPS) [dungeon]; Deathchill Armor (10764, -4.59 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.2 healing_power points (2.74 DPS) | yes | Enchanted Kodo Bracers (13119, -0.41 DPS) [world_drop]; Guardian Leather Bracers (4260, -0.52 DPS) [crafted]; Earthen Silk Cuffs (254019, -1.12 DPS, sim-verified) [crafted] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prefect's Gauntlet (250569, -0.40 DPS) [crafted]; Dreamweave Gloves (10019, -0.44 DPS) [crafted]; Gilded Handwraps (254021, -0.48 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.9 healing_power points (6.27 DPS) | yes | Prefect's Belt (250559, -1.12 DPS) [crafted]; Mender's Leather Belt (252523, -1.77 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -2.06 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 43.7 healing_power points (7.43 DPS) | yes | Stoneweaver Leggings (9407, -0.76 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.38 DPS) [crafted]; Pristine Leggings (253987, -1.52 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 38.6 healing_power points (6.56 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.14 DPS) [crafted]; Skycaller's Leather Shoes (252532, -2.17 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 21.5 healing_power points (3.66 DPS) | yes | Philanthropist's Ring (281635, -0.95 DPS) [quest]; Electrocutioner Lagnut (9447, -1.05 DPS) [dungeon]; Ring of Power Regulation (273030, -1.42 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 16.2 healing_power points (2.74 DPS) | yes | Electrocutioner Lagnut (9447, -0.14 DPS) [dungeon]; Philanthropist's Ring (281635, -0.28 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -0.51 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -5.64 DPS) [dungeon]; Staff of the Friar (3415, -5.92 DPS) [dungeon]; Staff of Jordan (873, -9.74 DPS, sim-verified) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battered Viking Shield (9403, -1.33 DPS) [dungeon]; Orb of Souls (249395, -1.94 DPS) [crafted]; Beacon of Hope (9393, -2.47 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 0000000000000000-055000000000000000-5530001011513051)

Set DPS (verified): 280.3. Weights run: 16.4s. Verify run: 222.2s. 1053 eligible items had no known source.

Race sweep: dwarf 280.4 ± 0.1 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.201, intellect=1.928 ± 0.032, spirit=0.854 ± 0.023, mp5=2.982 ± 0.032, crit=0.489 ± 0.025 per rating point (14 rating = 1%, 6.842 per %), spell_haste=not significant (2.176 ± 0.566)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcatcher Halo (10630) | Sunken Temple: Atal'ai Warrior [dungeon] | 66.2 healing_power points (15.86 DPS) | yes | Papal Fez (9431, -0.50 DPS, sim-verified) [dungeon]; Gemburst Circlet (10751, -1.21 DPS) [quest]; Soothsayer's Headdress (17740, -1.78 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 35.4 healing_power points (8.49 DPS) | yes | Gemshard Heart (17707, -0.84 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.20 DPS) [world_drop]; Mindburst Medallion (11196, -3.26 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironfeather Shoulders (15067, -0.48 DPS, sim-verified) [crafted]; Lead Surveyor's Mantle (11842, -2.27 DPS) [dungeon]; Rotgrip Mantle (17732, -2.74 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Expeditionary's Cape (271921, -0.72 DPS, sim-verified) [rep]; Caretaker's Cape (19531, -0.85 DPS) [rep]; Imperial Red Cloak (8248, -1.45 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 79.3 healing_power points (18.99 DPS) | yes | Embrace of the Wind Serpent (12462, -0.33 DPS) [world]; Stormcloth Vest (10020, -3.46 DPS) [crafted]; Feathered Breastplate (8349, -4.91 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Bracers (252543, -0.04 DPS) [crafted]; Mender's Mail Bracers (252573, -0.04 DPS) [crafted]; Aristocratic Cuffs (12546, -0.31 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.22 DPS) [crafted]; Silkweb Gloves (11634, -2.64 DPS) [dungeon] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.41 DPS, sim-verified) [dungeon]; Gilded Waistcord (254081, -1.63 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 65.9 healing_power points (15.78 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.63 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -2.32 DPS, sim-verified) [dungeon]; Warchief Kilt (7760, -4.40 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 49.2 healing_power points (11.79 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.21 DPS) [crafted]; Vinerot Sandals (17748, -1.27 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 34.5 healing_power points (8.28 DPS) | yes | Runemender's Seal (273031, -2.12 DPS) [world]; Woodseed Hoop (17768, -2.22 DPS) [quest]; Lorekeeper's Ring (19523, -2.54 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 25.9 healing_power points (6.21 DPS) | yes | Runemender's Seal (273031, -0.05 DPS) [world]; Woodseed Hoop (17768, -0.15 DPS) [quest]; Lorekeeper's Ring (19523, -0.48 DPS) [rep] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -16.87 DPS) [dungeon]; Resurgence Rod (17743, -18.80 DPS, sim-verified) [quest]; Dreamstaff (249454, -22.13 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soulcatcher Halo; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Robes of Insight; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 1053, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 374.1. Weights run: 19.3s. Verify run: 506.6s. 2377 eligible items had no known source.

Race sweep: dwarf 374.0 ± 0.2 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.141, intellect=1.615 ± 0.040, spirit=0.773 ± 0.029, mp5=2.490 ± 0.038, crit=0.492 ± 0.025 per rating point (14 rating = 1%, 6.885 per %), spell_haste=not significant (0.908 ± 0.687)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 112.9 healing_power points (32.34 DPS) | yes | Devout Crown (16693, -9.33 DPS, sim-verified) [dungeon]; Black Dragonscale Helm (252605, -12.79 DPS) [crafted]; The Postmaster's Band (13390, -14.26 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 54.8 healing_power points (15.70 DPS) | yes | Lady Maye's Pendant (14558, -3.45 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -4.17 DPS) [dungeon]; Animated Chain Necklace (18723, -4.69 DPS) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 59.6 healing_power points (17.07 DPS) | yes | Theramore Shoulderpads (271953, -0.69 DPS, sim-verified) [rep]; Theramore Shoulders (271954, -1.15 DPS) [rep]; Mantle of Lost Hope (22234, -1.68 DPS) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -0.88 DPS, sim-verified) [quest]; Gracious Cape (18743, -2.21 DPS) [dungeon]; Caretaker's Cape (19530, -2.52 DPS) [rep] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 73.0 healing_power points (20.91 DPS) | yes | Robes of the Exalted (13346, -0.21 DPS) [dungeon]; Robes of Insight (940, -0.35 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -0.89 DPS, sim-verified) [world] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 47.8 healing_power points (13.69 DPS) | yes | Funeral Cuffs (12626, -1.25 DPS, sim-verified) [dungeon]; Bleak Howler Armguards (13208, -2.13 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -2.13 DPS) [rep] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -2.41 DPS, sim-verified) [dungeon]; Mender's Leather Gauntlets (252551, -4.33 DPS) [crafted]; Mender's Mail Gauntlets (252587, -4.33 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Sash of Mercy (14553, -0.83 DPS, sim-verified) [world_drop]; Stormseeker's Girdle (272399, -1.66 DPS) [vendor]; Devout Belt (16696, -2.95 DPS) [dungeon] |
| legs | Legplates of the Chromatic Defier (12945) | Legplates of the Chromatic Defier [quest] | 84.5 healing_power points (24.21 DPS) | yes | Ironfeather Leggings (252486, -2.59 DPS, sim-verified) [crafted]; Leggings of Arcana (12756, -3.16 DPS) [quest]; Red Dragonscale Leggings (252603, -3.68 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 56.6 healing_power points (16.23 DPS) | yes | Faith Healer's Boots (22247, -0.99 DPS, sim-verified) [dungeon]; Firemoss Boots (22275, -2.17 DPS) [dungeon]; Mender's Leather Boots (252472, -3.12 DPS) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -2.29 DPS) [dungeon]; Blessed Band of Light (272407, -2.80 DPS) [vendor]; Naglering (11669, -11.78 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.60 DPS, sim-verified) [dungeon]; Blessed Band of Light (272407, -2.72 DPS) [vendor]; Seal of Rivendare (13345, -3.22 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.55 DPS) [world_drop]; Weakness Analyzer (272438, -0.55 DPS) [vendor]; Serenity Field (272439, -1.03 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Argent Crusader (13249, -2.69 DPS, sim-verified) [quest]; Hammer of the Grand Crusader (18717, -3.91 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Tidal Totem (272431) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Sustaining (23200, -5.95 DPS, sim-verified) [dungeon] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Loomguard Armbraces; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Legplates of the Chromatic Defier; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Ankh of Life; trinket2: Burst of Knowledge; ranged: Tidal Totem

No-known-source sample (15 of 2377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 895.5. Weights run: 12.0s. Verify run: 403.0s. 2377 eligible items had no known source.

Race sweep: dwarf 896.0 ± 0.6 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.144, intellect=1.276 ± 0.050, spirit=0.682 ± 0.041, mp5=2.234 ± 0.046, crit=0.696 ± 0.035 per rating point (14 rating = 1%, 9.737 per %), spell_haste=not significant (3.980 ± 1.007)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 105.5 healing_power points (25.90 DPS) | yes | Devout Crown (16693, -10.96 DPS) [dungeon]; Black Dragonscale Helm (252605, -11.61 DPS) [crafted]; Mooncloth Circlet (14140, -23.11 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 47.8 healing_power points (11.73 DPS) | yes | Animated Chain Necklace (18723, -2.54 DPS, sim-verified) [dungeon]; Tooth of Gnarr (13141, -2.96 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.08 DPS) [world_drop] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 50.4 healing_power points (12.38 DPS) | yes | Mantle of Lost Hope (22234, -0.36 DPS) [dungeon]; Theramore Shoulderpads (271953, -0.98 DPS) [rep]; Theramore Shoulders (271954, -0.98 DPS) [rep] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.55 DPS) [quest]; Gracious Cape (18743, -2.60 DPS) [dungeon]; Caretaker's Cape (19530, -3.87 DPS, sim-verified) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | 67.3 healing_power points (16.51 DPS) | yes | Embrace of the Wind Serpent (12462, -0.25 DPS) [world]; Mooncloth Vest (14138, -0.68 DPS) [crafted]; Robes of Insight (940, -0.89 DPS) [world_drop] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 44.6 healing_power points (10.96 DPS) | yes | Bleak Howler Armguards (13208, -2.11 DPS) [dungeon]; Funeral Cuffs (12626, -2.15 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -2.65 DPS, sim-verified) [rep] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -0.66 DPS) [dungeon]; Mender's Leather Gauntlets (252551, -2.41 DPS) [crafted]; Mender's Mail Gauntlets (252587, -2.41 DPS) [crafted] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 61.2 healing_power points (15.02 DPS) | yes | Stormseeker's Girdle (272399, -3.57 DPS) [vendor]; Wisdom of the Timbermaw (19047, -3.60 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -3.90 DPS) [crafted] |
| legs | Legplates of the Chromatic Defier (12945) | Legplates of the Chromatic Defier [quest] | 77.4 healing_power points (18.99 DPS) | yes | Ironfeather Leggings (252486, -1.64 DPS) [crafted]; Sentinel's Silk Leggings (237815, -3.23 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -3.23 DPS) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.7 healing_power points (12.45 DPS) | yes | Faith Healer's Boots (22247, -1.80 DPS, sim-verified) [dungeon]; Swarmtender's Footpads (275609, -1.90 DPS) [crafted]; Broodwatcher's Treaders (275621, -1.90 DPS) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.78 DPS) [dungeon]; Blessed Band of Light (272407, -2.01 DPS) [vendor]; Naglering (11669, -25.52 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.86 DPS) [vendor]; Maiden's Circle (13001, -3.19 DPS) [world_drop]; Rosewine Circle (13178, -3.44 DPS, sim-verified) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+12.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.74 DPS) [world_drop]; Burst of Knowledge (11832, -0.77 DPS) [dungeon]; Weakness Analyzer (272438, -1.50 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -3.28 DPS) [dungeon]; Argent Crusader (13249, -5.31 DPS, sim-verified) [quest]; Staff of Metanoia (22394, -5.89 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Tidal Totem (272431) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Sustaining (23200, -10.68 DPS, sim-verified) [dungeon] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Robes of the Exalted; wrist: Loomguard Armbraces; hands: Theramore Handwraps; waist: Sash of Mercy; legs: Legplates of the Chromatic Defier; feet: Mooncloth Boots; finger1: Fordring's Seal; finger2: Band of Mending; trinket2: Serenity Field; ranged: Tidal Totem

No-known-source sample (15 of 2377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-000000000000000000-2530000010000000)

Set DPS (verified): 72.9. Weights run: 7.2s. Verify run: 45.7s. 275 eligible items had no known source.

Race sweep: orc 72.9 ± 0.0, windshaper-skyborne 72.8 ± 0.0, troll 72.8 ± 0.0, tauren 72.8 ± 0.0 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.119, intellect=0.235 ± 0.014, spirit=0.025 ± 0.004, mp5=2.902 ± 0.076, crit=0.076 ± 0.007 per rating point (14 rating = 1%, 1.067 per %), spell_haste=not significant (0.038 ± 0.073)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.01 DPS) | yes | Wisdom's Leather Hood (252507, -0.07 DPS, sim-verified) [crafted]; Totemic Leather Hood (252448, -0.50 DPS) [crafted]; Trapper's Leather Hood (252505, -0.90 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.1 healing_power points (0.00 DPS) | yes | Scholarly Pendant (277203, -0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | sim-verified (72.9 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.07 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.17 DPS) [crafted]; Slime-encrusted Pads (6461, -0.28 DPS, sim-verified) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.26 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.34 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 16.3 healing_power points (0.91 DPS) | yes | Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Bloody Apron (6226, -0.18 DPS) [dungeon]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.9 healing_power points (0.55 DPS) | yes | Owlbeard Bracers (16981, -0.47 DPS) [quest]; Owl Bracers (4796, -0.49 DPS) [vendor]; Tabitha's Cuffs (251486, -0.75 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 11.7 healing_power points (0.66 DPS) | yes | Wisdom's Leather Gloves (252499, -0.06 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.32 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 11.9 healing_power points (0.67 DPS) | yes | Wisdom's Leather Belt (252433, -0.20 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.29 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.39 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 21.5 healing_power points (1.20 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.56 DPS) [crafted]; Totemic Leather Pants (252446, -0.76 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 12.2 healing_power points (0.68 DPS) | yes | Pristine Boots (253889, -0.14 DPS) [crafted]; Black Whelp Slippers (252424, -0.16 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -0.24 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 10.8 healing_power points (0.60 DPS) | yes | Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Loop of Sacrifice (281673, -0.54 DPS) [quest]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop] |
| finger2 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | 2.0 healing_power points (0.11 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS) [quest]; Volcanic Rock Ring (12053, -0.07 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.07 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | 45.2 healing_power points (2.53 DPS) | yes | Living Root (6631, +0.00 DPS, sim-verified) [dungeon]; Kris of Orgrimmar (15443, -0.68 DPS) [quest]; The Stitcher (279874, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Advisor's Ring; finger2: Furlfeather Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield

### Band 30 (orc, 0000000000000000-000000000000000000-5530001010411000)

Set DPS (verified): 134.5. Weights run: 7.7s. Verify run: 47.7s. 469 eligible items had no known source.

Race sweep: orc 134.5 ± 0.0, windshaper-skyborne 134.4 ± 0.0, troll 133.4 ± 0.0, tauren 133.3 ± 0.0 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.133, intellect=1.125 ± 0.018, spirit=0.663 ± 0.012, mp5=2.220 ± 0.034, crit=0.170 ± 0.009 per rating point (14 rating = 1%, 2.385 per %), spell_haste=not significant (-0.158 ± 0.102)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.0 healing_power points (7.09 DPS) | yes | Enduring Cap (3020, -2.58 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -2.87 DPS) [crafted]; Wisdom's Leather Helm (252515, -3.06 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.2 healing_power points (1.37 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.29 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.1 healing_power points (3.67 DPS) | yes | Mantle of Woe (7750, -0.12 DPS) [quest]; Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Tanned Shoulderpads (270023, -0.65 DPS) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.63 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -0.89 DPS) [vendor]; Windsong Drape (15468, -1.53 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 32.4 healing_power points (6.21 DPS) | yes | Pristine Gown (253961, -0.44 DPS, sim-verified) [crafted]; Wisdom's Leather Tunic (252511, -0.70 DPS) [crafted]; Beguiler Robes (7728, -0.88 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.7 healing_power points (3.00 DPS) | yes | Mindthrust Bracers (1974, -0.50 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.28 DPS) [world_drop]; Nightsky Wristbands (6407, -1.33 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.4 healing_power points (3.52 DPS) | yes | Naga Battle Gloves (888, -0.70 DPS, sim-verified) [dungeon]; Wisdom's Leather Gloves (252499, -0.74 DPS) [crafted]; Pristine Gloves (253913, -0.77 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 29.7 healing_power points (5.70 DPS) | yes | Skycaller's Leather Belt (252522, -2.11 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -2.69 DPS) [dungeon]; Pristine Sash (253925, -2.73 DPS) [crafted] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 33.5 healing_power points (6.43 DPS) | yes | Pristine Leggings (253987, -0.07 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.79 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.17 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.5 healing_power points (5.85 DPS) | yes | Wisdom's Leather Boots (252444, -2.67 DPS) [crafted]; Acidic Walkers (9454, -2.85 DPS, sim-verified) [dungeon]; Black Whelp Slippers (252424, -3.27 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.7 healing_power points (3.78 DPS) | yes | Electrocutioner Lagnut (9447, -1.17 DPS) [dungeon]; Ring of Power Regulation (273030, -1.36 DPS) [dungeon]; Snake Hoop (6750, -1.55 DPS) [quest] |
| finger2 | Solid Gold Band (279845) | Opportunistic Education [quest] | 14.3 healing_power points (2.75 DPS) | yes | Electrocutioner Lagnut (9447, -0.13 DPS, sim-verified) [dungeon]; Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Snake Hoop (6750, -0.52 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.54 DPS) [dungeon]; Taskmaster Axe (5194, -3.39 DPS) [dungeon]; Staff of the Friar (3415, -4.63 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 17.0 healing_power points (3.26 DPS) | yes | Heart of Agamaggan (6694, -0.33 DPS, sim-verified) [dungeon]; Orb of Mystic Insight (249394, -0.62 DPS) [crafted]; Orb of Mistmantle (13031, -1.02 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Solid Gold Band; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 469, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 0000000000000000-000000000000000000-5530001011513051)

Set DPS (verified): 208.7. Weights run: 14.9s. Verify run: 172.3s. 787 eligible items had no known source.

Race sweep: orc 208.6 ± 0.1, windshaper-skyborne 208.6 ± 0.1, troll 206.2 ± 0.1, tauren 206.2 ± 0.1 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.257, intellect=1.878 ± 0.025, spirit=1.137 ± 0.023, mp5=3.685 ± 0.028, crit=0.316 ± 0.023 per rating point (14 rating = 1%, 4.421 per %), spell_haste=not significant (0.461 ± 0.273)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 64.2 healing_power points (16.38 DPS) | yes | Whitemane's Chapeau (7720, -3.11 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -5.17 DPS) [dungeon]; Corpseshroud (10574, -5.48 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 35.0 healing_power points (8.94 DPS) | yes | Necklace of Calisea (1714, -4.60 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -4.87 DPS) [dungeon]; Darkspear Warding Pendant (272074, -5.58 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 46.4 healing_power points (11.85 DPS) | yes | Sheepshear Mantle (13115, -3.75 DPS, sim-verified) [world_drop]; Green Silken Shoulders (7057, -4.53 DPS) [crafted]; Mantle of Woe (7750, -4.82 DPS) [quest] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272077, -0.74 DPS, sim-verified) [vendor]; Blackforge Cape (6424, -1.05 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.1 healing_power points (17.13 DPS) | yes | Dreamweave Vest (10021, -4.15 DPS) [crafted]; Death's Head Vestment (10581, -4.92 DPS) [dungeon]; Deathchill Armor (10764, -5.44 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Radiant Silver Bracers (4545, -0.26 DPS) [quest]; Enchanted Kodo Bracers (13119, -0.30 DPS) [world_drop]; Mindthrust Bracers (1974, -0.49 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.54 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.96 DPS) [crafted]; Stormcloth Gloves (10011, -1.11 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.8 healing_power points (11.19 DPS) | yes | Mender's Leather Belt (252523, -2.04 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -2.22 DPS) [dungeon]; Prefect's Belt (250559, -2.22 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 51.2 healing_power points (13.07 DPS) | yes | Stoneweaver Leggings (9407, -0.92 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -2.32 DPS) [crafted]; Pristine Leggings (253987, -2.64 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 45.1 healing_power points (11.51 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.89 DPS) [crafted]; Skycaller's Leather Shoes (252532, -3.39 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 27.1 healing_power points (6.91 DPS) | yes | Advisor's Ring (19520, -1.80 DPS) [rep]; Philanthropist's Ring (281635, -1.97 DPS) [quest]; Welken Ring (5011, -2.49 DPS) [world_drop] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 22.1 healing_power points (5.64 DPS) | yes | Advisor's Ring (19520, -0.52 DPS) [rep]; Philanthropist's Ring (281635, -0.70 DPS) [quest]; Welken Ring (5011, -1.22 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -8.12 DPS) [dungeon]; Death Speaker Scepter (2816, -8.96 DPS) [dungeon]; Staff of Jordan (873, -9.94 DPS, sim-verified) [world_drop] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forcestone Buckler (17508, -1.51 DPS, sim-verified) [quest]; Beacon of Hope (9393, -1.65 DPS) [dungeon]; Battered Viking Shield (9403, -2.21 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger2: Darkspear Signet; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Prophetic Cane

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 0000000000000000-055000000000000000-5530001011513051)

Set DPS (verified): 282.2. Weights run: 16.4s. Verify run: 227.6s. 1021 eligible items had no known source.

Race sweep: orc 282.1 ± 0.1, windshaper-skyborne 282.0 ± 0.1, troll 279.6 ± 0.1, tauren 279.5 ± 0.1 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.201, intellect=1.928 ± 0.032, spirit=0.854 ± 0.023, mp5=2.982 ± 0.032, crit=0.489 ± 0.025 per rating point (14 rating = 1%, 6.842 per %), spell_haste=not significant (2.176 ± 0.566)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcatcher Halo (10630) | Sunken Temple: Atal'ai Warrior [dungeon] | 66.2 healing_power points (15.86 DPS) | yes | Papal Fez (9431, -0.90 DPS) [dungeon]; Gemburst Circlet (10751, -1.21 DPS) [quest]; Soothsayer's Headdress (17740, -1.78 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 35.4 healing_power points (8.49 DPS) | yes | Gemshard Heart (17707, -0.74 DPS) [dungeon]; Horizon Choker (13085, -1.20 DPS) [world_drop]; Mindburst Medallion (11196, -3.26 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironfeather Shoulders (15067, +0.00 DPS) [crafted]; Lead Surveyor's Mantle (11842, -2.27 DPS) [dungeon]; Rotgrip Mantle (17732, -2.74 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272076, +0.00 DPS) [vendor]; Deep Woodlands Cloak (19121, -0.32 DPS) [quest]; Battle Healer's Cloak (19527, -0.85 DPS) [rep] |
| chest | Robes of Insight (940) | World drop [world_drop] | 79.3 healing_power points (18.99 DPS) | yes | Embrace of the Wind Serpent (12462, -0.33 DPS) [world]; Stormcloth Vest (10020, -3.46 DPS) [crafted]; Feathered Breastplate (8349, -4.91 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 39.9 healing_power points (9.57 DPS) | yes | Nethergeld Cuffs (254061, -0.52 DPS) [crafted]; Mender's Leather Bracers (252543, -0.55 DPS) [crafted]; Mender's Mail Bracers (252573, -0.55 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.22 DPS) [crafted]; Silkweb Gloves (11634, -2.64 DPS) [dungeon] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Waistcord (254081, -1.63 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 65.9 healing_power points (15.78 DPS) | yes | Jinxed Hoodoo Kilt (9474, -1.52 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.63 DPS) [dungeon]; Stone Guard's Inscribed Legplates (220839, -2.37 DPS) [vendor] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 49.2 healing_power points (11.79 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.21 DPS) [crafted]; Vinerot Sandals (17748, -1.27 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 34.5 healing_power points (8.28 DPS) | yes | Cyclopean Band (11824, -2.07 DPS) [dungeon]; Woodseed Hoop (17768, -2.22 DPS) [quest]; Advisor's Ring (19519, -2.54 DPS) [rep] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Woodseed Hoop (17768, -0.10 DPS) [quest]; Advisor's Ring (19519, -0.42 DPS) [rep] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enamored Water Spirit (20503, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enamored Water Spirit (20503, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -16.48 DPS) [quest]; Charstone Dirk (17710, -16.87 DPS) [dungeon]; Dreamstaff (249454, -22.13 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soulcatcher Halo; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Runemender's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 1021, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 376.1. Weights run: 19.3s. Verify run: 515.5s. 2334 eligible items had no known source.

Race sweep: orc 376.0 ± 0.2, windshaper-skyborne 375.9 ± 0.2, troll 374.0 ± 0.2, tauren 372.8 ± 0.2 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.141, intellect=1.615 ± 0.040, spirit=0.773 ± 0.029, mp5=2.490 ± 0.038, crit=0.492 ± 0.025 per rating point (14 rating = 1%, 6.885 per %), spell_haste=not significant (0.908 ± 0.687)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 112.9 healing_power points (32.34 DPS) | yes | Devout Crown (16693, -9.61 DPS, sim-verified) [dungeon]; Warlord's Mail Skullcap (231669, -9.90 DPS) [vendor]; Champion's Mail Skullcap (227168, -12.66 DPS) [vendor] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 54.8 healing_power points (15.70 DPS) | yes | Lady Maye's Pendant (14558, -3.57 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -4.17 DPS) [dungeon]; Animated Chain Necklace (18723, -4.69 DPS) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 59.6 healing_power points (17.07 DPS) | yes | Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Champion's Mail Epaulets (227166, -0.74 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.96 DPS, sim-verified) [vendor] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.20 DPS, sim-verified) [quest]; Gracious Cape (18743, -2.21 DPS) [dungeon]; Battle Healer's Cloak (19526, -2.52 DPS) [rep] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 73.0 healing_power points (20.91 DPS) | yes | Warlord's Mail Chestguard (231668, +0.00 DPS) [vendor]; Legionnaire's Mail Chestguard (227171, -0.01 DPS) [vendor]; Embrace of the Wind Serpent (12462, -0.97 DPS, sim-verified) [world] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 47.8 healing_power points (13.69 DPS) | yes | Funeral Cuffs (12626, -1.75 DPS, sim-verified) [dungeon]; Bleak Howler Armguards (13208, -2.13 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -2.13 DPS) [rep] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -2.45 DPS, sim-verified) [dungeon]; Mender's Leather Gauntlets (252551, -4.33 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Sash of Mercy (14553, -0.48 DPS, sim-verified) [world_drop]; Stormseeker's Girdle (272399, -1.66 DPS) [vendor]; Devout Belt (16696, -2.95 DPS) [dungeon] |
| legs | Legplates of the Chromatic Defier (12945) | Legplates of the Chromatic Defier [quest] | 84.5 healing_power points (24.21 DPS) | yes | General's Mail Pants (231670, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -2.90 DPS, sim-verified) [crafted]; Legionnaire's Mail Pants (227167, -2.92 DPS) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 56.6 healing_power points (16.23 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Faith Healer's Boots (22247, -1.15 DPS, sim-verified) [dungeon]; Firemoss Boots (22275, -2.17 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -2.29 DPS) [dungeon]; Blessed Band of Light (272407, -2.80 DPS) [vendor]; Naglering (11669, -11.88 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.90 DPS, sim-verified) [dungeon]; Blessed Band of Light (272407, -2.72 DPS) [vendor]; Seal of Rivendare (13345, -3.22 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.55 DPS) [world_drop]; Weakness Analyzer (272438, -0.55 DPS) [vendor]; Serenity Field (272439, -0.91 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Argent Crusader (13249, -3.11 DPS, sim-verified) [quest]; Hammer of the Grand Crusader (18717, -3.91 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Tidal Totem (272431) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Sustaining (23200, -6.24 DPS, sim-verified) [dungeon] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Loomguard Armbraces; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Legplates of the Chromatic Defier; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Ankh of Life; trinket2: Burst of Knowledge; ranged: Tidal Totem

No-known-source sample (15 of 2334, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 900.3. Weights run: 12.0s. Verify run: 404.8s. 2334 eligible items had no known source.

Race sweep: windshaper-skyborne 904.8 ± 0.6, orc 900.4 ± 0.6, troll 900.0 ± 0.6, tauren 893.8 ± 0.5 (heal_score); windshaper-skyborne beats the guide race orc by 4.4 (0.48%) ± 0.8, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.144, intellect=1.276 ± 0.050, spirit=0.682 ± 0.041, mp5=2.234 ± 0.046, crit=0.696 ± 0.035 per rating point (14 rating = 1%, 9.737 per %), spell_haste=not significant (3.980 ± 1.007)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 105.5 healing_power points (25.90 DPS) | yes | Warlord's Mail Skullcap (231669, -7.39 DPS) [vendor]; Champion's Mail Skullcap (227168, -9.79 DPS) [vendor]; Mooncloth Circlet (14140, -26.01 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 47.8 healing_power points (11.73 DPS) | yes | Tooth of Gnarr (13141, -2.96 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.08 DPS) [world_drop]; Animated Chain Necklace (18723, -4.24 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 50.4 healing_power points (12.38 DPS) | yes | Champion's Mail Epaulets (227166, +0.00 DPS) [vendor]; Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Mantle of Lost Hope (22234, -1.77 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.55 DPS) [quest]; Gracious Cape (18743, -2.60 DPS) [dungeon]; Battle Healer's Cloak (19526, -4.32 DPS, sim-verified) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | 67.3 healing_power points (16.51 DPS) | yes | Legionnaire's Mail Chestguard (227171, +0.00 DPS) [vendor]; Warlord's Mail Chestguard (231668, +0.00 DPS) [vendor]; Embrace of the Wind Serpent (12462, -0.25 DPS) [world] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 44.6 healing_power points (10.96 DPS) | yes | Bleak Howler Armguards (13208, -2.11 DPS) [dungeon]; Funeral Cuffs (12626, -2.15 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -4.04 DPS, sim-verified) [rep] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Blood Guard's Mail Gloves (227169, -2.02 DPS) [vendor]; Hands of the Exalted Herald (12554, -2.36 DPS, sim-verified) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 61.2 healing_power points (15.02 DPS) | yes | Stormseeker's Girdle (272399, -3.57 DPS) [vendor]; Wisdom of the Timbermaw (19047, -3.74 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -3.90 DPS) [crafted] |
| legs | Legplates of the Chromatic Defier (12945) | Legplates of the Chromatic Defier [quest] | 77.4 healing_power points (18.99 DPS) | yes | General's Mail Pants (231670, +0.00 DPS) [vendor]; Legionnaire's Mail Pants (227167, -1.46 DPS) [vendor]; Ironfeather Leggings (252486, -3.38 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.7 healing_power points (12.45 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Faith Healer's Boots (22247, -1.33 DPS, sim-verified) [dungeon]; Blood Guard's Mail Boots (227170, -1.58 DPS) [vendor] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.78 DPS) [dungeon]; Blessed Band of Light (272407, -2.01 DPS) [vendor]; Naglering (11669, -25.66 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.86 DPS) [vendor]; Maiden's Circle (13001, -3.19 DPS) [world_drop]; Rosewine Circle (13178, -4.92 DPS, sim-verified) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+13.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.74 DPS) [world_drop]; Burst of Knowledge (11832, -0.77 DPS) [dungeon]; Weakness Analyzer (272438, -1.52 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -3.28 DPS) [dungeon]; Staff of Metanoia (22394, -5.89 DPS) [dungeon]; Argent Crusader (13249, -8.25 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Tidal Totem (272431) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Sustaining (23200, -11.88 DPS, sim-verified) [dungeon] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderguards; back: Drape of Recovery; chest: Robes of the Exalted; wrist: Loomguard Armbraces; hands: Raider Handwraps; waist: Sash of Mercy; legs: Legplates of the Chromatic Defier; feet: Mooncloth Boots; finger1: Fordring's Seal; finger2: Band of Mending; trinket2: Serenity Field; ranged: Tidal Totem

No-known-source sample (15 of 2334, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

