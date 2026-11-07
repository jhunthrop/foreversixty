# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-2530000010000000)

Set DPS (verified): 73.1. Weights run: 4.1s. Verify run: 24.1s. 260 eligible items had no known source.

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

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (dwarf, 0000000000000000-000000000000000000-5530001010411000)

Set DPS (verified): 132.3. Weights run: 4.4s. Verify run: 28.9s. 448 eligible items had no known source.

Race sweep: dwarf 132.3 ± 0.0 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.019, intellect=0.689 ± 0.004, spirit=0.457 ± 0.004, mp5=1.405 ± 0.004, crit=0.079 ± 0.003 per rating point (14 rating = 1%, 1.112 per %), spell_haste=-0.611 ± 0.064

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 35.7 healing_power points (7.43 DPS) | yes | Wisdom's Leather Helm (252515, -3.07 DPS) [crafted]; Filigreed Pristine Circlet (253975, -3.68 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -3.69 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.6 healing_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.09 DPS) [world]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Seal of the Expedition (271920, -0.24 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.2 healing_power points (3.16 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Magician's Mantle (12998, -0.83 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.66 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.93 DPS) [quest]; Prelacy Cape (7004, -1.34 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 27.6 healing_power points (5.73 DPS) | yes | Pristine Gown (253961, -0.55 DPS, sim-verified) [crafted]; Wisdom's Leather Tunic (252511, -0.56 DPS) [crafted]; Beguiler Robes (7728, -1.38 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.2 healing_power points (2.75 DPS) | yes | Mindthrust Bracers (1974, -0.47 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.60 DPS) [world_drop]; Nightsky Wristbands (6407, -1.60 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.1 healing_power points (3.55 DPS) | yes | Pristine Gloves (253913, -0.83 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.90 DPS) [crafted]; Naga Battle Gloves (888, -0.96 DPS, sim-verified) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 25.9 healing_power points (5.38 DPS) | yes | Skycaller's Leather Belt (252522, -2.42 DPS, sim-verified) [crafted]; Pristine Sash (253925, -2.52 DPS) [crafted]; Belt of Arugal (6392, -2.60 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | sim-verified (132.3 DPS) | yes | Pristine Leggings (253987, -0.13 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.64 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.05 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 26.7 healing_power points (5.54 DPS) | yes | Acidic Walkers (9454, -2.98 DPS) [dungeon]; Black Whelp Slippers (252424, -3.10 DPS) [crafted]; Wisdom's Leather Boots (252444, -3.50 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 healing_power points (2.54 DPS) | yes | Lorekeeper's Ring (19525, -0.50 DPS) [rep]; Snake Hoop (6750, -0.57 DPS) [quest]; Runemender's Seal (273031, -3.30 DPS, sim-verified) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Lorekeeper's Ring (19525, -0.08 DPS) [rep]; Runemender's Seal (273031, -0.10 DPS, sim-verified) [world]; Snake Hoop (6750, -0.16 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -2.19 DPS) [dungeon]; Staff of Westfall (2042, -2.87 DPS) [quest]; Staff of the Friar (3415, -4.97 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.7 healing_power points (3.27 DPS) | yes | Heart of Agamaggan (6694, -0.53 DPS, sim-verified) [dungeon]; Eye of Paleth (2943, -0.57 DPS) [quest]; Orb of Mystic Insight (249394, -0.96 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 0000000000000000-000000000000000000-5530001011513051)

Set DPS (verified): 203.4. Weights run: 8.1s. Verify run: 78.1s. 696 eligible items had no known source.

Race sweep: dwarf 203.4 ± 0.1 (heal_score); the guide race dwarf is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.322, intellect=1.989 ± 0.026, spirit=1.213 ± 0.024, mp5=3.950 ± 0.026, crit=0.343 ± 0.025 per rating point (14 rating = 1%, 4.803 per %), spell_haste=not significant (0.889 ± 0.340)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 66.7 healing_power points (13.11 DPS) | yes | Whitemane's Chapeau (7720, -2.52 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -4.09 DPS) [dungeon]; Corpseshroud (10574, -4.26 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 35.9 healing_power points (7.06 DPS) | yes | Prodigious Shadowshard Pendant (17773, -3.15 DPS) [quest]; Triune Amulet (7722, -3.82 DPS) [dungeon]; Necklace of Calisea (1714, -4.02 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 47.9 healing_power points (9.41 DPS) | yes | Sheepshear Mantle (13115, -3.27 DPS, sim-verified) [world_drop]; Green Silken Shoulders (7057, -3.53 DPS) [crafted]; Mantle of Doan (7712, -3.78 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Expeditionary's Cape (271922, -0.58 DPS, sim-verified) [rep]; Blackforge Cape (6424, -0.65 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.7 healing_power points (13.31 DPS) | yes | Dreamweave Vest (10021, -2.91 DPS) [crafted]; Robe of Power (7054, -3.96 DPS) [crafted]; Deathchill Armor (10764, -5.31 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanted Kodo Bracers (13119, -0.15 DPS) [world_drop]; Mindthrust Bracers (1974, -0.17 DPS, sim-verified) [dungeon]; Guardian Leather Bracers (4260, -0.57 DPS) [crafted] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.62 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.78 DPS) [crafted]; Stormcloth Gloves (10011, -0.79 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 45.2 healing_power points (8.88 DPS) | yes | Deathmage Sash (10771, -1.64 DPS) [dungeon]; Mender's Leather Belt (252523, -1.68 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.82 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 52.9 healing_power points (10.40 DPS) | yes | Stoneweaver Leggings (9407, -0.69 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.87 DPS) [crafted]; Pristine Leggings (253987, -2.15 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 46.4 healing_power points (9.12 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.50 DPS) [crafted]; Furen's Boots (13100, -2.57 DPS) [world_drop] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 20.9 healing_power points (4.10 DPS) | yes | Welken Ring (5011, -0.48 DPS) [world_drop]; Ring of Power Regulation (273030, -0.59 DPS) [dungeon]; Voodoo Band (1996, -0.65 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 19.9 healing_power points (3.92 DPS) | yes | Ring of Power Regulation (273030, -0.41 DPS) [dungeon]; Voodoo Band (1996, -0.47 DPS) [world]; Welken Ring (5011, -1.29 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -6.13 DPS) [dungeon]; Illusionary Rod (7713, -6.78 DPS) [dungeon]; Staff of Jordan (873, -7.82 DPS, sim-verified) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beacon of Hope (9393, -0.53 DPS, sim-verified) [dungeon]; Battered Viking Shield (9403, -1.33 DPS) [dungeon]; Ravager's Shield (14777, -2.10 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 0000000000000000-055000000000000000-5530001011513051)

Set DPS (verified): 270.9. Weights run: 9.7s. Verify run: 124.6s. 893 eligible items had no known source.

Race sweep: dwarf 270.9 ± 0.1 (heal_score); the guide race dwarf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.179, intellect=1.594 ± 0.028, spirit=0.716 ± 0.021, mp5=2.519 ± 0.028, crit=0.384 ± 0.021 per rating point (14 rating = 1%, 5.378 per %), spell_haste=2.628 ± 0.533

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 55.5 healing_power points (15.24 DPS) | yes | Soothsayer's Headdress (17740, -0.81 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -1.28 DPS) [crafted]; Cassandra's Grace (13102, -1.40 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.8 healing_power points (8.99 DPS) | yes | Horizon Choker (13085, -1.93 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -3.43 DPS) [dungeon]; Mindburst Medallion (11196, -3.73 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 52.6 healing_power points (14.43 DPS) | yes | Mender's Mail Shoulders (252569, -2.54 DPS) [crafted]; Nethergeld Shoulders (254049, -2.62 DPS) [crafted]; Mender's Leather Shoulders (252538, -3.06 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -0.54 DPS) [rep]; Caretaker's Cape (19531, -0.80 DPS, sim-verified) [rep]; Runecloth Cloak (13860, -1.68 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.7 healing_power points (17.48 DPS) | yes | Feathered Breastplate (8349, -3.07 DPS, sim-verified) [crafted]; Acumen Robes (17775, -3.52 DPS) [quest]; Robes of Insight (940, -3.60 DPS) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 34.7 healing_power points (9.53 DPS) | yes | Mender's Mail Bracers (252573, -0.08 DPS) [crafted]; Mender's Leather Bracers (252543, -0.24 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -1.79 DPS) [dungeon] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.31 DPS) [crafted]; Mender's Leather Gloves (252530, -4.51 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 49.1 healing_power points (13.48 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Waistcord (254081, -1.59 DPS) [crafted]; Dawnspire Cord (12466, -1.88 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 49.0 healing_power points (13.45 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.40 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.96 DPS) [dungeon]; Warchief Kilt (7760, -1.50 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 45.5 healing_power points (12.49 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Leather Shoes (252533, -1.47 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, -0.81 DPS) [dungeon]; Eye of Adaegus (5266, -1.08 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 23.0 healing_power points (6.32 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, -1.07 DPS) [dungeon]; Eye of Adaegus (5266, -1.34 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, +0.00 DPS) [rep]; Uther's Strength (11302, -0.50 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -17.59 DPS) [dungeon]; Resurgence Rod (17743, -18.09 DPS, sim-verified) [quest]; Dreamstaff (249454, -23.07 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Ironfeather Shoulders; back: Spritecaster Cape; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Mender's Leather Boots; finger1: Lorekeeper's Ring; finger2: Cyclopean Band; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 893, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 356.5. Weights run: 11.2s. Verify run: 316.9s. 2086 eligible items had no known source.

Race sweep: dwarf 356.7 ± 0.2 (heal_score); the guide race dwarf is best.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.031, intellect=1.987 ± 0.030, spirit=1.037 ± 0.025, mp5=2.927 ± 0.027, crit=0.338 ± 0.013 per rating point (14 rating = 1%, 4.736 per %), spell_haste=not significant (0.213 ± 0.275)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 123.3 healing_power points (18.67 DPS) | yes | Crown of the Penitent (13216, -6.38 DPS) [quest]; Black Dragonscale Helm (252605, -6.52 DPS) [crafted]; Devout Crown (16693, -9.02 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 69.6 healing_power points (10.55 DPS) | yes | Theramore Shoulders (271954, -0.61 DPS) [rep]; Theramore Shoulderpads (271953, -0.85 DPS, sim-verified) [rep]; Ironfeather Shoulders (15067, -1.00 DPS) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -0.85 DPS) [quest]; Frostweaver Cape (12968, -0.97 DPS) [dungeon]; Expeditionary's Cape (271907, -1.15 DPS, sim-verified) [rep] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Devout Robe (16690, -0.33 DPS) [dungeon]; Alanna's Embrace (13314, -0.75 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Mender's Leather Bracers (252543, -0.91 DPS) [crafted] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -1.35 DPS, sim-verified) [dungeon]; Mender's Leather Gauntlets (252551, -2.94 DPS) [crafted]; Mender's Mail Gauntlets (252587, -2.94 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 68.2 healing_power points (10.32 DPS) | yes | Sash of Mercy (14553, -0.72 DPS) [world_drop]; Stormseeker's Girdle (272399, -1.07 DPS, sim-verified) [vendor]; Devout Belt (16696, -1.22 DPS) [dungeon] |
| legs | Red Dragonscale Leggings (252603) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Ironfeather Leggings (252486, -0.77 DPS) [crafted]; Devout Skirt (16694, -0.90 DPS) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 65.3 healing_power points (9.89 DPS) | yes | Faith Healer's Boots (22247, -0.90 DPS, sim-verified) [dungeon]; Mender's Leather Boots (252472, -2.34 DPS) [crafted]; Mender's Mail Sabatons (252579, -2.34 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.03 DPS) [vendor]; Uther's Strength (11302, -3.63 DPS) [world_drop]; Second Wind (11819, -4.55 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -8.25 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -11.02 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Urgency (279249) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -1.52 DPS, sim-verified) [vendor] |

**New at 60:** head: Living Crown; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Red Dragonscale Leggings; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Ankh of Life; trinket2: Serenity Field; ranged: Totem of Urgency

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 878.4. Weights run: 7.0s. Verify run: 286.7s. 2086 eligible items had no known source.

Race sweep: dwarf 878.2 ± 0.6 (heal_score); the guide race dwarf is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.145, intellect=1.254 ± 0.050, spirit=0.668 ± 0.040, mp5=2.205 ± 0.045, crit=0.671 ± 0.034 per rating point (14 rating = 1%, 9.390 per %), spell_haste=4.112 ± 0.994

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 104.9 healing_power points (26.03 DPS) | yes | Devout Crown (16693, -11.12 DPS) [dungeon]; Crown of the Penitent (13216, -11.88 DPS) [quest]; Mooncloth Circlet (14140, -25.20 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.19 DPS) [world_drop] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 49.9 healing_power points (12.37 DPS) | yes | Theramore Shoulderpads (271953, -0.99 DPS) [rep]; Theramore Shoulders (271954, -0.99 DPS) [rep]; Royal Cap Spaulders (14548, -4.19 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Arcanoweave Cloak (272411, -3.42 DPS) [vendor]; Caretaker's Cape (19531, -3.43 DPS) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | 81.6 healing_power points (20.25 DPS) | yes | Stormcloth Vest (10020, -4.53 DPS) [crafted]; Devout Robe (16690, -5.33 DPS) [dungeon]; Mooncloth Vest (14138, -10.07 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.58 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.14 DPS) [crafted] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Leather Gauntlets (252551, -2.37 DPS) [crafted]; Mender's Mail Gauntlets (252587, -2.37 DPS) [crafted]; Hands of the Exalted Herald (12554, -2.90 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Stormseeker's Girdle (272399, -2.10 DPS) [vendor]; Mender's Leather Waistguard (252477, -2.36 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 70.0 healing_power points (17.36 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.65 DPS) [vendor]; Leggings of Arcana (12756, -1.70 DPS) [quest]; Sentinel's Silk Leggings (237815, -5.23 DPS, sim-verified) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.2 healing_power points (12.46 DPS) | yes | Faith Healer's Boots (22247, -0.29 DPS) [dungeon]; Swarmtender's Footpads (275609, -1.80 DPS) [crafted]; Broodwatcher's Treaders (275621, -1.80 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Band of Mending (22334, +0.00 DPS) [dungeon]; Maiden's Circle (13001, -1.38 DPS) [world_drop] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Maiden's Circle (13001, -3.10 DPS) [world_drop]; Rosewine Circle (13178, -3.95 DPS, sim-verified) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -4.96 DPS) [vendor]; Uther's Strength (11302, -5.95 DPS) [world_drop]; Second Wind (11819, -12.11 DPS, sim-verified) [dungeon] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, -1.45 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -15.82 DPS) [pvp]; Ironbark Staff (20069, -16.57 DPS) [rep]; Resurgence Rod (17743, -34.18 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Urgency (279249) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -2.70 DPS, sim-verified) [vendor] |

**New at 60:** head: Living Crown; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Robes of the Exalted; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Fordring's Seal; trinket1: Serenity Field; trinket2: Dormant Heart of the Mountain; ranged: Totem of Urgency

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-000000000000000000-2530000010000000)

Set DPS (verified): 72.9. Weights run: 4.1s. Verify run: 26.1s. 249 eligible items had no known source.

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

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads

### Band 30 (orc, 0000000000000000-000000000000000000-5530001010411000)

Set DPS (verified): 133.1. Weights run: 4.4s. Verify run: 29.8s. 442 eligible items had no known source.

Race sweep: orc 133.1 ± 0.0, windshaper-skyborne 133.0 ± 0.0, troll 131.9 ± 0.0, tauren 131.7 ± 0.0 (heal_score); the guide race orc is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.116, intellect=1.054 ± 0.013, spirit=0.645 ± 0.008, mp5=1.902 ± 0.024, crit=0.172 ± 0.008 per rating point (14 rating = 1%, 2.409 per %), spell_haste=not significant (-0.276 ± 0.085)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.9 healing_power points (9.00 DPS) | yes | Filigreed Pristine Circlet (253975, -3.63 DPS, sim-verified) [crafted]; Enduring Cap (3020, -3.78 DPS) [world_drop]; Wisdom's Leather Helm (252515, -3.87 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 6.8 healing_power points (1.66 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.12 DPS) [world]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.37 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.5 healing_power points (4.51 DPS) | yes | Death Speaker Mantle (6685, -0.22 DPS) [dungeon]; Magician's Mantle (12998, -0.98 DPS) [world_drop]; Feline Mantle (3748, -1.48 DPS) [dungeon] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.66 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -1.27 DPS) [vendor]; Windsong Drape (15468, -1.95 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.6 healing_power points (7.71 DPS) | yes | Pristine Gown (253961, -0.57 DPS, sim-verified) [crafted]; Wisdom's Leather Tunic (252511, -0.83 DPS) [crafted]; Beguiler Robes (7728, -1.17 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.7 healing_power points (3.59 DPS) | yes | Mindthrust Bracers (1974, -0.52 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.53 DPS) [world_drop]; Nightsky Wristbands (6407, -1.57 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.2 healing_power points (4.43 DPS) | yes | Naga Battle Gloves (888, -0.91 DPS, sim-verified) [dungeon]; Wisdom's Leather Gloves (252499, -0.96 DPS) [crafted]; Pristine Gloves (253913, -0.98 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 29.2 healing_power points (7.12 DPS) | yes | Skycaller's Leather Belt (252522, -2.49 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -3.37 DPS) [dungeon]; Pristine Sash (253925, -3.41 DPS) [crafted] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 32.9 healing_power points (8.03 DPS) | yes | Pristine Leggings (253987, -0.07 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.97 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.46 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.0 healing_power points (7.31 DPS) | yes | Wisdom's Leather Boots (252444, -3.40 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.40 DPS) [dungeon]; Black Whelp Slippers (252424, -4.09 DPS) [crafted] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 14.2 healing_power points (3.47 DPS) | yes | Runemender's Seal (273031, -0.17 DPS) [world]; Ring of Power Regulation (273030, -0.61 DPS) [dungeon]; Snake Hoop (6750, -0.72 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.5 healing_power points (3.30 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Ring of Power Regulation (273030, -0.44 DPS) [dungeon]; Snake Hoop (6750, -0.55 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -2.02 DPS) [dungeon]; Taskmaster Axe (5194, -4.36 DPS) [dungeon]; Staff of the Friar (3415, -4.96 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.9 healing_power points (4.12 DPS) | yes | Heart of Agamaggan (6694, -0.42 DPS, sim-verified) [dungeon]; Orb of Mystic Insight (249394, -0.87 DPS) [crafted]; Orb of Mistmantle (13031, -1.29 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 0000000000000000-000000000000000000-5530001011513051)

Set DPS (verified): 204.1. Weights run: 8.1s. Verify run: 82.1s. 679 eligible items had no known source.

Race sweep: windshaper-skyborne 204.2 ± 0.1, orc 204.1 ± 0.1, troll 202.4 ± 0.1, tauren 202.3 ± 0.1 (heal_score); windshaper-skyborne beats the guide race orc by 0.1 (0.04%) ± 0.1, inside error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.322, intellect=1.989 ± 0.026, spirit=1.213 ± 0.024, mp5=3.950 ± 0.026, crit=0.343 ± 0.025 per rating point (14 rating = 1%, 4.803 per %), spell_haste=not significant (0.889 ± 0.340)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 66.7 healing_power points (13.11 DPS) | yes | Whitemane's Chapeau (7720, -2.55 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -4.09 DPS) [dungeon]; Corpseshroud (10574, -4.26 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 35.9 healing_power points (7.06 DPS) | yes | Prodigious Shadowshard Pendant (17773, -3.15 DPS) [quest]; Triune Amulet (7722, -3.82 DPS) [dungeon]; Necklace of Calisea (1714, -4.19 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 47.9 healing_power points (9.41 DPS) | yes | Sheepshear Mantle (13115, -3.24 DPS, sim-verified) [world_drop]; Green Silken Shoulders (7057, -3.53 DPS) [crafted]; Mantle of Doan (7712, -3.78 DPS) [dungeon] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackforge Cape (6424, -0.65 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -0.88 DPS, sim-verified) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.7 healing_power points (13.31 DPS) | yes | Dreamweave Vest (10021, -2.91 DPS) [crafted]; Robe of Power (7054, -3.96 DPS) [crafted]; Deathchill Armor (10764, -5.12 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanted Kodo Bracers (13119, -0.15 DPS) [world_drop]; Radiant Silver Bracers (4545, -0.18 DPS) [quest]; Mindthrust Bracers (1974, -0.44 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.37 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.78 DPS) [crafted]; Stormcloth Gloves (10011, -0.79 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 45.2 healing_power points (8.88 DPS) | yes | Mender's Leather Belt (252523, -1.35 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.64 DPS) [dungeon]; Prefect's Belt (250559, -1.82 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 52.9 healing_power points (10.40 DPS) | yes | Stoneweaver Leggings (9407, -0.58 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.87 DPS) [crafted]; Pristine Leggings (253987, -2.15 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 46.4 healing_power points (9.12 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.50 DPS) [crafted]; Furen's Boots (13100, -2.57 DPS) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 23.7 healing_power points (4.66 DPS) | yes | Welken Ring (5011, -1.04 DPS) [world_drop]; Ring of Power Regulation (273030, -1.15 DPS) [dungeon]; Advisor's Ring (19520, -4.05 DPS, sim-verified) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Ring (19520, -0.19 DPS, sim-verified) [rep]; Welken Ring (5011, -0.30 DPS) [world_drop]; Ring of Power Regulation (273030, -0.41 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -6.13 DPS) [dungeon]; Illusionary Rod (7713, -6.78 DPS) [dungeon]; Staff of Jordan (873, -7.19 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Forcestone Buckler (17508, -0.29 DPS, sim-verified) [quest]; Battered Viking Shield (9403, -0.34 DPS) [dungeon]; Prophetic Cane (6803, -0.59 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 0000000000000000-055000000000000000-5530001011513051)

Set DPS (verified): 272.4. Weights run: 9.7s. Verify run: 133.0s. 862 eligible items had no known source.

Race sweep: orc 272.6 ± 0.1, windshaper-skyborne 272.4 ± 0.1, troll 270.1 ± 0.1, tauren 270.1 ± 0.1 (heal_score); the guide race orc is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.179, intellect=1.594 ± 0.028, spirit=0.716 ± 0.021, mp5=2.519 ± 0.028, crit=0.384 ± 0.021 per rating point (14 rating = 1%, 5.378 per %), spell_haste=2.628 ± 0.533

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 55.5 healing_power points (15.24 DPS) | yes | Soothsayer's Headdress (17740, -0.76 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -1.28 DPS) [crafted]; Cassandra's Grace (13102, -1.40 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.8 healing_power points (8.99 DPS) | yes | Horizon Choker (13085, -1.68 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -3.43 DPS) [dungeon]; Mindburst Medallion (11196, -3.73 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 52.6 healing_power points (14.43 DPS) | yes | Mender's Mail Shoulders (252569, -2.54 DPS) [crafted]; Nethergeld Shoulders (254049, -2.62 DPS) [crafted]; Mender's Leather Shoulders (252538, -3.05 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.36 DPS, sim-verified) [quest]; Battle Healer's Cloak (19527, -0.43 DPS) [rep]; Darkspear Raider's Cloak (272076, -0.54 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.7 healing_power points (17.48 DPS) | yes | Feathered Breastplate (8349, -2.64 DPS, sim-verified) [crafted]; Acumen Robes (17775, -3.52 DPS) [quest]; Robes of Insight (940, -3.60 DPS) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 34.7 healing_power points (9.53 DPS) | yes | Mender's Leather Bracers (252543, -0.08 DPS) [crafted]; Mender's Mail Bracers (252573, -0.08 DPS) [crafted]; Aristocratic Cuffs (12546, -1.79 DPS) [dungeon] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.31 DPS) [crafted]; Greenleaf Handwraps (19116, -3.36 DPS) [quest] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 49.1 healing_power points (13.48 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Waistcord (254081, -1.59 DPS) [crafted]; Dawnspire Cord (12466, -1.88 DPS) [dungeon] |
| legs | Stone Guard's Inscribed Legplates (220839) | Lady Palanseer [vendor] | 50.5 healing_power points (13.86 DPS) | yes | Senior Designer's Pantaloons (11841, -0.55 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.80 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -1.36 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 45.5 healing_power points (12.49 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Leather Shoes (252533, -1.47 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, -0.81 DPS) [dungeon]; Eye of Adaegus (5266, -1.08 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 23.0 healing_power points (6.32 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, -1.07 DPS) [dungeon]; Eye of Adaegus (5266, -1.34 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.32 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -17.59 DPS) [dungeon]; Resurgence Rod (17743, -17.92 DPS, sim-verified) [quest]; Dreamstaff (249454, -23.07 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Ironfeather Shoulders; back: Spritecaster Cape; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Stone Guard's Inscribed Legplates; feet: Mender's Leather Boots; finger1: Advisor's Ring; finger2: Cyclopean Band; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 862, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 359.0. Weights run: 11.2s. Verify run: 320.5s. 2043 eligible items had no known source.

Race sweep: orc 359.2 ± 0.2, windshaper-skyborne 358.4 ± 0.2, troll 356.2 ± 0.2, tauren 355.6 ± 0.2 (heal_score); the guide race orc is best.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.031, intellect=1.987 ± 0.030, spirit=1.037 ± 0.025, mp5=2.927 ± 0.027, crit=0.338 ± 0.013 per rating point (14 rating = 1%, 4.736 per %), spell_haste=not significant (0.213 ± 0.275)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 123.3 healing_power points (18.67 DPS) | yes | Warlord's Mail Skullcap (231669, -6.17 DPS) [vendor]; Crown of the Penitent (13216, -6.38 DPS) [quest]; Devout Crown (16693, -9.01 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 69.6 healing_power points (10.55 DPS) | yes | Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Darkspear Shoulders (272104, -0.61 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.85 DPS, sim-verified) [vendor] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -0.85 DPS) [quest]; Frostweaver Cape (12968, -0.97 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -1.19 DPS, sim-verified) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Warlord's Mail Chestguard (231668, +0.00 DPS) [vendor]; Devout Robe (16690, -0.33 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Mender's Leather Bracers (252543, -0.91 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gloves (231666, -0.47 DPS) [vendor]; Hands of the Exalted Herald (12554, -1.62 DPS, sim-verified) [dungeon]; Mender's Leather Gauntlets (252551, -2.94 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 68.2 healing_power points (10.32 DPS) | yes | Sash of Mercy (14553, -0.72 DPS) [world_drop]; Stormseeker's Girdle (272399, -1.09 DPS, sim-verified) [vendor]; Devout Belt (16696, -1.22 DPS) [dungeon] |
| legs | Red Dragonscale Leggings (252603) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; General's Mail Pants (231670, +0.00 DPS) [vendor]; General's Mail Leggings (231664, -0.48 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 65.3 healing_power points (9.89 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Faith Healer's Boots (22247, -0.46 DPS, sim-verified) [dungeon]; General's Mail Sabatons (231661, -1.75 DPS) [vendor] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.03 DPS) [vendor]; Uther's Strength (11302, -3.63 DPS) [world_drop]; Second Wind (11819, -4.39 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -8.39 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -11.02 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Urgency (279249) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -1.42 DPS, sim-verified) [vendor] |

**New at 60:** head: Living Crown; shoulder: Darkspear Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Red Dragonscale Leggings; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Ankh of Life; trinket2: Serenity Field; ranged: Totem of Urgency

No-known-source sample (15 of 2043, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 873.5. Weights run: 7.0s. Verify run: 290.0s. 2043 eligible items had no known source.

Race sweep: windshaper-skyborne 876.6 ± 0.6, orc 873.8 ± 0.5, troll 871.8 ± 0.6, tauren 867.3 ± 0.6 (heal_score); windshaper-skyborne beats the guide race orc by 2.8 (0.32%) ± 0.8, beyond error.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.145, intellect=1.254 ± 0.050, spirit=0.668 ± 0.040, mp5=2.205 ± 0.045, crit=0.671 ± 0.034 per rating point (14 rating = 1%, 9.390 per %), spell_haste=4.112 ± 0.994

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 104.9 healing_power points (26.03 DPS) | yes | Warlord's Mail Skullcap (231669, -7.50 DPS) [vendor]; Champion's Mail Skullcap (227168, -9.94 DPS) [vendor]; Mooncloth Circlet (14140, -24.69 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.19 DPS) [world_drop] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 49.9 healing_power points (12.37 DPS) | yes | Champion's Mail Epaulets (227166, +0.00 DPS) [vendor]; Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Royal Cap Spaulders (14548, -3.84 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raincaster Drape (12110, -2.93 DPS) [quest]; Arcanoweave Cloak (272411, -3.42 DPS) [vendor]; Battle Healer's Cloak (19526, -4.59 DPS, sim-verified) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | 81.6 healing_power points (20.25 DPS) | yes | Warlord's Mail Chestguard (231668, -0.78 DPS) [vendor]; Legionnaire's Mail Chestguard (227171, -3.14 DPS) [vendor]; Mooncloth Vest (14138, -8.77 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.58 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.14 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Blood Guard's Mail Gloves (227169, -2.03 DPS) [vendor]; Hands of the Exalted Herald (12554, -2.49 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Stormseeker's Girdle (272399, -2.10 DPS) [vendor]; Mender's Leather Waistguard (252477, -2.36 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 70.0 healing_power points (17.36 DPS) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, +0.00 DPS) [vendor]; Leggings of Arcana (12756, -4.97 DPS, sim-verified) [quest] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.2 healing_power points (12.46 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Faith Healer's Boots (22247, -0.29 DPS) [dungeon]; Blood Guard's Mail Boots (227170, -1.54 DPS) [vendor] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+23.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, -4.96 DPS) [dungeon]; Weakness Analyzer (272438, -4.96 DPS) [vendor]; Uther's Strength (11302, -5.95 DPS) [world_drop] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -15.82 DPS) [pvp]; Ironbark Staff (20220, -16.57 DPS) [rep]; Resurgence Rod (17743, -32.48 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Urgency (279249) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -2.78 DPS, sim-verified) [vendor] |

**New at 60:** head: Living Crown; shoulder: Darkspear Shoulderguards; back: Drape of Recovery; chest: Robes of the Exalted; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Serenity Field; trinket2: Dormant Heart of the Mountain; ranged: Totem of Urgency

No-known-source sample (15 of 2043, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

