# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 42.7. Weights run: 3.5s. Verify run: 3.1s. 198 eligible items had no known source.

Race sweep: gnome 42.7 ± 0.1, high-order-skyborne 41.3 ± 0.0, human 41.3 ± 0.0 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.989 ± 0.011, crit=0.088 ± 0.002 per rating point (14 rating = 1%, 1.234 per %), hit=0.339 ± 0.012 per rating point (10 rating = 1%, 3.386 per %), spell_haste=not significant (0.279 ± 0.116), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.860 ± 0.003, spirit=not significant (0.157 ± 0.066), mp5=0.523 ± 0.125

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 9.5 spell_power points (1.46 DPS) | yes | Shadow Goggles (4373, -0.55 DPS) [crafted]; Pristine Circlet (253949, -0.65 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.36 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.05 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.9 spell_power points (2.14 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.94 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.53 DPS) [crafted]; Slime-encrusted Pads (6461, -1.90 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.05 DPS) [crafted]; Sanguine Cape (14376, -0.16 DPS) [world_drop]; Caretaker's Cape (20428, -0.25 DPS) [rep] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 13.2 spell_power points (2.03 DPS) | yes | Filigreed Pristine Gown (253901, -0.51 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.65 DPS) [crafted]; Mystic's Wrap (14369, -0.96 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.5 spell_power points (1.00 DPS) | yes | Repurposed Hair Band (281256, -0.65 DPS) [quest]; Mystic's Bracelets (14366, -0.70 DPS) [world_drop]; Bright Bracers (3647, -0.75 DPS, sim-verified) [world_drop] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.0 spell_power points (1.38 DPS) | yes | Pristine Gloves (253913, -0.31 DPS) [crafted]; Gnoll Casting Gloves (892, -0.46 DPS) [world]; Serpent Gloves (5970, -0.50 DPS, sim-verified) [dungeon] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 10.0 spell_power points (1.54 DPS) | yes | Keller's Girdle (2911, -0.32 DPS) [world_drop]; Pristine Sash (253925, -0.38 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.46 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 16.0 spell_power points (2.47 DPS) | yes | Filigreed Pristine Leggings (253937, -0.65 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.26 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.39 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.0 spell_power points (1.69 DPS) | yes | Pearly Boots (253897, -0.52 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.77 DPS) [crafted]; Red Woolen Boots (4313, -0.95 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (1.07 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.62 DPS) [world_drop]; Furlfeather Ring (286745, -0.92 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.0 spell_power points (0.93 DPS) | yes | Lavishly Jeweled Ring (1156, -0.37 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.47 DPS) [world_drop]; Furlfeather Ring (286745, -0.78 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.8 spell_power points (4.74 DPS) | yes | Staff of the Blessed Seer (2271, -1.04 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.04 DPS) [dungeon]; Buzzer Blade (2169, -1.35 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 4.9 spell_power points (0.76 DPS) | yes | Tear of Grief (5611, -0.07 DPS) [quest]; Dwarven Tome (279898, -0.28 DPS) [quest]; Seer's Fine Stein (7608, -0.28 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -0.98 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -1.44 DPS) [quest]; Greater Magic Wand (11288, -1.47 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Skycaller

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 89.8. Weights run: 3.9s. Verify run: 7.5s. 326 eligible items had no known source.

Race sweep: gnome 89.8 ± 0.1, human 88.4 ± 0.1, high-order-skyborne 88.0 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.641 ± 0.015, crit=0.108 ± 0.003 per rating point (14 rating = 1%, 1.509 per %), hit=0.420 ± 0.023 per rating point (10 rating = 1%, 4.199 per %), spell_haste=1.336 ± 0.251, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.623 ± 0.002, spirit=not significant (0.471 ± 0.166), mp5=not significant (0.762 ± 0.201)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 13.8 spell_power points (3.73 DPS) | yes | Enchanter's Cowl (4322, -0.38 DPS) [crafted]; Nightsky Cowl (4039, -1.02 DPS) [world_drop]; Silk Headband (7050, -1.30 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.4 spell_power points (1.20 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world]; Seal of the Expedition (271920, -0.34 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (3.99 DPS) | yes | Death Speaker Mantle (6685, -0.46 DPS) [dungeon]; Magician's Mantle (12998, -1.08 DPS) [world_drop]; Invoker's Mantle (215365, -1.23 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Repairman's Cape (9605, -0.07 DPS) [quest]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Glowing Thresher Cape (6901, -0.28 DPS, sim-verified) [dungeon] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 20.5 spell_power points (5.53 DPS) | yes | Green Silk Armor (7065, -0.54 DPS, sim-verified) [crafted]; Mechbuilder's Overalls (9508, -1.29 DPS) [dungeon]; Pearly Gown (253969, -1.62 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.3 spell_power points (3.05 DPS) | yes | Nightsky Wristbands (6407, -1.63 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.66 DPS) [world_drop]; Mindthrust Bracers (1974, -1.93 DPS, sim-verified) [dungeon] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 10.1 spell_power points (2.73 DPS) | yes | Gloves of Meditation (4318, -0.49 DPS) [crafted]; Shilly Mitts (9609, -0.61 DPS) [quest]; Gloves of Old (9395, -1.45 DPS, sim-verified) [world_drop] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 13.3 spell_power points (3.59 DPS) | yes | Highlander's Cloth Girdle (20099, -0.10 DPS) [rep]; Crimson Silk Belt (7055, -0.75 DPS) [crafted]; Invoker's Cord (215366, -0.83 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 16.2 spell_power points (4.38 DPS) | yes | Pearly Leggings (253995, -0.18 DPS) [crafted]; Pristine Leggings (253987, -0.65 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.99 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 14.5 spell_power points (3.91 DPS) | yes | Gilded Slippers (254001, -0.30 DPS) [crafted]; Acidic Walkers (9454, -0.66 DPS) [dungeon]; Glinteye Slippers (273024, -1.26 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 9.7 spell_power points (2.61 DPS) | yes | Lorekeeper's Ring (19525, -0.31 DPS) [rep]; Ring of Power Regulation (273030, -0.37 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (2.49 DPS) | yes | Ring of Power Regulation (273030, -0.25 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.72 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -0.73 DPS) [crafted]; Rod of the Sleepwalker (1155, -1.27 DPS) [dungeon]; Glimmering Staff (249392, -3.29 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.8 spell_power points (2.93 DPS) | yes | Orb of Souls (249395, -1.09 DPS) [crafted]; Arcane Infused Rod (279838, -1.58 DPS, sim-verified) [quest]; Satyr's Rod (15962, -1.64 DPS) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.56 DPS) [crafted]; Thunderwood (13062, -0.64 DPS, sim-verified) [world_drop]; Wand of Mana Concentration (273043, -0.69 DPS) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 166.6. Weights run: 4.0s. Verify run: 3.8s. 456 eligible items had no known source.

Race sweep: gnome 166.6 ± 0.3, human 166.2 ± 0.3, high-order-skyborne 164.3 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.626 ± 0.022, crit=0.261 ± 0.006 per rating point (14 rating = 1%, 3.656 per %), hit=0.634 ± 0.035 per rating point (10 rating = 1%, 6.342 per %), spell_haste=2.712 ± 0.403, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.626 ± 0.002, spirit=not significant (0.260 ± 0.249), mp5=not significant (0.429 ± 0.271)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 22.6 spell_power points (7.12 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.52 DPS) [crafted]; Papal Fez (9431, -0.52 DPS) [dungeon]; Corpseshroud (10574, -1.18 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.0 spell_power points (3.78 DPS) | yes | Triune Amulet (7722, -0.69 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.95 DPS) [quest]; Necklace of Calisea (1714, -1.83 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.1 spell_power points (4.76 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.59 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.18 DPS) [crafted]; Expeditionary's Cape (271922, -0.38 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.3 spell_power points (8.57 DPS) | yes | Robe of Power (7054, -1.16 DPS) [crafted]; Robe of the Magi (1716, -1.22 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.97 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.3 spell_power points (3.23 DPS) | yes | Condor Bracers (15864, -1.03 DPS) [quest]; Arcane Runed Bracers (4744, -1.26 DPS, sim-verified) [quest]; Windchaser Cuffs (14429, -1.46 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.3 spell_power points (7.02 DPS) | yes | Red Mageweave Gloves (10018, -2.10 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.30 DPS) [crafted]; Golden Handwraps (254029, -2.89 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.5 spell_power points (5.19 DPS) | yes | Deathmage Sash (10771, -0.04 DPS) [dungeon]; Golden Cord (254045, -0.37 DPS) [crafted]; Gilded Cord (254037, -0.61 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.5 spell_power points (6.76 DPS) | yes | Crimson Silk Pantaloons (7062, -1.74 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.01 DPS) [dungeon]; Pearly Leggings (253995, -2.22 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.6 spell_power points (4.26 DPS) | yes | Gilded Slippers (254001, -0.36 DPS) [crafted]; Acidic Walkers (9454, -0.79 DPS) [dungeon]; Spidersilk Boots (4320, -1.27 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (4.13 DPS) | yes | Snake Hoop (6750, -1.26 DPS) [quest]; Runemender's Seal (273031, -1.51 DPS) [world]; Ring of Forlorn Spirits (2043, -1.61 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 10.3 spell_power points (3.23 DPS) | yes | Snake Hoop (6750, -0.36 DPS) [quest]; Runemender's Seal (273031, -0.62 DPS) [world]; Ring of Forlorn Spirits (2043, -0.72 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -10.12 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -10.57 DPS) [quest]; Hypnotic Blade (7714, -10.95 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.35 DPS) [dungeon]; Starfaller (13063, -0.67 DPS) [world_drop]; Twisted Nether Wand (249144, -0.68 DPS) [crafted] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 141005213100311531-30000000000000000-033000000000000000)

Set DPS (verified): 280.7. Weights run: 4.2s. Verify run: 11.1s. 600 eligible items had no known source.

Race sweep: human 280.8 ± 0.4, gnome 280.7 ± 0.4, high-order-skyborne 276.9 ± 0.4 (dps); human beats the guide race gnome by 0.1 (0.04%) ± 0.6, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.962 ± 0.036, crit=0.272 ± 0.008 per rating point (14 rating = 1%, 3.814 per %), hit=0.629 ± 0.047 per rating point (10 rating = 1%, 6.286 per %), spell_haste=not significant (1.578 ± 0.541), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.742 ± 0.003, spirit=not significant (0.830 ± 0.419), mp5=not significant (0.753 ± 0.397)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 40.6 spell_power points (9.95 DPS) | yes | Red Mageweave Headband (10033, -0.57 DPS) [crafted]; Gemburst Circlet (10751, -1.30 DPS) [quest]; Chief Architect's Monocle (11839, -2.14 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.8 spell_power points (5.34 DPS) | yes | Horizon Choker (13085, -1.22 DPS) [world_drop]; Mindburst Medallion (11196, -1.43 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.73 DPS) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 31.5 spell_power points (7.72 DPS) | yes | Rotgrip Mantle (17732, -0.29 DPS) [dungeon]; Nethershine Shoulders (254057, -1.52 DPS) [crafted]; Mantle of Lost Hope (22234, -2.18 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -1.75 DPS) [rep]; Runecloth Cloak (13860, -1.97 DPS) [crafted]; Featherskin Cape (10843, -2.28 DPS) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 45.5 spell_power points (11.16 DPS) | yes | Dreamweave Vest (10021, -1.78 DPS) [crafted]; Robes of Insight (940, -2.80 DPS) [world_drop]; Robe of Power (7054, -3.27 DPS) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 20.5 spell_power points (5.03 DPS) | yes | Aristocratic Cuffs (12546, -0.51 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.65 DPS) [crafted]; Bloodband Bracers (11469, -1.07 DPS) [quest] |
| hands | Golden Gloves (254103) | Tailoring [crafted] | 30.4 spell_power points (7.44 DPS) | yes | Runecloth Gloves (13863, -0.55 DPS) [crafted]; Dreamweave Gloves (10019, -0.66 DPS) [crafted]; Gilded Gloves (254095, -1.00 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.9 spell_power points (6.84 DPS) | yes | Dawnspire Cord (12466, -0.89 DPS) [dungeon]; Satyrmane Sash (17755, -1.05 DPS) [dungeon]; Ghostweave Cord (254073, -1.17 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kilt of the Atal'ai Prophet (10807, +0.00 DPS) [dungeon]; Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, -0.79 DPS) [world_drop] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 5 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mistwalker Boots (10629, +0.00 DPS) [dungeon]; Vinerot Sandals (17748, +0.00 DPS) [dungeon]; Golden Sandals (254115, +0.00 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 24.8 spell_power points (6.07 DPS) | yes | Eye of Adaegus (5266, -2.21 DPS) [world_drop]; Woodseed Hoop (17768, -2.24 DPS) [quest]; Lorekeeper's Ring (19523, -2.39 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.1 spell_power points (4.67 DPS) | yes | Eye of Adaegus (5266, -0.81 DPS) [world_drop]; Woodseed Hoop (17768, -0.84 DPS) [quest]; Lorekeeper's Ring (19523, -0.99 DPS) [rep] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.16 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 17.7 spell_power points (4.34 DPS) | yes | Enthralled Sphere (11625, -0.13 DPS) [dungeon]; Celestial Orb (7515, -0.44 DPS) [quest]; Orb of Lorica (11262, -0.51 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cairnstone Sliver (9654, +0.00 DPS) [quest]; Nether Force Wand (11263, +0.00 DPS) [quest]; Charged Lightning Rod (11860, -0.17 DPS) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Nethershine Cuffs; hands: Golden Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Lesser Eternal Wand

No-known-source sample (15 of 600, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 141005213100311531-30000000000000000-053500030000000000)

Set DPS (verified): 414.4. Weights run: 4.3s. Verify run: 24.4s. 1459 eligible items had no known source.

Race sweep: gnome 414.4 ± 0.6, human 413.5 ± 0.6, high-order-skyborne 409.1 ± 0.6 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.707 ± 0.032, crit=0.457 ± 0.012 per rating point (14 rating = 1%, 6.404 per %), hit=0.903 ± 0.048 per rating point (10 rating = 1%, 9.032 per %), spell_haste=not significant (2.297 ± 0.583), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.729 ± 0.003, spirit=not significant (0.567 ± 0.419), mp5=not significant (0.656 ± 0.405)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 49.4 spell_power points (13.15 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.56 DPS) [pvp]; Crimson Felt Hat (18727, -1.77 DPS) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 24.4 spell_power points (6.50 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.01 DPS) [world_drop]; Nacreous Shell Necklace (22403, -0.20 DPS) [dungeon]; Star of Mystaria (12103, -1.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 32.6 spell_power points (8.68 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.30 DPS) [pvp]; Burial Shawl (18681, -0.34 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.09 DPS) [world]; Spritecaster Cape (11623, -2.41 DPS) [dungeon]; Shroud of the Exile (15421, -3.14 DPS, sim-verified) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.02 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.36 DPS) [pvp]; The Postmaster's Tunic (13388, -8.36 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.6 spell_power points (8.42 DPS) | yes | Marshal's Silk Bracers (16438, -3.33 DPS) [pvp]; Runecloth Cuffs (254123, -3.61 DPS) [crafted]; Manacle Cuffs (11962, -3.73 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 33.9 spell_power points (9.02 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 36.0 spell_power points (9.57 DPS) | yes | Thuzadin Sash (18740, -2.46 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.65 DPS) [crafted]; Ban'thok Sash (11662, -5.09 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.9 spell_power points (13.81 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -2.56 DPS) [pvp]; Spiritshroud Leggings (12965, -3.33 DPS) [dungeon] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 29.3 spell_power points (7.80 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.22 DPS) [dungeon]; Knight-Lieutenant's Silk Walkers (227112, -0.33 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.74 DPS) [vendor]; Songstone of Ironforge (12543, -0.91 DPS) [quest]; Naglering (11669, -12.82 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Songstone of Ironforge (12543, -0.14 DPS) [quest]; The Postmaster's Seal (13392, -0.41 DPS) [dungeon]; Blessed Band of Light (272407, -2.16 DPS, sim-verified) [vendor] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -4.61 DPS, sim-verified) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (7.72 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -1.09 DPS) [pvp]; Tome of Arcane Domination (19308, -1.18 DPS) [rep] |
| ranged | Oblivion's Touch (18761) | Scholomance: Death Knight Darkreaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.33 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -0.63 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Maiden's Circle; finger2: Channeler's Ring; trinket1: Eye of the Beast; trinket2: Burst of Knowledge; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Oblivion's Touch

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 801.6. Weights run: 5.0s. Verify run: 25.1s. 1459 eligible items had no known source.

Race sweep: gnome 801.6 ± 0.8, human 797.3 ± 0.8, high-order-skyborne 791.0 ± 0.8 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.980 ± 0.043, crit=0.880 ± 0.035 per rating point (14 rating = 1%, 12.319 per %), hit=1.666 ± 0.103 per rating point (10 rating = 1%, 16.663 per %), spell_haste=3.550 ± 0.771, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.574 ± 0.002, spirit=not significant (0.583 ± 0.298), mp5=not significant (0.595 ± 0.355)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 57.2 spell_power points (30.04 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.47 DPS) [pvp]; Crimson Felt Hat (18727, -7.11 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 33.5 spell_power points (17.58 DPS) | yes | Star of Mystaria (12103, -0.47 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.04 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -2.12 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 42.1 spell_power points (22.07 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.85 DPS) [pvp]; Shroud of the Nathrezim (18720, -1.26 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -7.67 DPS) [quest]; Crystalline Threaded Cape (20697, -8.70 DPS) [world]; Cloak of Flames (3475, -11.79 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -5.87 DPS) [pvp]; Magister's Robes (16688, -18.96 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 33.9 spell_power points (17.80 DPS) | yes | Marshal's Silk Bracers (16438, -4.97 DPS) [pvp]; Manacle Cuffs (11962, -6.42 DPS) [quest]; Runecloth Cuffs (254123, -6.89 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 43.4 spell_power points (22.79 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Silk Gloves (16440, -0.92 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -0.92 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 42.5 spell_power points (22.31 DPS) | yes | Ban'thok Sash (11662, -2.13 DPS) [dungeon]; Clutch of Andros (13956, -3.80 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -5.41 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 66.4 spell_power points (34.86 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -7.11 DPS) [pvp]; Magister's Leggings (16687, -9.18 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 38.3 spell_power points (20.10 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -2.43 DPS) [dungeon]; Sergeant Major's Dreadweave Boots (220891, -2.53 DPS) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -3.18 DPS) [rep]; Maiden's Circle (13001, -3.51 DPS) [world_drop]; Naglering (11669, -17.95 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.23 DPS) [rep]; Maiden's Circle (13001, -1.56 DPS) [world_drop]; Songstone of Ironforge (12543, -3.40 DPS) [quest] |
| trinket1 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.87 DPS) [crafted]; Cannonball Runner (13382, -4.37 DPS) [dungeon]; Ankh of Life (1713, -5.07 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.58 DPS) [crafted]; Eye of the Beast (13968, -2.03 DPS, sim-verified) [quest]; Cannonball Runner (13382, -5.07 DPS) [dungeon] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -12.27 DPS, sim-verified) [world_drop]; Headmaster's Charge (13937, -13.47 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.02 DPS) [dungeon]; Brilliant Wand (249385, -0.05 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -2.56 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Smolderweb's Eye; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 41.9. Weights run: 3.5s. Verify run: 3.7s. 192 eligible items had no known source.

Race sweep: undead 41.9 ± 0.0, orc 40.3 ± 0.0, troll 39.9 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.989 ± 0.011, crit=0.088 ± 0.002 per rating point (14 rating = 1%, 1.234 per %), hit=0.339 ± 0.012 per rating point (10 rating = 1%, 3.386 per %), spell_haste=not significant (0.279 ± 0.116), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.860 ± 0.003, spirit=not significant (0.157 ± 0.066), mp5=0.523 ± 0.125

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 9.5 spell_power points (1.46 DPS) | yes | Shadow Goggles (4373, -0.55 DPS) [crafted]; Pristine Circlet (253949, -0.61 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.36 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.5 spell_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.9 spell_power points (2.14 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.84 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.53 DPS) [crafted]; Slime-encrusted Pads (6461, -1.90 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.05 DPS) [crafted]; Sanguine Cape (14376, -0.16 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.25 DPS) [rep] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 13.2 spell_power points (2.03 DPS) | yes | Filigreed Pristine Gown (253901, -0.48 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.65 DPS) [crafted]; Mystic's Wrap (14369, -0.96 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.5 spell_power points (1.00 DPS) | yes | Tabitha's Cuffs (251486, -0.54 DPS) [quest]; Owlbeard Bracers (16981, -0.54 DPS) [quest]; Bright Bracers (3647, -0.55 DPS, sim-verified) [world_drop] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.0 spell_power points (1.38 DPS) | yes | Pristine Gloves (253913, -0.31 DPS) [crafted]; Serpent Gloves (5970, -0.32 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.46 DPS) [world] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 10.0 spell_power points (1.54 DPS) | yes | Keller's Girdle (2911, -0.32 DPS) [world_drop]; Pristine Sash (253925, -0.36 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.46 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 16.0 spell_power points (2.47 DPS) | yes | Filigreed Pristine Leggings (253937, -0.61 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.26 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.39 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.0 spell_power points (1.69 DPS) | yes | Pearly Boots (253897, -0.48 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.77 DPS) [crafted]; Red Woolen Boots (4313, -0.95 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.0 spell_power points (0.93 DPS) | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Volcanic Rock Ring (12053, -0.47 DPS) [world_drop]; Furlfeather Ring (286745, -0.78 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.9 spell_power points (0.91 DPS) | yes | Loop of Sacrifice (281673, -0.15 DPS) [quest]; Volcanic Rock Ring (12053, -0.46 DPS) [world_drop]; Furlfeather Ring (286745, -0.76 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | 32.9 spell_power points (5.07 DPS) | yes | Evocator's Blade (2567, -0.33 DPS) [dungeon]; Staff of the Blessed Seer (2271, -0.52 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.82 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -1.02 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -1.56 DPS) [quest]; Greater Magic Wand (11288, -1.60 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Crescent Staff; ranged: Skycaller

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 89.5. Weights run: 3.9s. Verify run: 7.5s. 318 eligible items had no known source.

Race sweep: undead 89.5 ± 0.1, orc 88.2 ± 0.1, troll 87.1 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.641 ± 0.015, crit=0.108 ± 0.003 per rating point (14 rating = 1%, 1.509 per %), hit=0.420 ± 0.023 per rating point (10 rating = 1%, 4.199 per %), spell_haste=1.336 ± 0.251, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.623 ± 0.002, spirit=not significant (0.471 ± 0.166), mp5=not significant (0.762 ± 0.201)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.32 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.64 DPS) [world_drop]; Silk Headband (7050, -0.92 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.4 spell_power points (1.20 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world]; Darkspear Warding Pendant (272075, -0.34 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (3.99 DPS) | yes | Death Speaker Mantle (6685, -0.46 DPS) [dungeon]; Magician's Mantle (12998, -1.08 DPS) [world_drop]; Invoker's Mantle (215365, -1.23 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.03 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.09 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 20.5 spell_power points (5.53 DPS) | yes | Green Silk Armor (7065, -0.84 DPS) [crafted]; Mechbuilder's Overalls (9508, -1.29 DPS) [dungeon]; Pearly Gown (253969, -1.62 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.3 spell_power points (3.05 DPS) | yes | Mindthrust Bracers (1974, -1.07 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.63 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.66 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 10.1 spell_power points (2.73 DPS) | yes | Jutebraid Gloves (10654, -0.24 DPS) [quest]; Gloves of Old (9395, -0.34 DPS) [world_drop]; Gloves of Meditation (4318, -0.49 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Belt of Arugal (6392, -0.36 DPS, sim-verified) [dungeon]; Warsong Sash (16975, -0.52 DPS) [quest]; Crimson Silk Belt (7055, -0.66 DPS) [crafted] |
| legs | Pearly Leggings (253995) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.47 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.50 DPS, sim-verified) [dungeon]; Filigreed Pearly Leggings (253945, -0.81 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 14.5 spell_power points (3.91 DPS) | yes | Gilded Slippers (254001, -0.30 DPS) [crafted]; Acidic Walkers (9454, -0.66 DPS) [dungeon]; Glinteye Slippers (273024, -1.26 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 9.7 spell_power points (2.61 DPS) | yes | Advisor's Ring (19521, -0.31 DPS) [rep]; Ring of Power Regulation (273030, -0.37 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (2.49 DPS) | yes | Advisor's Ring (19521, -0.18 DPS) [rep]; Ring of Power Regulation (273030, -0.25 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -0.73 DPS) [crafted]; Rod of the Sleepwalker (1155, -1.27 DPS) [dungeon]; Glimmering Staff (249392, -2.37 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.8 spell_power points (2.93 DPS) | yes | Witch's Finger (16887, -1.54 DPS) [quest]; Satyr's Rod (15962, -1.64 DPS) [world_drop]; Orb of Souls (249395, -1.76 DPS, sim-verified) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.52 DPS) [world_drop]; Greater Mystic Wand (217287, -0.65 DPS) [crafted]; Wand of Mana Concentration (273043, -0.73 DPS) [world] |

**New at 30:** head: Enchanter's Cowl; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Pearly Leggings; feet: Golden Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 165.8. Weights run: 4.0s. Verify run: 3.6s. 448 eligible items had no known source.

Race sweep: undead 165.8 ± 0.3, orc 164.5 ± 0.3, troll 163.1 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.714 ± 0.022, crit=0.169 ± 0.004 per rating point (14 rating = 1%, 2.360 per %), hit=0.409 ± 0.028 per rating point (10 rating = 1%, 4.094 per %), spell_haste=not significant (1.016 ± 0.359), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.606 ± 0.003, spirit=not significant (0.475 ± 0.278), mp5=not significant (0.494 ± 0.289)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 24.4 spell_power points (5.85 DPS) | yes | Miner's Hat of the Deep (9429, -0.07 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.82 DPS) [crafted]; Corpseshroud (10574, -0.92 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (3.05 DPS) | yes | Triune Amulet (7722, -0.33 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.89 DPS) [quest]; Necklace of Calisea (1714, -1.05 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.3 spell_power points (3.90 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.21 DPS) [dungeon]; Mantle of Doan (7712, -0.29 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.02 DPS) [vendor]; Long Silken Cloak (4326, -0.07 DPS) [crafted]; Guardian Cloak (5965, -0.07 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 31.1 spell_power points (7.45 DPS) | yes | Robe of Power (7054, -1.13 DPS) [crafted]; Robe of the Magi (1716, -1.19 DPS, sim-verified) [world_drop]; Beguiler Robes (7728, -2.33 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.5 spell_power points (2.51 DPS) | yes | Radiant Silver Bracers (4545, -0.18 DPS) [quest]; Condor Bracers (15864, -0.84 DPS) [quest]; Windchaser Cuffs (14429, -0.97 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.2 spell_power points (5.80 DPS) | yes | Black Mageweave Gloves (10003, -2.20 DPS) [crafted]; Red Mageweave Gloves (10018, -2.22 DPS, sim-verified) [crafted]; Golden Handwraps (254029, -2.56 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.7 spell_power points (4.25 DPS) | yes | Golden Cord (254045, -0.16 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.20 DPS) [rep]; Gilded Cord (254037, -0.28 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.6 spell_power points (5.41 DPS) | yes | Crimson Silk Pantaloons (7062, -1.27 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.72 DPS, sim-verified) [dungeon] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 14.8 spell_power points (3.54 DPS) | yes | Gilded Slippers (254001, -0.21 DPS) [crafted]; Acidic Walkers (9454, -0.52 DPS) [dungeon]; Furen's Boots (13100, -0.52 DPS) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.6 spell_power points (3.25 DPS) | yes | Runemender's Seal (273031, -0.82 DPS) [world]; Snake Hoop (6750, -0.96 DPS) [quest]; Ring of Power Regulation (273030, -1.46 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 10.5 spell_power points (2.51 DPS) | yes | Runemender's Seal (273031, -0.08 DPS) [world]; Snake Hoop (6750, -0.22 DPS) [quest]; Ring of Power Regulation (273030, -0.72 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -8.02 DPS) [quest]; Hypnotic Blade (7714, -8.16 DPS) [dungeon]; Illusionary Rod (7713, -9.62 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.27 DPS) [dungeon]; Jaina's Firestarter (13064, -0.34 DPS) [world_drop]; Starfaller (13063, -0.51 DPS) [world_drop] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 141005213100311531-30000000000000000-033000000000000000)

Set DPS (verified): 277.6. Weights run: 4.2s. Verify run: 10.6s. 592 eligible items had no known source.

Race sweep: undead 277.6 ± 0.4, orc 274.6 ± 0.4, troll 272.2 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.962 ± 0.036, crit=0.272 ± 0.008 per rating point (14 rating = 1%, 3.814 per %), hit=0.629 ± 0.047 per rating point (10 rating = 1%, 6.286 per %), spell_haste=not significant (1.578 ± 0.541), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.742 ± 0.003, spirit=not significant (0.830 ± 0.419), mp5=not significant (0.753 ± 0.397)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 40.6 spell_power points (9.95 DPS) | yes | Red Mageweave Headband (10033, -0.57 DPS) [crafted]; Gemburst Circlet (10751, -1.30 DPS) [quest]; Chief Architect's Monocle (11839, -2.14 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.8 spell_power points (5.34 DPS) | yes | Mindburst Medallion (11196, -1.43 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.73 DPS) [dungeon]; Horizon Choker (13085, -2.14 DPS, sim-verified) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 31.5 spell_power points (7.72 DPS) | yes | Rotgrip Mantle (17732, -0.29 DPS) [dungeon]; Nethershine Shoulders (254057, -1.52 DPS) [crafted]; Mantle of Lost Hope (22234, -2.18 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272076, -1.75 DPS) [vendor]; Deep Woodlands Cloak (19121, -1.78 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.97 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 45.5 spell_power points (11.16 DPS) | yes | Dreamweave Vest (10021, -1.78 DPS) [crafted]; Robes of Insight (940, -2.80 DPS) [world_drop]; Robe of Power (7054, -3.27 DPS) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 20.5 spell_power points (5.03 DPS) | yes | Aristocratic Cuffs (12546, -0.51 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.65 DPS) [crafted]; Bloodband Bracers (11469, -1.07 DPS) [quest] |
| hands | Golden Gloves (254103) | Tailoring [crafted] | 30.4 spell_power points (7.44 DPS) | yes | Dreamweave Gloves (10019, -0.66 DPS) [crafted]; Gilded Gloves (254095, -1.00 DPS) [crafted]; Runecloth Gloves (13863, -1.80 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.9 spell_power points (6.84 DPS) | yes | Satyrmane Sash (17755, -1.05 DPS) [dungeon]; Ghostweave Cord (254073, -1.17 DPS) [crafted]; Dawnspire Cord (12466, -3.68 DPS, sim-verified) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 35.3 spell_power points (8.64 DPS) | yes | Senior Designer's Pantaloons (11841, -0.39 DPS) [dungeon]; Spellshock Leggings (9484, -0.65 DPS) [dungeon]; Dalewind Trousers (13008, -1.43 DPS) [world_drop] |
| feet | Golden Sandals (254115) | Tailoring [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Vinerot Sandals (17748, -0.13 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -0.13 DPS) [vendor]; Mistwalker Boots (10629, -2.21 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 24.8 spell_power points (6.07 DPS) | yes | Eye of Adaegus (5266, -2.21 DPS) [world_drop]; Woodseed Hoop (17768, -2.24 DPS) [quest]; Advisor's Ring (19519, -2.39 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.1 spell_power points (4.67 DPS) | yes | Woodseed Hoop (17768, -0.84 DPS) [quest]; Advisor's Ring (19519, -0.99 DPS) [rep]; Eye of Adaegus (5266, -2.40 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.31 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.39 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -7.41 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 17.7 spell_power points (4.34 DPS) | yes | Flask of Forest Mojo (19115, -0.25 DPS) [quest]; Celestial Orb (7515, -0.44 DPS) [quest]; Enthralled Sphere (11625, -2.15 DPS, sim-verified) [dungeon] |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Charged Lightning Rod (11860, -0.17 DPS) [quest]; Jaina's Firestarter (13064, -0.21 DPS) [world_drop] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Nethershine Cuffs; hands: Golden Gloves; waist: Ban'thok Sash; legs: Kilt of the Atal'ai Prophet; feet: Golden Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 141005213100311531-30000000000000000-053500030000000000)

Set DPS (verified): 410.4. Weights run: 4.3s. Verify run: 23.4s. 1453 eligible items had no known source.

Race sweep: undead 410.4 ± 0.5, orc 407.5 ± 0.5, troll 404.9 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.707 ± 0.032, crit=0.457 ± 0.012 per rating point (14 rating = 1%, 6.404 per %), hit=0.903 ± 0.048 per rating point (10 rating = 1%, 9.032 per %), spell_haste=not significant (2.297 ± 0.583), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.729 ± 0.003, spirit=not significant (0.567 ± 0.419), mp5=not significant (0.656 ± 0.405)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 49.4 spell_power points (13.15 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.56 DPS) [pvp]; Crimson Felt Hat (18727, -1.77 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Nacreous Shell Necklace (22403, -0.19 DPS) [dungeon]; Star of Mystaria (12103, -0.99 DPS) [dungeon] |
| shoulder | Darkspear Shoulderpads (272103) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Champion's Silk Mantle (227104, +0.00 DPS) [pvp]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.04 DPS) [quest]; Crystalline Threaded Cape (20697, -2.09 DPS) [world]; Spritecaster Cape (11623, -2.41 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.02 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.36 DPS) [pvp]; The Postmaster's Tunic (13388, -4.55 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.6 spell_power points (8.42 DPS) | yes | General's Silk Cuffs (16538, -3.33 DPS) [pvp]; Runecloth Cuffs (254123, -3.61 DPS) [crafted]; Manacle Cuffs (11962, -3.73 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 33.9 spell_power points (9.02 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; General's Silk Handwraps (231600, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 36.0 spell_power points (9.57 DPS) | yes | Ban'thok Sash (11662, -2.09 DPS) [dungeon]; Thuzadin Sash (18740, -2.46 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.65 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 47.1 spell_power points (12.54 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -1.29 DPS) [pvp]; Spiritshroud Leggings (12965, -2.05 DPS) [dungeon] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 29.3 spell_power points (7.80 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.22 DPS) [dungeon]; Blood Guard's Silk Walkers (227110, -0.33 DPS) [pvp] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, +0.00 DPS) [vendor]; Eye of Orgrimmar (12545, -0.14 DPS) [quest]; The Postmaster's Seal (13392, -0.41 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.74 DPS) [vendor]; Eye of Orgrimmar (12545, -0.91 DPS) [quest]; The Postmaster's Seal (13392, -1.18 DPS) [dungeon] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -1.58 DPS) [dungeon] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (7.72 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -1.09 DPS) [pvp]; Tome of Arcane Domination (19308, -1.18 DPS) [rep] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.17 DPS) [dungeon]; Brilliant Wand (249385, -0.50 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -0.80 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderpads; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Dragonrider Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Eye of the Beast; trinket2: Burst of Knowledge; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 789.3. Weights run: 5.0s. Verify run: 23.7s. 1453 eligible items had no known source.

Race sweep: undead 789.3 ± 0.8, orc 788.3 ± 0.8, troll 777.5 ± 0.8 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.980 ± 0.043, crit=0.880 ± 0.035 per rating point (14 rating = 1%, 12.319 per %), hit=1.666 ± 0.103 per rating point (10 rating = 1%, 16.663 per %), spell_haste=3.550 ± 0.771, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.574 ± 0.002, spirit=not significant (0.583 ± 0.298), mp5=not significant (0.595 ± 0.355)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 57.2 spell_power points (30.04 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.47 DPS) [pvp]; Crimson Felt Hat (18727, -6.10 DPS) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 33.5 spell_power points (17.58 DPS) | yes | Star of Mystaria (12103, -0.47 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.04 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -2.12 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 42.1 spell_power points (22.07 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -0.85 DPS) [pvp]; Shroud of the Nathrezim (18720, -1.26 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -7.67 DPS) [quest]; Crystalline Threaded Cape (20697, -8.70 DPS) [world]; Spritecaster Cape (11623, -8.99 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -5.87 DPS) [pvp]; Magister's Robes (16688, -10.48 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 33.9 spell_power points (17.80 DPS) | yes | General's Silk Cuffs (16538, -4.97 DPS) [pvp]; Manacle Cuffs (11962, -6.42 DPS) [quest]; Runecloth Cuffs (254123, -6.89 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 43.4 spell_power points (22.79 DPS) | yes | General's Silk Handguards (16540, -0.92 DPS) [vendor]; General's Silk Gauntlets (231599, -0.92 DPS) [vendor]; General's Silk Handwraps (231600, -0.92 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 42.5 spell_power points (22.31 DPS) | yes | Ban'thok Sash (11662, -2.13 DPS) [dungeon]; Clutch of Andros (13956, -3.80 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -5.41 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 52.4 spell_power points (27.52 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -1.84 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 38.3 spell_power points (20.10 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -2.43 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -2.53 DPS) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -3.18 DPS) [rep]; Maiden's Circle (13001, -3.51 DPS) [world_drop]; Eye of Orgrimmar (12545, -5.35 DPS) [quest] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.23 DPS) [rep]; Maiden's Circle (13001, -1.56 DPS) [world_drop]; Eye of Orgrimmar (12545, -3.40 DPS) [quest] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -0.70 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.58 DPS) [crafted]; Rune of the Guard Captain (19120, -3.32 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -8.88 DPS) [world_drop]; Headmaster's Charge (13937, -13.47 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.50 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

