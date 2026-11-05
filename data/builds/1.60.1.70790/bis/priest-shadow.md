# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 34.2. Weights run: 3.1s. Verify run: 3.7s. 183 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.823 ± 0.017, crit=0.139 ± 0.003 per rating point (14 rating = 1%, 1.949 per %), hit=0.450 ± 0.016 per rating point (10 rating = 1%, 4.497 per %), spell_haste=not significant (-0.035 ± 0.180), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.214 ± 0.203), mp5=not significant (0.598 ± 0.272)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (0.93 DPS) | yes | Pristine Circlet (253949, -0.42 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.47 DPS) [crafted]; Flying Tiger Goggles (4368, -0.85 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.4 spell_power points (0.04 DPS) | yes | Tarnished Locket (279870, -0.15 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.4 spell_power points (1.04 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.71 DPS) [crafted]; Slime-encrusted Pads (6461, -0.89 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.03 DPS) [crafted]; Caretaker's Cape (20428, -0.12 DPS) [rep]; Sanguine Cape (14376, -0.13 DPS) [world_drop] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 18.1 spell_power points (1.52 DPS) | yes | Filigreed Shadow Gown (253907, -0.19 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.70 DPS) [crafted]; Gray Woolen Robe (2585, -0.84 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 5.9 spell_power points (0.50 DPS) | yes | Repurposed Hair Band (281256, -0.32 DPS) [quest]; Mystic's Bracelets (14366, -0.36 DPS) [world_drop]; Bright Bracers (3647, -0.72 DPS, sim-verified) [world_drop] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 9.5 spell_power points (0.80 DPS) | yes | Serpent Gloves (5970, -0.21 DPS) [dungeon]; Pristine Gloves (253913, -0.25 DPS) [crafted]; Gnoll Casting Gloves (892, -0.29 DPS) [world] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 10.3 spell_power points (0.87 DPS) | yes | Pristine Sash (253925, -0.25 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.31 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.32 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 16.8 spell_power points (1.41 DPS) | yes | Filigreed Pristine Leggings (253937, -0.42 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.82 DPS) [world_drop]; Silk-threaded Trousers (1929, -0.82 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.3 spell_power points (0.87 DPS) | yes | Shadow Boots (253895, -0.12 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.41 DPS) [crafted]; Red Woolen Boots (4313, -0.44 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.6 spell_power points (0.56 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; Band of Purification (12996, -0.45 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.2 spell_power points (0.52 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; Band of Purification (12996, -0.41 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (34.2 DPS) | yes | Staff of the Blessed Seer (2271, -0.05 DPS) [dungeon]; Witching Stave (1484, -0.40 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -0.56 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 4.1 spell_power points (0.35 DPS) | yes | Tear of Grief (5611, -0.05 DPS) [quest]; Dwarven Tome (279898, -0.08 DPS) [quest]; Seer's Fine Stein (7608, -0.12 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.74 DPS, sim-verified) [world_drop]; Deepblaze (279896, -2.34 DPS) [quest]; Sizzle Stick (8071, -2.72 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 59.2. Weights run: 2.9s. Verify run: 4.2s. 302 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.551 ± 0.014, crit=0.129 ± 0.003 per rating point (14 rating = 1%, 1.809 per %), hit=0.347 ± 0.014 per rating point (10 rating = 1%, 3.470 per %), spell_haste=not significant (0.149 ± 0.123), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.584 ± 0.147), mp5=1.413 ± 0.210

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Hood (4323, -0.13 DPS) [crafted]; Holy Shroud (2721, -0.28 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.34 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.5 spell_power points (0.62 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Pendant of Myzrael (4614, -0.14 DPS) [dungeon]; Glowing Green Talisman (5002, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.92 DPS) | yes | Death Speaker Mantle (6685, -0.45 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.55 DPS) [world_drop]; Invoker's Mantle (215365, -0.58 DPS) [crafted] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Repairman's Cape (9605, -0.15 DPS) [quest]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Caretaker's Cape (19533, -0.18 DPS) [rep] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 20.3 spell_power points (2.78 DPS) | yes | Necrology Robes (2292, -0.48 DPS) [dungeon]; Shadow Gown (253967, -0.49 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.57 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.2 spell_power points (1.82 DPS) | yes | Nightsky Wristbands (6407, -1.12 DPS) [world_drop]; Mindthrust Bracers (1974, -1.19 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.21 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 12.9 spell_power points (1.76 DPS) | yes | Gloves of Old (9395, -0.42 DPS) [world_drop]; Gloves of Meditation (4318, -0.52 DPS) [crafted]; Hands of Darkness (7047, -0.97 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Belt of Arugal (6392, -0.22 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.38 DPS) [crafted]; Invoker's Cord (215366, -0.40 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 20.8 spell_power points (2.85 DPS) | yes | Filigreed Shadow Leggings (253943, -0.57 DPS) [crafted]; Necromancer Leggings (2277, -0.65 DPS) [world_drop]; Gaze Dreamer Pants (6903, -0.99 DPS, sim-verified) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 19.2 spell_power points (2.63 DPS) | yes | Gilded Slippers (254001, -0.85 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.02 DPS) [dungeon]; Glinteye Slippers (273024, -1.19 DPS) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.2 spell_power points (1.41 DPS) | yes | Snake Hoop (6750, -0.20 DPS) [quest]; Electrocutioner Lagnut (9447, -0.43 DPS) [dungeon]; Runemender's Seal (273031, -0.43 DPS) [world] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 9.8 spell_power points (1.35 DPS) | yes | Snake Hoop (6750, -0.15 DPS) [quest]; Electrocutioner Lagnut (9447, -0.38 DPS) [dungeon]; Runemender's Seal (273031, -0.38 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, -0.31 DPS) [crafted]; Rod of the Sleepwalker (1155, -0.54 DPS) [dungeon]; Soulstaff (249393, -2.90 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (1.41 DPS) | yes | Orb of Souls (249395, -0.38 DPS) [crafted]; Satyr's Rod (15962, -0.80 DPS) [world_drop]; Orb of Mistmantle (13031, -1.20 DPS, sim-verified) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, -0.26 DPS) [world]; Greater Mystic Wand (217287, -1.11 DPS) [crafted]; Gravestone Scepter (7001, -1.36 DPS, sim-verified) [quest] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Ring of Power Regulation; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 148.9. Weights run: 3.2s. Verify run: 2.9s. 389 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.148 ± 0.020, crit=0.255 ± 0.008 per rating point (14 rating = 1%, 3.568 per %), hit=0.400 ± 0.019 per rating point (10 rating = 1%, 4.001 per %), spell_haste=not significant (0.433 ± 0.170), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.917 ± 0.272), mp5=not significant (0.003 ± 0.303)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 35.8 spell_power points (6.03 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -1.74 DPS) [dungeon]; Augural Shroud (2620, -2.24 DPS) [world]; Whitemane's Chapeau (7720, -2.53 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.2 spell_power points (2.73 DPS) | yes | Triune Amulet (7722, -0.30 DPS) [dungeon]; Amberglow Talisman (10824, -1.18 DPS) [quest]; Necklace of Calisea (1714, -2.10 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.9 spell_power points (3.70 DPS) | yes | Green Silken Shoulders (7057, -0.22 DPS) [crafted]; Bloodmage Mantle (7684, -0.44 DPS) [dungeon]; Mantle of Doan (7712, -2.42 DPS, sim-verified) [dungeon] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-verified (148.9 DPS) | yes | Blackforge Cape (6424, -0.58 DPS) [dungeon]; Icy Cloak (4327, -0.64 DPS) [crafted]; Sergeant Major's Cape (16336, -0.66 DPS) [pvp] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 41.2 spell_power points (6.94 DPS) | yes | Robe of Power (7054, -1.02 DPS) [crafted]; Robe of the Magi (1716, -1.30 DPS) [world_drop]; Beguiler Robes (7728, -1.86 DPS) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 10.3 spell_power points (1.74 DPS) | yes | Nightsky Wristbands (6407, -0.12 DPS) [world_drop]; Silkstream Cuffs (16791, -0.16 DPS) [quest]; Mistscape Bracers (4045, -0.19 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 29.0 spell_power points (4.89 DPS) | yes | Shadoweave Gloves (10023, -1.06 DPS) [crafted]; Red Mageweave Gloves (10018, -1.10 DPS) [crafted]; Black Handwraps (254027, -1.18 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | 28.7 spell_power points (4.84 DPS) | yes | Deathmage Sash (10771, -0.75 DPS) [dungeon]; Gilded Cord (254037, -1.01 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.70 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.8 spell_power points (4.68 DPS) | yes | Shadow Leggings (253993, -0.19 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.82 DPS) [crafted]; Stoneweaver Leggings (9407, -3.15 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.4 spell_power points (5.13 DPS) | yes | Black Slippers (254007, -1.03 DPS, sim-verified) [crafted]; Furen's Boots (13100, -1.30 DPS) [world_drop]; Gilded Slippers (254001, -1.97 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.7 spell_power points (3.33 DPS) | yes | Welken Ring (5011, -1.28 DPS) [world_drop]; Snake Hoop (6750, -1.35 DPS) [quest]; Voodoo Band (1996, -1.51 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.7 spell_power points (2.65 DPS) | yes | Snake Hoop (6750, -0.67 DPS) [quest]; Voodoo Band (1996, -0.83 DPS) [world]; Welken Ring (5011, -2.75 DPS, sim-verified) [world_drop] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (148.9 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (148.9 DPS) | yes | Hypnotic Blade (7714, -7.11 DPS) [dungeon]; Glimmering Staff (249392, -8.68 DPS) [crafted]; Illusionary Rod (7713, -9.29 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (148.9 DPS) | yes | Goblin Igniter (5253, +0.00 DPS) [quest]; Jaina's Firestarter (13064, +0.00 DPS) [world_drop]; Burning Sliver (5249, -0.23 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Expeditionary's Cape; chest: Dreamweave Vest; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Black Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 389, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025102000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 220.8. Weights run: 3.2s. Verify run: 5.7s. 496 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.163 ± 0.017, crit=0.451 ± 0.009 per rating point (14 rating = 1%, 6.309 per %), hit=0.584 ± 0.030 per rating point (10 rating = 1%, 5.842 per %), spell_haste=not significant (0.455 ± 0.157), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.173 ± 0.132), mp5=1.100 ± 0.212

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 44.4 spell_power points (14.17 DPS) | yes | Red Mageweave Headband (10033, -0.69 DPS) [crafted]; Shadoweave Mask (10025, -2.80 DPS) [crafted]; Dreamweave Circlet (10041, -3.10 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.0 spell_power points (5.41 DPS) | yes | Glowing Eye of Mordresh (10769, -0.21 DPS) [dungeon]; Mindburst Medallion (11196, -1.00 DPS) [quest]; Seal of the Expedition (271918, -2.08 DPS) [rep] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.9 spell_power points (10.82 DPS) | yes | Netherpearl Shoulders (254055, -2.52 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.92 DPS) [vendor]; Red Mageweave Shoulders (10029, -3.03 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (220.8 DPS) | yes | Runecloth Cloak (13860, -1.18 DPS) [crafted]; Expeditionary's Cape (271921, -1.55 DPS) [rep]; Big Voodoo Cloak (8216, -2.09 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.3 spell_power points (13.48 DPS) | yes | Knight's Dreadweave Vest (220886, -3.60 DPS, sim-verified) [vendor]; Dreamweave Vest (10021, -3.63 DPS) [crafted]; Runecloth Robe (13858, -3.66 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 22.0 spell_power points (7.02 DPS) | yes | Nethergeld Cuffs (254061, -1.91 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -1.92 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.40 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 32.0 spell_power points (10.21 DPS) | yes | Runecloth Gloves (13863, -2.57 DPS, sim-verified) [crafted]; Dreamweave Gloves (10019, -2.60 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.73 DPS) [vendor] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 29.5 spell_power points (9.40 DPS) | yes | Dawnspire Cord (12466, -0.44 DPS) [dungeon]; Satyrmane Sash (17755, -1.22 DPS) [dungeon]; Deathmage Sash (10771, -1.60 DPS) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 40.0 spell_power points (12.74 DPS) | yes | Knight's Dreadweave Leggings (220888, -2.32 DPS, sim-verified) [vendor]; Red Mageweave Pants (10009, -3.83 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -4.12 DPS) [dungeon] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 30.5 spell_power points (9.72 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.96 DPS) [vendor]; Earthen Silk Slippers (254013, -1.99 DPS, sim-verified) [crafted]; Black Slippers (254007, -2.75 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.8 spell_power points (6.32 DPS) | yes | Philanthropist's Ring (281635, -1.28 DPS) [quest]; Band of the Unicorn (7553, -2.17 DPS) [world_drop]; Snake Hoop (6750, -2.55 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 16.4 spell_power points (5.23 DPS) | yes | Philanthropist's Ring (281635, -0.19 DPS) [quest]; Band of the Unicorn (7553, -1.08 DPS) [world_drop]; Snake Hoop (6750, -1.46 DPS) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (220.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -1.09 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.20 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (220.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -1.02 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.13 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (220.8 DPS) | yes | Glowing Brightwood Staff (812, -3.19 DPS, sim-verified) [world_drop]; Spellshifter Rod (9527, -4.15 DPS) [quest]; Kindling Stave (11750, -4.30 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-verified (220.8 DPS) | yes | Noxious Shooter (17745, -0.38 DPS) [dungeon]; Charged Lightning Rod (11860, -0.63 DPS) [quest]; Cairnstone Sliver (9654, -1.02 DPS) [quest] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; feet: Black Sandals; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Frozen Heart of the Mountain; ranged: Lesser Eternal Wand

No-known-source sample (15 of 496, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025105031030000000-00000000000000000-54342001201300051)

Set DPS (verified): 389.3. Weights run: 3.2s. Verify run: 21.4s. 1107 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.409 ± 0.033, crit=0.735 ± 0.015 per rating point (14 rating = 1%, 10.285 per %), hit=1.103 ± 0.049 per rating point (10 rating = 1%, 11.026 per %), spell_haste=not significant (0.464 ± 0.360), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.843 ± 0.474), mp5=2.979 ± 0.547

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (389.3 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Crown of the Penitent (13216, -16.42 DPS, sim-verified) [quest] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 26.1 spell_power points (7.88 DPS) | yes | Stormpike Sage's Pendant (19098, -0.03 DPS) [rep]; Horizon Choker (13085, -0.90 DPS) [world_drop]; Mindburst Medallion (11196, -2.24 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (389.3 DPS) | yes | Lieutenant Commander's Satin Epaulets (227136, +0.00 DPS) [vendor]; Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Theramore Shoulderpads (271953, -14.66 DPS, sim-verified) [rep] |
| back | Expeditionary's Cape (271907) | Theramore Expeditionary Force [rep] | sim-verified (389.3 DPS) | yes | Spritecaster Cape (11623, -0.03 DPS) [dungeon]; Fel Cape (279269, -0.70 DPS) [crafted]; Runecloth Cloak (13860, -2.21 DPS) [crafted] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 58.2 spell_power points (17.58 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Robe of Winter Night (14136, -22.26 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 39.2 spell_power points (11.83 DPS) | yes | Marshal's Satin Bracers (17606, -1.20 DPS) [pvp]; Netherpearl Cuffs (254067, -3.65 DPS) [crafted]; Runecloth Cuffs (254123, -4.25 DPS) [crafted] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-verified (389.3 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Theramore Handwraps (271945, -14.62 DPS, sim-verified) [rep] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (389.3 DPS) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Devout Belt (16696, -18.22 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 66.7 spell_power points (20.13 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Knight-Captain's Satin Leggings (227137, -2.01 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 38.7 spell_power points (11.68 DPS) | yes | Wolfrunner Shoes (13101, +0.00 DPS) [world_drop]; Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 32.9 spell_power points (9.94 DPS) | yes | Lorekeeper's Ring (19522, -2.12 DPS) [rep]; Band of the Hierophant (13096, -2.89 DPS) [world_drop] |
| finger2 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 26.1 spell_power points (7.88 DPS) | yes | Lorekeeper's Ring (19522, -0.05 DPS) [rep]; Band of the Hierophant (13096, -0.82 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (389.3 DPS) | yes | Enriched Seal (271903, -0.90 DPS) [rep]; Ankh of Life (1713, -2.34 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (389.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Enriched Seal (271905, -0.30 DPS) [rep] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (389.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -1.09 DPS) [world_drop]; Ironbark Staff (20069, -6.56 DPS, sim-verified) [rep] |
| off_hand | Tome of Shadow Force (19309) | Stormpike Guard [rep] | 31.0 spell_power points (9.36 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.60 DPS) [rep] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (389.3 DPS) | yes | Brilliant Wand (249385, -1.03 DPS) [crafted]; Stormrager (16997, -4.22 DPS) [quest]; Greater Eternal Wand (249237, -4.78 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Magister's Mantle; back: Expeditionary's Cape; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Sentinel's Silk Leggings; feet: Mooncloth Boots; finger1: Maiden's Circle; finger2: Deep Rooted Ring; trinket1: Shard of the Splithooves; main_hand: Hand of Edward the Odd; off_hand: Tome of Shadow Force; ranged: Torch of Light

No-known-source sample (15 of 1107, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 005105031030000000-00000000000000000-52352101201300251)

Set DPS (verified): 736.3. Weights run: 3.7s. Verify run: 19.2s. 1107 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.236 ± 0.007, crit=0.823 ± 0.016 per rating point (14 rating = 1%, 11.516 per %), hit=1.111 ± 0.050 per rating point (10 rating = 1%, 11.114 per %), spell_haste=1.916 ± 0.239, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 32.4 spell_power points (15.69 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.76 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.4 spell_power points (6.96 DPS) | yes | Glowing Eye of Mordresh (10769, -2.73 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.08 DPS) [dungeon]; Mindburst Medallion (11196, -3.37 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.6 spell_power points (15.31 DPS) | yes | Field Marshal's Satin Epaulets (231621, -0.90 DPS) [vendor]; Argent Shoulders (19059, -3.19 DPS) [crafted]; Felcloth Shoulders (14112, -3.41 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -5.38 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -6.07 DPS) [rep]; Cloak of Flames (3475, -9.15 DPS, sim-verified) [world_drop] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 42.8 spell_power points (20.76 DPS) | yes | Field Marshal's Satin Vestments (17605, -1.56 DPS) [vendor]; Field Marshal's Satin Robe (231618, -1.56 DPS) [vendor]; Felcloth Robe (14106, -3.49 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.9 spell_power points (11.58 DPS) | yes | Netherpearl Cuffs (254067, -4.48 DPS) [crafted]; Runecloth Cuffs (254123, -5.10 DPS) [crafted]; Arena Wristguards (18709, -6.00 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 34.4 spell_power points (16.67 DPS) | yes | Marshal's Satin Gloves (17608, -4.03 DPS) [vendor]; Marshal's Satin Grips (231617, -4.03 DPS) [vendor]; Earth Warder's Gloves (21318, -7.03 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 26.9 spell_power points (13.05 DPS) | yes | Black Waistcord (254087, -2.82 DPS) [crafted]; Stormpike Cloth Girdle (19094, -3.19 DPS) [rep]; Satyrmane Sash (17755, -5.12 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 57.4 spell_power points (27.82 DPS) | yes | Marshal's Satin Pants (17603, -9.91 DPS) [vendor]; Marshal's Satin Leggings (231619, -9.91 DPS) [vendor]; Felcloth Pants (14107, -13.84 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.63 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Omnicast Boots (11822, -0.57 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 22.6 spell_power points (10.97 DPS) | yes | Underworld Band (1980, -3.61 DPS) [world_drop]; Deep Rooted Ring (19109, -4.01 DPS) [world]; Lorekeeper's Ring (19522, -4.18 DPS) [rep] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (9.53 DPS) | yes | Underworld Band (1980, -2.17 DPS) [world_drop]; Deep Rooted Ring (19109, -2.57 DPS) [world]; Lorekeeper's Ring (19522, -2.74 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -1.29 DPS, sim-verified) [quest] |
| main_hand | Sageclaw (20070) | The League of Arathor [rep] | sim-verified (+7.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -3.13 DPS) [world_drop]; Ironbark Staff (20069, -7.23 DPS, sim-verified) [rep] |
| off_hand | Tome of Shadow Force (19309) | Stormpike Guard [rep] | 31.0 spell_power points (15.03 DPS) | yes | Therazane's Touch (19315, -0.97 DPS) [rep]; Grand Marshal's Tome of Power (234589, -1.42 DPS) [pvp]; Spirit of Aquementas (11904, -4.85 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Eternal Wand (249237, -0.80 DPS) [crafted]; Lesser Eternal Wand (249232, -1.28 DPS) [crafted]; Dreambough Wand (249234, -1.77 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Earthen Silk Slippers; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Sageclaw; off_hand: Tome of Shadow Force; ranged: Brilliant Wand

No-known-source sample (15 of 1107, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 36.6. Weights run: 3.1s. Verify run: 4.0s. 176 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.624 ± 0.013, crit=0.161 ± 0.004 per rating point (14 rating = 1%, 2.251 per %), hit=0.526 ± 0.014 per rating point (10 rating = 1%, 5.264 per %), spell_haste=not significant (0.141 ± 0.105), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=0.930 ± 0.177, mp5=2.079 ± 0.203

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (0.67 DPS) | yes | Pristine Circlet (253949, -0.30 DPS) [crafted]; Flying Tiger Goggles (4368, -0.44 DPS) [crafted]; Shadow Goggles (4373, -0.85 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Roadwatcher's Confidence (281265, -0.08 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.6 spell_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.27 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.43 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (20427, -0.17 DPS) [rep]; Pearl-clasped Cloak (5542, -0.23 DPS) [crafted]; Spirit Cloak (4792, -0.24 DPS) [vendor] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 17.1 spell_power points (1.04 DPS) | yes | Scholarly Robes (2034, -0.18 DPS) [dungeon]; Robe of the Moccasin (6465, -0.29 DPS) [dungeon]; Filigreed Shadow Gown (253907, -0.36 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.4 spell_power points (0.57 DPS) | yes | Crystalline Cuffs (14148, -0.40 DPS) [dungeon]; Bright Bracers (3647, -0.42 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.54 DPS, sim-verified) [quest] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 8.9 spell_power points (0.54 DPS) | yes | Magefist Gloves (12977, -0.12 DPS) [world_drop]; Bright Gloves (3066, -0.16 DPS) [world_drop]; Serpent Gloves (5970, -0.48 DPS, sim-verified) [dungeon] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 9.5 spell_power points (0.57 DPS) | yes | Pristine Sash (253925, -0.18 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Novice Ardent's Sash (253887, -0.53 DPS, sim-verified) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 18.5 spell_power points (1.12 DPS) | yes | Filigreed Pristine Leggings (253937, -0.45 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.52 DPS) [world_drop]; Colorful Kilt (10048, -0.53 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.5 spell_power points (0.57 DPS) | yes | Shadow Boots (253895, -0.10 DPS) [crafted]; Smoldering Boots (3076, -0.24 DPS) [world]; Red Woolen Boots (4313, -0.54 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.2 spell_power points (0.55 DPS) | yes | Loop of Sacrifice (281673, -0.37 DPS) [quest]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; Band of Purification (12996, -0.83 DPS, sim-verified) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Loop of Sacrifice (281673, -0.04 DPS) [quest]; Volcanic Rock Ring (12053, -0.11 DPS) [world_drop]; Band of Purification (12996, -0.18 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 37.3 spell_power points (2.26 DPS) | yes | Evocator's Blade (2567, -0.21 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.37 DPS, sim-verified) [quest]; Crescent Staff (6505, -0.42 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -1.03 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -2.66 DPS) [quest]; Greater Magic Wand (11288, -2.75 DPS) [crafted] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of the Blessed Seer; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 64.4. Weights run: 2.9s. Verify run: 4.1s. 295 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.551 ± 0.014, crit=0.129 ± 0.003 per rating point (14 rating = 1%, 1.809 per %), hit=0.347 ± 0.014 per rating point (10 rating = 1%, 3.470 per %), spell_haste=not significant (0.149 ± 0.123), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.584 ± 0.147), mp5=1.413 ± 0.210

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Hood (4323, -0.13 DPS) [crafted]; Enchanter's Cowl (4322, -0.34 DPS) [crafted]; Holy Shroud (2721, -0.64 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.5 spell_power points (0.62 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Pendant of Myzrael (4614, -0.14 DPS) [dungeon]; Glowing Green Talisman (5002, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.92 DPS) | yes | Death Speaker Mantle (6685, -0.54 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.55 DPS) [world_drop]; Invoker's Mantle (215365, -0.58 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.14 DPS) [crafted]; Battle Healer's Cloak (19529, -0.14 DPS) [rep]; Glowing Thresher Cape (6901, -0.27 DPS, sim-verified) [dungeon] |
| chest | Shadow Gown (253967) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.21 DPS, sim-verified) [dungeon]; Necrology Robes (2292, -0.33 DPS) [dungeon]; Green Silk Armor (7065, -0.42 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.2 spell_power points (1.82 DPS) | yes | Nightsky Wristbands (6407, -1.12 DPS) [world_drop]; Mindthrust Bracers (1974, -1.16 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.21 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 12.9 spell_power points (1.76 DPS) | yes | Gloves of Old (9395, -0.42 DPS) [world_drop]; Gloves of Meditation (4318, -0.52 DPS) [crafted]; Hands of Darkness (7047, -0.82 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Warsong Sash (16975, -0.23 DPS) [quest]; Belt of Arugal (6392, -0.35 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.38 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 20.8 spell_power points (2.85 DPS) | yes | Filigreed Shadow Leggings (253943, -0.57 DPS) [crafted]; Necromancer Leggings (2277, -0.65 DPS) [world_drop]; Gaze Dreamer Pants (6903, -1.12 DPS, sim-verified) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 19.2 spell_power points (2.63 DPS) | yes | Gilded Slippers (254001, -0.91 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.02 DPS) [dungeon]; Glinteye Slippers (273024, -1.19 DPS) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.2 spell_power points (1.41 DPS) | yes | Snake Hoop (6750, -0.20 DPS) [quest]; Electrocutioner Lagnut (9447, -0.43 DPS) [dungeon]; Runemender's Seal (273031, -0.43 DPS) [world] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 9.8 spell_power points (1.35 DPS) | yes | Snake Hoop (6750, -0.25 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.38 DPS) [dungeon]; Runemender's Seal (273031, -0.38 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, -0.31 DPS) [crafted]; Rod of the Sleepwalker (1155, -0.54 DPS) [dungeon]; Soulstaff (249393, -2.45 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (1.41 DPS) | yes | Orb of Souls (249395, -0.38 DPS) [crafted]; Orb of Mistmantle (13031, -0.72 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -0.80 DPS) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, -0.19 DPS) [world]; Gravestone Scepter (7001, -0.66 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -0.94 DPS) [crafted] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Shadow Gown; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Ring of Power Regulation; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 295, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 168.0. Weights run: 3.2s. Verify run: 2.7s. 382 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.148 ± 0.020, crit=0.255 ± 0.008 per rating point (14 rating = 1%, 3.568 per %), hit=0.400 ± 0.019 per rating point (10 rating = 1%, 4.001 per %), spell_haste=not significant (0.433 ± 0.170), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.917 ± 0.272), mp5=not significant (0.003 ± 0.303)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 35.8 spell_power points (6.03 DPS) | yes | Whitemane's Chapeau (7720, -1.69 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.74 DPS) [dungeon]; Augural Shroud (2620, -2.24 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.2 spell_power points (2.73 DPS) | yes | Triune Amulet (7722, -0.30 DPS) [dungeon]; Amberglow Talisman (10824, -1.18 DPS) [quest]; Necklace of Calisea (1714, -2.19 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.9 spell_power points (3.70 DPS) | yes | Green Silken Shoulders (7057, -0.22 DPS) [crafted]; Bloodmage Mantle (7684, -0.44 DPS) [dungeon]; Mantle of Doan (7712, -1.48 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (168.0 DPS) | yes | Icy Cloak (4327, -0.64 DPS) [crafted]; Cloak of Blight (6832, -0.74 DPS) [quest]; Blackforge Cape (6424, -1.08 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 41.2 spell_power points (6.94 DPS) | yes | Robe of Power (7054, -1.02 DPS) [crafted]; Robe of the Magi (1716, -1.30 DPS) [world_drop]; Beguiler Robes (7728, -1.86 DPS) [dungeon] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 13.2 spell_power points (2.22 DPS) | yes | Nightsky Wristbands (6407, -0.60 DPS) [world_drop]; Silkstream Cuffs (16791, -0.64 DPS) [quest]; Windchaser Cuffs (14429, -1.01 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 29.0 spell_power points (4.89 DPS) | yes | Shadoweave Gloves (10023, -1.06 DPS) [crafted]; Red Mageweave Gloves (10018, -1.10 DPS) [crafted]; Black Handwraps (254027, -1.18 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | 28.7 spell_power points (4.84 DPS) | yes | Deathmage Sash (10771, -0.75 DPS) [dungeon]; Gilded Cord (254037, -1.01 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.70 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.8 spell_power points (4.68 DPS) | yes | Shadow Leggings (253993, -0.19 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.82 DPS) [crafted]; Stoneweaver Leggings (9407, -2.93 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.4 spell_power points (5.13 DPS) | yes | Black Slippers (254007, -1.14 DPS, sim-verified) [crafted]; Furen's Boots (13100, -1.30 DPS) [world_drop]; Gilded Slippers (254001, -1.97 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.7 spell_power points (3.33 DPS) | yes | Welken Ring (5011, -1.28 DPS) [world_drop]; Snake Hoop (6750, -1.35 DPS) [quest]; Voodoo Band (1996, -1.51 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.7 spell_power points (2.65 DPS) | yes | Snake Hoop (6750, -0.67 DPS) [quest]; Voodoo Band (1996, -0.83 DPS) [world]; Welken Ring (5011, -2.99 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (168.0 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (168.0 DPS) | yes | Hypnotic Blade (7714, -7.11 DPS) [dungeon]; Glimmering Staff (249392, -8.68 DPS) [crafted]; Illusionary Rod (7713, -9.77 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-verified (168.0 DPS) | yes | Dancing Flame (6806, -1.00 DPS) [quest]; Goblin Igniter (5253, -1.05 DPS) [quest]; Twisted Nether Wand (249144, -1.15 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Dreamweave Vest; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Black Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025102000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 248.1. Weights run: 3.2s. Verify run: 6.0s. 489 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.163 ± 0.017, crit=0.451 ± 0.009 per rating point (14 rating = 1%, 6.309 per %), hit=0.584 ± 0.030 per rating point (10 rating = 1%, 5.842 per %), spell_haste=not significant (0.455 ± 0.157), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.173 ± 0.132), mp5=1.100 ± 0.212

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 44.4 spell_power points (14.17 DPS) | yes | Red Mageweave Headband (10033, -0.69 DPS) [crafted]; Shadoweave Mask (10025, -2.80 DPS) [crafted]; Dreamweave Circlet (10041, -3.10 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Horizon Choker (13085, -0.58 DPS, sim-verified) [world_drop]; Mindburst Medallion (11196, -0.79 DPS) [quest]; Darkspear Warding Pendant (272073, -1.86 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.9 spell_power points (10.82 DPS) | yes | Netherpearl Shoulders (254055, -2.10 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.92 DPS) [vendor]; Red Mageweave Shoulders (10029, -3.03 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.14 DPS) [dungeon]; Runecloth Cloak (13860, -1.33 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.70 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.3 spell_power points (13.48 DPS) | yes | Stone Guard's Dreadweave Vest (220904, -2.55 DPS, sim-verified) [vendor]; Dreamweave Vest (10021, -3.63 DPS) [crafted]; Runecloth Robe (13858, -3.66 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 22.0 spell_power points (7.02 DPS) | yes | Bloodband Bracers (11469, -1.92 DPS) [quest]; Nethergeld Cuffs (254061, -2.12 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -2.40 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 32.0 spell_power points (10.21 DPS) | yes | Dreamweave Gloves (10019, -2.60 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.73 DPS) [vendor]; Runecloth Gloves (13863, -2.81 DPS, sim-verified) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 29.5 spell_power points (9.40 DPS) | yes | Satyrmane Sash (17755, -1.22 DPS) [dungeon]; Deathmage Sash (10771, -1.60 DPS) [dungeon]; Dawnspire Cord (12466, -2.21 DPS, sim-verified) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 40.0 spell_power points (12.74 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -2.36 DPS, sim-verified) [vendor]; Red Mageweave Pants (10009, -3.83 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -4.12 DPS) [dungeon] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 30.5 spell_power points (9.72 DPS) | yes | Earthen Silk Slippers (254013, -1.68 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -1.96 DPS) [vendor]; Black Slippers (254007, -2.75 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.8 spell_power points (6.32 DPS) | yes | Philanthropist's Ring (281635, -1.28 DPS) [quest]; Band of the Unicorn (7553, -2.17 DPS) [world_drop]; Snake Hoop (6750, -2.55 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 16.4 spell_power points (5.23 DPS) | yes | Band of the Unicorn (7553, -1.08 DPS) [world_drop]; Snake Hoop (6750, -1.46 DPS) [quest]; Philanthropist's Ring (281635, -1.87 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.45 DPS) [quest]; Ankh of Life (1713, -1.09 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.37 DPS) [quest]; Ankh of Life (1713, -1.02 DPS) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -4.15 DPS) [quest]; Kindling Stave (11750, -4.30 DPS) [dungeon]; Glowing Brightwood Staff (812, -4.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charged Lightning Rod (11860, +0.00 DPS) [quest]; Jaina's Firestarter (13064, -0.55 DPS) [world_drop]; Nature's Breath (19118, -0.61 DPS) [quest] |

**New at 50:** head: Bad Mojo Mask; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; feet: Black Sandals; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Frozen Heart of the Mountain; ranged: Lesser Eternal Wand

No-known-source sample (15 of 489, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025105031030000000-00000000000000000-54342001201300051)

Set DPS (verified): 419.5. Weights run: 3.2s. Verify run: 22.5s. 1103 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.287 ± 0.031, crit=0.663 ± 0.016 per rating point (14 rating = 1%, 9.281 per %), hit=1.031 ± 0.042 per rating point (10 rating = 1%, 10.306 per %), spell_haste=not significant (0.864 ± 0.327), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.343 ± 0.452), mp5=2.141 ± 0.528

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Crown of the Penitent (13216, -10.15 DPS, sim-verified) [quest] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 24.9 spell_power points (4.99 DPS) | yes | Horizon Choker (13085, -1.10 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.52 DPS) [dungeon]; Frostwolf Advisor's Pendant (19096, -2.47 DPS, sim-verified) [rep] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Epaulets (227130, +0.00 DPS) [vendor]; Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -1.00 DPS, sim-verified) [vendor] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.14 DPS) [quest]; Darkspear Raider's Cloak (272063, -0.33 DPS) [vendor]; Spritecaster Cape (11623, -2.05 DPS, sim-verified) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robe of Winter Night (14136, -12.98 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 34.7 spell_power points (6.96 DPS) | yes | General's Satin Bracers (17619, -1.37 DPS) [pvp]; Runecloth Cuffs (254123, -2.17 DPS) [crafted]; Netherpearl Cuffs (254067, -2.20 DPS) [crafted] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; General's Satin Gloves (17620, +0.00 DPS) [vendor]; Raider Handwraps (272097, -9.19 DPS, sim-verified) [vendor] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.39 DPS) [crafted]; Devout Belt (16696, -10.47 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Satin Leggings (227131, +0.00 DPS) [vendor]; General's Satin Leggings (231614, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -8.68 DPS, sim-verified) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 35.4 spell_power points (7.11 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Black Sandals (254113, +0.00 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 29.1 spell_power points (5.83 DPS) | yes | Advisor's Ring (19518, -1.30 DPS) [rep]; Underworld Band (1980, -1.73 DPS) [world_drop] |
| finger2 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 24.9 spell_power points (4.99 DPS) | yes | Advisor's Ring (19518, -0.46 DPS) [rep]; Underworld Band (1980, -0.89 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, -0.43 DPS) [vendor]; Rune of the Guard Captain (19120, -1.13 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Rune of the Guard Captain (19120, -0.41 DPS) [quest]; Ankh of Life (1713, -1.48 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -1.11 DPS) [world_drop]; Ironbark Staff (20220, -5.52 DPS, sim-verified) [rep] |
| off_hand | Tome of Shadow Force (19309) | Frostwolf Clan [rep] | 31.0 spell_power points (6.22 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.40 DPS) [rep] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torch of Light (279246, -0.23 DPS) [crafted]; Stormrager (16997, -2.06 DPS) [quest]; Greater Eternal Wand (249237, -2.32 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Deep Rooted Ring; trinket1: Shard of the Splithooves; main_hand: Hand of Edward the Odd; off_hand: Tome of Shadow Force; ranged: Brilliant Wand

No-known-source sample (15 of 1103, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 005105031030000000-00000000000000000-52352101201300251)

Set DPS (verified): 730.3. Weights run: 3.7s. Verify run: 17.8s. 1103 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.236 ± 0.007, crit=0.823 ± 0.016 per rating point (14 rating = 1%, 11.516 per %), hit=1.111 ± 0.050 per rating point (10 rating = 1%, 11.114 per %), spell_haste=1.916 ± 0.239, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 32.4 spell_power points (15.69 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.98 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.4 spell_power points (6.96 DPS) | yes | Glowing Eye of Mordresh (10769, -2.83 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.08 DPS) [dungeon]; Mindburst Medallion (11196, -3.37 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.6 spell_power points (15.31 DPS) | yes | Warlord's Satin Epaulets (231611, -0.90 DPS) [vendor]; Argent Shoulders (19059, -3.19 DPS) [crafted]; Felcloth Shoulders (14112, -4.40 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-verified (730.3 DPS) | yes | Deep Woodlands Cloak (19121, -6.01 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -6.07 DPS) [rep]; Spritecaster Cape (11623, -6.77 DPS, sim-verified) [dungeon] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 42.8 spell_power points (20.76 DPS) | yes | Warlord's Satin Robes (231612, -1.56 DPS) [pvp]; Felcloth Robe (14106, -3.62 DPS, sim-verified) [crafted]; Warlord's Satin Tunic (231632, -7.38 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.9 spell_power points (11.58 DPS) | yes | Netherpearl Cuffs (254067, -4.48 DPS) [crafted]; Runecloth Cuffs (254123, -5.10 DPS) [crafted]; Arena Wristguards (18709, -6.00 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 34.4 spell_power points (16.67 DPS) | yes | General's Satin Gloves (17620, -4.03 DPS) [vendor]; General's Satin Grips (231613, -4.03 DPS) [vendor]; Earth Warder's Gloves (21318, -7.33 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 26.9 spell_power points (13.05 DPS) | yes | Black Waistcord (254087, -2.82 DPS) [crafted]; Frostwolf Cloth Belt (19090, -3.19 DPS) [rep]; Satyrmane Sash (17755, -5.12 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 32.5 spell_power points (15.74 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Stone Guard's Dreadweave Leggings (220906, -2.97 DPS) [vendor]; Felcloth Pants (14107, -3.11 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.63 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Omnicast Boots (11822, -0.57 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 22.6 spell_power points (10.97 DPS) | yes | Underworld Band (1980, -3.61 DPS) [world_drop]; Deep Rooted Ring (19109, -4.01 DPS) [world]; Advisor's Ring (19518, -4.18 DPS) [rep] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (9.53 DPS) | yes | Underworld Band (1980, -2.26 DPS, sim-verified) [world_drop]; Deep Rooted Ring (19109, -2.57 DPS) [world]; Advisor's Ring (19518, -2.74 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (730.3 DPS) | yes | Uther's Strength (11302, -1.94 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (730.3 DPS) | yes | Uther's Strength (11302, -0.86 DPS) [world_drop]; Ankh of Life (1713, -2.21 DPS, sim-verified) [world_drop] |
| main_hand | Mindfang (20214) | The Defilers [rep] | sim-verified (730.3 DPS) | yes | Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -4.11 DPS, sim-verified) [world_drop] |
| off_hand | Tome of Shadow Force (19309) | Frostwolf Clan [rep] | 31.0 spell_power points (15.03 DPS) | yes | Therazane's Touch (19315, -0.97 DPS) [rep]; High Warlord's Tome of Destruction (234563, -1.42 DPS) [pvp]; Spirit of Aquementas (11904, -4.85 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (730.3 DPS) | yes | Greater Eternal Wand (249237, -0.80 DPS) [crafted]; Lesser Eternal Wand (249232, -1.28 DPS) [crafted]; Dreambough Wand (249234, -1.77 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Earthen Silk Slippers; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Mindfang; off_hand: Tome of Shadow Force; ranged: Brilliant Wand

No-known-source sample (15 of 1103, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

