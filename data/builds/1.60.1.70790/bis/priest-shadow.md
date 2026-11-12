# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 020000000000000000-00000000000000000-40500000000000000)

Set DPS (verified): 38.4. Weights run: 2.6s. Verify run: 3.5s. 203 eligible items had no known source.

Race sweep: human 38.4 ± 0.0, dwarf 38.4 ± 0.0, gnome 38.3 ± 0.0, night-elf 38.2 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.695 ± 0.018, crit=0.147 ± 0.003 per rating point (14 rating = 1%, 2.058 per %), hit=0.457 ± 0.015 per rating point (10 rating = 1%, 4.573 per %), spell_haste=0.859 ± 0.206, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.133 ± 0.214), mp5=1.444 ± 0.274

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (1.00 DPS) | yes | Pristine Circlet (253949, -0.45 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.69 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.3 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.35 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.66 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Pearl-clasped Cloak (5542, -0.14 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 17.5 spell_power points (1.60 DPS) | yes | Filigreed Shadow Gown (253907, -0.35 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.82 DPS) [crafted]; Gray Woolen Robe (2585, -0.91 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.8 spell_power points (0.71 DPS) | yes | Mystic's Bracelets (14366, -0.59 DPS) [world_drop]; Repurposed Hair Band (281256, -0.59 DPS) [quest]; Bright Bracers (3647, -1.08 DPS, sim-verified) [world_drop] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 9.1 spell_power points (0.83 DPS) | yes | Serpent Gloves (5970, -0.18 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.27 DPS) [crafted]; Gnoll Casting Gloves (892, -0.28 DPS) [world] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 9.8 spell_power points (0.89 DPS) | yes | Pristine Sash (253925, -0.27 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.34 DPS) [crafted]; Keller's Girdle (2911, -0.39 DPS) [world_drop] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 15.2 spell_power points (1.39 DPS) | yes | Filigreed Pristine Leggings (253937, -0.45 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.75 DPS) [dungeon]; Colorful Kilt (10048, -0.93 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.8 spell_power points (0.89 DPS) | yes | Shadow Boots (253895, -0.15 DPS) [crafted]; Pristine Boots (253889, -0.43 DPS) [crafted]; Red Woolen Boots (4313, -0.53 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.9 spell_power points (0.72 DPS) | yes | Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop]; Furlfeather Ring (286745, -0.63 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.4 spell_power points (0.58 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Furlfeather Ring (286745, -0.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.13 DPS) [dungeon]; Emberstone Staff (5201, -0.49 DPS) [dungeon]; Witching Stave (1484, -0.51 DPS, sim-verified) [dungeon] |
| off_hand | Tear of Grief (5611) (or Mathystra Channeling Orb (281291)) | For Love Eternal [quest] | 3.5 spell_power points (0.32 DPS) | yes | Mathystra Channeling Orb (281291, +0.00 DPS) [quest]; Dwarven Tome (279898, -0.04 DPS) [quest]; Spellbinder Orb (15926, -0.13 DPS) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -1.07 DPS, sim-verified) [world_drop]; Deepblaze (279896, -3.00 DPS) [quest]; Sizzle Stick (8071, -3.52 DPS) [quest] |

**New at 20:** head: Shadow Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Tear of Grief; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 020000000000000000-00000000000000000-50542001200000000)

Set DPS (verified): 62.7. Weights run: 2.6s. Verify run: 3.6s. 325 eligible items had no known source.

Race sweep: human 62.7 ± 0.1, gnome 62.6 ± 0.1, night-elf 61.8 ± 0.1, dwarf 61.8 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.497 ± 0.011, crit=0.151 ± 0.003 per rating point (14 rating = 1%, 2.113 per %), hit=0.450 ± 0.013 per rating point (10 rating = 1%, 4.504 per %), spell_haste=not significant (0.049 ± 0.060), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=-2.096 ± 0.130, mp5=-1.301 ± 0.175

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 14.0 spell_power points (1.27 DPS) | yes | Holy Shroud (2721, -0.27 DPS) [world_drop]; Shadow Circlet (253955, -0.27 DPS) [crafted]; Shadow Hood (4323, -0.46 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.0 spell_power points (0.27 DPS) | yes | Crystal Starfire Medallion (5003, -0.09 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.09 DPS) [world_drop]; Seal of the Expedition (271920, -0.51 DPS, sim-verified) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.5 spell_power points (1.22 DPS) | yes | Invoker's Mantle (215365, -0.36 DPS) [crafted]; Magician's Mantle (12998, -0.36 DPS) [world_drop]; Death Speaker Mantle (6685, -0.64 DPS, sim-verified) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Prelacy Cape (7004, -0.09 DPS) [quest] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 16.5 spell_power points (1.50 DPS) | yes | Green Silk Armor (7065, -0.09 DPS) [crafted]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Shadow Gown (253967, -0.24 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.82 DPS) | yes | Nightsky Wristbands (6407, -0.55 DPS) [world_drop]; Stonecloth Bindings (14416, -0.59 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.24 DPS, sim-verified) [world_drop] |
| hands | Hands of Darkness (7047) | Tailoring [crafted] | sim-verified (62.7 DPS) | yes | Shadow Gloves (253919, -0.05 DPS) [crafted]; Shilly Mitts (9609, -0.18 DPS) [quest]; Ephemeral Grips (273034, -0.27 DPS, sim-verified) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.5 spell_power points (1.14 DPS) | yes | Invoker's Cord (215366, -0.27 DPS) [crafted]; Crimson Silk Belt (7055, -0.27 DPS) [crafted]; Belt of Arugal (6392, -0.38 DPS, sim-verified) [dungeon] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 17.5 spell_power points (1.59 DPS) | yes | Filigreed Shadow Leggings (253943, -0.32 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.50 DPS) [dungeon]; Necromancer Leggings (2277, -0.61 DPS, sim-verified) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 16.5 spell_power points (1.50 DPS) | yes | Spidersilk Boots (4320, -0.68 DPS) [crafted]; Acidic Walkers (9454, -0.68 DPS) [dungeon]; Gilded Slippers (254001, -0.80 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.86 DPS) | yes | Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Minor Channeling Ring (1449, -0.32 DPS) [quest]; Runemender's Seal (273031, -0.32 DPS) [world] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.64 DPS) | yes | Minor Channeling Ring (1449, -0.09 DPS) [quest]; Runemender's Seal (273031, -0.09 DPS) [world]; Ring of Power Regulation (273030, -0.50 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Shade (2549, -0.27 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.36 DPS) [world]; Glimmering Staff (249392, -1.28 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.0 spell_power points (0.91 DPS) | yes | Eye of Paleth (2943, -0.54 DPS) [quest]; Orb of Souls (249395, -0.54 DPS) [crafted]; Orb of Mistmantle (13031, -1.07 DPS, sim-verified) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, -1.53 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -1.53 DPS) [dungeon]; Gravestone Scepter (7001, -1.72 DPS) [quest] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; wrist: Spidertank Oilrag; hands: Hands of Darkness; waist: Highlander's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 325, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish

### Band 40 (human, 020000000000000000-00000000000000000-52542001201300040)

Set DPS (verified): 91.8. Weights run: 2.6s. Verify run: 4.0s. 461 eligible items had no known source.

Race sweep: gnome 94.7 ± 0.1, human 91.8 ± 0.1, night-elf 91.7 ± 0.1, dwarf 91.6 ± 0.1 (dps); gnome beats the guide race human by 2.9 (3.21%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.079 ± 0.024, crit=0.240 ± 0.006 per rating point (14 rating = 1%, 3.358 per %), hit=0.671 ± 0.029 per rating point (10 rating = 1%, 6.708 per %), spell_haste=not significant (0.285 ± 0.218), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.267 ± 0.269), mp5=not significant (0.533 ± 0.345)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Papal Fez (9431, -0.16 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.25 DPS, sim-verified) [dungeon]; Augural Shroud (2620, -0.79 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.6 spell_power points (2.15 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.95 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.05 DPS) [dungeon]; Seal of the Expedition (271919, -1.11 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.0 spell_power points (2.89 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Bloodmage Mantle (7684, -0.32 DPS) [dungeon]; Death Speaker Mantle (6685, -0.43 DPS) [dungeon] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.07 DPS) [crafted]; Blackforge Cape (6424, -0.45 DPS) [dungeon]; Long Silken Cloak (4326, -0.79 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.10 DPS) [crafted]; Robe of the Magi (1716, -0.46 DPS, sim-verified) [world_drop]; Green Silk Armor (7065, -0.64 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.6 spell_power points (1.46 DPS) | yes | Windchaser Cuffs (14429, -0.12 DPS) [world_drop]; Arcane Runed Bracers (4744, -0.22 DPS) [quest]; Mistscape Bracers (4045, -0.27 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Gloves (10018, -0.07 DPS) [crafted]; Black Handwraps (254027, -0.10 DPS) [crafted]; Shadoweave Gloves (10023, -0.26 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.2 spell_power points (3.19 DPS) | yes | Black Cord (254043, -0.08 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.67 DPS) [rep]; Gilded Cord (254037, -0.90 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.0 spell_power points (3.71 DPS) | yes | Necromancer Leggings (2277, -0.70 DPS) [world_drop]; Shadow Leggings (253993, -0.74 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.91 DPS, sim-verified) [crafted] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 20.6 spell_power points (2.83 DPS) | yes | Gilded Slippers (254001, -0.84 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.95 DPS) [dungeon]; Spidersilk Boots (4320, -1.27 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.4 spell_power points (2.67 DPS) | yes | Lorekeeper's Ring (19524, -1.21 DPS) [rep]; Runemender's Seal (273031, -1.36 DPS) [world]; Philanthropist's Ring (281635, -2.98 DPS, sim-verified) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lorekeeper's Ring (19524, -0.41 DPS) [rep]; Philanthropist's Ring (281635, -0.54 DPS, sim-verified) [quest]; Runemender's Seal (273031, -0.56 DPS) [world] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -4.97 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -5.36 DPS) [dungeon]; Glimmering Staff (249392, -6.29 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.57 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -2.36 DPS) [dungeon]; Umbral Wand (5216, -2.73 DPS) [dungeon] |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Expeditionary's Cape; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; finger1: Underworld Band; finger2: Snake Hoop; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 025003000000000000-00000000000000000-52542001201300051)

Set DPS (verified): 218.8. Weights run: 2.6s. Verify run: 7.5s. 606 eligible items had no known source.

Race sweep: gnome 225.1 ± 0.3, human 218.8 ± 0.3, dwarf 216.8 ± 0.3, night-elf 216.7 ± 0.3 (dps); gnome beats the guide race human by 6.3 (2.89%) ± 0.4, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.516 ± 0.028, crit=0.471 ± 0.012 per rating point (14 rating = 1%, 6.588 per %), hit=0.716 ± 0.035 per rating point (10 rating = 1%, 7.157 per %), spell_haste=not significant (0.405 ± 0.323), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.211 ± 0.217), mp5=2.043 ± 0.471

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 51.8 spell_power points (12.66 DPS) | yes | Red Mageweave Headband (10033, -0.61 DPS) [crafted]; Soulcatcher Halo (10630, -1.94 DPS) [dungeon]; Chief Architect's Monocle (11839, -2.03 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 21.2 spell_power points (5.18 DPS) | yes | Glowing Eye of Mordresh (10769, -0.51 DPS) [dungeon]; Gemshard Heart (17707, -0.63 DPS) [dungeon]; Mindburst Medallion (11196, -1.50 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Kentic Amice (11624, -0.03 DPS) [dungeon]; Rotgrip Mantle (17732, -0.54 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.00 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -0.46 DPS) [rep]; Runecloth Cloak (13860, -0.48 DPS) [crafted]; Big Voodoo Cloak (8216, -1.09 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 47.4 spell_power points (11.57 DPS) | yes | Runecloth Robe (13858, -2.15 DPS, sim-verified) [crafted]; Acumen Robes (17775, -2.60 DPS) [quest]; Knight's Dreadweave Vest (220886, -2.96 DPS) [vendor] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Aristocratic Cuffs (12546, -0.66 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -1.21 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.32 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 34.2 spell_power points (8.34 DPS) | yes | Silkweb Gloves (11634, -1.06 DPS) [dungeon]; Sergeant Major's Dreadweave Gloves (220890, -1.84 DPS) [vendor]; Red Mageweave Gloves (10018, -1.95 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Waistcord (254087, -0.41 DPS) [crafted]; Deathmage Sash (10771, -1.12 DPS) [dungeon]; Dawnspire Cord (12466, -1.24 DPS, sim-verified) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 44.2 spell_power points (10.79 DPS) | yes | Spellshock Leggings (9484, -1.47 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -1.81 DPS) [vendor]; Red Mageweave Pants (10009, -2.93 DPS) [crafted] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 33.6 spell_power points (8.22 DPS) | yes | Coldstone Slippers (18697, -1.04 DPS) [dungeon]; Sergeant Major's Dreadweave Boots (220891, -1.18 DPS) [vendor]; Southsea Mojo Boots (20641, -2.19 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 26.2 spell_power points (6.39 DPS) | yes | Woodseed Hoop (17768, -1.22 DPS) [quest]; Lorekeeper's Ring (19523, -1.46 DPS) [rep]; Mindseye Circle (10634, -6.18 DPS, sim-verified) [dungeon] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Woodseed Hoop (17768, -0.10 DPS) [quest]; Lorekeeper's Ring (19523, -0.34 DPS) [rep]; Mindseye Circle (10634, -0.53 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.11 DPS) [world_drop]; Blessed Prayer Beads (19990, -0.60 DPS) [quest] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.03 DPS) [world_drop]; Blessed Prayer Beads (19990, -1.52 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -1.73 DPS) [world_drop]; Radiant Staff (249453, -3.97 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Woestave (20082) | Blood of Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pyric Caduceus (11748, -2.28 DPS) [dungeon]; Lesser Eternal Wand (249232, -2.63 DPS) [crafted]; Noxious Shooter (17745, -5.83 DPS, sim-verified) [dungeon] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Robes of Insight; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Ban'thok Sash; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Underworld Band; trinket1: Frozen Heart of the Mountain; trinket2: Enriched Seal; ranged: Woestave

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 025105031030000000-00000000000000000-52542001201300051)

Set DPS (verified): 418.7. Weights run: 2.6s. Verify run: 22.1s. 1523 eligible items had no known source.

Race sweep: gnome 420.4 ± 0.5, human 418.7 ± 0.5, night-elf 399.9 ± 0.3, dwarf 399.6 ± 0.3 (dps); gnome beats the guide race human by 1.6 (0.39%) ± 0.7, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.473 ± 0.026, crit=0.728 ± 0.015 per rating point (14 rating = 1%, 10.190 per %), hit=1.040 ± 0.042 per rating point (10 rating = 1%, 10.405 per %), spell_haste=not significant (-0.056 ± 0.279), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=1.743 ± 0.382, mp5=2.202 ± 0.420

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Star of Mystaria (12103, -0.93 DPS) [dungeon]; Dark Advisor's Pendant (18691, -2.26 DPS) [dungeon] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Nathrezim (18720, +0.00 DPS) [dungeon]; Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Theramore Shoulderpads (271953, +0.00 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, +0.00 DPS) [quest]; Expeditionary's Cape (271907, -1.30 DPS) [rep]; Spritecaster Cape (11623, -1.53 DPS) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Demonskin Gloves (13181, +0.00 DPS) [dungeon]; Hands of Power (13253, -0.57 DPS) [dungeon]; Hands of the Exalted Herald (12554, -0.61 DPS) [dungeon] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devout Belt (16696, +0.00 DPS) [dungeon]; Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Penitent's Cinch (272394, +0.00 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Haunting Specter Leggings (11929, +0.00 DPS) [dungeon]; Magister's Leggings (16687, +0.00 DPS) [dungeon]; Marshal's Satin Pants (17603, +0.00 DPS) [vendor] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -2.08 DPS) [dungeon]; Brainlash (6440, -3.30 DPS) [dungeon]; Blessed Band of Light (272407, -3.33 DPS) [vendor] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -0.14 DPS) [dungeon]; Brainlash (6440, -1.36 DPS) [dungeon]; Blessed Band of Light (272407, -1.39 DPS) [vendor] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.06 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -0.54 DPS) [dungeon] |
| off_hand | Tome of Shadow Force (19309) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Book of the Dead (13353, +0.00 DPS) [dungeon]; Penelope's Rose (15805, +0.00 DPS) [quest]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Wand of Eternal Light (22254, -1.54 DPS) [dungeon] |

**New at 60:** head: Magister's Crown; neck: Lady Maye's Pendant; shoulder: Magister's Mantle; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Magister's Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Second Wind; trinket2: Burst of Knowledge; main_hand: Hand of Edward the Odd; off_hand: Tome of Shadow Force; ranged: Brilliant Wand

No-known-source sample (15 of 1523, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005105031030000000-00000000000000000-50052201211312251)

Set DPS (verified): 798.2. Weights run: 3.0s. Verify run: 17.4s. 1523 eligible items had no known source.

Race sweep: human 798.2 ± 0.5, night-elf 797.8 ± 0.5, dwarf 796.7 ± 0.5, gnome 795.3 ± 0.5 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.266 ± 0.008, crit=0.955 ± 0.018 per rating point (14 rating = 1%, 13.374 per %), hit=1.236 ± 0.052 per rating point (10 rating = 1%, 12.360 per %), spell_haste=1.427 ± 0.264, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 45.4 spell_power points (21.73 DPS) | yes | Field Marshal's Headdress (17602, -2.35 DPS) [vendor]; Field Marshal's Satin Crown (231616, -2.35 DPS) [vendor]; Felcloth Hood (14111, -4.76 DPS, sim-verified) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.5 spell_power points (11.25 DPS) | yes | Dark Advisor's Pendant (18691, -0.65 DPS) [dungeon]; Emberfury Talisman (12929, -1.97 DPS) [dungeon]; Star of Mystaria (12103, -2.66 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 33.8 spell_power points (16.21 DPS) | yes | Field Marshal's Satin Epaulets (231621, -1.68 DPS) [vendor]; Argent Shoulders (19059, -4.23 DPS) [crafted]; Felcloth Shoulders (14112, -4.25 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.42 DPS, sim-verified) [world]; Spritecaster Cape (11623, -7.14 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -7.28 DPS) [dungeon] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 43.2 spell_power points (20.70 DPS) | yes | Field Marshal's Satin Vestments (17605, -1.31 DPS) [vendor]; Field Marshal's Satin Robe (231618, -1.31 DPS) [vendor]; Felcloth Robe (14106, -3.70 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.1 spell_power points (11.56 DPS) | yes | Netherpearl Cuffs (254067, -4.44 DPS) [crafted]; Runecloth Cuffs (254123, -5.01 DPS) [crafted]; Arena Wristguards (18709, -5.15 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 38.4 spell_power points (18.40 DPS) | yes | Marshal's Satin Gloves (17608, -5.72 DPS) [vendor]; Marshal's Satin Grips (231617, -5.72 DPS) [vendor]; Hands of Power (13253, -6.51 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 29.0 spell_power points (13.88 DPS) | yes | Ban'thok Sash (11662, -3.28 DPS, sim-verified) [dungeon]; Black Waistcord (254087, -3.63 DPS) [crafted]; Stormpike Cloth Girdle (19094, -3.98 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 61.4 spell_power points (29.42 DPS) | yes | Marshal's Satin Pants (17603, -11.41 DPS) [vendor]; Marshal's Satin Leggings (231619, -11.41 DPS) [vendor]; Skyshroud Leggings (13170, -13.58 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 23.2 spell_power points (11.11 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -3.00 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.62 DPS) [vendor]; Maiden's Circle (13001, -2.81 DPS) [world_drop]; Naglering (11669, -12.15 DPS, sim-verified) [dungeon] |
| finger2 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -2.16 DPS) [quest]; Maiden's Circle (13001, -2.16 DPS) [world_drop]; Blessed Band of Light (272407, -3.26 DPS, sim-verified) [vendor] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -2.70 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.32 DPS, sim-verified) [crafted]; Serenity Field (272439, -5.27 DPS) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.33 DPS) [pvp]; Headmaster's Charge (13937, -13.55 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -16.71 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Woestave (20082, -0.64 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, -0.64 DPS) [dungeon]; Oblivion's Touch (18761, -0.70 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Don Julio's Band; finger2: Don Mauricio's Band of Domination; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1523, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 020000000000000000-00000000000000000-40500000000000000)

Set DPS (verified): 41.1. Weights run: 2.6s. Verify run: 3.3s. 196 eligible items had no known source.

Race sweep: undead 41.1 ± 0.0, troll 37.3 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.695 ± 0.018, crit=0.147 ± 0.003 per rating point (14 rating = 1%, 2.058 per %), hit=0.457 ± 0.015 per rating point (10 rating = 1%, 4.573 per %), spell_haste=0.859 ± 0.206, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.133 ± 0.214), mp5=1.444 ± 0.274

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (1.00 DPS) | yes | Pristine Circlet (253949, -0.46 DPS) [crafted]; Shadow Goggles (4373, -0.69 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.3 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.32 DPS) [crafted]; Slime-encrusted Pads (6461, -0.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.66 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, +0.00 DPS) [crafted]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Battle Healer's Cloak (20427, -0.09 DPS) [rep] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 17.5 spell_power points (1.60 DPS) | yes | Filigreed Shadow Gown (253907, -0.46 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.82 DPS) [crafted]; Gray Woolen Robe (2585, -0.91 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.8 spell_power points (0.71 DPS) | yes | Tabitha's Cuffs (251486, -0.44 DPS) [quest]; Bright Bracers (3647, -0.46 DPS) [world_drop]; Owlbeard Bracers (16981, -0.50 DPS) [quest] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 9.1 spell_power points (0.83 DPS) | yes | Serpent Gloves (5970, -0.19 DPS) [dungeon]; Pristine Gloves (253913, -0.27 DPS) [crafted]; Gnoll Casting Gloves (892, -0.28 DPS) [world] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 9.8 spell_power points (0.89 DPS) | yes | Pristine Sash (253925, -0.27 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.34 DPS) [crafted]; Keller's Girdle (2911, -0.39 DPS) [world_drop] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 15.2 spell_power points (1.39 DPS) | yes | Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted]; Silk-threaded Trousers (1929, -0.75 DPS) [dungeon]; Colorful Kilt (10048, -0.93 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.8 spell_power points (0.89 DPS) | yes | Shadow Boots (253895, -0.15 DPS) [crafted]; Pristine Boots (253889, -0.43 DPS) [crafted]; Red Woolen Boots (4313, -0.53 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.9 spell_power points (0.72 DPS) | yes | Loop of Sacrifice (281673, -0.40 DPS) [quest]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop]; Furlfeather Ring (286745, -0.63 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.2 spell_power points (0.38 DPS) | yes | Loop of Sacrifice (281673, -0.06 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Furlfeather Ring (286745, -0.29 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Witching Stave (1484, +0.00 DPS) [dungeon]; Evocator's Blade (2567, +0.00 DPS) [dungeon]; Crescent Staff (6505, +0.00 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.56 DPS) [world_drop]; Sizzle Stick (8071, -3.32 DPS) [quest]; Greater Magic Wand (11288, -3.43 DPS) [crafted] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of the Blessed Seer; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (undead, 020000000000000000-00000000000000000-50542001200000000)

Set DPS (verified): 68.7. Weights run: 2.6s. Verify run: 3.7s. 318 eligible items had no known source.

Race sweep: undead 68.7 ± 0.1, troll 61.9 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.497 ± 0.011, crit=0.151 ± 0.003 per rating point (14 rating = 1%, 2.113 per %), hit=0.450 ± 0.013 per rating point (10 rating = 1%, 4.504 per %), spell_haste=not significant (0.049 ± 0.060), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=-2.096 ± 0.130, mp5=-1.301 ± 0.175

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 14.0 spell_power points (1.27 DPS) | yes | Shadow Hood (4323, -0.14 DPS) [crafted]; Holy Shroud (2721, -0.27 DPS) [world_drop]; Shadow Circlet (253955, -0.27 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (68.7 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS) [world_drop]; Darkspear Warding Pendant (272075, +0.00 DPS) [vendor]; Archmagister's Faceted Pendant (273049, +0.00 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.5 spell_power points (1.22 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Invoker's Mantle (215365, -0.36 DPS) [crafted]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (68.7 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Battle Healer's Cloak (19529, -0.09 DPS) [rep] |
| chest | Shadow Gown (253967) | Tailoring [crafted] | sim-verified (68.7 DPS) | yes | Necrology Robes (2292, +0.00 DPS) [dungeon]; Green Silk Armor (7065, -0.09 DPS) [crafted]; Beguiler Robes (7728, -0.14 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.82 DPS) | yes | Glowing Magical Bracelets (13106, -0.46 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.55 DPS) [quest]; Nightsky Wristbands (6407, -0.55 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 12.5 spell_power points (1.13 DPS) | yes | Hands of Darkness (7047, -0.32 DPS) [crafted]; Shadow Gloves (253919, -0.36 DPS) [crafted]; Jutebraid Gloves (10654, -0.36 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.5 spell_power points (1.14 DPS) | yes | Warsong Sash (16975, -0.14 DPS) [quest]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Invoker's Cord (215366, -0.27 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 17.5 spell_power points (1.59 DPS) | yes | Necromancer Leggings (2277, -0.18 DPS) [world_drop]; Filigreed Shadow Leggings (253943, -0.32 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.50 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 16.5 spell_power points (1.50 DPS) | yes | Gilded Slippers (254001, -0.55 DPS) [crafted]; Spidersilk Boots (4320, -0.68 DPS) [crafted]; Acidic Walkers (9454, -0.68 DPS) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (68.7 DPS) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Runemender's Seal (273031, -0.00 DPS) [world]; Solid Gold Band (279845, -0.18 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.64 DPS) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Runemender's Seal (273031, -0.09 DPS) [world]; Solid Gold Band (279845, -0.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, -0.05 DPS) [crafted]; Staff of the Shade (2549, -0.27 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.36 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.0 spell_power points (0.91 DPS) | yes | Orb of Mistmantle (13031, -0.36 DPS) [world_drop]; Orb of Souls (249395, -0.54 DPS) [crafted]; Witch's Finger (16887, -0.55 DPS) [quest] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, -0.80 DPS) [world]; Gyromatic Icemaker (9489, -1.37 DPS) [dungeon]; Gravestone Scepter (7001, -1.47 DPS) [quest] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Shadow Gown; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Ring of Power Regulation; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (undead, 020000000000000000-00000000000000000-52542001201300040)

Set DPS (verified): 104.6. Weights run: 2.6s. Verify run: 3.5s. 454 eligible items had no known source.

Race sweep: undead 104.6 ± 0.1, troll 90.7 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.079 ± 0.024, crit=0.240 ± 0.006 per rating point (14 rating = 1%, 3.358 per %), hit=0.671 ± 0.029 per rating point (10 rating = 1%, 6.708 per %), spell_haste=not significant (0.285 ± 0.218), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.267 ± 0.269), mp5=not significant (0.533 ± 0.345)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 30.4 spell_power points (4.17 DPS) | yes | Papal Fez (9431, -0.55 DPS) [dungeon]; Corpseshroud (10574, -1.09 DPS, sim-verified) [dungeon]; Augural Shroud (2620, -1.18 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.6 spell_power points (2.15 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.81 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.05 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.11 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.0 spell_power points (2.89 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Bloodmage Mantle (7684, -0.32 DPS) [dungeon]; Death Speaker Mantle (6685, -0.43 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.07 DPS) [crafted]; Blackforge Cape (6424, -0.45 DPS) [dungeon]; Long Silken Cloak (4326, -0.53 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.10 DPS) [crafted]; Robe of the Magi (1716, -0.56 DPS, sim-verified) [world_drop]; Green Silk Armor (7065, -0.64 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Windchaser Cuffs (14429, -0.12 DPS) [world_drop]; Mistscape Bracers (4045, -0.27 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.73 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Gloves (10018, -0.07 DPS) [crafted]; Black Handwraps (254027, -0.10 DPS) [crafted]; Shadoweave Gloves (10023, -0.61 DPS, sim-verified) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Cloth Girdle (20166, -0.59 DPS) [rep]; Deathmage Sash (10771, -0.68 DPS, sim-verified) [dungeon]; Gilded Cord (254037, -0.82 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.0 spell_power points (3.71 DPS) | yes | Necromancer Leggings (2277, -0.70 DPS) [world_drop]; Shadow Leggings (253993, -0.74 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.93 DPS, sim-verified) [crafted] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 20.6 spell_power points (2.83 DPS) | yes | Gilded Slippers (254001, -0.94 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.95 DPS) [dungeon]; Spidersilk Boots (4320, -1.27 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.4 spell_power points (2.67 DPS) | yes | Snake Hoop (6750, -0.80 DPS) [quest]; Advisor's Ring (19520, -1.21 DPS) [rep]; Runemender's Seal (273031, -1.36 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.12 DPS) | yes | Snake Hoop (6750, -0.25 DPS) [quest]; Advisor's Ring (19520, -0.66 DPS) [rep]; Runemender's Seal (273031, -0.81 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -5.36 DPS) [dungeon]; Illusionary Rod (7713, -5.80 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -6.29 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.07 DPS) [dungeon]; Earthen Rod (9381, -2.17 DPS) [dungeon]; Twisted Nether Wand (249144, -2.46 DPS) [crafted] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Black Cord; legs: Red Mageweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 025003000000000000-00000000000000000-52542001201300051)

Set DPS (verified): 251.2. Weights run: 2.6s. Verify run: 7.0s. 599 eligible items had no known source.

Race sweep: undead 251.2 ± 0.4, troll 215.4 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.516 ± 0.028, crit=0.471 ± 0.012 per rating point (14 rating = 1%, 6.588 per %), hit=0.716 ± 0.035 per rating point (10 rating = 1%, 7.157 per %), spell_haste=not significant (0.405 ± 0.323), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.211 ± 0.217), mp5=2.043 ± 0.471

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 51.8 spell_power points (12.66 DPS) | yes | Red Mageweave Headband (10033, -0.61 DPS) [crafted]; Soulcatcher Halo (10630, -1.94 DPS) [dungeon]; Chief Architect's Monocle (11839, -2.03 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 21.2 spell_power points (5.18 DPS) | yes | Glowing Eye of Mordresh (10769, -0.51 DPS) [dungeon]; Gemshard Heart (17707, -0.63 DPS) [dungeon]; Mindburst Medallion (11196, -1.50 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 40.3 spell_power points (9.84 DPS) | yes | Mantle of Lost Hope (22234, -1.57 DPS) [dungeon]; Kentic Amice (11624, -1.61 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.58 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.62 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -1.08 DPS) [vendor]; Runecloth Cloak (13860, -1.10 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 47.4 spell_power points (11.57 DPS) | yes | Runecloth Robe (13858, -1.51 DPS, sim-verified) [crafted]; Acumen Robes (17775, -2.60 DPS) [quest]; Stone Guard's Dreadweave Vest (220904, -2.96 DPS) [vendor] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodband Bracers (11469, -1.21 DPS) [quest]; Aristocratic Cuffs (12546, -1.23 DPS, sim-verified) [dungeon]; Shizzle's Nozzle Wiper (11917, -1.32 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 34.2 spell_power points (8.34 DPS) | yes | Silkweb Gloves (11634, -1.74 DPS, sim-verified) [dungeon]; First Sergeant's Dreadweave Gloves (220908, -1.84 DPS) [vendor]; Red Mageweave Gloves (10018, -1.95 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Waistcord (254087, -0.41 DPS) [crafted]; Deathmage Sash (10771, -1.12 DPS) [dungeon]; Dawnspire Cord (12466, -1.59 DPS, sim-verified) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 44.2 spell_power points (10.79 DPS) | yes | Spellshock Leggings (9484, -1.47 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -1.81 DPS) [vendor]; Red Mageweave Pants (10009, -2.93 DPS) [crafted] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 33.6 spell_power points (8.22 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -1.18 DPS) [vendor]; Coldstone Slippers (18697, -1.72 DPS, sim-verified) [dungeon]; Southsea Mojo Boots (20641, -2.19 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 26.2 spell_power points (6.39 DPS) | yes | Woodseed Hoop (17768, -1.22 DPS) [quest]; Advisor's Ring (19519, -1.46 DPS) [rep]; Mindseye Circle (10634, -5.81 DPS, sim-verified) [dungeon] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Woodseed Hoop (17768, -0.10 DPS) [quest]; Advisor's Ring (19519, -0.34 DPS) [rep]; Mindseye Circle (10634, -1.38 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.11 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.35 DPS) [quest]; Blessed Prayer Beads (19990, -0.60 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.03 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.27 DPS) [quest]; Blessed Prayer Beads (19990, -1.52 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Radiant Staff (249453, +0.00 DPS) [crafted]; Staff of Jordan (873, -2.35 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 22.1 spell_power points (5.40 DPS) | yes | Flask of Forest Mojo (19115, -0.24 DPS) [quest]; Enthralled Sphere (11625, -0.61 DPS) [dungeon]; Prophetic Cane (6803, -0.98 DPS) [quest] |
| ranged | Woestave (20082) | Blood of Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -1.85 DPS) [crafted]; Pyric Caduceus (11748, -2.45 DPS) [dungeon]; Noxious Shooter (17745, -4.39 DPS, sim-verified) [dungeon] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Robes of Insight; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Ban'thok Sash; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Underworld Band; trinket1: Frozen Heart of the Mountain; trinket2: Darkspear Voodoo Seal; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Woestave

No-known-source sample (15 of 599, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 025105031030000000-00000000000000000-52542001201300051)

Set DPS (verified): 445.2. Weights run: 2.6s. Verify run: 19.6s. 1518 eligible items had no known source.

Race sweep: undead 445.2 ± 0.5, troll 387.6 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.364 ± 0.035, crit=0.714 ± 0.016 per rating point (14 rating = 1%, 9.991 per %), hit=1.256 ± 0.045 per rating point (10 rating = 1%, 12.563 per %), spell_haste=not significant (0.163 ± 0.227), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.049 ± 0.490), mp5=not significant (2.016 ± 0.533)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 36.9 spell_power points (6.95 DPS) | yes | Nacreous Shell Necklace (22403, -0.84 DPS) [dungeon]; Dark Advisor's Pendant (18691, -1.13 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -1.20 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Champion's Satin Epaulets (227130, -0.05 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -2.24 DPS) [dungeon]; Royal Tribunal Cloak (13376, -2.38 DPS) [dungeon]; Shroud of the Exile (15421, -2.60 DPS) [quest] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 56.4 spell_power points (10.61 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Legionnaire's Satin Robe (227135, -0.53 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 32.9 spell_power points (6.20 DPS) | yes | Manacle Cuffs (11962, -1.03 DPS) [quest]; General's Satin Bracers (17619, -1.32 DPS) [pvp]; Funeral Cuffs (12626, -1.47 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, +0.00 DPS) [vendor]; General's Satin Gloves (17620, -0.36 DPS) [vendor]; General's Satin Grips (231613, -0.36 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Clutch of Andros (13956, +0.00 DPS) [dungeon]; Magister's Belt (16685, -0.48 DPS) [dungeon]; Dustfeather Sash (12589, -1.13 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 53.9 spell_power points (10.15 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Legionnaire's Satin Leggings (227131, -0.07 DPS) [vendor]; General's Satin Legguards (231634, -0.99 DPS) [vendor] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; Kayser's Boots of Precision (22231, +0.00 DPS) [dungeon]; General's Satin Treads (231610, +0.00 DPS) [vendor] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Seal of Rivendare (13345, +0.00 DPS) [dungeon]; Channeler's Ring (272406, +0.00 DPS) [vendor] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Seal of Rivendare (13345, +0.00 DPS) [dungeon]; Channeler's Ring (272406, +0.00 DPS) [vendor] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Smolderweb's Eye (13213, -0.67 DPS) [dungeon]; Shard of the Splithooves (10659, -0.76 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trindlehaven Staff (13161, +0.00 DPS) [dungeon]; Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | Tome of Shadow Force (19309) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Book of the Dead (13353, +0.00 DPS) [dungeon]; High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Wand of Eternal Light (22254, -1.00 DPS) [dungeon] |

**New at 60:** head: The Postmaster's Band; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Outrider's Silk Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Maiden's Circle; trinket1: Second Wind; trinket2: Burst of Knowledge; main_hand: Hand of Edward the Odd; off_hand: Tome of Shadow Force; ranged: Brilliant Wand

No-known-source sample (15 of 1518, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005105031030000000-00000000000000000-50052201211312251)

Set DPS (verified): 785.6. Weights run: 3.0s. Verify run: 16.3s. 1518 eligible items had no known source.

Race sweep: undead 785.6 ± 0.5, troll 777.3 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.266 ± 0.008, crit=0.955 ± 0.018 per rating point (14 rating = 1%, 13.374 per %), hit=1.236 ± 0.052 per rating point (10 rating = 1%, 12.360 per %), spell_haste=1.427 ± 0.264, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 45.4 spell_power points (21.73 DPS) | yes | Warlord's Satin Cowl (17623, -2.35 DPS) [vendor]; Warlord's Satin Crown (231615, -2.35 DPS) [vendor]; Felcloth Hood (14111, -5.53 DPS, sim-verified) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.5 spell_power points (11.25 DPS) | yes | Dark Advisor's Pendant (18691, -0.65 DPS) [dungeon]; Emberfury Talisman (12929, -1.97 DPS) [dungeon]; Star of Mystaria (12103, -2.66 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 33.8 spell_power points (16.21 DPS) | yes | Warlord's Satin Epaulets (231611, -1.68 DPS) [vendor]; Felcloth Shoulders (14112, -3.51 DPS, sim-verified) [crafted]; Argent Shoulders (19059, -4.23 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.39 DPS, sim-verified) [world]; Spritecaster Cape (11623, -7.14 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -7.28 DPS) [dungeon] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 43.2 spell_power points (20.70 DPS) | yes | Warlord's Satin Robes (231612, -1.31 DPS) [pvp]; Felcloth Robe (14106, -3.10 DPS, sim-verified) [crafted]; Warlord's Satin Tunic (231632, -7.06 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.1 spell_power points (11.56 DPS) | yes | Netherpearl Cuffs (254067, -4.44 DPS) [crafted]; Runecloth Cuffs (254123, -5.01 DPS) [crafted]; Arena Wristguards (18709, -5.15 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 38.4 spell_power points (18.40 DPS) | yes | General's Satin Gloves (17620, -5.72 DPS) [vendor]; General's Satin Grips (231613, -5.72 DPS) [vendor]; Hands of Power (13253, -6.28 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 29.0 spell_power points (13.88 DPS) | yes | Ban'thok Sash (11662, -3.03 DPS, sim-verified) [dungeon]; Black Waistcord (254087, -3.63 DPS) [crafted]; Frostwolf Cloth Belt (19090, -3.98 DPS) [rep] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 33.2 spell_power points (15.90 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.06 DPS) [rep]; Felcloth Pants (14107, -1.92 DPS) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 23.2 spell_power points (11.11 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; First Sergeant's Dreadweave Boots (220909, -0.21 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.62 DPS) [vendor]; Maiden's Circle (13001, -2.81 DPS) [world_drop]; Naglering (11669, -13.63 DPS, sim-verified) [dungeon] |
| finger2 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.16 DPS) [quest]; Maiden's Circle (13001, -2.16 DPS) [world_drop]; Blessed Band of Light (272407, -3.02 DPS, sim-verified) [vendor] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+13.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -2.70 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.29 DPS) [crafted]; Serenity Field (272439, -5.33 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.33 DPS) [pvp]; Headmaster's Charge (13937, -13.89 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -16.71 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Woestave (20082, -0.64 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, -0.64 DPS) [dungeon]; Oblivion's Touch (18761, -0.70 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Skyshroud Leggings; feet: Omnicast Boots; finger1: Don Julio's Band; finger2: Don Mauricio's Band of Domination; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1518, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

