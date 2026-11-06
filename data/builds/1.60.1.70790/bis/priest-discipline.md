# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 45.9. Weights run: 13.0s. Verify run: 64.0s. 186 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.034, intellect=0.785 ± 0.006, spirit=0.439 ± 0.007, mp5=2.230 ± 0.016, crit=0.051 ± 0.003 per rating point (14 rating = 1%, 0.721 per %), spell_haste=not significant (-0.089 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.00 DPS) | yes | Shadow Goggles (4373, -0.83 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.81 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.9 healing_power points (0.10 DPS) | yes | Tarnished Locket (279870, -0.28 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.1 healing_power points (1.34 DPS) | yes | Slime-encrusted Pads (6461, -0.60 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.68 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (45.9 DPS) | yes | Heavy Woolen Cloak (4311, -0.41 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.61 DPS) [crafted]; Sanguine Cape (14376, -0.75 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 18.2 healing_power points (2.03 DPS) | yes | Corsair's Overshirt (5202, -0.26 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.54 DPS) [dungeon]; Bloody Apron (6226, -0.58 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.6 healing_power points (1.18 DPS) | yes | Repurposed Hair Band (281256, -0.91 DPS) [quest]; Seer's Cuffs (3645, -1.00 DPS) [dungeon]; Bright Bracers (3647, -1.43 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.4 healing_power points (1.48 DPS) | yes | Gnoll Casting Gloves (892, -0.82 DPS) [world]; Magefist Gloves (12977, -0.85 DPS) [world_drop]; Serpent Gloves (5970, -0.87 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.1 healing_power points (1.57 DPS) | yes | Novice Ardent's Sash (253887, -0.40 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.87 DPS) [crafted]; Keller's Girdle (2911, -0.87 DPS) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 24.5 healing_power points (2.72 DPS) | yes | Darkweave Breeches (12987, -1.80 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -1.92 DPS) [crafted]; Silk-threaded Trousers (1929, -1.94 DPS) [dungeon] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.4 healing_power points (1.26 DPS) | yes | Spidersilk Boots (4320, -0.14 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.57 DPS) [crafted]; Kimbra Boots (6191, -0.77 DPS) [quest] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.5 healing_power points (1.05 DPS) | yes | Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Band of Purification (12996, -0.76 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.79 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.6 healing_power points (0.73 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS, sim-verified) [dungeon]; Band of Purification (12996, -0.44 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.47 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 54.6 healing_power points (6.06 DPS) | yes | Living Root (6631, -1.14 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -2.46 DPS) [dungeon]; Evocator's Blade (2567, -2.53 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.22 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Flaring Baton (5326, -0.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall; ranged: Sizzle Stick

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 89.8. Weights run: 15.0s. Verify run: 72.6s. 308 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.026, intellect=0.837 ± 0.008, spirit=1.139 ± 0.009, mp5=2.561 ± 0.009, crit=0.086 ± 0.003 per rating point (14 rating = 1%, 1.201 per %), spell_haste=not significant (0.294 ± 0.080)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.8 healing_power points (5.22 DPS) | yes | Pristine Circlet (253949, -2.86 DPS) [crafted]; Filigreed Pristine Circlet (253975, -3.13 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -3.16 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.9 healing_power points (1.04 DPS) | yes | Glowing Green Talisman (5002, -0.14 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.15 DPS) [world_drop]; Pendant of Myzrael (4614, -0.35 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.5 healing_power points (2.17 DPS) | yes | Death Speaker Mantle (6685, -0.12 DPS, sim-verified) [dungeon]; Mantle of Honor (3560, -0.35 DPS) [quest]; Magician's Mantle (12998, -0.52 DPS) [world_drop] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-verified (89.8 DPS) | yes | Glowing Thresher Cape (6901, -0.12 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.44 DPS) [quest]; Prelacy Cape (7004, -0.94 DPS) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 30.4 healing_power points (3.99 DPS) | yes | Death Speaker Robes (6682, -0.14 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.30 DPS) [dungeon]; Robes of Arugal (6324, -0.73 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.7 healing_power points (2.19 DPS) | yes | Mindthrust Bracers (1974, -0.51 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.31 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 17.9 healing_power points (2.34 DPS) | yes | Pristine Gloves (253913, -0.57 DPS) [crafted]; Gloves of Meditation (4318, -0.64 DPS) [crafted]; Truefaith Gloves (7049, -0.78 DPS, sim-verified) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 17.2 healing_power points (2.25 DPS) | yes | Highlander's Cloth Girdle (20099, -0.48 DPS) [rep]; Pristine Sash (253925, -0.66 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.69 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 33.6 healing_power points (4.40 DPS) | yes | Earthen Leggings (253999, -0.79 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.09 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.48 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.4 healing_power points (3.99 DPS) | yes | Boots of the Enchanter (4325, -2.14 DPS) [crafted]; Glinteye Slippers (273024, -2.17 DPS) [dungeon]; Acidic Walkers (9454, -2.36 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.0 healing_power points (2.22 DPS) | yes | Ring of Power Regulation (273030, -0.43 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.64 DPS) [rep]; The Queen's Jewel (13094, -0.81 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 17.0 healing_power points (2.22 DPS) | yes | Ring of Power Regulation (273030, -0.37 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19525, -0.64 DPS) [rep]; The Queen's Jewel (13094, -0.81 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (89.8 DPS) | yes | Odo's Ley Staff (6318, -0.31 DPS) [dungeon]; Staff of Westfall (2042, -1.18 DPS) [quest]; Staff of the Friar (3415, -2.94 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 19.8 healing_power points (2.60 DPS) | yes | Eye of Paleth (2943, -0.90 DPS) [quest]; Orb of Mystic Insight (249394, -1.02 DPS) [crafted]; Orb of Mistmantle (13031, -1.14 DPS, sim-verified) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.5 healing_power points (1.11 DPS) | yes | Wand of Eventide (5214, -0.45 DPS) [world_drop]; Greater Mystic Wand (217287, -0.45 DPS) [crafted]; Thunderwood (13062, -1.01 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 005005031330011430-20000000000000000-00000000000000000)

Set DPS (verified): 131.6. Weights run: 25.2s. Verify run: 124.8s. 402 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.158, intellect=1.445 ± 0.032, spirit=1.282 ± 0.031, mp5=3.723 ± 0.035, crit=0.402 ± 0.023 per rating point (14 rating = 1%, 5.633 per %), spell_haste=not significant (2.210 ± 0.573)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 58.1 healing_power points (7.13 DPS) | yes | Whitemane's Chapeau (7720, -0.97 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.14 DPS) [world_drop]; Miner's Hat of the Deep (9429, -2.54 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.6 healing_power points (3.87 DPS) | yes | Necklace of Calisea (1714, -1.76 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -1.79 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.10 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.8 healing_power points (5.01 DPS) | yes | Mantle of Doan (7712, -1.82 DPS) [dungeon]; Green Silken Shoulders (7057, -2.07 DPS) [crafted]; Earthen Silk Shoulders (254033, -2.41 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.41 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -0.88 DPS) [dungeon]; Repairman's Cape (9605, -0.94 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 68.3 healing_power points (8.38 DPS) | yes | Dreamweave Vest (10021, -2.79 DPS, sim-verified) [crafted]; Robe of Power (7054, -3.27 DPS) [crafted]; Civinad Robes (9623, -3.30 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.2 healing_power points (2.48 DPS) | yes | Mindthrust Bracers (1974, -0.31 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.46 DPS) [crafted]; Windchaser Cuffs (14429, -0.88 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 32.8 healing_power points (4.02 DPS) | yes | Stormcloth Gloves (10011, -0.54 DPS) [crafted]; Earthen Silk Gloves (254017, -0.55 DPS) [crafted]; Gilded Handwraps (254021, -0.66 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.3 healing_power points (5.06 DPS) | yes | Deathmage Sash (10771, -2.08 DPS, sim-verified) [dungeon]; Windchaser Cinch (14435, -2.23 DPS) [world_drop]; Sutarn's Ring (13105, -2.35 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 46.9 healing_power points (5.76 DPS) | yes | Pristine Leggings (253987, -1.49 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -1.56 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.86 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 35.2 healing_power points (4.33 DPS) | yes | Thoughtcast Boots (10578, -0.68 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.02 DPS) [crafted]; Furen's Boots (13100, -1.56 DPS, sim-verified) [world_drop] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 20.2 healing_power points (2.48 DPS) | yes | Runemender's Seal (273031, -0.27 DPS) [world]; Philanthropist's Ring (281635, -0.36 DPS) [quest]; Ring of Power Regulation (273030, -0.37 DPS) [dungeon] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.0 healing_power points (2.21 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Philanthropist's Ring (281635, -0.09 DPS) [quest]; Ring of Power Regulation (273030, -0.10 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -3.86 DPS, sim-verified) [world_drop]; Staff of the Friar (3415, -3.99 DPS) [dungeon]; Death Speaker Scepter (2816, -4.45 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 27.1 healing_power points (3.33 DPS) | yes | Orb of Lorica (11262, -0.31 DPS, sim-verified) [quest]; Orb of Souls (249395, -0.79 DPS) [crafted]; Orb of Mystic Insight (249394, -1.41 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-verified (131.6 DPS) | yes | Wand of Mana Concentration (273043, -0.09 DPS, sim-verified) [world]; Captain Rackmore's Tiller (16789, -0.33 DPS) [quest]; Goblin Igniter (5253, -0.35 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Lorekeeper's Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 190.6. Weights run: 26.5s. Verify run: 246.4s. 517 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.146, intellect=1.253 ± 0.033, spirit=0.883 ± 0.030, mp5=2.839 ± 0.033, crit=0.514 ± 0.025 per rating point (14 rating = 1%, 7.199 per %), spell_haste=not significant (1.577 ± 0.671)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Gemburst Circlet (10751, -0.70 DPS) [quest]; Cassandra's Grace (13102, -0.92 DPS, sim-verified) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.44 DPS) [quest]; Gemshard Heart (17707, -0.45 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.50 DPS, sim-verified) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 41.3 healing_power points (5.75 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Kentic Amice (11624, -0.80 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.87 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (19532, -0.61 DPS) [rep]; Expeditionary's Cape (271921, -0.68 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.1 healing_power points (9.06 DPS) | yes | Knight's Satin Armor (220892, -0.93 DPS, sim-verified) [vendor]; Ghostweave Vest (14141, -2.18 DPS) [crafted]; Embrace of the Wind Serpent (12462, -2.40 DPS) [world] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, -1.27 DPS) [dungeon]; Bloodband Bracers (11469, -1.99 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.16 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 48.6 healing_power points (6.77 DPS) | yes | Earthenweave Gloves (254075, -1.79 DPS) [crafted]; Runecloth Gloves (13863, -2.42 DPS) [crafted]; Sergeant Major's Satin Gloves (220897, -3.21 DPS, sim-verified) [vendor] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Dawnspire Cord (12466, -0.21 DPS) [dungeon]; Earthenweave Cord (254077, -0.22 DPS) [crafted]; Gilded Waistcord (254081, -0.92 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Kilt of the Atal'ai Prophet (10807, -0.73 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.26 DPS) [dungeon]; Knight's Satin Leggings (220893, -2.18 DPS, sim-verified) [vendor] |
| feet | Earthenweave Boots (254093) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Sergeant Major's Satin Boots (220895, -0.09 DPS) [vendor]; Gilded Slippers (254001, -0.47 DPS) [crafted]; Gilded Sandals (254107, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 23.4 healing_power points (3.25 DPS) | yes | Cyclopean Band (11824, -0.29 DPS) [dungeon]; Philanthropist's Ring (281635, -0.99 DPS) [quest]; Brainlash (6440, -3.94 DPS, sim-verified) [dungeon] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-verified (190.5 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.26 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.25 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.49 DPS) [quest] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.37 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -0.99 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -5.78 DPS, sim-verified) [quest]; Charstone Dirk (17710, -8.19 DPS) [dungeon]; Dreamstaff (249454, -10.10 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Jaina's Firestarter (13064, -0.07 DPS) [world_drop]; Charged Lightning Rod (11860, -0.09 DPS) [quest]; Wand of Mana Concentration (273043, -0.61 DPS, sim-verified) [world] |

**New at 50:** neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Spritecaster Cape; wrist: Nethergeld Cuffs; hands: Gilded Gloves; legs: Senior Designer's Pantaloons; feet: Earthenweave Boots; finger1: Lorekeeper's Ring; finger2: Eye of Adaegus; trinket2: Enriched Seal; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 517, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 288.1. Weights run: 27.5s. Verify run: 728.8s. 1359 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.146, intellect=1.601 ± 0.069, spirit=3.066 ± 0.063, mp5=6.371 ± 0.067, crit=0.392 ± 0.026 per rating point (14 rating = 1%, 5.491 per %), spell_haste=not significant (-2.755 ± 1.073)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 104.4 healing_power points (12.00 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -5.20 DPS, sim-verified) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 61.1 healing_power points (7.02 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.37 DPS, sim-verified) [world_drop]; Heart of the Fiend (13960, -0.81 DPS) [dungeon]; Animated Chain Necklace (18723, -1.11 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -4.91 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, -0.47 DPS) [world]; Caretaker's Cape (19530, -0.68 DPS) [rep]; Frostweaver Cape (12968, -2.42 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 122.1 healing_power points (14.02 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Embrace of the Wind Serpent (12462, -0.33 DPS) [world]; Ghostweave Vest (14141, -3.61 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, -2.87 DPS, sim-verified) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.67 DPS) [crafted]; Magister's Gloves (16684, -1.24 DPS) [dungeon]; Hands of the Exalted Herald (12554, -2.38 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Penitent's Cinch (272394, +0.00 DPS) [vendor]; Wisdom of the Timbermaw (19047, -3.06 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 118.5 healing_power points (13.62 DPS) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -1.33 DPS) [pvp]; Haunting Specter Leggings (11929, -7.89 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 86.3 healing_power points (9.91 DPS) | yes | Faith Healer's Boots (22247, -0.30 DPS, sim-verified) [dungeon]; Devout Sandals (16691, -0.82 DPS) [dungeon]; Marshal's Satin Walkers (231627, -1.55 DPS) [vendor] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (288.0 DPS) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, +0.00 DPS) [dungeon]; Band of the Hierophant (13096, -0.67 DPS) [world_drop] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Band of the Hierophant (13096, -1.25 DPS) [world_drop]; The Postmaster's Seal (13392, -3.25 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, -0.57 DPS) [rep]; Evonice's Landin' Pilla (18951, -0.70 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.43 DPS) [quest]; Enriched Seal (271903, -1.17 DPS) [rep]; Second Wind (11819, -3.16 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -4.03 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -4.82 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 50.7 healing_power points (5.83 DPS) | yes | Cairnstone Sliver (9654, -2.60 DPS, sim-verified) [quest]; Brilliant Wand (249385, -3.51 DPS) [crafted]; Wand of Mana Concentration (273043, -3.63 DPS) [world] |

**New at 60:** head: Devout Crown; neck: Lady Maye's Pendant; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Band of Mending; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 596.1. Weights run: 18.7s. Verify run: 526.5s. 1359 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.497, intellect=1.734 ± 0.090, spirit=1.606 ± 0.088, mp5=4.128 ± 0.101, crit=1.114 ± 0.078 per rating point (14 rating = 1%, 15.591 per %), spell_haste=not significant (1.384 ± 2.027)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 85.7 healing_power points (16.38 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Crown of the Penitent (13216, -14.49 DPS, sim-verified) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -13.37 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -1.66 DPS) [rep]; Expeditionary's Cape (271907, -1.94 DPS) [rep]; Frostweaver Cape (12968, -3.13 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 98.7 healing_power points (18.86 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Field Marshal's Satin Vestments (17605, -0.12 DPS) [vendor]; Robes of the Exalted (13346, -1.26 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.15 DPS) [crafted]; Dryad's Wrist Bindings (19595, -11.34 DPS, sim-verified) [rep] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Marshal's Satin Handwraps (231623, -0.82 DPS) [vendor]; Devout Gloves (16692, -2.14 DPS) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Penitent's Cinch (272394, +0.00 DPS) [vendor]; Wisdom of the Timbermaw (19047, -10.56 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 86.9 healing_power points (16.62 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -10.69 DPS, sim-verified) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 69.1 healing_power points (13.22 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -0.59 DPS) [dungeon]; Devout Sandals (16691, -2.58 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-verified (596.0 DPS) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+6.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Second Wind (11819, +0.00 DPS) [dungeon]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -3.29 DPS) [quest]; Weakness Analyzer (272438, -3.82 DPS) [vendor]; Second Wind (11819, -5.22 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -11.98 DPS) [pvp]; Ironbark Staff (20069, -13.29 DPS) [rep]; Resurgence Rod (17743, -14.15 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 34.7 healing_power points (6.63 DPS) | yes | Oblivion's Touch (18761, -2.98 DPS) [dungeon]; Bonecreeper Stylus (13938, -3.20 DPS) [dungeon]; Brilliant Wand (249385, -3.72 DPS, sim-verified) [crafted] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Theramore Handwraps; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 45.4. Weights run: 13.0s. Verify run: 67.6s. 179 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.159, intellect=1.183 ± 0.014, spirit=1.108 ± 0.014, mp5=3.213 ± 0.015, crit=0.117 ± 0.009 per rating point (14 rating = 1%, 1.641 per %), spell_haste=not significant (0.179 ± 0.183)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.93 DPS) | yes | Shadow Goggles (4373, -0.58 DPS) [crafted]; Flying Tiger Goggles (4368, -1.46 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | sim-verified (45.4 DPS) | yes | Scholarly Pendant (277203, -0.12 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.6 healing_power points (1.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.64 DPS) [crafted]; Slime-encrusted Pads (6461, -0.65 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.25 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.30 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.61 DPS) [crafted]; Sanguine Cape (14376, -0.70 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 22.2 healing_power points (2.39 DPS) | yes | Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Corsair's Overshirt (5202, -0.26 DPS) [dungeon]; Scholarly Robes (2034, -0.64 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 15.6 healing_power points (1.67 DPS) | yes | Bright Bracers (3647, -1.16 DPS) [world_drop]; Featherbead Bracers (15452, -1.29 DPS) [quest]; Seer's Cuffs (3645, -1.31 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 14.5 healing_power points (1.56 DPS) | yes | Magefist Gloves (12977, -0.45 DPS) [world_drop]; Bright Gloves (3066, -0.58 DPS) [world_drop]; Serpent Gloves (5970, -0.81 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.7 healing_power points (1.69 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.67 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.88 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 29.5 healing_power points (3.17 DPS) | yes | Darkweave Breeches (12987, -1.57 DPS) [world_drop]; Filigreed Silky Leggings (253939, -1.93 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.93 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 12.5 healing_power points (1.35 DPS) | yes | Spidersilk Boots (4320, -0.09 DPS) [crafted]; Red Woolen Boots (4313, -0.32 DPS) [crafted]; Bluegill Sandals (1560, -0.62 DPS) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 11.4 healing_power points (1.23 DPS) | yes | Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon]; Loop of Sacrifice (281673, -0.59 DPS) [quest]; Volcanic Rock Ring (12053, -0.85 DPS) [world_drop] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 50.5 healing_power points (5.43 DPS) | yes | Staff of the Blessed Seer (2271, -1.23 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -1.45 DPS) [quest]; Evocator's Blade (2567, -1.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-verified (45.4 DPS) | yes | Wand of Decay (5252, +0.00 DPS) [quest]; Flaring Baton (5326, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Living Root; ranged: Sizzle Stick

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 94.6. Weights run: 15.0s. Verify run: 74.4s. 301 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.026, intellect=0.837 ± 0.008, spirit=1.139 ± 0.009, mp5=2.561 ± 0.009, crit=0.086 ± 0.003 per rating point (14 rating = 1%, 1.201 per %), spell_haste=not significant (0.294 ± 0.080)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.8 healing_power points (5.22 DPS) | yes | Pristine Circlet (253949, -2.86 DPS) [crafted]; Filigreed Pristine Circlet (253975, -3.06 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -3.16 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.9 healing_power points (1.04 DPS) | yes | Glowing Green Talisman (5002, -0.14 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.15 DPS) [world_drop]; Pendant of Myzrael (4614, -0.24 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.5 healing_power points (2.17 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Ghostly Mantle (3324, -0.48 DPS) [quest]; Magician's Mantle (12998, -0.52 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-verified (94.6 DPS) | yes | Glowing Thresher Cape (6901, -0.19 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.98 DPS) [vendor]; Windsong Drape (15468, -1.05 DPS) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 30.4 healing_power points (3.99 DPS) | yes | Death Speaker Robes (6682, -0.11 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.30 DPS) [dungeon]; Robes of Arugal (6324, -0.73 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.7 healing_power points (2.19 DPS) | yes | Mindthrust Bracers (1974, -0.55 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.31 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 17.9 healing_power points (2.34 DPS) | yes | Truefaith Gloves (7049, -0.55 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.57 DPS) [crafted]; Gloves of Meditation (4318, -0.64 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 17.2 healing_power points (2.25 DPS) | yes | Pristine Sash (253925, -0.46 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.48 DPS) [rep]; Novice Ardent's Sash (253887, -0.69 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 33.6 healing_power points (4.40 DPS) | yes | Earthen Leggings (253999, -0.79 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.97 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.48 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.4 healing_power points (3.99 DPS) | yes | Boots of the Enchanter (4325, -2.14 DPS) [crafted]; Glinteye Slippers (273024, -2.17 DPS) [dungeon]; Acidic Walkers (9454, -2.26 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.0 healing_power points (2.22 DPS) | yes | Solid Gold Band (279845, -0.04 DPS) [quest]; Ring of Power Regulation (273030, -0.43 DPS) [dungeon]; Advisor's Ring (19521, -0.64 DPS) [rep] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 17.0 healing_power points (2.22 DPS) | yes | Solid Gold Band (279845, -0.08 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -0.43 DPS) [dungeon]; Advisor's Ring (19521, -0.64 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (94.6 DPS) | yes | Odo's Ley Staff (6318, -0.31 DPS) [dungeon]; Living Root (6631, -2.27 DPS) [dungeon]; Staff of the Friar (3415, -2.98 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 19.8 healing_power points (2.60 DPS) | yes | Orb of Mystic Insight (249394, -1.02 DPS) [crafted]; Orb of Mistmantle (13031, -1.03 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -1.64 DPS) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.5 healing_power points (1.11 DPS) | yes | Wand of Eventide (5214, -0.45 DPS) [world_drop]; Greater Mystic Wand (217287, -0.45 DPS) [crafted]; Thunderwood (13062, -0.75 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 005005031330011430-20000000000000000-00000000000000000)

Set DPS (verified): 140.7. Weights run: 25.2s. Verify run: 128.6s. 395 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.158, intellect=1.445 ± 0.032, spirit=1.282 ± 0.031, mp5=3.723 ± 0.035, crit=0.402 ± 0.023 per rating point (14 rating = 1%, 5.633 per %), spell_haste=not significant (2.210 ± 0.573)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 58.1 healing_power points (7.13 DPS) | yes | Whitemane's Chapeau (7720, -0.20 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.14 DPS) [world_drop]; Miner's Hat of the Deep (9429, -2.54 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.6 healing_power points (3.87 DPS) | yes | Necklace of Calisea (1714, -0.90 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -1.79 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.10 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.8 healing_power points (5.01 DPS) | yes | Earthen Silk Shoulders (254033, -1.07 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -1.82 DPS) [dungeon]; Green Silken Shoulders (7057, -2.07 DPS) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.42 DPS) [vendor]; Glowing Thresher Cape (6901, -0.88 DPS) [dungeon]; Blackforge Cape (6424, -0.95 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 68.3 healing_power points (8.38 DPS) | yes | Dreamweave Vest (10021, -2.09 DPS, sim-verified) [crafted]; Robe of Power (7054, -3.27 DPS) [crafted]; Civinad Robes (9623, -3.30 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.2 healing_power points (2.48 DPS) | yes | Earthen Silk Cuffs (254019, -0.46 DPS) [crafted]; Mindthrust Bracers (1974, -0.53 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -0.57 DPS) [quest] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Gloves (10011, -0.46 DPS) [crafted]; Earthen Silk Gloves (254017, -0.47 DPS) [crafted]; Dreamweave Gloves (10019, -0.73 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.3 healing_power points (5.06 DPS) | yes | Deathmage Sash (10771, -1.17 DPS, sim-verified) [dungeon]; Windchaser Cinch (14435, -2.23 DPS) [world_drop]; Sutarn's Ring (13105, -2.35 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 46.9 healing_power points (5.76 DPS) | yes | Pristine Leggings (253987, -0.42 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -1.56 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.86 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 35.2 healing_power points (4.33 DPS) | yes | Furen's Boots (13100, -0.57 DPS) [world_drop]; Thoughtcast Boots (10578, -0.68 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.02 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 22.3 healing_power points (2.74 DPS) | yes | Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Runemender's Seal (273031, -0.54 DPS) [world]; Solid Gold Band (279845, -0.60 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 20.2 healing_power points (2.48 DPS) | yes | Electrocutioner Lagnut (9447, -0.27 DPS) [dungeon]; Runemender's Seal (273031, -0.27 DPS) [world]; Solid Gold Band (279845, -0.34 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -3.75 DPS, sim-verified) [world_drop]; Staff of the Friar (3415, -3.99 DPS) [dungeon]; Death Speaker Scepter (2816, -4.45 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 27.1 healing_power points (3.33 DPS) | yes | Orb of Souls (249395, -0.18 DPS, sim-verified) [crafted]; Omega Orb (7749, -0.85 DPS) [quest]; Prophetic Cane (6803, -1.20 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Captain Rackmore's Tiller (16789, -0.33 DPS) [quest]; Dancing Flame (6806, -0.34 DPS) [quest]; Wand of Mana Concentration (273043, -0.49 DPS, sim-verified) [world] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 202.5. Weights run: 26.5s. Verify run: 240.3s. 510 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.146, intellect=1.253 ± 0.033, spirit=0.883 ± 0.030, mp5=2.839 ± 0.033, crit=0.514 ± 0.025 per rating point (14 rating = 1%, 7.199 per %), spell_haste=not significant (1.577 ± 0.671)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 51.9 healing_power points (7.23 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Papal Fez (9431, -0.36 DPS, sim-verified) [dungeon]; Gemburst Circlet (10751, -0.80 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.0 healing_power points (4.18 DPS) | yes | Horizon Choker (13085, -0.74 DPS, sim-verified) [world_drop]; Mindburst Medallion (11196, -1.68 DPS) [quest]; Gemshard Heart (17707, -1.70 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 41.3 healing_power points (5.75 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Kentic Amice (11624, -0.80 DPS) [dungeon]; Inquisitor's Shawl (19507, -1.62 DPS, sim-verified) [dungeon] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.07 DPS) [dungeon]; Deep Woodlands Cloak (19121, -0.56 DPS) [quest]; Darkspear Raider's Cloak (272076, -0.74 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.1 healing_power points (9.06 DPS) | yes | Ghostweave Vest (14141, -2.18 DPS) [crafted]; Embrace of the Wind Serpent (12462, -2.40 DPS) [world]; Stone Guard's Satin Armor (220903, -2.49 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 33.2 healing_power points (4.62 DPS) | yes | Aristocratic Cuffs (12546, -1.48 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -1.99 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.16 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 48.6 healing_power points (6.77 DPS) | yes | Greenleaf Handwraps (19116, -1.78 DPS) [quest]; Earthenweave Gloves (254075, -1.79 DPS) [crafted]; First Sergeant's Satin Gloves (220898, -3.70 DPS, sim-verified) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 40.3 healing_power points (5.61 DPS) | yes | Gilded Cord (254037, -0.41 DPS) [crafted]; Dawnspire Cord (12466, -0.62 DPS) [dungeon]; Earthenweave Cord (254077, -0.64 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Kilt of the Atal'ai Prophet (10807, -0.73 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.26 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -1.97 DPS, sim-verified) [vendor] |
| feet | Earthenweave Boots (254093) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | First Sergeant's Satin Boots (220900, -0.09 DPS) [vendor]; Gilded Slippers (254001, -0.47 DPS) [crafted]; Gilded Sandals (254107, -0.51 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 23.4 healing_power points (3.25 DPS) | yes | Cyclopean Band (11824, -0.29 DPS) [dungeon]; Eye of Adaegus (5266, -0.73 DPS) [world_drop]; Brainlash (6440, -4.64 DPS, sim-verified) [dungeon] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | sim-verified (202.5 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon]; Eye of Adaegus (5266, -0.24 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.25 DPS) [quest]; Alchemist's Stone (13503, -0.98 DPS) [crafted] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.36 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -1.48 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -6.33 DPS, sim-verified) [quest]; Charstone Dirk (17710, -8.19 DPS) [dungeon]; Dreamstaff (249454, -10.10 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 10.7 healing_power points (1.49 DPS) | yes | Charged Lightning Rod (11860, -0.09 DPS) [quest]; Jaina's Firestarter (13064, -0.23 DPS, sim-verified) [world_drop]; Dancing Flame (6806, -0.23 DPS) [quest] |

**New at 50:** head: Cassandra's Grace; shoulder: Nethergeld Shoulders; back: Battle Healer's Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Senior Designer's Pantaloons; feet: Earthenweave Boots; finger1: Advisor's Ring; finger2: Darkspear Signet; trinket2: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 510, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 289.7. Weights run: 27.5s. Verify run: 761.6s. 1354 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.146, intellect=1.601 ± 0.069, spirit=3.066 ± 0.063, mp5=6.371 ± 0.067, crit=0.392 ± 0.026 per rating point (14 rating = 1%, 5.491 per %), spell_haste=not significant (-2.755 ± 1.073)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 104.4 healing_power points (12.00 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -5.76 DPS, sim-verified) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 61.1 healing_power points (7.02 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.45 DPS, sim-verified) [world_drop]; Heart of the Fiend (13960, -0.81 DPS) [dungeon]; Animated Chain Necklace (18723, -1.11 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (289.7 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -7.90 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-verified (289.7 DPS) | yes | Featherskin Cape (10843, -0.47 DPS) [world]; Battle Healer's Cloak (19526, -0.68 DPS) [rep]; Frostweaver Cape (12968, -2.67 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 122.1 healing_power points (14.02 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Embrace of the Wind Serpent (12462, -0.33 DPS) [world]; Ghostweave Vest (14141, -8.15 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (289.7 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Wyrmthalak's Shackles (13958, -8.37 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (289.7 DPS) | yes | Swarmtender's Gloves (275607, -0.67 DPS) [crafted]; Magister's Gloves (16684, -1.24 DPS) [dungeon]; Hands of the Exalted Herald (12554, -6.82 DPS, sim-verified) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 96.6 healing_power points (11.10 DPS) | yes | General's Satin Cinch (17621, -1.58 DPS) [pvp]; Devout Belt (16696, -2.98 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -3.92 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 118.5 healing_power points (13.62 DPS) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -1.33 DPS) [pvp]; Haunting Specter Leggings (11929, -11.59 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (289.7 DPS) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; General's Satin Walkers (231630, -0.73 DPS) [vendor]; Mooncloth Boots (15802, -6.28 DPS, sim-verified) [crafted] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-verified (289.7 DPS) | yes | Darkspear Signet (272068, -0.40 DPS) [vendor]; The Postmaster's Seal (13392, -0.45 DPS) [dungeon]; Naglering (11669, -11.51 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (289.7 DPS) | yes | The Postmaster's Seal (13392, -0.42 DPS) [dungeon]; Blessed Band of Light (272407, -0.58 DPS) [vendor]; Darkspear Signet (272068, -4.80 DPS, sim-verified) [vendor] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (289.7 DPS) | yes | Ankh of Life (1713, -0.16 DPS) [world_drop]; Darkspear Voodoo Seal (272061, -0.73 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.87 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (289.7 DPS) | yes | Ankh of Life (1713, -0.60 DPS) [world_drop]; Darkspear Voodoo Seal (272061, -1.17 DPS) [vendor]; Second Wind (11819, -3.55 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (289.7 DPS) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Resurgence Rod (17743, -4.82 DPS) [quest]; Hand of Edward the Odd (2243, -5.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 50.7 healing_power points (5.83 DPS) | yes | Wand of Mana Concentration (273043, -3.63 DPS) [world]; Jaina's Firestarter (13064, -3.67 DPS) [world_drop]; Brilliant Wand (249385, -6.73 DPS, sim-verified) [crafted] |

**New at 60:** head: Devout Crown; neck: Lady Maye's Pendant; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Penitent's Cinch; legs: Devout Skirt; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Band of Mending; trinket1: Shard of the Splithooves; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1354, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 592.7. Weights run: 18.7s. Verify run: 543.3s. 1354 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.497, intellect=1.734 ± 0.090, spirit=1.606 ± 0.088, mp5=4.128 ± 0.101, crit=1.114 ± 0.078 per rating point (14 rating = 1%, 15.591 per %), spell_haste=not significant (1.384 ± 2.027)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 85.7 healing_power points (16.38 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Crown of the Penitent (13216, -14.82 DPS, sim-verified) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -12.08 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -1.66 DPS) [rep]; Darkspear Raider's Cloak (272063, -1.94 DPS) [vendor]; Frostweaver Cape (12968, -4.96 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 98.7 healing_power points (18.86 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Warlord's Satin Robes (231612, -0.12 DPS) [pvp]; Robes of the Exalted (13346, -3.03 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.15 DPS) [crafted]; Dryad's Wrist Bindings (19595, -12.06 DPS, sim-verified) [rep] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Handwraps (231633, -0.82 DPS) [vendor]; Hands of the Exalted Herald (12554, -1.42 DPS, sim-verified) [dungeon]; Devout Gloves (16692, -2.14 DPS) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 69.7 healing_power points (13.33 DPS) | yes | General's Satin Cinch (17621, -1.10 DPS) [pvp]; Devout Belt (16696, -1.83 DPS) [dungeon]; Penitent's Cinch (272394, -4.80 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Outrider's Silk Leggings (22747, -17.74 DPS, sim-verified) [rep] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; General's Satin Walkers (231630, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -11.69 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-verified (592.6 DPS) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Second Wind (11819, +0.00 DPS) [dungeon]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -3.29 DPS) [quest]; Weakness Analyzer (272438, -3.82 DPS) [vendor]; Second Wind (11819, -5.58 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -11.98 DPS) [pvp]; Ironbark Staff (20220, -13.29 DPS) [rep]; Resurgence Rod (17743, -18.11 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 34.7 healing_power points (6.63 DPS) | yes | Oblivion's Touch (18761, -2.98 DPS) [dungeon]; Bonecreeper Stylus (13938, -3.20 DPS) [dungeon]; Brilliant Wand (249385, -5.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Devout Crown; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1354, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

