# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 50320001000000000-00000000000000000-00000000000000000)

Set DPS (verified): 51.3. Weights run: 7.0s. Verify run: 53.3s. 281 eligible items had no known source.

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

### Band 30 (human, 50320215002100000-00000000000000000-00000000000000000)

Set DPS (verified): 132.4. Weights run: 6.5s. Verify run: 38.3s. 504 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.054, intellect=1.660 ± 0.014, spirit=1.325 ± 0.011, mp5=4.523 ± 0.022, crit=0.565 ± 0.017 per rating point (14 rating = 1%, 7.913 per %), spell_haste=not significant (-0.359 ± 0.096)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 40.9 healing_power points (6.53 DPS) | yes | Nightsky Cowl (4039, -2.30 DPS) [world_drop]; Acolyte's Chain Helm (250501, -2.55 DPS) [crafted]; Enduring Cap (3020, -4.59 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 11.9 healing_power points (1.90 DPS) | yes | Crystal Starfire Medallion (5003, -0.20 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.32 DPS) [world]; Seal of the Expedition (271920, -0.58 DPS) [rep] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 24.3 healing_power points (3.87 DPS) | yes | Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Mantle of Honor (3560, -0.54 DPS) [quest]; Magician's Mantle (12998, -0.69 DPS) [world_drop] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.11 DPS) [dungeon]; Expeditionary's Cape (271923, -0.17 DPS) [rep]; Repairman's Cape (9605, -0.70 DPS, sim-verified) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.21 DPS) [crafted]; Beguiler Robes (7728, -0.29 DPS, sim-verified) [dungeon]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.6 healing_power points (3.60 DPS) | yes | Mindthrust Bracers (1974, -0.90 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.38 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.48 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 24.2 healing_power points (3.86 DPS) | yes | Truefaith Gloves (7049, -0.67 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.69 DPS) [dungeon]; Acolyte's Gloves (250511, -0.94 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 36.9 healing_power points (5.89 DPS) | yes | Prefect's Belt (250559, -0.28 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -1.86 DPS) [crafted]; Belt of Arugal (6392, -2.60 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 41.9 healing_power points (6.68 DPS) | yes | Pristine Leggings (253987, -0.23 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.58 DPS) [crafted]; Wisdom's Leather Pants (252503, -1.06 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 36.9 healing_power points (5.89 DPS) | yes | Acolyte's Boots (250506, -2.81 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.81 DPS) [crafted]; Acidic Walkers (9454, -3.16 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 19.6 healing_power points (3.12 DPS) | yes | Runemender's Seal (273031, -0.21 DPS) [world]; Lorekeeper's Ring (19525, -0.56 DPS) [rep]; Snake Hoop (6750, -0.84 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.3 healing_power points (2.92 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Lorekeeper's Ring (19525, -0.36 DPS) [rep]; Snake Hoop (6750, -0.63 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crested Scepter (3414, -5.22 DPS) [dungeon]; Pysan's Old Greatsword (1975, -6.22 DPS) [dungeon]; Taskmaster Axe (5194, -7.96 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 20.9 healing_power points (3.34 DPS) | yes | Orb of Mystic Insight (249394, -0.94 DPS, sim-verified) [crafted]; Heart of Agamaggan (6694, -0.95 DPS) [dungeon]; Orb of Mistmantle (13031, -1.06 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 504, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 50320215012101251-00000000000000000-00000000000000000)

Set DPS (verified): 199.0. Weights run: 8.7s. Verify run: 40.7s. 691 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.275, intellect=1.003 ± 0.026, spirit=0.860 ± 0.024, mp5=3.254 ± 0.036, crit=0.884 ± 0.033 per rating point (14 rating = 1%, 12.379 per %), spell_haste=2.651 ± 0.470

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 46.8 healing_power points (10.08 DPS) | yes | Whitemane's Chapeau (7720, -1.41 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.86 DPS) [world_drop]; Earthen Silk Hood (254015, -3.49 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.0 healing_power points (6.03 DPS) | yes | Triune Amulet (7722, -3.11 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -3.23 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -3.88 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.0 healing_power points (7.55 DPS) | yes | Earthen Silk Shoulders (254033, -2.42 DPS, sim-verified) [crafted]; Sheepshear Mantle (13115, -2.98 DPS) [world_drop]; Mantle of Doan (7712, -3.26 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Expeditionary's Cape (271922, -1.69 DPS) [rep]; Glowing Thresher Cape (6901, -1.81 DPS) [dungeon]; Repairman's Cape (9605, -1.94 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.9 healing_power points (13.97 DPS) | yes | Deathchill Armor (10764, -4.97 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -5.56 DPS) [crafted]; Civinad Robes (9623, -6.15 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.8 healing_power points (4.04 DPS) | yes | Earthen Silk Cuffs (254019, -0.96 DPS) [crafted]; Mindthrust Bracers (1974, -1.04 DPS, sim-verified) [dungeon]; Enchanted Kodo Bracers (13119, -1.32 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -0.43 DPS) [crafted]; Prefect's Gauntlet (250569, -0.43 DPS) [crafted]; Gilded Handwraps (254021, -0.50 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.2 healing_power points (7.58 DPS) | yes | Mender's Leather Belt (252523, -1.26 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.33 DPS) [crafted]; Deathmage Sash (10771, -2.83 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 42.1 healing_power points (9.06 DPS) | yes | Stoneweaver Leggings (9407, -0.93 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.73 DPS) [crafted]; Pristine Leggings (253987, -1.88 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 37.0 healing_power points (7.98 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Prefect's Boots (250549, -0.00 DPS) [crafted]; Gilded Slippers (254001, -1.42 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 18.8 healing_power points (4.04 DPS) | yes | Electrocutioner Lagnut (9447, -0.81 DPS) [dungeon]; Runemender's Seal (273031, -0.81 DPS) [world]; Philanthropist's Ring (281635, -0.81 DPS) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 15.8 healing_power points (3.39 DPS) | yes | Electrocutioner Lagnut (9447, -0.16 DPS) [dungeon]; Runemender's Seal (273031, -0.16 DPS) [world]; Philanthropist's Ring (281635, -0.16 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -6.82 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -10.55 DPS) [dungeon]; Bookmaker's Scepter (4122, -13.97 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of Lorica (11262, -0.61 DPS, sim-verified) [quest]; Battered Viking Shield (9403, -0.83 DPS) [dungeon]; Orb of Souls (249395, -1.57 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Lorekeeper's Ring; finger2: Ring of Power Regulation; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 691, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 50320215012101251-50320000000000000-00000000000000000)

Set DPS (verified): 214.8. Weights run: 14.5s. Verify run: 195.2s. 883 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=1.607 ± 0.018, spirit=1.263 ± 0.015, mp5=4.432 ± 0.018, crit=0.667 ± 0.026 per rating point (14 rating = 1%, 9.332 per %), spell_haste=not significant (0.220 ± 0.144)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 60.7 healing_power points (8.05 DPS) | yes | Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Gemburst Circlet (10751, -0.22 DPS) [quest]; Soothsayer's Headdress (17740, -0.46 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.9 healing_power points (4.36 DPS) | yes | Gemshard Heart (17707, -1.22 DPS) [dungeon]; Mindburst Medallion (11196, -1.45 DPS) [quest]; Horizon Choker (13085, -1.54 DPS, sim-verified) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 57.2 healing_power points (7.59 DPS) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, +0.00 DPS) [vendor]; Living Shoulders (15061, -1.30 DPS) [crafted]; Lead Surveyor's Mantle (11842, -2.15 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19531, -0.22 DPS) [rep]; Expeditionary's Cape (271921, -0.32 DPS) [rep]; Featherskin Cape (10843, -0.78 DPS) [world] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 69.7 healing_power points (9.25 DPS) | yes | Knight's Imbued Armor (220813, -0.28 DPS) [vendor]; Feathered Breastplate (8349, -0.58 DPS) [crafted]; Stormcloth Vest (10020, -1.66 DPS, sim-verified) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 37.8 healing_power points (5.02 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.26 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.22 DPS) [crafted]; Prefect's Gloves (250554, -0.55 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 49.3 healing_power points (6.54 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.43 DPS) [crafted]; Gilded Waistcord (254081, -0.77 DPS) [crafted] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 67.6 healing_power points (8.97 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.32 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -1.53 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -2.05 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 45.7 healing_power points (6.06 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.03 DPS) [crafted]; Mender's Leather Shoes (252533, -0.21 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 30.4 healing_power points (4.04 DPS) | yes | Chivalrous Signet (20505, -0.55 DPS) [quest]; Cyclopean Band (11824, -0.68 DPS) [dungeon]; Eye of Adaegus (5266, -0.75 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 29.7 healing_power points (3.94 DPS) | yes | Cyclopean Band (11824, -0.59 DPS) [dungeon]; Eye of Adaegus (5266, -0.65 DPS) [world_drop]; Chivalrous Signet (20505, -1.84 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+6.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.34 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.67 DPS) [quest] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.55 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -1.60 DPS) [quest]; Evonice's Landin' Pilla (18951, -2.42 DPS, sim-verified) [quest] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -4.79 DPS) [dungeon]; Taskmaster Axe (5194, -6.50 DPS) [dungeon]; Inventor's Focal Sword (17719, -6.89 DPS, sim-verified) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gizlock's Hypertech Buckler (17718, +0.00 DPS) [dungeon]; Cloud Stone (17737, -0.05 DPS) [dungeon]; Orb of Lorica (11262, -0.14 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket2: Enriched Seal; off_hand: Twisting Essence Jar

No-known-source sample (15 of 883, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 338.1. Weights run: 15.3s. Verify run: 891.4s. 2115 eligible items had no known source.

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

Set DPS (verified): 743.8. Weights run: 8.0s. Verify run: 361.2s. 2115 eligible items had no known source.

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

Set DPS (verified): 50.3. Weights run: 7.0s. Verify run: 52.7s. 268 eligible items had no known source.

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

### Band 30 (undead, 50320215002100000-00000000000000000-00000000000000000)

Set DPS (verified): 127.6. Weights run: 6.5s. Verify run: 40.7s. 491 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.054, intellect=1.660 ± 0.014, spirit=1.325 ± 0.011, mp5=4.523 ± 0.022, crit=0.565 ± 0.017 per rating point (14 rating = 1%, 7.913 per %), spell_haste=not significant (-0.359 ± 0.096)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 40.9 healing_power points (6.53 DPS) | yes | Nightsky Cowl (4039, -2.30 DPS) [world_drop]; Acolyte's Chain Helm (250501, -2.55 DPS) [crafted]; Enduring Cap (3020, -2.77 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 11.9 healing_power points (1.90 DPS) | yes | Crystal Starfire Medallion (5003, -0.31 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.32 DPS) [world]; Darkspear Warding Pendant (272075, -0.58 DPS) [vendor] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 24.3 healing_power points (3.87 DPS) | yes | Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Magician's Mantle (12998, -0.69 DPS) [world_drop]; Nightsky Mantle (4718, -0.69 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, -0.17 DPS) [vendor]; Glowing Thresher Cape (6901, -0.49 DPS, sim-verified) [dungeon]; Windsong Drape (15468, -1.28 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 39.5 healing_power points (6.31 DPS) | yes | Death Speaker Robes (6682, -0.40 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.41 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.58 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.6 healing_power points (3.60 DPS) | yes | Mindthrust Bracers (1974, -0.42 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.38 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.48 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 24.2 healing_power points (3.86 DPS) | yes | Truefaith Gloves (7049, -0.64 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.69 DPS) [dungeon]; Acolyte's Gloves (250511, -0.94 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 36.9 healing_power points (5.89 DPS) | yes | Prefect's Belt (250559, -0.67 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -1.86 DPS) [crafted]; Belt of Arugal (6392, -2.60 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 41.9 healing_power points (6.68 DPS) | yes | Pristine Leggings (253987, -0.42 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.58 DPS) [crafted]; Wisdom's Leather Pants (252503, -1.06 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 36.9 healing_power points (5.89 DPS) | yes | Acolyte's Boots (250506, -2.81 DPS) [crafted]; Wisdom's Leather Boots (252444, -2.81 DPS) [crafted]; Acidic Walkers (9454, -3.18 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 19.6 healing_power points (3.12 DPS) | yes | Runemender's Seal (273031, -0.21 DPS) [world]; Darkspear Signet (272071, -0.24 DPS) [vendor]; Solid Gold Band (279845, -0.31 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.3 healing_power points (2.92 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Darkspear Signet (272071, -0.03 DPS) [vendor]; Solid Gold Band (279845, -0.10 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crested Scepter (3414, -5.22 DPS) [dungeon]; The Stitcher (279874, -6.06 DPS) [quest]; Taskmaster Axe (5194, -7.77 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 20.9 healing_power points (3.34 DPS) | yes | Orb of Mystic Insight (249394, -0.75 DPS, sim-verified) [crafted]; Heart of Agamaggan (6694, -0.95 DPS) [dungeon]; Orb of Mistmantle (13031, -1.06 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Battle Healer's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 491, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 50320215012101251-00000000000000000-00000000000000000)

Set DPS (verified): 195.0. Weights run: 8.7s. Verify run: 42.3s. 675 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.275, intellect=1.003 ± 0.026, spirit=0.860 ± 0.024, mp5=3.254 ± 0.036, crit=0.884 ± 0.033 per rating point (14 rating = 1%, 12.379 per %), spell_haste=2.651 ± 0.470

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 46.8 healing_power points (10.08 DPS) | yes | Whitemane's Chapeau (7720, -1.34 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.86 DPS) [world_drop]; Earthen Silk Hood (254015, -3.49 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.0 healing_power points (6.03 DPS) | yes | Triune Amulet (7722, -3.06 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -3.23 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -3.88 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.0 healing_power points (7.55 DPS) | yes | Earthen Silk Shoulders (254033, -2.39 DPS, sim-verified) [crafted]; Sheepshear Mantle (13115, -2.98 DPS) [world_drop]; Mantle of Doan (7712, -3.26 DPS) [dungeon] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272077, -1.69 DPS) [vendor]; Glowing Thresher Cape (6901, -1.81 DPS) [dungeon]; Icy Cloak (4327, -2.18 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.9 healing_power points (13.97 DPS) | yes | Deathchill Armor (10764, -4.85 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -5.56 DPS) [crafted]; Civinad Robes (9623, -6.15 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.8 healing_power points (4.04 DPS) | yes | Mindthrust Bracers (1974, -0.63 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.96 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.32 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.30 DPS, sim-verified) [crafted]; Dreamweave Gloves (10019, -0.43 DPS) [crafted]; Prefect's Gauntlet (250569, -0.43 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.2 healing_power points (7.58 DPS) | yes | Mender's Leather Belt (252523, -1.12 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -1.33 DPS) [crafted]; Deathmage Sash (10771, -2.83 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 42.1 healing_power points (9.06 DPS) | yes | Stoneweaver Leggings (9407, -0.52 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.73 DPS) [crafted]; Pristine Leggings (253987, -1.88 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 37.0 healing_power points (7.98 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Prefect's Boots (250549, -0.00 DPS) [crafted]; Gilded Slippers (254001, -1.42 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 19.5 healing_power points (4.20 DPS) | yes | Ring of Power Regulation (273030, -0.81 DPS) [dungeon]; Solid Gold Band (279845, -0.91 DPS) [quest]; Electrocutioner Lagnut (9447, -0.97 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 18.8 healing_power points (4.04 DPS) | yes | Ring of Power Regulation (273030, -0.56 DPS, sim-verified) [dungeon]; Solid Gold Band (279845, -0.75 DPS) [quest]; Electrocutioner Lagnut (9447, -0.81 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -6.36 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -10.55 DPS) [dungeon]; Skullbreaker (17039, -14.18 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forcestone Buckler (17508, -0.12 DPS) [quest]; Battered Viking Shield (9403, -0.83 DPS) [dungeon]; Omega Orb (7749, -1.44 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 675, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 50320215012101251-50320000000000000-00000000000000000)

Set DPS (verified): 204.9. Weights run: 14.5s. Verify run: 189.6s. 883 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.197, intellect=1.295 ± 0.027, spirit=1.314 ± 0.018, mp5=4.642 ± 0.029, crit=1.228 ± 0.039 per rating point (14 rating = 1%, 17.193 per %), spell_haste=2.019 ± 0.435

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 55.8 healing_power points (14.61 DPS) | yes | Eye of Theradras (17715, -0.24 DPS) [dungeon]; Gemburst Circlet (10751, -0.31 DPS) [quest]; Cassandra's Grace (13102, -0.96 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.4 healing_power points (7.94 DPS) | yes | Horizon Choker (13085, -1.89 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -2.49 DPS) [dungeon]; Mindburst Medallion (11196, -2.62 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 51.4 healing_power points (13.45 DPS) | yes | Living Shoulders (15061, -1.71 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -1.72 DPS) [crafted]; Mender's Mail Shoulders (252569, -1.72 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.43 DPS, sim-verified) [dungeon]; Featherskin Cape (10843, -1.31 DPS) [world]; Darkspear Raider's Cloak (272076, -1.36 DPS) [vendor] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 72.4 healing_power points (18.95 DPS) | yes | Stormcloth Vest (10020, -1.94 DPS, sim-verified) [crafted]; Feathered Breastplate (8349, -2.34 DPS) [crafted]; Embrace of the Wind Serpent (12462, -2.87 DPS) [world] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 36.0 healing_power points (9.40 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.08 DPS) [crafted]; Prefect's Wristguards (250584, -0.68 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.53 DPS) [crafted]; Prefect's Gloves (250554, -1.37 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 45.5 healing_power points (11.91 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.68 DPS) [crafted]; Gilded Waistcord (254081, -1.28 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 57.7 healing_power points (15.10 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.24 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -1.38 DPS, sim-verified) [dungeon]; Warchief Kilt (7760, -1.98 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 42.7 healing_power points (11.16 DPS) | yes | Mender's Mail Boots (252565, -0.10 DPS) [crafted]; Mender's Leather Boots (252472, -0.11 DPS) [crafted]; Mender's Leather Shoes (252533, -0.41 DPS, sim-verified) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 32.5 healing_power points (8.50 DPS) | yes | Brainlash (6440, -1.70 DPS) [dungeon]; Eye of Adaegus (5266, -2.34 DPS) [world_drop]; Cyclopean Band (11824, -2.40 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 30.6 healing_power points (8.00 DPS) | yes | Brainlash (6440, -1.64 DPS, sim-verified) [dungeon]; Eye of Adaegus (5266, -1.84 DPS) [world_drop]; Cyclopean Band (11824, -1.90 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.69 DPS) [quest]; Alchemist's Stone (13503, -2.75 DPS) [crafted] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.89 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -2.63 DPS) [quest]; Alchemist's Stone (13503, -4.70 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inventor's Focal Sword (17719, -6.60 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -9.55 DPS) [dungeon]; Taskmaster Axe (5194, -12.82 DPS) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gizlock's Hypertech Buckler (17718, +0.00 DPS) [dungeon]; Forcestone Buckler (17508, -0.26 DPS) [quest]; Beacon of Hope (9393, -0.59 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Ironfeather Shoulders; back: Battle Healer's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket2: Darkspear Voodoo Seal; off_hand: Twisting Essence Jar

No-known-source sample (15 of 883, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 321.8. Weights run: 15.3s. Verify run: 883.7s. 2141 eligible items had no known source.

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

Set DPS (verified): 734.2. Weights run: 8.0s. Verify run: 378.2s. 2141 eligible items had no known source.

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

