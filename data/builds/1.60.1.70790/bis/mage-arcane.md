# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 47.9. Weights run: 3.3s. Verify run: 2.7s. 181 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.817 ± 0.013, crit=0.079 ± 0.002 per rating point (14 rating = 1%, 1.105 per %), hit=0.328 ± 0.013 per rating point (10 rating = 1%, 3.281 per %), spell_haste=not significant (0.117 ± 0.044), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.953 ± 0.003, spirit=not significant (0.158 ± 0.047), mp5=0.727 ± 0.085

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 10.5 spell_power points (2.11 DPS) | yes | Pristine Circlet (253949, -0.95 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.10 DPS) [crafted]; Flying Tiger Goggles (4368, -1.98 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.06 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.4 spell_power points (2.49 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.89 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.68 DPS) [crafted]; Slime-encrusted Pads (6461, -2.05 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (47.9 DPS) | yes | Pearl-clasped Cloak (5542, -0.04 DPS) [crafted]; Caretaker's Cape (20428, -0.26 DPS) [rep]; Sanguine Cape (14376, -0.27 DPS) [world_drop] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 13.1 spell_power points (2.64 DPS) | yes | Filigreed Pristine Gown (253901, -0.75 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.02 DPS) [crafted]; Mystic's Wrap (14369, -1.49 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.3 spell_power points (1.26 DPS) | yes | Bright Bracers (3647, -0.72 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.87 DPS) [quest]; Mystic's Bracelets (14366, -0.93 DPS) [world_drop] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.1 spell_power points (1.84 DPS) | yes | Serpent Gloves (5970, -0.50 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.54 DPS) [crafted]; Gnoll Casting Gloves (892, -0.63 DPS) [world] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 9.9 spell_power points (2.00 DPS) | yes | Pristine Sash (253925, -0.56 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.69 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.70 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 16.0 spell_power points (3.22 DPS) | yes | Filigreed Pristine Leggings (253937, -0.95 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.82 DPS) [dungeon]; Darkweave Breeches (12987, -1.88 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.3 spell_power points (2.07 DPS) | yes | Pearly Boots (253897, -0.40 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.97 DPS) [crafted]; Red Woolen Boots (4313, -1.10 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.6 spell_power points (1.34 DPS) | yes | Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.84 DPS) [world_drop]; Furlfeather Ring (286745, -1.13 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.5 spell_power points (1.30 DPS) | yes | Lavishly Jeweled Ring (1156, -0.49 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.81 DPS) [world_drop]; Furlfeather Ring (286745, -1.10 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.4 spell_power points (6.12 DPS) | yes | Staff of the Blessed Seer (2271, -1.19 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.29 DPS) [dungeon]; Buzzer Blade (2169, -1.70 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 4.1 spell_power points (0.82 DPS) | yes | Tear of Grief (5611, -0.10 DPS) [quest]; Dwarven Tome (279898, -0.19 DPS) [quest]; Seer's Fine Stein (7608, -0.30 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-verified (47.9 DPS) | yes | Sizzle Stick (8071, -0.98 DPS, sim-verified) [quest]; Greater Magic Wand (11288, -1.02 DPS) [crafted]; Cookie's Stirring Rod (5198, -1.10 DPS) [dungeon] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Skycaller

No-known-source sample (15 of 181, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 141005213100300000-00000000000000000-000000000000000000)

Set DPS (verified): 91.1. Weights run: 3.5s. Verify run: 5.7s. 309 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.010, crit=0.067 ± 0.002 per rating point (14 rating = 1%, 0.937 per %), hit=0.243 ± 0.012 per rating point (10 rating = 1%, 2.431 per %), spell_haste=0.856 ± 0.108, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.593 ± 0.003, spirit=not significant (0.217 ± 0.081), mp5=not significant (0.200 ± 0.116)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.3 spell_power points (2.89 DPS) | yes | Enchanter's Cowl (4322, -0.18 DPS) [crafted]; Silk Headband (7050, -0.78 DPS) [crafted]; Filigreed Pearly Circlet (253983, -0.94 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.78 DPS) | yes | Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.11 DPS) [world_drop]; Seal of the Expedition (271920, -0.13 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.29 DPS) | yes | Death Speaker Mantle (6685, -0.44 DPS) [dungeon]; Magician's Mantle (12998, -0.94 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.02 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.02 DPS) [quest]; Heavy Woolen Cloak (4311, -0.05 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 17.4 spell_power points (4.09 DPS) | yes | Mechbuilder's Overalls (9508, -0.64 DPS) [dungeon]; Green Silk Armor (7065, -0.66 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.01 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.6 spell_power points (2.26 DPS) | yes | Glowing Magical Bracelets (13106, -1.14 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.32 DPS) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 9.2 spell_power points (2.17 DPS) | yes | Serpent Gloves (5970, -0.52 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.55 DPS) [quest]; Truefaith Gloves (7049, -0.60 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (2.98 DPS) | yes | Belt of Arugal (6392, -0.21 DPS) [dungeon]; Crimson Silk Belt (7055, -0.65 DPS) [crafted]; Invoker's Cord (215366, -0.68 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 14.0 spell_power points (3.28 DPS) | yes | Pearly Leggings (253995, -0.16 DPS) [crafted]; Pristine Leggings (253987, -0.47 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.76 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 12.5 spell_power points (2.93 DPS) | yes | Gilded Slippers (254001, -0.17 DPS) [crafted]; Acidic Walkers (9454, -0.51 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (2.06 DPS) | yes | Ring of Power Regulation (273030, -0.51 DPS) [dungeon]; Minor Channeling Ring (1449, -0.63 DPS) [quest]; Electrocutioner Lagnut (9447, -1.00 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.4 spell_power points (1.74 DPS) | yes | Ring of Power Regulation (273030, -0.25 DPS, sim-verified) [dungeon]; Minor Channeling Ring (1449, -0.30 DPS) [quest]; Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (91.1 DPS) | yes | Soulstaff (249393, -1.05 DPS) [crafted]; Tendonscraper Dagger (273053, -1.14 DPS) [world]; Glimmering Staff (249392, -2.82 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (2.43 DPS) | yes | Arcane Infused Rod (279838, -0.62 DPS, sim-verified) [quest]; Orb of Souls (249395, -1.18 DPS) [crafted]; Eye of Paleth (2943, -1.49 DPS) [quest] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-verified (91.1 DPS) | yes | Thunderwood (13062, -0.11 DPS) [world_drop]; Greater Mystic Wand (217287, -0.29 DPS) [crafted]; Wand of Eventide (5214, -0.38 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 173.5. Weights run: 4.0s. Verify run: 5.7s. 400 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.757 ± 0.022, crit=0.168 ± 0.004 per rating point (14 rating = 1%, 2.347 per %), hit=0.414 ± 0.026 per rating point (10 rating = 1%, 4.143 per %), spell_haste=not significant (1.110 ± 0.309), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.621 ± 0.003, spirit=not significant (0.416 ± 0.283), mp5=not significant (0.124 ± 0.286)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 24.6 spell_power points (6.23 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.91 DPS) [crafted]; Whitemane's Chapeau (7720, -1.28 DPS) [dungeon]; Augural Shroud (2620, -1.53 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.1 spell_power points (3.31 DPS) | yes | Triune Amulet (7722, -0.54 DPS) [dungeon]; Necklace of Calisea (1714, -1.23 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.39 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.8 spell_power points (4.27 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Mantle of Doan (7712, -0.50 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Expeditionary's Cape (271922, +0.00 DPS) [rep]; Sporid Cape (6629, -2.63 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 30.6 spell_power points (7.76 DPS) | yes | Robe of the Magi (1716, -0.51 DPS) [world_drop]; Robe of Power (7054, -1.07 DPS) [crafted]; Beguiler Robes (7728, -2.34 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.4 spell_power points (2.37 DPS) | yes | Arcane Runed Bracers (4744, -0.09 DPS) [quest]; Condor Bracers (15864, -0.60 DPS) [quest]; Windchaser Cuffs (14429, -0.65 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.9 spell_power points (6.07 DPS) | yes | Red Mageweave Gloves (10018, -1.51 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.26 DPS) [crafted]; Golden Handwraps (254029, -2.52 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.3 spell_power points (4.65 DPS) | yes | Golden Cord (254045, -0.28 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.34 DPS) [rep]; Gilded Cord (254037, -0.46 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.1 spell_power points (5.85 DPS) | yes | Crimson Silk Pantaloons (7062, -1.33 DPS) [crafted]; Stoneweaver Leggings (9407, -1.75 DPS, sim-verified) [dungeon]; Pearly Leggings (253995, -1.78 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.9 spell_power points (6.82 DPS) | yes | Gilded Slippers (254001, -3.28 DPS) [crafted]; Acidic Walkers (9454, -3.60 DPS) [dungeon]; Golden Slippers (254009, -4.32 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (3.49 DPS) | yes | Ring of Forlorn Spirits (2043, -1.47 DPS) [quest]; Reedknot Ring (9622, -1.72 DPS) [quest]; Snake Hoop (6750, -3.93 DPS, sim-verified) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of Forlorn Spirits (2043, -0.35 DPS) [quest]; Snake Hoop (6750, -0.46 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.60 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -8.48 DPS) [quest]; Hypnotic Blade (7714, -8.67 DPS) [dungeon]; Illusionary Rod (7713, -9.61 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, -0.37 DPS) [world_drop]; Jaina's Firestarter (13064, -0.48 DPS) [world_drop]; Fizzle's Zippy Lighter (6729, -0.53 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 141005213100301531-40000000000000000-051000000000000000)

Set DPS (verified): 255.0. Weights run: 3.9s. Verify run: 7.4s. 512 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.661 ± 0.026, crit=0.335 ± 0.008 per rating point (14 rating = 1%, 4.690 per %), hit=0.785 ± 0.042 per rating point (10 rating = 1%, 7.846 per %), spell_haste=2.231 ± 0.550, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.639 ± 0.002, spirit=not significant (0.328 ± 0.245), mp5=not significant (0.206 ± 0.268)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 32.2 spell_power points (11.78 DPS) | yes | Dreamweave Circlet (10041, -0.25 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.91 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.05 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.0 spell_power points (7.30 DPS) | yes | Glowing Eye of Mordresh (10769, -2.47 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.06 DPS) [quest]; Horizon Choker (13085, -3.44 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 24.9 spell_power points (9.10 DPS) | yes | Kentic Amice (11624, -0.12 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.29 DPS) [vendor]; Nethershine Shoulders (254057, -2.48 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.63 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -3.29 DPS) [crafted]; Expeditionary's Cape (271921, -3.31 DPS) [rep] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 32.2 spell_power points (11.78 DPS) | yes | Dreamweave Vest (10021, -1.34 DPS) [crafted]; Robe of the Magi (1716, -1.69 DPS) [world_drop]; Robe of Power (7054, -2.80 DPS) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 14.6 spell_power points (5.33 DPS) | yes | Nethergeld Cuffs (254061, -0.48 DPS) [crafted]; Bloodband Bracers (11469, -0.97 DPS) [quest]; Aristocratic Cuffs (12546, -0.98 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.9 spell_power points (8.39 DPS) | yes | Golden Gloves (254103, -0.57 DPS) [crafted]; Runecloth Gloves (13863, -0.75 DPS) [crafted]; Gilded Gloves (254095, -1.36 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Dawnspire Cord (12466, -0.75 DPS) [dungeon]; Golden Waistcord (254089, -0.92 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | sim-verified (255.0 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.76 DPS) [dungeon]; Red Mageweave Pants (10009, -0.98 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.3 spell_power points (9.62 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.65 DPS) [vendor]; Golden Sandals (254115, -2.77 DPS) [crafted]; Gilded Sandals (254107, -3.42 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.9 spell_power points (5.46 DPS) | yes | Band of the Unicorn (7553, -0.71 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.77 DPS) [rep]; Brainlash (6440, -1.24 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (4.86 DPS) | yes | Band of the Unicorn (7553, -0.11 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.18 DPS) [rep]; Brainlash (6440, -0.64 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.39 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.38 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.62 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.24 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.48 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -5.90 DPS) [world_drop]; Spellshifter Rod (9527, -6.97 DPS) [quest]; Kindling Stave (11750, -8.68 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -0.68 DPS) [dungeon]; Nether Force Wand (11263, -0.91 DPS) [quest]; Charged Lightning Rod (11860, -0.95 DPS) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethershine Cuffs; waist: Satyrmane Sash; legs: Knight's Dreadweave Leggings; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Ankh of Life; ranged: Lesser Eternal Wand

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 141005213100301531-40000000000000000-053500030000000000)

Set DPS (verified): 414.7. Weights run: 4.0s. Verify run: 25.4s. 1299 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.815 ± 0.043, crit=0.579 ± 0.014 per rating point (14 rating = 1%, 8.103 per %), hit=1.158 ± 0.067 per rating point (10 rating = 1%, 11.578 per %), spell_haste=4.854 ± 0.869, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.608 ± 0.002, spirit=not significant (0.803 ± 0.375), mp5=not significant (0.907 ± 0.395)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreadmist Mask (16698) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.7 spell_power points (14.08 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.16 DPS) [pvp]; Shroud of the Nathrezim (18720, -0.45 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -3.38 DPS, sim-verified) [dungeon]; Crystalline Threaded Cape (20697, -4.27 DPS) [world]; Shroud of Arcane Mastery (22330, -5.35 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.30 DPS) [pvp]; Alanna's Embrace (13314, -7.87 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 34.1 spell_power points (13.46 DPS) | yes | Marshal's Silk Bracers (16438, -4.19 DPS) [pvp]; Wyrmthalak's Shackles (13958, -5.82 DPS) [quest]; Runecloth Cuffs (254123, -5.91 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 38.6 spell_power points (15.20 DPS) | yes | Marshal's Silk Sash (16439, -3.06 DPS) [pvp]; Wisdom of the Timbermaw (19047, -3.52 DPS) [crafted]; Ban'thok Sash (11662, -5.93 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 56.4 spell_power points (22.22 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.70 DPS) [pvp]; Magister's Leggings (16687, -3.88 DPS) [dungeon] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | sim-verified (414.7 DPS) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Mooncloth Boots (15802, -0.37 DPS) [crafted] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.56 DPS) [vendor]; Songstone of Ironforge (12543, -0.98 DPS) [quest]; Naglering (11669, -12.42 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.63 DPS) [quest]; Blessed Band of Light (272407, -2.95 DPS, sim-verified) [vendor]; Deep Rooted Ring (19109, -3.30 DPS) [world] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+10.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.62 DPS) [crafted]; Ankh of Life (1713, -0.93 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.57 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.21 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.57 DPS) [crafted]; Ankh of Life (1713, -4.88 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (+7.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -3.46 DPS) [world_drop]; Ironbark Staff (20069, -7.18 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.44 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -1.02 DPS) [pvp]; Tome of the Lost (22253, -2.09 DPS) [dungeon] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torch of Light (279246, +0.00 DPS) [crafted]; Bonecreeper Stylus (13938, -0.01 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.23 DPS) [dungeon] |

**New at 60:** head: Dreadmist Mask; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Wolfrunner Shoes; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 800.9. Weights run: 4.8s. Verify run: 26.7s. 1299 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.882 ± 0.048, crit=0.928 ± 0.039 per rating point (14 rating = 1%, 12.995 per %), hit=1.683 ± 0.115 per rating point (10 rating = 1%, 16.835 per %), spell_haste=5.932 ± 0.684, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.565 ± 0.002, spirit=not significant (0.242 ± 0.323), mp5=not significant (0.318 ± 0.365)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cap of the Scarlet Savant (12752, +0.00 DPS) [quest]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Arcane Crystal Pendant (20037, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 41.5 spell_power points (21.82 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.06 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.47 DPS) [pvp]; Shroud of the Nathrezim (18720, -7.50 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -8.61 DPS) [world]; Shroud of Arcane Mastery (22330, -9.45 DPS, sim-verified) [dungeon]; Spritecaster Cape (11623, -10.08 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.60 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -7.04 DPS) [pvp]; Alanna's Embrace (13314, -15.04 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 30.7 spell_power points (16.18 DPS) | yes | Runecloth Cuffs (254123, -5.75 DPS) [crafted]; Marshal's Silk Bracers (16438, -6.09 DPS) [pvp]; Dreadmist Bracers (16703, -7.81 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 43.8 spell_power points (23.05 DPS) | yes | Marshal's Silk Gloves (16440, -2.64 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.64 DPS) [vendor]; Hands of Power (13253, -6.81 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.2 spell_power points (21.14 DPS) | yes | Ban'thok Sash (11662, -3.24 DPS, sim-verified) [dungeon]; Clutch of Andros (13956, -3.46 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -4.14 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 66.8 spell_power points (35.16 DPS) | yes | Marshal's Silk Leggings (231605, -1.98 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -8.74 DPS) [pvp]; Skyshroud Leggings (13170, -13.55 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP reward · Alliance [vendor] | 32.8 spell_power points (17.25 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.15 DPS) [dungeon]; Dragonrider Boots (18102, -1.70 DPS, sim-verified) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.23 DPS) [vendor]; Maiden's Circle (13001, -2.89 DPS) [world_drop]; Naglering (11669, -22.05 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -2.21 DPS) [world_drop]; Songstone of Ironforge (12543, -2.47 DPS) [quest]; Blessed Band of Light (272407, -4.16 DPS, sim-verified) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+17.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -3.16 DPS) [world_drop]; Ankh of Life (1713, -4.79 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.60 DPS) [crafted]; Serenity Field (272439, -4.21 DPS) [vendor]; Uther's Strength (11302, -8.42 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -12.37 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -22.97 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (800.9 DPS) | yes | Bonecreeper Stylus (13938, -0.31 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.15 DPS) [dungeon]; Oblivion's Touch (18761, -2.85 DPS) [dungeon] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 46.7. Weights run: 3.3s. Verify run: 4.0s. 175 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.817 ± 0.013, crit=0.079 ± 0.002 per rating point (14 rating = 1%, 1.105 per %), hit=0.328 ± 0.013 per rating point (10 rating = 1%, 3.281 per %), spell_haste=not significant (0.117 ± 0.044), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.953 ± 0.003, spirit=not significant (0.158 ± 0.047), mp5=0.727 ± 0.085

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 10.5 spell_power points (2.11 DPS) | yes | Pristine Circlet (253949, -0.87 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.10 DPS) [crafted]; Flying Tiger Goggles (4368, -1.98 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.5 spell_power points (0.10 DPS) | yes | Scholarly Pendant (277203, -0.03 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.4 spell_power points (2.49 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.91 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.68 DPS) [crafted]; Slime-encrusted Pads (6461, -2.05 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.04 DPS) [crafted]; Battle Healer's Cloak (20427, -0.26 DPS) [rep]; Sanguine Cape (14376, -0.27 DPS) [world_drop] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 13.1 spell_power points (2.64 DPS) | yes | Filigreed Pristine Gown (253901, -0.70 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.02 DPS) [crafted]; Mystic's Wrap (14369, -1.49 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.3 spell_power points (1.26 DPS) | yes | Bright Bracers (3647, -0.53 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.66 DPS) [quest]; Owlbeard Bracers (16981, -0.73 DPS) [quest] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.1 spell_power points (1.84 DPS) | yes | Serpent Gloves (5970, -0.52 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.54 DPS) [crafted]; Gnoll Casting Gloves (892, -0.63 DPS) [world] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 9.9 spell_power points (2.00 DPS) | yes | Pristine Sash (253925, -0.52 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.69 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.70 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 16.0 spell_power points (3.22 DPS) | yes | Filigreed Pristine Leggings (253937, -0.87 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.82 DPS) [dungeon]; Darkweave Breeches (12987, -1.88 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.3 spell_power points (2.07 DPS) | yes | Pearly Boots (253897, -0.53 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.97 DPS) [crafted]; Red Woolen Boots (4313, -1.10 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.5 spell_power points (1.30 DPS) | yes | Loop of Sacrifice (281673, -0.48 DPS) [quest]; Volcanic Rock Ring (12053, -0.81 DPS) [world_drop]; Furlfeather Ring (286745, -1.10 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.9 spell_power points (0.99 DPS) | yes | Loop of Sacrifice (281673, -0.16 DPS) [quest]; Volcanic Rock Ring (12053, -0.49 DPS) [world_drop]; Furlfeather Ring (286745, -0.79 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Crescent Staff (6505, -0.12 DPS, sim-verified) [quest]; Staff of the Blessed Seer (2271, -0.17 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.57 DPS) [quest] |
| off_hand | Seer's Fine Stein (7608) | World drop [world_drop] | 2.6 spell_power points (0.53 DPS) | yes | Spellbinder Orb (15926, -0.03 DPS) [dungeon]; Tork Wrench (11855, -0.20 DPS) [quest]; Ancestral Orb (15944, -0.36 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sizzle Stick (8071, -1.02 DPS, sim-verified) [quest]; Cookie's Stirring Rod (5198, -1.08 DPS) [dungeon]; Greater Magic Wand (11288, -1.14 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Seer's Fine Stein; ranged: Skycaller

No-known-source sample (15 of 175, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance

### Band 30 (undead, 141005213100300000-00000000000000000-000000000000000000)

Set DPS (verified): 91.0. Weights run: 3.5s. Verify run: 5.7s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.010, crit=0.067 ± 0.002 per rating point (14 rating = 1%, 0.937 per %), hit=0.243 ± 0.012 per rating point (10 rating = 1%, 2.431 per %), spell_haste=0.856 ± 0.108, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.593 ± 0.003, spirit=not significant (0.217 ± 0.081), mp5=not significant (0.200 ± 0.116)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.26 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.60 DPS) [crafted]; Filigreed Pearly Circlet (253983, -0.76 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.78 DPS) | yes | Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.11 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.13 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.29 DPS) | yes | Death Speaker Mantle (6685, -0.44 DPS) [dungeon]; Magician's Mantle (12998, -0.94 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, -0.18 DPS) [vendor]; Hillman's Cloak (3719, -0.20 DPS) [crafted]; Sporid Cape (6629, -0.94 DPS, sim-verified) [dungeon] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 17.4 spell_power points (4.09 DPS) | yes | Green Silk Armor (7065, -0.28 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.64 DPS) [dungeon]; Death Speaker Robes (6682, -1.01 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.6 spell_power points (2.26 DPS) | yes | Nightsky Wristbands (6407, -1.32 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.41 DPS, sim-verified) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 9.2 spell_power points (2.17 DPS) | yes | Jutebraid Gloves (10654, -0.11 DPS) [quest]; Serpent Gloves (5970, -0.52 DPS) [dungeon]; Truefaith Gloves (7049, -0.60 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.7 spell_power points (2.98 DPS) | yes | Belt of Arugal (6392, -0.21 DPS) [dungeon]; Warsong Sash (16975, -0.39 DPS) [quest]; Crimson Silk Belt (7055, -0.65 DPS) [crafted] |
| legs | Pearly Leggings (253995) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.31 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.60 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.70 DPS, sim-verified) [dungeon] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 12.5 spell_power points (2.93 DPS) | yes | Gilded Slippers (254001, -0.17 DPS) [crafted]; Acidic Walkers (9454, -0.51 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (2.06 DPS) | yes | Ring of Power Regulation (273030, -0.51 DPS) [dungeon]; Solid Gold Band (279845, -0.87 DPS) [quest]; Electrocutioner Lagnut (9447, -1.00 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.4 spell_power points (1.74 DPS) | yes | Ring of Power Regulation (273030, -0.19 DPS) [dungeon]; Solid Gold Band (279845, -0.54 DPS) [quest]; Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -1.05 DPS) [crafted]; Tendonscraper Dagger (273053, -1.14 DPS) [world]; Glimmering Staff (249392, -2.87 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (2.43 DPS) | yes | Witch's Finger (16887, -1.38 DPS) [quest]; Satyr's Rod (15962, -1.54 DPS) [world_drop]; Orb of Souls (249395, -1.83 DPS, sim-verified) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.20 DPS) [world_drop]; Greater Mystic Wand (217287, -0.35 DPS) [crafted]; Wand of Eventide (5214, -0.49 DPS) [world_drop] |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Pearly Leggings; feet: Golden Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 173.0. Weights run: 4.0s. Verify run: 4.2s. 392 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.571 ± 0.022, crit=0.267 ± 0.006 per rating point (14 rating = 1%, 3.740 per %), hit=0.605 ± 0.034 per rating point (10 rating = 1%, 6.055 per %), spell_haste=2.320 ± 0.360, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.626 ± 0.002, spirit=not significant (0.676 ± 0.220), mp5=not significant (0.868 ± 0.248)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.8 spell_power points (7.83 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.92 DPS) [crafted]; Whitemane's Chapeau (7720, -0.98 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.69 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 12.7 spell_power points (4.19 DPS) | yes | Glowing Eye of Mordresh (10769, -0.38 DPS) [dungeon]; Necklace of Calisea (1714, -1.32 DPS) [dungeon]; Amberglow Talisman (10824, -1.97 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Mantle of Doan (7712, -0.71 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19528, -0.55 DPS) [rep]; Darkspear Raider's Cloak (272077, -0.68 DPS) [vendor]; Long Silken Cloak (4326, -0.72 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 32.6 spell_power points (10.73 DPS) | yes | Robe of the Magi (1716, -1.25 DPS) [world_drop]; Robe of Power (7054, -2.09 DPS) [crafted]; Beguiler Robes (7728, -3.73 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.6 spell_power points (3.82 DPS) | yes | Radiant Silver Bracers (4545, -1.23 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.39 DPS) [crafted]; Condor Bracers (15864, -1.52 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 25.0 spell_power points (8.23 DPS) | yes | Red Mageweave Gloves (10018, -2.00 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -3.30 DPS) [crafted]; Golden Handwraps (254029, -4.03 DPS) [crafted] |
| waist | Golden Cord (254045) | Tailoring [crafted] | 17.4 spell_power points (5.72 DPS) | yes | Gilded Cord (254037, -0.25 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.36 DPS) [rep]; Star Belt (4329, -0.56 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.9 spell_power points (6.87 DPS) | yes | Gaze Dreamer Pants (6903, -0.91 DPS) [dungeon]; Black Mageweave Leggings (9999, -1.12 DPS) [crafted]; Stoneweaver Leggings (9407, -1.85 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.7 spell_power points (9.46 DPS) | yes | Furen's Boots (13100, -4.58 DPS) [world_drop]; Thoughtcast Boots (10578, -4.92 DPS) [dungeon]; Golden Slippers (254009, -5.10 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (4.23 DPS) | yes | Snake Hoop (6750, -1.32 DPS) [quest]; Ring of Power Regulation (273030, -1.40 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.69 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 11.6 spell_power points (3.82 DPS) | yes | Snake Hoop (6750, -0.90 DPS) [quest]; Ring of Power Regulation (273030, -0.99 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.28 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -12.07 DPS) [dungeon]; Celestial Stave (9517, -12.08 DPS) [quest]; Illusionary Rod (7713, -12.15 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Twisted Nether Wand (249144, -0.53 DPS) [crafted]; Starfaller (13063, -0.55 DPS) [world_drop]; Dancing Flame (6806, -1.15 DPS, sim-verified) [quest] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Golden Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 141005213100301531-40000000000000000-051000000000000000)

Set DPS (verified): 256.0. Weights run: 3.9s. Verify run: 6.2s. 504 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.931 ± 0.035, crit=0.241 ± 0.006 per rating point (14 rating = 1%, 3.378 per %), hit=0.637 ± 0.039 per rating point (10 rating = 1%, 6.366 per %), spell_haste=not significant (1.423 ± 0.511), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.619 ± 0.003, spirit=not significant (0.959 ± 0.413), mp5=not significant (0.863 ± 0.398)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 41.8 spell_power points (10.41 DPS) | yes | Red Mageweave Headband (10033, -1.05 DPS) [crafted]; Gemburst Circlet (10751, -1.18 DPS) [quest]; Soulcatcher Halo (10630, -2.23 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.6 spell_power points (5.37 DPS) | yes | Mindburst Medallion (11196, -1.30 DPS) [quest]; Gemshard Heart (17707, -1.62 DPS) [dungeon]; Horizon Choker (13085, -3.74 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (256.0 DPS) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Nethershine Shoulders (254057, -1.50 DPS) [crafted]; Nethergeld Shoulders (254049, -1.88 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.58 DPS, sim-verified) [quest]; Featherskin Cape (10843, -1.80 DPS) [world]; Darkspear Raider's Cloak (272076, -1.87 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Robes of Insight (940, -0.53 DPS) [world_drop]; Acumen Robes (17775, -0.54 DPS) [quest]; Embrace of the Wind Serpent (12462, -1.12 DPS, sim-verified) [world] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Nethergeld Cuffs (254061, -0.26 DPS) [crafted]; Bloodband Bracers (11469, -0.77 DPS) [quest]; Aristocratic Cuffs (12546, -1.41 DPS, sim-verified) [dungeon] |
| hands | Runecloth Gloves (13863) | Tailoring [crafted] | 29.0 spell_power points (7.22 DPS) | yes | Golden Gloves (254103, -0.06 DPS) [crafted]; Dreamweave Gloves (10019, -0.14 DPS) [crafted]; Gilded Gloves (254095, -0.50 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 28.6 spell_power points (7.12 DPS) | yes | Dawnspire Cord (12466, -1.22 DPS) [dungeon]; Satyrmane Sash (17755, -1.32 DPS) [dungeon]; Ghostweave Cord (254073, -4.55 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 37.0 spell_power points (9.22 DPS) | yes | Senior Designer's Pantaloons (11841, -0.31 DPS) [dungeon]; Dalewind Trousers (13008, -1.16 DPS) [world_drop]; Spellshock Leggings (9484, -1.17 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.7 spell_power points (7.65 DPS) | yes | Vinerot Sandals (17748, -1.68 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -1.98 DPS) [vendor]; Mistwalker Boots (10629, -5.68 DPS, sim-verified) [dungeon] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.4 spell_power points (4.82 DPS) | yes | Eye of Adaegus (5266, -0.56 DPS) [world_drop]; Advisor's Ring (19519, -0.97 DPS) [rep]; Philanthropist's Ring (281635, -1.17 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 18.8 spell_power points (4.67 DPS) | yes | Eye of Adaegus (5266, -0.41 DPS) [world_drop]; Advisor's Ring (19519, -0.82 DPS) [rep]; Philanthropist's Ring (281635, -1.02 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.32 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.38 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.67 DPS, sim-verified) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -3.02 DPS) [world_drop]; Spellshifter Rod (9527, -5.56 DPS) [quest]; Spire of Hakkar (10844, -6.20 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Ledger from Tanaris [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, +0.00 DPS) [world_drop]; Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Nether Force Wand (11263, -0.09 DPS) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Nethershine Cuffs; hands: Runecloth Gloves; waist: Ban'thok Sash; legs: Kilt of the Atal'ai Prophet; finger1: Cyclopean Band; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Charged Lightning Rod

No-known-source sample (15 of 504, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 141005213100301531-40000000000000000-053500030000000000)

Set DPS (verified): 408.9. Weights run: 4.0s. Verify run: 25.6s. 1293 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.815 ± 0.043, crit=0.579 ± 0.014 per rating point (14 rating = 1%, 8.103 per %), hit=1.158 ± 0.067 per rating point (10 rating = 1%, 11.578 per %), spell_haste=4.854 ± 0.869, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.608 ± 0.002, spirit=not significant (0.803 ± 0.375), mp5=not significant (0.907 ± 0.395)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreadmist Mask (16698) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.7 spell_power points (14.08 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -0.16 DPS) [pvp]; Shroud of the Nathrezim (18720, -0.45 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.27 DPS) [world]; Shroud of Arcane Mastery (22330, -5.35 DPS) [dungeon]; Spritecaster Cape (11623, -7.54 DPS, sim-verified) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.30 DPS) [pvp]; Alanna's Embrace (13314, -7.51 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 34.1 spell_power points (13.46 DPS) | yes | General's Silk Cuffs (16538, -4.19 DPS) [pvp]; Wyrmthalak's Shackles (13958, -5.82 DPS) [quest]; Runecloth Cuffs (254123, -5.91 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; General's Silk Handwraps (231600, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 38.6 spell_power points (15.20 DPS) | yes | General's Silk Sash (16537, -3.06 DPS) [pvp]; Wisdom of the Timbermaw (19047, -3.52 DPS) [crafted]; Ban'thok Sash (11662, -6.30 DPS, sim-verified) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 51.5 spell_power points (20.31 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -1.79 DPS) [pvp]; Magister's Leggings (16687, -3.96 DPS, sim-verified) [dungeon] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | sim-verified (408.9 DPS) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Mooncloth Boots (15802, -0.37 DPS) [crafted] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.56 DPS) [vendor]; Eye of Orgrimmar (12545, -0.98 DPS) [quest]; Naglering (11669, -15.32 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.63 DPS) [quest]; Deep Rooted Ring (19109, -3.30 DPS) [world]; Blessed Band of Light (272407, -3.88 DPS, sim-verified) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.62 DPS) [crafted]; Ankh of Life (1713, -0.93 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.54 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -3.15 DPS) [vendor]; Frozen Heart of the Mountain (249469, -3.91 DPS, sim-verified) [crafted]; Ankh of Life (1713, -4.88 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -3.46 DPS) [world_drop]; Ironbark Staff (20220, -9.87 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (11.44 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -1.02 DPS) [pvp]; Tome of the Lost (22253, -2.09 DPS) [dungeon] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torch of Light (279246, +0.00 DPS) [crafted]; Bonecreeper Stylus (13938, -0.01 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.23 DPS) [dungeon] |

**New at 60:** head: Dreadmist Mask; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Wolfrunner Shoes; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 787.2. Weights run: 4.8s. Verify run: 27.6s. 1293 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.846 ± 0.044, crit=0.579 ± 0.026 per rating point (14 rating = 1%, 8.110 per %), hit=1.186 ± 0.075 per rating point (10 rating = 1%, 11.857 per %), spell_haste=3.133 ± 0.631, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.567 ± 0.002, spirit=not significant (0.542 ± 0.476), mp5=not significant (1.276 ± 0.493)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreadmist Mask (16698) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 27.6 spell_power points (12.16 DPS) | yes | Star of Mystaria (12103, -1.68 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.70 DPS) [world_drop]; Arcane Crystal Pendant (20037, -2.88 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.1 spell_power points (15.90 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.56 DPS) [vendor]; Champion's Silk Mantle (227104, -0.67 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.95 DPS) [world]; Spritecaster Cape (11623, -5.41 DPS) [dungeon]; Cloak of Flames (3475, -12.80 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -5.31 DPS) [pvp]; Alanna's Embrace (13314, -11.00 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 32.6 spell_power points (14.34 DPS) | yes | General's Silk Cuffs (16538, -4.87 DPS) [pvp]; Runecloth Cuffs (254123, -5.77 DPS) [crafted]; Dreadmist Bracers (16703, -6.85 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Hands of Power (13253, -5.09 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 39.3 spell_power points (17.30 DPS) | yes | Ban'thok Sash (11662, -2.70 DPS) [dungeon]; Clutch of Andros (13956, -5.01 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -5.10 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 49.5 spell_power points (21.79 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -1.45 DPS) [pvp]; Magister's Leggings (16687, -5.71 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (787.2 DPS) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.30 DPS) [pvp] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.81 DPS) [vendor]; Eye of Orgrimmar (12545, -0.86 DPS) [quest]; Naglering (11669, -22.01 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.43 DPS) [vendor]; Eye of Orgrimmar (12545, -0.48 DPS) [quest]; Deep Rooted Ring (19109, -2.96 DPS) [world] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+16.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Rune of the Guard Captain (19120, -1.63 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Rune of the Guard Captain (19120, -1.04 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -0.10 DPS) [pvp]; Whiteout Staff (19101, -12.93 DPS) [rep]; Staff of Jordan (873, -14.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.21 DPS) [dungeon]; Torch of Light (279246, -1.15 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -1.69 DPS) [dungeon] |

**New at 60:** head: Dreadmist Mask; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Burst of Knowledge; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

