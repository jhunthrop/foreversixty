# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 020000000000000000-00000000000000000-40500000000000000)

Set DPS (verified): 41.2. Weights run: 3.3s. Verify run: 2.8s. 203 eligible items had no known source.

Race sweep: human 41.2 ± 0.0, gnome 41.0 ± 0.0, night-elf 40.5 ± 0.0, dwarf 40.5 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.854 ± 0.017, crit=0.160 ± 0.003 per rating point (14 rating = 1%, 2.233 per %), hit=0.491 ± 0.016 per rating point (10 rating = 1%, 4.914 per %), spell_haste=not significant (0.456 ± 0.305), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.431 ± 0.219), mp5=not significant (1.050 ± 0.268)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (0.99 DPS) | yes | Pristine Circlet (253949, -0.45 DPS) [crafted]; Flying Tiger Goggles (4368, -0.83 DPS) [crafted]; Shadow Goggles (4373, -1.06 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.9 spell_power points (0.08 DPS) | yes | Tarnished Locket (279870, -0.25 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.7 spell_power points (1.14 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.69 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.78 DPS) [crafted]; Slime-encrusted Pads (6461, -0.86 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.17 DPS) [rep]; Sanguine Cape (14376, -0.21 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.42 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 18.3 spell_power points (1.64 DPS) | yes | Filigreed Shadow Gown (253907, -0.44 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.69 DPS) [crafted]; Gray Woolen Robe (2585, -0.90 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.4 spell_power points (0.67 DPS) | yes | Repurposed Hair Band (281256, -0.44 DPS) [quest]; Seer's Cuffs (3645, -0.51 DPS) [dungeon]; Bright Bracers (3647, -1.48 DPS, sim-verified) [world_drop] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 9.6 spell_power points (0.86 DPS) | yes | Pristine Gloves (253913, -0.27 DPS) [crafted]; Gnoll Casting Gloves (892, -0.32 DPS) [world]; Serpent Gloves (5970, -0.55 DPS, sim-verified) [dungeon] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 10.4 spell_power points (0.94 DPS) | yes | Pristine Sash (253925, -0.28 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.32 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.35 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 17.9 spell_power points (1.60 DPS) | yes | Filigreed Pristine Leggings (253937, -0.47 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.83 DPS) [world_drop]; Colorful Kilt (10048, -0.96 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.4 spell_power points (0.94 DPS) | yes | Shadow Boots (253895, -0.17 DPS) [crafted]; Red Woolen Boots (4313, -0.38 DPS) [crafted]; Pristine Boots (253889, -0.44 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.1 spell_power points (0.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Band of Purification (12996, -0.41 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.7 spell_power points (0.60 DPS) | yes | Band of Purification (12996, -0.37 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.41 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 32.3 spell_power points (2.90 DPS) | yes | Witching Stave (1484, -0.09 DPS) [dungeon]; Evocator's Blade (2567, -0.21 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -0.75 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -0.85 DPS, sim-verified) [dungeon]; Deepblaze (279896, -2.75 DPS) [quest]; Sizzle Stick (8071, -3.37 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of the Blessed Seer; ranged: Skycaller

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 020000000000000000-00000000000000000-50542001200000000)

Set DPS (verified): 66.1. Weights run: 3.1s. Verify run: 4.9s. 325 eligible items had no known source.

Race sweep: gnome 66.2 ± 0.1, human 66.1 ± 0.1, dwarf 65.4 ± 0.1, night-elf 65.3 ± 0.1 (dps); gnome beats the guide race human by 0.1 (0.10%) ± 0.1, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.441 ± 0.011, crit=0.158 ± 0.003 per rating point (14 rating = 1%, 2.212 per %), hit=0.443 ± 0.013 per rating point (10 rating = 1%, 4.431 per %), spell_haste=not significant (0.020 ± 0.057), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=-1.776 ± 0.134, mp5=not significant (-0.539 ± 0.173)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 14.0 spell_power points (1.34 DPS) | yes | Shadow Hood (4323, -0.21 DPS) [crafted]; Holy Shroud (2721, -0.29 DPS) [world_drop]; Shadow Circlet (253955, -0.29 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.6 spell_power points (0.25 DPS) | yes | Seal of the Expedition (271920, -0.04 DPS) [rep]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (1.24 DPS) | yes | Invoker's Mantle (215365, -0.36 DPS) [crafted]; Magician's Mantle (12998, -0.38 DPS) [world_drop]; Death Speaker Mantle (6685, -0.48 DPS, sim-verified) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.10 DPS) [crafted]; Caretaker's Cape (19533, -0.10 DPS) [rep] |
| chest | Shadow Gown (253967) | Tailoring [crafted] | sim-verified (66.1 DPS) | yes | Green Silk Armor (7065, -0.13 DPS) [crafted]; Beguiler Robes (7728, -0.17 DPS) [dungeon]; Necrology Robes (2292, -0.28 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Nightsky Wristbands (6407, -0.61 DPS) [world_drop]; Stonecloth Bindings (14416, -0.65 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.98 DPS, sim-verified) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 12.1 spell_power points (1.15 DPS) | yes | Shadow Gloves (253919, -0.36 DPS) [crafted]; Serpent Gloves (5970, -0.49 DPS) [dungeon]; Hands of Darkness (7047, -0.65 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.3 spell_power points (1.18 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.30 DPS) [crafted]; Crimson Silk Belt (7055, -0.31 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 17.1 spell_power points (1.63 DPS) | yes | Filigreed Shadow Leggings (253943, -0.33 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.49 DPS) [dungeon]; Necromancer Leggings (2277, -0.92 DPS, sim-verified) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 16.1 spell_power points (1.54 DPS) | yes | Spidersilk Boots (4320, -0.70 DPS) [crafted]; Acidic Walkers (9454, -0.72 DPS) [dungeon]; Gilded Slippers (254001, -0.80 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.2 spell_power points (0.78 DPS) | yes | Ring of Power Regulation (273030, -0.21 DPS) [dungeon]; Minor Channeling Ring (1449, -0.22 DPS) [quest]; Runemender's Seal (273031, -0.24 DPS) [world] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.67 DPS) | yes | Ring of Power Regulation (273030, -0.10 DPS) [dungeon]; Minor Channeling Ring (1449, -0.11 DPS) [quest]; Runemender's Seal (273031, -0.13 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Shade (2549, -0.29 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.38 DPS) [world]; Glimmering Staff (249392, -1.89 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.6 spell_power points (0.92 DPS) | yes | Orb of Mistmantle (13031, -0.51 DPS, sim-verified) [world_drop]; Eye of Paleth (2943, -0.54 DPS) [quest]; Orb of Souls (249395, -0.54 DPS) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, -0.87 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -1.79 DPS) [dungeon]; Gravestone Scepter (7001, -2.09 DPS) [quest] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Shadow Gown; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 325, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish

### Band 40 (human, 020000000000000000-00000000000000000-52542001201300040)

Set DPS (verified): 95.7. Weights run: 3.3s. Verify run: 3.5s. 461 eligible items had no known source.

Race sweep: gnome 96.9 ± 0.1, human 95.7 ± 0.1, dwarf 94.5 ± 0.1, night-elf 94.2 ± 0.1 (dps); gnome beats the guide race human by 1.3 (1.33%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.860 ± 0.024, crit=0.236 ± 0.006 per rating point (14 rating = 1%, 3.302 per %), hit=0.668 ± 0.027 per rating point (10 rating = 1%, 6.679 per %), spell_haste=not significant (0.397 ± 0.239), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.574 ± 0.224), mp5=1.719 ± 0.406

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 27.8 spell_power points (3.63 DPS) | yes | Miner's Hat of the Deep (9429, -0.15 DPS) [dungeon]; Corpseshroud (10574, -0.58 DPS) [dungeon]; Whitemane's Chapeau (7720, -0.69 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.9 spell_power points (1.81 DPS) | yes | Necklace of Calisea (1714, -0.50 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.64 DPS) [quest]; Triune Amulet (7722, -0.92 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.2 spell_power points (2.37 DPS) | yes | Mantle of Doan (7712, -0.19 DPS) [dungeon]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Green Silken Shoulders (7057, -0.66 DPS, sim-verified) [crafted] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.19 DPS) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Icy Cloak (4327, -1.17 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 33.8 spell_power points (4.41 DPS) | yes | Robe of Power (7054, -0.64 DPS) [crafted]; Robe of the Magi (1716, -0.81 DPS, sim-verified) [world_drop]; Shadoweave Robe (10004, -1.14 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.2 spell_power points (1.85 DPS) | yes | Arcane Runed Bracers (4744, -0.67 DPS) [quest]; Windchaser Cuffs (14429, -0.84 DPS) [world_drop]; Mindthrust Bracers (1974, -1.21 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 25.5 spell_power points (3.32 DPS) | yes | Shadoweave Gloves (10023, -0.55 DPS, sim-verified) [crafted]; Black Handwraps (254027, -0.71 DPS) [crafted]; Red Mageweave Gloves (10018, -0.76 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | 24.3 spell_power points (3.18 DPS) | yes | Deathmage Sash (10771, -0.61 DPS, sim-verified) [dungeon]; Gilded Cord (254037, -0.78 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.90 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.3 spell_power points (3.17 DPS) | yes | Shadoweave Pants (10002, -0.43 DPS) [crafted]; Stoneweaver Leggings (9407, -0.46 DPS) [dungeon]; Shadow Leggings (253993, -0.94 DPS, sim-verified) [crafted] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 21.3 spell_power points (2.78 DPS) | yes | Furen's Boots (13100, -0.80 DPS) [world_drop]; Gilded Slippers (254001, -0.80 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.93 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 18.3 spell_power points (2.39 DPS) | yes | Lorekeeper's Ring (19524, -0.54 DPS) [rep]; Runemender's Seal (273031, -0.87 DPS) [world]; Ring of Power Regulation (273030, -0.93 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.87 DPS) | yes | Runemender's Seal (273031, -0.35 DPS) [world]; Ring of Power Regulation (273030, -0.41 DPS) [dungeon]; Lorekeeper's Ring (19524, -0.70 DPS, sim-verified) [rep] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -5.27 DPS) [dungeon]; Illusionary Rod (7713, -5.87 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -6.46 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Plaguerot Sprig (10766) | Razorfen Downs: Plaguemaw the Rotting [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, +0.00 DPS) [world_drop]; Earthen Rod (9381, -2.70 DPS) [dungeon]; Umbral Wand (5216, -3.05 DPS) [dungeon] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Expeditionary's Cape; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Black Cord; legs: Red Mageweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Plaguerot Sprig

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 025003000000000000-00000000000000000-52542001201300051)

Set DPS (verified): 228.4. Weights run: 3.1s. Verify run: 5.6s. 608 eligible items had no known source.

Race sweep: gnome 232.1 ± 0.3, human 228.4 ± 0.3, night-elf 223.9 ± 0.3, dwarf 223.2 ± 0.3 (dps); gnome beats the guide race human by 3.8 (1.65%) ± 0.4, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.540 ± 0.027, crit=0.568 ± 0.013 per rating point (14 rating = 1%, 7.949 per %), hit=0.820 ± 0.035 per rating point (10 rating = 1%, 8.196 per %), spell_haste=not significant (-0.029 ± 0.265), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.159 ± 0.227), mp5=2.064 ± 0.347

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 52.3 spell_power points (13.36 DPS) | yes | Red Mageweave Headband (10033, -0.65 DPS) [crafted]; Chief Architect's Monocle (11839, -1.71 DPS) [dungeon]; Soulcatcher Halo (10630, -2.00 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 22.2 spell_power points (5.66 DPS) | yes | Glowing Eye of Mordresh (10769, -0.73 DPS) [dungeon]; Gemshard Heart (17707, -0.85 DPS) [dungeon]; Mindburst Medallion (11196, -1.57 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 40.7 spell_power points (10.39 DPS) | yes | Kentic Amice (11624, -1.47 DPS) [dungeon]; Mantle of Lost Hope (22234, -1.66 DPS) [dungeon]; Netherpearl Shoulders (254055, -2.66 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -0.47 DPS) [rep]; Runecloth Cloak (13860, -0.73 DPS) [crafted]; Big Voodoo Cloak (8216, -1.36 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 47.9 spell_power points (12.24 DPS) | yes | Runecloth Robe (13858, -2.39 DPS, sim-verified) [crafted]; Acumen Robes (17775, -2.77 DPS) [quest]; Knight's Dreadweave Vest (220886, -2.82 DPS) [vendor] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 27.1 spell_power points (6.91 DPS) | yes | Netherpearl Cuffs (254067, -0.64 DPS) [crafted]; Bloodband Bracers (11469, -1.98 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.08 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 35.7 spell_power points (9.10 DPS) | yes | Silkweb Gloves (11634, -1.54 DPS, sim-verified) [dungeon]; Runecloth Gloves (13863, -2.14 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.25 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 35.6 spell_power points (9.08 DPS) | yes | Dawnspire Cord (12466, -0.09 DPS) [dungeon]; Black Waistcord (254087, -0.70 DPS) [crafted]; Deathmage Sash (10771, -1.40 DPS) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 44.5 spell_power points (11.35 DPS) | yes | Spellshock Leggings (9484, -1.55 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -1.73 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.78 DPS) [dungeon] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 33.9 spell_power points (8.64 DPS) | yes | Coldstone Slippers (18697, -1.03 DPS) [dungeon]; Sergeant Major's Dreadweave Boots (220891, -1.79 DPS, sim-verified) [vendor]; Southsea Mojo Boots (20641, -2.28 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 27.2 spell_power points (6.94 DPS) | yes | Underworld Band (1980, -1.40 DPS) [world_drop]; Woodseed Hoop (17768, -1.48 DPS) [quest]; Cyclopean Band (11824, -1.73 DPS) [dungeon] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 22.6 spell_power points (5.76 DPS) | yes | Underworld Band (1980, -0.22 DPS) [world_drop]; Woodseed Hoop (17768, -0.30 DPS) [quest]; Cyclopean Band (11824, -0.55 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.35 DPS) [world_drop]; Blessed Prayer Beads (19990, -0.86 DPS) [quest]; Ankh of Life (1713, -1.40 DPS) [world_drop] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.10 DPS) [world_drop]; Blessed Prayer Beads (19990, -1.61 DPS) [quest]; Ankh of Life (1713, -2.15 DPS) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -2.19 DPS) [world_drop]; Kindling Stave (11750, -4.66 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Woestave (20082) | Blood of Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -2.21 DPS) [dungeon]; Lesser Eternal Wand (249232, -3.62 DPS) [crafted]; Pyric Caduceus (11748, -8.24 DPS, sim-verified) [dungeon] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Black Gloves; waist: Ban'thok Sash; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Frozen Heart of the Mountain; trinket2: Enriched Seal; ranged: Woestave

No-known-source sample (15 of 608, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 025105031030000000-00000000000000000-52542001201300051)

Set DPS (verified): 426.1. Weights run: 3.2s. Verify run: 22.9s. 1525 eligible items had no known source.

Race sweep: gnome 426.3 ± 0.4, human 426.1 ± 0.4, night-elf 420.9 ± 0.4, dwarf 420.7 ± 0.4 (dps); gnome beats the guide race human by 0.2 (0.04%) ± 0.6, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.976 ± 0.025, crit=0.739 ± 0.015 per rating point (14 rating = 1%, 10.340 per %), hit=0.908 ± 0.039 per rating point (10 rating = 1%, 9.083 per %), spell_haste=not significant (0.361 ± 0.288), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.378 ± 0.341), mp5=not significant (1.498 ± 0.419)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Spellweaver's Turban (22267, -8.77 DPS, sim-verified) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 29.5 spell_power points (9.51 DPS) | yes | Dark Advisor's Pendant (18691, -0.56 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.71 DPS) [world_drop]; Nacreous Shell Necklace (22403, -1.18 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Epaulets (227136, -1.07 DPS) [vendor]; Theramore Shoulderpads (271953, -1.23 DPS, sim-verified) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.89 DPS) [world]; Spritecaster Cape (11623, -3.47 DPS) [dungeon]; Shroud of the Exile (15421, -4.05 DPS, sim-verified) [quest] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Robe of Winter Night (14136, -10.28 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 32.5 spell_power points (10.45 DPS) | yes | Marshal's Satin Bracers (17606, -3.26 DPS) [pvp]; Netherpearl Cuffs (254067, -3.46 DPS) [crafted]; Manacle Cuffs (11962, -3.48 DPS) [quest] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Gloves of Spell Mastery (14146, -8.20 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Highlander's Cloth Girdle (20047, +0.00 DPS) [rep]; Ban'thok Sash (11662, -10.52 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, +0.00 DPS) [vendor] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 33.6 spell_power points (10.83 DPS) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -3.20 DPS) [quest]; Seal of Rivendare (13345, -3.28 DPS) [dungeon]; Naglering (11669, -13.80 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.73 DPS) [quest]; Seal of Rivendare (13345, -0.81 DPS) [dungeon]; Brainlash (6440, -1.43 DPS) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+9.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Smolderweb's Eye (13213, -0.93 DPS) [dungeon] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -0.99 DPS) [dungeon]; Ironbark Staff (20069, -7.29 DPS, sim-verified) [rep] |
| off_hand | Book of the Dead (13353) | Stratholme: Balnazzar [dungeon] | 34.4 spell_power points (11.08 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Therazane's Touch (19315, -1.75 DPS) [rep]; Tome of Shadow Force (19309, -2.12 DPS, sim-verified) [rep] |
| ranged | Oblivion's Touch (18761) | Scholomance: Death Knight Darkreaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brilliant Wand (249385, -0.98 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -1.11 DPS) [dungeon]; Bonecreeper Stylus (13938, -2.00 DPS, sim-verified) [dungeon] |

**New at 60:** head: Magister's Crown; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Magister's Leggings; feet: Dragonrider Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Second Wind; trinket2: Burst of Knowledge; main_hand: Hand of Edward the Odd; off_hand: Book of the Dead; ranged: Oblivion's Touch

No-known-source sample (15 of 1525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005105031030000000-00000000000000000-50052201211312251)

Set DPS (verified): 797.8. Weights run: 3.5s. Verify run: 21.2s. 1525 eligible items had no known source.

Race sweep: human 797.8 ± 0.5, night-elf 797.4 ± 0.5, dwarf 796.8 ± 0.5, gnome 794.6 ± 0.5 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.210 ± 0.006, crit=0.736 ± 0.014 per rating point (14 rating = 1%, 10.300 per %), hit=0.832 ± 0.036 per rating point (10 rating = 1%, 8.316 per %), spell_haste=0.973 ± 0.176, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 41.3 spell_power points (17.69 DPS) | yes | Field Marshal's Headdress (17602, -1.04 DPS) [vendor]; Field Marshal's Satin Crown (231616, -1.04 DPS) [vendor]; Felcloth Hood (14111, -5.29 DPS, sim-verified) [crafted] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 21.7 spell_power points (9.28 DPS) | yes | Nacreous Shell Necklace (22403, -1.15 DPS) [dungeon]; Emberfury Talisman (12929, -2.31 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.86 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.0 spell_power points (12.86 DPS) | yes | Field Marshal's Satin Epaulets (231621, -0.35 DPS) [vendor]; Argent Shoulders (19059, -2.16 DPS) [crafted]; Felcloth Shoulders (14112, -4.35 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.67 DPS, sim-verified) [world]; Spritecaster Cape (11623, -4.60 DPS) [dungeon]; Shroud of the Exile (15421, -5.09 DPS) [quest] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 42.5 spell_power points (18.21 DPS) | yes | Field Marshal's Satin Vestments (17605, -1.56 DPS) [vendor]; Field Marshal's Satin Robe (231618, -1.56 DPS) [vendor]; Felcloth Robe (14106, -3.35 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.7 spell_power points (10.14 DPS) | yes | Netherpearl Cuffs (254067, -3.94 DPS) [crafted]; Runecloth Cuffs (254123, -4.53 DPS) [crafted]; Arena Wristguards (18709, -5.73 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.7 spell_power points (13.58 DPS) | yes | Marshal's Satin Gloves (17608, -2.56 DPS) [vendor]; Marshal's Satin Grips (231617, -2.56 DPS) [vendor]; Hands of Power (13253, -6.13 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 25.6 spell_power points (10.95 DPS) | yes | Ban'thok Sash (11662, -1.35 DPS) [dungeon]; Black Waistcord (254087, -2.00 DPS) [crafted]; Stormpike Cloth Girdle (19094, -2.34 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.7 spell_power points (23.42 DPS) | yes | Marshal's Satin Pants (17603, -7.83 DPS) [vendor]; Marshal's Satin Leggings (231619, -7.83 DPS) [vendor]; Skyshroud Leggings (13170, -13.44 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 22.5 spell_power points (9.65 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Venomspew Footpads (275606, -0.23 DPS) [crafted] |
| finger1 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.78 DPS) [quest]; Maiden's Circle (13001, -0.78 DPS) [world_drop]; Naglering (11669, -11.69 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.34 DPS) [quest]; Maiden's Circle (13001, -0.34 DPS) [world_drop]; Don Julio's Band (19325, -0.70 DPS) [rep] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Smolderweb's Eye (13213, -4.80 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.35 DPS, sim-verified) [crafted]; Smolderweb's Eye (13213, -4.15 DPS) [dungeon]; Serenity Field (272439, -4.71 DPS) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.28 DPS) [pvp]; Headmaster's Charge (13937, -14.01 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -15.02 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Woestave (20082, -0.45 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, -0.45 DPS) [dungeon]; Oblivion's Touch (18761, -0.68 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Don Mauricio's Band of Domination; finger2: Channeler's Ring; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 020000000000000000-00000000000000000-40500000000000000)

Set DPS (verified): 44.1. Weights run: 3.3s. Verify run: 2.5s. 196 eligible items had no known source.

Race sweep: undead 44.1 ± 0.0, troll 40.0 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.854 ± 0.017, crit=0.160 ± 0.003 per rating point (14 rating = 1%, 2.233 per %), hit=0.491 ± 0.016 per rating point (10 rating = 1%, 4.914 per %), spell_haste=not significant (0.456 ± 0.305), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.431 ± 0.219), mp5=not significant (1.050 ± 0.268)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (0.99 DPS) | yes | Pristine Circlet (253949, -0.45 DPS) [crafted]; Flying Tiger Goggles (4368, -0.83 DPS) [crafted]; Shadow Goggles (4373, -0.95 DPS, sim-verified) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.3 spell_power points (0.12 DPS) | yes | Scholarly Pendant (277203, -0.04 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.7 spell_power points (1.14 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.73 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.78 DPS) [crafted]; Slime-encrusted Pads (6461, -0.86 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (20427, -0.17 DPS) [rep]; Sanguine Cape (14376, -0.21 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.36 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 18.3 spell_power points (1.64 DPS) | yes | Filigreed Shadow Gown (253907, -0.54 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.69 DPS) [crafted]; Gray Woolen Robe (2585, -0.90 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.4 spell_power points (0.67 DPS) | yes | Tabitha's Cuffs (251486, -0.40 DPS) [quest]; Owlbeard Bracers (16981, -0.42 DPS) [quest]; Bright Bracers (3647, -0.84 DPS, sim-verified) [world_drop] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 9.6 spell_power points (0.86 DPS) | yes | Pristine Gloves (253913, -0.27 DPS) [crafted]; Gnoll Casting Gloves (892, -0.32 DPS) [world]; Serpent Gloves (5970, -0.45 DPS, sim-verified) [dungeon] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 10.4 spell_power points (0.94 DPS) | yes | Pristine Sash (253925, -0.31 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.32 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.35 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 17.9 spell_power points (1.60 DPS) | yes | Filigreed Pristine Leggings (253937, -0.51 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.83 DPS) [world_drop]; Colorful Kilt (10048, -0.96 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.4 spell_power points (0.94 DPS) | yes | Shadow Boots (253895, -0.33 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.38 DPS) [crafted]; Pristine Boots (253889, -0.44 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.1 spell_power points (0.64 DPS) | yes | Loop of Sacrifice (281673, -0.25 DPS) [quest]; Band of Purification (12996, -0.41 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.1 spell_power points (0.46 DPS) | yes | Band of Purification (12996, -0.23 DPS) [world_drop]; Loop of Sacrifice (281673, -0.23 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 32.3 spell_power points (2.90 DPS) | yes | Evocator's Blade (2567, -0.04 DPS) [dungeon]; Witching Stave (1484, -0.09 DPS) [dungeon]; Crescent Staff (6505, -0.49 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -0.28 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -3.28 DPS) [quest]; Greater Magic Wand (11288, -3.41 DPS) [crafted] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of the Blessed Seer; ranged: Skycaller

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (undead, 020000000000000000-00000000000000000-50542001200000000)

Set DPS (verified): 71.8. Weights run: 3.1s. Verify run: 4.6s. 318 eligible items had no known source.

Race sweep: undead 71.8 ± 0.1, troll 64.7 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.618 ± 0.017, crit=0.170 ± 0.004 per rating point (14 rating = 1%, 2.381 per %), hit=0.463 ± 0.019 per rating point (10 rating = 1%, 4.634 per %), spell_haste=not significant (0.238 ± 0.174), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=0.844 ± 0.209, mp5=3.013 ± 0.310

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Hood (4323, -0.03 DPS) [crafted]; Enchanter's Cowl (4322, -0.25 DPS) [crafted]; Holy Shroud (2721, -0.88 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.8 spell_power points (0.79 DPS) | yes | Glowing Green Talisman (5002, -0.11 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.11 DPS) [world_drop]; Pendant of Myzrael (4614, -0.36 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.6 spell_power points (1.96 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.54 DPS) [world_drop]; Invoker's Mantle (215365, -0.60 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272078, -0.12 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Glowing Thresher Cape (6901, -0.46 DPS, sim-verified) [dungeon] |
| chest | Shadow Gown (253967) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Necrology Robes (2292, -0.49 DPS) [dungeon]; Green Silk Armor (7065, -0.49 DPS) [crafted]; Beguiler Robes (7728, -0.52 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.0 spell_power points (2.43 DPS) | yes | Mindthrust Bracers (1974, -1.09 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.59 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.76 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 13.3 spell_power points (1.80 DPS) | yes | Gloves of Old (9395, -0.01 DPS) [world_drop]; Gloves of Meditation (4318, -0.33 DPS) [crafted]; Hands of Darkness (7047, -0.38 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Warsong Sash (16975, -0.25 DPS) [quest]; Crimson Silk Belt (7055, -0.34 DPS) [crafted]; Belt of Arugal (6392, -0.60 DPS, sim-verified) [dungeon] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 22.5 spell_power points (3.04 DPS) | yes | Gaze Dreamer Pants (6903, -0.40 DPS) [dungeon]; Filigreed Shadow Leggings (253943, -0.60 DPS) [crafted]; Leech Pants (6910, -1.28 DPS, sim-verified) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 20.7 spell_power points (2.79 DPS) | yes | Gilded Slippers (254001, -0.92 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.00 DPS) [dungeon]; Glinteye Slippers (273024, -1.16 DPS) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 15.0 spell_power points (2.03 DPS) | yes | Runemender's Seal (273031, -0.44 DPS) [world]; Snake Hoop (6750, -0.80 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 13.0 spell_power points (1.76 DPS) | yes | Runemender's Seal (273031, -0.36 DPS, sim-verified) [world]; Snake Hoop (6750, -0.53 DPS) [quest]; Electrocutioner Lagnut (9447, -0.55 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rod of the Sleepwalker (1155, -0.28 DPS) [dungeon]; Glimmering Staff (249392, -0.35 DPS) [crafted]; Soulstaff (249393, -2.12 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.7 spell_power points (1.44 DPS) | yes | Orb of Souls (249395, -0.22 DPS) [crafted]; Orb of Mistmantle (13031, -0.49 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -0.72 DPS) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, -0.64 DPS, sim-verified) [world]; Gravestone Scepter (7001, -1.30 DPS) [quest]; Gyromatic Icemaker (9489, -1.70 DPS) [dungeon] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Shadow Gown; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Ring of Power Regulation; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (undead, 020000000000000000-00000000000000000-52542001201300040)

Set DPS (verified): 107.1. Weights run: 3.3s. Verify run: 4.0s. 454 eligible items had no known source.

Race sweep: undead 107.1 ± 0.1, troll 94.1 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.860 ± 0.024, crit=0.236 ± 0.006 per rating point (14 rating = 1%, 3.302 per %), hit=0.668 ± 0.027 per rating point (10 rating = 1%, 6.679 per %), spell_haste=not significant (0.397 ± 0.239), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.574 ± 0.224), mp5=1.719 ± 0.406

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Papal Fez (9431, -0.38 DPS, sim-verified) [dungeon]; Corpseshroud (10574, -0.43 DPS) [dungeon]; Whitemane's Chapeau (7720, -0.54 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.9 spell_power points (1.81 DPS) | yes | Triune Amulet (7722, -0.24 DPS) [dungeon]; Necklace of Calisea (1714, -0.50 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.64 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.2 spell_power points (2.37 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Mantle of Doan (7712, -0.19 DPS) [dungeon]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Icy Cloak (4327, -0.17 DPS) [crafted]; Long Silken Cloak (4326, -0.19 DPS) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 33.8 spell_power points (4.41 DPS) | yes | Robe of the Magi (1716, -0.49 DPS) [world_drop]; Robe of Power (7054, -0.64 DPS) [crafted]; Shadoweave Robe (10004, -1.14 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.2 spell_power points (1.85 DPS) | yes | Radiant Silver Bracers (4545, -0.43 DPS) [quest]; Mindthrust Bracers (1974, -0.61 DPS) [dungeon]; Windchaser Cuffs (14429, -0.84 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 25.5 spell_power points (3.32 DPS) | yes | Shadoweave Gloves (10023, -0.54 DPS) [crafted]; Black Handwraps (254027, -0.71 DPS) [crafted]; Red Mageweave Gloves (10018, -0.76 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | 24.3 spell_power points (3.18 DPS) | yes | Deathmage Sash (10771, -0.58 DPS) [dungeon]; Gilded Cord (254037, -0.78 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.90 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.3 spell_power points (3.17 DPS) | yes | Shadow Leggings (253993, -0.19 DPS) [crafted]; Shadoweave Pants (10002, -0.43 DPS) [crafted]; Stoneweaver Leggings (9407, -0.46 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 21.3 spell_power points (2.78 DPS) | yes | Furen's Boots (13100, -0.80 DPS) [world_drop]; Gilded Slippers (254001, -0.91 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.93 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 18.3 spell_power points (2.39 DPS) | yes | Advisor's Ring (19520, -0.54 DPS) [rep]; Runemender's Seal (273031, -0.87 DPS) [world]; Ring of Power Regulation (273030, -0.93 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.87 DPS) | yes | Advisor's Ring (19520, -0.02 DPS) [rep]; Runemender's Seal (273031, -0.35 DPS) [world]; Ring of Power Regulation (273030, -0.41 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -5.26 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -5.27 DPS) [dungeon]; Glimmering Staff (249392, -6.46 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.12 DPS) [dungeon]; Earthen Rod (9381, -2.60 DPS) [dungeon]; Umbral Wand (5216, -2.94 DPS) [dungeon] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Black Cord; legs: Red Mageweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 025003000000000000-00000000000000000-52542001201300051)

Set DPS (verified): 257.3. Weights run: 3.1s. Verify run: 5.9s. 601 eligible items had no known source.

Race sweep: undead 257.3 ± 0.3, troll 221.6 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.540 ± 0.027, crit=0.568 ± 0.013 per rating point (14 rating = 1%, 7.949 per %), hit=0.820 ± 0.035 per rating point (10 rating = 1%, 8.196 per %), spell_haste=not significant (-0.029 ± 0.265), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.159 ± 0.227), mp5=2.064 ± 0.347

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 52.3 spell_power points (13.36 DPS) | yes | Red Mageweave Headband (10033, -0.65 DPS) [crafted]; Chief Architect's Monocle (11839, -1.71 DPS) [dungeon]; Soulcatcher Halo (10630, -2.00 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 22.2 spell_power points (5.66 DPS) | yes | Glowing Eye of Mordresh (10769, -0.73 DPS) [dungeon]; Gemshard Heart (17707, -0.85 DPS) [dungeon]; Mindburst Medallion (11196, -1.57 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 40.7 spell_power points (10.39 DPS) | yes | Kentic Amice (11624, -1.47 DPS) [dungeon]; Mantle of Lost Hope (22234, -1.66 DPS) [dungeon]; Netherpearl Shoulders (254055, -2.66 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.42 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -0.90 DPS) [vendor]; Runecloth Cloak (13860, -1.16 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 47.9 spell_power points (12.24 DPS) | yes | Runecloth Robe (13858, -2.34 DPS, sim-verified) [crafted]; Acumen Robes (17775, -2.77 DPS) [quest]; Stone Guard's Dreadweave Vest (220904, -2.82 DPS) [vendor] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Aristocratic Cuffs (12546, -0.69 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -1.34 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.43 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 35.7 spell_power points (9.10 DPS) | yes | Runecloth Gloves (13863, -2.14 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.25 DPS) [vendor]; Silkweb Gloves (11634, -2.28 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 35.6 spell_power points (9.08 DPS) | yes | Black Waistcord (254087, -0.70 DPS) [crafted]; Deathmage Sash (10771, -1.40 DPS) [dungeon]; Dawnspire Cord (12466, -1.83 DPS, sim-verified) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 44.5 spell_power points (11.35 DPS) | yes | Spellshock Leggings (9484, -1.55 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -1.78 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.78 DPS) [dungeon] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 33.9 spell_power points (8.64 DPS) | yes | Coldstone Slippers (18697, -1.03 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -1.79 DPS, sim-verified) [vendor]; Southsea Mojo Boots (20641, -2.28 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 27.2 spell_power points (6.94 DPS) | yes | Woodseed Hoop (17768, -1.48 DPS) [quest]; Cyclopean Band (11824, -1.73 DPS) [dungeon]; Mindseye Circle (10634, -6.63 DPS, sim-verified) [dungeon] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Woodseed Hoop (17768, -0.08 DPS) [quest]; Cyclopean Band (11824, -0.33 DPS) [dungeon]; Mindseye Circle (10634, -1.08 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.10 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.17 DPS) [quest]; Blessed Prayer Beads (19990, -1.61 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.35 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.42 DPS) [quest]; Blessed Prayer Beads (19990, -0.86 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -0.68 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 22.3 spell_power points (5.70 DPS) | yes | Flask of Forest Mojo (19115, -0.26 DPS) [quest]; Enthralled Sphere (11625, -0.51 DPS) [dungeon]; Cloud Stone (17737, -0.88 DPS) [dungeon] |
| ranged | Woestave (20082) | Blood of Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pyric Caduceus (11748, -2.30 DPS) [dungeon]; Lesser Eternal Wand (249232, -2.59 DPS) [crafted]; Noxious Shooter (17745, -6.66 DPS, sim-verified) [dungeon] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Robes of Insight; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Ban'thok Sash; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Underworld Band; trinket1: Darkspear Voodoo Seal; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Woestave

No-known-source sample (15 of 601, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 025105031030000000-00000000000000000-52542001201300051)

Set DPS (verified): 451.9. Weights run: 3.2s. Verify run: 18.7s. 1520 eligible items had no known source.

Race sweep: undead 451.9 ± 0.4, troll 401.8 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.976 ± 0.025, crit=0.739 ± 0.015 per rating point (14 rating = 1%, 10.340 per %), hit=0.908 ± 0.039 per rating point (10 rating = 1%, 9.083 per %), spell_haste=not significant (0.361 ± 0.288), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.378 ± 0.341), mp5=not significant (1.498 ± 0.419)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 47.0 spell_power points (15.13 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -4.03 DPS, sim-verified) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 29.5 spell_power points (9.51 DPS) | yes | Dark Advisor's Pendant (18691, -0.56 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.71 DPS) [world_drop]; Nacreous Shell Necklace (22403, -1.18 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 40.0 spell_power points (12.89 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Champion's Satin Epaulets (227130, -1.07 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.52 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.89 DPS) [world]; Spritecaster Cape (11623, -3.47 DPS) [dungeon]; Shroud of the Exile (15421, -4.15 DPS, sim-verified) [quest] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 51.7 spell_power points (16.65 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Alanna's Embrace (13314, -4.61 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 32.5 spell_power points (10.45 DPS) | yes | General's Satin Bracers (17619, -3.26 DPS) [pvp]; Netherpearl Cuffs (254067, -3.46 DPS) [crafted]; Manacle Cuffs (11962, -3.48 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 39.4 spell_power points (12.70 DPS) | yes | General's Satin Gloves (17620, -1.21 DPS) [vendor]; General's Satin Grips (231613, -1.21 DPS) [vendor]; Hands of Power (13253, -1.49 DPS) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 30.8 spell_power points (9.93 DPS) | yes | Defiler's Cloth Girdle (20163, -0.21 DPS) [rep]; General's Satin Cinch (17621, -0.88 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.89 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 50.3 spell_power points (16.21 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Legionnaire's Satin Leggings (227131, -1.83 DPS) [vendor]; Skyshroud Leggings (13170, -4.17 DPS, sim-verified) [dungeon] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 33.6 spell_power points (10.83 DPS) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -3.20 DPS) [quest]; Seal of Rivendare (13345, -3.28 DPS) [dungeon]; Naglering (11669, -15.34 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.73 DPS) [quest]; Seal of Rivendare (13345, -0.81 DPS) [dungeon]; Brainlash (6440, -1.43 DPS) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+14.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Smolderweb's Eye (13213, -0.93 DPS) [dungeon]; Eye of the Beast (13968, -6.65 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -0.99 DPS) [dungeon]; Ironbark Staff (20220, -9.49 DPS, sim-verified) [rep] |
| off_hand | Book of the Dead (13353) | Stratholme: Balnazzar [dungeon] | 34.4 spell_power points (11.08 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Therazane's Touch (19315, -1.75 DPS) [rep]; Tome of Shadow Force (19309, -3.47 DPS, sim-verified) [rep] |
| ranged | Oblivion's Touch (18761) | Scholomance: Death Knight Darkreaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.59 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -1.28 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Outrider's Silk Leggings; feet: Dragonrider Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Second Wind; trinket2: Burst of Knowledge; main_hand: Hand of Edward the Odd; off_hand: Book of the Dead; ranged: Oblivion's Touch

No-known-source sample (15 of 1520, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005105031030000000-00000000000000000-50052201211312251)

Set DPS (verified): 784.7. Weights run: 3.5s. Verify run: 20.8s. 1520 eligible items had no known source.

Race sweep: undead 784.7 ± 0.5, troll 776.9 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.271 ± 0.008, crit=0.939 ± 0.017 per rating point (14 rating = 1%, 13.145 per %), hit=1.150 ± 0.049 per rating point (10 rating = 1%, 11.503 per %), spell_haste=1.521 ± 0.266, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 44.5 spell_power points (21.50 DPS) | yes | Warlord's Satin Cowl (17623, -1.89 DPS) [vendor]; Warlord's Satin Crown (231615, -1.89 DPS) [vendor]; Felcloth Hood (14111, -4.15 DPS, sim-verified) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 22.7 spell_power points (10.95 DPS) | yes | Dark Advisor's Pendant (18691, -0.24 DPS) [dungeon]; Emberfury Talisman (12929, -1.70 DPS) [dungeon]; Star of Mystaria (12103, -2.81 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 33.7 spell_power points (16.26 DPS) | yes | Warlord's Satin Epaulets (231611, -1.57 DPS) [vendor]; Felcloth Shoulders (14112, -3.66 DPS, sim-verified) [crafted]; Argent Shoulders (19059, -4.19 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.86 DPS, sim-verified) [world]; Spritecaster Cape (11623, -6.78 DPS) [dungeon]; Shroud of the Exile (15421, -7.23 DPS) [quest] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 43.3 spell_power points (20.89 DPS) | yes | Warlord's Satin Robes (231612, -1.29 DPS) [pvp]; Felcloth Robe (14106, -3.66 DPS, sim-verified) [crafted]; Warlord's Satin Tunic (231632, -7.08 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.2 spell_power points (11.68 DPS) | yes | Netherpearl Cuffs (254067, -4.48 DPS) [crafted]; Runecloth Cuffs (254123, -5.05 DPS) [crafted]; Arena Wristguards (18709, -5.33 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 38.0 spell_power points (18.36 DPS) | yes | General's Satin Gloves (17620, -5.54 DPS) [vendor]; General's Satin Grips (231613, -5.54 DPS) [vendor]; Hands of Power (13253, -6.80 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 28.8 spell_power points (13.90 DPS) | yes | Ban'thok Sash (11662, -1.24 DPS) [dungeon]; Black Waistcord (254087, -3.54 DPS) [crafted]; Frostwolf Cloth Belt (19090, -3.89 DPS) [rep] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 33.3 spell_power points (16.06 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.05 DPS) [rep]; Felcloth Pants (14107, -1.93 DPS) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 23.3 spell_power points (11.23 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Black Sandals (254113, -0.39 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.24 DPS) [dungeon]; Maiden's Circle (13001, -2.29 DPS) [world_drop]; Naglering (11669, -10.84 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.07 DPS) [dungeon]; Eye of Orgrimmar (12545, -2.13 DPS) [quest]; Maiden's Circle (13001, -2.13 DPS) [world_drop] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -3.14 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.29 DPS, sim-verified) [crafted]; Rune of the Guard Captain (19120, -4.81 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.33 DPS) [pvp]; Headmaster's Charge (13937, -14.09 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -17.24 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Woestave (20082, -0.65 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, -0.65 DPS) [dungeon]; Oblivion's Touch (18761, -0.70 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Skyshroud Leggings; feet: Omnicast Boots; finger1: Don Julio's Band; finger2: Channeler's Ring; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1520, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

