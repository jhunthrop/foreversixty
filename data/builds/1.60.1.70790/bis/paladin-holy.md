# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 50.1. Weights run: 5.6s. Verify run: 43.2s. 312 eligible items had no known source.

Race sweep: human 50.1 ± 0.0, dwarf 48.2 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=0.455 ± 0.001, spirit=0.239 ± 0.001, mp5=1.175 ± 0.003, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.449 per %), spell_haste=0.906 ± 0.026

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (3.36 DPS) | yes | Pristine Circlet (253949, -0.48 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.64 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.60 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.5 healing_power points (0.08 DPS) | yes | - |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.1 healing_power points (1.46 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.36 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.82 DPS) [crafted]; Slime-encrusted Pads (6461, -0.89 DPS) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (50.1 DPS) | yes | Heavy Woolen Cloak (4311, -0.72 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.98 DPS) [crafted]; Black Whelp Cloak (7283, -1.04 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Acolyte's Chain Shirt (250491, +0.00 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.32 DPS) [crafted]; Bloody Apron (6226, -0.80 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 5.8 healing_power points (0.93 DPS) | yes | Owl Bracers (4796, -0.56 DPS) [vendor]; Rift Bracers (5943, -0.61 DPS) [dungeon]; Bright Bracers (3647, -0.64 DPS) [world_drop] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | sim-verified (50.1 DPS) | yes | Pristine Gloves (253913, +0.00 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.07 DPS) [crafted]; Serpent Gloves (5970, -0.84 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 12.8 healing_power points (2.05 DPS) | yes | Acolyte's Chain Belt (250516, -0.32 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.48 DPS) [crafted]; Novice Ardent's Sash (253887, -0.76 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | sim-verified (50.1 DPS) | yes | Wisdom's Leather Pants (252503, -0.08 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.68 DPS) [crafted] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 13.3 healing_power points (2.13 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.39 DPS) [crafted]; Black Whelp Slippers (252424, -0.39 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.4 healing_power points (1.18 DPS) | yes | Lavishly Jeweled Ring (1156, -0.74 DPS) [dungeon]; Band of Purification (12996, -0.95 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.96 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.9 healing_power points (0.95 DPS) | yes | Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; Band of Purification (12996, -0.72 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monastic Hammer (270005, -5.45 DPS) [quest]; Rhahk'Zor's Hammer (5187, -6.23 DPS) [dungeon]; Thornroot Club (5587, -6.55 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 312, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 30 (human, 50320015000100000-40000000000000000-00000000000000000)

Set DPS (verified): 112.6. Weights run: 6.3s. Verify run: 41.3s. 537 eligible items had no known source.

Race sweep: human 112.6 ± 0.0, dwarf 109.1 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.654 ± 0.006, spirit=0.979 ± 0.004, mp5=5.512 ± 0.008, crit=0.323 ± 0.007 per rating point (14 rating = 1%, 4.515 per %), spell_haste=not significant (-0.079 ± 0.024)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.9 healing_power points (8.64 DPS) | yes | Enduring Cap (3020, -1.47 DPS, sim-verified) [world_drop]; Acolyte's Chain Helm (250501, -3.08 DPS) [crafted]; Nightsky Cowl (4039, -3.14 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 10.5 healing_power points (2.34 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.12 DPS, sim-verified) [world]; Crystal Starfire Medallion (5003, -0.22 DPS) [world_drop]; Seal of the Expedition (271920, -0.50 DPS) [rep] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 24.2 healing_power points (5.38 DPS) | yes | Bloodmage Mantle (7684, -0.07 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.96 DPS) [world_drop]; Mantle of Honor (3560, -1.28 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-verified (112.6 DPS) | yes | Prelacy Cape (7004, -0.01 DPS) [quest]; Expeditionary's Cape (271923, -0.17 DPS) [rep]; Repairman's Cape (9605, -0.24 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 38.2 healing_power points (8.49 DPS) | yes | Beguiler Robes (7728, -0.38 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.60 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.98 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 25.5 healing_power points (5.67 DPS) | yes | Mindthrust Bracers (1974, -0.19 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.73 DPS) [world_drop]; Nightsky Wristbands (6407, -2.82 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 20.0 healing_power points (4.45 DPS) | yes | Truefaith Gloves (7049, -0.01 DPS) [crafted]; Acolyte's Gloves (250511, -0.39 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.42 DPS) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 34.8 healing_power points (7.73 DPS) | yes | Prefect's Belt (250559, -0.28 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -2.43 DPS) [crafted]; Highlander's Lizardhide Girdle (20105, -3.32 DPS) [rep] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 39.4 healing_power points (8.76 DPS) | yes | Pristine Leggings (253987, -0.14 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.27 DPS) [crafted]; Wisdom's Leather Pants (252503, -1.24 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 35.5 healing_power points (7.89 DPS) | yes | Acidic Walkers (9454, -2.77 DPS, sim-verified) [dungeon]; Acolyte's Boots (250506, -3.60 DPS) [crafted]; Wisdom's Leather Boots (252444, -3.60 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 25.8 healing_power points (5.73 DPS) | yes | Ring of Power Regulation (273030, -0.72 DPS) [dungeon]; Lorekeeper's Ring (19525, -1.72 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 24.4 healing_power points (5.43 DPS) | yes | Ring of Power Regulation (273030, -0.21 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19525, -1.42 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (112.6 DPS) | yes | Taskmaster Axe (5194, -7.25 DPS, sim-verified) [dungeon]; Crested Scepter (3414, -7.26 DPS) [dungeon]; Pysan's Old Greatsword (1975, -8.67 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.9 healing_power points (4.19 DPS) | yes | Orb of Mystic Insight (249394, -0.43 DPS, sim-verified) [crafted]; Heart of Agamaggan (6694, -0.86 DPS) [dungeon]; Seedcloud Buckler (6630, -1.24 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 537, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 50320115000100000-50000000000000000-05030000000000000)

Set DPS (verified): 193.6. Weights run: 5.1s. Verify run: 26.1s. 808 eligible items had no known source.

Race sweep: human 193.6 ± 0.0, dwarf 191.5 ± 0.1 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.091, intellect=1.012 ± 0.019, spirit=0.555 ± 0.011, mp5=3.146 ± 0.072, crit=0.316 ± 0.009 per rating point (14 rating = 1%, 4.418 per %), spell_haste=1.327 ± 0.098

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 44.2 healing_power points (12.53 DPS) | yes | Whitemane's Chapeau (7720, -0.62 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.23 DPS) [world_drop]; Miner's Hat of the Deep (9429, -4.25 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.1 healing_power points (7.96 DPS) | yes | Triune Amulet (7722, -2.64 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -4.85 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -5.41 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.2 healing_power points (9.96 DPS) | yes | Earthen Silk Shoulders (254033, -1.25 DPS, sim-verified) [crafted]; Sheepshear Mantle (13115, -4.43 DPS) [world_drop]; Green Silken Shoulders (7057, -4.54 DPS) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Prelacy Cape (7004, -1.83 DPS) [quest]; Expeditionary's Cape (271922, -2.10 DPS) [rep]; Repairman's Cape (9605, -2.54 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.4 healing_power points (17.69 DPS) | yes | Deathchill Armor (10764, -6.85 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -7.81 DPS) [crafted]; Death's Head Vestment (10581, -8.25 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.4 healing_power points (5.22 DPS) | yes | Earthen Silk Cuffs (254019, -1.60 DPS) [crafted]; Mindthrust Bracers (1974, -1.64 DPS, sim-verified) [dungeon]; Guardian Leather Bracers (4260, -1.80 DPS) [crafted] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Gloves (252530, -0.15 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.28 DPS) [crafted]; Gloves of the Greatfather (17721, -0.87 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 33.4 healing_power points (9.47 DPS) | yes | Prefect's Belt (250559, -1.23 DPS) [crafted]; Mender's Leather Belt (252523, -1.42 DPS) [crafted]; Deathmage Sash (10771, -3.19 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 37.8 healing_power points (10.72 DPS) | yes | Stoneweaver Leggings (9407, -0.76 DPS, sim-verified) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -1.35 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.66 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 37.1 healing_power points (10.52 DPS) | yes | Mender's Mail Boots (252565, -0.61 DPS) [crafted]; Mender's Leather Shoes (252533, -0.81 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -2.21 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.0 healing_power points (5.37 DPS) | yes | Snake Hoop (6750, -0.56 DPS) [quest]; Ring of Power Regulation (273030, -1.00 DPS) [dungeon]; Philanthropist's Ring (281635, -1.10 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 18.4 healing_power points (5.22 DPS) | yes | Snake Hoop (6750, -0.46 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -0.85 DPS) [dungeon]; Philanthropist's Ring (281635, -0.96 DPS) [quest] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -5.32 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -13.88 DPS) [dungeon]; Bookmaker's Scepter (4122, -17.67 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Lorica (11262, -0.51 DPS, sim-verified) [quest]; Battered Viking Shield (9403, -1.61 DPS) [dungeon]; Orb of Souls (249395, -2.24 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Gilded Handwraps; waist: Gilded Cord; legs: Warchief Kilt; feet: Prefect's Boots; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 808, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 50320215012101230-50000000000000000-05030000000000000)

Set DPS (verified): 204.6. Weights run: 12.2s. Verify run: 168.3s. 1047 eligible items had no known source.

Race sweep: human 204.5 ± 0.2, dwarf 199.4 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=1.548 ± 0.016, spirit=0.797 ± 0.013, mp5=4.226 ± 0.014, crit=0.333 ± 0.016 per rating point (14 rating = 1%, 4.667 per %), spell_haste=not significant (0.277 ± 0.106)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcatcher Halo (10630) | Sunken Temple: Atal'ai Warrior [dungeon] | 56.7 healing_power points (7.46 DPS) | yes | Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Papal Fez (9431, -0.16 DPS) [dungeon]; Gemburst Circlet (10751, -0.40 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.4 healing_power points (4.26 DPS) | yes | Gemshard Heart (17707, -0.77 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -0.99 DPS) [world_drop]; Mindburst Medallion (11196, -1.72 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 59.9 healing_power points (7.88 DPS) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, -0.53 DPS) [vendor]; Prefect's Pauldrons (250579, -2.03 DPS) [crafted]; Ironfeather Shoulders (15067, -3.39 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Expeditionary's Cape (271921, -0.15 DPS) [rep]; Spritecaster Cape (11623, -0.34 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.71 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 72.4 healing_power points (9.52 DPS) | yes | Knight's Imbued Armor (220813, -0.84 DPS) [vendor]; Stormcloth Vest (10020, -1.05 DPS) [crafted]; Robes of Insight (940, -1.25 DPS, sim-verified) [world_drop] |
| wrist | Prefect's Wristguards (250584) | Blacksmithing [crafted] | 35.4 healing_power points (4.65 DPS) | yes | Nethergeld Cuffs (254061, -0.07 DPS) [crafted]; Mender's Leather Bracers (252543, -0.10 DPS) [crafted]; Mender's Mail Bracers (252573, -0.10 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Prefect's Gloves (250554, -0.01 DPS) [crafted]; Gilded Gloves (254095, -0.16 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 48.6 healing_power points (6.39 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.15 DPS) [dungeon]; Prefect's Waistguard (250574, -0.41 DPS) [crafted] |
| legs | Knight's Imbued Leggings (220809) | PvP rank 6 · Knight · Alliance [vendor] | 66.0 healing_power points (8.68 DPS) | yes | Senior Designer's Pantaloons (11841, -1.24 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Kilt (9474, -1.50 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.95 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 45.0 healing_power points (5.92 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.01 DPS) [crafted]; Prefect's Boots (250549, -0.41 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 30.5 healing_power points (4.01 DPS) | yes | Chivalrous Signet (20505, -0.21 DPS) [quest]; Runemender's Seal (273031, -0.87 DPS) [world]; Cyclopean Band (11824, -0.98 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 28.9 healing_power points (3.80 DPS) | yes | Runemender's Seal (273031, -0.66 DPS) [world]; Cyclopean Band (11824, -0.77 DPS) [dungeon]; Chivalrous Signet (20505, -0.78 DPS, sim-verified) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.41 DPS) [world_drop]; Sanctified Orb (20512, -0.81 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.73 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanctified Orb (20512, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.21 DPS) [quest]; Uther's Strength (11302, -0.89 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -4.29 DPS) [dungeon]; Inventor's Focal Sword (17719, -4.70 DPS) [dungeon]; Death Speaker Scepter (2816, -8.34 DPS, sim-verified) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Twisting Essence Jar (249456, -2.63 DPS) [crafted]; Beacon of Hope (9393, -2.66 DPS) [dungeon]; Orb of Lorica (11262, -3.54 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** head: Soulcatcher Halo; shoulder: Mantle of Lost Hope; back: Caretaker's Cape; chest: Embrace of the Wind Serpent; wrist: Prefect's Wristguards; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1047, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 275.2. Weights run: 13.5s. Verify run: 775.9s. 2436 eligible items had no known source.

Race sweep: human 275.1 ± 0.3, dwarf 266.8 ± 0.3 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.374, intellect=2.874 ± 0.064, spirit=1.982 ± 0.031, mp5=11.113 ± 0.062, crit=1.670 ± 0.085 per rating point (14 rating = 1%, 23.381 per %), spell_haste=not significant (3.153 ± 1.181)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 152.5 healing_power points (16.94 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, +0.00 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, +0.00 DPS) [vendor]; Crown of the Penitent (13216, -8.40 DPS, sim-verified) [quest] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 89.6 healing_power points (9.95 DPS) | yes | Tooth of Gnarr (13141, -0.65 DPS, sim-verified) [dungeon]; Stormpike Sage's Pendant (19098, -1.82 DPS) [rep]; Lady Maye's Pendant (14558, -2.66 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Epaulets (20061, +0.00 DPS) [rep]; Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Mantle of Lost Hope (22234, -7.48 DPS, sim-verified) [dungeon] |
| back | Gracious Cape (18743) | Stratholme: Ghostly Citizen [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Drape of Recovery (272413, -2.04 DPS) [vendor]; Expeditionary's Cape (271907, -3.99 DPS) [rep]; Shroud of the Exile (15421, -4.35 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Breastplate of Salvation (250601, -6.02 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brazecore Armguards (13179, +0.00 DPS) [dungeon]; Morlune's Bracer (18741, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, -5.96 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Gauntlets (13244, +0.00 DPS) [dungeon]; Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Theramore Handwraps (271945, -7.33 DPS, sim-verified) [rep] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Grimgore Noose (13403, -0.42 DPS) [dungeon]; Feralsurge Girdle (18104, -7.45 DPS, sim-verified) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Legguards (227150, +0.00 DPS) [vendor]; Leggings of Arcana (12756, -1.55 DPS) [quest]; Ghostloom Leggings (14545, -5.06 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -8.44 DPS, sim-verified) [crafted] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -3.07 DPS) [dungeon]; Lorekeeper's Ring (19522, -3.58 DPS) [rep]; Naglering (11669, -13.55 DPS, sim-verified) [dungeon] |
| finger2 | Band of Rumination (18103) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.61 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19522, -2.28 DPS) [rep]; Seal of Rivendare (13345, -2.45 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -3.70 DPS) [rep]; Ankh of Life (1713, -7.24 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, -1.04 DPS, sim-verified) [rep]; Ankh of Life (1713, -4.77 DPS) [world_drop] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -10.69 DPS) [world_drop]; The Hammer of Grace (11923, -13.27 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Gracious Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Band of Rumination; trinket1: Burst of Knowledge; trinket2: Shard of the Splithooves; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2436, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 719.6. Weights run: 7.8s. Verify run: 357.9s. 2436 eligible items had no known source.

Race sweep: human 719.4 ± 0.7, dwarf 706.1 ± 0.7 (heal_score); the guide race human is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.384, intellect=1.017 ± 0.045, spirit=0.357 ± 0.041, mp5=1.909 ± 0.053, crit=0.962 ± 0.058 per rating point (14 rating = 1%, 13.475 per %), spell_haste=not significant (3.068 ± 0.922)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 96.3 healing_power points (58.19 DPS) | yes | Field Marshal's Lamellar Helmet (231640, -5.88 DPS) [vendor]; Lieutenant Commander's Lamellar Helmet (227149, -12.35 DPS) [vendor]; Mooncloth Circlet (14140, -17.78 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 39.9 healing_power points (24.10 DPS) | yes | Tooth of Gnarr (13141, -4.81 DPS) [dungeon]; Animated Chain Necklace (18723, -4.88 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -5.77 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Theramore Mantle (271957, -25.91 DPS, sim-verified) [rep] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -3.71 DPS, sim-verified) [rep]; Gracious Cape (18743, -7.89 DPS) [dungeon]; Shroud of the Exile (15421, -8.44 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Breastplate of Salvation (250601, -23.24 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Funeral Cuffs (12626, +0.00 DPS) [dungeon]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Loomguard Armbraces (13969, -22.65 DPS, sim-verified) [dungeon] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Hands of the Exalted Herald (12554, -1.72 DPS, sim-verified) [dungeon]; Prefect's Gloves (250554, -3.19 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Belt of the Ordained (18702, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -24.57 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legguards (231639, -5.34 DPS) [vendor]; Ironfeather Leggings (252486, -6.56 DPS, sim-verified) [crafted]; Knight-Captain's Lamellar Legguards (227150, -11.97 DPS) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Marshal's Lamellar Greaves (231646, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, -25.79 DPS, sim-verified) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -1.72 DPS) [dungeon]; Blessed Band of Light (272407, -5.59 DPS) [vendor]; Naglering (11669, -16.85 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -0.97 DPS) [vendor]; Maiden's Circle (13001, -4.95 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+11.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, -29.56 DPS) [vendor]; Taskmaster Axe (5194, -52.47 DPS, sim-verified) [dungeon]; Grand Marshal's Demolisher (234568, -71.11 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Devout Robe; wrist: Devout Bracers; hands: Theramore Handwraps; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Fordring's Seal; finger2: Rosewine Circle; trinket1: Dormant Heart of the Mountain; trinket2: Burst of Knowledge; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2436, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 49.3. Weights run: 5.6s. Verify run: 41.9s. 299 eligible items had no known source.

Race sweep: undead 49.3 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=0.949 ± 0.005, spirit=0.708 ± 0.004, mp5=3.245 ± 0.008, crit=0.107 ± 0.004 per rating point (14 rating = 1%, 1.505 per %), spell_haste=4.520 ± 0.121

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (3.55 DPS) | yes | Pristine Circlet (253949, -0.55 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.68 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.69 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.1 healing_power points (0.36 DPS) | yes | Scholarly Pendant (277203, -0.18 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.5 healing_power points (2.29 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.80 DPS) [crafted]; Serpent's Shoulders (5404, -1.45 DPS) [dungeon]; Slime-encrusted Pads (6461, -1.58 DPS, sim-verified) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.88 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.94 DPS) [crafted]; Sanguine Cape (14376, -1.12 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 21.9 healing_power points (3.70 DPS) | yes | Acolyte's Chain Shirt (250491, -0.07 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.34 DPS) [crafted]; Corsair's Overshirt (5202, -0.88 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.5 healing_power points (2.45 DPS) | yes | Owl Bracers (4796, -1.65 DPS) [vendor]; Drakewing Bands (12999, -1.73 DPS) [world_drop]; Garrison Cuffs (270003, -1.75 DPS, sim-verified) [quest] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.7 healing_power points (2.49 DPS) | yes | Wisdom's Leather Gloves (252499, -0.16 DPS) [crafted]; Pristine Gloves (253913, -0.28 DPS, sim-verified) [crafted]; Stormrider's Leather Gloves (252498, -1.01 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.8 healing_power points (2.50 DPS) | yes | Acolyte's Chain Belt (250516, -0.35 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.51 DPS) [crafted]; Novice Ardent's Sash (253887, -0.65 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | sim-verified (49.3 DPS) | yes | Wisdom's Leather Pants (252503, -0.09 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.54 DPS) [crafted] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.7 healing_power points (2.66 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.50 DPS) [crafted]; Black Whelp Slippers (252424, -0.50 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 11.5 healing_power points (1.94 DPS) | yes | Loop of Sacrifice (281673, -1.14 DPS) [quest]; Band of Purification (12996, -1.22 DPS) [world_drop]; Volcanic Rock Ring (12053, -1.46 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.7 healing_power points (0.96 DPS) | yes | Loop of Sacrifice (281673, -0.19 DPS, sim-verified) [quest]; Band of Purification (12996, -0.24 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.48 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | 50.7 healing_power points (8.56 DPS) | yes | The Stitcher (279874, -2.54 DPS, sim-verified) [quest]; Pathfinder's Clearway (277288, -5.54 DPS) [quest]; Sword of the Fallen (286977, -7.12 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 299, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (undead, 50320015000100000-40000000000000000-00000000000000000)

Set DPS (verified): 110.1. Weights run: 6.3s. Verify run: 42.6s. 524 eligible items had no known source.

Race sweep: undead 110.1 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.031, intellect=1.586 ± 0.008, spirit=0.815 ± 0.004, mp5=4.340 ± 0.022, crit=0.182 ± 0.005 per rating point (14 rating = 1%, 2.551 per %), spell_haste=not significant (0.048 ± 0.017)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.9 healing_power points (5.47 DPS) | yes | Enduring Cap (3020, -1.05 DPS, sim-verified) [world_drop]; Acolyte's Chain Helm (250501, -1.86 DPS) [crafted]; Nightsky Cowl (4039, -2.13 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.6 healing_power points (1.39 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.01 DPS) [world]; Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.24 DPS) [vendor] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 23.5 healing_power points (3.39 DPS) | yes | Bloodmage Mantle (7684, -0.25 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.60 DPS) [world_drop]; Tanned Shoulderpads (270023, -0.71 DPS) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, -0.10 DPS, sim-verified) [vendor]; Glowing Thresher Cape (6901, -0.40 DPS) [dungeon]; Resilient Cape (14400, -0.97 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 37.5 healing_power points (5.41 DPS) | yes | Pristine Gown (253961, -0.45 DPS) [crafted]; Beguiler Robes (7728, -0.47 DPS, sim-verified) [dungeon]; Acolyte's Silvered Chain Shirt (250521, -0.63 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.0 healing_power points (3.18 DPS) | yes | Mindthrust Bracers (1974, -0.15 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.35 DPS) [world_drop]; Nightsky Wristbands (6407, -1.45 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.8 healing_power points (2.85 DPS) | yes | Acolyte's Gloves (250511, -0.22 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.30 DPS) [world_drop]; Hotshot Pilot's Gloves (9491, -0.43 DPS) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 33.4 healing_power points (4.82 DPS) | yes | Prefect's Belt (250559, -0.16 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -1.54 DPS) [crafted]; Defiler's Lizardhide Girdle (20172, -2.08 DPS) [rep] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 37.8 healing_power points (5.46 DPS) | yes | Acolyte's Silvered Chain Leggings (250526, -0.02 DPS) [crafted]; Pristine Leggings (253987, -0.09 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.73 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.4 healing_power points (4.96 DPS) | yes | Acolyte's Boots (250506, -2.23 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.23 DPS) [crafted]; Acidic Walkers (9454, -2.79 DPS, sim-verified) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 24.2 healing_power points (3.50 DPS) | yes | Darkspear Signet (272071, -0.99 DPS) [vendor]; Advisor's Ring (19521, -1.23 DPS) [rep]; Snake Hoop (6750, -5.45 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.11 DPS, sim-verified) [quest]; Darkspear Signet (272071, -0.24 DPS) [vendor]; Advisor's Ring (19521, -0.48 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crested Scepter (3414, -4.74 DPS) [dungeon]; The Stitcher (279874, -5.49 DPS) [quest]; Taskmaster Axe (5194, -7.10 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 17.9 healing_power points (2.58 DPS) | yes | Orb of Mystic Insight (249394, -0.20 DPS) [crafted]; Heart of Agamaggan (6694, -0.42 DPS) [dungeon]; Seedcloud Buckler (6630, -0.72 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 524, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 50320115000100000-50000000000000000-05030000000000000)

Set DPS (verified): 191.9. Weights run: 5.1s. Verify run: 26.6s. 792 eligible items had no known source.

Race sweep: undead 191.9 ± 0.1 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.091, intellect=1.012 ± 0.019, spirit=0.555 ± 0.011, mp5=3.146 ± 0.072, crit=0.316 ± 0.009 per rating point (14 rating = 1%, 4.418 per %), spell_haste=1.327 ± 0.098

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 44.2 healing_power points (12.53 DPS) | yes | Whitemane's Chapeau (7720, -0.65 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.23 DPS) [world_drop]; Miner's Hat of the Deep (9429, -4.25 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.1 healing_power points (7.96 DPS) | yes | Triune Amulet (7722, -2.58 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -4.85 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -5.41 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.2 healing_power points (9.96 DPS) | yes | Earthen Silk Shoulders (254033, -1.38 DPS, sim-verified) [crafted]; Sheepshear Mantle (13115, -4.43 DPS) [world_drop]; Green Silken Shoulders (7057, -4.54 DPS) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272077, -2.10 DPS) [vendor]; Glowing Thresher Cape (6901, -2.64 DPS) [dungeon]; Long Silken Cloak (4326, -2.75 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.4 healing_power points (17.69 DPS) | yes | Deathchill Armor (10764, -6.43 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -7.81 DPS) [crafted]; Death's Head Vestment (10581, -8.25 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.4 healing_power points (5.22 DPS) | yes | Mindthrust Bracers (1974, -1.59 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -1.60 DPS) [crafted]; Radiant Silver Bracers (4545, -1.80 DPS) [quest] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Gloves (252530, -0.14 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.28 DPS) [crafted]; Gloves of the Greatfather (17721, -0.87 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 33.4 healing_power points (9.47 DPS) | yes | Prefect's Belt (250559, -0.08 DPS, sim-verified) [crafted]; Mender's Leather Belt (252523, -1.42 DPS) [crafted]; Deathmage Sash (10771, -3.19 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 37.8 healing_power points (10.72 DPS) | yes | Stoneweaver Leggings (9407, -0.77 DPS, sim-verified) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -1.35 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.66 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 37.1 healing_power points (10.52 DPS) | yes | Mender's Mail Boots (252565, -0.61 DPS) [crafted]; Mender's Leather Shoes (252533, -0.82 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -2.21 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.0 healing_power points (5.37 DPS) | yes | Snake Hoop (6750, -0.56 DPS) [quest]; Ring of Power Regulation (273030, -1.00 DPS) [dungeon]; Darkspear Signet (272070, -2.29 DPS, sim-verified) [vendor] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.42 DPS) [quest]; Ring of Power Regulation (273030, -0.85 DPS) [dungeon]; Darkspear Signet (272070, -1.48 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -5.15 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -13.88 DPS) [dungeon]; Skullbreaker (17039, -17.96 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Prophetic Cane (6803, -0.47 DPS, sim-verified) [quest]; Forcestone Buckler (17508, -0.50 DPS) [quest]; Battered Viking Shield (9403, -1.61 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Gilded Handwraps; waist: Gilded Cord; legs: Warchief Kilt; feet: Prefect's Boots; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 792, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 50320215012101230-50000000000000000-05030000000000000)

Set DPS (verified): 199.9. Weights run: 12.2s. Verify run: 162.5s. 1047 eligible items had no known source.

Race sweep: undead 199.8 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.207, intellect=1.535 ± 0.025, spirit=0.937 ± 0.014, mp5=5.536 ± 0.027, crit=0.644 ± 0.032 per rating point (14 rating = 1%, 9.017 per %), spell_haste=2.055 ± 0.384

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.5 healing_power points (12.53 DPS) | yes | Soulcatcher Halo (10630, -0.03 DPS) [dungeon]; Gemburst Circlet (10751, -0.13 DPS) [quest]; Soothsayer's Headdress (17740, -0.67 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.3 healing_power points (7.16 DPS) | yes | Gemshard Heart (17707, -0.75 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.56 DPS) [world_drop]; Mindburst Medallion (11196, -2.75 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 65.0 healing_power points (14.42 DPS) | yes | Ironfeather Shoulders (15067, -3.05 DPS, sim-verified) [crafted]; Lead Surveyor's Mantle (11842, -4.46 DPS) [dungeon]; Mender's Leather Shoulders (252538, -4.58 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19527, -0.27 DPS) [rep]; Darkspear Raider's Cloak (272076, -0.59 DPS) [vendor]; Deep Woodlands Cloak (19121, -0.67 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 76.5 healing_power points (16.95 DPS) | yes | Ghostweave Vest (14141, -1.22 DPS, sim-verified) [crafted]; Robes of Insight (940, -1.47 DPS) [world_drop]; Stormcloth Vest (10020, -2.43 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 35.4 healing_power points (7.85 DPS) | yes | Mender's Leather Bracers (252543, -0.01 DPS) [crafted]; Mender's Mail Bracers (252573, -0.01 DPS) [crafted]; Prefect's Wristguards (250584, -0.03 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Prefect's Gloves (250554, -0.30 DPS) [crafted]; Gilded Gloves (254095, -0.31 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 48.4 healing_power points (10.73 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.28 DPS) [dungeon]; Prefect's Waistguard (250574, -0.68 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 64.6 healing_power points (14.32 DPS) | yes | Jinxed Hoodoo Kilt (9474, -1.88 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -2.46 DPS) [dungeon]; Warchief Kilt (7760, -4.05 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 44.9 healing_power points (9.95 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.02 DPS) [crafted]; Coldstone Slippers (18697, -0.28 DPS) [dungeon] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 38.8 healing_power points (8.59 DPS) | yes | Brainlash (6440, -1.71 DPS) [dungeon]; Runemender's Seal (273031, -3.10 DPS) [world]; Cyclopean Band (11824, -3.38 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 34.1 healing_power points (7.57 DPS) | yes | Brainlash (6440, -0.97 DPS, sim-verified) [dungeon]; Runemender's Seal (273031, -2.08 DPS) [world]; Cyclopean Band (11824, -2.36 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.15 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.06 DPS) [quest]; Alchemist's Stone (13503, -5.31 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.42 DPS) [quest]; Uther's Strength (11302, -0.96 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -1.66 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -7.42 DPS) [dungeon]; Arbiter's Blade (11784, -7.50 DPS) [dungeon]; Inventor's Focal Sword (17719, -7.75 DPS, sim-verified) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prophetic Cane (6803, -4.49 DPS, sim-verified) [quest]; Twisting Essence Jar (249456, -5.32 DPS) [crafted]; Beacon of Hope (9393, -5.49 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Mender's Leather Boots; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1047, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 267.8. Weights run: 13.5s. Verify run: 780.5s. 2461 eligible items had no known source.

Race sweep: undead 267.7 ± 0.3 (heal_score); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.038, intellect=1.914 ± 0.027, spirit=1.363 ± 0.016, mp5=4.918 ± 0.020, crit=0.564 ± 0.028 per rating point (14 rating = 1%, 7.897 per %), spell_haste=not significant (0.059 ± 0.107)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 126.5 healing_power points (13.72 DPS) | yes | Devout Crown (16693, -4.35 DPS) [dungeon]; Whitesoul Helm (12633, -4.59 DPS) [crafted]; Black Dragonscale Helm (252605, -4.98 DPS) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 66.1 healing_power points (7.16 DPS) | yes | Tooth of Gnarr (13141, -1.66 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.03 DPS) [world_drop]; Animated Chain Necklace (18723, -2.55 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Epaulets (20176, +0.00 DPS) [rep]; Mantle of Lost Hope (22234, +0.00 DPS) [dungeon]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gracious Cape (18743, -0.08 DPS) [dungeon]; Shroud of the Exile (15421, -1.08 DPS) [quest]; Darkspear Raider's Cloak (272063, -1.29 DPS) [vendor] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Mooncloth Vest (14138, +0.00 DPS) [crafted]; Breastplate of Salvation (250601, +0.00 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Morlune's Bracer (18741, +0.00 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Raider Handwraps (272097, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Feralsurge Girdle (18104, -0.12 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legplates of the Chromatic Defier (12945, -2.58 DPS) [quest]; Leggings of Arcana (12756, -2.94 DPS) [quest]; Ghostloom Leggings (14545, -3.85 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Firemoss Boots (22275, +0.00 DPS) [dungeon] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.10 DPS) [dungeon]; Fordring's Seal (16058, -0.54 DPS) [quest]; Blessed Band of Light (272407, -1.26 DPS) [vendor] |
| finger2 | Darkspear Signet (272068) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -0.26 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.07 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.60 DPS) [vendor]; Serenity Field (272439, -1.99 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -8.61 DPS) [world_drop]; The Hammer of Grace (11923, -9.10 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Darkspear Signet; trinket1: Burst of Knowledge; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2461, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 705.3. Weights run: 7.8s. Verify run: 361.4s. 2461 eligible items had no known source.

Race sweep: undead 705.6 ± 0.7 (heal_score); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.384, intellect=1.017 ± 0.045, spirit=0.357 ± 0.041, mp5=1.909 ± 0.053, crit=0.962 ± 0.058 per rating point (14 rating = 1%, 13.475 per %), spell_haste=not significant (3.068 ± 0.922)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 96.3 healing_power points (58.19 DPS) | yes | Mooncloth Circlet (14140, -18.60 DPS, sim-verified) [crafted]; Whitesoul Helm (12633, -24.59 DPS) [crafted]; Black Dragonscale Helm (252605, -26.69 DPS) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 39.9 healing_power points (24.10 DPS) | yes | Animated Chain Necklace (18723, -4.32 DPS, sim-verified) [dungeon]; Tooth of Gnarr (13141, -4.81 DPS) [dungeon]; Lady Maye's Pendant (14558, -5.77 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawnbringer Shoulders (12625, +0.00 DPS) [crafted]; Mantle of Lost Hope (22234, +0.00 DPS) [dungeon]; Darkspear Mantle (272107, -26.74 DPS, sim-verified) [vendor] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -5.17 DPS, sim-verified) [rep]; Gracious Cape (18743, -7.89 DPS) [dungeon]; Raincaster Drape (12110, -8.10 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of Insight (940, +0.00 DPS) [world_drop]; Stormcloth Vest (10020, +0.00 DPS) [crafted]; Breastplate of Salvation (250601, -24.97 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Funeral Cuffs (12626, +0.00 DPS) [dungeon]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Loomguard Armbraces (13969, -23.06 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -2.89 DPS, sim-verified) [dungeon]; Prefect's Gloves (250554, -3.19 DPS) [crafted]; Mender's Leather Gauntlets (252551, -4.90 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of the Ordained (18702, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Sash of Mercy (14553, -25.06 DPS, sim-verified) [world_drop] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironfeather Leggings (252486, -6.59 DPS, sim-verified) [crafted]; Legplates of the Chromatic Defier (12945, -16.08 DPS) [quest]; Ghoul Skin Leggings (18682, -23.15 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Broodwatcher's Treaders (275621, +0.00 DPS) [crafted]; Swarmtender's Footpads (275609, -28.27 DPS, sim-verified) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -1.72 DPS) [dungeon]; Blessed Band of Light (272407, -5.59 DPS) [vendor]; Naglering (11669, -19.58 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -0.97 DPS) [vendor]; Maiden's Circle (13001, -4.95 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+13.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Eye of the Beast (13968, -1.90 DPS, sim-verified) [quest] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, -29.56 DPS) [vendor]; Hand of Edward the Odd (2243, -45.06 DPS, sim-verified) [world_drop]; High Warlord's Destroyer (234546, -71.11 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Devout Robe; wrist: Devout Bracers; hands: Raider Handwraps; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Fordring's Seal; finger2: Rosewine Circle; trinket1: Dormant Heart of the Mountain; trinket2: Burst of Knowledge; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2461, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

