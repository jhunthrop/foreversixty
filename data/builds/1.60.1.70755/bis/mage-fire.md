# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 37.8. Weights run: 2.8s. Verify run: 2.3s. 178 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.114 ± 0.016, crit=0.161 ± 0.007 per rating point (14 rating = 1%, 2.255 per %), hit=0.334 ± 0.016 per rating point (10 rating = 1%, 3.340 per %), spell_haste=not significant (0.637 ± 0.201), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.123 ± 0.060), mp5=not significant (0.649 ± 0.173)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.36 DPS) | yes | Pristine Circlet (253949, -0.62 DPS) [crafted]; Shadow Goggles (4373, -1.07 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.30 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.2 spell_power points (0.03 DPS) | yes | Tarnished Locket (279870, -0.34 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.0 spell_power points (1.86 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.93 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.36 DPS) [crafted]; Slime-encrusted Pads (6461, -1.62 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-verified (37.8 DPS) | yes | Sanguine Cape (14376, -0.11 DPS) [world_drop]; Caretaker's Cape (20428, -0.26 DPS) [rep]; Heavy Woolen Cloak (4311, -0.43 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.9 spell_power points (1.85 DPS) | yes | Filigreed Pristine Gown (253901, -0.57 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.66 DPS) [crafted]; Mystic's Wrap (14369, -0.88 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.5 spell_power points (0.93 DPS) | yes | Repurposed Hair Band (281256, -0.62 DPS) [quest]; Mystic's Bracelets (14366, -0.65 DPS) [world_drop]; Bright Bracers (3647, -0.83 DPS, sim-verified) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 10.3 spell_power points (1.28 DPS) | yes | Pristine Gloves (253913, -0.37 DPS) [crafted]; Serpent Gloves (5970, -0.41 DPS) [dungeon]; Phoenix Gloves (4331, -0.45 DPS, sim-verified) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 11.5 spell_power points (1.42 DPS) | yes | Pristine Sash (253925, -0.37 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.51 DPS) [crafted]; Keller's Girdle (2911, -0.85 DPS, sim-verified) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 18.2 spell_power points (2.25 DPS) | yes | Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted]; Phoenix Pants (4317, -0.90 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.19 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.5 spell_power points (1.42 DPS) | yes | Pristine Boots (253889, -0.63 DPS) [crafted]; Flame Boots (253893, -0.71 DPS, sim-verified) [crafted]; Kimbra Boots (6191, -0.82 DPS) [quest] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.2 spell_power points (0.89 DPS) | yes | Lorekeeper's Ring (20431, -0.11 DPS) [rep]; Volcanic Rock Ring (12053, -0.48 DPS) [world_drop]; The 1 Ring (8350, -0.74 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.7 spell_power points (0.83 DPS) | yes | Lorekeeper's Ring (20431, -0.05 DPS) [rep]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; The 1 Ring (8350, -0.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.8 spell_power points (3.81 DPS) | yes | Emberstone Staff (5201, -0.85 DPS) [dungeon]; Staff of Westfall (2042, -1.05 DPS) [quest]; Staff of the Blessed Seer (2271, -1.49 DPS, sim-verified) [dungeon] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.74 DPS) | yes | Tear of Grief (5611, -0.10 DPS) [quest]; Seer's Fine Stein (7608, -0.31 DPS) [world_drop]; Mathystra Channeling Orb (281291, -0.51 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-verified (37.8 DPS) | yes | Skycaller (12984, -0.41 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -1.06 DPS) [quest]; Deepblaze (279896, -1.06 DPS) [quest] |

**New at 20:** head: Flame Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 178, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff; 20440 Protector's Sword

### Band 30 (gnome, 000000000000000000-51521001230100000-000000000000000000)

