# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-2530000010000000)

Set DPS (verified): 73.1. Weights run: 4.5s. Verify run: 23.1s. 286 eligible items had no known source.

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
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 10.8 healing_power points (0.60 DPS) | yes | Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Furlfeather Ring (286745, -0.55 DPS) [world]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 healing_power points (0.31 DPS) | yes | Furlfeather Ring (286745, -0.25 DPS) [world]; Lavishly Jeweled Ring (1156, -0.26 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 49.3 healing_power points (2.76 DPS) | yes | Taskmaster Axe (5194, -0.18 DPS, sim-verified) [dungeon]; Living Root (6631, -0.24 DPS) [dungeon]; Evocator's Blade (2567, -1.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (dwarf, 0000000000000000-000000000000000000-5530001010411000)

Set DPS (verified): 134.7. Weights run: 4.9s. Verify run: 28.6s. 475 eligible items had no known source.

Race sweep: dwarf 134.7 ± 0.0 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.019, intellect=0.689 ± 0.004, spirit=0.457 ± 0.004, mp5=1.405 ± 0.004, crit=0.079 ± 0.003 per rating point (14 rating = 1%, 1.112 per %), spell_haste=-0.611 ± 0.064

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 35.7 healing_power points (7.43 DPS) | yes | Filigreed Pristine Circlet (253975, -2.62 DPS, sim-verified) [crafted]; Wisdom's Leather Helm (252515, -3.07 DPS) [crafted]; Pristine Circlet (253949, -3.69 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.6 healing_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.09 DPS) [world]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Seal of the Expedition (271920, -0.24 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.2 healing_power points (3.16 DPS) | yes | Death Speaker Mantle (6685, -0.12 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.83 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prelacy Cape (7004, -0.13 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.87 DPS) [dungeon]; Repairman's Cape (9605, -0.93 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 27.6 healing_power points (5.73 DPS) | yes | Pristine Gown (253961, -0.27 DPS, sim-verified) [crafted]; Wisdom's Leather Tunic (252511, -0.56 DPS) [crafted]; Beguiler Robes (7728, -1.38 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.2 healing_power points (2.75 DPS) | yes | Mindthrust Bracers (1974, -0.52 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.60 DPS) [world_drop]; Nightsky Wristbands (6407, -1.60 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.1 healing_power points (3.55 DPS) | yes | Naga Battle Gloves (888, -0.47 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.83 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.90 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 25.9 healing_power points (5.38 DPS) | yes | Skycaller's Leather Belt (252522, -1.73 DPS, sim-verified) [crafted]; Pristine Sash (253925, -2.52 DPS) [crafted]; Belt of Arugal (6392, -2.60 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 29.1 healing_power points (6.05 DPS) | yes | Wisdom's Leather Leggings (252519, -0.02 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.65 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.07 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 26.7 healing_power points (5.54 DPS) | yes | Wisdom's Leather Boots (252444, -2.52 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.98 DPS) [dungeon]; Black Whelp Slippers (252424, -3.10 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 16.3 healing_power points (3.40 DPS) | yes | Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Ring of Power Regulation (273030, -1.27 DPS) [dungeon]; Lorekeeper's Ring (19525, -1.36 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 14.0 healing_power points (2.92 DPS) | yes | Ring of Power Regulation (273030, -0.79 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.84 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19525, -0.88 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -2.19 DPS) [dungeon]; Staff of Westfall (2042, -2.87 DPS) [quest]; Staff of the Friar (3415, -4.08 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.7 healing_power points (3.27 DPS) | yes | Heart of Agamaggan (6694, -0.24 DPS, sim-verified) [dungeon]; Eye of Paleth (2943, -0.57 DPS) [quest]; Orb of Mystic Insight (249394, -0.96 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 0000000000000000-000000000000000000-5530001011513051)

Set DPS (verified): 207.8. Weights run: 8.5s. Verify run: 51.8s. 804 eligible items had no known source.

Race sweep: dwarf 207.7 ± 0.1 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.246, intellect=1.342 ± 0.020, spirit=0.819 ± 0.017, mp5=2.710 ± 0.020, crit=0.259 ± 0.019 per rating point (14 rating = 1%, 3.625 per %), spell_haste=not significant (-0.034 ± 0.273)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 52.2 healing_power points (14.90 DPS) | yes | Whitemane's Chapeau (7720, -2.25 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -4.07 DPS) [world_drop]; Miner's Hat of the Deep (9429, -4.96 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.7 healing_power points (8.78 DPS) | yes | Necklace of Calisea (1714, -3.49 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -4.85 DPS) [dungeon]; Seal of the Expedition (271919, -6.09 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 39.4 healing_power points (11.26 DPS) | yes | Sheepshear Mantle (13115, -3.06 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -4.25 DPS) [crafted]; Green Silken Shoulders (7057, -4.76 DPS) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (207.7 DPS) | yes | Expeditionary's Cape (271922, -0.44 DPS, sim-verified) [rep]; Prelacy Cape (7004, -1.76 DPS) [quest]; Repairman's Cape (9605, -2.28 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.6 healing_power points (18.43 DPS) | yes | Deathchill Armor (10764, -4.84 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -6.57 DPS) [crafted]; Death's Head Vestment (10581, -7.19 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.1 healing_power points (4.89 DPS) | yes | Mindthrust Bracers (1974, -0.27 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.87 DPS) [crafted]; Guardian Leather Bracers (4260, -0.88 DPS) [crafted] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-verified (207.7 DPS) | yes | Gilded Handwraps (254021, -0.54 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.77 DPS) [crafted]; Dreamweave Gloves (10019, -1.13 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 37.6 healing_power points (10.75 DPS) | yes | Mender's Leather Belt (252523, -1.63 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.79 DPS) [crafted]; Deathmage Sash (10771, -3.01 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 43.5 healing_power points (12.43 DPS) | yes | Stoneweaver Leggings (9407, -0.56 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -2.11 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -2.22 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 39.1 healing_power points (11.17 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.84 DPS) [crafted]; Skycaller's Leather Shoes (252532, -3.61 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.8 healing_power points (6.51 DPS) | yes | Lorekeeper's Ring (19524, -1.61 DPS) [rep]; Philanthropist's Ring (281635, -1.74 DPS) [quest]; Electrocutioner Lagnut (9447, -2.30 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 21.1 healing_power points (6.03 DPS) | yes | Lorekeeper's Ring (19524, -0.98 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -1.26 DPS) [quest]; Electrocutioner Lagnut (9447, -1.82 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (207.7 DPS) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (207.7 DPS) | yes | Staff of Jordan (873, -8.90 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -9.30 DPS) [dungeon]; Staff of the Friar (3415, -9.91 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (207.7 DPS) | yes | Beacon of Hope (9393, -1.83 DPS, sim-verified) [dungeon]; Battered Viking Shield (9403, -2.92 DPS) [dungeon]; Orb of Souls (249395, -3.85 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 0000000000000000-055000000000000000-5530001011513051)

Set DPS (verified): 281.1. Weights run: 10.1s. Verify run: 131.2s. 1050 eligible items had no known source.

Race sweep: dwarf 281.0 ± 0.1 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.151, intellect=1.349 ± 0.024, spirit=0.583 ± 0.018, mp5=2.087 ± 0.024, crit=0.338 ± 0.018 per rating point (14 rating = 1%, 4.730 per %), spell_haste=2.342 ± 0.449

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcatcher Halo (10630) | Sunken Temple: Atal'ai Warrior [dungeon] | 51.7 healing_power points (17.57 DPS) | yes | Cassandra's Grace (13102, -0.84 DPS) [world_drop]; Soothsayer's Headdress (17740, -1.30 DPS) [dungeon]; Papal Fez (9431, -1.33 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.8 healing_power points (10.46 DPS) | yes | Gemshard Heart (17707, -0.67 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -3.25 DPS) [world_drop]; Mindburst Medallion (11196, -4.68 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 49.2 healing_power points (16.71 DPS) | yes | Ironfeather Shoulders (15067, -0.87 DPS) [crafted]; Nethergeld Shoulders (254049, -3.02 DPS) [crafted]; Mender's Leather Shoulders (252538, -3.05 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19531, -1.09 DPS, sim-verified) [rep]; Expeditionary's Cape (271921, -1.29 DPS) [rep] |
| chest | Robes of Insight (940) | World drop [world_drop] | 65.4 healing_power points (22.20 DPS) | yes | Stormcloth Vest (10020, -3.66 DPS, sim-verified) [crafted]; Feathered Breastplate (8349, -6.41 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -7.29 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 32.4 healing_power points (10.99 DPS) | yes | Mender's Leather Bracers (252543, -0.14 DPS) [crafted]; Mender's Mail Bracers (252573, -0.14 DPS) [crafted]; Aristocratic Cuffs (12546, -0.38 DPS) [dungeon] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.42 DPS) [crafted]; Jumanza Grips (18083, -4.28 DPS) [dungeon] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 46.2 healing_power points (15.69 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.87 DPS) [dungeon]; Gilded Waistcord (254081, -1.71 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 56.9 healing_power points (19.34 DPS) | yes | Jinxed Hoodoo Kilt (9474, -2.25 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -4.47 DPS) [dungeon]; Warchief Kilt (7760, -5.67 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 43.1 healing_power points (14.65 DPS) | yes | Mender's Leather Boots (252472, -0.10 DPS) [crafted]; Mender's Mail Sabatons (252579, -0.10 DPS) [crafted]; Mender's Leather Shoes (252533, -1.91 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 27.4 healing_power points (9.31 DPS) | yes | Cyclopean Band (11824, -2.25 DPS) [dungeon]; Lorekeeper's Ring (19523, -2.40 DPS) [rep]; Snake Hoop (6750, -2.68 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 21.2 healing_power points (7.19 DPS) | yes | Cyclopean Band (11824, -0.14 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.28 DPS) [rep]; Snake Hoop (6750, -0.56 DPS) [quest] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, +0.00 DPS) [rep]; Uther's Strength (11302, -0.82 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -7.51 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -20.18 DPS) [quest]; Dreamstaff (249454, -26.59 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soulcatcher Halo; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Robes of Insight; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Brainlash; finger2: Runemender's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 1050, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 364.6. Weights run: 11.8s. Verify run: 398.3s. 2350 eligible items had no known source.

Race sweep: dwarf 364.6 ± 0.2 (heal_score); the guide race dwarf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.253, intellect=2.545 ± 0.065, spirit=1.281 ± 0.049, mp5=3.995 ± 0.060, crit=0.798 ± 0.047 per rating point (14 rating = 1%, 11.168 per %), spell_haste=not significant (3.007 ± 1.267)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 136.7 healing_power points (23.16 DPS) | yes | Devout Crown (16693, -6.17 DPS) [dungeon]; Horns of Eranikus (10833, -7.49 DPS) [dungeon]; Black Dragonscale Helm (252605, -11.02 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 76.6 healing_power points (12.97 DPS) | yes | Lady Maye's Pendant (14558, -3.44 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -4.91 DPS) [dungeon]; Animated Chain Necklace (18723, -5.87 DPS) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 84.7 healing_power points (14.34 DPS) | yes | Theramore Shoulders (271954, -0.68 DPS) [rep]; Theramore Shoulderpads (271953, -0.93 DPS, sim-verified) [rep]; Ironfeather Shoulders (15067, -1.45 DPS) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Shroud of the Exile (15421, -0.10 DPS) [quest]; Royal Tribunal Cloak (13376, -0.21 DPS) [dungeon]; Expeditionary's Cape (271907, -1.65 DPS, sim-verified) [rep] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 101.2 healing_power points (17.14 DPS) | yes | Devout Robe (16690, -0.76 DPS, sim-verified) [dungeon]; Robes of Insight (940, -1.21 DPS) [world_drop]; Magister's Robes (16688, -1.31 DPS) [dungeon] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 58.8 healing_power points (9.96 DPS) | yes | Funeral Cuffs (12626, -0.88 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -1.26 DPS) [rep]; Bleak Howler Armguards (13208, -1.99 DPS, sim-verified) [dungeon] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Hands of the Exalted Herald (12554, -2.24 DPS) [dungeon]; Theramore Handguards (271950, -3.94 DPS) [rep]; Mender's Leather Gauntlets (252551, -4.62 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 78.2 healing_power points (13.23 DPS) | yes | Devout Belt (16696, -0.80 DPS) [dungeon]; Stormseeker's Girdle (272399, -1.29 DPS, sim-verified) [vendor]; Grimgore Noose (13403, -1.45 DPS) [dungeon] |
| legs | Red Dragonscale Leggings (252603) | Leatherworking [crafted] | sim-verified (364.6 DPS) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Ironfeather Leggings (252486, -1.09 DPS) [crafted]; Ghostloom Leggings (14545, -1.47 DPS) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 76.3 healing_power points (12.92 DPS) | yes | Faith Healer's Boots (22247, -0.70 DPS, sim-verified) [dungeon]; Firemoss Boots (22275, -1.77 DPS) [dungeon]; Dragonrider Boots (18102, -2.98 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.89 DPS) [dungeon]; Rosewine Circle (13178, -1.44 DPS) [dungeon]; Naglering (11669, -11.94 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.73 DPS) [dungeon]; Channeler's Ring (272406, -1.08 DPS) [vendor]; Seal of Rivendare (13345, -2.83 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep]; Serenity Field (272439, -1.09 DPS, sim-verified) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.35 DPS) [quest]; Serenity Field (272439, -1.86 DPS) [vendor]; Enriched Seal (271903, -2.03 DPS) [rep] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hammer of the Grand Crusader (18717, -2.10 DPS) [dungeon]; Argent Crusader (13249, -3.30 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Loomguard Armbraces; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Red Dragonscale Leggings; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Ankh of Life; trinket2: Burst of Knowledge

No-known-source sample (15 of 2350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 865.3. Weights run: 7.5s. Verify run: 198.1s. 2350 eligible items had no known source.

Race sweep: dwarf 864.6 ± 0.6 (heal_score); the guide race dwarf is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.441, intellect=1.148 ± 0.057, spirit=0.464 ± 0.045, mp5=1.470 ± 0.058, crit=0.656 ± 0.060 per rating point (14 rating = 1%, 9.181 per %), spell_haste=5.412 ± 1.346

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 100.2 healing_power points (64.12 DPS) | yes | Mooncloth Circlet (14140, -25.89 DPS, sim-verified) [crafted]; Devout Crown (16693, -29.22 DPS) [dungeon]; Black Dragonscale Helm (252605, -30.31 DPS) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 43.3 healing_power points (27.73 DPS) | yes | Animated Chain Necklace (18723, -5.82 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -6.72 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -9.04 DPS) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 47.0 healing_power points (30.09 DPS) | yes | Theramore Shoulderpads (271953, -2.56 DPS) [rep]; Theramore Shoulders (271954, -2.56 DPS) [rep]; Mantle of Lost Hope (22234, -3.03 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -4.43 DPS, sim-verified) [rep]; Shroud of the Exile (15421, -6.49 DPS) [quest]; Arcanoweave Cloak (272411, -7.48 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.7 healing_power points (39.50 DPS) | yes | Robes of the Exalted (13346, -0.59 DPS) [dungeon]; Robes of Insight (940, -0.74 DPS) [world_drop]; Mooncloth Vest (14138, -2.51 DPS) [crafted] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 41.4 healing_power points (26.53 DPS) | yes | Funeral Cuffs (12626, -4.72 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -5.26 DPS, sim-verified) [rep]; Bleak Howler Armguards (13208, -6.15 DPS) [dungeon] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -4.23 DPS, sim-verified) [dungeon]; Mender's Leather Gauntlets (252551, -6.08 DPS) [crafted]; Mender's Mail Gauntlets (252587, -6.08 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Mender's Leather Waistguard (252477, -4.79 DPS) [crafted]; Mender's Mail Belt (252591, -4.79 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 67.9 healing_power points (43.43 DPS) | yes | Sentinel's Lizardhide Pants (237817, -3.84 DPS) [vendor]; Sentinel's Silk Leggings (237815, -4.79 DPS, sim-verified) [vendor]; Ghoul Skin Leggings (18682, -6.71 DPS) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 46.1 healing_power points (29.52 DPS) | yes | Faith Healer's Boots (22247, -0.49 DPS) [dungeon]; Swarmtender's Footpads (275609, -1.99 DPS) [crafted]; Broodwatcher's Treaders (275621, -1.99 DPS) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -1.32 DPS) [dungeon]; Rosewine Circle (13178, -7.09 DPS) [dungeon]; Naglering (11669, -27.09 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Rosewine Circle (13178, -1.21 DPS) [dungeon]; Maiden's Circle (13001, -3.77 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+13.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.13 DPS, sim-verified) [vendor]; Uther's Strength (11302, -1.92 DPS) [world_drop]; Second Wind (11819, -3.84 DPS) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -9.33 DPS, sim-verified) [quest]; Hammer of the Grand Crusader (18717, -9.43 DPS) [dungeon]; Staff of Metanoia (22394, -16.03 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Stormcloth Vest; wrist: Loomguard Armbraces; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Mooncloth Boots; finger1: Fordring's Seal; finger2: Blessed Band of Light; trinket2: Serenity Field

No-known-source sample (15 of 2350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-000000000000000000-2530000010000000)

Set DPS (verified): 72.8. Weights run: 4.5s. Verify run: 23.9s. 275 eligible items had no known source.

Race sweep: orc 72.8 ± 0.0, windshaper-skyborne 72.7 ± 0.0, troll 72.7 ± 0.0, tauren 72.7 ± 0.0 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.119, intellect=0.235 ± 0.014, spirit=0.025 ± 0.004, mp5=2.902 ± 0.076, crit=0.076 ± 0.007 per rating point (14 rating = 1%, 1.067 per %), spell_haste=not significant (0.038 ± 0.073)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.01 DPS) | yes | Wisdom's Leather Hood (252507, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.50 DPS) [crafted]; Trapper's Leather Hood (252505, -0.90 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.1 healing_power points (0.00 DPS) | yes | Scholarly Pendant (277203, -0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.07 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.17 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Black Whelp Cloak (7283, -0.34 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 16.3 healing_power points (0.91 DPS) | yes | Filigreed Pristine Gown (253901, -0.11 DPS) [crafted]; Bloody Apron (6226, -0.18 DPS) [dungeon]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.9 healing_power points (0.55 DPS) | yes | Tabitha's Cuffs (251486, -0.39 DPS) [quest]; Owlbeard Bracers (16981, -0.47 DPS) [quest]; Owl Bracers (4796, -0.49 DPS) [vendor] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 11.7 healing_power points (0.66 DPS) | yes | Wisdom's Leather Gloves (252499, -0.04 DPS) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.32 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 11.9 healing_power points (0.67 DPS) | yes | Wisdom's Leather Belt (252433, -0.17 DPS) [crafted]; Novice Ardent's Sash (253887, -0.29 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.39 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 21.5 healing_power points (1.20 DPS) | yes | Filigreed Pristine Leggings (253937, -0.11 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.56 DPS) [crafted]; Totemic Leather Pants (252446, -0.76 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 12.2 healing_power points (0.68 DPS) | yes | Black Whelp Slippers (252424, -0.13 DPS) [crafted]; Pristine Boots (253889, -0.14 DPS) [crafted]; Spidersilk Boots (4320, -0.24 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 10.8 healing_power points (0.60 DPS) | yes | Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Loop of Sacrifice (281673, -0.54 DPS) [quest]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop] |
| finger2 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | sim-verified (72.8 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Loop of Sacrifice (281673, +0.00 DPS) [quest]; Volcanic Rock Ring (12053, -0.02 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | 45.2 healing_power points (2.53 DPS) | yes | Living Root (6631, -0.00 DPS) [dungeon]; Kris of Orgrimmar (15443, -0.68 DPS) [quest]; The Stitcher (279874, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Advisor's Ring; finger2: Furlfeather Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield

### Band 30 (orc, 0000000000000000-000000000000000000-5530001010411000)

Set DPS (verified): 135.1. Weights run: 4.9s. Verify run: 29.6s. 469 eligible items had no known source.

Race sweep: orc 135.2 ± 0.0, windshaper-skyborne 135.0 ± 0.0, troll 134.4 ± 0.0, tauren 134.4 ± 0.0 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.019, intellect=0.689 ± 0.004, spirit=0.457 ± 0.004, mp5=1.405 ± 0.004, crit=0.079 ± 0.003 per rating point (14 rating = 1%, 1.112 per %), spell_haste=-0.611 ± 0.064

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 35.7 healing_power points (7.43 DPS) | yes | Filigreed Pristine Circlet (253975, -2.53 DPS, sim-verified) [crafted]; Wisdom's Leather Helm (252515, -3.07 DPS) [crafted]; Pristine Circlet (253949, -3.69 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.6 healing_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.09 DPS) [world]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.24 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.2 healing_power points (3.16 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS, sim-verified) [dungeon]; Tanned Shoulderpads (270023, -0.43 DPS) [quest]; Magician's Mantle (12998, -0.83 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.55 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -1.65 DPS) [vendor]; Windsong Drape (15468, -1.66 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 27.6 healing_power points (5.73 DPS) | yes | Pristine Gown (253961, -0.23 DPS, sim-verified) [crafted]; Wisdom's Leather Tunic (252511, -0.56 DPS) [crafted]; Beguiler Robes (7728, -1.38 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.2 healing_power points (2.75 DPS) | yes | Mindthrust Bracers (1974, -0.54 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.60 DPS) [world_drop]; Nightsky Wristbands (6407, -1.60 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.1 healing_power points (3.55 DPS) | yes | Naga Battle Gloves (888, -0.46 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.83 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.90 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 25.9 healing_power points (5.38 DPS) | yes | Skycaller's Leather Belt (252522, -1.56 DPS, sim-verified) [crafted]; Pristine Sash (253925, -2.52 DPS) [crafted]; Belt of Arugal (6392, -2.60 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 29.1 healing_power points (6.05 DPS) | yes | Wisdom's Leather Leggings (252519, -0.02 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.65 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.07 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 26.7 healing_power points (5.54 DPS) | yes | Wisdom's Leather Boots (252444, -2.32 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.98 DPS) [dungeon]; Black Whelp Slippers (252424, -3.10 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 16.3 healing_power points (3.40 DPS) | yes | Solid Gold Band (279845, -0.63 DPS) [quest]; Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Ring of Power Regulation (273030, -1.27 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 14.0 healing_power points (2.92 DPS) | yes | Electrocutioner Lagnut (9447, -0.38 DPS) [dungeon]; Solid Gold Band (279845, -0.64 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -0.79 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -2.19 DPS) [dungeon]; Staff of the Friar (3415, -3.96 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -4.02 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.7 healing_power points (3.27 DPS) | yes | Heart of Agamaggan (6694, -0.23 DPS, sim-verified) [dungeon]; Orb of Mystic Insight (249394, -0.96 DPS) [crafted]; Orb of Mistmantle (13031, -1.02 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 469, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 0000000000000000-000000000000000000-5530001011513051)

Set DPS (verified): 208.1. Weights run: 8.5s. Verify run: 53.8s. 787 eligible items had no known source.

Race sweep: windshaper-skyborne 208.2 ± 0.1, orc 208.2 ± 0.1, troll 206.4 ± 0.1, tauren 206.3 ± 0.1 (heal_score); windshaper-skyborne beats the guide race orc by 0.1 (0.03%) ± 0.1, inside error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.246, intellect=1.342 ± 0.020, spirit=0.819 ± 0.017, mp5=2.710 ± 0.020, crit=0.259 ± 0.019 per rating point (14 rating = 1%, 3.625 per %), spell_haste=not significant (-0.034 ± 0.273)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 52.2 healing_power points (14.90 DPS) | yes | Whitemane's Chapeau (7720, -2.38 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -4.07 DPS) [world_drop]; Miner's Hat of the Deep (9429, -4.96 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.7 healing_power points (8.78 DPS) | yes | Necklace of Calisea (1714, -3.82 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -4.85 DPS) [dungeon]; Darkspear Warding Pendant (272074, -6.09 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 39.4 healing_power points (11.26 DPS) | yes | Sheepshear Mantle (13115, -3.12 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -4.25 DPS) [crafted]; Green Silken Shoulders (7057, -4.76 DPS) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.56 DPS, sim-verified) [vendor]; Blackforge Cape (6424, -2.31 DPS) [dungeon]; Glowing Thresher Cape (6901, -2.44 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.6 healing_power points (18.43 DPS) | yes | Deathchill Armor (10764, -4.79 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -6.57 DPS) [crafted]; Death's Head Vestment (10581, -7.19 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.1 healing_power points (4.89 DPS) | yes | Mindthrust Bracers (1974, -0.21 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -0.68 DPS) [quest]; Earthen Silk Cuffs (254019, -0.87 DPS) [crafted] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.26 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.77 DPS) [crafted]; Dreamweave Gloves (10019, -1.13 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 37.6 healing_power points (10.75 DPS) | yes | Mender's Leather Belt (252523, -1.45 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.79 DPS) [crafted]; Deathmage Sash (10771, -3.01 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 43.5 healing_power points (12.43 DPS) | yes | Stoneweaver Leggings (9407, -0.50 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -2.11 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -2.22 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 39.1 healing_power points (11.17 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.84 DPS) [crafted]; Skycaller's Leather Shoes (252532, -3.61 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.8 healing_power points (6.51 DPS) | yes | Advisor's Ring (19520, -1.61 DPS) [rep]; Philanthropist's Ring (281635, -1.74 DPS) [quest]; Darkspear Signet (272070, -1.86 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 21.1 healing_power points (6.03 DPS) | yes | Advisor's Ring (19520, -0.98 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -1.26 DPS) [quest]; Darkspear Signet (272070, -1.39 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -8.43 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -9.30 DPS) [dungeon]; Staff of the Friar (3415, -9.91 DPS) [dungeon] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beacon of Hope (9393, -1.14 DPS, sim-verified) [dungeon]; Forcestone Buckler (17508, -1.19 DPS) [quest]; Battered Viking Shield (9403, -2.15 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Prophetic Cane

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 0000000000000000-055000000000000000-5530001011513051)

Set DPS (verified): 282.0. Weights run: 10.1s. Verify run: 138.6s. 1018 eligible items had no known source.

Race sweep: windshaper-skyborne 282.1 ± 0.1, orc 281.9 ± 0.1, troll 279.8 ± 0.1, tauren 279.6 ± 0.1 (heal_score); windshaper-skyborne beats the guide race orc by 0.2 (0.07%) ± 0.2, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.151, intellect=1.349 ± 0.024, spirit=0.583 ± 0.018, mp5=2.087 ± 0.024, crit=0.338 ± 0.018 per rating point (14 rating = 1%, 4.730 per %), spell_haste=2.342 ± 0.449

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcatcher Halo (10630) | Sunken Temple: Atal'ai Warrior [dungeon] | 51.7 healing_power points (17.57 DPS) | yes | Cassandra's Grace (13102, -0.84 DPS) [world_drop]; Papal Fez (9431, -0.93 DPS, sim-verified) [dungeon]; Soothsayer's Headdress (17740, -1.30 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.8 healing_power points (10.46 DPS) | yes | Gemshard Heart (17707, -0.76 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -3.25 DPS) [world_drop]; Mindburst Medallion (11196, -4.68 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 49.2 healing_power points (16.71 DPS) | yes | Ironfeather Shoulders (15067, -0.87 DPS) [crafted]; Nethergeld Shoulders (254049, -3.02 DPS) [crafted]; Mender's Leather Shoulders (252538, -3.05 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.49 DPS) [quest]; Battle Healer's Cloak (19527, -0.82 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272076, -1.29 DPS) [vendor] |
| chest | Robes of Insight (940) | World drop [world_drop] | 65.4 healing_power points (22.20 DPS) | yes | Stormcloth Vest (10020, -3.33 DPS, sim-verified) [crafted]; Stone Guard's Inscribed Chestpiece (220838, -6.27 DPS) [vendor]; Feathered Breastplate (8349, -6.41 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 32.4 healing_power points (10.99 DPS) | yes | Mender's Leather Bracers (252543, -0.14 DPS) [crafted]; Mender's Mail Bracers (252573, -0.14 DPS) [crafted]; Aristocratic Cuffs (12546, -0.38 DPS) [dungeon] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.42 DPS) [crafted]; Greenleaf Handwraps (19116, -3.92 DPS) [quest] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 46.2 healing_power points (15.69 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.87 DPS) [dungeon]; Gilded Waistcord (254081, -1.71 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 56.9 healing_power points (19.34 DPS) | yes | Jinxed Hoodoo Kilt (9474, -2.06 DPS, sim-verified) [dungeon]; Stone Guard's Inscribed Legplates (220839, -3.41 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -4.47 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 43.1 healing_power points (14.65 DPS) | yes | Mender's Leather Boots (252472, -0.10 DPS) [crafted]; Mender's Mail Sabatons (252579, -0.10 DPS) [crafted]; Mender's Leather Shoes (252533, -1.91 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 27.4 healing_power points (9.31 DPS) | yes | Cyclopean Band (11824, -2.25 DPS) [dungeon]; Advisor's Ring (19519, -2.40 DPS) [rep]; Snake Hoop (6750, -2.68 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 21.2 healing_power points (7.19 DPS) | yes | Cyclopean Band (11824, -0.14 DPS) [dungeon]; Advisor's Ring (19519, -0.28 DPS) [rep]; Snake Hoop (6750, -0.56 DPS) [quest] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enamored Water Spirit (20503, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enamored Water Spirit (20503, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.53 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -7.41 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -20.18 DPS) [quest]; Dreamstaff (249454, -26.59 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soulcatcher Halo; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Robes of Insight; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Brainlash; finger2: Runemender's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 1018, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 366.8. Weights run: 11.8s. Verify run: 429.1s. 2307 eligible items had no known source.

Race sweep: orc 366.8 ± 0.2, windshaper-skyborne 366.2 ± 0.2, troll 364.3 ± 0.2, tauren 363.6 ± 0.2 (heal_score); the guide race orc is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.253, intellect=2.545 ± 0.065, spirit=1.281 ± 0.049, mp5=3.995 ± 0.060, crit=0.798 ± 0.047 per rating point (14 rating = 1%, 11.168 per %), spell_haste=not significant (3.007 ± 1.267)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 136.7 healing_power points (23.16 DPS) | yes | Devout Crown (16693, -6.17 DPS) [dungeon]; Warlord's Mail Skullcap (231669, -6.49 DPS) [vendor]; Black Dragonscale Helm (252605, -10.67 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 76.6 healing_power points (12.97 DPS) | yes | Lady Maye's Pendant (14558, -3.07 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -4.91 DPS) [dungeon]; Animated Chain Necklace (18723, -5.87 DPS) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 84.7 healing_power points (14.34 DPS) | yes | Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Darkspear Shoulders (272104, -0.68 DPS) [vendor]; Darkspear Shoulderpads (272103, -1.00 DPS, sim-verified) [vendor] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Shroud of the Exile (15421, -0.10 DPS) [quest]; Royal Tribunal Cloak (13376, -0.21 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -2.06 DPS, sim-verified) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 101.2 healing_power points (17.14 DPS) | yes | Warlord's Mail Chestguard (231668, +0.00 DPS) [vendor]; Devout Robe (16690, -0.56 DPS, sim-verified) [dungeon]; Robes of Insight (940, -1.21 DPS) [world_drop] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 58.8 healing_power points (9.96 DPS) | yes | Funeral Cuffs (12626, -0.88 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -1.26 DPS) [rep]; Bleak Howler Armguards (13208, -1.61 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Gloves (231666, -0.57 DPS) [vendor]; Hands of the Exalted Herald (12554, -2.24 DPS) [dungeon]; General's Mail Gauntlets (231660, -3.79 DPS) [pvp] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 78.2 healing_power points (13.23 DPS) | yes | General's Mail Waistband (16575, -0.29 DPS) [pvp]; Devout Belt (16696, -0.80 DPS) [dungeon]; Stormseeker's Girdle (272399, -0.96 DPS, sim-verified) [vendor] |
| legs | Red Dragonscale Leggings (252603) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; General's Mail Pants (231670, +0.00 DPS) [vendor]; General's Mail Leggings (231664, -0.20 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 76.3 healing_power points (12.92 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Faith Healer's Boots (22247, -0.36 DPS, sim-verified) [dungeon]; Firemoss Boots (22275, -1.77 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.89 DPS) [dungeon]; Rosewine Circle (13178, -1.44 DPS) [dungeon]; Naglering (11669, -11.88 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.73 DPS) [dungeon]; Channeler's Ring (272406, -1.08 DPS) [vendor]; Seal of Rivendare (13345, -2.52 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -1.35 DPS) [quest]; Serenity Field (272439, -1.86 DPS) [vendor]; Darkspear Voodoo Seal (272061, -2.03 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Serenity Field (272439, -0.72 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hammer of the Grand Crusader (18717, -2.10 DPS) [dungeon]; Argent Crusader (13249, -2.88 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Loomguard Armbraces; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Red Dragonscale Leggings; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Burst of Knowledge

No-known-source sample (15 of 2307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 869.1. Weights run: 7.5s. Verify run: 201.1s. 2307 eligible items had no known source.

Race sweep: windshaper-skyborne 870.9 ± 0.6, orc 868.8 ± 0.6, troll 866.8 ± 0.6, tauren 863.4 ± 0.6 (heal_score); windshaper-skyborne beats the guide race orc by 2.1 (0.24%) ± 0.8, beyond error.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.441, intellect=1.148 ± 0.057, spirit=0.464 ± 0.045, mp5=1.470 ± 0.058, crit=0.656 ± 0.060 per rating point (14 rating = 1%, 9.181 per %), spell_haste=5.412 ± 1.346

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 100.2 healing_power points (64.12 DPS) | yes | Warlord's Mail Skullcap (231669, -17.58 DPS) [vendor]; Champion's Mail Skullcap (227168, -24.76 DPS) [vendor]; Mooncloth Circlet (14140, -26.09 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 43.3 healing_power points (27.73 DPS) | yes | Animated Chain Necklace (18723, -5.24 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -6.72 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -9.04 DPS) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 47.0 healing_power points (30.09 DPS) | yes | Champion's Mail Epaulets (227166, +0.00 DPS) [vendor]; Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Mantle of Lost Hope (22234, -1.17 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -4.58 DPS, sim-verified) [rep]; Shroud of the Exile (15421, -6.49 DPS) [quest]; Arcanoweave Cloak (272411, -7.48 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.7 healing_power points (39.50 DPS) | yes | Legionnaire's Mail Chestguard (227171, +0.00 DPS) [vendor]; Warlord's Mail Chestguard (231668, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.59 DPS) [dungeon] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 41.4 healing_power points (26.53 DPS) | yes | Dryad's Wrist Bindings (19595, -3.87 DPS, sim-verified) [rep]; Funeral Cuffs (12626, -4.72 DPS) [dungeon]; Bleak Howler Armguards (13208, -6.15 DPS) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -3.49 DPS, sim-verified) [dungeon]; Blood Guard's Mail Gloves (227169, -4.15 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Mender's Leather Waistguard (252477, -4.79 DPS) [crafted]; Mender's Mail Belt (252591, -4.79 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 67.9 healing_power points (43.43 DPS) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, +0.00 DPS) [vendor]; Ghoul Skin Leggings (18682, -6.39 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 46.1 healing_power points (29.52 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Faith Healer's Boots (22247, -0.49 DPS) [dungeon]; Swarmtender's Footpads (275609, -1.99 DPS) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -1.32 DPS) [dungeon]; Rosewine Circle (13178, -7.09 DPS) [dungeon]; Naglering (11669, -25.15 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Rosewine Circle (13178, -1.21 DPS) [dungeon]; Maiden's Circle (13001, -3.77 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.92 DPS) [world_drop]; Weakness Analyzer (272438, -1.92 DPS) [vendor]; Second Wind (11819, -3.31 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -7.82 DPS, sim-verified) [quest]; Hammer of the Grand Crusader (18717, -9.43 DPS) [dungeon]; Staff of Metanoia (22394, -16.03 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderguards; back: Drape of Recovery; chest: Stormcloth Vest; wrist: Loomguard Armbraces; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Mooncloth Boots; finger1: Fordring's Seal; finger2: Blessed Band of Light; trinket2: Serenity Field

No-known-source sample (15 of 2307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

