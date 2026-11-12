# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 51.3. Weights run: 5.2s. Verify run: 37.3s. 312 eligible items had no known source.

Race sweep: human 51.3 ± 0.0, dwarf 49.3 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=1.023 ± 0.008, spirit=0.775 ± 0.005, mp5=3.250 ± 0.009, crit=0.194 ± 0.009 per rating point (14 rating = 1%, 2.712 per %), spell_haste=3.246 ± 0.103

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (3.61 DPS) | yes | Pristine Circlet (253949, -0.57 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.69 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.72 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.5 healing_power points (0.27 DPS) | yes | Tarnished Locket (279870, -0.13 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.2 healing_power points (2.44 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.88 DPS) [crafted]; Serpent's Shoulders (5404, -1.51 DPS) [dungeon]; Slime-encrusted Pads (6461, -1.53 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.80 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.94 DPS) [crafted]; Sanguine Cape (14376, -1.11 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 22.4 healing_power points (3.86 DPS) | yes | Acolyte's Chain Shirt (250491, -0.23 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.34 DPS) [crafted]; Corsair's Overshirt (5202, -0.90 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.9 healing_power points (2.56 DPS) | yes | Drakewing Bands (12999, -1.76 DPS) [world_drop]; Bright Bracers (3647, -1.85 DPS) [world_drop]; Owl Bracers (4796, -2.62 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 15.1 healing_power points (2.60 DPS) | yes | Wisdom's Leather Gloves (252499, -0.17 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.18 DPS) [crafted]; Stormrider's Leather Gloves (252498, -1.04 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.1 healing_power points (2.60 DPS) | yes | Acolyte's Chain Belt (250516, -0.37 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.52 DPS) [crafted]; Novice Ardent's Sash (253887, -0.64 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Pants (252503, -0.07 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.16 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.54 DPS) [crafted] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 16.1 healing_power points (2.77 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.52 DPS) [crafted]; Black Whelp Slippers (252424, -0.52 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 11.5 healing_power points (1.98 DPS) | yes | Lavishly Jeweled Ring (1156, -0.92 DPS) [dungeon]; Band of Purification (12996, -1.18 DPS) [world_drop]; Volcanic Rock Ring (12053, -1.45 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 healing_power points (1.21 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS, sim-verified) [dungeon]; Band of Purification (12996, -0.41 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.68 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monastic Hammer (270005, -5.49 DPS, sim-verified) [quest]; Verigan's Fist (6953, -6.42 DPS) [quest]; Sword of the Fallen (286977, -7.22 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 312, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 30 (human, 50320015000100000-40000000000000000-00000000000000000)

Set DPS (verified): 137.1. Weights run: 4.5s. Verify run: 24.6s. 537 eligible items had no known source.

Race sweep: human 137.1 ± 0.1, dwarf 132.1 ± 0.1 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=1.891 ± 0.016, spirit=1.419 ± 0.012, mp5=4.787 ± 0.023, crit=0.598 ± 0.018 per rating point (14 rating = 1%, 8.373 per %), spell_haste=-0.471 ± 0.113

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 41.5 healing_power points (6.29 DPS) | yes | Nightsky Cowl (4039, -1.78 DPS) [world_drop]; Dokebi Hat (14584, -2.06 DPS) [world_drop]; Enduring Cap (3020, -7.00 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.2 healing_power points (2.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.21 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.29 DPS) [world]; Seal of the Expedition (271920, -0.57 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mantle of Honor (3560, -0.43 DPS) [quest]; Death Speaker Mantle (6685, -0.52 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.57 DPS) [world_drop] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Repairman's Cape (9605, -0.11 DPS) [quest]; Caretaker's Cape (19533, -0.13 DPS) [rep]; Expeditionary's Cape (271923, -2.20 DPS, sim-verified) [rep] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.29 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.44 DPS) [crafted]; Beguiler Robes (7728, -2.05 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Wristbands (6407, -1.18 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.25 DPS) [world_drop]; Mindthrust Bracers (1974, -1.70 DPS, sim-verified) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 26.5 healing_power points (4.01 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.65 DPS) [dungeon]; Truefaith Gloves (7049, -0.88 DPS) [crafted]; Shilly Mitts (9609, -1.02 DPS) [quest] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Belt (252523, -0.49 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -1.24 DPS) [crafted]; Silver-lined Belt (13011, -1.90 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Leggings (252519, -0.17 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.33 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.80 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.9 healing_power points (5.89 DPS) | yes | Acolyte's Boots (250506, -2.80 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.80 DPS) [crafted]; Acidic Walkers (9454, -3.39 DPS, sim-verified) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 30.3 healing_power points (4.59 DPS) | yes | Ring of Power Regulation (273030, -1.50 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.72 DPS) [dungeon]; Lorekeeper's Ring (19525, -2.08 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 29.2 healing_power points (4.42 DPS) | yes | Ring of Power Regulation (273030, -0.36 DPS, sim-verified) [dungeon]; Electrocutioner Lagnut (9447, -1.55 DPS) [dungeon]; Lorekeeper's Ring (19525, -1.91 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crested Scepter (3414, -4.88 DPS) [dungeon]; Pysan's Old Greatsword (1975, -5.91 DPS) [dungeon]; Taskmaster Axe (5194, -8.31 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 21.5 healing_power points (3.26 DPS) | yes | Heart of Agamaggan (6694, -0.99 DPS) [dungeon]; Orb of Mistmantle (13031, -1.04 DPS) [world_drop]; Orb of Mystic Insight (249394, -1.27 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Prelacy Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Prefect's Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 537, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 50320115000100000-50000000000000000-05030000000000000)

Set DPS (verified): 219.3. Weights run: 4.9s. Verify run: 19.8s. 808 eligible items had no known source.

Race sweep: human 219.2 ± 0.1, dwarf 212.6 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.055, intellect=1.243 ± 0.025, spirit=1.236 ± 0.024, mp5=3.229 ± 0.063, crit=0.793 ± 0.024 per rating point (14 rating = 1%, 11.098 per %), spell_haste=1.934 ± 0.181

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 54.3 healing_power points (13.98 DPS) | yes | Whitemane's Chapeau (7720, -1.04 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.57 DPS) [world_drop]; Earthen Silk Hood (254015, -5.13 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.9 healing_power points (7.72 DPS) | yes | Triune Amulet (7722, -3.43 DPS) [dungeon]; Amberglow Talisman (10824, -4.53 DPS) [quest]; Necklace of Calisea (1714, -4.65 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 38.2 healing_power points (9.84 DPS) | yes | Earthen Silk Shoulders (254033, -1.75 DPS, sim-verified) [crafted]; Sheepshear Mantle (13115, -2.80 DPS) [world_drop]; Mantle of Doan (7712, -3.56 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prelacy Cape (7004, -1.49 DPS) [quest]; Glowing Thresher Cape (6901, -1.88 DPS) [dungeon]; Expeditionary's Cape (271922, -3.64 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.9 healing_power points (17.50 DPS) | yes | Dreamweave Vest (10021, -5.51 DPS) [crafted]; Death's Head Vestment (10581, -6.15 DPS) [dungeon]; Deathchill Armor (10764, -8.60 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.7 healing_power points (4.82 DPS) | yes | Earthen Silk Cuffs (254019, -0.65 DPS) [crafted]; Mindthrust Bracers (1974, -0.72 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.80 DPS, sim-verified) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.38 DPS) [crafted]; Prefect's Gauntlet (250569, -0.64 DPS) [crafted]; Dreamweave Gloves (10019, -0.69 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 39.4 healing_power points (10.15 DPS) | yes | Mender's Leather Belt (252523, -0.95 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -2.23 DPS) [crafted]; Deathmage Sash (10771, -3.53 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 48.8 healing_power points (12.57 DPS) | yes | Stoneweaver Leggings (9407, -0.94 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -2.68 DPS) [crafted]; Pristine Leggings (253987, -3.06 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 41.4 healing_power points (10.66 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Prefect's Boots (250549, -0.56 DPS) [crafted]; Gilded Slippers (254001, -1.99 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 25.1 healing_power points (6.47 DPS) | yes | Electrocutioner Lagnut (9447, -1.92 DPS) [dungeon]; Philanthropist's Ring (281635, -2.29 DPS) [quest]; Snake Hoop (6750, -3.99 DPS, sim-verified) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.16 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.27 DPS) [dungeon]; Philanthropist's Ring (281635, -0.64 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -7.31 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -12.63 DPS) [dungeon]; Bookmaker's Scepter (4122, -17.12 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Lorica (11262, -0.27 DPS, sim-verified) [quest]; Battered Viking Shield (9403, -0.41 DPS) [dungeon]; Orb of Souls (249395, -1.68 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 808, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 50320215012101230-50000000000000000-05030000000000000)

Set DPS (verified): 257.3. Weights run: 9.4s. Verify run: 113.5s. 1047 eligible items had no known source.

Race sweep: human 257.0 ± 0.3, dwarf 243.1 ± 0.3 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.331, intellect=1.772 ± 0.039, spirit=2.085 ± 0.031, mp5=6.243 ± 0.045, crit=1.594 ± 0.057 per rating point (14 rating = 1%, 22.316 per %), spell_haste=not significant (0.945 ± 0.674)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 80.7 healing_power points (18.20 DPS) | yes | Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Engineer's Guild Headpiece (9534, -1.30 DPS) [quest]; Eye of Theradras (17715, -2.03 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 34.2 healing_power points (7.71 DPS) | yes | Lei of Lilies (1315, -0.65 DPS) [world_drop]; Gemshard Heart (17707, -0.73 DPS) [dungeon]; Horizon Choker (13085, -2.01 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 70.5 healing_power points (15.90 DPS) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, -0.90 DPS) [vendor]; Ironfeather Shoulders (15067, -1.99 DPS, sim-verified) [crafted]; Lead Surveyor's Mantle (11842, -2.68 DPS) [dungeon] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.60 DPS, sim-verified) [dungeon]; Expeditionary's Cape (271921, -1.14 DPS) [rep]; Caretaker's Cape (19531, -1.30 DPS) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 113.3 healing_power points (25.56 DPS) | yes | Ghostweave Vest (14141, -4.78 DPS, sim-verified) [crafted]; Feathered Breastplate (8349, -5.98 DPS) [crafted]; Knight's Imbued Armor (220813, -8.18 DPS) [vendor] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 43.9 healing_power points (9.91 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Slimescale Bracers (10632, -0.16 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.24 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.52 DPS) [crafted]; Jumanza Grips (18083, -2.33 DPS) [dungeon] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Earthenweave Cord (254077, -0.52 DPS) [crafted]; Dawnspire Cord (12466, -0.84 DPS, sim-verified) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 86.9 healing_power points (19.62 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.92 DPS) [dungeon]; Knight's Imbued Leggings (220809, -2.24 DPS) [vendor]; Jinxed Hoodoo Kilt (9474, -2.42 DPS, sim-verified) [dungeon] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Shoes (252533, -0.86 DPS) [crafted]; Mender's Mail Boots (252565, -0.86 DPS) [crafted]; Sandals of the Insurgent (13111, -1.36 DPS, sim-verified) [world_drop] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 39.1 healing_power points (8.83 DPS) | yes | Eye of Adaegus (5266, -0.79 DPS) [world_drop]; Runemender's Seal (273031, -1.11 DPS) [world]; Snake Hoop (6750, -1.39 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 37.0 healing_power points (8.34 DPS) | yes | Runemender's Seal (273031, -0.62 DPS) [world]; Snake Hoop (6750, -0.90 DPS) [quest]; Eye of Adaegus (5266, -1.69 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -0.94 DPS) [quest]; Uther's Strength (11302, -1.59 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -1.88 DPS) [quest] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.29 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -2.34 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.28 DPS) [quest] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inventor's Focal Sword (17719, -9.42 DPS, sim-verified) [dungeon]; Arbiter's Blade (11784, -9.44 DPS) [dungeon]; Death Speaker Scepter (2816, -9.63 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of Lorica (11262, -3.80 DPS, sim-verified) [quest]; Twisting Essence Jar (249456, -4.51 DPS) [crafted]; Mountainside Buckler (13082, -4.92 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; shoulder: Mantle of Lost Hope; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Vinerot Sandals; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket2: Enriched Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1047, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 349.8. Weights run: 10.6s. Verify run: 538.2s. 2436 eligible items had no known source.

Race sweep: human 349.5 ± 0.5, dwarf 331.1 ± 0.5 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.215, intellect=1.692 ± 0.052, spirit=2.454 ± 0.042, mp5=7.646 ± 0.058, crit=2.112 ± 0.075 per rating point (14 rating = 1%, 29.570 per %), spell_haste=not significant (1.163 ± 0.754)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Lamellar Faceguard (16474, +0.00 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, +0.00 DPS) [vendor]; Living Crown (252561, -0.96 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 73.0 healing_power points (15.62 DPS) | yes | Tooth of Gnarr (13141, -3.75 DPS, sim-verified) [dungeon]; Animated Chain Necklace (18723, -4.88 DPS) [dungeon]; Dragon's Eye (10829, -4.96 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of Lost Hope (22234, +0.00 DPS) [dungeon]; Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Shroud of the Nathrezim (18720, -28.96 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Gracious Cape (18743, -1.26 DPS, sim-verified) [dungeon]; Shroud of the Exile (15421, -3.31 DPS) [quest]; Featherskin Cape (10843, -3.37 DPS) [world] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 144.7 healing_power points (30.95 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, -2.27 DPS) [pvp]; Embrace of the Wind Serpent (12462, -5.34 DPS, sim-verified) [world]; Ghostweave Vest (14141, -5.44 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Morlune's Bracer (18741, +0.00 DPS) [dungeon]; Bleak Howler Armguards (13208, -27.74 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Theramore Handwraps (271945, -0.85 DPS) [rep]; Hands of the Exalted Herald (12554, -28.49 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feralsurge Girdle (18104, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Sash of Mercy (14553, -26.19 DPS, sim-verified) [world_drop] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legplates of the Chromatic Defier (12945, -4.96 DPS) [quest]; Knight-Captain's Lamellar Legguards (227150, -5.60 DPS) [vendor]; Leggings of Arcana (12756, -5.96 DPS, sim-verified) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -28.90 DPS, sim-verified) [crafted] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -2.63 DPS) [dungeon]; Blessed Band of Light (272407, -4.12 DPS) [vendor]; Naglering (11669, -15.58 DPS, sim-verified) [dungeon] |
| finger2 | Band of Rumination (18103) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -2.19 DPS) [dungeon]; Blessed Band of Light (272407, -3.67 DPS) [vendor]; The Postmaster's Seal (13392, -4.34 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+12.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -4.91 DPS) [rep]; Ankh of Life (1713, -6.79 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -7.84 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, -1.64 DPS) [rep]; Ankh of Life (1713, -3.51 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.56 DPS) [quest] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -2.57 DPS, sim-verified) [world_drop]; The Hammer of Grace (11923, -20.69 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Breastplate of Salvation; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Band of Rumination; trinket1: Burst of Knowledge; trinket2: Shard of the Splithooves; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2436, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 797.6. Weights run: 11.0s. Verify run: 219.7s. 2436 eligible items had no known source.

Race sweep: human 797.4 ± 0.6, dwarf 778.6 ± 0.6 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.179, intellect=1.436 ± 0.060, spirit=1.404 ± 0.058, mp5=3.558 ± 0.063, crit=1.564 ± 0.074 per rating point (14 rating = 1%, 21.893 per %), spell_haste=not significant (-0.241 ± 1.121)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 118.5 healing_power points (30.50 DPS) | yes | Field Marshal's Lamellar Helmet (231640, -2.52 DPS) [vendor]; Field Marshal's Lamellar Faceguard (16474, -6.46 DPS) [vendor]; Whitesoul Helm (12633, -13.50 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 57.9 healing_power points (14.90 DPS) | yes | Tooth of Gnarr (13141, -4.23 DPS) [dungeon]; Animated Chain Necklace (18723, -4.92 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -5.05 DPS) [world_drop] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 63.1 healing_power points (16.24 DPS) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Theramore Mantle (271957, -1.08 DPS) [rep] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -2.06 DPS) [rep]; Shroud of the Exile (15421, -2.33 DPS) [quest]; Gracious Cape (18743, -2.71 DPS, sim-verified) [dungeon] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 106.4 healing_power points (27.40 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Knight-Captain's Lamellar Chestplate (227151, -4.69 DPS) [vendor]; Embrace of the Wind Serpent (12462, -5.10 DPS, sim-verified) [world] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 49.7 healing_power points (12.80 DPS) | yes | Morlune's Bracer (18741, -0.53 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.55 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -1.65 DPS) [rep] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Theramore Handwraps (271945, -0.96 DPS) [rep]; Mender's Leather Gauntlets (252551, -2.56 DPS) [crafted] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 69.9 healing_power points (17.98 DPS) | yes | Wisdom of the Timbermaw (19047, -1.14 DPS, sim-verified) [crafted]; Belt of the Ordained (18702, -4.31 DPS) [dungeon]; Devout Belt (16696, -4.50 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legguards (231639, -4.31 DPS) [vendor]; Sentinel's Silk Leggings (237815, -7.27 DPS) [vendor]; Legplates of the Chromatic Defier (12945, -8.21 DPS, sim-verified) [quest] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 62.4 healing_power points (16.05 DPS) | yes | Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Marshal's Lamellar Boots (16472, -0.38 DPS) [vendor]; Faith Healer's Boots (22247, -1.10 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.06 DPS) [dungeon]; Blessed Band of Light (272407, -1.87 DPS) [vendor]; Naglering (11669, -16.49 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.87 DPS) [vendor]; Band of Rumination (18103, -1.53 DPS) [dungeon]; Rosewine Circle (13178, -3.13 DPS, sim-verified) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+11.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, -16.57 DPS) [vendor]; Grand Marshal's Demolisher (234568, -33.36 DPS) [pvp]; Hand of Edward the Odd (2243, -42.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Shroud of the Nathrezim; back: Drape of Recovery; chest: Breastplate of Salvation; wrist: Loomguard Armbraces; hands: Hands of the Exalted Herald; waist: Sash of Mercy; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Eye of the Beast; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2436, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 50.3. Weights run: 5.2s. Verify run: 36.5s. 299 eligible items had no known source.

Race sweep: undead 50.3 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=1.023 ± 0.008, spirit=0.775 ± 0.005, mp5=3.250 ± 0.009, crit=0.194 ± 0.009 per rating point (14 rating = 1%, 2.712 per %), spell_haste=3.246 ± 0.103

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (3.61 DPS) | yes | Pristine Circlet (253949, -0.53 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.69 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.72 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.3 healing_power points (0.40 DPS) | yes | Scholarly Pendant (277203, -0.13 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.2 healing_power points (2.44 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.88 DPS) [crafted]; Slime-encrusted Pads (6461, -1.45 DPS, sim-verified) [dungeon]; Serpent's Shoulders (5404, -1.51 DPS) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.73 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.94 DPS) [crafted]; Sanguine Cape (14376, -1.11 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 22.4 healing_power points (3.86 DPS) | yes | Acolyte's Chain Shirt (250491, -0.23 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.34 DPS) [crafted]; Corsair's Overshirt (5202, -0.90 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.9 healing_power points (2.56 DPS) | yes | Owl Bracers (4796, -1.68 DPS) [vendor]; Garrison Cuffs (270003, -1.68 DPS, sim-verified) [quest]; Drakewing Bands (12999, -1.76 DPS) [world_drop] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 15.1 healing_power points (2.60 DPS) | yes | Wisdom's Leather Gloves (252499, -0.15 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.18 DPS) [crafted]; Stormrider's Leather Gloves (252498, -1.04 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.1 healing_power points (2.60 DPS) | yes | Acolyte's Chain Belt (250516, -0.36 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.52 DPS) [crafted]; Novice Ardent's Sash (253887, -0.64 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Pants (252503, -0.14 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.16 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.54 DPS) [crafted] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 16.1 healing_power points (2.77 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.52 DPS) [crafted]; Black Whelp Slippers (252424, -0.52 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 11.5 healing_power points (1.98 DPS) | yes | Loop of Sacrifice (281673, -1.10 DPS) [quest]; Band of Purification (12996, -1.18 DPS) [world_drop]; Volcanic Rock Ring (12053, -1.45 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.1 healing_power points (1.06 DPS) | yes | Loop of Sacrifice (281673, -0.15 DPS, sim-verified) [quest]; Band of Purification (12996, -0.26 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | 51.2 healing_power points (8.81 DPS) | yes | The Stitcher (279874, -2.36 DPS, sim-verified) [quest]; Pathfinder's Clearway (277288, -5.70 DPS) [quest]; Sword of the Fallen (286977, -7.22 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 299, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (undead, 50320015000100000-40000000000000000-00000000000000000)

Set DPS (verified): 133.5. Weights run: 4.5s. Verify run: 26.0s. 524 eligible items had no known source.

Race sweep: undead 133.4 ± 0.1 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=1.891 ± 0.016, spirit=1.419 ± 0.012, mp5=4.787 ± 0.023, crit=0.598 ± 0.018 per rating point (14 rating = 1%, 8.373 per %), spell_haste=-0.471 ± 0.113

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 41.5 healing_power points (6.29 DPS) | yes | Nightsky Cowl (4039, -1.78 DPS) [world_drop]; Dokebi Hat (14584, -2.06 DPS) [world_drop]; Enduring Cap (3020, -5.70 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.2 healing_power points (2.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.21 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.29 DPS) [world]; Darkspear Warding Pendant (272075, -0.57 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Mantle (6685, -0.30 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.57 DPS) [world_drop]; Magician's Mantle (12998, -0.61 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.05 DPS) [dungeon]; Resilient Cape (14400, -1.11 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -1.79 DPS, sim-verified) [vendor] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.29 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.44 DPS) [crafted]; Beguiler Robes (7728, -1.09 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindthrust Bracers (1974, -1.00 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.18 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.25 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 26.5 healing_power points (4.01 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.19 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.88 DPS) [crafted]; Acolyte's Gloves (250511, -1.06 DPS) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Belt (252523, -0.17 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -1.24 DPS) [crafted]; Silver-lined Belt (13011, -1.90 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 44.2 healing_power points (6.69 DPS) | yes | Pristine Leggings (253987, -0.28 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.61 DPS) [crafted]; Wisdom's Leather Pants (252503, -1.08 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.9 healing_power points (5.89 DPS) | yes | Acolyte's Boots (250506, -2.80 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.80 DPS) [crafted]; Acidic Walkers (9454, -3.14 DPS, sim-verified) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 30.3 healing_power points (4.59 DPS) | yes | Ring of Power Regulation (273030, -1.50 DPS) [dungeon]; Darkspear Signet (272071, -1.69 DPS) [vendor]; Electrocutioner Lagnut (9447, -1.72 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 29.2 healing_power points (4.42 DPS) | yes | Ring of Power Regulation (273030, -0.77 DPS, sim-verified) [dungeon]; Darkspear Signet (272071, -1.52 DPS) [vendor]; Electrocutioner Lagnut (9447, -1.55 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crested Scepter (3414, -4.88 DPS) [dungeon]; The Stitcher (279874, -5.75 DPS) [quest]; Taskmaster Axe (5194, -8.09 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 21.5 healing_power points (3.26 DPS) | yes | Orb of Mystic Insight (249394, -0.89 DPS, sim-verified) [crafted]; Witch's Finger (16887, -0.97 DPS) [quest]; Heart of Agamaggan (6694, -0.99 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Prefect's Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 524, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 50320115000100000-50000000000000000-05030000000000000)

Set DPS (verified): 215.5. Weights run: 4.9s. Verify run: 19.1s. 792 eligible items had no known source.

Race sweep: undead 215.4 ± 0.1 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.055, intellect=1.243 ± 0.025, spirit=1.236 ± 0.024, mp5=3.229 ± 0.063, crit=0.793 ± 0.024 per rating point (14 rating = 1%, 11.098 per %), spell_haste=1.934 ± 0.181

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 54.3 healing_power points (13.98 DPS) | yes | Whitemane's Chapeau (7720, -1.97 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.57 DPS) [world_drop]; Earthen Silk Hood (254015, -5.13 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.9 healing_power points (7.72 DPS) | yes | Triune Amulet (7722, -3.43 DPS) [dungeon]; Necklace of Calisea (1714, -3.87 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -4.53 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 38.2 healing_power points (9.84 DPS) | yes | Sheepshear Mantle (13115, -2.80 DPS) [world_drop]; Earthen Silk Shoulders (254033, -2.90 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -3.56 DPS) [dungeon] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -1.88 DPS) [dungeon]; Cloak of Blight (6832, -2.09 DPS) [quest]; Darkspear Raider's Cloak (272077, -2.35 DPS, sim-verified) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.9 healing_power points (17.50 DPS) | yes | Dreamweave Vest (10021, -5.51 DPS) [crafted]; Death's Head Vestment (10581, -6.15 DPS) [dungeon]; Deathchill Armor (10764, -6.88 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.7 healing_power points (4.82 DPS) | yes | Earthen Silk Cuffs (254019, -0.65 DPS) [crafted]; Mindthrust Bracers (1974, -0.72 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.24 DPS, sim-verified) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.38 DPS) [crafted]; Dreamweave Gloves (10019, -0.56 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.64 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 39.4 healing_power points (10.15 DPS) | yes | Mender's Leather Belt (252523, -1.35 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -2.23 DPS) [crafted]; Deathmage Sash (10771, -3.53 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 48.8 healing_power points (12.57 DPS) | yes | Stoneweaver Leggings (9407, -0.94 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -2.68 DPS) [crafted]; Pristine Leggings (253987, -3.06 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 41.4 healing_power points (10.66 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Prefect's Boots (250549, -0.56 DPS) [crafted]; Gilded Slippers (254001, -1.99 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 25.1 healing_power points (6.47 DPS) | yes | Darkspear Signet (272070, -1.48 DPS) [vendor]; Advisor's Ring (19520, -1.66 DPS) [rep]; Electrocutioner Lagnut (9447, -1.92 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 23.4 healing_power points (6.02 DPS) | yes | Darkspear Signet (272070, -0.77 DPS, sim-verified) [vendor]; Advisor's Ring (19520, -1.20 DPS) [rep]; Electrocutioner Lagnut (9447, -1.47 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -8.39 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -12.63 DPS) [dungeon]; Skullbreaker (17039, -17.44 DPS) [quest] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beacon of Hope (9393, -0.26 DPS) [dungeon]; Forcestone Buckler (17508, -0.45 DPS, sim-verified) [quest]; Battered Viking Shield (9403, -0.67 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Prophetic Cane

No-known-source sample (15 of 792, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 50320215012101230-50000000000000000-05030000000000000)

Set DPS (verified): 245.7. Weights run: 9.4s. Verify run: 110.1s. 1047 eligible items had no known source.

Race sweep: undead 245.4 ± 0.3 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.331, intellect=1.772 ± 0.039, spirit=2.085 ± 0.031, mp5=6.243 ± 0.045, crit=1.594 ± 0.057 per rating point (14 rating = 1%, 22.316 per %), spell_haste=not significant (0.945 ± 0.674)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 80.7 healing_power points (18.20 DPS) | yes | Eye of Theradras (17715, -1.06 DPS, sim-verified) [dungeon]; Engineer's Guild Headpiece (9534, -1.30 DPS) [quest]; Papal Fez (9431, -2.20 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 34.2 healing_power points (7.71 DPS) | yes | Lei of Lilies (1315, -0.65 DPS) [world_drop]; Gemshard Heart (17707, -0.73 DPS) [dungeon]; Horizon Choker (13085, -2.12 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 70.5 healing_power points (15.90 DPS) | yes | Ironfeather Shoulders (15067, -1.73 DPS, sim-verified) [crafted]; Lead Surveyor's Mantle (11842, -2.68 DPS) [dungeon]; Living Shoulders (15061, -2.79 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272076, -0.43 DPS) [vendor]; Battle Healer's Cloak (19527, -0.59 DPS) [rep]; Featherskin Cape (10843, -0.76 DPS, sim-verified) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 113.3 healing_power points (25.56 DPS) | yes | Ghostweave Vest (14141, -4.13 DPS, sim-verified) [crafted]; Feathered Breastplate (8349, -5.98 DPS) [crafted]; Robes of Insight (940, -8.52 DPS) [world_drop] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 43.9 healing_power points (9.91 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Slimescale Bracers (10632, -0.16 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.24 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.52 DPS) [crafted]; Jumanza Grips (18083, -2.33 DPS) [dungeon] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Earthenweave Cord (254077, -0.52 DPS) [crafted]; Dawnspire Cord (12466, -0.91 DPS, sim-verified) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 86.9 healing_power points (19.62 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.92 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -2.01 DPS, sim-verified) [dungeon]; Rainstrider Leggings (11123, -3.32 DPS) [quest] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Shoes (252533, -0.86 DPS) [crafted]; Mender's Mail Boots (252565, -0.86 DPS) [crafted]; Sandals of the Insurgent (13111, -1.75 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 43.7 healing_power points (9.86 DPS) | yes | Advisor's Ring (19519, -1.52 DPS) [rep]; Eye of Adaegus (5266, -1.82 DPS) [world_drop]; Runemender's Seal (273031, -2.14 DPS) [world] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 39.1 healing_power points (8.83 DPS) | yes | Advisor's Ring (19519, -0.49 DPS) [rep]; Eye of Adaegus (5266, -0.79 DPS) [world_drop]; Runemender's Seal (273031, -1.11 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -0.94 DPS) [quest]; Uther's Strength (11302, -1.59 DPS) [world_drop]; Alchemist's Stone (13503, -3.76 DPS) [crafted] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.68 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -2.34 DPS) [quest]; Alchemist's Stone (13503, -5.16 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inventor's Focal Sword (17719, -8.72 DPS, sim-verified) [dungeon]; Arbiter's Blade (11784, -9.44 DPS) [dungeon]; Death Speaker Scepter (2816, -9.63 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Twisting Essence Jar (249456, -3.45 DPS, sim-verified) [crafted]; Forcestone Buckler (17508, -4.74 DPS) [quest]; Mountainside Buckler (13082, -4.92 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Vinerot Sandals; finger1: Darkspear Signet; finger2: Brainlash; trinket2: Darkspear Voodoo Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1047, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 334.5. Weights run: 10.6s. Verify run: 552.3s. 2461 eligible items had no known source.

Race sweep: undead 334.3 ± 0.5 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.030, intellect=1.875 ± 0.031, spirit=1.892 ± 0.023, mp5=5.381 ± 0.030, crit=0.813 ± 0.034 per rating point (14 rating = 1%, 11.387 per %), spell_haste=not significant (-0.090 ± 0.218)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 133.2 healing_power points (16.36 DPS) | yes | Whitesoul Helm (12633, -5.13 DPS) [crafted]; Black Dragonscale Helm (252605, -6.21 DPS) [crafted]; Devout Crown (16693, -12.56 DPS, sim-verified) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 70.7 healing_power points (8.68 DPS) | yes | Lady Maye's Pendant (14558, -2.95 DPS) [world_drop]; Animated Chain Necklace (18723, -3.00 DPS) [dungeon]; Tooth of Gnarr (13141, -4.20 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironfeather Shoulders (15067, +0.00 DPS) [crafted]; Darkspear Shoulderguards (272958, -0.10 DPS) [vendor]; Mantle of Lost Hope (22234, -10.73 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.05 DPS) [quest]; Darkspear Raider's Cloak (272063, -1.37 DPS) [vendor]; Gracious Cape (18743, -1.89 DPS, sim-verified) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Breastplate of Salvation (250601, -3.63 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Morlune's Bracer (18741, +0.00 DPS) [dungeon]; Bleak Howler Armguards (13208, -10.00 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -10.57 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Grimgore Noose (13403, -0.23 DPS) [dungeon]; Sash of Mercy (14553, -8.67 DPS, sim-verified) [world_drop] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legplates of the Chromatic Defier (12945, -2.87 DPS) [quest]; Devout Skirt (16694, -4.63 DPS) [dungeon]; Leggings of Arcana (12756, -7.38 DPS, sim-verified) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Wolfrunner Shoes (13101, -0.92 DPS) [world_drop]; Mooncloth Boots (15802, -11.07 DPS, sim-verified) [crafted] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.10 DPS) [dungeon]; Fordring's Seal (16058, -0.99 DPS) [quest]; Naglering (11669, -16.05 DPS, sim-verified) [dungeon] |
| finger2 | Darkspear Signet (272068) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.02 DPS) [quest]; Blessed Band of Light (272407, -0.28 DPS) [vendor]; Band of Mending (22334, -2.03 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+13.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -1.32 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.98 DPS) [vendor]; Serenity Field (272439, -2.71 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Serenity Field (272439, -0.21 DPS) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -2.60 DPS, sim-verified) [world_drop]; The Hammer of Grace (11923, -11.29 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Darkspear Signet; trinket1: Burst of Knowledge; trinket2: Ankh of Life; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2461, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 780.8. Weights run: 11.0s. Verify run: 222.7s. 2461 eligible items had no known source.

Race sweep: undead 779.9 ± 0.6 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.179, intellect=1.436 ± 0.060, spirit=1.404 ± 0.058, mp5=3.558 ± 0.063, crit=1.564 ± 0.074 per rating point (14 rating = 1%, 21.893 per %), spell_haste=not significant (-0.241 ± 1.121)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 118.5 healing_power points (30.50 DPS) | yes | Black Dragonscale Helm (252605, -10.82 DPS) [crafted]; Devout Crown (16693, -11.06 DPS) [dungeon]; Whitesoul Helm (12633, -15.98 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 57.9 healing_power points (14.90 DPS) | yes | Tooth of Gnarr (13141, -4.23 DPS) [dungeon]; Lady Maye's Pendant (14558, -5.05 DPS) [world_drop]; Animated Chain Necklace (18723, -6.63 DPS, sim-verified) [dungeon] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 63.1 healing_power points (16.24 DPS) | yes | Darkspear Mantle (272107, -1.08 DPS) [vendor]; Dawnbringer Shoulders (12625, -1.30 DPS) [crafted]; Mantle of the Timbermaw (19050, -1.43 DPS) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -2.06 DPS) [rep]; Raincaster Drape (12110, -2.18 DPS) [quest]; Gracious Cape (18743, -3.24 DPS, sim-verified) [dungeon] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 106.4 healing_power points (27.40 DPS) | yes | Robes of the Exalted (13346, -5.11 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -6.15 DPS, sim-verified) [world]; Mooncloth Vest (14138, -6.44 DPS) [crafted] |
| wrist | Loomguard Armbraces (13969) | Scholomance: Kirtonos the Herald [dungeon] | 49.7 healing_power points (12.80 DPS) | yes | Morlune's Bracer (18741, -0.53 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.55 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -1.65 DPS) [rep] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, -0.96 DPS) [vendor]; Mender's Leather Gauntlets (252551, -2.56 DPS) [crafted]; Mender's Mail Gauntlets (252587, -2.56 DPS) [crafted] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 69.9 healing_power points (17.98 DPS) | yes | Wisdom of the Timbermaw (19047, -1.16 DPS, sim-verified) [crafted]; Belt of the Ordained (18702, -4.31 DPS) [dungeon]; Devout Belt (16696, -4.50 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, -7.60 DPS) [quest]; Ironfeather Leggings (252486, -8.40 DPS) [crafted]; Legplates of the Chromatic Defier (12945, -9.26 DPS, sim-verified) [quest] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 62.4 healing_power points (16.05 DPS) | yes | Faith Healer's Boots (22247, -1.10 DPS) [dungeon]; Devout Sandals (16691, -3.38 DPS) [dungeon]; Firemoss Boots (22275, -4.25 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.06 DPS) [dungeon]; Blessed Band of Light (272407, -1.87 DPS) [vendor]; Naglering (11669, -17.46 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.87 DPS) [vendor]; Band of Rumination (18103, -1.53 DPS) [dungeon]; Rosewine Circle (13178, -2.19 DPS, sim-verified) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Serenity Field (272439, -4.75 DPS, sim-verified) [vendor] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -2.33 DPS, sim-verified) [dungeon] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, -16.57 DPS) [vendor]; High Warlord's Destroyer (234546, -33.36 DPS) [pvp]; Hand of Edward the Odd (2243, -45.25 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Shroud of the Nathrezim; back: Drape of Recovery; chest: Breastplate of Salvation; wrist: Loomguard Armbraces; hands: Hands of the Exalted Herald; waist: Sash of Mercy; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Eye of the Beast; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 2461, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

