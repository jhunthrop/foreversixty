# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5320000010000000)

Set DPS (verified): 61.0. Weights run: 9.6s. Verify run: 49.3s. 250 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.113, intellect=1.544 ± 0.011, spirit=0.791 ± 0.008, mp5=2.738 ± 0.014, crit=0.140 ± 0.008 per rating point (14 rating = 1%, 1.959 per %), spell_haste=not significant (0.001 ± 0.093)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (3.22 DPS) | yes | Wisdom's Leather Hood (252507, -0.19 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.99 DPS) [crafted]; Trapper's Leather Hood (252505, -1.01 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.6 healing_power points (0.28 DPS) | yes | Tarnished Locket (279870, -0.28 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 18.9 healing_power points (3.38 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.37 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.91 DPS) [dungeon]; Serpent's Shoulders (5404, -2.39 DPS) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (61.0 DPS) | yes | Heavy Woolen Cloak (4311, -0.63 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.71 DPS) [crafted]; Sanguine Cape (14376, -0.79 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 25.1 healing_power points (4.49 DPS) | yes | Filigreed Pristine Gown (253901, -0.34 DPS, sim-verified) [crafted]; Stormrider's Leather Armor (252492, -1.25 DPS) [crafted]; Corsair's Overshirt (5202, -1.39 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 15.9 healing_power points (2.85 DPS) | yes | Owl Bracers (4796, -1.43 DPS, sim-verified) [vendor]; Bright Bracers (3647, -1.75 DPS) [world_drop]; Drakewing Bands (12999, -2.00 DPS) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 16.2 healing_power points (2.89 DPS) | yes | Pristine Gloves (253913, -0.13 DPS, sim-verified) [crafted]; Stormrider's Leather Gloves (252498, -0.89 DPS) [crafted]; Magefist Gloves (12977, -0.95 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.2 healing_power points (3.07 DPS) | yes | Wisdom's Leather Belt (252433, -0.53 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.75 DPS) [crafted]; Keller's Girdle (2911, -0.86 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 32.4 healing_power points (5.80 DPS) | yes | Filigreed Pristine Leggings (253937, -0.34 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -1.79 DPS) [crafted]; Darkweave Breeches (12987, -3.02 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 18.7 healing_power points (3.35 DPS) | yes | Black Whelp Slippers (252424, -0.60 DPS, sim-verified) [crafted]; Stormrider's Leather Boots (252443, -0.89 DPS) [crafted]; Pristine Boots (253889, -0.91 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 10.5 healing_power points (1.87 DPS) | yes | Minor Channeling Ring (1449, -0.43 DPS) [quest]; Band of Purification (12996, -1.02 DPS) [world_drop]; Volcanic Rock Ring (12053, -1.04 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.3 healing_power points (1.66 DPS) | yes | Minor Channeling Ring (1449, -0.27 DPS, sim-verified) [quest]; Band of Purification (12996, -0.81 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.83 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 60.5 healing_power points (10.81 DPS) | yes | Taskmaster Axe (5194, -1.63 DPS, sim-verified) [dungeon]; Living Root (6631, -2.06 DPS) [dungeon]; Staff of the Blessed Seer (2271, -4.39 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lorekeeper's Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 250, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde

### Band 30 (dwarf, 0000000000000000-000000000000000000-5530003310010000)

Set DPS (verified): 93.2. Weights run: 10.4s. Verify run: 47.5s. 434 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.108, intellect=1.474 ± 0.013, spirit=0.837 ± 0.011, mp5=2.607 ± 0.014, crit=0.189 ± 0.010 per rating point (14 rating = 1%, 2.650 per %), spell_haste=not significant (0.155 ± 0.112)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.0 healing_power points (7.47 DPS) | yes | Enduring Cap (3020, -0.74 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -3.15 DPS) [crafted]; Nightsky Cowl (4039, -3.17 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.2 healing_power points (1.82 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop]; Seal of the Expedition (271920, -0.37 DPS) [rep] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (93.3 DPS) | yes | Bloodmage Mantle (7684, -0.11 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.78 DPS) [world_drop]; Mantle of Honor (3560, -1.19 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Repairman's Cape (9605, -0.11 DPS, sim-verified) [quest]; Expeditionary's Cape (271923, -0.40 DPS) [rep]; Glowing Thresher Cape (6901, -0.52 DPS) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 36.2 healing_power points (7.11 DPS) | yes | Pristine Gown (253961, -0.50 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.56 DPS) [dungeon]; Wisdom's Leather Tunic (252511, -0.85 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.8 healing_power points (3.30 DPS) | yes | Mindthrust Bracers (1974, -0.10 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -0.99 DPS) [world_drop]; Nightsky Wristbands (6407, -1.07 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.4 healing_power points (3.81 DPS) | yes | Gloves of Old (9395, -0.39 DPS) [world_drop]; Hotshot Pilot's Gloves (9491, -0.68 DPS) [dungeon]; Wisdom's Leather Gloves (252499, -0.69 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 32.9 healing_power points (6.45 DPS) | yes | Skycaller's Leather Belt (252522, -1.79 DPS, sim-verified) [crafted]; Highlander's Lizardhide Girdle (20105, -2.98 DPS) [rep]; Highlander's Mail Girdle (20120, -2.98 DPS) [vendor] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 37.2 healing_power points (7.30 DPS) | yes | Pristine Leggings (253987, -0.15 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.98 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.37 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.7 healing_power points (6.61 DPS) | yes | Acidic Walkers (9454, -2.21 DPS, sim-verified) [dungeon]; Wisdom's Leather Boots (252444, -3.00 DPS) [crafted]; Black Whelp Slippers (252424, -3.69 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.9 healing_power points (2.92 DPS) | yes | Ring of Power Regulation (273030, -0.20 DPS) [dungeon]; Snake Hoop (6750, -0.29 DPS) [quest]; Lorekeeper's Ring (19525, -0.52 DPS) [rep] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 14.9 healing_power points (2.92 DPS) | yes | Ring of Power Regulation (273030, -0.07 DPS, sim-verified) [dungeon]; Snake Hoop (6750, -0.29 DPS) [quest]; Lorekeeper's Ring (19525, -0.52 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.17 DPS) [dungeon]; Staff of Westfall (2042, -1.49 DPS) [quest]; Staff of the Friar (3415, -2.96 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.0 healing_power points (3.54 DPS) | yes | Orb of Mystic Insight (249394, -0.22 DPS, sim-verified) [crafted]; Heart of Agamaggan (6694, -0.59 DPS) [dungeon]; Eye of Paleth (2943, -0.99 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (dwarf, 0000000000000000-000000000000000000-5530003315510000)

Set DPS (verified): 235.6. Weights run: 6.8s. Verify run: 64.4s. 661 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.241, intellect=1.270 ± 0.018, spirit=0.667 ± 0.016, mp5=2.863 ± 0.024, crit=0.221 ± 0.017 per rating point (14 rating = 1%, 3.098 per %), spell_haste=1.381 ± 0.296

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 49.6 healing_power points (13.28 DPS) | yes | Whitemane's Chapeau (7720, -2.22 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.37 DPS) [world_drop]; Earthen Silk Hood (254015, -5.60 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.2 healing_power points (8.07 DPS) | yes | Triune Amulet (7722, -4.68 DPS) [dungeon]; Seal of the Expedition (271919, -5.69 DPS) [rep]; Necklace of Calisea (1714, -9.20 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 38.5 healing_power points (10.31 DPS) | yes | Earthen Silk Shoulders (254033, -4.06 DPS) [crafted]; Green Silken Shoulders (7057, -4.43 DPS) [crafted]; Sheepshear Mantle (13115, -9.21 DPS, sim-verified) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Repairman's Cape (9605, -2.19 DPS) [quest]; Blackforge Cape (6424, -2.28 DPS) [dungeon]; Expeditionary's Cape (271922, -5.78 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.3 healing_power points (16.95 DPS) | yes | Dreamweave Vest (10021, -6.57 DPS) [crafted]; Robe of Power (7054, -7.70 DPS) [crafted]; Deathchill Armor (10764, -16.34 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.6 healing_power points (4.71 DPS) | yes | Guardian Leather Bracers (4260, -1.06 DPS) [crafted]; Earthen Silk Cuffs (254019, -1.14 DPS) [crafted]; Mindthrust Bracers (1974, -2.66 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.41 DPS) [crafted]; Prefect's Gauntlet (250569, -0.68 DPS) [crafted]; Dreamweave Gloves (10019, -1.25 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.2 healing_power points (9.68 DPS) | yes | Prefect's Belt (250559, -1.03 DPS, sim-verified) [crafted]; Mender's Leather Belt (252523, -1.48 DPS) [crafted]; Deathmage Sash (10771, -2.71 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 41.0 healing_power points (10.96 DPS) | yes | Stoneweaver Leggings (9407, -1.44 DPS, sim-verified) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -1.55 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.71 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 37.6 healing_power points (10.05 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.61 DPS) [crafted]; Skycaller's Leather Shoes (252532, -3.30 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 17.6 healing_power points (4.71 DPS) | yes | Ring of Power Regulation (273030, -0.80 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.05 DPS) [dungeon]; Runemender's Seal (273031, -1.05 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.4 healing_power points (4.38 DPS) | yes | Electrocutioner Lagnut (9447, -0.72 DPS) [dungeon]; Runemender's Seal (273031, -0.72 DPS) [world]; Ring of Power Regulation (273030, -1.40 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -8.39 DPS) [dungeon]; Staff of the Friar (3415, -9.51 DPS) [dungeon]; Staff of Jordan (873, -21.46 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-verified (235.6 DPS) | yes | Battered Viking Shield (9403, -1.34 DPS) [dungeon]; Orb of Lorica (11262, -1.40 DPS, sim-verified) [quest]; Orb of Souls (249395, -2.05 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 661, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (dwarf, 0000000000000000-000000000000000000-5530003315513151)

Set DPS (verified): 386.7. Weights run: 10.5s. Verify run: 80.4s. 848 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.164, intellect=0.657 ± 0.019, spirit=0.588 ± 0.018, mp5=2.357 ± 0.022, crit=0.291 ± 0.018 per rating point (14 rating = 1%, 4.073 per %), spell_haste=1.810 ± 0.349

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 49.3 healing_power points (28.91 DPS) | yes | Soothsayer's Headdress (17740, -6.88 DPS) [dungeon]; Papal Fez (9431, -7.17 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -7.49 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.3 healing_power points (14.81 DPS) | yes | Mindburst Medallion (11196, -7.52 DPS, sim-verified) [quest]; Triune Amulet (7722, -7.71 DPS) [dungeon]; Horizon Choker (13085, -8.04 DPS) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 38.6 healing_power points (22.67 DPS) | yes | Nethergeld Shoulders (254049, -2.41 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -2.71 DPS) [crafted]; Mender's Mail Shoulders (252569, -2.71 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-verified (386.8 DPS) | yes | Spritecaster Cape (11623, -2.46 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -6.61 DPS) [crafted]; Expeditionary's Cape (271921, -7.85 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.7 healing_power points (36.78 DPS) | yes | Grimlok's Tribal Vestments (9415, -13.26 DPS, sim-verified) [dungeon]; Feathered Breastplate (8349, -13.50 DPS) [crafted]; Ghostweave Vest (14141, -13.54 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 27.5 healing_power points (16.15 DPS) | yes | Mender's Mail Bracers (252573, -0.24 DPS) [crafted]; Mender's Leather Bracers (252543, -0.56 DPS, sim-verified) [crafted]; Skycaller's Leather Bracers (252542, -5.87 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-verified (386.8 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -1.13 DPS) [crafted]; Earthenweave Gloves (254075, -5.71 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 37.9 healing_power points (22.22 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Waistcord (254081, -1.74 DPS) [crafted]; Earthenweave Cord (254077, -3.18 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 40.2 healing_power points (23.57 DPS) | yes | Warchief Kilt (7760, -2.61 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Kilt (9474, -3.78 DPS) [dungeon]; Stoneweaver Leggings (9407, -4.62 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 36.9 healing_power points (21.65 DPS) | yes | Mender's Mail Sabatons (252579, -0.99 DPS) [crafted]; Mender's Leather Boots (252472, -1.13 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.46 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 21.4 healing_power points (12.57 DPS) | yes | Electrocutioner Lagnut (9447, -4.88 DPS) [dungeon]; Runemender's Seal (273031, -4.88 DPS) [world]; Ring of Power Regulation (273030, -4.90 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 healing_power points (7.79 DPS) | yes | Electrocutioner Lagnut (9447, -0.10 DPS) [dungeon]; Runemender's Seal (273031, -0.10 DPS) [world]; Ring of Power Regulation (273030, -0.12 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (386.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (386.8 DPS) | yes | Enriched Seal (271903, -4.36 DPS, sim-verified) [rep]; Ankh of Life (1713, -6.42 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (386.8 DPS) | yes | Resurgence Rod (17743, -20.92 DPS, sim-verified) [quest]; Charstone Dirk (17710, -27.22 DPS) [dungeon]; Dreamstaff (249454, -34.12 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Cassandra's Grace; shoulder: Living Shoulders; back: Caretaker's Cape; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Lorekeeper's Ring; trinket1: Dormant Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 848, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 483.1. Weights run: 9.5s. Verify run: 235.2s. 1668 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.069, intellect=0.911 ± 0.029, spirit=0.736 ± 0.028, mp5=2.858 ± 0.029, crit=0.359 ± 0.023 per rating point (14 rating = 1%, 5.029 per %), spell_haste=not significant (1.132 ± 0.527)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 99.7 healing_power points (53.10 DPS) | yes | Mooncloth Circlet (14140, -18.44 DPS, sim-verified) [crafted]; Earthenweave Crown (254137, -24.74 DPS) [crafted]; Devout Crown (16693, -24.92 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.3 healing_power points (14.53 DPS) | yes | Archlight Talisman (15856, -3.06 DPS, sim-verified) [quest]; Stormpike Sage's Pendant (19098, -3.59 DPS) [rep]; Horizon Choker (13085, -6.17 DPS) [world_drop] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 45.1 healing_power points (24.02 DPS) | yes | Eternity Pauldrons (279255, -0.92 DPS, sim-verified) [crafted]; Theramore Shoulderguards (272957, -2.40 DPS) [rep]; Living Shoulders (15061, -2.41 DPS) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (483.1 DPS) | yes | Spritecaster Cape (11623, -4.26 DPS) [dungeon]; Expeditionary's Cape (271907, -6.87 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.9 healing_power points (34.02 DPS) | yes | Red Dragonscale Breastplate (15047, -3.67 DPS) [crafted]; Mooncloth Vest (14138, -5.77 DPS, sim-verified) [crafted]; Devout Robe (16690, -5.85 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 34.4 healing_power points (18.34 DPS) | yes | Nethergeld Cuffs (254061, -1.88 DPS, sim-verified) [crafted]; Mender's Leather Bracers (252543, -2.47 DPS) [crafted]; Mender's Mail Bracers (252573, -2.47 DPS) [crafted] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-verified (483.1 DPS) | yes | Mender's Leather Gauntlets (252551, -0.55 DPS, sim-verified) [crafted]; Mender's Mail Gauntlets (252587, -1.47 DPS) [crafted]; Gilded Gloves (254095, -2.45 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 50.4 healing_power points (26.83 DPS) | yes | Elderwild Waistcord (279252, -2.89 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -5.03 DPS) [crafted]; Mender's Mail Belt (252591, -5.03 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 59.4 healing_power points (31.64 DPS) | yes | Elderwild Pants (279254, -1.42 DPS, sim-verified) [crafted]; Devout Skirt (16694, -2.57 DPS) [dungeon]; Sentinel's Silk Leggings (22752, -3.60 DPS) [rep] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 46.3 healing_power points (24.67 DPS) | yes | Swarmtender's Footpads (275609, -1.42 DPS, sim-verified) [crafted]; Broodwatcher's Treaders (275621, -1.77 DPS) [crafted]; Gilded Sandals (254107, -3.79 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 28.8 healing_power points (15.33 DPS) | yes | Deep Rooted Ring (19109, -4.09 DPS) [world]; Band of the Hierophant (13096, -6.17 DPS) [world_drop]; Ring of Power Regulation (273030, -7.57 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 25.4 healing_power points (13.54 DPS) | yes | Deep Rooted Ring (19109, -2.30 DPS) [world]; Band of the Hierophant (13096, -4.38 DPS) [world_drop]; Ring of Power Regulation (273030, -5.78 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (483.1 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep]; Uther's Strength (11302, -10.57 DPS, sim-verified) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (483.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, -1.52 DPS) [rep]; Ankh of Life (1713, -4.43 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (483.1 DPS) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -11.33 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -21.09 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Urgency (279249) | Enchanting [crafted] | sim-verified (483.1 DPS) | yes | Totem of Thunder (228176, -1.31 DPS, sim-verified) [crafted] |

**New at 60:** head: Living Crown; shoulder: Mooncloth Shoulders; back: Caretaker's Cape; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Mooncloth Boots; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; trinket2: Shard of the Splithooves; ranged: Totem of Urgency

No-known-source sample (15 of 1668, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 793.1. Weights run: 5.6s. Verify run: 122.6s. 1668 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.165, intellect=0.698 ± 0.032, spirit=0.248 ± 0.029, mp5=0.773 ± 0.033, crit=0.524 ± 0.032 per rating point (14 rating = 1%, 7.341 per %), spell_haste=5.943 ± 0.735

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 89.0 healing_power points (45.76 DPS) | yes | Mooncloth Circlet (14140, -17.89 DPS, sim-verified) [crafted]; Cassandra's Grace (13102, -22.00 DPS) [world_drop]; Earthenweave Crown (254137, -22.90 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.6 healing_power points (13.15 DPS) | yes | Archlight Talisman (15856, -3.18 DPS, sim-verified) [quest]; Mindburst Medallion (11196, -7.27 DPS) [quest]; Horizon Choker (13085, -7.61 DPS) [world_drop] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 39.7 healing_power points (20.42 DPS) | yes | Eternity Pauldrons (279255, -2.70 DPS) [crafted]; Living Shoulders (15061, -2.83 DPS) [crafted]; Theramore Shoulderguards (272957, -3.08 DPS, sim-verified) [rep] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -4.27 DPS) [dungeon]; Runecloth Cloak (13860, -6.89 DPS) [crafted]; Stormpike Sage's Cloak (19086, -7.19 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.0 healing_power points (30.82 DPS) | yes | Red Dragonscale Breastplate (15047, -1.53 DPS, sim-verified) [crafted]; Mooncloth Vest (14138, -7.86 DPS) [crafted]; Earthenweave Vest (254121, -8.10 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.3 healing_power points (15.07 DPS) | yes | Nethergeld Cuffs (254061, -1.57 DPS, sim-verified) [crafted]; Mender's Leather Bracers (252543, -2.03 DPS) [crafted]; Mender's Mail Bracers (252573, -2.03 DPS) [crafted] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Green Dragonscale Gauntlets (20296, -1.03 DPS, sim-verified) [crafted]; Mender's Leather Gauntlets (252551, -1.71 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.71 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 43.0 healing_power points (22.10 DPS) | yes | Mender's Leather Waistguard (252477, -2.04 DPS, sim-verified) [crafted]; Mender's Mail Belt (252591, -2.38 DPS) [crafted]; Gilded Waistcord (254081, -3.97 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 57.9 healing_power points (29.76 DPS) | yes | Sentinel's Silk Leggings (237815, -2.09 DPS, sim-verified) [vendor]; Sentinel's Lizardhide Pants (237817, -2.18 DPS) [vendor]; Elderwild Pants (279254, -6.14 DPS) [crafted] |
| feet | Swarmtender's Footpads (275609) (or Broodwatcher's Treaders (275621)) | Tailoring [crafted] | 43.0 healing_power points (22.10 DPS) | yes | Broodwatcher's Treaders (275621, +0.00 DPS) [crafted]; Gilded Sandals (254107, -2.94 DPS) [crafted]; Mooncloth Boots (15802, -3.08 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.4 healing_power points (12.53 DPS) | yes | Lorekeeper's Ring (19522, -3.74 DPS) [rep]; Philanthropist's Ring (281635, -5.59 DPS) [quest] |
| finger2 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 19.0 healing_power points (9.75 DPS) | yes | Lorekeeper's Ring (19522, -0.97 DPS) [rep]; Philanthropist's Ring (281635, -2.82 DPS) [quest] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+7.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (793.2 DPS) | yes | Ankh of Life (1713, -5.77 DPS, sim-verified) [world_drop]; Shard of the Splithooves (10659, -6.87 DPS) [quest]; Enriched Seal (271903, -7.26 DPS) [rep] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Charstone Dirk (17710, -3.01 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -26.71 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Urgency (279249) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -6.77 DPS, sim-verified) [crafted] |

**New at 60:** head: Living Crown; shoulder: Mooncloth Shoulders; back: Caretaker's Cape; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Swarmtender's Footpads; finger1: Maiden's Circle; finger2: Deep Rooted Ring; ranged: Totem of Urgency

No-known-source sample (15 of 1668, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5320000010000000)

Set DPS (verified): 58.8. Weights run: 9.6s. Verify run: 48.7s. 239 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.113, intellect=1.544 ± 0.011, spirit=0.791 ± 0.008, mp5=2.738 ± 0.014, crit=0.140 ± 0.008 per rating point (14 rating = 1%, 1.959 per %), spell_haste=not significant (0.001 ± 0.093)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (3.22 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.99 DPS) [crafted]; Trapper's Leather Hood (252505, -1.01 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.4 healing_power points (0.42 DPS) | yes | Scholarly Pendant (277203, -0.14 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 18.9 healing_power points (3.38 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.29 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.91 DPS) [dungeon]; Serpent's Shoulders (5404, -2.39 DPS) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-verified (58.8 DPS) | yes | Heavy Woolen Cloak (4311, -0.50 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.71 DPS) [crafted]; Sanguine Cape (14376, -0.79 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 25.1 healing_power points (4.49 DPS) | yes | Filigreed Pristine Gown (253901, -0.30 DPS, sim-verified) [crafted]; Stormrider's Leather Armor (252492, -1.25 DPS) [crafted]; Corsair's Overshirt (5202, -1.39 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 15.9 healing_power points (2.85 DPS) | yes | Owl Bracers (4796, -1.34 DPS, sim-verified) [vendor]; Bright Bracers (3647, -1.75 DPS) [world_drop]; Drakewing Bands (12999, -2.00 DPS) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 16.2 healing_power points (2.89 DPS) | yes | Pristine Gloves (253913, -0.09 DPS, sim-verified) [crafted]; Stormrider's Leather Gloves (252498, -0.89 DPS) [crafted]; Magefist Gloves (12977, -0.95 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.2 healing_power points (3.07 DPS) | yes | Wisdom's Leather Belt (252433, -0.46 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.75 DPS) [crafted]; Keller's Girdle (2911, -0.86 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 32.4 healing_power points (5.80 DPS) | yes | Filigreed Pristine Leggings (253937, -0.30 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -1.79 DPS) [crafted]; Darkweave Breeches (12987, -3.02 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 18.7 healing_power points (3.35 DPS) | yes | Black Whelp Slippers (252424, -0.59 DPS, sim-verified) [crafted]; Stormrider's Leather Boots (252443, -0.89 DPS) [crafted]; Pristine Boots (253889, -0.91 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 10.5 healing_power points (1.87 DPS) | yes | Loop of Sacrifice (281673, -0.49 DPS) [quest]; Band of Purification (12996, -1.02 DPS) [world_drop]; Volcanic Rock Ring (12053, -1.04 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.3 healing_power points (1.66 DPS) | yes | Loop of Sacrifice (281673, -0.26 DPS, sim-verified) [quest]; Band of Purification (12996, -0.81 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.83 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | 51.3 healing_power points (9.18 DPS) | yes | Living Root (6631, -0.38 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -2.60 DPS) [quest]; Staff of the Blessed Seer (2271, -2.76 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon

### Band 30 (tauren, 0000000000000000-000000000000000000-5530003310010000)

Set DPS (verified): 92.6. Weights run: 10.4s. Verify run: 46.2s. 428 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.108, intellect=1.474 ± 0.013, spirit=0.837 ± 0.011, mp5=2.607 ± 0.014, crit=0.189 ± 0.010 per rating point (14 rating = 1%, 2.650 per %), spell_haste=not significant (0.155 ± 0.112)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.0 healing_power points (7.47 DPS) | yes | Enduring Cap (3020, -0.79 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -3.15 DPS) [crafted]; Nightsky Cowl (4039, -3.17 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.2 healing_power points (1.82 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS, sim-verified) [world]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.37 DPS) [vendor] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (92.6 DPS) | yes | Bloodmage Mantle (7684, -0.08 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.78 DPS) [world_drop]; Nightsky Mantle (4718, -1.22 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, -0.40 DPS) [vendor]; Glowing Thresher Cape (6901, -0.52 DPS) [dungeon]; Resilient Cape (14400, -1.47 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 36.2 healing_power points (7.11 DPS) | yes | Beguiler Robes (7728, -0.56 DPS) [dungeon]; Pristine Gown (253961, -0.59 DPS, sim-verified) [crafted]; Wisdom's Leather Tunic (252511, -0.85 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.8 healing_power points (3.30 DPS) | yes | Mindthrust Bracers (1974, -0.10 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -0.99 DPS) [world_drop]; Nightsky Wristbands (6407, -1.07 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.4 healing_power points (3.81 DPS) | yes | Gloves of Old (9395, -0.39 DPS) [world_drop]; Hotshot Pilot's Gloves (9491, -0.68 DPS) [dungeon]; Wisdom's Leather Gloves (252499, -0.69 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 32.9 healing_power points (6.45 DPS) | yes | Skycaller's Leather Belt (252522, -1.72 DPS, sim-verified) [crafted]; Defiler's Lizardhide Girdle (20172, -2.98 DPS) [rep]; Defiler's Mail Girdle (20197, -2.98 DPS) [rep] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 37.2 healing_power points (7.30 DPS) | yes | Pristine Leggings (253987, -0.16 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.98 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.37 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.7 healing_power points (6.61 DPS) | yes | Acidic Walkers (9454, -2.16 DPS, sim-verified) [dungeon]; Wisdom's Leather Boots (252444, -3.00 DPS) [crafted]; Black Whelp Slippers (252424, -3.69 DPS) [crafted] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 15.2 healing_power points (2.98 DPS) | yes | Runemender's Seal (273031, -0.06 DPS) [world]; Ring of Power Regulation (273030, -0.27 DPS) [dungeon]; Snake Hoop (6750, -0.36 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.9 healing_power points (2.92 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Ring of Power Regulation (273030, -0.20 DPS) [dungeon]; Snake Hoop (6750, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.17 DPS) [dungeon]; Staff of the Friar (3415, -2.96 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -3.20 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.0 healing_power points (3.54 DPS) | yes | Orb of Mystic Insight (249394, -0.17 DPS, sim-verified) [crafted]; Heart of Agamaggan (6694, -0.59 DPS) [dungeon]; Seedcloud Buckler (6630, -1.11 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 428, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 0000000000000000-000000000000000000-5530003315510000)

Set DPS (verified): 235.2. Weights run: 6.8s. Verify run: 62.6s. 644 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.241, intellect=1.270 ± 0.018, spirit=0.667 ± 0.016, mp5=2.863 ± 0.024, crit=0.221 ± 0.017 per rating point (14 rating = 1%, 3.098 per %), spell_haste=1.381 ± 0.296

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 49.6 healing_power points (13.28 DPS) | yes | Whitemane's Chapeau (7720, -2.38 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.37 DPS) [world_drop]; Earthen Silk Hood (254015, -5.60 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.2 healing_power points (8.07 DPS) | yes | Triune Amulet (7722, -4.68 DPS) [dungeon]; Darkspear Warding Pendant (272074, -5.69 DPS) [vendor]; Necklace of Calisea (1714, -9.05 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 38.5 healing_power points (10.31 DPS) | yes | Earthen Silk Shoulders (254033, -4.06 DPS) [crafted]; Green Silken Shoulders (7057, -4.43 DPS) [crafted]; Sheepshear Mantle (13115, -9.14 DPS, sim-verified) [world_drop] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackforge Cape (6424, -2.28 DPS) [dungeon]; Long Silken Cloak (4326, -2.40 DPS) [crafted]; Darkspear Raider's Cloak (272077, -5.73 DPS, sim-verified) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.3 healing_power points (16.95 DPS) | yes | Dreamweave Vest (10021, -6.57 DPS) [crafted]; Robe of Power (7054, -7.70 DPS) [crafted]; Deathchill Armor (10764, -16.24 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.6 healing_power points (4.71 DPS) | yes | Radiant Silver Bracers (4545, -0.92 DPS) [quest]; Guardian Leather Bracers (4260, -1.06 DPS) [crafted]; Mindthrust Bracers (1974, -2.60 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.41 DPS) [crafted]; Prefect's Gauntlet (250569, -0.68 DPS) [crafted]; Dreamweave Gloves (10019, -1.25 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.2 healing_power points (9.68 DPS) | yes | Prefect's Belt (250559, -1.14 DPS, sim-verified) [crafted]; Mender's Leather Belt (252523, -1.48 DPS) [crafted]; Deathmage Sash (10771, -2.71 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 41.0 healing_power points (10.96 DPS) | yes | Stoneweaver Leggings (9407, -1.25 DPS, sim-verified) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -1.55 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.71 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 37.6 healing_power points (10.05 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.61 DPS) [crafted]; Skycaller's Leather Shoes (252532, -3.30 DPS) [crafted] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 17.6 healing_power points (4.71 DPS) | yes | Ring of Power Regulation (273030, -0.80 DPS) [dungeon]; Solid Gold Band (279845, -0.87 DPS) [quest]; Darkspear Signet (272070, -5.93 DPS, sim-verified) [vendor] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (235.2 DPS) | yes | Ring of Power Regulation (273030, -0.47 DPS) [dungeon]; Solid Gold Band (279845, -0.54 DPS) [quest]; Darkspear Signet (272070, -1.69 DPS, sim-verified) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -8.39 DPS) [dungeon]; Staff of the Friar (3415, -9.51 DPS) [dungeon]; Staff of Jordan (873, -24.08 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forcestone Buckler (17508, -1.06 DPS, sim-verified) [quest]; Battered Viking Shield (9403, -1.34 DPS) [dungeon]; Omega Orb (7749, -1.89 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Advisor's Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 644, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (tauren, 0000000000000000-000000000000000000-5530003315513151)

Set DPS (verified): 387.1. Weights run: 10.5s. Verify run: 85.9s. 819 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.164, intellect=0.657 ± 0.019, spirit=0.588 ± 0.018, mp5=2.357 ± 0.022, crit=0.291 ± 0.018 per rating point (14 rating = 1%, 4.073 per %), spell_haste=1.810 ± 0.349

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 49.3 healing_power points (28.91 DPS) | yes | Papal Fez (9431, -6.61 DPS, sim-verified) [dungeon]; Soothsayer's Headdress (17740, -6.88 DPS) [dungeon]; Holy Shroud (2721, -7.49 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.3 healing_power points (14.81 DPS) | yes | Mindburst Medallion (11196, -7.01 DPS, sim-verified) [quest]; Triune Amulet (7722, -7.71 DPS) [dungeon]; Horizon Choker (13085, -8.04 DPS) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 38.6 healing_power points (22.67 DPS) | yes | Nethergeld Shoulders (254049, -2.32 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -2.71 DPS) [crafted]; Mender's Mail Shoulders (252569, -2.71 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-verified (387.2 DPS) | yes | Spritecaster Cape (11623, -2.29 DPS, sim-verified) [dungeon]; Deep Woodlands Cloak (19121, -4.47 DPS) [quest]; Runecloth Cloak (13860, -6.61 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.7 healing_power points (36.78 DPS) | yes | Grimlok's Tribal Vestments (9415, -12.69 DPS, sim-verified) [dungeon]; Feathered Breastplate (8349, -13.50 DPS) [crafted]; Ghostweave Vest (14141, -13.54 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 27.5 healing_power points (16.15 DPS) | yes | Mender's Leather Bracers (252543, -0.24 DPS) [crafted]; Mender's Mail Bracers (252573, -0.24 DPS) [crafted]; Skycaller's Leather Bracers (252542, -5.87 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-verified (387.2 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -1.13 DPS) [crafted]; Earthenweave Gloves (254075, -5.71 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 37.9 healing_power points (22.22 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Waistcord (254081, -1.74 DPS) [crafted]; Earthenweave Cord (254077, -3.18 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 40.2 healing_power points (23.57 DPS) | yes | Stone Guard's Inscribed Legplates (220839, -1.31 DPS) [vendor]; Warchief Kilt (7760, -2.35 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -3.78 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 36.9 healing_power points (21.65 DPS) | yes | Mender's Mail Sabatons (252579, -0.99 DPS) [crafted]; Mender's Leather Boots (252472, -1.00 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.46 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 21.4 healing_power points (12.57 DPS) | yes | Solid Gold Band (279845, -4.39 DPS) [quest]; Philanthropist's Ring (281635, -4.78 DPS) [quest]; Electrocutioner Lagnut (9447, -4.88 DPS) [dungeon] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 16.5 healing_power points (9.68 DPS) | yes | Solid Gold Band (279845, -1.50 DPS) [quest]; Philanthropist's Ring (281635, -1.89 DPS) [quest]; Electrocutioner Lagnut (9447, -1.99 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (387.2 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (387.2 DPS) | yes | Darkspear Voodoo Seal (272061, -4.03 DPS, sim-verified) [vendor]; Ankh of Life (1713, -6.42 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (387.2 DPS) | yes | Resurgence Rod (17743, -20.12 DPS, sim-verified) [quest]; Charstone Dirk (17710, -27.22 DPS) [dungeon]; Dreamstaff (249454, -34.12 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Cassandra's Grace; shoulder: Living Shoulders; back: Battle Healer's Cloak; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Advisor's Ring; finger2: Darkspear Signet; trinket1: Dormant Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 819, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (tauren, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 481.5. Weights run: 9.5s. Verify run: 243.7s. 1628 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=0.657 ± 0.006, spirit=0.326 ± 0.006, mp5=1.249 ± 0.007, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.595 per %), spell_haste=not significant (0.032 ± 0.083)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 89.4 healing_power points (36.27 DPS) | yes | Warlord's Mail Skullcap (231669, -13.24 DPS) [vendor]; Cassandra's Grace (13102, -17.23 DPS) [world_drop]; Mooncloth Circlet (14140, -19.15 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.3 healing_power points (10.25 DPS) | yes | Archlight Talisman (15856, -3.68 DPS, sim-verified) [quest]; Mindburst Medallion (11196, -5.55 DPS) [quest]; Frostwolf Advisor's Pendant (19096, -5.55 DPS) [rep] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 40.6 healing_power points (16.46 DPS) | yes | Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Champion's Mail Epaulets (227166, -0.37 DPS) [vendor]; Eternity Pauldrons (279255, -0.81 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (481.5 DPS) | yes | Spritecaster Cape (11623, -3.53 DPS) [dungeon]; Deep Woodlands Cloak (19121, -4.34 DPS) [quest]; Runecloth Cloak (13860, -5.82 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.6 healing_power points (24.59 DPS) | yes | Warlord's Mail Chestguard (231668, -0.76 DPS) [vendor]; Legionnaire's Mail Chestguard (227171, -4.13 DPS) [vendor]; Red Dragonscale Breastplate (15047, -4.54 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.5 healing_power points (11.98 DPS) | yes | Mender's Leather Bracers (252543, -1.62 DPS) [crafted]; Mender's Mail Bracers (252573, -1.62 DPS) [crafted]; Nethergeld Cuffs (254061, -2.40 DPS, sim-verified) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (481.5 DPS) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Mender's Leather Gauntlets (252551, -0.77 DPS) [crafted]; Green Dragonscale Gauntlets (20296, -5.03 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 43.1 healing_power points (17.50 DPS) | yes | Mender's Mail Belt (252591, -2.13 DPS) [crafted]; Elderwild Waistcord (279252, -2.95 DPS) [crafted]; Mender's Leather Waistguard (252477, -5.18 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.4 healing_power points (20.86 DPS) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, +0.00 DPS) [vendor]; Elderwild Pants (279254, -0.69 DPS, sim-verified) [crafted] |
| feet | Swarmtender's Footpads (275609) (or Broodwatcher's Treaders (275621)) | Tailoring [crafted] | 43.0 healing_power points (17.44 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Broodwatcher's Treaders (275621, +0.00 DPS) [crafted]; Mooncloth Boots (15802, -2.25 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.6 healing_power points (9.96 DPS) | yes | Deep Rooted Ring (19109, -2.43 DPS) [world]; Philanthropist's Ring (281635, -4.57 DPS) [quest]; Band of the Unicorn (7553, -4.69 DPS) [world_drop] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 19.0 healing_power points (7.71 DPS) | yes | Philanthropist's Ring (281635, -2.32 DPS) [quest]; Band of the Unicorn (7553, -2.43 DPS) [world_drop]; Deep Rooted Ring (19109, -3.42 DPS, sim-verified) [world] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (481.5 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (481.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -1.45 DPS) [world_drop]; Darkspear Voodoo Seal (272061, -1.75 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (481.5 DPS) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -11.25 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -18.77 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Urgency (279249) | Enchanting [crafted] | sim-verified (481.5 DPS) | yes | Totem of Thunder (228176, -1.19 DPS, sim-verified) [crafted] |

**New at 60:** head: Living Crown; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Swarmtender's Footpads; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket2: Shard of the Splithooves; ranged: Totem of Urgency

No-known-source sample (15 of 1628, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (tauren, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 792.1. Weights run: 5.6s. Verify run: 122.5s. 1628 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.165, intellect=0.698 ± 0.032, spirit=0.248 ± 0.029, mp5=0.773 ± 0.033, crit=0.524 ± 0.032 per rating point (14 rating = 1%, 7.341 per %), spell_haste=5.943 ± 0.735

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 89.0 healing_power points (45.76 DPS) | yes | Warlord's Mail Skullcap (231669, -13.28 DPS) [vendor]; Mooncloth Circlet (14140, -17.95 DPS, sim-verified) [crafted]; Champion's Mail Skullcap (227168, -19.52 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.6 healing_power points (13.15 DPS) | yes | Archlight Talisman (15856, -3.34 DPS, sim-verified) [quest]; Mindburst Medallion (11196, -7.27 DPS) [quest]; Horizon Choker (13085, -7.61 DPS) [world_drop] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 39.7 healing_power points (20.42 DPS) | yes | Champion's Mail Epaulets (227166, +0.00 DPS) [vendor]; Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, -3.17 DPS, sim-verified) [vendor] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -4.27 DPS) [dungeon]; Deep Woodlands Cloak (19121, -4.99 DPS) [quest]; Runecloth Cloak (13860, -6.89 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.0 healing_power points (30.82 DPS) | yes | Warlord's Mail Chestguard (231668, +0.00 DPS) [vendor]; Legionnaire's Mail Chestguard (227171, -1.58 DPS) [vendor]; Red Dragonscale Breastplate (15047, -1.63 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.3 healing_power points (15.07 DPS) | yes | Nethergeld Cuffs (254061, -0.88 DPS, sim-verified) [crafted]; Mender's Leather Bracers (252543, -2.03 DPS) [crafted]; Mender's Mail Bracers (252573, -2.03 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Blood Guard's Mail Gloves (227169, -0.12 DPS) [vendor]; Green Dragonscale Gauntlets (20296, -0.92 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 43.0 healing_power points (22.10 DPS) | yes | Mender's Leather Waistguard (252477, -1.75 DPS, sim-verified) [crafted]; Mender's Mail Belt (252591, -2.38 DPS) [crafted]; Gilded Waistcord (254081, -3.97 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 57.9 healing_power points (29.76 DPS) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, +0.00 DPS) [vendor]; Elderwild Pants (279254, -5.93 DPS, sim-verified) [crafted] |
| feet | Swarmtender's Footpads (275609) (or Broodwatcher's Treaders (275621)) | Tailoring [crafted] | 43.0 healing_power points (22.10 DPS) | yes | General's Mail Boots (231667, +0.00 DPS) [pvp]; Broodwatcher's Treaders (275621, +0.00 DPS) [crafted]; Blood Guard's Mail Boots (227170, -2.89 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.4 healing_power points (12.53 DPS) | yes | Advisor's Ring (19518, -3.74 DPS) [rep]; Philanthropist's Ring (281635, -5.59 DPS) [quest] |
| finger2 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 19.0 healing_power points (9.75 DPS) | yes | Advisor's Ring (19518, -0.51 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -2.82 DPS) [quest] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (792.1 DPS) | yes | Ankh of Life (1713, -5.85 DPS, sim-verified) [world_drop]; Shard of the Splithooves (10659, -6.87 DPS) [quest]; Darkspear Voodoo Seal (272061, -7.26 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Charstone Dirk (17710, -2.76 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -26.71 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Urgency (279249) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -7.14 DPS, sim-verified) [crafted] |

**New at 60:** head: Living Crown; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Swarmtender's Footpads; finger1: Maiden's Circle; finger2: Deep Rooted Ring; ranged: Totem of Urgency

No-known-source sample (15 of 1628, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

