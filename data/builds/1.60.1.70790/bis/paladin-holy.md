# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 51.3. Weights run: 4.1s. Verify run: 31.2s. 281 eligible items had no known source.

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

No-known-source sample (15 of 281, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 30 (human, 50320015000100000-40000000000000000-00000000000000000)

Set DPS (verified): 134.6. Weights run: 3.6s. Verify run: 21.0s. 504 eligible items had no known source.

Race sweep: human 134.6 ± 0.1, dwarf 128.8 ± 0.1 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=1.891 ± 0.016, spirit=1.419 ± 0.012, mp5=4.787 ± 0.023, crit=0.598 ± 0.018 per rating point (14 rating = 1%, 8.373 per %), spell_haste=-0.471 ± 0.113

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 41.5 healing_power points (6.29 DPS) | yes | Enduring Cap (3020, -0.20 DPS) [world_drop]; Nightsky Cowl (4039, -1.78 DPS) [world_drop]; Dokebi Hat (14584, -2.06 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.2 healing_power points (2.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.21 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.29 DPS) [world]; Seal of the Expedition (271920, -0.57 DPS) [rep] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Mantle of Honor (3560, -0.55 DPS) [quest]; Nightsky Mantle (4718, -0.69 DPS) [world_drop] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Repairman's Cape (9605, +0.00 DPS) [quest]; Expeditionary's Cape (271923, +0.00 DPS) [rep]; Glowing Thresher Cape (6901, -0.05 DPS) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beguiler Robes (7728, +0.00 DPS) [dungeon]; Pristine Gown (253961, -0.29 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.44 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Nightsky Wristbands (6407, -1.18 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.25 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 26.5 healing_power points (4.01 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.65 DPS) [dungeon]; Truefaith Gloves (7049, -0.88 DPS) [crafted]; Acolyte's Gloves (250511, -1.06 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prefect's Belt (250559, -0.55 DPS) [crafted]; Skycaller's Leather Belt (252522, -1.79 DPS) [crafted]; Silver-lined Belt (13011, -2.45 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 44.2 healing_power points (6.69 DPS) | yes | Pristine Leggings (253987, -0.28 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.61 DPS) [crafted]; Wisdom's Leather Pants (252503, -1.08 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.9 healing_power points (5.89 DPS) | yes | Acidic Walkers (9454, -1.99 DPS) [dungeon]; Acolyte's Boots (250506, -2.80 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.80 DPS) [crafted] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 20.4 healing_power points (3.08 DPS) | yes | Runemender's Seal (273031, -0.22 DPS) [world]; Lorekeeper's Ring (19525, -0.57 DPS) [rep]; Snake Hoop (6750, -0.74 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.9 healing_power points (2.87 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Lorekeeper's Ring (19525, -0.36 DPS) [rep]; Snake Hoop (6750, -0.53 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taskmaster Axe (5194, -1.76 DPS) [dungeon]; Crested Scepter (3414, -4.88 DPS) [dungeon]; Pysan's Old Greatsword (1975, -5.91 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 21.5 healing_power points (3.26 DPS) | yes | Orb of Mystic Insight (249394, -0.48 DPS) [crafted]; Heart of Agamaggan (6694, -0.99 DPS) [dungeon]; Orb of Mistmantle (13031, -1.04 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 504, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 50320115000100000-50000000000000000-05030000000000000)

Set DPS (verified): 217.9. Weights run: 4.0s. Verify run: 25.6s. 691 eligible items had no known source.

Race sweep: human 217.9 ± 0.1, dwarf 210.2 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.086, intellect=1.532 ± 0.014, spirit=1.462 ± 0.013, mp5=4.606 ± 0.025, crit=0.527 ± 0.018 per rating point (14 rating = 1%, 7.384 per %), spell_haste=not significant (0.236 ± 0.171)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 61.2 healing_power points (9.81 DPS) | yes | Whitemane's Chapeau (7720, -1.60 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.11 DPS) [world_drop]; Miner's Hat of the Deep (9429, -3.29 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.3 healing_power points (5.17 DPS) | yes | Triune Amulet (7722, -2.25 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.71 DPS) [quest]; Necklace of Calisea (1714, -3.91 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 41.9 healing_power points (6.72 DPS) | yes | Earthen Silk Shoulders (254033, -1.96 DPS) [crafted]; Mantle of Doan (7712, -2.25 DPS) [dungeon]; Sheepshear Mantle (13115, -3.97 DPS, sim-verified) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -1.06 DPS) [dungeon]; Blackforge Cape (6424, -1.15 DPS) [dungeon]; Expeditionary's Cape (271922, -2.64 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 69.7 healing_power points (11.17 DPS) | yes | Dreamweave Vest (10021, -2.79 DPS) [crafted]; Robe of Power (7054, -4.10 DPS) [crafted]; Deathchill Armor (10764, -7.18 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.8 healing_power points (3.66 DPS) | yes | Enchanted Kodo Bracers (13119, -0.33 DPS) [world_drop]; Earthen Silk Cuffs (254019, -0.88 DPS) [crafted]; Mindthrust Bracers (1974, -1.12 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.33 DPS) [crafted]; Bonefingers (10765, -0.41 DPS) [dungeon]; Dreamweave Gloves (10019, -0.50 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.0 healing_power points (6.89 DPS) | yes | Mender's Leather Belt (252523, -1.31 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.65 DPS) [crafted]; Deathmage Sash (10771, -2.09 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 53.7 healing_power points (8.60 DPS) | yes | Stoneweaver Leggings (9407, -0.96 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.88 DPS) [crafted]; Pristine Leggings (253987, -2.18 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 45.0 healing_power points (7.20 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Prefect's Boots (250549, -0.51 DPS) [crafted]; Gilded Slippers (254001, -1.34 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 22.8 healing_power points (3.66 DPS) | yes | Runemender's Seal (273031, -0.57 DPS) [world]; Welken Ring (5011, -0.79 DPS) [world_drop]; Ring of Power Regulation (273030, -4.43 DPS, sim-verified) [dungeon] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Welken Ring (5011, -0.21 DPS) [world_drop]; Ring of Power Regulation (273030, -0.32 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taskmaster Axe (5194, -7.85 DPS) [dungeon]; Death Speaker Scepter (2816, -8.55 DPS, sim-verified) [dungeon]; Bookmaker's Scepter (4122, -10.66 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of Lorica (11262, -0.08 DPS) [quest]; Orb of Souls (249395, -0.97 DPS) [crafted]; Battered Viking Shield (9403, -1.06 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 691, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 50320215012101230-50000000000000000-05030000000000000)

Set DPS (verified): 240.6. Weights run: 7.8s. Verify run: 102.5s. 883 eligible items had no known source.

Race sweep: human 240.2 ± 0.3, dwarf 228.0 ± 0.3 (heal_score); the guide race human is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.207, intellect=1.099 ± 0.023, spirit=1.351 ± 0.017, mp5=4.046 ± 0.026, crit=0.926 ± 0.034 per rating point (14 rating = 1%, 12.965 per %), spell_haste=not significant (1.409 ± 0.437)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 56.2 healing_power points (18.31 DPS) | yes | Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Papal Fez (9431, -1.08 DPS) [dungeon]; Gemburst Circlet (10751, -1.35 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.8 healing_power points (9.39 DPS) | yes | Horizon Choker (13085, -1.99 DPS, sim-verified) [world_drop]; Lei of Lilies (1315, -2.78 DPS) [world_drop]; Mindburst Medallion (11196, -3.08 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, +0.00 DPS) [vendor]; Living Shoulders (15061, -0.66 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -1.45 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.46 DPS) [dungeon]; Featherskin Cape (10843, -1.78 DPS) [world]; Expeditionary's Cape (271921, -2.60 DPS) [rep] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Vest (10020, -1.75 DPS, sim-verified) [crafted]; Feathered Breastplate (8349, -1.95 DPS) [crafted]; Knight's Imbued Armor (220813, -2.91 DPS) [vendor] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 34.8 healing_power points (11.35 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.11 DPS) [crafted]; Prefect's Wristguards (250584, -0.98 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.73 DPS) [crafted]; Prefect's Gloves (250554, -1.94 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 43.2 healing_power points (14.09 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.72 DPS) [crafted]; Earthenweave Cord (254077, -0.76 DPS) [crafted] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 59.5 healing_power points (19.39 DPS) | yes | Senior Designer's Pantaloons (11841, -1.68 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Kilt (9474, -2.22 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -3.06 DPS) [dungeon] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 41.2 healing_power points (13.42 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.08 DPS) [crafted]; Earthenweave Boots (254093, -0.10 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 28.2 healing_power points (9.19 DPS) | yes | Chivalrous Signet (20505, -1.79 DPS) [quest]; Cyclopean Band (11824, -1.98 DPS) [dungeon]; Brainlash (6440, -9.87 DPS, sim-verified) [dungeon] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Chivalrous Signet (20505, -0.04 DPS) [quest]; Cyclopean Band (11824, -0.23 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+10.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.88 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.76 DPS) [quest] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.73 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.59 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -3.07 DPS) [quest] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inventor's Focal Sword (17719, -9.91 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -12.00 DPS) [dungeon]; Taskmaster Axe (5194, -15.98 DPS) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gizlock's Hypertech Buckler (17718, -0.86 DPS) [dungeon]; Battered Viking Shield (9403, -1.08 DPS) [dungeon]; Beacon of Hope (9393, -1.46 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Cassandra's Grace; shoulder: Ironfeather Shoulders; back: Caretaker's Cape; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; finger1: Lorekeeper's Ring; finger2: Eye of Adaegus; trinket2: Enriched Seal; off_hand: Twisting Essence Jar

No-known-source sample (15 of 883, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 338.1. Weights run: 8.9s. Verify run: 519.0s. 2115 eligible items had no known source.

Race sweep: human 337.7 ± 0.5, dwarf 320.0 ± 0.5 (heal_score); the guide race human is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.030, intellect=1.875 ± 0.031, spirit=1.892 ± 0.023, mp5=5.381 ± 0.030, crit=0.813 ± 0.034 per rating point (14 rating = 1%, 11.387 per %), spell_haste=not significant (-0.090 ± 0.218)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 133.2 healing_power points (16.36 DPS) | yes | Field Marshal's Lamellar Helmet (231640, -2.50 DPS) [vendor]; Lieutenant Commander's Lamellar Helmet (227149, -3.02 DPS) [vendor]; Crown of the Penitent (13216, -10.31 DPS, sim-verified) [quest] |
| neck | Stormpike Sage's Pendant (19098) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Royal Cap Spaulders (14548, -11.65 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostweaver Cape (12968, -0.90 DPS) [dungeon]; Shroud of the Exile (15421, -1.37 DPS) [quest]; Gracious Cape (18743, -2.28 DPS, sim-verified) [dungeon] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 128.7 healing_power points (15.81 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, -1.94 DPS) [pvp]; Knight-Captain's Lamellar Chestplate (227151, -2.94 DPS) [vendor]; Mooncloth Vest (14138, -9.96 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Bleak Howler Armguards (13208, -9.92 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Theramore Handwraps (271945, +0.00 DPS) [rep]; Hands of the Exalted Herald (12554, -11.16 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Elderwild Waistcord (279252, -0.18 DPS) [crafted]; Wisdom of the Timbermaw (19047, -10.65 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Knight-Captain's Lamellar Legguards (227150, +0.00 DPS) [vendor]; Martyr's Legplates (250600, -3.74 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -10.58 DPS, sim-verified) [crafted] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.05 DPS) [dungeon]; Fordring's Seal (16058, -1.17 DPS) [quest]; Naglering (11669, -15.82 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon]; Seal of Rivendare (13345, -0.12 DPS) [dungeon] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (+9.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -0.66 DPS) [rep]; Ankh of Life (1713, -1.18 DPS) [world_drop]; Second Wind (11819, -1.26 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, -1.85 DPS) [rep]; Second Wind (11819, -2.46 DPS) [dungeon]; Ankh of Life (1713, -3.07 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Sageblade (22383, -3.44 DPS) [crafted]; Hand of Righteousness (7721, -5.54 DPS, sim-verified) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, -0.14 DPS) [vendor]; Thaurissan's Royal Scepter (11928, -5.83 DPS, sim-verified) [dungeon] |
| ranged | Libram of Infusion (279248) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Stormpike Sage's Pendant; shoulder: Devout Mantle; back: Drape of Recovery; chest: Breastplate of Salvation; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Blessed Band of Light; trinket1: Shard of the Splithooves; trinket2: Serenity Field; main_hand: Hand of Edward the Odd; off_hand: Lei of the Lifegiver; ranged: Libram of Infusion

No-known-source sample (15 of 2115, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 743.8. Weights run: 4.6s. Verify run: 208.8s. 2115 eligible items had no known source.

Race sweep: human 744.3 ± 0.7, dwarf 723.3 ± 0.6 (heal_score); the guide race human is best.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.667, intellect=1.606 ± 0.078, spirit=1.252 ± 0.075, mp5=2.594 ± 0.084, crit=1.713 ± 0.088 per rating point (14 rating = 1%, 23.986 per %), spell_haste=not significant (5.667 ± 1.563)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 119.4 healing_power points (42.21 DPS) | yes | Field Marshal's Lamellar Helmet (231640, -2.81 DPS) [vendor]; Field Marshal's Lamellar Faceguard (16474, -8.96 DPS) [vendor]; Black Dragonscale Helm (252605, -13.74 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 66.7 healing_power points (23.57 DPS) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.88 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostweaver Cape (12968, -2.50 DPS) [dungeon]; Caretaker's Cape (19530, -2.85 DPS, sim-verified) [rep]; Expeditionary's Cape (271907, -2.89 DPS) [rep] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 103.8 healing_power points (36.68 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Robes of the Exalted (13346, -5.95 DPS, sim-verified) [dungeon]; Knight-Captain's Lamellar Chestplate (227151, -6.19 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Mender's Leather Bracers (252543, -2.07 DPS) [crafted] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Mender's Leather Gauntlets (252551, -3.76 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Belt of the Ordained (18702, -3.50 DPS) [dungeon]; Devout Belt (16696, -3.66 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legguards (231639, -3.86 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -7.77 DPS) [vendor]; Sentinel's Silk Leggings (237815, -9.10 DPS, sim-verified) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 62.8 healing_power points (22.18 DPS) | yes | Knight-Lieutenant's Lamellar Greaves (227153, -0.54 DPS) [vendor]; Faith Healer's Boots (22247, -1.20 DPS, sim-verified) [dungeon]; Marshal's Lamellar Greaves (231646, -1.44 DPS) [vendor] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+13.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -7.07 DPS) [vendor]; Second Wind (11819, -7.80 DPS, sim-verified) [dungeon]; Uther's Strength (11302, -8.48 DPS) [world_drop] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Grand Marshal's Demolisher (234568, -1.03 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Infusion (279248) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, -1.49 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; shoulder: Shroud of the Nathrezim; back: Drape of Recovery; chest: Breastplate of Salvation; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; main_hand: Hammer of the Grand Crusader; ranged: Libram of Infusion

No-known-source sample (15 of 2115, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 50.3. Weights run: 4.1s. Verify run: 31.1s. 268 eligible items had no known source.

Race sweep: undead 50.3 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=1.023 ± 0.008, spirit=0.775 ± 0.005, mp5=3.250 ± 0.009, crit=0.194 ± 0.009 per rating point (14 rating = 1%, 2.712 per %), spell_haste=3.246 ± 0.103

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (3.61 DPS) | yes | Pristine Circlet (253949, -0.53 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.69 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.72 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.3 healing_power points (0.40 DPS) | yes | Scholarly Pendant (277203, -0.13 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.2 healing_power points (2.44 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.88 DPS) [crafted]; Slime-encrusted Pads (6461, -1.45 DPS, sim-verified) [dungeon]; Serpent's Shoulders (5404, -1.51 DPS) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.73 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.94 DPS) [crafted]; Sanguine Cape (14376, -1.11 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 22.4 healing_power points (3.86 DPS) | yes | Acolyte's Chain Shirt (250491, -0.23 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.34 DPS) [crafted]; Corsair's Overshirt (5202, -0.90 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.9 healing_power points (2.56 DPS) | yes | Drakewing Bands (12999, -1.76 DPS) [world_drop]; Bright Bracers (3647, -1.85 DPS) [world_drop]; Owl Bracers (4796, -2.45 DPS, sim-verified) [vendor] |
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

No-known-source sample (15 of 268, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (undead, 50320015000100000-40000000000000000-00000000000000000)

Set DPS (verified): 130.0. Weights run: 3.6s. Verify run: 22.3s. 491 eligible items had no known source.

Race sweep: undead 130.0 ± 0.1 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=1.891 ± 0.016, spirit=1.419 ± 0.012, mp5=4.787 ± 0.023, crit=0.598 ± 0.018 per rating point (14 rating = 1%, 8.373 per %), spell_haste=-0.471 ± 0.113

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 41.5 healing_power points (6.29 DPS) | yes | Nightsky Cowl (4039, -1.78 DPS) [world_drop]; Dokebi Hat (14584, -2.06 DPS) [world_drop]; Enduring Cap (3020, -5.18 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.2 healing_power points (2.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.16 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.29 DPS) [world]; Darkspear Warding Pendant (272075, -0.57 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Mantle (6685, -0.15 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.57 DPS) [world_drop]; Magician's Mantle (12998, -0.61 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.05 DPS) [dungeon]; Resilient Cape (14400, -1.11 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -1.57 DPS, sim-verified) [vendor] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.29 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.44 DPS) [crafted]; Beguiler Robes (7728, -0.90 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindthrust Bracers (1974, -0.79 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.18 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.25 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 26.5 healing_power points (4.01 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.37 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.88 DPS) [crafted]; Acolyte's Gloves (250511, -1.06 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 38.9 healing_power points (5.88 DPS) | yes | Prefect's Belt (250559, -0.55 DPS) [crafted]; Skycaller's Leather Belt (252522, -1.79 DPS) [crafted]; Silver-lined Belt (13011, -2.45 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 44.2 healing_power points (6.69 DPS) | yes | Pristine Leggings (253987, -0.28 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.61 DPS) [crafted]; Wisdom's Leather Pants (252503, -1.08 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.9 healing_power points (5.89 DPS) | yes | Acolyte's Boots (250506, -2.80 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.80 DPS) [crafted]; Acidic Walkers (9454, -3.21 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 20.4 healing_power points (3.08 DPS) | yes | Runemender's Seal (273031, -0.22 DPS) [world]; Solid Gold Band (279845, -0.34 DPS) [quest]; Darkspear Signet (272071, -3.28 DPS, sim-verified) [vendor] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, -0.13 DPS) [quest]; Darkspear Signet (272071, -0.59 DPS, sim-verified) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crested Scepter (3414, -4.88 DPS) [dungeon]; The Stitcher (279874, -5.75 DPS) [quest]; Taskmaster Axe (5194, -8.02 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 21.5 healing_power points (3.26 DPS) | yes | Orb of Mystic Insight (249394, -0.82 DPS, sim-verified) [crafted]; Witch's Finger (16887, -0.97 DPS) [quest]; Heart of Agamaggan (6694, -0.99 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 491, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 50320115000100000-50000000000000000-05030000000000000)

Set DPS (verified): 213.4. Weights run: 4.0s. Verify run: 24.8s. 675 eligible items had no known source.

Race sweep: undead 213.3 ± 0.1 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.086, intellect=1.532 ± 0.014, spirit=1.462 ± 0.013, mp5=4.606 ± 0.025, crit=0.527 ± 0.018 per rating point (14 rating = 1%, 7.384 per %), spell_haste=not significant (0.236 ± 0.171)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 61.2 healing_power points (9.81 DPS) | yes | Whitemane's Chapeau (7720, -1.89 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.11 DPS) [world_drop]; Miner's Hat of the Deep (9429, -3.29 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.3 healing_power points (5.17 DPS) | yes | Triune Amulet (7722, -2.25 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.71 DPS) [quest]; Necklace of Calisea (1714, -3.82 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 41.9 healing_power points (6.72 DPS) | yes | Earthen Silk Shoulders (254033, -1.96 DPS) [crafted]; Mantle of Doan (7712, -2.25 DPS) [dungeon]; Sheepshear Mantle (13115, -3.65 DPS, sim-verified) [world_drop] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (213.3 DPS) | yes | Cloak of Blight (6832, -1.01 DPS) [quest]; Glowing Thresher Cape (6901, -1.06 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -2.20 DPS, sim-verified) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 69.7 healing_power points (11.17 DPS) | yes | Dreamweave Vest (10021, -2.79 DPS) [crafted]; Robe of Power (7054, -4.10 DPS) [crafted]; Deathchill Armor (10764, -6.66 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.8 healing_power points (3.66 DPS) | yes | Enchanted Kodo Bracers (13119, -0.33 DPS) [world_drop]; Mindthrust Bracers (1974, -0.86 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.88 DPS) [crafted] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-verified (213.3 DPS) | yes | Gilded Handwraps (254021, -0.33 DPS) [crafted]; Bonefingers (10765, -0.41 DPS) [dungeon]; Dreamweave Gloves (10019, -0.56 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.0 healing_power points (6.89 DPS) | yes | Mender's Leather Belt (252523, -1.27 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.65 DPS) [crafted]; Deathmage Sash (10771, -2.09 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 53.7 healing_power points (8.60 DPS) | yes | Stoneweaver Leggings (9407, -0.82 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.88 DPS) [crafted]; Pristine Leggings (253987, -2.18 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 45.0 healing_power points (7.20 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Prefect's Boots (250549, -0.51 DPS) [crafted]; Gilded Slippers (254001, -1.34 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 27.6 healing_power points (4.43 DPS) | yes | Ring of Power Regulation (273030, -1.25 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.35 DPS) [dungeon]; Runemender's Seal (273031, -1.35 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 22.8 healing_power points (3.66 DPS) | yes | Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon]; Runemender's Seal (273031, -0.57 DPS) [world]; Ring of Power Regulation (273030, -0.65 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (213.3 DPS) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (213.3 DPS) | yes | Taskmaster Axe (5194, -7.85 DPS) [dungeon]; Death Speaker Scepter (2816, -8.26 DPS, sim-verified) [dungeon]; Skullbreaker (17039, -10.90 DPS) [quest] |
| off_hand | Forcestone Buckler (17508) | Compendium of the Fallen [quest] | sim-verified (213.3 DPS) | yes | Beacon of Hope (9393, -0.30 DPS) [dungeon]; Battered Viking Shield (9403, -0.33 DPS) [dungeon]; Omega Orb (7749, -1.10 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Forcestone Buckler

No-known-source sample (15 of 675, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 50320215012101230-50000000000000000-05030000000000000)

Set DPS (verified): 231.3. Weights run: 7.8s. Verify run: 99.7s. 883 eligible items had no known source.

Race sweep: undead 231.3 ± 0.3 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.207, intellect=1.099 ± 0.023, spirit=1.351 ± 0.017, mp5=4.046 ± 0.026, crit=0.926 ± 0.034 per rating point (14 rating = 1%, 12.965 per %), spell_haste=not significant (1.409 ± 0.437)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 56.2 healing_power points (18.31 DPS) | yes | Papal Fez (9431, -1.08 DPS) [dungeon]; Gemburst Circlet (10751, -1.35 DPS) [quest]; Soothsayer's Headdress (17740, -1.91 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.8 healing_power points (9.39 DPS) | yes | Horizon Choker (13085, -2.16 DPS, sim-verified) [world_drop]; Lei of Lilies (1315, -2.78 DPS) [world_drop]; Mindburst Medallion (11196, -3.08 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Living Shoulders (15061, -0.91 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -1.45 DPS) [crafted]; Mender's Mail Shoulders (252569, -1.45 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.46 DPS) [dungeon]; Featherskin Cape (10843, -1.78 DPS) [world]; Darkspear Raider's Cloak (272076, -2.60 DPS) [vendor] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Feathered Breastplate (8349, -1.95 DPS) [crafted]; Stormcloth Vest (10020, -2.06 DPS, sim-verified) [crafted]; Embrace of the Wind Serpent (12462, -2.99 DPS) [world] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 34.8 healing_power points (11.35 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.11 DPS) [crafted]; Prefect's Wristguards (250584, -0.98 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.73 DPS) [crafted]; Prefect's Gloves (250554, -1.94 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 43.2 healing_power points (14.09 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.72 DPS) [crafted]; Earthenweave Cord (254077, -0.76 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 57.0 healing_power points (18.59 DPS) | yes | Jinxed Hoodoo Kilt (9474, -1.14 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -2.25 DPS) [dungeon]; Warchief Kilt (7760, -2.45 DPS) [dungeon] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 41.2 healing_power points (13.42 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.08 DPS) [crafted]; Earthenweave Boots (254093, -0.10 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 28.3 healing_power points (9.24 DPS) | yes | Brainlash (6440, -1.66 DPS) [dungeon]; Eye of Adaegus (5266, -1.80 DPS) [world_drop]; Cyclopean Band (11824, -2.03 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 28.2 healing_power points (9.19 DPS) | yes | Brainlash (6440, -1.28 DPS, sim-verified) [dungeon]; Eye of Adaegus (5266, -1.75 DPS) [world_drop]; Cyclopean Band (11824, -1.98 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.88 DPS) [quest]; Alchemist's Stone (13503, -3.52 DPS) [crafted] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -2.19 DPS) [quest]; Uther's Strength (11302, -2.94 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -4.83 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inventor's Focal Sword (17719, -8.79 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -12.00 DPS) [dungeon]; Taskmaster Axe (5194, -15.98 DPS) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forcestone Buckler (17508, -0.33 DPS) [quest]; Beacon of Hope (9393, -0.78 DPS) [dungeon]; Gizlock's Hypertech Buckler (17718, -0.86 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Cassandra's Grace; shoulder: Ironfeather Shoulders; back: Battle Healer's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket2: Darkspear Voodoo Seal; off_hand: Twisting Essence Jar

No-known-source sample (15 of 883, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 321.8. Weights run: 8.9s. Verify run: 513.1s. 2141 eligible items had no known source.

Race sweep: undead 321.6 ± 0.5 (heal_score); the guide race undead is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.030, intellect=1.875 ± 0.031, spirit=1.892 ± 0.023, mp5=5.381 ± 0.030, crit=0.813 ± 0.034 per rating point (14 rating = 1%, 11.387 per %), spell_haste=not significant (-0.090 ± 0.218)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 133.2 healing_power points (16.36 DPS) | yes | Devout Crown (16693, -4.90 DPS) [dungeon]; Whitesoul Helm (12633, -5.13 DPS) [crafted]; Crown of the Penitent (13216, -10.42 DPS, sim-verified) [quest] |
| neck | Frostwolf Advisor's Pendant (19096) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Cap Spaulders (14548, +0.00 DPS) [dungeon]; Darkspear Shoulderguards (272958, -0.10 DPS) [vendor]; Ironfeather Shoulders (15067, -13.59 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostweaver Cape (12968, -0.90 DPS) [dungeon]; Shroud of the Exile (15421, -1.37 DPS) [quest]; Gracious Cape (18743, -2.76 DPS, sim-verified) [dungeon] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 128.7 healing_power points (15.81 DPS) | yes | Robes of the Exalted (13346, -3.75 DPS) [dungeon]; Ironfeather Breastplate (15066, -3.96 DPS) [crafted]; Mooncloth Vest (14138, -8.92 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Bleak Howler Armguards (13208, -10.22 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -10.04 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Elderwild Waistcord (279252, -0.18 DPS) [crafted]; Wisdom of the Timbermaw (19047, -9.65 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Ghostloom Leggings (14545, -0.62 DPS) [dungeon]; Martyr's Legplates (250600, -3.71 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Wolfrunner Shoes (13101, -0.92 DPS) [world_drop]; Mooncloth Boots (15802, -10.38 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Darkspear Signet (272068) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, -0.25 DPS) [quest]; Band of Mending (22334, -1.60 DPS, sim-verified) [dungeon] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, -0.66 DPS) [vendor]; Second Wind (11819, -1.26 DPS) [dungeon]; Ankh of Life (1713, -2.26 DPS, sim-verified) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -2.37 DPS) [world_drop]; Second Wind (11819, -2.46 DPS) [dungeon]; Darkspear Voodoo Seal (272061, -4.10 DPS, sim-verified) [vendor] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Sageblade (22383, -3.44 DPS) [crafted]; Hand of Righteousness (7721, -5.75 DPS, sim-verified) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, -0.14 DPS) [vendor]; Thaurissan's Royal Scepter (11928, -5.14 DPS, sim-verified) [dungeon] |
| ranged | Libram of Infusion (279248) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Frostwolf Advisor's Pendant; shoulder: Devout Mantle; back: Drape of Recovery; chest: Breastplate of Salvation; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Darkspear Signet; trinket1: Shard of the Splithooves; trinket2: Serenity Field; main_hand: Hand of Edward the Odd; off_hand: Lei of the Lifegiver; ranged: Libram of Infusion

No-known-source sample (15 of 2141, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 734.2. Weights run: 4.6s. Verify run: 217.9s. 2141 eligible items had no known source.

Race sweep: undead 734.1 ± 0.7 (heal_score); the guide race undead is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.667, intellect=1.606 ± 0.078, spirit=1.252 ± 0.075, mp5=2.594 ± 0.084, crit=1.713 ± 0.088 per rating point (14 rating = 1%, 23.986 per %), spell_haste=not significant (5.667 ± 1.563)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 119.4 healing_power points (42.21 DPS) | yes | Black Dragonscale Helm (252605, -11.23 DPS, sim-verified) [crafted]; Whitesoul Helm (12633, -14.69 DPS) [crafted]; Devout Crown (16693, -14.88 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Darkspear Mantle (272107, +0.00 DPS) [vendor]; Shroud of the Nathrezim (18720, -19.88 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -1.80 DPS, sim-verified) [rep]; Frostweaver Cape (12968, -2.50 DPS) [dungeon]; Raincaster Drape (12110, -2.87 DPS) [quest] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 103.8 healing_power points (36.68 DPS) | yes | Robes of the Exalted (13346, -2.11 DPS, sim-verified) [dungeon]; Mooncloth Vest (14138, -7.74 DPS) [crafted]; Devout Robe (16690, -9.36 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Loomguard Armbraces (13969, +0.00 DPS) [dungeon]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Bleak Howler Armguards (13208, -21.88 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -17.89 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of the Ordained (18702, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Sash of Mercy (14553, -19.84 DPS, sim-verified) [world_drop] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironfeather Leggings (252486, -7.34 DPS, sim-verified) [crafted]; Leggings of Arcana (12756, -10.72 DPS) [quest]; Devout Skirt (16694, -14.54 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Mender's Leather Boots (252472, -0.95 DPS) [crafted]; Mooncloth Boots (15802, -18.08 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+13.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -7.07 DPS) [vendor]; Second Wind (11819, -7.09 DPS, sim-verified) [dungeon]; Uther's Strength (11302, -8.48 DPS) [world_drop] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; High Warlord's Destroyer (234546, -1.03 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Infusion (279248) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, -1.57 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Drape of Recovery; chest: Breastplate of Salvation; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; main_hand: Hammer of the Grand Crusader; ranged: Libram of Infusion

No-known-source sample (15 of 2141, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

