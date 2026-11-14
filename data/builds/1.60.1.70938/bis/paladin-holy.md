# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 48.6. Weights run: 7.7s. Verify run: 55.5s. 312 eligible items had no known source.

Race sweep: human 48.6 ± 0.0, dwarf 46.8 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.030, intellect=1.005 ± 0.006, spirit=0.883 ± 0.004, mp5=3.917 ± 0.010, crit=0.112 ± 0.005 per rating point (14 rating = 1%, 1.570 per %), spell_haste=5.956 ± 0.172

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (2.82 DPS) | yes | Pristine Circlet (253949, -0.46 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.54 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.34 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.8 healing_power points (0.24 DPS) | yes | Tarnished Locket (279870, -0.27 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.0 healing_power points (1.89 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.68 DPS) [crafted]; Serpent's Shoulders (5404, -1.06 DPS) [dungeon]; Slime-encrusted Pads (6461, -1.43 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.75 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.77 DPS) [crafted]; Sanguine Cape (14376, -0.91 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 22.7 healing_power points (3.05 DPS) | yes | Acolyte's Chain Shirt (250491, -0.22 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.27 DPS) [crafted]; Corsair's Overshirt (5202, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.8 healing_power points (2.25 DPS) | yes | Owl Bracers (4796, -1.58 DPS) [vendor]; Bright Bracers (3647, -1.71 DPS) [world_drop]; Drakewing Bands (12999, -1.90 DPS, sim-verified) [world_drop] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 15.0 healing_power points (2.02 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.14 DPS, sim-verified) [crafted]; Stormrider's Leather Gloves (252498, -0.81 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.0 healing_power points (2.02 DPS) | yes | Acolyte's Chain Belt (250516, -0.30 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.40 DPS) [crafted]; Novice Ardent's Sash (253887, -0.45 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | sim-verified (48.6 DPS) | yes | Wisdom's Leather Pants (252503, -0.04 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.06 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.14 DPS) [crafted] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 16.0 healing_power points (2.15 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.40 DPS) [crafted]; Black Whelp Slippers (252424, -0.40 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 12.8 healing_power points (1.72 DPS) | yes | Lavishly Jeweled Ring (1156, -0.91 DPS) [dungeon]; Band of Purification (12996, -1.01 DPS) [world_drop]; Volcanic Rock Ring (12053, -1.32 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 healing_power points (0.94 DPS) | yes | Band of Purification (12996, -0.23 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.40 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.54 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monastic Hammer (270005, -5.11 DPS) [quest]; Sword of the Fallen (286977, -5.78 DPS) [dungeon]; Verigan's Fist (6953, -6.73 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 312, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 30 (human, 50320015000100000-40000000000000000-00000000000000000)

Set DPS (verified): 106.6. Weights run: 8.7s. Verify run: 60.4s. 537 eligible items had no known source.

Race sweep: human 106.6 ± 0.0, dwarf 103.9 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.049, intellect=2.262 ± 0.011, spirit=1.555 ± 0.007, mp5=8.511 ± 0.016, crit=0.412 ± 0.011 per rating point (14 rating = 1%, 5.767 per %), spell_haste=-1.291 ± 0.137

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Enduring Cap (3020, -0.66 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -1.02 DPS) [world_drop]; Dokebi Hat (14584, -1.33 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.3 healing_power points (2.09 DPS) | yes | Crystal Starfire Medallion (5003, -0.17 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.23 DPS) [world]; Seal of the Expedition (271920, -0.54 DPS) [rep] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 30.9 healing_power points (4.23 DPS) | yes | Bloodmage Mantle (7684, -0.21 DPS) [dungeon]; Mantle of Honor (3560, -0.57 DPS) [quest]; Nightsky Mantle (4718, -0.69 DPS) [world_drop] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Repairman's Cape (9605, -0.29 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.46 DPS) [dungeon]; Caretaker's Cape (19533, -0.49 DPS) [rep] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.39 DPS) [crafted]; Beguiler Robes (7728, -0.42 DPS, sim-verified) [dungeon]; Wisdom's Leather Tunic (252511, -0.55 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindthrust Bracers (1974, -0.07 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -2.23 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.25 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 30.0 healing_power points (4.11 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.17 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -1.12 DPS) [crafted]; Acolyte's Gloves (250511, -1.19 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 41.9 healing_power points (5.74 DPS) | yes | Prefect's Belt (250559, -0.13 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -1.66 DPS) [crafted]; Highlander's Lizardhide Girdle (20105, -2.02 DPS) [rep] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 47.7 healing_power points (6.54 DPS) | yes | Pristine Leggings (253987, -0.29 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.63 DPS) [crafted]; Leech Pants (6910, -0.89 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.1 healing_power points (5.76 DPS) | yes | Acidic Walkers (9454, -2.45 DPS, sim-verified) [dungeon]; Acolyte's Boots (250506, -2.70 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.70 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 31.9 healing_power points (4.37 DPS) | yes | Lorekeeper's Ring (19525, -1.08 DPS) [rep]; Electrocutioner Lagnut (9447, -1.65 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 31.5 healing_power points (4.32 DPS) | yes | Lorekeeper's Ring (19525, -0.66 DPS, sim-verified) [rep]; Electrocutioner Lagnut (9447, -1.60 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crested Scepter (3414, -4.31 DPS) [dungeon]; Verigan's Fist (6953, -5.33 DPS) [quest]; Taskmaster Axe (5194, -5.63 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 22.3 healing_power points (3.06 DPS) | yes | Orb of Mystic Insight (249394, -0.29 DPS, sim-verified) [crafted]; Worn Turtle Shell Shield (6447, -0.73 DPS) [dungeon]; Satyr's Rod (15962, -0.77 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Expeditionary's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 537, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 50320115000100000-50000000000000000-05030000000000000)

Set DPS (verified): 193.8. Weights run: 6.6s. Verify run: 40.2s. 808 eligible items had no known source.

Race sweep: human 193.8 ± 0.0, dwarf 191.1 ± 0.1 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.039, intellect=1.645 ± 0.027, spirit=0.973 ± 0.015, mp5=4.758 ± 0.105, crit=0.492 ± 0.011 per rating point (14 rating = 1%, 6.885 per %), spell_haste=not significant (-0.214 ± 0.100)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 58.7 healing_power points (11.64 DPS) | yes | Whitemane's Chapeau (7720, -0.50 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -3.72 DPS) [dungeon]; Holy Shroud (2721, -3.94 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.2 healing_power points (6.57 DPS) | yes | Triune Amulet (7722, -3.64 DPS) [dungeon]; Necklace of Calisea (1714, -3.88 DPS, sim-verified) [dungeon]; Seal of the Expedition (271919, -4.29 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 43.4 healing_power points (8.60 DPS) | yes | Green Silken Shoulders (7057, -3.43 DPS) [crafted]; Earthen Silk Shoulders (254033, -3.49 DPS) [crafted]; Sheepshear Mantle (13115, -4.22 DPS, sim-verified) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackforge Cape (6424, -1.15 DPS) [dungeon]; Expeditionary's Cape (271923, -1.35 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.8 healing_power points (13.04 DPS) | yes | Dreamweave Vest (10021, -3.84 DPS) [crafted]; Death's Head Vestment (10581, -4.36 DPS) [dungeon]; Deathchill Armor (10764, -7.25 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 23.3 healing_power points (4.61 DPS) | yes | Enchanted Kodo Bracers (13119, -1.38 DPS) [world_drop]; Guardian Leather Bracers (4260, -1.47 DPS) [crafted]; Mindthrust Bracers (1974, -1.67 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonefingers (10765, -0.39 DPS) [dungeon]; Gilded Handwraps (254021, -0.45 DPS) [crafted]; Gloves of Holy Might (867, -3.79 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.0 healing_power points (8.13 DPS) | yes | Mender's Leather Belt (252523, -0.64 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.48 DPS) [crafted]; Deathmage Sash (10771, -1.85 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 47.5 healing_power points (9.42 DPS) | yes | Stoneweaver Leggings (9407, -0.70 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.62 DPS) [crafted]; Pristine Leggings (253987, -1.81 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 42.8 healing_power points (8.48 DPS) | yes | Mender's Mail Boots (252565, -0.09 DPS) [crafted]; Mender's Leather Shoes (252533, -0.81 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.47 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 24.7 healing_power points (4.90 DPS) | yes | Ring of Power Regulation (273030, -0.88 DPS) [dungeon]; Philanthropist's Ring (281635, -1.29 DPS) [quest]; Electrocutioner Lagnut (9447, -1.76 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 23.3 healing_power points (4.61 DPS) | yes | Ring of Power Regulation (273030, -0.61 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -1.00 DPS) [quest]; Electrocutioner Lagnut (9447, -1.48 DPS) [dungeon] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -5.17 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -9.71 DPS) [dungeon]; Bookmaker's Scepter (4122, -12.27 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-verified (193.7 DPS) | yes | Orb of Lorica (11262, -0.44 DPS, sim-verified) [quest]; Battered Viking Shield (9403, -0.63 DPS) [dungeon]; Orb of Souls (249395, -1.40 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Prefect's Boots; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 808, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 50320215012101230-50000000000000000-05030000000000000)

Set DPS (verified): 205.9. Weights run: 16.4s. Verify run: 231.3s. 1050 eligible items had no known source.

Race sweep: human 205.7 ± 0.2, dwarf 199.9 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.165, intellect=1.253 ± 0.022, spirit=0.750 ± 0.012, mp5=4.522 ± 0.023, crit=0.551 ± 0.027 per rating point (14 rating = 1%, 7.719 per %), spell_haste=not significant (1.227 ± 0.359)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 50.7 healing_power points (13.79 DPS) | yes | Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Papal Fez (9431, -0.19 DPS) [dungeon]; Soulcatcher Halo (10630, -0.39 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.0 healing_power points (8.16 DPS) | yes | Gemshard Heart (17707, -1.12 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.58 DPS) [world_drop]; Mindburst Medallion (11196, -3.47 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 57.9 healing_power points (15.73 DPS) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, -1.09 DPS) [vendor]; Ironfeather Shoulders (15067, -3.30 DPS, sim-verified) [crafted]; Prefect's Pauldrons (250579, -4.44 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Expeditionary's Cape (271921, -1.42 DPS) [rep]; Featherskin Cape (10843, -1.70 DPS) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 68.0 healing_power points (18.49 DPS) | yes | Knight's Imbued Armor (220813, -1.26 DPS) [vendor]; Robes of Insight (940, -1.35 DPS) [world_drop]; Stormcloth Vest (10020, -1.89 DPS, sim-verified) [crafted] |
| wrist | Prefect's Wristguards (250584) | Blacksmithing [crafted] | 33.0 healing_power points (8.97 DPS) | yes | Nethergeld Cuffs (254061, -0.14 DPS) [crafted]; Mender's Leather Bracers (252543, -0.20 DPS) [crafted]; Mender's Mail Bracers (252573, -0.20 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Prefect's Gloves (250554, -0.07 DPS) [crafted]; Gilded Gloves (254095, -0.41 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 45.0 healing_power points (12.24 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.68 DPS) [crafted]; Dawnspire Cord (12466, -0.88 DPS) [dungeon] |
| legs | Knight's Imbued Leggings (220809) | PvP rank 6 · Knight · Alliance [vendor] | 63.4 healing_power points (17.22 DPS) | yes | Senior Designer's Pantaloons (11841, -1.47 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Kilt (9474, -3.42 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -4.98 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 42.3 healing_power points (11.49 DPS) | yes | Mender's Leather Boots (252472, -0.13 DPS) [crafted]; Mender's Mail Sabatons (252579, -0.13 DPS) [crafted]; Prefect's Boots (250549, -0.82 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 30.1 healing_power points (8.18 DPS) | yes | Chivalrous Signet (20505, -0.90 DPS) [quest]; Runemender's Seal (273031, -2.47 DPS) [world]; Cyclopean Band (11824, -2.53 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 27.3 healing_power points (7.41 DPS) | yes | Chivalrous Signet (20505, -0.31 DPS, sim-verified) [quest]; Runemender's Seal (273031, -1.70 DPS) [world]; Cyclopean Band (11824, -1.77 DPS) [dungeon] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.25 DPS) [world_drop]; Sanctified Orb (20512, -2.07 DPS) [quest]; Evonice's Landin' Pilla (18951, -4.11 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanctified Orb (20512, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.41 DPS) [quest]; Uther's Strength (11302, -1.88 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -8.65 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -8.77 DPS) [dungeon]; Arbiter's Blade (11784, -9.17 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of Lorica (11262, -4.22 DPS, sim-verified) [quest]; Beacon of Hope (9393, -5.06 DPS) [dungeon]; Twisting Essence Jar (249456, -5.06 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Cassandra's Grace; shoulder: Mantle of Lost Hope; back: Caretaker's Cape; chest: Embrace of the Wind Serpent; wrist: Prefect's Wristguards; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Gilded Sandals; finger1: Lorekeeper's Ring; finger2: Brainlash; trinket1: Enriched Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1050, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 276.0. Weights run: 17.8s. Verify run: 901.7s. 2437 eligible items had no known source.

Race sweep: human 275.8 ± 0.3, dwarf 267.9 ± 0.3 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.138, intellect=1.376 ± 0.029, spirit=0.925 ± 0.015, mp5=5.577 ± 0.025, crit=0.846 ± 0.044 per rating point (14 rating = 1%, 11.839 per %), spell_haste=not significant (1.491 ± 0.558)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 110.7 healing_power points (26.08 DPS) | yes | Field Marshal's Lamellar Helmet (231640, -1.71 DPS) [vendor]; Lieutenant Commander's Lamellar Helmet (227149, -2.37 DPS) [vendor]; Crown of the Penitent (13216, -7.02 DPS, sim-verified) [quest] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 52.0 healing_power points (12.25 DPS) | yes | Tooth of Gnarr (13141, -1.21 DPS) [dungeon]; Animated Chain Necklace (18723, -2.95 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.50 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Mantle of Lost Hope (22234, -7.66 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gracious Cape (18743, -0.76 DPS) [dungeon]; Caretaker's Cape (19530, -4.69 DPS) [rep]; Shroud of the Exile (15421, -4.96 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Breastplate of Salvation (250601, -5.93 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brazecore Armguards (13179, +0.00 DPS) [dungeon]; Morlune's Bracer (18741, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, -5.49 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Theramore Handwraps (271945, -7.52 DPS, sim-verified) [rep] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Feralsurge Girdle (18104, -6.46 DPS, sim-verified) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Legguards (227150, -2.42 DPS) [vendor]; Ghostloom Leggings (14545, -3.15 DPS, sim-verified) [dungeon]; Marshal's Lamellar Legguards (231639, -4.37 DPS) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -8.43 DPS, sim-verified) [crafted] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -2.88 DPS) [quest]; Band of Rumination (18103, -4.18 DPS) [dungeon]; Naglering (11669, -12.61 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.22 DPS) [quest]; Band of Rumination (18103, -1.52 DPS) [dungeon]; Blessed Band of Light (272407, -1.83 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+10.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -3.94 DPS) [rep]; Serenity Field (272439, -5.56 DPS) [vendor]; Weakness Analyzer (272438, -6.27 DPS) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, -1.33 DPS, sim-verified) [rep]; Serenity Field (272439, -2.94 DPS) [vendor]; Weakness Analyzer (272438, -3.64 DPS) [vendor] |
| main_hand | The Hammer of Grace (11923) | Blackrock Depths: Chest of The Seven [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hammer of the Grand Crusader (18717, -1.23 DPS, sim-verified) [dungeon]; Hand of Edward the Odd (2243, -1.62 DPS) [world_drop] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 64.9 healing_power points (15.28 DPS) | yes | Grand Marshal's Tome of Restoration (234590, -0.38 DPS) [pvp]; Grand Marshal's Barricade (235473, -1.68 DPS) [vendor]; Gizlock's Hypertech Buckler (17718, -3.73 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Band of Mending; trinket1: Burst of Knowledge; trinket2: Shard of the Splithooves; main_hand: The Hammer of Grace; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 2437, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 721.0. Weights run: 9.9s. Verify run: 483.7s. 2437 eligible items had no known source.

Race sweep: human 722.1 ± 0.7, dwarf 706.3 ± 0.7 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.594, intellect=1.380 ± 0.065, spirit=0.420 ± 0.059, mp5=2.695 ± 0.073, crit=1.276 ± 0.077 per rating point (14 rating = 1%, 17.867 per %), spell_haste=not significant (1.863 ± 1.315)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 103.7 healing_power points (45.03 DPS) | yes | Field Marshal's Lamellar Helmet (231640, -1.21 DPS) [vendor]; Lieutenant Commander's Lamellar Helmet (227149, -7.56 DPS) [vendor]; Black Dragonscale Helm (252605, -11.62 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 47.0 healing_power points (20.42 DPS) | yes | Lady Maye's Pendant (14558, -4.26 DPS) [world_drop]; Animated Chain Necklace (18723, -4.82 DPS) [dungeon]; Tooth of Gnarr (13141, -5.20 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Theramore Shoulderguards (272957, -26.13 DPS, sim-verified) [rep] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gracious Cape (18743, -4.06 DPS, sim-verified) [dungeon]; Caretaker's Cape (19530, -5.39 DPS) [rep]; Shroud of the Exile (15421, -5.66 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Breastplate of Salvation (250601, -24.64 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Funeral Cuffs (12626, +0.00 DPS) [dungeon]; Morlune's Bracer (18741, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, -22.27 DPS, sim-verified) [dungeon] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Hands of the Exalted Herald (12554, -3.61 DPS, sim-verified) [dungeon]; Prefect's Gloves (250554, -4.81 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of the Ordained (18702, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Sash of Mercy (14553, -25.37 DPS, sim-verified) [world_drop] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legguards (231639, -1.48 DPS) [vendor]; Sentinel's Silk Leggings (237815, -5.54 DPS, sim-verified) [vendor]; Knight-Captain's Lamellar Legguards (227150, -6.76 DPS) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Marshal's Lamellar Greaves (231646, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -23.51 DPS, sim-verified) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -2.22 DPS) [dungeon]; Blessed Band of Light (272407, -4.74 DPS) [vendor]; Naglering (11669, -19.04 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -2.63 DPS, sim-verified) [dungeon]; Blessed Band of Light (272407, -3.83 DPS) [vendor]; Channeler's Ring (272406, -4.99 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -0.24 DPS) [vendor]; Eye of the Beast (13968, -1.37 DPS, sim-verified) [quest]; Weakness Analyzer (272438, -1.54 DPS) [vendor] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -5.15 DPS, sim-verified) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, -22.49 DPS) [vendor]; Hand of Edward the Odd (2243, -47.05 DPS, sim-verified) [world_drop]; Grand Marshal's Demolisher (234568, -50.60 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Devout Robe; wrist: Devout Bracers; hands: Theramore Handwraps; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Fordring's Seal; finger2: Band of Mending; trinket1: Burst of Knowledge; trinket2: Dormant Heart of the Mountain; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2437, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 47.9. Weights run: 7.7s. Verify run: 50.7s. 299 eligible items had no known source.

Race sweep: undead 47.9 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.030, intellect=1.005 ± 0.006, spirit=0.883 ± 0.004, mp5=3.917 ± 0.010, crit=0.112 ± 0.005 per rating point (14 rating = 1%, 1.570 per %), spell_haste=5.956 ± 0.172

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (2.82 DPS) | yes | Pristine Circlet (253949, -0.46 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.54 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.34 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.6 healing_power points (0.36 DPS) | yes | Scholarly Pendant (277203, -0.25 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.0 healing_power points (1.89 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.68 DPS) [crafted]; Serpent's Shoulders (5404, -1.06 DPS) [dungeon]; Slime-encrusted Pads (6461, -1.47 DPS, sim-verified) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-verified (47.9 DPS) | yes | Heavy Woolen Cloak (4311, -0.76 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.77 DPS) [crafted]; Sanguine Cape (14376, -0.91 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 22.7 healing_power points (3.05 DPS) | yes | Acolyte's Chain Shirt (250491, -0.22 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.27 DPS) [crafted]; Corsair's Overshirt (5202, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.8 healing_power points (2.25 DPS) | yes | Owl Bracers (4796, -1.58 DPS) [vendor]; Bright Bracers (3647, -1.71 DPS) [world_drop]; Drakewing Bands (12999, -1.92 DPS, sim-verified) [world_drop] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 15.0 healing_power points (2.02 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.20 DPS, sim-verified) [crafted]; Stormrider's Leather Gloves (252498, -0.81 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.0 healing_power points (2.02 DPS) | yes | Acolyte's Chain Belt (250516, -0.31 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.40 DPS) [crafted]; Novice Ardent's Sash (253887, -0.45 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 29.6 healing_power points (3.97 DPS) | yes | Acolyte's Chain Leggings (250496, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.34 DPS) [crafted] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 16.0 healing_power points (2.15 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.40 DPS) [crafted]; Black Whelp Slippers (252424, -0.40 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 12.8 healing_power points (1.72 DPS) | yes | Band of Purification (12996, -1.01 DPS) [world_drop]; Loop of Sacrifice (281673, -1.05 DPS) [quest]; Volcanic Rock Ring (12053, -1.32 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.0 healing_power points (0.81 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; Band of Purification (12996, -0.70 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | 52.1 healing_power points (6.99 DPS) | yes | The Stitcher (279874, -2.59 DPS, sim-verified) [quest]; Pathfinder's Clearway (277288, -4.57 DPS) [quest]; Sword of the Fallen (286977, -5.78 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Acolyte's Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 299, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (undead, 50320015000100000-40000000000000000-00000000000000000)

Set DPS (verified): 104.7. Weights run: 8.7s. Verify run: 61.2s. 524 eligible items had no known source.

Race sweep: undead 104.7 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.049, intellect=2.262 ± 0.011, spirit=1.555 ± 0.007, mp5=8.511 ± 0.016, crit=0.412 ± 0.011 per rating point (14 rating = 1%, 5.767 per %), spell_haste=-1.291 ± 0.137

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Enduring Cap (3020, -0.56 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -1.02 DPS) [world_drop]; Dokebi Hat (14584, -1.33 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.3 healing_power points (2.09 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.23 DPS) [world]; Crystal Starfire Medallion (5003, -0.31 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.54 DPS) [vendor] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 30.9 healing_power points (4.23 DPS) | yes | Bloodmage Mantle (7684, -0.11 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.69 DPS) [world_drop]; Slime-encrusted Pads (6461, -0.73 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19529, -0.49 DPS) [rep]; Glowing Thresher Cape (6901, -0.65 DPS, sim-verified) [dungeon]; Resilient Cape (14400, -1.26 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.39 DPS) [crafted]; Beguiler Robes (7728, -0.43 DPS, sim-verified) [dungeon]; Wisdom's Leather Tunic (252511, -0.55 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 36.8 healing_power points (5.05 DPS) | yes | Spidertank Oilrag (9448, -0.12 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -2.55 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.57 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 30.0 healing_power points (4.11 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.56 DPS) [dungeon]; Truefaith Gloves (7049, -1.12 DPS) [crafted]; Acolyte's Gloves (250511, -1.19 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 41.9 healing_power points (5.74 DPS) | yes | Prefect's Belt (250559, -0.56 DPS) [crafted]; Skycaller's Leather Belt (252522, -1.66 DPS) [crafted]; Defiler's Lizardhide Girdle (20172, -2.02 DPS) [rep] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 47.7 healing_power points (6.54 DPS) | yes | Pristine Leggings (253987, -0.29 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.63 DPS) [crafted]; Leech Pants (6910, -0.78 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.1 healing_power points (5.76 DPS) | yes | Acidic Walkers (9454, -2.47 DPS, sim-verified) [dungeon]; Acolyte's Boots (250506, -2.70 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.70 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 31.9 healing_power points (4.37 DPS) | yes | Advisor's Ring (19521, -1.08 DPS) [rep]; Electrocutioner Lagnut (9447, -1.65 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 31.5 healing_power points (4.32 DPS) | yes | Advisor's Ring (19521, -0.74 DPS, sim-verified) [rep]; Electrocutioner Lagnut (9447, -1.60 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crested Scepter (3414, -4.31 DPS) [dungeon]; The Stitcher (279874, -5.21 DPS) [quest]; Taskmaster Axe (5194, -5.72 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 22.3 healing_power points (3.06 DPS) | yes | Orb of Mystic Insight (249394, -0.36 DPS, sim-verified) [crafted]; Witch's Finger (16887, -0.58 DPS) [quest]; Worn Turtle Shell Shield (6447, -0.73 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 524, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 50320115000100000-50000000000000000-05030000000000000)

Set DPS (verified): 192.4. Weights run: 6.6s. Verify run: 40.6s. 792 eligible items had no known source.

Race sweep: undead 192.4 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.039, intellect=1.645 ± 0.027, spirit=0.973 ± 0.015, mp5=4.758 ± 0.105, crit=0.492 ± 0.011 per rating point (14 rating = 1%, 6.885 per %), spell_haste=not significant (-0.214 ± 0.100)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 58.7 healing_power points (11.64 DPS) | yes | Whitemane's Chapeau (7720, -1.53 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -3.72 DPS) [dungeon]; Holy Shroud (2721, -3.94 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.2 healing_power points (6.57 DPS) | yes | Triune Amulet (7722, -3.64 DPS) [dungeon]; Necklace of Calisea (1714, -3.74 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -4.29 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 43.4 healing_power points (8.60 DPS) | yes | Green Silken Shoulders (7057, -3.43 DPS) [crafted]; Earthen Silk Shoulders (254033, -3.49 DPS) [crafted]; Sheepshear Mantle (13115, -3.99 DPS, sim-verified) [world_drop] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (192.4 DPS) | yes | Blackforge Cape (6424, -1.15 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -1.35 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.8 healing_power points (13.04 DPS) | yes | Dreamweave Vest (10021, -3.84 DPS) [crafted]; Death's Head Vestment (10581, -4.36 DPS) [dungeon]; Deathchill Armor (10764, -7.00 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 23.3 healing_power points (4.61 DPS) | yes | Radiant Silver Bracers (4545, -1.21 DPS) [quest]; Enchanted Kodo Bracers (13119, -1.38 DPS) [world_drop]; Mindthrust Bracers (1974, -1.41 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-verified (192.4 DPS) | yes | Gilded Handwraps (254021, -0.45 DPS) [crafted]; Prefect's Gauntlet (250569, -0.65 DPS) [crafted]; Bonefingers (10765, -3.81 DPS, sim-verified) [dungeon] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.0 healing_power points (8.13 DPS) | yes | Mender's Leather Belt (252523, -0.91 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.48 DPS) [crafted]; Deathmage Sash (10771, -1.85 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 47.5 healing_power points (9.42 DPS) | yes | Stoneweaver Leggings (9407, -0.62 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.62 DPS) [crafted]; Pristine Leggings (253987, -1.81 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 42.8 healing_power points (8.48 DPS) | yes | Mender's Mail Boots (252565, -0.09 DPS) [crafted]; Mender's Leather Shoes (252533, -0.66 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.47 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 24.7 healing_power points (4.90 DPS) | yes | Ring of Power Regulation (273030, -0.88 DPS) [dungeon]; Philanthropist's Ring (281635, -1.29 DPS) [quest]; Solid Gold Band (279845, -1.75 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 23.3 healing_power points (4.61 DPS) | yes | Ring of Power Regulation (273030, -0.61 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -1.00 DPS) [quest]; Solid Gold Band (279845, -1.47 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (192.4 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (192.4 DPS) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (192.4 DPS) | yes | Death Speaker Scepter (2816, -5.88 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -9.71 DPS) [dungeon]; Skullbreaker (17039, -12.60 DPS) [quest] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-verified (192.4 DPS) | yes | Beacon of Hope (9393, -0.73 DPS, sim-verified) [dungeon]; Forcestone Buckler (17508, -1.06 DPS) [quest]; Omega Orb (7749, -1.56 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; wrist: Spidertank Oilrag; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Prefect's Boots; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Prophetic Cane

No-known-source sample (15 of 792, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 50320215012101230-50000000000000000-05030000000000000)

Set DPS (verified): 199.5. Weights run: 16.4s. Verify run: 223.6s. 1050 eligible items had no known source.

Race sweep: undead 199.2 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.165, intellect=1.253 ± 0.022, spirit=0.750 ± 0.012, mp5=4.522 ± 0.023, crit=0.551 ± 0.027 per rating point (14 rating = 1%, 7.719 per %), spell_haste=not significant (1.227 ± 0.359)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulcatcher Halo (10630, -0.20 DPS) [dungeon]; Cassandra's Grace (13102, -0.27 DPS, sim-verified) [world_drop]; Gemburst Circlet (10751, -0.61 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.0 healing_power points (8.16 DPS) | yes | Gemshard Heart (17707, -1.14 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.58 DPS) [world_drop]; Mindburst Medallion (11196, -3.47 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 57.9 healing_power points (15.73 DPS) | yes | Ironfeather Shoulders (15067, -3.28 DPS, sim-verified) [crafted]; Prefect's Pauldrons (250579, -4.44 DPS) [crafted]; Living Shoulders (15061, -4.65 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Deep Woodlands Cloak (19121, -0.88 DPS) [quest]; Darkspear Raider's Cloak (272076, -1.42 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 68.0 healing_power points (18.49 DPS) | yes | Robes of Insight (940, -1.35 DPS) [world_drop]; Stormcloth Vest (10020, -1.47 DPS, sim-verified) [crafted]; Ghostweave Vest (14141, -2.13 DPS) [crafted] |
| wrist | Prefect's Wristguards (250584) | Blacksmithing [crafted] | 33.0 healing_power points (8.97 DPS) | yes | Nethergeld Cuffs (254061, -0.14 DPS) [crafted]; Mender's Leather Bracers (252543, -0.20 DPS) [crafted]; Mender's Mail Bracers (252573, -0.20 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Prefect's Gloves (250554, -0.07 DPS) [crafted]; Gilded Gloves (254095, -0.41 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 45.0 healing_power points (12.24 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.68 DPS) [crafted]; Dawnspire Cord (12466, -0.88 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 59.3 healing_power points (16.11 DPS) | yes | Jinxed Hoodoo Kilt (9474, -2.13 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -3.87 DPS) [dungeon]; Warchief Kilt (7760, -4.69 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 42.3 healing_power points (11.49 DPS) | yes | Mender's Mail Sabatons (252579, -0.13 DPS) [crafted]; Mender's Leather Boots (252472, -0.28 DPS, sim-verified) [crafted]; Prefect's Boots (250549, -0.82 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 30.1 healing_power points (8.18 DPS) | yes | Runemender's Seal (273031, -2.47 DPS) [world]; Cyclopean Band (11824, -2.53 DPS) [dungeon]; Ring of Power Regulation (273030, -2.86 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 27.3 healing_power points (7.41 DPS) | yes | Runemender's Seal (273031, -1.70 DPS, sim-verified) [world]; Cyclopean Band (11824, -1.77 DPS) [dungeon]; Ring of Power Regulation (273030, -2.10 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.25 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.11 DPS) [quest]; Alchemist's Stone (13503, -5.33 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.41 DPS) [quest]; Alchemist's Stone (13503, -1.63 DPS) [crafted]; Uther's Strength (11302, -1.75 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -8.29 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -8.77 DPS) [dungeon]; Arbiter's Blade (11784, -9.17 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prophetic Cane (6803, -4.73 DPS, sim-verified) [quest]; Beacon of Hope (9393, -5.06 DPS) [dungeon]; Twisting Essence Jar (249456, -5.06 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Mantle of Lost Hope; back: Battle Healer's Cloak; chest: Embrace of the Wind Serpent; wrist: Prefect's Wristguards; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Advisor's Ring; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1050, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 267.9. Weights run: 17.8s. Verify run: 847.1s. 2462 eligible items had no known source.

Race sweep: undead 267.9 ± 0.3 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.138, intellect=1.376 ± 0.029, spirit=0.925 ± 0.015, mp5=5.577 ± 0.025, crit=0.846 ± 0.044 per rating point (14 rating = 1%, 11.839 per %), spell_haste=not significant (1.491 ± 0.558)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 110.7 healing_power points (26.08 DPS) | yes | Crown of the Penitent (13216, -7.17 DPS, sim-verified) [quest]; Whitesoul Helm (12633, -9.70 DPS) [crafted]; Devout Crown (16693, -10.32 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 52.0 healing_power points (12.25 DPS) | yes | Tooth of Gnarr (13141, -0.93 DPS, sim-verified) [dungeon]; Animated Chain Necklace (18723, -2.95 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.50 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawnbringer Shoulders (12625, +0.00 DPS) [crafted]; Defiler's Epaulets (20176, +0.00 DPS) [rep]; Mantle of Lost Hope (22234, -7.81 DPS, sim-verified) [dungeon] |
| back | Gracious Cape (18743) | Stratholme: Ghostly Citizen [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Drape of Recovery (272413, +0.00 DPS) [vendor]; Battle Healer's Cloak (19526, -3.93 DPS) [rep]; Shroud of the Exile (15421, -4.21 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Breastplate of Salvation (250601, -6.47 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brazecore Armguards (13179, +0.00 DPS) [dungeon]; Morlune's Bracer (18741, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, -6.25 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Gilded Gauntlets (13244, +0.00 DPS) [dungeon]; Raider Handwraps (272097, -8.36 DPS, sim-verified) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Feralsurge Girdle (18104, -7.15 DPS, sim-verified) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ghostloom Leggings (14545, -3.51 DPS, sim-verified) [dungeon]; Legplates of the Chromatic Defier (12945, -5.96 DPS) [quest]; Outrider's Mail Leggings (22676, -7.85 DPS) [rep] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Firemoss Boots (22275, +0.00 DPS) [dungeon]; Mooncloth Boots (15802, -8.65 DPS, sim-verified) [crafted] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -2.88 DPS) [quest]; Band of Rumination (18103, -4.18 DPS) [dungeon]; Naglering (11669, -13.32 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.22 DPS) [quest]; Band of Rumination (18103, -1.52 DPS) [dungeon]; Blessed Band of Light (272407, -1.83 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -3.94 DPS) [vendor]; Serenity Field (272439, -5.56 DPS) [vendor]; Weakness Analyzer (272438, -6.27 DPS) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, -1.76 DPS, sim-verified) [vendor]; Serenity Field (272439, -2.94 DPS) [vendor]; Weakness Analyzer (272438, -3.64 DPS) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; The Hammer of Grace (11923, -14.72 DPS) [dungeon]; Hand of Edward the Odd (2243, -16.34 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Gracious Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Band of Mending; trinket1: Burst of Knowledge; trinket2: Shard of the Splithooves; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 705.2. Weights run: 9.9s. Verify run: 471.6s. 2462 eligible items had no known source.

Race sweep: undead 705.9 ± 0.7 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.594, intellect=1.380 ± 0.065, spirit=0.420 ± 0.059, mp5=2.695 ± 0.073, crit=1.276 ± 0.077 per rating point (14 rating = 1%, 17.867 per %), spell_haste=not significant (1.863 ± 1.315)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 103.7 healing_power points (45.03 DPS) | yes | Black Dragonscale Helm (252605, -14.51 DPS) [crafted]; Whitesoul Helm (12633, -18.11 DPS) [crafted]; Devout Crown (16693, -19.23 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 47.0 healing_power points (20.42 DPS) | yes | Tooth of Gnarr (13141, -3.81 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.26 DPS) [world_drop]; Animated Chain Necklace (18723, -4.82 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Darkspear Mantle (272107, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gracious Cape (18743, -3.89 DPS) [dungeon]; Battle Healer's Cloak (19526, -5.39 DPS) [rep]; Shroud of the Exile (15421, -5.66 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of Insight (940, +0.00 DPS) [world_drop]; Mooncloth Vest (14138, +0.00 DPS) [crafted]; Breastplate of Salvation (250601, +0.00 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Funeral Cuffs (12626, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Morlune's Bracer (18741, +0.00 DPS) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -3.16 DPS) [dungeon]; Prefect's Gloves (250554, -4.81 DPS) [crafted]; Mender's Leather Gauntlets (252551, -6.10 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Belt of the Ordained (18702, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironfeather Leggings (252486, -8.32 DPS) [crafted]; Legplates of the Chromatic Defier (12945, -11.21 DPS) [quest]; Outrider's Mail Leggings (22676, -15.40 DPS) [rep] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Firemoss Boots (22275, +0.00 DPS) [dungeon] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.91 DPS) [dungeon]; Blessed Band of Light (272407, -4.74 DPS) [vendor]; Channeler's Ring (272406, -5.90 DPS) [vendor] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -2.52 DPS) [vendor]; Channeler's Ring (272406, -3.68 DPS) [vendor] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -0.24 DPS) [vendor]; Uther's Strength (11302, -1.54 DPS) [world_drop]; Weakness Analyzer (272438, -1.54 DPS) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, -22.49 DPS) [vendor]; High Warlord's Destroyer (234546, -50.60 DPS) [pvp]; Taskmaster Axe (5194, -56.29 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Devout Robe; wrist: Devout Bracers; hands: Raider Handwraps; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Fordring's Seal; finger2: Rosewine Circle; trinket1: Dormant Heart of the Mountain; trinket2: Burst of Knowledge; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

