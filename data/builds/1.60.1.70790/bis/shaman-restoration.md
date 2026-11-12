# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-2530000010000000)

Set DPS (verified): 73.1. Weights run: 5.2s. Verify run: 26.6s. 286 eligible items had no known source.

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

Set DPS (verified): 134.7. Weights run: 5.5s. Verify run: 32.7s. 475 eligible items had no known source.

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

Set DPS (verified): 210.4. Weights run: 10.1s. Verify run: 60.2s. 804 eligible items had no known source.

Race sweep: dwarf 210.4 ± 0.1 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.211, intellect=1.550 ± 0.021, spirit=0.930 ± 0.018, mp5=2.956 ± 0.022, crit=0.245 ± 0.018 per rating point (14 rating = 1%, 3.428 per %), spell_haste=not significant (0.322 ± 0.226)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.7 healing_power points (17.99 DPS) | yes | Whitemane's Chapeau (7720, -3.16 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -5.75 DPS) [world_drop]; Miner's Hat of the Deep (9429, -5.83 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.4 healing_power points (10.28 DPS) | yes | Necklace of Calisea (1714, -4.31 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -5.67 DPS) [dungeon]; Seal of the Expedition (271919, -6.84 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 42.1 healing_power points (13.37 DPS) | yes | Sheepshear Mantle (13115, -3.19 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -5.30 DPS) [crafted]; Green Silken Shoulders (7057, -5.42 DPS) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.60 DPS) [rep]; Prelacy Cape (7004, -1.93 DPS) [quest]; Blackforge Cape (6424, -2.07 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.4 healing_power points (20.76 DPS) | yes | Deathchill Armor (10764, -4.91 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -6.49 DPS) [crafted]; Death's Head Vestment (10581, -7.28 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.9 healing_power points (5.67 DPS) | yes | Mindthrust Bracers (1974, -0.25 DPS, sim-verified) [dungeon]; Enchanted Kodo Bracers (13119, -0.75 DPS) [world_drop]; Guardian Leather Bracers (4260, -0.82 DPS) [crafted] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.68 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.98 DPS) [crafted]; Dreamweave Gloves (10019, -1.34 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 40.0 healing_power points (12.68 DPS) | yes | Mender's Leather Belt (252523, -1.77 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -2.26 DPS) [crafted]; Deathmage Sash (10771, -3.09 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 46.3 healing_power points (14.69 DPS) | yes | Stoneweaver Leggings (9407, -0.59 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -2.53 DPS) [crafted]; Pristine Leggings (253987, -2.80 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 41.4 healing_power points (13.12 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -2.15 DPS) [crafted]; Skycaller's Leather Shoes (252532, -4.08 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 24.8 healing_power points (7.87 DPS) | yes | Lorekeeper's Ring (19524, -2.20 DPS) [rep]; Philanthropist's Ring (281635, -2.24 DPS) [quest]; Electrocutioner Lagnut (9447, -2.95 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 23.4 healing_power points (7.41 DPS) | yes | Lorekeeper's Ring (19524, -1.51 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -1.78 DPS) [quest]; Electrocutioner Lagnut (9447, -2.49 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -10.10 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -10.61 DPS) [dungeon]; Staff of the Friar (3415, -10.67 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beacon of Hope (9393, -2.54 DPS, sim-verified) [dungeon]; Battered Viking Shield (9403, -3.55 DPS) [dungeon]; Orb of Souls (249395, -4.73 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 0000000000000000-055000000000000000-5530001011513051)

Set DPS (verified): 281.3. Weights run: 11.0s. Verify run: 142.6s. 1050 eligible items had no known source.

Race sweep: dwarf 281.4 ± 0.1 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.175, intellect=1.589 ± 0.027, spirit=0.709 ± 0.020, mp5=2.418 ± 0.026, crit=0.399 ± 0.022 per rating point (14 rating = 1%, 5.586 per %), spell_haste=2.467 ± 0.503

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcatcher Halo (10630) | Sunken Temple: Atal'ai Warrior [dungeon] | 57.7 healing_power points (16.67 DPS) | yes | Papal Fez (9431, -0.77 DPS, sim-verified) [dungeon]; Gemburst Circlet (10751, -1.43 DPS) [quest]; Soothsayer's Headdress (17740, -1.51 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.7 healing_power points (9.44 DPS) | yes | Gemshard Heart (17707, -0.90 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.20 DPS) [world_drop]; Mindburst Medallion (11196, -3.94 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 53.2 healing_power points (15.35 DPS) | yes | Ironfeather Shoulders (15067, -0.20 DPS) [crafted]; Lead Surveyor's Mantle (11842, -2.75 DPS) [dungeon]; Mender's Leather Shoulders (252538, -2.86 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -0.58 DPS) [rep]; Caretaker's Cape (19531, -0.82 DPS, sim-verified) [rep]; Runecloth Cloak (13860, -1.75 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 71.1 healing_power points (20.54 DPS) | yes | Embrace of the Wind Serpent (12462, -0.48 DPS, sim-verified) [world]; Stormcloth Vest (10020, -2.16 DPS) [crafted]; Feathered Breastplate (8349, -5.55 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 34.8 healing_power points (10.06 DPS) | yes | Mender's Leather Bracers (252543, -0.13 DPS) [crafted]; Mender's Mail Bracers (252573, -0.13 DPS) [crafted]; Nethergeld Cuffs (254061, -0.28 DPS, sim-verified) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.32 DPS) [crafted]; Silkweb Gloves (11634, -3.69 DPS) [dungeon] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 49.1 healing_power points (14.17 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.25 DPS) [dungeon]; Gilded Waistcord (254081, -1.66 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 60.9 healing_power points (17.58 DPS) | yes | Jinxed Hoodoo Kilt (9474, -2.14 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -3.04 DPS) [dungeon]; Warchief Kilt (7760, -5.03 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 45.5 healing_power points (13.13 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Leather Shoes (252533, -1.56 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 30.4 healing_power points (8.79 DPS) | yes | Cyclopean Band (11824, -2.16 DPS) [dungeon]; Snake Hoop (6750, -2.41 DPS) [quest]; Woodseed Hoop (17768, -2.47 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 23.5 healing_power points (6.78 DPS) | yes | Cyclopean Band (11824, -0.16 DPS) [dungeon]; Snake Hoop (6750, -0.41 DPS) [quest]; Woodseed Hoop (17768, -0.46 DPS) [quest] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, +0.00 DPS) [rep]; Uther's Strength (11302, -0.34 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -18.52 DPS) [dungeon]; Resurgence Rod (17743, -18.96 DPS, sim-verified) [quest]; Dreamstaff (249454, -24.25 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soulcatcher Halo; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Runemender's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 1050, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 373.9. Weights run: 12.8s. Verify run: 299.5s. 2374 eligible items had no known source.

Race sweep: dwarf 373.9 ± 0.2 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.157, intellect=1.859 ± 0.046, spirit=0.918 ± 0.033, mp5=2.902 ± 0.043, crit=0.552 ± 0.030 per rating point (14 rating = 1%, 7.734 per %), spell_haste=not significant (1.604 ± 0.760)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 119.3 healing_power points (29.61 DPS) | yes | Devout Crown (16693, -10.07 DPS, sim-verified) [dungeon]; Black Dragonscale Helm (252605, -10.16 DPS) [crafted]; The Postmaster's Band (13390, -12.08 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 60.6 healing_power points (15.05 DPS) | yes | Lady Maye's Pendant (14558, -3.73 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -4.09 DPS) [dungeon]; Animated Chain Necklace (18723, -5.26 DPS) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 66.2 healing_power points (16.43 DPS) | yes | Theramore Shoulders (271954, -0.99 DPS) [rep]; Theramore Shoulderpads (271953, -1.17 DPS, sim-verified) [rep]; Ironfeather Shoulders (15067, -1.65 DPS) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.39 DPS) [quest]; Gracious Cape (18743, -1.64 DPS, sim-verified) [dungeon]; Expeditionary's Cape (271907, -1.83 DPS) [rep] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 80.6 healing_power points (20.01 DPS) | yes | Devout Robe (16690, -0.56 DPS) [dungeon]; Robes of the Exalted (13346, -0.69 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -0.86 DPS, sim-verified) [world] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 50.7 healing_power points (12.59 DPS) | yes | Funeral Cuffs (12626, -1.66 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -1.84 DPS) [rep]; Bleak Howler Armguards (13208, -1.95 DPS, sim-verified) [dungeon] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -2.67 DPS, sim-verified) [dungeon]; Mender's Leather Gauntlets (252551, -4.52 DPS) [crafted]; Mender's Mail Gauntlets (252587, -4.52 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 65.3 healing_power points (16.20 DPS) | yes | Sash of Mercy (14553, -1.02 DPS, sim-verified) [world_drop]; Stormseeker's Girdle (272399, -1.16 DPS) [vendor]; Devout Belt (16696, -2.20 DPS) [dungeon] |
| legs | Legplates of the Chromatic Defier (12945) | Legplates of the Chromatic Defier [quest] | 91.2 healing_power points (22.63 DPS) | yes | Leggings of Arcana (12756, -2.16 DPS, sim-verified) [quest]; Red Dragonscale Leggings (252603, -2.90 DPS) [crafted]; Ironfeather Leggings (252486, -2.98 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 62.0 healing_power points (15.38 DPS) | yes | Faith Healer's Boots (22247, -1.35 DPS, sim-verified) [dungeon]; Firemoss Boots (22275, -2.11 DPS) [dungeon]; Mender's Leather Boots (252472, -3.35 DPS) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -2.01 DPS) [dungeon]; Seal of Rivendare (13345, -2.46 DPS) [dungeon]; Naglering (11669, -12.00 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.99 DPS, sim-verified) [dungeon]; Seal of Rivendare (13345, -2.12 DPS) [dungeon]; Blessed Band of Light (272407, -2.43 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -0.61 DPS, sim-verified) [vendor]; Uther's Strength (11302, -1.29 DPS) [world_drop]; Weakness Analyzer (272438, -1.29 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Argent Crusader (13249, -3.17 DPS, sim-verified) [quest]; Hammer of the Grand Crusader (18717, -3.28 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Tidal Totem (272431) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Sustaining (23200, -6.16 DPS, sim-verified) [dungeon] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Loomguard Armbraces; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Legplates of the Chromatic Defier; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Ankh of Life; trinket2: Burst of Knowledge; ranged: Tidal Totem

No-known-source sample (15 of 2374, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 889.7. Weights run: 8.2s. Verify run: 285.7s. 2374 eligible items had no known source.

Race sweep: dwarf 890.0 ± 0.6 (heal_score); the guide race dwarf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.145, intellect=1.254 ± 0.050, spirit=0.668 ± 0.040, mp5=2.205 ± 0.045, crit=0.671 ± 0.034 per rating point (14 rating = 1%, 9.390 per %), spell_haste=4.112 ± 0.994

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 104.9 healing_power points (26.03 DPS) | yes | Devout Crown (16693, -11.12 DPS) [dungeon]; Black Dragonscale Helm (252605, -11.88 DPS) [crafted]; Mooncloth Circlet (14140, -23.35 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 47.3 healing_power points (11.72 DPS) | yes | Animated Chain Necklace (18723, -1.22 DPS, sim-verified) [dungeon]; Tooth of Gnarr (13141, -2.94 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.08 DPS) [world_drop] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 49.9 healing_power points (12.37 DPS) | yes | Mantle of Lost Hope (22234, -0.31 DPS) [dungeon]; Theramore Shoulderpads (271953, -0.99 DPS) [rep]; Theramore Shoulders (271954, -0.99 DPS) [rep] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.63 DPS) [quest]; Gracious Cape (18743, -2.68 DPS) [dungeon]; Caretaker's Cape (19530, -3.58 DPS, sim-verified) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Vest (14138, -0.45 DPS) [crafted]; Stormcloth Vest (10020, -0.56 DPS) [crafted]; Robes of the Exalted (13346, -1.16 DPS, sim-verified) [dungeon] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 44.4 healing_power points (11.01 DPS) | yes | Bleak Howler Armguards (13208, -2.17 DPS) [dungeon]; Funeral Cuffs (12626, -2.19 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -2.69 DPS, sim-verified) [rep] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -0.63 DPS) [dungeon]; Mender's Leather Gauntlets (252551, -2.37 DPS) [crafted]; Mender's Mail Gauntlets (252587, -2.37 DPS) [crafted] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 61.0 healing_power points (15.14 DPS) | yes | Wisdom of the Timbermaw (19047, -3.39 DPS, sim-verified) [crafted]; Stormseeker's Girdle (272399, -3.70 DPS) [vendor]; Mender's Leather Waistguard (252477, -3.96 DPS) [crafted] |
| legs | Legplates of the Chromatic Defier (12945) | Legplates of the Chromatic Defier [quest] | 76.8 healing_power points (19.04 DPS) | yes | Ironfeather Leggings (252486, -0.98 DPS, sim-verified) [crafted]; Sentinel's Silk Leggings (237815, -3.33 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -3.33 DPS) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.2 healing_power points (12.46 DPS) | yes | Swarmtender's Footpads (275609, -1.80 DPS) [crafted]; Broodwatcher's Treaders (275621, -1.80 DPS) [crafted]; Faith Healer's Boots (22247, -2.01 DPS, sim-verified) [dungeon] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.18 DPS) [dungeon]; Blessed Band of Light (272407, -2.03 DPS) [vendor]; Naglering (11669, -25.28 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -0.22 DPS) [vendor]; Maiden's Circle (13001, -1.60 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+11.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.74 DPS) [world_drop]; Weakness Analyzer (272438, -0.74 DPS) [vendor]; Burst of Knowledge (11832, -0.83 DPS) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -3.33 DPS) [dungeon]; Argent Crusader (13249, -5.77 DPS, sim-verified) [quest]; Staff of Metanoia (22394, -5.92 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Tidal Totem (272431) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Sustaining (23200, -10.38 DPS, sim-verified) [dungeon] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Embrace of the Wind Serpent; wrist: Loomguard Armbraces; hands: Theramore Handwraps; waist: Sash of Mercy; legs: Legplates of the Chromatic Defier; feet: Mooncloth Boots; finger1: Fordring's Seal; finger2: Rosewine Circle; trinket2: Serenity Field; ranged: Tidal Totem

No-known-source sample (15 of 2374, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-000000000000000000-2530000010000000)

Set DPS (verified): 72.8. Weights run: 5.2s. Verify run: 27.6s. 275 eligible items had no known source.

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

Set DPS (verified): 135.1. Weights run: 5.5s. Verify run: 34.0s. 469 eligible items had no known source.

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

Set DPS (verified): 210.8. Weights run: 10.1s. Verify run: 62.7s. 787 eligible items had no known source.

Race sweep: orc 210.8 ± 0.1, windshaper-skyborne 210.7 ± 0.1, troll 208.6 ± 0.1, tauren 208.4 ± 0.1 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.211, intellect=1.550 ± 0.021, spirit=0.930 ± 0.018, mp5=2.956 ± 0.022, crit=0.245 ± 0.018 per rating point (14 rating = 1%, 3.428 per %), spell_haste=not significant (0.322 ± 0.226)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.7 healing_power points (17.99 DPS) | yes | Whitemane's Chapeau (7720, -3.22 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -5.75 DPS) [world_drop]; Miner's Hat of the Deep (9429, -5.83 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.4 healing_power points (10.28 DPS) | yes | Necklace of Calisea (1714, -4.35 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -5.67 DPS) [dungeon]; Darkspear Warding Pendant (272074, -6.84 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 42.1 healing_power points (13.37 DPS) | yes | Sheepshear Mantle (13115, -3.43 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -5.30 DPS) [crafted]; Green Silken Shoulders (7057, -5.42 DPS) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.60 DPS) [vendor]; Blackforge Cape (6424, -2.07 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.4 healing_power points (20.76 DPS) | yes | Deathchill Armor (10764, -5.35 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -6.49 DPS) [crafted]; Death's Head Vestment (10581, -7.28 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.9 healing_power points (5.67 DPS) | yes | Mindthrust Bracers (1974, -0.40 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.47 DPS) [quest]; Enchanted Kodo Bracers (13119, -0.75 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.59 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.98 DPS) [crafted]; Dreamweave Gloves (10019, -1.34 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 40.0 healing_power points (12.68 DPS) | yes | Mender's Leather Belt (252523, -2.13 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -2.26 DPS) [crafted]; Deathmage Sash (10771, -3.09 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 46.3 healing_power points (14.69 DPS) | yes | Stoneweaver Leggings (9407, -0.45 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -2.53 DPS) [crafted]; Pristine Leggings (253987, -2.80 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 41.4 healing_power points (13.12 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -2.15 DPS) [crafted]; Skycaller's Leather Shoes (252532, -4.08 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 24.8 healing_power points (7.87 DPS) | yes | Advisor's Ring (19520, -2.20 DPS) [rep]; Philanthropist's Ring (281635, -2.24 DPS) [quest]; Darkspear Signet (272070, -2.24 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 23.4 healing_power points (7.41 DPS) | yes | Philanthropist's Ring (281635, -1.78 DPS) [quest]; Darkspear Signet (272070, -1.78 DPS) [vendor]; Advisor's Ring (19520, -1.79 DPS, sim-verified) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -9.56 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -10.61 DPS) [dungeon]; Staff of the Friar (3415, -10.67 DPS) [dungeon] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forcestone Buckler (17508, -1.57 DPS) [quest]; Beacon of Hope (9393, -1.74 DPS, sim-verified) [dungeon]; Battered Viking Shield (9403, -2.57 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Prophetic Cane

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 0000000000000000-055000000000000000-5530001011513051)

Set DPS (verified): 282.5. Weights run: 11.0s. Verify run: 150.5s. 1018 eligible items had no known source.

Race sweep: windshaper-skyborne 282.7 ± 0.1, orc 282.5 ± 0.1, troll 280.2 ± 0.1, tauren 280.0 ± 0.1 (heal_score); windshaper-skyborne beats the guide race orc by 0.1 (0.05%) ± 0.2, inside error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.175, intellect=1.589 ± 0.027, spirit=0.709 ± 0.020, mp5=2.418 ± 0.026, crit=0.399 ± 0.022 per rating point (14 rating = 1%, 5.586 per %), spell_haste=2.467 ± 0.503

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcatcher Halo (10630) | Sunken Temple: Atal'ai Warrior [dungeon] | 57.7 healing_power points (16.67 DPS) | yes | Papal Fez (9431, -0.89 DPS, sim-verified) [dungeon]; Gemburst Circlet (10751, -1.43 DPS) [quest]; Soothsayer's Headdress (17740, -1.51 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.7 healing_power points (9.44 DPS) | yes | Gemshard Heart (17707, -0.99 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.20 DPS) [world_drop]; Mindburst Medallion (11196, -3.94 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 53.2 healing_power points (15.35 DPS) | yes | Ironfeather Shoulders (15067, -0.20 DPS) [crafted]; Lead Surveyor's Mantle (11842, -2.75 DPS) [dungeon]; Mender's Leather Shoulders (252538, -2.86 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.43 DPS) [quest]; Battle Healer's Cloak (19527, -0.44 DPS) [rep]; Darkspear Raider's Cloak (272076, -0.58 DPS) [vendor] |
| chest | Robes of Insight (940) | World drop [world_drop] | 71.1 healing_power points (20.54 DPS) | yes | Embrace of the Wind Serpent (12462, -0.57 DPS, sim-verified) [world]; Stormcloth Vest (10020, -2.16 DPS) [crafted]; Feathered Breastplate (8349, -5.55 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 34.8 healing_power points (10.06 DPS) | yes | Mender's Leather Bracers (252543, -0.13 DPS) [crafted]; Mender's Mail Bracers (252573, -0.13 DPS) [crafted]; Nethergeld Cuffs (254061, -0.39 DPS, sim-verified) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.32 DPS) [crafted]; Greenleaf Handwraps (19116, -3.52 DPS) [quest] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 49.1 healing_power points (14.17 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.25 DPS) [dungeon]; Gilded Waistcord (254081, -1.66 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 60.9 healing_power points (17.58 DPS) | yes | Jinxed Hoodoo Kilt (9474, -2.09 DPS, sim-verified) [dungeon]; Stone Guard's Inscribed Legplates (220839, -2.95 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -3.04 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 45.5 healing_power points (13.13 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Leather Shoes (252533, -1.56 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 30.4 healing_power points (8.79 DPS) | yes | Cyclopean Band (11824, -2.16 DPS) [dungeon]; Snake Hoop (6750, -2.41 DPS) [quest]; Woodseed Hoop (17768, -2.47 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 23.5 healing_power points (6.78 DPS) | yes | Cyclopean Band (11824, -0.16 DPS) [dungeon]; Snake Hoop (6750, -0.41 DPS) [quest]; Woodseed Hoop (17768, -0.46 DPS) [quest] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enamored Water Spirit (20503, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enamored Water Spirit (20503, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.47 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -18.52 DPS) [dungeon]; Resurgence Rod (17743, -18.83 DPS, sim-verified) [quest]; Dreamstaff (249454, -24.25 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soulcatcher Halo; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Runemender's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 1018, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 375.9. Weights run: 12.8s. Verify run: 304.2s. 2331 eligible items had no known source.

Race sweep: orc 376.0 ± 0.2, windshaper-skyborne 375.6 ± 0.2, troll 373.7 ± 0.2, tauren 372.5 ± 0.2 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.157, intellect=1.859 ± 0.046, spirit=0.918 ± 0.033, mp5=2.902 ± 0.043, crit=0.552 ± 0.030 per rating point (14 rating = 1%, 7.734 per %), spell_haste=not significant (1.604 ± 0.760)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 119.3 healing_power points (29.61 DPS) | yes | Warlord's Mail Skullcap (231669, -8.93 DPS) [vendor]; Devout Crown (16693, -10.00 DPS, sim-verified) [dungeon]; Black Dragonscale Helm (252605, -10.16 DPS) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 60.6 healing_power points (15.05 DPS) | yes | Lady Maye's Pendant (14558, -3.01 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -4.09 DPS) [dungeon]; Animated Chain Necklace (18723, -5.26 DPS) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 66.2 healing_power points (16.43 DPS) | yes | Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.77 DPS, sim-verified) [vendor]; Darkspear Shoulders (272104, -0.99 DPS) [vendor] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-verified (376.0 DPS) | yes | Gracious Cape (18743, -1.32 DPS, sim-verified) [dungeon]; Shroud of the Exile (15421, -1.39 DPS) [quest]; Darkspear Raider's Cloak (272063, -1.83 DPS) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 80.6 healing_power points (20.01 DPS) | yes | Warlord's Mail Chestguard (231668, +0.00 DPS) [vendor]; Devout Robe (16690, -0.56 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -0.97 DPS, sim-verified) [world] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 50.7 healing_power points (12.59 DPS) | yes | Bleak Howler Armguards (13208, -1.28 DPS, sim-verified) [dungeon]; Funeral Cuffs (12626, -1.66 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -1.84 DPS) [rep] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (376.0 DPS) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -2.57 DPS, sim-verified) [dungeon]; General's Mail Gauntlets (231660, -4.36 DPS) [pvp] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 65.3 healing_power points (16.20 DPS) | yes | Sash of Mercy (14553, -0.90 DPS, sim-verified) [world_drop]; Stormseeker's Girdle (272399, -1.16 DPS) [vendor]; Devout Belt (16696, -2.20 DPS) [dungeon] |
| legs | Legplates of the Chromatic Defier (12945) | Legplates of the Chromatic Defier [quest] | 91.2 healing_power points (22.63 DPS) | yes | General's Mail Pants (231670, +0.00 DPS) [vendor]; Leggings of Arcana (12756, -2.11 DPS, sim-verified) [quest]; General's Mail Leggings (231664, -2.69 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 62.0 healing_power points (15.38 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Faith Healer's Boots (22247, -1.23 DPS, sim-verified) [dungeon]; Firemoss Boots (22275, -2.11 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (376.0 DPS) | yes | Rosewine Circle (13178, -2.01 DPS) [dungeon]; Seal of Rivendare (13345, -2.46 DPS) [dungeon]; Naglering (11669, -12.08 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (376.0 DPS) | yes | Rosewine Circle (13178, -1.55 DPS, sim-verified) [dungeon]; Seal of Rivendare (13345, -2.12 DPS) [dungeon]; Blessed Band of Light (272407, -2.43 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (376.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (376.0 DPS) | yes | Serenity Field (272439, -0.55 DPS) [vendor]; Weakness Analyzer (272438, -1.29 DPS) [vendor]; Shard of the Splithooves (10659, -1.34 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (376.0 DPS) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Argent Crusader (13249, -2.80 DPS, sim-verified) [quest]; Hammer of the Grand Crusader (18717, -3.28 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Tidal Totem (272431) | Pix Xizzix [vendor] | sim-verified (376.0 DPS) | yes | Totem of Sustaining (23200, -6.43 DPS, sim-verified) [dungeon] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Loomguard Armbraces; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Legplates of the Chromatic Defier; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Ankh of Life; trinket2: Burst of Knowledge; ranged: Tidal Totem

No-known-source sample (15 of 2331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 895.8. Weights run: 8.2s. Verify run: 283.6s. 2331 eligible items had no known source.

Race sweep: windshaper-skyborne 899.0 ± 0.6, orc 895.7 ± 0.6, troll 894.5 ± 0.5, tauren 890.2 ± 0.6 (heal_score); windshaper-skyborne beats the guide race orc by 3.3 (0.37%) ± 0.8, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.145, intellect=1.254 ± 0.050, spirit=0.668 ± 0.040, mp5=2.205 ± 0.045, crit=0.671 ± 0.034 per rating point (14 rating = 1%, 9.390 per %), spell_haste=4.112 ± 0.994

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 104.9 healing_power points (26.03 DPS) | yes | Warlord's Mail Skullcap (231669, -7.50 DPS) [vendor]; Champion's Mail Skullcap (227168, -9.94 DPS) [vendor]; Mooncloth Circlet (14140, -26.17 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 47.3 healing_power points (11.72 DPS) | yes | Tooth of Gnarr (13141, -2.94 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.08 DPS) [world_drop]; Animated Chain Necklace (18723, -3.10 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 49.9 healing_power points (12.37 DPS) | yes | Mantle of Lost Hope (22234, +0.00 DPS) [dungeon]; Champion's Mail Epaulets (227166, +0.00 DPS) [vendor]; Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.63 DPS) [quest]; Gracious Cape (18743, -2.68 DPS) [dungeon]; Battle Healer's Cloak (19526, -4.56 DPS, sim-verified) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | 66.7 healing_power points (16.56 DPS) | yes | Legionnaire's Mail Chestguard (227171, +0.00 DPS) [vendor]; Warlord's Mail Chestguard (231668, +0.00 DPS) [vendor]; Embrace of the Wind Serpent (12462, -0.29 DPS) [world] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 44.4 healing_power points (11.01 DPS) | yes | Bleak Howler Armguards (13208, -2.17 DPS) [dungeon]; Funeral Cuffs (12626, -2.19 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -3.13 DPS, sim-verified) [rep] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -1.02 DPS, sim-verified) [dungeon]; Blood Guard's Mail Gloves (227169, -2.03 DPS) [vendor] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 61.0 healing_power points (15.14 DPS) | yes | Wisdom of the Timbermaw (19047, -2.99 DPS, sim-verified) [crafted]; Stormseeker's Girdle (272399, -3.70 DPS) [vendor]; Mender's Leather Waistguard (252477, -3.96 DPS) [crafted] |
| legs | Legplates of the Chromatic Defier (12945) | Legplates of the Chromatic Defier [quest] | 76.8 healing_power points (19.04 DPS) | yes | General's Mail Pants (231670, +0.00 DPS) [vendor]; Legionnaire's Mail Pants (227167, -1.50 DPS) [vendor]; Ironfeather Leggings (252486, -1.99 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.2 healing_power points (12.46 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Faith Healer's Boots (22247, -0.79 DPS) [dungeon]; Blood Guard's Mail Boots (227170, -1.54 DPS) [vendor] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.18 DPS) [dungeon]; Blessed Band of Light (272407, -2.03 DPS) [vendor]; Naglering (11669, -25.23 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -0.22 DPS) [vendor]; Maiden's Circle (13001, -1.60 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.74 DPS) [world_drop]; Burst of Knowledge (11832, -0.83 DPS) [dungeon]; Weakness Analyzer (272438, -1.50 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -3.33 DPS) [dungeon]; Staff of Metanoia (22394, -5.92 DPS) [dungeon]; Argent Crusader (13249, -7.29 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Tidal Totem (272431) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Sustaining (23200, -11.97 DPS, sim-verified) [dungeon] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderguards; back: Drape of Recovery; chest: Robes of the Exalted; wrist: Loomguard Armbraces; hands: Raider Handwraps; waist: Sash of Mercy; legs: Legplates of the Chromatic Defier; feet: Mooncloth Boots; finger1: Fordring's Seal; finger2: Rosewine Circle; trinket2: Serenity Field; ranged: Tidal Totem

No-known-source sample (15 of 2331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

