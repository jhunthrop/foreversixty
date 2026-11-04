# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 50320100000000000-00000000000000000-00000000000000000)

Set DPS (verified): 51.4. Weights run: 7.0s. Verify run: 52.7s. 265 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=0.393 ± 0.001, spirit=0.218 ± 0.001, mp5=1.041 ± 0.003, crit=0.025 ± 0.001 per rating point (14 rating = 1%, 0.352 per %), spell_haste=0.754 ± 0.024

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (3.80 DPS) | yes | Pristine Circlet (253949, -0.62 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.72 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.81 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.4 healing_power points (0.08 DPS) | yes | Tarnished Locket (279870, -1.55 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 healing_power points (1.54 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.69 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.82 DPS) [crafted]; Slime-encrusted Pads (6461, -0.98 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.31 DPS) [crafted]; Black Whelp Cloak (7283, -0.34 DPS) [crafted]; Caretaker's Cape (20428, -0.70 DPS, sim-verified) [rep] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Acolyte's Chain Shirt (250491, -0.19 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.36 DPS) [crafted]; Bloody Apron (6226, -0.84 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 5.1 healing_power points (0.92 DPS) | yes | Owl Bracers (4796, -0.56 DPS) [vendor]; Bright Bracers (3647, -0.64 DPS) [world_drop]; Rift Bracers (5943, -3.50 DPS, sim-verified) [dungeon] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Gloves (252499, -0.07 DPS) [crafted]; Pristine Gloves (253913, -0.09 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.90 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 12.6 healing_power points (2.27 DPS) | yes | Acolyte's Chain Belt (250516, -0.39 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.54 DPS) [crafted]; Novice Ardent's Sash (253887, -0.86 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 23.8 healing_power points (4.29 DPS) | yes | Wisdom's Leather Pants (252503, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.90 DPS) [crafted] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 13.0 healing_power points (2.34 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.43 DPS) [crafted]; Black Whelp Slippers (252424, -0.43 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.1 healing_power points (1.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.85 DPS) [dungeon]; Furlfeather Ring (286745, -0.92 DPS) [world]; Band of Purification (12996, -1.04 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 healing_power points (1.05 DPS) | yes | Lavishly Jeweled Ring (1156, -0.39 DPS, sim-verified) [dungeon]; Furlfeather Ring (286745, -0.68 DPS) [world]; Band of Purification (12996, -0.81 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monastic Hammer (270005, -6.41 DPS, sim-verified) [quest]; Rhahk'Zor's Hammer (5187, -7.00 DPS) [dungeon]; Thornroot Club (5587, -7.37 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 265, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs

### Band 30 (human, 50320215011100000-00000000000000000-00000000000000000)

Set DPS (verified): 122.4. Weights run: 4.1s. Verify run: 30.8s. 479 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.050, intellect=0.726 ± 0.014, spirit=0.550 ± 0.011, mp5=1.476 ± 0.025, crit=0.535 ± 0.015 per rating point (14 rating = 1%, 7.487 per %), spell_haste=1.681 ± 0.115

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.3 healing_power points (9.27 DPS) | yes | Acolyte's Chain Helm (250501, -3.05 DPS, sim-verified) [crafted]; Sunblaze Coif (5819, -3.65 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -3.65 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.1 healing_power points (1.30 DPS) | yes | Crystal Starfire Medallion (5003, -0.14 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.19 DPS) [world]; Seal of the Expedition (271920, -0.38 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.5 healing_power points (3.97 DPS) | yes | Death Speaker Mantle (6685, -0.40 DPS) [dungeon]; Magician's Mantle (12998, -1.02 DPS) [world_drop]; Invoker's Mantle (215365, -1.25 DPS) [crafted] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.54 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -1.09 DPS) [quest]; Prelacy Cape (7004, -1.67 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 28.0 healing_power points (7.15 DPS) | yes | Pristine Gown (253961, -0.22 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Shirt (250521, -0.23 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.4 healing_power points (3.43 DPS) | yes | Mindthrust Bracers (1974, -0.57 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.90 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.95 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.2 healing_power points (4.39 DPS) | yes | Naga Battle Gloves (888, -0.84 DPS, sim-verified) [dungeon]; Acolyte's Gloves (250511, -0.91 DPS) [crafted]; Pristine Gloves (253913, -1.02 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Prefect's Belt (250559, -0.38 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -2.58 DPS) [crafted]; Belt of Arugal (6392, -3.25 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.03 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.46 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.86 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 27.3 healing_power points (6.97 DPS) | yes | Wisdom's Leather Boots (252444, -3.23 DPS) [crafted]; Acolyte's Boots (250506, -3.35 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.65 DPS) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.8 healing_power points (3.28 DPS) | yes | Ring of Power Regulation (273030, -0.62 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.74 DPS) [rep]; Snake Hoop (6750, -0.82 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 12.8 healing_power points (3.28 DPS) | yes | Ring of Power Regulation (273030, -0.62 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.74 DPS) [rep]; Snake Hoop (6750, -0.82 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taskmaster Axe (5194, -7.23 DPS, sim-verified) [dungeon]; Crested Scepter (3414, -8.83 DPS) [dungeon]; Pysan's Old Greatsword (1975, -9.96 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.3 healing_power points (4.16 DPS) | yes | Eye of Paleth (2943, -0.84 DPS) [quest]; Heart of Agamaggan (6694, -0.95 DPS, sim-verified) [dungeon]; Orb of Mystic Insight (249394, -1.26 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 479, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 50320215012101251-00000000000000000-00000000000000000)

Set DPS (verified): 206.8. Weights run: 5.0s. Verify run: 57.4s. 657 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.170, intellect=0.160 ± 0.015, spirit=0.159 ± 0.014, mp5=1.329 ± 0.037, crit=0.348 ± 0.018 per rating point (14 rating = 1%, 4.876 per %), spell_haste=4.503 ± 0.252

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 34.0 healing_power points (6.65 DPS) | yes | Papal Fez (9431, -0.79 DPS, sim-verified) [dungeon]; Acolyte's Chain Helm (250501, -1.75 DPS) [crafted]; Earthen Silk Hood (254015, -2.03 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.3 healing_power points (4.17 DPS) | yes | Triune Amulet (7722, -2.44 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -3.73 DPS) [dungeon]; Amberglow Talisman (10824, -3.86 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.1 healing_power points (4.72 DPS) | yes | Earthen Silk Shoulders (254033, -1.14 DPS, sim-verified) [crafted]; Shining Mithril Pauldrons (250541, -2.56 DPS) [crafted]; Bloodmage Mantle (7684, -2.67 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (206.8 DPS) | yes | Glowing Thresher Cape (6901, -2.06 DPS) [dungeon]; Icy Cloak (4327, -2.12 DPS) [crafted]; Prelacy Cape (7004, -2.22 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 59.3 healing_power points (11.61 DPS) | yes | Civinad Robes (9623, -6.12 DPS, sim-verified) [quest]; Deathchill Armor (10764, -6.94 DPS) [dungeon]; Robe of the Magi (1716, -6.96 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.0 healing_power points (2.54 DPS) | yes | Arcane Runed Bracers (4744, -0.78 DPS) [quest]; Earthen Silk Cuffs (254019, -1.09 DPS, sim-verified) [crafted]; Condor Bracers (15864, -1.17 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (206.8 DPS) | yes | Gilded Handwraps (254021, -0.41 DPS, sim-verified) [crafted]; Mender's Leather Gloves (252530, -0.70 DPS) [crafted]; Prefect's Gauntlet (250569, -0.76 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 24.2 healing_power points (4.75 DPS) | yes | Mender's Leather Belt (252523, -0.65 DPS) [crafted]; Prefect's Belt (250559, -0.74 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20098, -1.88 DPS) [rep] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 27.2 healing_power points (5.33 DPS) | yes | Pristine Leggings (253987, -0.64 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.77 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.99 DPS, sim-verified) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 29.4 healing_power points (5.77 DPS) | yes | Mender's Leather Shoes (252533, -0.36 DPS, sim-verified) [crafted]; Mender's Mail Boots (252565, -0.63 DPS) [crafted]; Gilded Slippers (254001, -1.51 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 13.0 healing_power points (2.54 DPS) | yes | Electrocutioner Lagnut (9447, -0.56 DPS) [dungeon]; Runemender's Seal (273031, -0.56 DPS) [world]; Ring of Power Regulation (273030, -0.59 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 healing_power points (2.12 DPS) | yes | Electrocutioner Lagnut (9447, -0.14 DPS) [dungeon]; Runemender's Seal (273031, -0.14 DPS) [world]; Ring of Power Regulation (273030, -0.16 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (206.8 DPS) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (206.8 DPS) | yes | Death Speaker Scepter (2816, -5.86 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -9.60 DPS) [dungeon]; Crested Scepter (3414, -12.33 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-verified (206.8 DPS) | yes | Battered Viking Shield (9403, -1.58 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -1.69 DPS) [dungeon]; Heart of Agamaggan (6694, -1.77 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Gloves of the Greatfather; waist: Gilded Cord; legs: Warchief Kilt; feet: Prefect's Boots; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 657, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (human, 50320215012101251-50320000000000000-00000000000000000)

Set DPS (verified): 249.6. Weights run: 8.6s. Verify run: 74.4s. 840 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.015, intellect=0.902 ± 0.014, spirit=0.910 ± 0.007, mp5=3.425 ± 0.015, crit=0.514 ± 0.016 per rating point (14 rating = 1%, 7.190 per %), spell_haste=not significant (-0.016 ± 0.135)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 52.2 healing_power points (11.19 DPS) | yes | Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Soothsayer's Headdress (17740, -1.80 DPS) [dungeon]; Papal Fez (9431, -1.92 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.2 healing_power points (5.84 DPS) | yes | Mindburst Medallion (11196, -2.41 DPS) [quest]; Triune Amulet (7722, -2.75 DPS) [dungeon]; Horizon Choker (13085, -3.61 DPS, sim-verified) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 42.8 healing_power points (9.18 DPS) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -0.48 DPS, sim-verified) [crafted]; Nethergeld Shoulders (254049, -0.89 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.51 DPS, sim-verified) [dungeon]; Expeditionary's Cape (271921, -2.20 DPS) [rep]; Runecloth Cloak (13860, -2.41 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.3 healing_power points (14.00 DPS) | yes | Knight's Imbued Armor (220813, -2.33 DPS) [vendor]; Feathered Breastplate (8349, -3.31 DPS) [crafted]; Ghostweave Vest (14141, -6.74 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 30.9 healing_power points (6.62 DPS) | yes | Mender's Leather Bracers (252543, -0.02 DPS) [crafted]; Mender's Mail Bracers (252573, -0.02 DPS) [crafted]; Prefect's Wristguards (250584, -0.14 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.43 DPS) [crafted]; Prefect's Gloves (250554, -0.51 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 40.8 healing_power points (8.75 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.39 DPS) [crafted]; Gilded Waistcord (254081, -0.79 DPS) [crafted] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 54.4 healing_power points (11.67 DPS) | yes | Senior Designer's Pantaloons (11841, -2.05 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Kilt (9474, -2.60 DPS) [dungeon]; Warchief Kilt (7760, -2.63 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 39.1 healing_power points (8.39 DPS) | yes | Mender's Leather Boots (252472, -0.26 DPS) [crafted]; Mender's Mail Sabatons (252579, -0.26 DPS) [crafted]; Mender's Leather Shoes (252533, -0.52 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 25.7 healing_power points (5.51 DPS) | yes | Ring of Power Regulation (273030, -2.02 DPS) [dungeon]; Electrocutioner Lagnut (9447, -2.21 DPS) [dungeon]; Runemender's Seal (273031, -2.21 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 16.3 healing_power points (3.50 DPS) | yes | Ring of Power Regulation (273030, -0.01 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.21 DPS) [dungeon]; Runemender's Seal (273031, -0.21 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, +0.00 DPS) [rep]; Evonice's Landin' Pilla (18951, -0.39 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (249.7 DPS) | yes | Enriched Seal (271903, -1.29 DPS, sim-verified) [rep]; Evonice's Landin' Pilla (18951, -1.91 DPS) [quest] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inventor's Focal Sword (17719, -7.31 DPS) [dungeon]; Galgann's Firehammer (9419, -10.14 DPS) [dungeon]; Death Speaker Scepter (2816, -10.22 DPS, sim-verified) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beacon of Hope (9393, -0.14 DPS) [dungeon]; Battered Viking Shield (9403, -0.90 DPS) [dungeon]; Orb of Lorica (11262, -1.13 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** head: Cassandra's Grace; shoulder: Living Shoulders; back: Caretaker's Cape; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Gilded Sandals; finger1: Lorekeeper's Ring; finger2: Eye of Adaegus; trinket2: Uther's Strength; off_hand: Twisting Essence Jar

No-known-source sample (15 of 840, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 393.1. Weights run: 6.2s. Verify run: 240.0s. 1646 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.040, intellect=0.871 ± 0.022, spirit=0.996 ± 0.019, mp5=2.918 ± 0.026, crit=1.062 ± 0.032 per rating point (14 rating = 1%, 14.873 per %), spell_haste=5.395 ± 0.408

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 102.6 healing_power points (37.57 DPS) | yes | Field Marshal's Lamellar Helmet (231640, -5.38 DPS) [vendor]; Lieutenant Commander's Lamellar Helmet (227149, -8.92 DPS) [vendor]; Mooncloth Circlet (14140, -13.37 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.0 healing_power points (9.87 DPS) | yes | Archlight Talisman (15856, -1.92 DPS, sim-verified) [quest]; Stormpike Sage's Pendant (19098, -2.41 DPS) [rep]; Mindburst Medallion (11196, -3.94 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Dawnbringer Shoulders (12625, -15.65 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -3.21 DPS) [dungeon]; Expeditionary's Cape (271907, -5.15 DPS) [rep] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 85.1 healing_power points (31.14 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Knight-Captain's Lamellar Chestplate (227151, -3.54 DPS) [vendor]; Stormcloth Vest (10020, -6.51 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+17.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Mender's Leather Gauntlets (252551, -18.58 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -15.89 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legguards (231639, -7.26 DPS) [vendor]; Knight-Captain's Lamellar Legguards (227150, -9.83 DPS) [vendor]; Sentinel's Silk Leggings (237815, -11.30 DPS, sim-verified) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -16.93 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 30.1 healing_power points (11.01 DPS) | yes | Deep Rooted Ring (19109, -3.43 DPS) [world]; Band of the Hierophant (13096, -3.81 DPS) [world_drop]; Eye of Adaegus (5266, -4.72 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 25.7 healing_power points (9.40 DPS) | yes | Deep Rooted Ring (19109, -1.82 DPS) [world]; Band of the Hierophant (13096, -2.20 DPS) [world_drop]; Eye of Adaegus (5266, -3.11 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+6.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (393.0 DPS) | yes | Enriched Seal (271903, -1.25 DPS) [rep]; Shard of the Splithooves (10659, -1.49 DPS, sim-verified) [quest]; Ankh of Life (1713, -2.21 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Righteousness (7721, -3.90 DPS, sim-verified) [dungeon]; Sageblade (22383, -4.41 DPS) [crafted] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Evergreen Shield (279262, -3.94 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Caretaker's Cape; chest: Breastplate of Salvation; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; trinket1: Dormant Heart of the Mountain; main_hand: Hand of Edward the Odd; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1646, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 701.5. Weights run: 4.2s. Verify run: 213.0s. 1646 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.443, intellect=0.991 ± 0.065, spirit=0.732 ± 0.062, mp5=1.639 ± 0.066, crit=0.991 ± 0.071 per rating point (14 rating = 1%, 13.876 per %), spell_haste=6.084 ± 1.303

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 101.1 healing_power points (45.65 DPS) | yes | Field Marshal's Lamellar Helmet (231640, -6.99 DPS) [vendor]; Lieutenant Commander's Lamellar Helmet (227149, -12.22 DPS) [vendor]; Mooncloth Circlet (14140, -17.21 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.9 healing_power points (12.61 DPS) | yes | Archlight Talisman (15856, -3.04 DPS, sim-verified) [quest]; Horizon Choker (13085, -5.02 DPS) [world_drop]; Stormpike Sage's Pendant (19098, -5.18 DPS) [rep] |
| shoulder | Dawnbringer Shoulders (12625) | Blacksmithing [crafted] | 51.3 healing_power points (23.18 DPS) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Theramore Mantle (271957, -0.20 DPS) [rep] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -3.39 DPS) [dungeon]; Expeditionary's Cape (271907, -5.24 DPS) [rep] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 78.0 healing_power points (35.24 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Knight-Captain's Lamellar Chestplate (227151, -3.16 DPS) [vendor]; Stormcloth Vest (10020, -7.54 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 35.1 healing_power points (15.83 DPS) | yes | Prefect's Wristguards (250584, -1.91 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -2.01 DPS) [crafted]; Mender's Leather Bracers (252543, -2.13 DPS) [crafted] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Mender's Mail Gauntlets (252587, -1.92 DPS) [crafted]; Mender's Leather Gauntlets (252551, -2.85 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 51.5 healing_power points (23.24 DPS) | yes | Elderwild Waistcord (279252, -3.42 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -4.32 DPS) [crafted]; Mender's Mail Belt (252591, -4.32 DPS) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legguards (231639, -6.94 DPS) [vendor]; Ironfeather Leggings (252486, -10.35 DPS, sim-verified) [crafted]; Knight-Captain's Lamellar Legguards (227150, -12.65 DPS) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 47.4 healing_power points (21.40 DPS) | yes | Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Marshal's Lamellar Greaves (231646, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, -3.32 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 29.3 healing_power points (13.25 DPS) | yes | Lorekeeper's Ring (19522, -3.96 DPS) [rep]; Band of the Hierophant (13096, -5.13 DPS) [world_drop] |
| finger2 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 21.9 healing_power points (9.89 DPS) | yes | Lorekeeper's Ring (19522, -0.61 DPS) [rep]; Band of the Hierophant (13096, -1.78 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+8.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (701.7 DPS) | yes | Enriched Seal (271903, -2.76 DPS, sim-verified) [rep]; Shard of the Splithooves (10659, -3.69 DPS) [quest]; Ankh of Life (1713, -4.16 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Righteousness (7721, -4.67 DPS) [dungeon]; Sageblade (22383, -6.05 DPS, sim-verified) [crafted] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Evergreen Shield (279262, -1.38 DPS, sim-verified) [crafted] |
| ranged | Libram of Infusion (279248) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, -1.08 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; shoulder: Dawnbringer Shoulders; back: Caretaker's Cape; chest: Breastplate of Salvation; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Maiden's Circle; finger2: Deep Rooted Ring; trinket1: Dormant Heart of the Mountain; main_hand: Hand of Edward the Odd; off_hand: Lei of the Lifegiver; ranged: Libram of Infusion

No-known-source sample (15 of 1646, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (undead, 50320100000000000-00000000000000000-00000000000000000)

Set DPS (verified): 51.4. Weights run: 7.0s. Verify run: 52.3s. 252 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.040, intellect=0.815 ± 0.005, spirit=0.691 ± 0.004, mp5=2.985 ± 0.005, crit=0.085 ± 0.003 per rating point (14 rating = 1%, 1.191 per %), spell_haste=4.539 ± 0.124

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (3.68 DPS) | yes | Pristine Circlet (253949, -0.61 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.70 DPS) [crafted]; Crusader's Silvered Chain Helm (250532, -1.75 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.1 healing_power points (0.36 DPS) | yes | Scholarly Pendant (277203, -0.17 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.3 healing_power points (2.16 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.71 DPS) [crafted]; Serpent's Shoulders (5404, -1.31 DPS) [dungeon]; Slime-encrusted Pads (6461, -1.49 DPS, sim-verified) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.91 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -1.04 DPS) [crafted]; Sanguine Cape (14376, -1.25 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 21.1 healing_power points (3.70 DPS) | yes | Acolyte's Chain Shirt (250491, -0.19 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.35 DPS) [crafted]; Corsair's Overshirt (5202, -0.81 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 13.0 healing_power points (2.28 DPS) | yes | Owl Bracers (4796, -1.57 DPS) [vendor]; Garrison Cuffs (270003, -1.69 DPS) [quest]; Drakewing Bands (12999, -1.79 DPS, sim-verified) [world_drop] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.1 healing_power points (2.47 DPS) | yes | Wisdom's Leather Gloves (252499, -0.14 DPS) [crafted]; Pristine Gloves (253913, -0.25 DPS, sim-verified) [crafted]; Stormrider's Leather Gloves (252498, -1.02 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.3 healing_power points (2.50 DPS) | yes | Acolyte's Chain Belt (250516, -0.41 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.53 DPS) [crafted]; Novice Ardent's Sash (253887, -0.66 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | sim-verified (51.4 DPS) | yes | Wisdom's Leather Pants (252503, -0.08 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.59 DPS) [crafted] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.1 healing_power points (2.64 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.49 DPS) [crafted]; Black Whelp Slippers (252424, -0.49 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 11.0 healing_power points (1.92 DPS) | yes | Band of Purification (12996, -1.20 DPS) [world_drop]; Loop of Sacrifice (281673, -1.21 DPS) [quest]; Volcanic Rock Ring (12053, -1.49 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.9 healing_power points (0.86 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop]; Band of Purification (12996, -0.67 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | 50.5 healing_power points (8.85 DPS) | yes | The Stitcher (279874, -2.77 DPS, sim-verified) [quest]; Pathfinder's Clearway (277288, -5.76 DPS) [quest]; Rhahk'Zor's Hammer (5187, -7.45 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 252, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 20430 Legionnaire's Sword; 23173 Abomination Skin Leggings; 205218 Libram of Discovery

### Band 30 (undead, 50320215011100000-00000000000000000-00000000000000000)

Set DPS (verified): 118.6. Weights run: 4.1s. Verify run: 35.5s. 466 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.050, intellect=0.726 ± 0.014, spirit=0.550 ± 0.011, mp5=1.476 ± 0.025, crit=0.535 ± 0.015 per rating point (14 rating = 1%, 7.487 per %), spell_haste=1.681 ± 0.115

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.3 healing_power points (9.27 DPS) | yes | Acolyte's Chain Helm (250501, -2.71 DPS, sim-verified) [crafted]; Sunblaze Coif (5819, -3.65 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -3.65 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.1 healing_power points (1.30 DPS) | yes | Crystal Starfire Medallion (5003, -0.19 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.19 DPS) [world]; Darkspear Warding Pendant (272075, -0.38 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.5 healing_power points (3.97 DPS) | yes | Death Speaker Mantle (6685, -0.40 DPS) [dungeon]; Magician's Mantle (12998, -1.02 DPS) [world_drop]; Invoker's Mantle (215365, -1.25 DPS) [crafted] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.52 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -1.98 DPS) [vendor]; Windsong Drape (15468, -2.04 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 28.0 healing_power points (7.15 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, -0.23 DPS) [crafted]; Pristine Gown (253961, -0.42 DPS, sim-verified) [crafted]; Wisdom's Leather Tunic (252511, -0.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.4 healing_power points (3.43 DPS) | yes | Mindthrust Bracers (1974, -0.73 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.90 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.95 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.2 healing_power points (4.39 DPS) | yes | Naga Battle Gloves (888, -0.85 DPS, sim-verified) [dungeon]; Acolyte's Gloves (250511, -0.91 DPS) [crafted]; Pristine Gloves (253913, -1.02 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Prefect's Belt (250559, -0.15 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -2.58 DPS) [crafted]; Belt of Arugal (6392, -3.25 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.03 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.17 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.86 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 27.3 healing_power points (6.97 DPS) | yes | Acolyte's Boots (250506, -3.10 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -3.23 DPS) [crafted]; Acidic Walkers (9454, -3.65 DPS) [dungeon] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 13.7 healing_power points (3.51 DPS) | yes | Runemender's Seal (273031, -0.23 DPS) [world]; Ring of Power Regulation (273030, -0.85 DPS) [dungeon]; Advisor's Ring (19521, -0.97 DPS) [rep] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.8 healing_power points (3.28 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Ring of Power Regulation (273030, -0.62 DPS) [dungeon]; Advisor's Ring (19521, -0.74 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taskmaster Axe (5194, -6.47 DPS, sim-verified) [dungeon]; Crested Scepter (3414, -8.83 DPS) [dungeon]; The Stitcher (279874, -9.71 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.3 healing_power points (4.16 DPS) | yes | Heart of Agamaggan (6694, -0.69 DPS, sim-verified) [dungeon]; Orb of Mystic Insight (249394, -1.26 DPS) [crafted]; Orb of Mistmantle (13031, -1.30 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 466, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (undead, 50320215012101251-00000000000000000-00000000000000000)

Set DPS (verified): 204.5. Weights run: 5.0s. Verify run: 58.7s. 641 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.170, intellect=0.160 ± 0.015, spirit=0.159 ± 0.014, mp5=1.329 ± 0.037, crit=0.348 ± 0.018 per rating point (14 rating = 1%, 4.876 per %), spell_haste=4.503 ± 0.252

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 34.0 healing_power points (6.65 DPS) | yes | Papal Fez (9431, -0.62 DPS, sim-verified) [dungeon]; Acolyte's Chain Helm (250501, -1.75 DPS) [crafted]; Earthen Silk Hood (254015, -2.03 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.3 healing_power points (4.17 DPS) | yes | Triune Amulet (7722, -2.57 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -3.73 DPS) [dungeon]; Amberglow Talisman (10824, -3.86 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.1 healing_power points (4.72 DPS) | yes | Earthen Silk Shoulders (254033, -1.35 DPS, sim-verified) [crafted]; Shining Mithril Pauldrons (250541, -2.56 DPS) [crafted]; Naga Priestess's Mantle (279889, -2.57 DPS) [quest] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (204.5 DPS) | yes | Glowing Thresher Cape (6901, -2.06 DPS) [dungeon]; Icy Cloak (4327, -2.12 DPS) [crafted]; Long Silken Cloak (4326, -2.35 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 59.3 healing_power points (11.61 DPS) | yes | Civinad Robes (9623, -6.01 DPS, sim-verified) [quest]; Deathchill Armor (10764, -6.94 DPS) [dungeon]; Robe of the Magi (1716, -6.96 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.0 healing_power points (2.54 DPS) | yes | Condor Bracers (15864, -1.17 DPS) [quest]; Guardian Leather Bracers (4260, -1.18 DPS) [crafted]; Earthen Silk Cuffs (254019, -1.57 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (204.5 DPS) | yes | Gilded Handwraps (254021, -0.22 DPS, sim-verified) [crafted]; Mender's Leather Gloves (252530, -0.70 DPS) [crafted]; Prefect's Gauntlet (250569, -0.76 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 24.2 healing_power points (4.75 DPS) | yes | Mender's Leather Belt (252523, -0.65 DPS) [crafted]; Prefect's Belt (250559, -0.76 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20166, -1.88 DPS) [rep] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 27.2 healing_power points (5.33 DPS) | yes | Pristine Leggings (253987, -0.64 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.77 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -1.32 DPS, sim-verified) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 29.4 healing_power points (5.77 DPS) | yes | Mender's Leather Shoes (252533, -0.46 DPS, sim-verified) [crafted]; Mender's Mail Boots (252565, -0.63 DPS) [crafted]; Gilded Slippers (254001, -1.51 DPS) [crafted] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 13.0 healing_power points (2.54 DPS) | yes | Philanthropist's Ring (281635, -0.43 DPS) [quest]; Electrocutioner Lagnut (9447, -0.56 DPS) [dungeon]; Runemender's Seal (273031, -0.56 DPS) [world] |
| finger2 | Solid Gold Band (279845) | Opportunistic Education [quest] | 11.8 healing_power points (2.31 DPS) | yes | Philanthropist's Ring (281635, -0.28 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Runemender's Seal (273031, -0.33 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (204.5 DPS) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (204.5 DPS) | yes | Death Speaker Scepter (2816, -6.08 DPS, sim-verified) [dungeon]; Taskmaster Axe (5194, -9.60 DPS) [dungeon]; Crested Scepter (3414, -12.33 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-verified (204.5 DPS) | yes | Forcestone Buckler (17508, -0.59 DPS, sim-verified) [quest]; Heart of Agamaggan (6694, -1.50 DPS) [dungeon]; Battered Viking Shield (9403, -1.58 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Gloves of the Greatfather; waist: Gilded Cord; legs: Warchief Kilt; feet: Prefect's Boots; finger1: Advisor's Ring; finger2: Solid Gold Band; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 641, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (undead, 50320215012101251-50320000000000000-00000000000000000)

Set DPS (verified): 241.0. Weights run: 8.6s. Verify run: 73.3s. 837 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.015, intellect=0.902 ± 0.014, spirit=0.910 ± 0.007, mp5=3.425 ± 0.015, crit=0.514 ± 0.016 per rating point (14 rating = 1%, 7.190 per %), spell_haste=not significant (-0.016 ± 0.135)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 52.2 healing_power points (11.19 DPS) | yes | Soothsayer's Headdress (17740, -1.80 DPS) [dungeon]; Papal Fez (9431, -2.10 DPS, sim-verified) [dungeon]; Stormcloth Headband (10032, -2.16 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.2 healing_power points (5.84 DPS) | yes | Mindburst Medallion (11196, -2.41 DPS) [quest]; Triune Amulet (7722, -2.75 DPS) [dungeon]; Horizon Choker (13085, -4.42 DPS, sim-verified) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 42.8 healing_power points (9.18 DPS) | yes | Ironfeather Shoulders (15067, -0.54 DPS) [crafted]; Nethergeld Shoulders (254049, -0.89 DPS) [crafted]; Mender's Leather Shoulders (252538, -0.93 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.77 DPS, sim-verified) [dungeon]; Deep Woodlands Cloak (19121, -1.57 DPS) [quest]; Darkspear Raider's Cloak (272076, -2.20 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.3 healing_power points (14.00 DPS) | yes | Feathered Breastplate (8349, -3.31 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -4.08 DPS) [dungeon]; Ghostweave Vest (14141, -6.79 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 30.9 healing_power points (6.62 DPS) | yes | Mender's Leather Bracers (252543, -0.02 DPS) [crafted]; Mender's Mail Bracers (252573, -0.02 DPS) [crafted]; Prefect's Wristguards (250584, -0.14 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.43 DPS) [crafted]; Prefect's Gloves (250554, -0.51 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 40.8 healing_power points (8.75 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.39 DPS) [crafted]; Gilded Waistcord (254081, -0.79 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 47.7 healing_power points (10.23 DPS) | yes | Warchief Kilt (7760, -1.19 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -1.68 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.95 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 39.1 healing_power points (8.39 DPS) | yes | Mender's Mail Sabatons (252579, -0.26 DPS) [crafted]; Mender's Leather Boots (252472, -0.35 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -0.52 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 25.7 healing_power points (5.51 DPS) | yes | Eye of Adaegus (5266, -2.01 DPS) [world_drop]; Ring of Power Regulation (273030, -2.02 DPS) [dungeon]; Solid Gold Band (279845, -2.18 DPS) [quest] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 24.0 healing_power points (5.14 DPS) | yes | Eye of Adaegus (5266, -1.64 DPS) [world_drop]; Ring of Power Regulation (273030, -1.65 DPS) [dungeon]; Solid Gold Band (279845, -1.81 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -1.72 DPS) [quest]; Alchemist's Stone (13503, -2.89 DPS) [crafted]; Ankh of Life (1713, -5.45 DPS, sim-verified) [world_drop] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (241.0 DPS) | yes | Ankh of Life (1713, -0.70 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -1.91 DPS) [quest]; Alchemist's Stone (13503, -3.08 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Inventor's Focal Sword (17719, -7.31 DPS) [dungeon]; Galgann's Firehammer (9419, -10.14 DPS) [dungeon]; Death Speaker Scepter (2816, -10.23 DPS, sim-verified) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beacon of Hope (9393, -0.14 DPS) [dungeon]; Forcestone Buckler (17508, -0.21 DPS) [quest]; Battered Viking Shield (9403, -0.90 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Cassandra's Grace; shoulder: Living Shoulders; back: Battle Healer's Cloak; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Advisor's Ring; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; off_hand: Twisting Essence Jar

No-known-source sample (15 of 837, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 381.0. Weights run: 6.2s. Verify run: 240.5s. 1670 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.040, intellect=0.871 ± 0.022, spirit=0.996 ± 0.019, mp5=2.918 ± 0.026, crit=1.062 ± 0.032 per rating point (14 rating = 1%, 14.873 per %), spell_haste=5.395 ± 0.408

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 102.6 healing_power points (37.57 DPS) | yes | Mooncloth Circlet (14140, -12.43 DPS, sim-verified) [crafted]; Whitesoul Helm (12633, -14.51 DPS) [crafted]; Earthenweave Crown (254137, -16.37 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.0 healing_power points (9.87 DPS) | yes | Archlight Talisman (15856, -2.27 DPS, sim-verified) [quest]; Frostwolf Advisor's Pendant (19096, -2.41 DPS) [rep]; Mindburst Medallion (11196, -3.94 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Darkspear Mantle (272107, +0.00 DPS) [vendor]; Dawnbringer Shoulders (12625, -15.66 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -3.21 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -5.15 DPS) [vendor]; Deep Woodlands Cloak (19121, -5.17 DPS) [quest] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 85.1 healing_power points (31.14 DPS) | yes | Stormcloth Vest (10020, -4.91 DPS, sim-verified) [crafted]; Mooncloth Vest (14138, -8.32 DPS) [crafted]; Ironfeather Breastplate (15066, -9.42 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+17.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Mender's Leather Gauntlets (252551, -18.12 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -15.40 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironfeather Leggings (252486, -10.07 DPS, sim-verified) [crafted]; Elderwild Pants (279254, -14.36 DPS) [crafted]; Devout Skirt (16694, -15.80 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Broodwatcher's Treaders (275621, +0.00 DPS) [crafted]; Mooncloth Boots (15802, -15.97 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 30.1 healing_power points (11.01 DPS) | yes | Deep Rooted Ring (19109, -3.43 DPS) [world]; Darkspear Signet (272069, -3.53 DPS) [vendor] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Signet (272068, -0.65 DPS, sim-verified) [vendor]; Deep Rooted Ring (19109, -1.82 DPS) [world] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Ankh of Life (1713, -13.15 DPS, sim-verified) [world_drop] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -0.18 DPS) [quest]; Ankh of Life (1713, -0.91 DPS, sim-verified) [world_drop]; Darkspear Voodoo Seal (272061, -1.25 DPS) [vendor] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Righteousness (7721, -3.04 DPS, sim-verified) [dungeon]; Sageblade (22383, -4.41 DPS) [crafted] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Evergreen Shield (279262, -3.70 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Battle Healer's Cloak; chest: Breastplate of Salvation; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket1: Dormant Heart of the Mountain; main_hand: Hand of Edward the Odd; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1670, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 688.1. Weights run: 4.2s. Verify run: 211.4s. 1670 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.443, intellect=0.991 ± 0.065, spirit=0.732 ± 0.062, mp5=1.639 ± 0.066, crit=0.991 ± 0.071 per rating point (14 rating = 1%, 13.876 per %), spell_haste=6.084 ± 1.303

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 101.1 healing_power points (45.65 DPS) | yes | Mooncloth Circlet (14140, -15.58 DPS, sim-verified) [crafted]; Whitesoul Helm (12633, -18.17 DPS) [crafted]; Devout Crown (16693, -20.92 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.9 healing_power points (12.61 DPS) | yes | Archlight Talisman (15856, -1.79 DPS, sim-verified) [quest]; Horizon Choker (13085, -5.02 DPS) [world_drop]; Frostwolf Advisor's Pendant (19096, -5.18 DPS) [rep] |
| shoulder | Darkspear Mantle (272107) | Creeg Bothunk [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Dawnbringer Shoulders (12625, -1.14 DPS, sim-verified) [crafted]; Mooncloth Shoulders (14139, -2.63 DPS) [crafted]; Mantle of the Timbermaw (19050, -3.21 DPS) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -3.39 DPS) [dungeon]; Deep Woodlands Cloak (19121, -4.94 DPS) [quest]; Darkspear Raider's Cloak (272063, -5.24 DPS) [vendor] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 78.0 healing_power points (35.24 DPS) | yes | Stormcloth Vest (10020, -7.70 DPS, sim-verified) [crafted]; Mooncloth Vest (14138, -8.27 DPS) [crafted]; Red Dragonscale Breastplate (15047, -9.50 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 35.1 healing_power points (15.83 DPS) | yes | Nethergeld Cuffs (254061, -2.01 DPS) [crafted]; Mender's Leather Bracers (252543, -2.13 DPS) [crafted]; Prefect's Wristguards (250584, -2.30 DPS, sim-verified) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, -1.92 DPS) [crafted]; Prefect's Gloves (250554, -2.19 DPS) [crafted]; Mender's Leather Gauntlets (252551, -2.50 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 51.5 healing_power points (23.24 DPS) | yes | Mender's Leather Waistguard (252477, -4.32 DPS) [crafted]; Mender's Mail Belt (252591, -4.32 DPS) [crafted]; Elderwild Waistcord (279252, -4.57 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironfeather Leggings (252486, -9.91 DPS, sim-verified) [crafted]; Elderwild Pants (279254, -19.18 DPS) [crafted]; Devout Skirt (16694, -20.03 DPS) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 47.4 healing_power points (21.40 DPS) | yes | Broodwatcher's Treaders (275621, -1.99 DPS) [crafted]; Swarmtender's Footpads (275609, -3.31 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -3.38 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 29.3 healing_power points (13.25 DPS) | yes | Advisor's Ring (19518, -3.96 DPS) [rep]; Band of the Hierophant (13096, -5.13 DPS) [world_drop] |
| finger2 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 21.9 healing_power points (9.89 DPS) | yes | Advisor's Ring (19518, -0.61 DPS) [rep]; Band of the Hierophant (13096, -1.78 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+7.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -3.67 DPS, sim-verified) [quest]; Ankh of Life (1713, -4.16 DPS) [world_drop]; Darkspear Voodoo Seal (272061, -4.43 DPS) [vendor] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Sageblade (22383, -4.26 DPS, sim-verified) [crafted]; Hand of Righteousness (7721, -4.67 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Evergreen Shield (279262, -2.73 DPS, sim-verified) [crafted] |
| ranged | Libram of Infusion (279248) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, -1.06 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; shoulder: Darkspear Mantle; back: Battle Healer's Cloak; chest: Breastplate of Salvation; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Maiden's Circle; finger2: Deep Rooted Ring; trinket1: Dormant Heart of the Mountain; main_hand: Hand of Edward the Odd; off_hand: Lei of the Lifegiver; ranged: Libram of Infusion

No-known-source sample (15 of 1670, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