Set DPS (verified): 76.3. Weights run: 3.2s. Verify run: 5.2s. 303 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.535 ± 0.011, crit=0.139 ± 0.008 per rating point (14 rating = 1%, 1.949 per %), hit=0.267 ± 0.014 per rating point (10 rating = 1%, 2.673 per %), spell_haste=0.483 ± 0.052, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.043 ± 0.043), mp5=not significant (0.202 ± 0.076)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (2.38 DPS) | yes | Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Holy Shroud (2721, -0.47 DPS) [world_drop]; Flame Circlet (253953, -0.51 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.2 spell_power points (0.55 DPS) | yes | Seal of the Expedition (271920, -0.09 DPS) [rep]; Kaleidoscope Chain (13084, -0.15 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.8 spell_power points (2.35 DPS) | yes | Death Speaker Mantle (6685, -0.23 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.68 DPS) [world_drop]; Invoker's Mantle (215365, -0.71 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Expeditionary's Cape (271923, -0.10 DPS) [rep]; Heavy Woolen Cloak (4311, -0.14 DPS) [crafted] |
| chest | Flame Gown (253965) | Tailoring [crafted] | 16.9 spell_power points (2.88 DPS) | yes | Green Silk Armor (7065, -0.16 DPS) [crafted]; Beguiler Robes (7728, -0.20 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.43 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (2.21 DPS) | yes | Spidertank Oilrag (9448, -0.58 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.49 DPS) [world_drop]; Nightsky Wristbands (6407, -1.65 DPS) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (1.53 DPS) | yes | Flame Gloves (253917, -0.07 DPS) [crafted]; Serpent Gloves (5970, -0.34 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.38 DPS) [quest] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Crimson Silk Belt (7055, -0.18 DPS) [crafted]; Invoker's Cord (215366, -0.20 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.27 DPS, sim-verified) [rep] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.0 spell_power points (3.06 DPS) | yes | Filigreed Flame Leggings (253941, -0.73 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.84 DPS) [world]; Gaze Dreamer Pants (6903, -0.95 DPS) [dungeon] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 16.9 spell_power points (2.88 DPS) | yes | Acidic Walkers (9454, -1.27 DPS) [dungeon]; Gilded Slippers (254001, -1.29 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.32 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.7 spell_power points (1.48 DPS) | yes | Ring of Power Regulation (273030, -0.35 DPS) [dungeon]; Minor Channeling Ring (1449, -0.44 DPS) [quest]; Electrocutioner Lagnut (9447, -0.92 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.4 spell_power points (1.26 DPS) | yes | Ring of Power Regulation (273030, -0.14 DPS) [dungeon]; Minor Channeling Ring (1449, -0.23 DPS) [quest]; Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.71 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.71 DPS) [world]; Glimmering Staff (249392, -2.89 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (2.21 DPS) | yes | Orb of Mystic Insight (249394, -0.48 DPS) [crafted]; Antipodean Rod (2879, -1.19 DPS) [world_drop]; Orb of Souls (249395, -1.49 DPS) [crafted] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, +0.00 DPS) [world_drop]; Wand of Mana Concentration (273043, +0.00 DPS) [world]; Greater Mystic Wand (217287, -0.13 DPS) [crafted] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Flame Gown; wrist: Phoenix Bindings; hands: Phoenix Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 303, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 000000000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 163.9. Weights run: 3.5s. Verify run: 5.1s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.749 ± 0.021, crit=0.373 ± 0.022 per rating point (14 rating = 1%, 5.224 per %), hit=0.622 ± 0.037 per rating point (10 rating = 1%, 6.219 per %), spell_haste=not significant (0.660 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=1.034 ± 0.184, mp5=1.226 ± 0.231

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 30.0 spell_power points (10.74 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -2.41 DPS) [dungeon]; Whitemane's Chapeau (7720, -3.11 DPS, sim-verified) [dungeon]; Spellpower Goggles Xtreme (10502, -3.23 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -0.18 DPS) [dungeon]; Amberglow Talisman (10824, -0.95 DPS) [quest]; Triune Amulet (7722, -1.09 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Mantle of Doan (7712, -1.10 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.30 DPS) [rep]; Glowing Thresher Cape (6901, -0.69 DPS) [dungeon]; Caretaker's Cape (19532, -0.73 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 39.2 spell_power points (14.03 DPS) | yes | Robe of the Magi (1716, -2.70 DPS) [world_drop]; Robe of Power (7054, -2.85 DPS) [crafted]; Cindercloth Robe (10042, -4.37 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.65 DPS) | yes | Earthen Silk Cuffs (254019, -1.37 DPS) [crafted]; Arcane Runed Bracers (4744, -1.43 DPS) [quest]; Spidertank Oilrag (9448, -1.45 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 28.2 spell_power points (10.10 DPS) | yes | Fiery Handwraps (254025, -1.91 DPS, sim-verified) [crafted]; Crimson Silk Gloves (7064, -3.48 DPS) [crafted]; Red Mageweave Gloves (10018, -3.49 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 26.2 spell_power points (9.37 DPS) | yes | Gilded Cord (254037, -2.15 DPS) [crafted]; Deathmage Sash (10771, -2.85 DPS) [dungeon]; Scorching Sash (4117, -2.99 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.51 DPS) [crafted]; Black Mageweave Leggings (9999, -0.69 DPS) [crafted]; Stoneweaver Leggings (9407, -1.70 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 31.2 spell_power points (11.17 DPS) | yes | Fiery Slippers (254005, -3.32 DPS, sim-verified) [crafted]; Furen's Boots (13100, -3.38 DPS) [world_drop]; Thoughtcast Boots (10578, -4.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (4.92 DPS) | yes | Welken Ring (5011, -0.99 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.25 DPS) [dungeon]; Runemender's Seal (273031, -1.25 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 12.7 spell_power points (4.53 DPS) | yes | Electrocutioner Lagnut (9447, -0.87 DPS) [dungeon]; Runemender's Seal (273031, -0.87 DPS) [world]; Welken Ring (5011, -3.15 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -11.39 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -14.86 DPS) [dungeon]; Celestial Stave (9517, -15.26 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -1.22 DPS) [world_drop]; Goblin Igniter (5253, -2.00 DPS) [quest]; Twisted Nether Wand (249144, -2.70 DPS, sim-verified) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 252010000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 210.8. Weights run: 3.5s. Verify run: 9.5s. 491 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.475 ± 0.027, crit=0.216 ± 0.024 per rating point (14 rating = 1%, 3.017 per %), hit=0.385 ± 0.040 per rating point (10 rating = 1%, 3.849 per %), spell_haste=not significant (0.329 ± 0.238), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.546 ± 0.271), mp5=not significant (0.481 ± 0.349)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 32.3 spell_power points (8.43 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -1.39 DPS) [crafted]; Red Mageweave Headband (10033, -1.79 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.50 DPS) [vendor] |
| neck | Mindburst Medallion (11196) | Tremors of the Earth [quest] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Eye of Mordresh (10769, -0.20 DPS) [dungeon]; Horizon Choker (13085, -0.72 DPS) [world_drop]; Triune Amulet (7722, -0.75 DPS, sim-verified) [dungeon] |
| shoulder | Netherflame Shoulders (254053) | Tailoring [crafted] | 24.1 spell_power points (6.29 DPS) | yes | Rotgrip Mantle (17732, -0.66 DPS) [dungeon]; Flameseer Mantle (11310, -1.21 DPS) [dungeon]; Nethergeld Shoulders (254049, -2.09 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -0.82 DPS, sim-verified) [crafted]; Cloak of Fire (14134, -1.60 DPS) [crafted]; Runecloth Cloak (13860, -1.91 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Acumen Robes (17775, -0.37 DPS) [quest]; Robe of the Magi (1716, -0.61 DPS) [world_drop]; Cindercloth Vest (14042, -1.74 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.1 spell_power points (4.98 DPS) | yes | Phoenix Bindings (210781, -1.58 DPS) [crafted]; Bloodband Bracers (11469, -2.13 DPS) [quest]; Nethergeld Cuffs (254061, -2.18 DPS, sim-verified) [crafted] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 28.1 spell_power points (7.34 DPS) | yes | Cindercloth Gloves (14043, -1.33 DPS) [crafted]; Runecloth Gloves (13863, -1.81 DPS) [crafted]; Dreamweave Gloves (10019, -2.45 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 23.3 spell_power points (6.08 DPS) | yes | Fiery Cord (254041, -0.57 DPS) [crafted]; Ghostweave Cord (254073, -0.85 DPS) [crafted]; Satyrmane Sash (17755, -1.18 DPS) [dungeon] |
| legs | Cindercloth Leggings (12256) | Vinasia [vendor] | 27.2 spell_power points (7.10 DPS) | yes | Senior Designer's Pantaloons (11841, -0.53 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.52 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -1.69 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (7.26 DPS) | yes | Fiery Sandals (254111, -0.93 DPS) [crafted]; Cindercloth Boots (10044, -1.78 DPS) [crafted]; Fiery Slippers (254005, -2.43 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 13.9 spell_power points (3.64 DPS) | yes | Eye of Adaegus (5266, -1.18 DPS) [world_drop]; Snake Hoop (6750, -1.45 DPS) [quest]; Band of the Unicorn (7553, -5.00 DPS, sim-verified) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Band of the Unicorn (7553, -0.70 DPS, sim-verified) [world_drop]; Eye of Adaegus (5266, -0.77 DPS) [world_drop]; Snake Hoop (6750, -1.04 DPS) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -1.19 DPS, sim-verified) [crafted] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -2.61 DPS) [world]; Glowing Brightwood Staff (812, -3.14 DPS) [world_drop]; Kindling Stave (11750, -5.30 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -2.81 DPS) [dungeon]; Ragefire Wand (7513, -2.92 DPS) [quest]; Lesser Eternal Wand (249232, -3.65 DPS, sim-verified) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Mindburst Medallion; shoulder: Netherflame Shoulders; back: Spritecaster Cape; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Cindercloth Leggings; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Enriched Seal; ranged: Pyric Caduceus

No-known-source sample (15 of 491, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 252015000000200000-51521001230133031-003000000000000000)

Set DPS (verified): 352.4. Weights run: 3.8s. Verify run: 23.3s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.497 ± 0.036, crit=0.541 ± 0.032 per rating point (14 rating = 1%, 7.576 per %), hit=1.090 ± 0.058 per rating point (10 rating = 1%, 10.903 per %), spell_haste=not significant (0.976 ± 0.434), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.755 ± 0.463), mp5=not significant (0.367 ± 0.456)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (352.4 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -12.38 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.0 spell_power points (3.94 DPS) | yes | Triune Amulet (7722, -0.86 DPS) [dungeon]; Mindburst Medallion (11196, -0.98 DPS) [quest]; Lady Alizabeth's Pendant (13002, -1.31 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (352.4 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -12.58 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (352.4 DPS) | yes | Cindercloth Cloak (14044, -1.05 DPS) [crafted]; Caretaker's Cape (19530, -1.50 DPS) [rep]; Mageflame Cloak (13007, -2.83 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (352.4 DPS) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -2.96 DPS) [pvp]; Robes of Fiery Devastation (279266, -14.34 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.3 spell_power points (7.26 DPS) | yes | Netherflame Cuffs (254065, -2.56 DPS) [crafted]; Marshal's Silk Bracers (16438, -3.31 DPS) [pvp]; Runecloth Cuffs (254123, -3.55 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (352.4 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Fiery Gloves (254099, -14.19 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (352.4 DPS) | yes | Highlander's Cloth Girdle (20047, +0.00 DPS) [rep]; Fiery Waistcord (254085, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -16.48 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.1 spell_power points (12.10 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -2.56 DPS) [crafted]; Knight-Captain's Silk Legguards (227109, -2.63 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (352.4 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Venomspew Footpads (275606, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -13.24 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 26.0 spell_power points (6.04 DPS) | yes | Deep Rooted Ring (19109, -2.10 DPS) [world]; Lorekeeper's Ring (19522, -2.45 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 18.5 spell_power points (4.29 DPS) | yes | Deep Rooted Ring (19109, -0.35 DPS) [world]; Lorekeeper's Ring (19522, -0.70 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (352.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest]; Uther's Strength (11302, -0.71 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (352.4 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.01 DPS) [world_drop] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, -0.54 DPS) [rep]; Sageblade (22383, -1.17 DPS) [crafted]; Grand Marshal's Stave (234571, -1.24 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (352.4 DPS) | yes | Brilliant Wand (249385, -5.20 DPS) [crafted]; Pyric Caduceus (11748, -6.52 DPS) [dungeon]; Stormrager (16997, -6.65 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Sentinel's Silk Leggings; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Don Julio's Band; trinket1: Ankh of Life; trinket2: Thunderbrew's Boot Flask; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 253005000000200000-50520201230133031-003000000000000000)

Set DPS (verified): 737.9. Weights run: 4.3s. Verify run: 16.9s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.372 ± 0.030, crit=0.866 ± 0.043 per rating point (14 rating = 1%, 12.119 per %), hit=1.653 ± 0.077 per rating point (10 rating = 1%, 16.526 per %), spell_haste=4.201 ± 0.433, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.002, spirit=not significant (-0.022 ± 0.189), mp5=not significant (-0.007 ± 0.262)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.31 DPS) | yes | Red Mageweave Headband (10033, +0.00 DPS) [crafted]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.7 spell_power points (9.49 DPS) | yes | Glowing Eye of Mordresh (10769, -4.40 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -4.52 DPS) [quest]; Triune Amulet (7722, -4.66 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.0 spell_power points (20.51 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -1.66 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -2.04 DPS) [pvp]; Shroud of the Nathrezim (18720, -4.66 DPS, sim-verified) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | sim-verified (737.9 DPS) | yes | Spritecaster Cape (11623, -2.88 DPS) [dungeon]; Cindercloth Cloak (14044, -3.04 DPS) [crafted]; Cloak of Flames (3475, -3.96 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (737.9 DPS) | yes | Field Marshal's Silk Vestments (231603, -3.11 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -10.35 DPS) [pvp]; Robes of Fiery Devastation (279266, -15.97 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.0 spell_power points (15.09 DPS) | yes | Netherflame Cuffs (254065, -5.66 DPS) [crafted]; Runecloth Cuffs (254123, -6.20 DPS) [crafted]; Phoenix Bindings (210781, -7.23 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.0 spell_power points (22.32 DPS) | yes | Marshal's Silk Gloves (16440, -3.32 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -3.32 DPS) [vendor]; Fiery Gloves (254099, -9.44 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 28.3 spell_power points (17.12 DPS) | yes | Fiery Waistcord (254085, -3.63 DPS) [crafted]; Stormpike Cloth Girdle (19094, -4.01 DPS) [rep]; Satyrmane Sash (17755, -6.42 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 60.0 spell_power points (36.22 DPS) | yes | Marshal's Silk Leggings (231605, -6.28 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -12.39 DPS) [pvp]; Cindercloth Pants (14045, -17.82 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP reward · Alliance [vendor] | 27.9 spell_power points (16.84 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -1.90 DPS) [pvp]; Omnicast Boots (11822, -2.06 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 28.6 spell_power points (17.30 DPS) | yes | Blackstone Ring (17713, -7.32 DPS) [dungeon]; Deep Rooted Ring (19109, -7.81 DPS) [world]; Lorekeeper's Ring (19522, -8.85 DPS) [rep] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 20.6 spell_power points (12.45 DPS) | yes | Deep Rooted Ring (19109, -2.95 DPS) [world]; Lorekeeper's Ring (19522, -3.99 DPS) [rep]; Blackstone Ring (17713, -7.19 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (737.9 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (737.9 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (737.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Elemental Mage Staff (944, -1.87 DPS) [world_drop]; Ironbark Staff (20069, -21.41 DPS, sim-verified) [rep] |
| off_hand | Tome of Fiery Arcana (19311) | Stormpike Guard [rep] | 37.0 spell_power points (22.35 DPS) | yes | Grand Marshal's Tome of Power (234589, -4.33 DPS) [pvp]; Therazane's Touch (19315, -4.91 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -9.66 DPS) [quest] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (737.9 DPS) | yes | Brilliant Wand (249385, -0.83 DPS) [crafted]; Ragefire Wand (7513, -1.76 DPS) [quest]; Greater Eternal Wand (249237, -2.42 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Sageblade; off_hand: Tome of Fiery Arcana

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 36.3. Weights run: 2.8s. Verify run: 3.1s. 172 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.114 ± 0.016, crit=0.161 ± 0.007 per rating point (14 rating = 1%, 2.255 per %), hit=0.334 ± 0.016 per rating point (10 rating = 1%, 3.340 per %), spell_haste=not significant (0.637 ± 0.201), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.123 ± 0.060), mp5=not significant (0.649 ± 0.173)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.36 DPS) | yes | Pristine Circlet (253949, -0.62 DPS) [crafted]; Shadow Goggles (4373, -0.67 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.30 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.4 spell_power points (0.05 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.0 spell_power points (1.86 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.69 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.36 DPS) [crafted]; Slime-encrusted Pads (6461, -1.62 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.00 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.14 DPS, sim-verified) [crafted]; Battle Healer's Cloak (20427, -0.15 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.9 spell_power points (1.85 DPS) | yes | Filigreed Pristine Gown (253901, -0.53 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.66 DPS) [crafted]; Mystic's Wrap (14369, -0.88 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.5 spell_power points (0.93 DPS) | yes | Bright Bracers (3647, -0.26 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.52 DPS) [quest]; Owlbeard Bracers (16981, -0.53 DPS) [quest] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 10.3 spell_power points (1.28 DPS) | yes | Phoenix Gloves (4331, -0.17 DPS) [crafted]; Pristine Gloves (253913, -0.37 DPS) [crafted]; Serpent Gloves (5970, -0.41 DPS) [dungeon] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 11.5 spell_power points (1.42 DPS) | yes | Keller's Girdle (2911, -0.22 DPS, sim-verified) [world_drop]; Pristine Sash (253925, -0.37 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.51 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 18.2 spell_power points (2.25 DPS) | yes | Phoenix Pants (4317, -0.26 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted]; Darkweave Breeches (12987, -1.19 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.5 spell_power points (1.42 DPS) | yes | Flame Boots (253893, -0.26 DPS) [crafted]; Pristine Boots (253889, -0.63 DPS) [crafted]; Red Woolen Boots (4313, -0.85 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.7 spell_power points (0.83 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; The 1 Ring (8350, -0.67 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.3 spell_power points (0.78 DPS) | yes | Loop of Sacrifice (281673, -0.15 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.63 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.20 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.45 DPS) [quest]; Crescent Staff (6505, -0.93 DPS, sim-verified) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.74 DPS) | yes | Seer's Fine Stein (7608, -0.31 DPS) [world_drop]; Spellbinder Orb (15926, -0.33 DPS) [dungeon]; Tork Wrench (11855, -0.47 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.24 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -1.21 DPS) [quest]; Greater Magic Wand (11288, -1.26 DPS) [crafted] |

**New at 20:** head: Flame Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 172, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (orc, 000000000000000000-51521001230100000-000000000000000000)

Set DPS (verified): 75.2. Weights run: 3.2s. Verify run: 5.5s. 295 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.535 ± 0.011, crit=0.139 ± 0.008 per rating point (14 rating = 1%, 1.949 per %), hit=0.267 ± 0.014 per rating point (10 rating = 1%, 2.673 per %), spell_haste=0.483 ± 0.052, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.043 ± 0.043), mp5=not significant (0.202 ± 0.076)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.02 DPS) [world_drop]; Flame Circlet (253953, -0.06 DPS) [crafted]; Filigreed Flame Circlet (253979, -0.52 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.2 spell_power points (0.55 DPS) | yes | Darkspear Warding Pendant (272075, -0.09 DPS) [vendor]; Kaleidoscope Chain (13084, -0.15 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.8 spell_power points (2.35 DPS) | yes | Death Speaker Mantle (6685, -0.33 DPS) [dungeon]; Magician's Mantle (12998, -0.68 DPS) [world_drop]; Invoker's Mantle (215365, -0.71 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.03 DPS) [crafted]; Darkspear Raider's Cloak (272078, -0.13 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.17 DPS) [crafted] |
| chest | Flame Gown (253965) | Tailoring [crafted] | 16.9 spell_power points (2.88 DPS) | yes | Green Silk Armor (7065, -0.16 DPS) [crafted]; Beguiler Robes (7728, -0.20 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.43 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (2.21 DPS) | yes | Spidertank Oilrag (9448, -0.93 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.49 DPS) [world_drop]; Nightsky Wristbands (6407, -1.65 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gloves (253917, -0.01 DPS) [crafted]; Serpent Gloves (5970, -0.29 DPS) [dungeon]; Phoenix Gloves (4331, -0.56 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.6 spell_power points (2.15 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Crimson Silk Belt (7055, -0.49 DPS) [crafted]; Warsong Sash (16975, -0.83 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.0 spell_power points (3.06 DPS) | yes | Filigreed Flame Leggings (253941, -0.74 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.84 DPS) [world]; Gaze Dreamer Pants (6903, -0.95 DPS) [dungeon] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 16.9 spell_power points (2.88 DPS) | yes | Gilded Slippers (254001, -1.26 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.27 DPS) [dungeon]; Spidersilk Boots (4320, -1.32 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.7 spell_power points (1.48 DPS) | yes | Ring of Power Regulation (273030, -0.35 DPS) [dungeon]; Solid Gold Band (279845, -0.76 DPS) [quest]; Electrocutioner Lagnut (9447, -0.92 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.4 spell_power points (1.26 DPS) | yes | Ring of Power Regulation (273030, -0.14 DPS) [dungeon]; Solid Gold Band (279845, -0.54 DPS) [quest]; Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.71 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.71 DPS) [world]; Glimmering Staff (249392, -2.07 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (2.21 DPS) | yes | Orb of Mystic Insight (249394, -0.48 DPS) [crafted]; Antipodean Rod (2879, -1.19 DPS) [world_drop]; Witch's Finger (16887, -1.49 DPS) [quest] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, +0.00 DPS) [world_drop]; Wand of Mana Concentration (273043, +0.00 DPS) [world]; Greater Mystic Wand (217287, -0.12 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Flame Gown; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 295, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19545 Scout's Blade

### Band 40 (orc, 000000000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 161.3. Weights run: 3.5s. Verify run: 5.0s. 379 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.749 ± 0.021, crit=0.373 ± 0.022 per rating point (14 rating = 1%, 5.224 per %), hit=0.622 ± 0.037 per rating point (10 rating = 1%, 6.219 per %), spell_haste=not significant (0.660 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=1.034 ± 0.184, mp5=1.226 ± 0.231

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 30.0 spell_power points (10.74 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -2.41 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -3.23 DPS) [crafted]; Whitemane's Chapeau (7720, -3.36 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -0.18 DPS) [dungeon]; Amberglow Talisman (10824, -0.95 DPS) [quest]; Triune Amulet (7722, -2.36 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Mantle of Doan (7712, -1.88 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272077, -0.30 DPS) [vendor]; Glowing Thresher Cape (6901, -0.69 DPS) [dungeon]; Cloak of Blight (6832, -2.03 DPS, sim-verified) [quest] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 39.2 spell_power points (14.03 DPS) | yes | Robe of the Magi (1716, -1.65 DPS, sim-verified) [world_drop]; Robe of Power (7054, -2.85 DPS) [crafted]; Cindercloth Robe (10042, -4.37 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.65 DPS) | yes | Spidertank Oilrag (9448, -0.12 DPS) [dungeon]; Radiant Silver Bracers (4545, -1.08 DPS) [quest]; Earthen Silk Cuffs (254019, -1.37 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 28.2 spell_power points (10.10 DPS) | yes | Fiery Handwraps (254025, -3.22 DPS) [crafted]; Crimson Silk Gloves (7064, -3.48 DPS) [crafted]; Red Mageweave Gloves (10018, -3.49 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 26.2 spell_power points (9.37 DPS) | yes | Gilded Cord (254037, -2.15 DPS) [crafted]; Deathmage Sash (10771, -2.85 DPS) [dungeon]; Scorching Sash (4117, -3.39 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.51 DPS) [crafted]; Black Mageweave Leggings (9999, -0.69 DPS) [crafted]; Stoneweaver Leggings (9407, -1.87 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 31.2 spell_power points (11.17 DPS) | yes | Fiery Slippers (254005, -1.89 DPS, sim-verified) [crafted]; Furen's Boots (13100, -3.38 DPS) [world_drop]; Thoughtcast Boots (10578, -4.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (4.92 DPS) | yes | Welken Ring (5011, -0.99 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.25 DPS) [dungeon]; Runemender's Seal (273031, -1.25 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 12.7 spell_power points (4.53 DPS) | yes | Electrocutioner Lagnut (9447, -0.87 DPS) [dungeon]; Runemender's Seal (273031, -0.87 DPS) [world]; Welken Ring (5011, -1.84 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -12.56 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -14.86 DPS) [dungeon]; Celestial Stave (9517, -15.26 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -1.19 DPS) [world_drop]; Dancing Flame (6806, -1.21 DPS, sim-verified) [quest]; Twisted Nether Wand (249144, -2.04 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 252010000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 203.0. Weights run: 3.5s. Verify run: 8.2s. 483 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.047 ± 0.031, crit=0.453 ± 0.031 per rating point (14 rating = 1%, 6.343 per %), hit=0.764 ± 0.053 per rating point (10 rating = 1%, 7.639 per %), spell_haste=not significant (0.845 ± 0.292), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.002, spirit=not significant (0.425 ± 0.219), mp5=not significant (1.282 ± 0.356)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Dreadweave Hat (220907, -1.31 DPS) [vendor]; Eye of Theradras (17715, -2.18 DPS) [dungeon]; Red Mageweave Headband (10033, -2.55 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 16.4 spell_power points (5.85 DPS) | yes | Glowing Eye of Mordresh (10769, -0.35 DPS) [dungeon]; Mindburst Medallion (11196, -0.70 DPS) [quest]; Triune Amulet (7722, -1.93 DPS) [dungeon] |
| shoulder | Netherflame Shoulders (254053) | Tailoring [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Dreadweave Mantle (220905, -1.49 DPS) [vendor]; Rotgrip Mantle (17732, -1.59 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.87 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -0.52 DPS) [crafted]; Runecloth Cloak (13860, -1.95 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.24 DPS, sim-verified) [quest] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -1.06 DPS) [world_drop]; Robe of Power (7054, -1.22 DPS) [crafted]; Acumen Robes (17775, -1.73 DPS, sim-verified) [quest] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.5 spell_power points (8.03 DPS) | yes | Nethergeld Cuffs (254061, -2.19 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -2.42 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -3.08 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 32.9 spell_power points (11.76 DPS) | yes | Runecloth Gloves (13863, -2.55 DPS, sim-verified) [crafted]; Dreamweave Gloves (10019, -2.76 DPS) [crafted]; Gilded Gloves (254095, -3.22 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 28.4 spell_power points (10.17 DPS) | yes | Fiery Cord (254041, -1.25 DPS) [crafted]; Satyrmane Sash (17755, -1.41 DPS) [dungeon]; Dawnspire Cord (12466, -3.10 DPS, sim-verified) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Senior Designer's Pantaloons (11841, -0.90 DPS) [dungeon]; Red Mageweave Pants (10009, -1.05 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -1.41 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | First Sergeant's Dreadweave Boots (220909, -0.68 DPS) [vendor]; Fiery Slippers (254005, -1.77 DPS) [crafted]; Fiery Sandals (254111, -1.77 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 17.1 spell_power points (6.13 DPS) | yes | Band of the Unicorn (7553, -1.48 DPS) [world_drop]; Eye of Adaegus (5266, -2.06 DPS) [world_drop]; Snake Hoop (6750, -2.11 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (5.45 DPS) | yes | Band of the Unicorn (7553, -0.80 DPS) [world_drop]; Eye of Adaegus (5266, -1.38 DPS) [world_drop]; Snake Hoop (6750, -1.43 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, -0.17 DPS) [vendor]; Uther's Strength (11302, -0.31 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.55 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -4.46 DPS) [world_drop]; Kindling Stave (11750, -5.05 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -6.64 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -3.38 DPS) [crafted]; Noxious Shooter (17745, -3.74 DPS) [dungeon]; Ragefire Wand (7513, -3.93 DPS, sim-verified) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Horizon Choker; shoulder: Netherflame Shoulders; back: Spritecaster Cape; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Kilt of the Atal'ai Prophet; finger1: Advisor's Ring; finger2: Philanthropist's Ring; trinket2: Frozen Heart of the Mountain; ranged: Pyric Caduceus

No-known-source sample (15 of 483, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 252015000000200000-51521001230133031-003000000000000000)

Set DPS (verified): 340.9. Weights run: 3.8s. Verify run: 23.0s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.497 ± 0.036, crit=0.541 ± 0.032 per rating point (14 rating = 1%, 7.576 per %), hit=1.090 ± 0.058 per rating point (10 rating = 1%, 10.903 per %), spell_haste=not significant (0.976 ± 0.434), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.755 ± 0.463), mp5=not significant (0.367 ± 0.456)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -14.43 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.0 spell_power points (3.94 DPS) | yes | Triune Amulet (7722, -0.86 DPS) [dungeon]; Mindburst Medallion (11196, -0.98 DPS) [quest]; Lady Alizabeth's Pendant (13002, -1.31 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Silk Mantle (227104, +0.00 DPS) [pvp]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -10.50 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mageflame Cloak (13007, -0.12 DPS) [world_drop]; Cindercloth Cloak (14044, -1.05 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.17 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.96 DPS) [pvp]; Robes of Fiery Devastation (279266, -15.26 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.3 spell_power points (7.26 DPS) | yes | Netherflame Cuffs (254065, -2.56 DPS) [crafted]; General's Silk Cuffs (16538, -3.31 DPS) [pvp]; Runecloth Cuffs (254123, -3.55 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Fiery Gloves (254099, -12.56 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Cloth Girdle (20163, +0.00 DPS) [rep]; Fiery Waistcord (254085, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -12.97 DPS, sim-verified) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 45.0 spell_power points (10.45 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.97 DPS) [pvp]; Ghostweave Pants (14144, -2.68 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Venomspew Footpads (275606, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -12.60 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 26.0 spell_power points (6.04 DPS) | yes | Deep Rooted Ring (19109, -2.10 DPS) [world]; Advisor's Ring (19518, -2.45 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 18.5 spell_power points (4.29 DPS) | yes | Deep Rooted Ring (19109, -0.35 DPS) [world]; Advisor's Ring (19518, -0.70 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest]; Uther's Strength (11302, -0.71 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -0.02 DPS) [quest]; Uther's Strength (11302, -0.38 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Sageblade (22383, -0.63 DPS) [crafted]; High Warlord's War Staff (234549, -0.70 DPS) [pvp]; Elemental Mage Staff (944, -1.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brilliant Wand (249385, -5.12 DPS) [crafted]; Pyric Caduceus (11748, -6.41 DPS) [dungeon]; Stormrager (16997, -6.59 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Outrider's Silk Leggings; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Don Julio's Band; trinket2: Rune of the Guard Captain; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 253005000000200000-50520201230133031-003000000000000000)

Set DPS (verified): 723.9. Weights run: 4.3s. Verify run: 17.0s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.372 ± 0.030, crit=0.866 ± 0.043 per rating point (14 rating = 1%, 12.119 per %), hit=1.653 ± 0.077 per rating point (10 rating = 1%, 16.526 per %), spell_haste=4.201 ± 0.433, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.002, spirit=not significant (-0.022 ± 0.189), mp5=not significant (-0.007 ± 0.262)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.31 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -4.13 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.7 spell_power points (9.49 DPS) | yes | Glowing Eye of Mordresh (10769, -3.85 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -4.52 DPS) [quest]; Triune Amulet (7722, -4.66 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.0 spell_power points (20.51 DPS) | yes | Champion's Silk Mantle (227104, -1.66 DPS) [pvp]; Warlord's Silk Amice (231594, -2.04 DPS) [pvp]; Shroud of the Nathrezim (18720, -6.20 DPS, sim-verified) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | sim-verified (723.9 DPS) | yes | Spritecaster Cape (11623, -2.88 DPS) [dungeon]; Cindercloth Cloak (14044, -3.04 DPS) [crafted]; Cloak of Flames (3475, -4.20 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (723.9 DPS) | yes | Warlord's Silk Raiment (231596, -3.11 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -10.35 DPS) [pvp]; Robes of Fiery Devastation (279266, -17.20 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.0 spell_power points (15.09 DPS) | yes | Netherflame Cuffs (254065, -5.66 DPS) [crafted]; Runecloth Cuffs (254123, -6.20 DPS) [crafted]; Phoenix Bindings (210781, -7.23 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.0 spell_power points (22.32 DPS) | yes | General's Silk Handguards (16540, -3.32 DPS) [vendor]; General's Silk Gauntlets (231599, -3.32 DPS) [vendor]; Fiery Gloves (254099, -8.54 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 28.3 spell_power points (17.12 DPS) | yes | Fiery Waistcord (254085, -3.63 DPS) [crafted]; Frostwolf Cloth Belt (19090, -4.01 DPS) [rep]; Satyrmane Sash (17755, -6.42 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 35.1 spell_power points (21.18 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -4.00 DPS, sim-verified) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP reward · Horde [vendor] | 27.9 spell_power points (16.84 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.90 DPS) [pvp]; Omnicast Boots (11822, -4.67 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 28.6 spell_power points (17.30 DPS) | yes | Blackstone Ring (17713, -7.32 DPS) [dungeon]; Deep Rooted Ring (19109, -7.81 DPS) [world]; Advisor's Ring (19518, -8.85 DPS) [rep] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 20.6 spell_power points (12.45 DPS) | yes | Blackstone Ring (17713, -2.46 DPS) [dungeon]; Deep Rooted Ring (19109, -2.95 DPS) [world]; Advisor's Ring (19518, -3.99 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (723.9 DPS) | yes | Uther's Strength (11302, -5.36 DPS) [world_drop]; Ankh of Life (1713, -7.51 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (723.9 DPS) | yes | Uther's Strength (11302, -3.36 DPS) [world_drop]; Ankh of Life (1713, -5.28 DPS, sim-verified) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (723.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Elemental Mage Staff (944, -1.87 DPS) [world_drop]; Ironbark Staff (20220, -21.67 DPS, sim-verified) [rep] |
| off_hand | Tome of Fiery Arcana (19311) | Frostwolf Clan [rep] | 37.0 spell_power points (22.35 DPS) | yes | High Warlord's Tome of Destruction (234563, -4.33 DPS) [pvp]; Therazane's Touch (19315, -4.88 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -9.66 DPS) [quest] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (723.9 DPS) | yes | Brilliant Wand (249385, -0.84 DPS) [crafted]; Ragefire Wand (7513, -1.76 DPS) [quest]; Greater Eternal Wand (249237, -2.42 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Tome of Fiery Arcana

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

