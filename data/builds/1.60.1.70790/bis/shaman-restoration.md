# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5320000010000000)

Set DPS (verified): 72.8. Weights run: 7.1s. Verify run: 33.5s. 260 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.111, intellect=1.988 ± 0.050, spirit=0.358 ± 0.015, mp5=4.171 ± 0.103, crit=0.105 ± 0.008 per rating point (14 rating = 1%, 1.464 per %), spell_haste=not significant (0.172 ± 0.102)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.34 DPS) | yes | Wisdom's Leather Hood (252507, -0.06 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -0.16 DPS) [crafted]; Shadow Goggles (4373, -0.44 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.7 healing_power points (0.05 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 22.9 healing_power points (1.70 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.29 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.77 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.41 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanguine Cape (14376, -0.13 DPS) [world_drop]; Heavy Woolen Cloak (4311, -0.32 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.43 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.0 healing_power points (1.93 DPS) | yes | Filigreed Pristine Gown (253901, -0.16 DPS, sim-verified) [crafted]; Stormrider's Leather Armor (252492, -0.52 DPS) [crafted]; Gray Woolen Robe (2585, -0.90 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.5 healing_power points (1.67 DPS) | yes | Owl Bracers (4796, -0.49 DPS, sim-verified) [vendor]; Bright Bracers (3647, -1.08 DPS) [world_drop]; Inscribed Leather Bracers (3205, -1.23 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Gloves (252499, -0.04 DPS, sim-verified) [crafted]; Stormrider's Leather Gloves (252498, -0.30 DPS) [crafted]; Magefist Gloves (12977, -0.42 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 19.0 healing_power points (1.41 DPS) | yes | Wisdom's Leather Belt (252433, -0.22 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.23 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.44 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.4 healing_power points (2.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.16 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.74 DPS) [crafted]; Darkweave Breeches (12987, -1.29 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.9 healing_power points (1.56 DPS) | yes | Black Whelp Slippers (252424, -0.20 DPS, sim-verified) [crafted]; Stormrider's Leather Boots (252443, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.44 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 13.3 healing_power points (0.99 DPS) | yes | Volcanic Rock Ring (12053, -0.55 DPS) [world_drop]; The 1 Ring (8350, -0.82 DPS) [world]; Lavishly Jeweled Ring (1156, -0.84 DPS, sim-verified) [dungeon] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Volcanic Rock Ring (12053, -0.22 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.23 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 60.1 healing_power points (4.47 DPS) | yes | Taskmaster Axe (5194, -0.55 DPS, sim-verified) [dungeon]; Living Root (6631, -0.99 DPS) [dungeon]; Evocator's Blade (2567, -1.96 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (dwarf, 0000000000000000-000000000000000000-5530003310010000)

Set DPS (verified): 116.2. Weights run: 8.4s. Verify run: 62.0s. 448 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.133, intellect=1.506 ± 0.009, spirit=0.823 ± 0.009, mp5=2.652 ± 0.014, crit=0.178 ± 0.009 per rating point (14 rating = 1%, 2.493 per %), spell_haste=-0.630 ± 0.112

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.9 healing_power points (8.07 DPS) | yes | Enduring Cap (3020, -2.30 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -3.35 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -3.39 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.3 healing_power points (1.98 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.06 DPS) [world]; Crystal Starfire Medallion (5003, -0.17 DPS) [world_drop]; Seal of the Expedition (271920, -0.38 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Mantle (6685, -0.07 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.85 DPS) [world_drop]; Mantle of Honor (3560, -1.33 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271923, -0.38 DPS) [rep]; Repairman's Cape (9605, -0.46 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.58 DPS) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 36.6 healing_power points (7.78 DPS) | yes | Pristine Gown (253961, -0.58 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.62 DPS) [dungeon]; Wisdom's Leather Tunic (252511, -0.98 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.0 healing_power points (3.61 DPS) | yes | Mindthrust Bracers (1974, -0.45 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.04 DPS) [world_drop]; Nightsky Wristbands (6407, -1.16 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.5 healing_power points (4.15 DPS) | yes | Gloves of Old (9395, -0.60 DPS, sim-verified) [world_drop]; Hotshot Pilot's Gloves (9491, -0.71 DPS) [dungeon]; Wisdom's Leather Gloves (252499, -0.74 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 33.0 healing_power points (7.01 DPS) | yes | Skycaller's Leather Belt (252522, -2.35 DPS, sim-verified) [crafted]; Highlander's Lizardhide Girdle (20105, -3.17 DPS) [rep]; Highlander's Mail Girdle (20120, -3.17 DPS) [vendor] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 37.3 healing_power points (7.93 DPS) | yes | Pristine Leggings (253987, -0.10 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -1.06 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.48 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.8 healing_power points (7.20 DPS) | yes | Acidic Walkers (9454, -3.08 DPS, sim-verified) [dungeon]; Wisdom's Leather Boots (252444, -3.25 DPS) [crafted]; Black Whelp Slippers (252424, -4.00 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.8 healing_power points (3.14 DPS) | yes | Snake Hoop (6750, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.52 DPS) [rep]; Runemender's Seal (273031, -3.32 DPS, sim-verified) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.09 DPS) [quest]; Runemender's Seal (273031, -0.11 DPS, sim-verified) [world]; Lorekeeper's Ring (19525, -0.35 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.30 DPS) [dungeon]; Staff of Westfall (2042, -1.60 DPS) [quest]; Staff of the Friar (3415, -4.63 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 17.9 healing_power points (3.81 DPS) | yes | Orb of Mystic Insight (249394, -0.42 DPS, sim-verified) [crafted]; Heart of Agamaggan (6694, -0.62 DPS) [dungeon]; Eye of Paleth (2943, -1.05 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 0000000000000000-000000000000000000-5530003315510000)

Set DPS (verified): 164.0. Weights run: 10.4s. Verify run: 118.2s. 696 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.122, intellect=1.124 ± 0.019, spirit=0.604 ± 0.018, mp5=2.472 ± 0.018, crit=0.200 ± 0.014 per rating point (14 rating = 1%, 2.806 per %), spell_haste=not significant (0.429 ± 0.289)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 46.6 healing_power points (11.31 DPS) | yes | Whitemane's Chapeau (7720, -2.16 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.41 DPS) [world_drop]; Earthen Silk Hood (254015, -4.50 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.0 healing_power points (7.05 DPS) | yes | Triune Amulet (7722, -3.58 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -4.11 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -4.31 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 36.6 healing_power points (8.90 DPS) | yes | Earthen Silk Shoulders (254033, -3.39 DPS, sim-verified) [crafted]; Sheepshear Mantle (13115, -3.65 DPS) [world_drop]; Green Silken Shoulders (7057, -3.95 DPS) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Expeditionary's Cape (271922, -1.52 DPS) [rep]; Repairman's Cape (9605, -2.10 DPS) [quest]; Glowing Thresher Cape (6901, -2.23 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.8 healing_power points (15.27 DPS) | yes | Deathchill Armor (10764, -3.35 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -6.38 DPS) [crafted]; Civinad Robes (9623, -7.19 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.4 healing_power points (3.99 DPS) | yes | Mindthrust Bracers (1974, -0.22 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.82 DPS) [crafted]; Guardian Leather Bracers (4260, -0.89 DPS) [crafted] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, -0.50 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.55 DPS) [crafted]; Dreamweave Gloves (10019, -1.07 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 34.6 healing_power points (8.41 DPS) | yes | Prefect's Belt (250559, -1.19 DPS, sim-verified) [crafted]; Mender's Leather Belt (252523, -1.28 DPS) [crafted]; Deathmage Sash (10771, -2.61 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 39.2 healing_power points (9.53 DPS) | yes | Stoneweaver Leggings (9407, -0.46 DPS, sim-verified) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -1.27 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.48 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 36.1 healing_power points (8.77 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.41 DPS) [crafted]; Skycaller's Leather Shoes (252532, -2.97 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 16.4 healing_power points (3.99 DPS) | yes | Ring of Power Regulation (273030, -0.73 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Runemender's Seal (273031, -0.77 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.6 healing_power points (3.80 DPS) | yes | Ring of Power Regulation (273030, -0.18 DPS, sim-verified) [dungeon]; Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon]; Runemender's Seal (273031, -0.58 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -5.31 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -7.49 DPS) [dungeon]; Staff of the Friar (3415, -8.80 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Beacon of Hope (9393, -0.40 DPS, sim-verified) [dungeon]; Battered Viking Shield (9403, -1.02 DPS) [dungeon]; Orb of Souls (249395, -1.61 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 0000000000000000-000000000000000000-5530003315513151)

Set DPS (verified): 259.5. Weights run: 18.0s. Verify run: 213.4s. 893 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.139, intellect=1.635 ± 0.025, spirit=0.799 ± 0.021, mp5=2.743 ± 0.025, crit=0.330 ± 0.019 per rating point (14 rating = 1%, 4.624 per %), spell_haste=not significant (0.483 ± 0.362)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.0 healing_power points (14.01 DPS) | yes | Soothsayer's Headdress (17740, -0.71 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -1.30 DPS) [crafted]; Cassandra's Grace (13102, -1.42 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.1 healing_power points (8.13 DPS) | yes | Horizon Choker (13085, -1.54 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -2.93 DPS) [dungeon]; Mindburst Medallion (11196, -3.26 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 54.1 healing_power points (13.30 DPS) | yes | Mender's Mail Shoulders (252569, -2.40 DPS) [crafted]; Lead Surveyor's Mantle (11842, -2.41 DPS) [dungeon]; Mender's Leather Shoulders (252538, -2.56 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -0.29 DPS, sim-verified) [rep]; Caretaker's Cape (19531, -0.44 DPS) [rep]; Runecloth Cloak (13860, -1.60 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.4 healing_power points (15.83 DPS) | yes | Feathered Breastplate (8349, -2.33 DPS, sim-verified) [crafted]; Robes of Insight (940, -2.84 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -3.10 DPS) [world] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 35.4 healing_power points (8.71 DPS) | yes | Mender's Leather Bracers (252543, -0.05 DPS) [crafted]; Mender's Mail Bracers (252573, -0.05 DPS) [crafted]; Aristocratic Cuffs (12546, -1.50 DPS) [dungeon] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.29 DPS) [crafted]; Ogreseer Fists (11665, -4.22 DPS) [dungeon] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 49.6 healing_power points (12.20 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Waistcord (254081, -1.45 DPS) [crafted]; Dawnspire Cord (12466, -1.61 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 50.8 healing_power points (12.49 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.25 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.81 DPS) [dungeon]; Warchief Kilt (7760, -1.43 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 46.0 healing_power points (11.30 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.07 DPS) [crafted]; Mender's Leather Shoes (252533, -1.22 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-verified (259.5 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, -0.82 DPS) [dungeon]; Eye of Adaegus (5266, -0.88 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 23.6 healing_power points (5.81 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, -0.99 DPS) [dungeon]; Eye of Adaegus (5266, -1.04 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, +0.00 DPS) [rep]; Uther's Strength (11302, -0.39 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -15.43 DPS, sim-verified) [quest]; Charstone Dirk (17710, -16.01 DPS) [dungeon]; Dreamstaff (249454, -20.75 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Ironfeather Shoulders; back: Spritecaster Cape; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Mender's Leather Boots; finger1: Lorekeeper's Ring; finger2: Cyclopean Band; trinket2: Dormant Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 893, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 356.5. Weights run: 19.1s. Verify run: 531.5s. 2086 eligible items had no known source.

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

**New at 60:** head: Living Crown; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Mooncloth Vest; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Red Dragonscale Leggings; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket2: Serenity Field; ranged: Totem of Urgency

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 878.4. Weights run: 11.9s. Verify run: 486.8s. 2086 eligible items had no known source.

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

**New at 60:** head: Living Crown; shoulder: Theramore Shoulderguards; back: Drape of Recovery; chest: Robes of the Exalted; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Fordring's Seal; trinket1: Serenity Field; ranged: Totem of Urgency

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-000000000000000000-5320000010000000)

Set DPS (verified): 72.2. Weights run: 7.1s. Verify run: 32.9s. 249 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=0.918 ± 0.003, spirit=0.252 ± 0.003, mp5=1.638 ± 0.012, crit=0.049 ± 0.002 per rating point (14 rating = 1%, 0.689 per %), spell_haste=0.804 ± 0.068

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.78 DPS) | yes | Wisdom's Leather Hood (252507, -0.14 DPS, sim-verified) [crafted]; Totemic Leather Hood (252448, -1.39 DPS) [crafted]; Trapper's Leather Hood (252505, -1.65 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.8 healing_power points (0.12 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.3 healing_power points (2.05 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.21 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.29 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.43 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-verified (72.2 DPS) | yes | Heavy Woolen Cloak (4311, -0.66 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.73 DPS) [crafted]; Sanguine Cape (14376, -0.90 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 20.3 healing_power points (3.14 DPS) | yes | Filigreed Pristine Gown (253901, -0.29 DPS, sim-verified) [crafted]; Stormrider's Leather Armor (252492, -1.08 DPS) [crafted]; Corsair's Overshirt (5202, -1.13 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.5 healing_power points (1.47 DPS) | yes | Bright Bracers (3647, -0.90 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.01 DPS) [quest]; Owl Bracers (4796, -1.53 DPS, sim-verified) [vendor] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.8 healing_power points (2.13 DPS) | yes | Wisdom's Leather Gloves (252499, -0.01 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.79 DPS) [crafted]; Serpent Gloves (5970, -1.04 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.7 healing_power points (2.27 DPS) | yes | Wisdom's Leather Belt (252433, -0.44 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.80 DPS) [crafted]; Stormrider's Leather Belt (252432, -1.08 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 26.5 healing_power points (4.10 DPS) | yes | Filigreed Pristine Leggings (253937, -0.29 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -1.55 DPS) [crafted]; Totemic Leather Pants (252446, -2.45 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 15.6 healing_power points (2.41 DPS) | yes | Black Whelp Slippers (252424, -0.46 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.59 DPS) [crafted]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 8.3 healing_power points (1.28 DPS) | yes | Loop of Sacrifice (281673, -0.57 DPS) [quest]; Volcanic Rock Ring (12053, -0.85 DPS) [world_drop]; Furlfeather Ring (286745, -0.97 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.5 healing_power points (0.85 DPS) | yes | Loop of Sacrifice (281673, -0.15 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop]; Furlfeather Ring (286745, -0.54 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Taskmaster Axe (5194) | The Deadmines: Sneed [dungeon] | 47.0 healing_power points (7.27 DPS) | yes | Living Root (6631, -0.19 DPS, sim-verified) [dungeon]; Kris of Orgrimmar (15443, -2.17 DPS) [quest]; Crescent Staff (6505, -2.25 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Taskmaster Axe

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads

### Band 30 (orc, 0000000000000000-000000000000000000-5530003310010000)

Set DPS (verified): 116.8. Weights run: 8.4s. Verify run: 60.7s. 442 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=0.785 ± 0.004, spirit=0.507 ± 0.004, mp5=1.546 ± 0.004, crit=0.080 ± 0.003 per rating point (14 rating = 1%, 1.115 per %), spell_haste=-1.244 ± 0.075

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (6.37 DPS) | yes | Wisdom's Leather Helm (252515, -2.66 DPS) [crafted]; Pristine Circlet (253949, -3.19 DPS) [crafted]; Filigreed Pristine Circlet (253975, -3.56 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.2 healing_power points (0.91 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Crystal Starfire Medallion (5003, -0.09 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.22 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.1 healing_power points (2.84 DPS) | yes | Death Speaker Mantle (6685, -0.25 DPS) [dungeon]; Magician's Mantle (12998, -0.71 DPS) [world_drop]; Invoker's Mantle (215365, -0.91 DPS) [crafted] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-verified (116.8 DPS) | yes | Glowing Thresher Cape (6901, -0.73 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -1.28 DPS) [vendor]; Windsong Drape (15468, -1.41 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 28.6 healing_power points (5.06 DPS) | yes | Wisdom's Leather Tunic (252511, -0.51 DPS) [crafted]; Pristine Gown (253961, -0.56 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.09 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.6 healing_power points (2.41 DPS) | yes | Mindthrust Bracers (1974, -0.47 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.30 DPS) [world_drop]; Nightsky Wristbands (6407, -1.31 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.4 healing_power points (3.07 DPS) | yes | Pristine Gloves (253913, -0.71 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.75 DPS) [crafted]; Naga Battle Gloves (888, -0.98 DPS, sim-verified) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 26.8 healing_power points (4.73 DPS) | yes | Pristine Sash (253925, -2.23 DPS) [crafted]; Belt of Arugal (6392, -2.27 DPS) [dungeon]; Skycaller's Leather Belt (252522, -2.38 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 30.0 healing_power points (5.31 DPS) | yes | Pristine Leggings (253987, -0.07 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.58 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.94 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 27.5 healing_power points (4.87 DPS) | yes | Acidic Walkers (9454, -2.51 DPS) [dungeon]; Black Whelp Slippers (252424, -2.72 DPS) [crafted]; Wisdom's Leather Boots (252444, -3.38 DPS, sim-verified) [crafted] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 13.5 healing_power points (2.39 DPS) | yes | Runemender's Seal (273031, -0.17 DPS) [world]; Ring of Power Regulation (273030, -0.51 DPS) [dungeon]; Advisor's Ring (19521, -0.61 DPS) [rep] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.5 healing_power points (2.22 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Ring of Power Regulation (273030, -0.34 DPS) [dungeon]; Advisor's Ring (19521, -0.43 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (116.8 DPS) | yes | Odo's Ley Staff (6318, -1.75 DPS) [dungeon]; Taskmaster Axe (5194, -3.35 DPS) [dungeon]; Staff of the Friar (3415, -4.63 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.0 healing_power points (2.84 DPS) | yes | Heart of Agamaggan (6694, -0.55 DPS, sim-verified) [dungeon]; Orb of Mystic Insight (249394, -0.77 DPS) [crafted]; Orb of Mistmantle (13031, -0.89 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 0000000000000000-000000000000000000-5530003315510000)

Set DPS (verified): 165.0. Weights run: 10.4s. Verify run: 116.2s. 679 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.122, intellect=1.124 ± 0.019, spirit=0.604 ± 0.018, mp5=2.472 ± 0.018, crit=0.200 ± 0.014 per rating point (14 rating = 1%, 2.806 per %), spell_haste=not significant (0.429 ± 0.289)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 46.6 healing_power points (11.31 DPS) | yes | Whitemane's Chapeau (7720, -2.17 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.41 DPS) [world_drop]; Earthen Silk Hood (254015, -4.50 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.0 healing_power points (7.05 DPS) | yes | Triune Amulet (7722, -3.75 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -4.11 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -4.31 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 36.6 healing_power points (8.90 DPS) | yes | Earthen Silk Shoulders (254033, -3.36 DPS, sim-verified) [crafted]; Sheepshear Mantle (13115, -3.65 DPS) [world_drop]; Green Silken Shoulders (7057, -3.95 DPS) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272077, -1.52 DPS) [vendor]; Glowing Thresher Cape (6901, -2.23 DPS) [dungeon]; Long Silken Cloak (4326, -2.28 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.8 healing_power points (15.27 DPS) | yes | Deathchill Armor (10764, -3.54 DPS, sim-verified) [dungeon]; Dreamweave Vest (10021, -6.38 DPS) [crafted]; Civinad Robes (9623, -7.19 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.4 healing_power points (3.99 DPS) | yes | Mindthrust Bracers (1974, -0.20 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.82 DPS) [crafted]; Radiant Silver Bracers (4545, -0.83 DPS) [quest] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prefect's Gauntlet (250569, -0.55 DPS) [crafted]; Gilded Handwraps (254021, -0.56 DPS, sim-verified) [crafted]; Dreamweave Gloves (10019, -1.07 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 34.6 healing_power points (8.41 DPS) | yes | Mender's Leather Belt (252523, -1.28 DPS) [crafted]; Prefect's Belt (250559, -1.60 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -2.61 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 39.2 healing_power points (9.53 DPS) | yes | Stoneweaver Leggings (9407, -0.72 DPS, sim-verified) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -1.27 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.48 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 36.1 healing_power points (8.77 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -1.41 DPS) [crafted]; Skycaller's Leather Shoes (252532, -2.97 DPS) [crafted] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 16.4 healing_power points (3.99 DPS) | yes | Solid Gold Band (279845, -0.58 DPS) [quest]; Ring of Power Regulation (273030, -0.73 DPS) [dungeon]; Philanthropist's Ring (281635, -3.83 DPS, sim-verified) [quest] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Solid Gold Band (279845, -0.20 DPS) [quest]; Ring of Power Regulation (273030, -0.34 DPS) [dungeon]; Philanthropist's Ring (281635, -0.44 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -5.62 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -7.49 DPS) [dungeon]; Staff of the Friar (3415, -8.80 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forcestone Buckler (17508, -0.38 DPS) [quest]; Battered Viking Shield (9403, -1.31 DPS) [dungeon]; Orb of Souls (249395, -1.89 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Advisor's Ring; finger2: Darkspear Signet; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 0000000000000000-000000000000000000-5530003315513151)

Set DPS (verified): 260.7. Weights run: 18.0s. Verify run: 224.6s. 862 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.059, intellect=1.465 ± 0.017, spirit=0.984 ± 0.014, mp5=2.709 ± 0.017, crit=0.193 ± 0.008 per rating point (14 rating = 1%, 2.709 per %), spell_haste=not significant (0.154 ± 0.184)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 55.8 healing_power points (9.08 DPS) | yes | Soothsayer's Headdress (17740, -0.47 DPS) [dungeon]; Gemburst Circlet (10751, -0.68 DPS) [quest]; Cassandra's Grace (13102, -1.85 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.7 healing_power points (5.17 DPS) | yes | Horizon Choker (13085, -1.62 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -1.82 DPS) [dungeon]; Mindburst Medallion (11196, -1.96 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 52.2 healing_power points (8.50 DPS) | yes | Mender's Mail Shoulders (252569, -1.32 DPS) [crafted]; Living Shoulders (15061, -1.36 DPS) [crafted]; Mender's Leather Shoulders (252538, -2.61 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19527, -0.47 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272076, -0.53 DPS) [vendor]; Deep Woodlands Cloak (19121, -0.57 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.9 healing_power points (10.73 DPS) | yes | Embrace of the Wind Serpent (12462, -1.86 DPS) [world]; Robes of Insight (940, -2.36 DPS) [world_drop]; Feathered Breastplate (8349, -2.43 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 35.2 healing_power points (5.73 DPS) | yes | Mender's Leather Bracers (252543, -0.00 DPS) [crafted]; Mender's Mail Bracers (252573, -0.00 DPS) [crafted]; Aristocratic Cuffs (12546, -1.19 DPS) [dungeon] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.25 DPS) [crafted]; Greenleaf Handwraps (19116, -2.43 DPS) [quest] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 47.6 healing_power points (7.75 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Waistcord (254081, -0.88 DPS) [crafted]; Dawnspire Cord (12466, -1.26 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 53.0 healing_power points (8.63 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.47 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.59 DPS) [dungeon]; Warchief Kilt (7760, -1.04 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 44.2 healing_power points (7.20 DPS) | yes | Mender's Leather Boots (252472, -0.01 DPS) [crafted]; Mender's Mail Sabatons (252579, -0.01 DPS) [crafted]; Mender's Leather Shoes (252533, -0.50 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Eye of Adaegus (5266, -0.36 DPS) [world_drop]; Darkspear Signet (272069, -0.63 DPS) [vendor] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 23.2 healing_power points (3.78 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Eye of Adaegus (5266, -0.42 DPS) [world_drop]; Darkspear Signet (272069, -0.69 DPS) [vendor] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.32 DPS) [quest]; Uther's Strength (11302, -0.50 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -10.39 DPS) [dungeon]; Dreamstaff (249454, -12.65 DPS) [crafted]; Resurgence Rod (17743, -15.78 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Ironfeather Shoulders; back: Spritecaster Cape; wrist: Nethergeld Cuffs; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Advisor's Ring; finger2: Cyclopean Band; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 862, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 359.0. Weights run: 19.1s. Verify run: 523.5s. 2043 eligible items had no known source.

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

Set DPS (verified): 873.5. Weights run: 11.9s. Verify run: 487.8s. 2043 eligible items had no known source.

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

