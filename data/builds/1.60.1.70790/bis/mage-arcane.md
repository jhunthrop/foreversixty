# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 47.9. Weights run: 2.4s. Verify run: 2.6s. 198 eligible items had no known source.

Race sweep: gnome 47.9 ± 0.1, human 46.4 ± 0.1, high-order-skyborne 46.3 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.767 ± 0.013, crit=0.078 ± 0.002 per rating point (14 rating = 1%, 1.094 per %), hit=0.312 ± 0.012 per rating point (10 rating = 1%, 3.118 per %), spell_haste=not significant (0.156 ± 0.046), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.956 ± 0.003, spirit=not significant (0.036 ± 0.052), mp5=0.572 ± 0.088

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 10.5 spell_power points (2.14 DPS) | yes | Pristine Circlet (253949, -0.96 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.31 DPS) [crafted]; Flying Tiger Goggles (4368, -2.11 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.1 spell_power points (0.01 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.9 spell_power points (2.42 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.73 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.61 DPS) [crafted]; Slime-encrusted Pads (6461, -2.07 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.03 DPS) [crafted]; Caretaker's Cape (20428, -0.25 DPS) [rep]; Sanguine Cape (14376, -0.25 DPS) [world_drop] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 12.5 spell_power points (2.55 DPS) | yes | Filigreed Pristine Gown (253901, -0.77 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.96 DPS) [crafted]; Mystic's Wrap (14369, -1.46 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 5.6 spell_power points (1.13 DPS) | yes | Bright Bracers (3647, -0.47 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.80 DPS) [quest]; Mystic's Bracelets (14366, -0.82 DPS) [world_drop] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.0 spell_power points (1.83 DPS) | yes | Serpent Gloves (5970, -0.33 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.55 DPS) [crafted]; Gnoll Casting Gloves (892, -0.61 DPS) [world] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 9.8 spell_power points (1.98 DPS) | yes | Pristine Sash (253925, -0.57 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.70 DPS) [crafted]; Keller's Girdle (2911, -0.74 DPS) [world_drop] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 15.3 spell_power points (3.10 DPS) | yes | Filigreed Pristine Leggings (253937, -0.96 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.68 DPS) [dungeon]; Darkweave Breeches (12987, -1.97 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.1 spell_power points (2.05 DPS) | yes | Pearly Boots (253897, -0.46 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.97 DPS) [crafted]; Red Woolen Boots (4313, -1.20 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (1.33 DPS) | yes | Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.86 DPS) [world_drop]; Furlfeather Ring (286745, -1.12 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.1 spell_power points (1.25 DPS) | yes | Lavishly Jeweled Ring (1156, -0.65 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.78 DPS) [world_drop]; Furlfeather Ring (286745, -1.05 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 29.7 spell_power points (6.04 DPS) | yes | Staff of the Blessed Seer (2271, -1.03 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.16 DPS) [dungeon]; Buzzer Blade (2169, -1.57 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 3.8 spell_power points (0.78 DPS) | yes | Tear of Grief (5611, -0.02 DPS) [quest]; Dwarven Tome (279898, -0.16 DPS) [quest]; Seer's Fine Stein (7608, -0.30 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sizzle Stick (8071, -0.97 DPS, sim-verified) [quest]; Greater Magic Wand (11288, -1.01 DPS) [crafted]; Torchlight Wand (5240, -1.11 DPS) [quest] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Skycaller

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 93.0. Weights run: 2.5s. Verify run: 3.0s. 326 eligible items had no known source.

Race sweep: gnome 93.0 ± 0.1, high-order-skyborne 91.5 ± 0.1, human 91.5 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.010, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.924 per %), hit=0.235 ± 0.012 per rating point (10 rating = 1%, 2.352 per %), spell_haste=0.844 ± 0.143, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.583 ± 0.003, spirit=not significant (0.101 ± 0.088), mp5=not significant (0.232 ± 0.126)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.6 spell_power points (2.73 DPS) | yes | Enchanter's Cowl (4322, -0.01 DPS) [crafted]; Silk Headband (7050, -0.61 DPS) [crafted]; Filigreed Pearly Circlet (253983, -0.81 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.78 DPS) | yes | Seal of the Expedition (271920, -0.13 DPS) [rep]; Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.19 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.29 DPS) | yes | Death Speaker Mantle (6685, -0.64 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.94 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Expeditionary's Cape (271923, -0.06 DPS) [rep]; Prelacy Cape (7004, -0.09 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 16.5 spell_power points (3.87 DPS) | yes | Mechbuilder's Overalls (9508, -0.43 DPS) [dungeon]; Death Speaker Robes (6682, -0.79 DPS) [dungeon]; Green Silk Armor (7065, -0.90 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.7 spell_power points (2.28 DPS) | yes | Nightsky Wristbands (6407, -1.43 DPS) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.94 DPS, sim-verified) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 9.1 spell_power points (2.15 DPS) | yes | Shilly Mitts (9609, -0.29 DPS) [quest]; Serpent Gloves (5970, -0.50 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.53 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (2.98 DPS) | yes | Belt of Arugal (6392, -0.35 DPS) [dungeon]; Crimson Silk Belt (7055, -0.65 DPS) [crafted]; Invoker's Cord (215366, -0.68 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.9 spell_power points (3.04 DPS) | yes | Pearly Leggings (253995, -0.09 DPS) [crafted]; Pristine Leggings (253987, -0.36 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.65 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 11.9 spell_power points (2.79 DPS) | yes | Gilded Slippers (254001, -0.14 DPS) [crafted]; Acidic Walkers (9454, -0.47 DPS) [dungeon]; Spidersilk Boots (4320, -0.62 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.6 spell_power points (2.49 DPS) | yes | Runemender's Seal (273031, -0.84 DPS) [world]; Ring of Power Regulation (273030, -0.92 DPS) [dungeon]; Minor Channeling Ring (1449, -1.05 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.5 spell_power points (1.75 DPS) | yes | Runemender's Seal (273031, -0.10 DPS) [world]; Ring of Power Regulation (273030, -0.18 DPS) [dungeon]; Minor Channeling Ring (1449, -0.32 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.04 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.04 DPS) [world]; Glimmering Staff (249392, -3.01 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (2.43 DPS) | yes | Orb of Souls (249395, -1.35 DPS) [crafted]; Arcane Infused Rod (279838, -1.35 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.49 DPS) [quest] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.12 DPS) [world_drop]; Greater Mystic Wand (217287, -0.29 DPS) [crafted]; Wand of Eventide (5214, -0.40 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 175.4. Weights run: 2.5s. Verify run: 4.6s. 456 eligible items had no known source.

Race sweep: gnome 175.4 ± 0.3, human 174.3 ± 0.3, high-order-skyborne 173.7 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.557 ± 0.023, crit=0.268 ± 0.006 per rating point (14 rating = 1%, 3.747 per %), hit=0.665 ± 0.034 per rating point (10 rating = 1%, 6.653 per %), spell_haste=2.866 ± 0.400, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.635 ± 0.002, spirit=not significant (0.045 ± 0.219), mp5=not significant (0.181 ± 0.263)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 21.5 spell_power points (7.03 DPS) | yes | Papal Fez (9431, -1.17 DPS) [dungeon]; Corpseshroud (10574, -1.27 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.98 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.5 spell_power points (3.75 DPS) | yes | Triune Amulet (7722, -1.03 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.79 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -2.37 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.2 spell_power points (4.66 DPS) | yes | Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.55 DPS) [quest]; Green Silken Shoulders (7057, -1.26 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.49 DPS) [crafted]; Expeditionary's Cape (271922, -0.81 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.81 DPS, sim-verified) [world_drop]; Robe of Power (7054, -0.85 DPS) [crafted]; Elemental Raiment (9434, -0.87 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.5 spell_power points (3.12 DPS) | yes | Condor Bracers (15864, -0.83 DPS) [quest]; Windchaser Cuffs (14429, -1.48 DPS) [world_drop]; Arcane Runed Bracers (4744, -1.74 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (6.72 DPS) | yes | Black Mageweave Gloves (10003, -1.81 DPS) [crafted]; Red Mageweave Gloves (10018, -2.31 DPS, sim-verified) [crafted]; Golden Handwraps (254029, -2.54 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.2 spell_power points (5.31 DPS) | yes | Deathmage Sash (10771, -0.29 DPS) [dungeon]; Golden Cord (254045, -0.85 DPS) [crafted]; Star Belt (4329, -1.00 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.7 spell_power points (6.77 DPS) | yes | Crimson Silk Pantaloons (7062, -1.44 DPS, sim-verified) [crafted]; Pearly Leggings (253995, -2.51 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.71 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.3 spell_power points (7.96 DPS) | yes | Golden Slippers (254009, -4.24 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -4.33 DPS) [crafted]; Acidic Walkers (9454, -4.80 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (4.18 DPS) | yes | Lorekeeper's Ring (19524, -1.06 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.57 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.2 spell_power points (3.34 DPS) | yes | Lorekeeper's Ring (19524, -0.22 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.73 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -10.18 DPS) [quest]; Hypnotic Blade (7714, -10.79 DPS) [dungeon]; Illusionary Rod (7713, -11.10 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Twisted Nether Wand (249144, -0.58 DPS) [crafted]; Starfaller (13063, -0.61 DPS) [world_drop]; Plaguerot Sprig (10766, -1.92 DPS, sim-verified) [dungeon] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 141005213100301531-40000000000000000-033000000000000000)

Set DPS (verified): 278.5. Weights run: 2.5s. Verify run: 3.6s. 598 eligible items had no known source.

Race sweep: gnome 278.5 ± 0.4, human 277.3 ± 0.4, high-order-skyborne 274.6 ± 0.4 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.644 ± 0.026, crit=0.356 ± 0.008 per rating point (14 rating = 1%, 4.977 per %), hit=0.761 ± 0.040 per rating point (10 rating = 1%, 7.610 per %), spell_haste=3.255 ± 0.410, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.640 ± 0.002, spirit=not significant (0.547 ± 0.223), mp5=not significant (0.500 ± 0.252)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 34.0 spell_power points (12.60 DPS) | yes | Red Mageweave Headband (10033, -0.79 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -2.60 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.70 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.9 spell_power points (7.36 DPS) | yes | Mindburst Medallion (11196, -2.69 DPS) [quest]; Talvash's Enhancing Necklace (7673, -2.83 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -2.86 DPS) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.7 spell_power points (9.50 DPS) | yes | Rotgrip Mantle (17732, -0.39 DPS) [dungeon]; Nethershine Shoulders (254057, -2.35 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.55 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.75 DPS, sim-verified) [crafted]; Expeditionary's Cape (271921, -3.48 DPS) [rep]; Featherskin Cape (10843, -3.68 DPS) [world] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 31.5 spell_power points (11.65 DPS) | yes | Robe of the Magi (1716, -1.06 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -1.52 DPS) [world]; Robes of Insight (940, -1.85 DPS) [world_drop] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 15.6 spell_power points (5.77 DPS) | yes | Nethergeld Cuffs (254061, -0.49 DPS) [crafted]; Aristocratic Cuffs (12546, -0.71 DPS) [dungeon]; Bloodband Bracers (11469, -1.16 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.4 spell_power points (9.04 DPS) | yes | Golden Gloves (254103, -0.53 DPS) [crafted]; Runecloth Gloves (13863, -0.62 DPS) [crafted]; Gilded Gloves (254095, -1.33 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.0 spell_power points (9.65 DPS) | yes | Satyrmane Sash (17755, -2.08 DPS) [dungeon]; Ghostweave Cord (254073, -2.23 DPS) [crafted]; Dawnspire Cord (12466, -2.90 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 29.4 spell_power points (10.91 DPS) | yes | Senior Designer's Pantaloons (11841, -1.14 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -1.75 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.85 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (10.31 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -2.38 DPS) [vendor]; Mistwalker Boots (10629, -3.33 DPS) [dungeon]; Golden Sandals (254115, -3.42 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 20.2 spell_power points (7.47 DPS) | yes | Snake Hoop (6750, -2.16 DPS) [quest]; Lorekeeper's Ring (19523, -2.29 DPS) [rep]; Philanthropist's Ring (281635, -2.58 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.7 spell_power points (5.82 DPS) | yes | Snake Hoop (6750, -0.50 DPS) [quest]; Lorekeeper's Ring (19523, -0.63 DPS) [rep]; Philanthropist's Ring (281635, -0.92 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.10 DPS) [world_drop]; Uther's Strength (11302, -0.31 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.51 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -8.23 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.2 spell_power points (5.61 DPS) | yes | Celestial Orb (7515, -0.08 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.43 DPS) [dungeon]; Enthralled Sphere (11625, -0.64 DPS) [dungeon] |
| ranged | Charged Lightning Rod (11860) | Ledger from Tanaris [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nether Force Wand (11263, +0.00 DPS) [quest]; Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Dreambough Wand (249234, -0.11 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Nethershine Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Charged Lightning Rod

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 141005213100301531-40000000000000000-053500030000000000)

Set DPS (verified): 419.7. Weights run: 2.8s. Verify run: 19.2s. 1442 eligible items had no known source.

Race sweep: gnome 419.7 ± 0.6, human 418.3 ± 0.6, high-order-skyborne 415.6 ± 0.6 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.725 ± 0.032, crit=0.436 ± 0.011 per rating point (14 rating = 1%, 6.104 per %), hit=0.807 ± 0.044 per rating point (10 rating = 1%, 8.074 per %), spell_haste=3.532 ± 0.520, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.594 ± 0.003, spirit=not significant (0.377 ± 0.363), mp5=not significant (0.543 ± 0.349)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Crimson Felt Hat (18727, -10.57 DPS, sim-verified) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 24.8 spell_power points (7.26 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.57 DPS) [world_drop]; Arcane Crystal Pendant (20037, -1.30 DPS) [quest]; Nacreous Shell Necklace (22403, -4.09 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 32.5 spell_power points (9.54 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Burial Shawl (18681, -0.27 DPS) [dungeon]; Theramore Shoulderpads (271953, -0.28 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.34 DPS) [quest]; Spritecaster Cape (11623, -2.71 DPS) [dungeon]; Crystalline Threaded Cape (20697, -7.11 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.32 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.95 DPS) [pvp]; Alanna's Embrace (13314, -13.53 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 30.4 spell_power points (8.92 DPS) | yes | Runecloth Cuffs (254123, -3.57 DPS) [crafted]; Manacle Cuffs (11962, -3.68 DPS) [quest]; Marshal's Silk Bracers (16438, -3.78 DPS) [pvp] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 32.5 spell_power points (9.54 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 36.4 spell_power points (10.67 DPS) | yes | Stormpike Cloth Girdle (19094, -3.27 DPS) [rep]; Thuzadin Sash (18740, -3.39 DPS) [dungeon]; Ban'thok Sash (11662, -6.77 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.5 spell_power points (15.08 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -2.97 DPS) [pvp]; Skyshroud Leggings (13170, -3.74 DPS) [dungeon] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -7.50 DPS, sim-verified) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Blessed Band of Light (272407, +0.00 DPS) [vendor]; Maiden's Circle (13001, -8.21 DPS, sim-verified) [world_drop] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Blessed Band of Light (272407, -0.05 DPS) [vendor]; Songstone of Ironforge (12543, -5.94 DPS, sim-verified) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -2.91 DPS) [dungeon]; Serenity Field (272439, -3.22 DPS) [vendor]; Frozen Heart of the Mountain (249469, -4.16 DPS, sim-verified) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Blade of Eternal Darkness (17780, -4.08 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (8.50 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -1.13 DPS) [pvp]; Tome of the Lost (22253, -4.73 DPS, sim-verified) [dungeon] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.61 DPS) [dungeon]; Oblivion's Touch (18761, -1.59 DPS) [dungeon] |

**New at 60:** head: The Postmaster's Band; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 804.0. Weights run: 3.0s. Verify run: 20.3s. 1442 eligible items had no known source.

Race sweep: gnome 804.0 ± 0.9, human 799.5 ± 0.9, high-order-skyborne 792.5 ± 0.9 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.871 ± 0.043, crit=0.548 ± 0.026 per rating point (14 rating = 1%, 7.671 per %), hit=1.223 ± 0.078 per rating point (10 rating = 1%, 12.232 per %), spell_haste=not significant (2.359 ± 0.642), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.569 ± 0.002, spirit=not significant (0.589 ± 0.487), mp5=not significant (0.623 ± 0.488)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 44.6 spell_power points (19.62 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; The Postmaster's Band (13390, -0.85 DPS) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 28.2 spell_power points (12.41 DPS) | yes | Lady Maye's Pendant (14558, -0.29 DPS) [world_drop]; Star of Mystaria (12103, -0.92 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -4.23 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.0 spell_power points (15.84 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Theramore Shoulderpads (271953, -0.21 DPS) [rep]; Lieutenant Commander's Silk Mantle (227102, -0.61 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.15 DPS) [world]; Spritecaster Cape (11623, -5.47 DPS) [dungeon]; Shroud of the Exile (15421, -7.40 DPS, sim-verified) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -5.15 DPS) [pvp]; The Postmaster's Tunic (13388, -17.47 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 33.1 spell_power points (14.56 DPS) | yes | Marshal's Silk Bracers (16438, -4.69 DPS) [pvp]; Manacle Cuffs (11962, -5.73 DPS) [quest]; Runecloth Cuffs (254123, -5.89 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Hands of Power (13253, -5.52 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 39.9 spell_power points (17.56 DPS) | yes | Thuzadin Sash (18740, -4.83 DPS) [dungeon]; Marshal's Silk Sash (16439, -4.86 DPS) [pvp]; Ban'thok Sash (11662, -6.53 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 56.0 spell_power points (24.66 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -4.24 DPS) [pvp]; Magister's Leggings (16687, -5.55 DPS) [dungeon] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 31.9 spell_power points (14.05 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Kayser's Boots of Precision (22231, -0.01 DPS) [dungeon]; Omnicast Boots (11822, -0.65 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.96 DPS) [vendor]; Songstone of Ironforge (12543, -2.22 DPS) [quest]; Naglering (11669, -20.49 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.55 DPS) [quest]; Seal of Rivendare (13345, -2.12 DPS) [dungeon]; Blessed Band of Light (272407, -4.33 DPS, sim-verified) [vendor] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -5.58 DPS, sim-verified) [dungeon] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.54 DPS) [crafted]; Ankh of Life (1713, -2.27 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -0.01 DPS) [pvp]; Headmaster's Charge (13937, -6.44 DPS) [dungeon]; Kindling Stave (11750, -21.62 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.24 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.76 DPS) [dungeon]; Oblivion's Touch (18761, -2.40 DPS) [dungeon] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Eye of the Beast; trinket2: Smolderweb's Eye; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 46.9. Weights run: 2.4s. Verify run: 2.9s. 192 eligible items had no known source.

Race sweep: undead 46.9 ± 0.1, orc 45.4 ± 0.1, troll 44.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.571 ± 0.007, crit=0.056 ± 0.001 per rating point (14 rating = 1%, 0.779 per %), hit=0.224 ± 0.008 per rating point (10 rating = 1%, 2.241 per %), spell_haste=0.311 ± 0.055, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.855 ± 0.005, spirit=not significant (0.122 ± 0.044), mp5=0.256 ± 0.063

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 9.4 spell_power points (1.33 DPS) | yes | Shadow Goggles (4373, -0.82 DPS) [crafted]; Pristine Circlet (253949, -0.87 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.26 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.4 spell_power points (0.05 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.1 spell_power points (1.43 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.87 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.91 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.32 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.11 DPS) [crafted]; Battle Healer's Cloak (20427, -0.18 DPS) [rep]; Black Whelp Cloak (7283, -0.21 DPS) [crafted] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 10.9 spell_power points (1.54 DPS) | yes | Gray Woolen Robe (2585, -0.57 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.70 DPS, sim-verified) [crafted]; Green Woolen Robe (6243, -0.88 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 3.6 spell_power points (0.51 DPS) | yes | Bright Bracers (3647, -0.19 DPS) [world_drop]; Owlbeard Bracers (16981, -0.21 DPS) [quest]; Tabitha's Cuffs (251486, -0.61 DPS, sim-verified) [quest] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 7.7 spell_power points (1.09 DPS) | yes | Gnoll Casting Gloves (892, -0.24 DPS) [world]; Pristine Gloves (253913, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.52 DPS, sim-verified) [dungeon] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 8.3 spell_power points (1.17 DPS) | yes | Novice Arcanist's Sash (253885, -0.36 DPS) [crafted]; Pristine Sash (253925, -0.52 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.52 DPS) [world_drop] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 13.3 spell_power points (1.88 DPS) | yes | Filigreed Pristine Leggings (253937, -0.87 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.89 DPS) [dungeon]; Colorful Kilt (10048, -1.09 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.3 spell_power points (1.31 DPS) | yes | Pearly Boots (253897, -0.53 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.65 DPS) [crafted]; Red Woolen Boots (4313, -0.66 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.5 spell_power points (0.78 DPS) | yes | Loop of Sacrifice (281673, -0.38 DPS) [quest]; Volcanic Rock Ring (12053, -0.54 DPS) [world_drop]; Furlfeather Ring (286745, -0.64 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.4 spell_power points (0.48 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop]; Furlfeather Ring (286745, -0.34 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.08 DPS) [dungeon]; Crescent Staff (6505, -0.12 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.36 DPS) [quest] |
| off_hand | Seer's Fine Stein (7608) | World drop [world_drop] | 1.8 spell_power points (0.26 DPS) | yes | Spellbinder Orb (15926, -0.02 DPS) [dungeon]; Tork Wrench (11855, -0.10 DPS) [quest]; Ancestral Orb (15944, -0.18 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -0.65 DPS) [dungeon]; Greater Magic Wand (11288, -0.84 DPS) [crafted]; Sizzle Stick (8071, -1.02 DPS, sim-verified) [quest] |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Seer's Fine Stein; ranged: Skycaller

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 92.9. Weights run: 2.5s. Verify run: 3.7s. 318 eligible items had no known source.

Race sweep: undead 92.9 ± 0.1, orc 91.7 ± 0.1, troll 90.5 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.010, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.924 per %), hit=0.235 ± 0.012 per rating point (10 rating = 1%, 2.352 per %), spell_haste=0.844 ± 0.143, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.583 ± 0.003, spirit=not significant (0.101 ± 0.088), mp5=not significant (0.232 ± 0.126)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.6 spell_power points (2.73 DPS) | yes | Enchanter's Cowl (4322, -0.01 DPS) [crafted]; Silk Headband (7050, -0.61 DPS) [crafted]; Filigreed Pearly Circlet (253983, -0.81 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.78 DPS) | yes | Darkspear Warding Pendant (272075, -0.13 DPS) [vendor]; Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.19 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.29 DPS) | yes | Death Speaker Mantle (6685, -0.44 DPS) [dungeon]; Magician's Mantle (12998, -0.94 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.09 DPS) [crafted]; Darkspear Raider's Cloak (272078, -0.16 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.24 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 16.5 spell_power points (3.87 DPS) | yes | Green Silk Armor (7065, -0.06 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.43 DPS) [dungeon]; Death Speaker Robes (6682, -0.79 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.7 spell_power points (2.28 DPS) | yes | Glowing Magical Bracelets (13106, -1.16 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.43 DPS) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 9.1 spell_power points (2.15 DPS) | yes | Jutebraid Gloves (10654, -0.08 DPS) [quest]; Serpent Gloves (5970, -0.50 DPS) [dungeon]; Truefaith Gloves (7049, -0.58 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.7 spell_power points (2.98 DPS) | yes | Warsong Sash (16975, -0.39 DPS) [quest]; Belt of Arugal (6392, -0.55 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.65 DPS) [crafted] |
| legs | Pearly Leggings (253995) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.27 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.45 DPS, sim-verified) [dungeon]; Filigreed Pearly Leggings (253945, -0.57 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 11.9 spell_power points (2.79 DPS) | yes | Gilded Slippers (254001, -0.14 DPS) [crafted]; Acidic Walkers (9454, -0.47 DPS) [dungeon]; Spidersilk Boots (4320, -0.62 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.6 spell_power points (2.49 DPS) | yes | Runemender's Seal (273031, -0.84 DPS) [world]; Ring of Power Regulation (273030, -0.92 DPS) [dungeon]; Solid Gold Band (279845, -1.43 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.5 spell_power points (1.75 DPS) | yes | Runemender's Seal (273031, -0.10 DPS) [world]; Ring of Power Regulation (273030, -0.18 DPS) [dungeon]; Solid Gold Band (279845, -0.70 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.04 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.04 DPS) [world]; Glimmering Staff (249392, -2.50 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (2.43 DPS) | yes | Witch's Finger (16887, -1.38 DPS) [quest]; Satyr's Rod (15962, -1.60 DPS) [world_drop]; Orb of Souls (249395, -1.63 DPS, sim-verified) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.16 DPS) [world_drop]; Greater Mystic Wand (217287, -0.32 DPS) [crafted]; Wand of Eventide (5214, -0.45 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Pearly Leggings; feet: Golden Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 174.6. Weights run: 2.5s. Verify run: 4.8s. 448 eligible items had no known source.

Race sweep: undead 174.6 ± 0.3, orc 173.0 ± 0.3, troll 171.8 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.744 ± 0.022, crit=0.169 ± 0.004 per rating point (14 rating = 1%, 2.366 per %), hit=0.417 ± 0.025 per rating point (10 rating = 1%, 4.165 per %), spell_haste=1.334 ± 0.327, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.619 ± 0.003, spirit=not significant (0.809 ± 0.280), mp5=not significant (0.904 ± 0.280)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Whitemane's Chapeau (7720, -0.13 DPS) [dungeon]; Papal Fez (9431, -0.55 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.86 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -0.52 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.52 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.99 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.12 DPS) [crafted]; Bloodmage Mantle (7684, -0.25 DPS) [dungeon]; Mantle of Doan (7712, -0.85 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.11 DPS) [vendor]; Cloak of Blight (6832, -0.34 DPS) [quest]; Battle Healer's Cloak (19528, -0.46 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 36.0 spell_power points (9.06 DPS) | yes | Robe of the Magi (1716, -1.39 DPS) [world_drop]; Robe of Power (7054, -1.67 DPS) [crafted]; Beguiler Robes (7728, -2.92 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.7 spell_power points (2.95 DPS) | yes | Radiant Silver Bracers (4545, -0.44 DPS) [quest]; Earthen Silk Cuffs (254019, -0.92 DPS) [crafted]; Condor Bracers (15864, -1.19 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 26.6 spell_power points (6.70 DPS) | yes | Red Mageweave Gloves (10018, -1.72 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.93 DPS) [crafted]; Golden Handwraps (254029, -3.21 DPS) [crafted] |
| waist | Golden Cord (254045) | Tailoring [crafted] | 19.5 spell_power points (4.90 DPS) | yes | Gilded Cord (254037, -0.17 DPS) [crafted]; Deathmage Sash (10771, -0.33 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.63 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.9 spell_power points (5.77 DPS) | yes | Black Mageweave Leggings (9999, -0.90 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.92 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.66 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 29.7 spell_power points (7.46 DPS) | yes | Golden Slippers (254009, -3.31 DPS) [crafted]; Thoughtcast Boots (10578, -3.56 DPS) [dungeon]; Furen's Boots (13100, -5.96 DPS, sim-verified) [world_drop] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 16.9 spell_power points (4.24 DPS) | yes | Runemender's Seal (273031, -0.94 DPS) [world]; Advisor's Ring (19520, -1.30 DPS) [rep]; Welken Ring (5011, -1.88 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (3.45 DPS) | yes | Advisor's Ring (19520, -0.50 DPS) [rep]; Welken Ring (5011, -1.09 DPS) [world_drop]; Runemender's Seal (273031, -1.86 DPS, sim-verified) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -9.27 DPS) [dungeon]; Celestial Stave (9517, -9.37 DPS) [quest]; Illusionary Rod (7713, -11.28 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -0.19 DPS) [world_drop]; Dancing Flame (6806, -0.34 DPS) [quest]; Starfaller (13063, -0.38 DPS) [world_drop] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Golden Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 141005213100301531-40000000000000000-033000000000000000)

Set DPS (verified): 274.4. Weights run: 2.5s. Verify run: 4.2s. 590 eligible items had no known source.

Race sweep: undead 274.4 ± 0.4, orc 272.1 ± 0.4, troll 270.3 ± 0.4 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.644 ± 0.026, crit=0.356 ± 0.008 per rating point (14 rating = 1%, 4.977 per %), hit=0.761 ± 0.040 per rating point (10 rating = 1%, 7.610 per %), spell_haste=3.255 ± 0.410, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.640 ± 0.002, spirit=not significant (0.547 ± 0.223), mp5=not significant (0.500 ± 0.252)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 34.0 spell_power points (12.60 DPS) | yes | Red Mageweave Headband (10033, -0.79 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -2.60 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.70 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.9 spell_power points (7.36 DPS) | yes | Jarkal's Enhancing Necklace (7888, -2.07 DPS, sim-verified) [quest]; Mindburst Medallion (11196, -2.69 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.86 DPS) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.7 spell_power points (9.50 DPS) | yes | Rotgrip Mantle (17732, -0.39 DPS) [dungeon]; Nethershine Shoulders (254057, -2.35 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.55 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.24 DPS) [quest]; Runecloth Cloak (13860, -2.59 DPS) [crafted]; Darkspear Raider's Cloak (272076, -3.48 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 31.5 spell_power points (11.65 DPS) | yes | Robe of the Magi (1716, -1.06 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -1.52 DPS) [world]; Robes of Insight (940, -1.85 DPS) [world_drop] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 15.6 spell_power points (5.77 DPS) | yes | Nethergeld Cuffs (254061, -0.49 DPS) [crafted]; Aristocratic Cuffs (12546, -0.71 DPS) [dungeon]; Bloodband Bracers (11469, -1.16 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.4 spell_power points (9.04 DPS) | yes | Golden Gloves (254103, -0.53 DPS) [crafted]; Runecloth Gloves (13863, -0.62 DPS) [crafted]; Gilded Gloves (254095, -1.33 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.0 spell_power points (9.65 DPS) | yes | Ghostweave Cord (254073, -2.23 DPS) [crafted]; Satyrmane Sash (17755, -2.55 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -2.90 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -0.61 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.71 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (10.31 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.38 DPS) [vendor]; Mistwalker Boots (10629, -3.33 DPS) [dungeon]; Golden Sandals (254115, -3.42 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 20.2 spell_power points (7.47 DPS) | yes | Snake Hoop (6750, -2.16 DPS) [quest]; Advisor's Ring (19519, -2.29 DPS) [rep]; Philanthropist's Ring (281635, -2.58 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.7 spell_power points (5.82 DPS) | yes | Snake Hoop (6750, -0.50 DPS) [quest]; Advisor's Ring (19519, -0.63 DPS) [rep]; Philanthropist's Ring (281635, -0.92 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -0.51 DPS) [quest]; Rune of the Guard Captain (19120, -0.56 DPS) [quest]; Ankh of Life (1713, -0.87 DPS, sim-verified) [world_drop] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Alchemist's Stone (13503, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -0.20 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -8.11 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.2 spell_power points (5.61 DPS) | yes | Celestial Orb (7515, -0.08 DPS) [quest]; Flask of Forest Mojo (19115, -0.37 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.43 DPS) [dungeon] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nether Force Wand (11263, -0.72 DPS) [quest]; Charged Lightning Rod (11860, -0.77 DPS) [quest]; Noxious Shooter (17745, -0.80 DPS) [dungeon] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Nethershine Cuffs; waist: Ban'thok Sash; legs: Senior Designer's Pantaloons; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Lesser Eternal Wand

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 141005213100301531-40000000000000000-053500030000000000)

Set DPS (verified): 415.3. Weights run: 2.8s. Verify run: 19.4s. 1436 eligible items had no known source.

Race sweep: undead 415.3 ± 0.6, orc 411.2 ± 0.6, troll 407.6 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.725 ± 0.032, crit=0.436 ± 0.011 per rating point (14 rating = 1%, 6.104 per %), hit=0.807 ± 0.044 per rating point (10 rating = 1%, 8.074 per %), spell_haste=3.532 ± 0.520, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.594 ± 0.003, spirit=not significant (0.377 ± 0.363), mp5=not significant (0.543 ± 0.349)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Crimson Felt Hat (18727, -8.58 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.01 DPS) [world_drop]; Arcane Crystal Pendant (20037, -0.74 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 32.5 spell_power points (9.54 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Burial Shawl (18681, -0.27 DPS) [dungeon]; Darkspear Shoulderpads (272103, -0.28 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.34 DPS) [quest]; Spritecaster Cape (11623, -2.71 DPS) [dungeon]; Crystalline Threaded Cape (20697, -2.89 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.32 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.95 DPS) [pvp]; Alanna's Embrace (13314, -11.90 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 30.4 spell_power points (8.92 DPS) | yes | Runecloth Cuffs (254123, -3.57 DPS) [crafted]; Manacle Cuffs (11962, -3.68 DPS) [quest]; General's Silk Cuffs (16538, -3.78 DPS) [pvp] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 32.5 spell_power points (9.54 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 36.4 spell_power points (10.67 DPS) | yes | Ban'thok Sash (11662, -2.68 DPS, sim-verified) [dungeon]; Frostwolf Cloth Belt (19090, -3.27 DPS) [rep]; Thuzadin Sash (18740, -3.39 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 45.6 spell_power points (13.35 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -1.24 DPS) [pvp]; Skyshroud Leggings (13170, -6.36 DPS, sim-verified) [dungeon] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -7.49 DPS, sim-verified) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Blessed Band of Light (272407, +0.00 DPS) [vendor]; Maiden's Circle (13001, -7.83 DPS, sim-verified) [world_drop] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Maiden's Circle (13001, +0.00 DPS) [world_drop]; Blessed Band of Light (272407, -0.05 DPS) [vendor] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -2.91 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.15 DPS) [crafted]; Draconic Infused Emblem (22268, -6.91 DPS, sim-verified) [dungeon] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Sageblade (22383, -4.20 DPS, sim-verified) [crafted] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (8.50 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -1.13 DPS) [pvp]; Tome of the Lost (22253, -3.25 DPS, sim-verified) [dungeon] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.63 DPS) [dungeon]; Oblivion's Touch (18761, -1.60 DPS) [dungeon] |

**New at 60:** head: The Postmaster's Band; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 791.5. Weights run: 3.0s. Verify run: 21.9s. 1436 eligible items had no known source.

Race sweep: undead 791.5 ± 0.9, orc 789.6 ± 0.9, troll 780.3 ± 0.9 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.871 ± 0.043, crit=0.548 ± 0.026 per rating point (14 rating = 1%, 7.671 per %), hit=1.223 ± 0.078 per rating point (10 rating = 1%, 12.232 per %), spell_haste=not significant (2.359 ± 0.642), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.569 ± 0.002, spirit=not significant (0.589 ± 0.487), mp5=not significant (0.623 ± 0.488)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 44.6 spell_power points (19.62 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; The Postmaster's Band (13390, -0.85 DPS) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 28.2 spell_power points (12.41 DPS) | yes | Lady Maye's Pendant (14558, -0.29 DPS) [world_drop]; Star of Mystaria (12103, -0.92 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -4.06 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.0 spell_power points (15.84 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.21 DPS) [vendor]; Champion's Silk Mantle (227104, -0.61 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -4.56 DPS) [quest]; Crystalline Threaded Cape (20697, -5.15 DPS) [world]; Cloak of Flames (3475, -12.95 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -5.15 DPS) [pvp]; The Postmaster's Tunic (13388, -16.32 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 33.1 spell_power points (14.56 DPS) | yes | General's Silk Cuffs (16538, -4.69 DPS) [pvp]; Manacle Cuffs (11962, -5.73 DPS) [quest]; Runecloth Cuffs (254123, -5.89 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Hands of Power (13253, -5.56 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 39.9 spell_power points (17.56 DPS) | yes | Ban'thok Sash (11662, -3.06 DPS) [dungeon]; Thuzadin Sash (18740, -4.83 DPS) [dungeon]; General's Silk Sash (16537, -4.86 DPS) [pvp] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 50.4 spell_power points (22.19 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -1.77 DPS) [pvp]; Magister's Leggings (16687, -4.06 DPS, sim-verified) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.64 DPS) [dungeon]; Dragonrider Boots (18102, -2.07 DPS, sim-verified) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -0.66 DPS) [world_drop]; Eye of Orgrimmar (12545, -2.22 DPS) [quest]; Naglering (11669, -12.92 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Orgrimmar (12545, -1.25 DPS) [quest]; Seal of Rivendare (13345, -1.82 DPS) [dungeon]; Maiden's Circle (13001, -2.37 DPS, sim-verified) [world_drop] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -5.22 DPS, sim-verified) [crafted] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.54 DPS) [crafted] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -0.01 DPS) [pvp]; Headmaster's Charge (13937, -6.44 DPS) [dungeon]; Kindling Stave (11750, -17.00 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.24 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.76 DPS) [dungeon]; Oblivion's Touch (18761, -2.40 DPS) [dungeon] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Smolderweb's Eye; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

