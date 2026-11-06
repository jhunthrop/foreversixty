# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 48.8. Weights run: 2.7s. Verify run: 3.2s. 181 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.978 ± 0.007, crit=0.093 ± 0.003 per rating point (14 rating = 1%, 1.295 per %), hit=0.263 ± 0.008 per rating point (10 rating = 1%, 2.631 per %), spell_haste=0.623 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=2.134 ± 0.106

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.37 DPS) | yes | Pristine Circlet (253949, -1.08 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.32 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (2.98 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.17 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.60 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -2.11 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.02 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.21 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.22 DPS) [crafted] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 13.9 spell_power points (3.00 DPS) | yes | Filigreed Pristine Gown (253901, -0.86 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.08 DPS) [crafted]; Mystic's Wrap (14369, -1.52 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 spell_power points (2.44 DPS) | yes | Bright Bracers (3647, -0.41 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -2.01 DPS) [world_drop]; Repurposed Hair Band (281256, -2.01 DPS) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.14 DPS) | yes | Serpent Gloves (5970, -0.38 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.65 DPS) [crafted]; Gnoll Casting Gloves (892, -0.85 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Pristine Sash (253925, -0.65 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.67 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.86 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 16.9 spell_power points (3.64 DPS) | yes | Filigreed Pristine Leggings (253937, -1.08 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -2.13 DPS) [dungeon]; Darkweave Breeches (12987, -2.16 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Silky Boots (253891, -0.38 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.07 DPS) [crafted]; Red Woolen Boots (4313, -1.49 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.3 spell_power points (2.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.73 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.37 DPS) [world_drop]; Furlfeather Ring (286745, -1.78 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (1.50 DPS) | yes | Lavishly Jeweled Ring (1156, -0.30 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.87 DPS) [world_drop]; Furlfeather Ring (286745, -1.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.0 spell_power points (6.46 DPS) | yes | Emberstone Staff (5201, -1.28 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.45 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -1.72 DPS) [dungeon] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.29 DPS) | yes | Mathystra Channeling Orb (281291, -0.24 DPS) [quest]; Tear of Grief (5611, -0.43 DPS, sim-verified) [quest]; Dwarven Tome (279898, -0.65 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.47 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.68 DPS) [quest]; Greater Magic Wand (11288, -0.71 DPS) [crafted] |

**New at 20:** head: Silky Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 181, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 82.7. Weights run: 2.6s. Verify run: 2.2s. 309 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.301 ± 0.005, crit=0.111 ± 0.004 per rating point (14 rating = 1%, 1.552 per %), hit=0.208 ± 0.006 per rating point (10 rating = 1%, 2.076 per %), spell_haste=1.103 ± 0.061, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.179 ± 0.038, mp5=0.287 ± 0.042

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (2.89 DPS) | yes | Holy Shroud (2721, -0.40 DPS) [world_drop]; Silky Circlet (253951, -0.62 DPS) [crafted]; Azure Silk Hood (7048, -0.83 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 1.9 spell_power points (0.40 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.02 DPS) [world]; Crystal Starfire Medallion (5003, -0.04 DPS) [world_drop]; Seal of the Expedition (271920, -0.09 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (2.42 DPS) | yes | Death Speaker Mantle (6685, -0.51 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.66 DPS) [crafted]; Magician's Mantle (12998, -0.83 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-verified (82.7 DPS) | yes | Hillman's Cloak (3719, -1.14 DPS) [crafted]; Heavy Woolen Cloak (4311, -1.20 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.61 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 15.8 spell_power points (3.26 DPS) | yes | Beguiler Robes (7728, -0.37 DPS) [dungeon]; Tree Bark Jacket (1486, -0.58 DPS) [dungeon]; Green Silk Armor (7065, -0.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.9 spell_power points (2.03 DPS) | yes | Mindthrust Bracers (1974, -1.55 DPS) [dungeon]; Nightsky Wristbands (6407, -1.55 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.02 DPS, sim-verified) [world_drop] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.5 spell_power points (2.17 DPS) | yes | Silky Gloves (253915, -0.63 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.73 DPS) [dungeon]; Gloves of Meditation (4318, -0.88 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 spell_power points (2.46 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.70 DPS) [crafted]; Silky Sash (253927, -0.76 DPS) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 17.0 spell_power points (3.51 DPS) | yes | Filigreed Silky Leggings (253939, -0.72 DPS) [crafted]; Azure Silk Pants (7046, -1.07 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.61 DPS, sim-verified) [dungeon] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 15.8 spell_power points (3.26 DPS) | yes | Spidersilk Boots (4320, -1.57 DPS) [crafted]; Acidic Walkers (9454, -1.59 DPS) [dungeon]; Gilded Slippers (254001, -1.60 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.6 spell_power points (1.56 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Minor Channeling Ring (1449, -0.41 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.5 spell_power points (1.55 DPS) | yes | Ring of Power Regulation (273030, -0.13 DPS) [dungeon]; Minor Channeling Ring (1449, -0.39 DPS) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (82.7 DPS) | yes | Hydrocane (9452, -0.97 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.97 DPS) [world]; Glimmering Staff (249392, -3.11 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.82 DPS) | yes | Orb of Souls (249395, -0.77 DPS) [crafted]; Eye of Paleth (2943, -0.99 DPS) [quest]; Antipodean Rod (2879, -1.26 DPS, sim-verified) [world_drop] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-verified (82.7 DPS) | yes | Greater Mystic Wand (217287, -0.06 DPS) [crafted]; Wand of Eventide (5214, -0.31 DPS) [world_drop]; Starfaller (13063, -0.55 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Highlander's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 177.8. Weights run: 2.8s. Verify run: 4.8s. 400 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.349 ± 0.011, crit=0.274 ± 0.009 per rating point (14 rating = 1%, 3.831 per %), hit=0.550 ± 0.022 per rating point (10 rating = 1%, 5.497 per %), spell_haste=4.888 ± 0.192, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.362 ± 0.081, mp5=0.576 ± 0.096

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.97 DPS) | yes | Papal Fez (9431, -1.44 DPS) [dungeon]; Whitemane's Chapeau (7720, -2.34 DPS) [dungeon]; Augural Shroud (2620, -2.47 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -1.58 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.83 DPS) [dungeon]; Amberglow Talisman (10824, -2.34 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mantle of Doan (7712, -0.06 DPS) [dungeon]; Green Silken Shoulders (7057, -0.11 DPS) [crafted]; Azure Shoulders (7060, -1.32 DPS, sim-verified) [crafted] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Leather Cloak (7377, -0.03 DPS) [crafted]; Icy Cloak (4327, -0.73 DPS) [crafted]; Caretaker's Cape (19532, -1.24 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 26.2 spell_power points (9.95 DPS) | yes | Robe of the Magi (1716, -0.12 DPS) [world_drop]; Robe of Power (7054, -1.95 DPS) [crafted]; Elemental Raiment (9434, -1.98 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.7 spell_power points (4.07 DPS) | yes | Arcane Runed Bracers (4744, -1.32 DPS, sim-verified) [quest]; Condor Bracers (15864, -1.42 DPS) [quest]; Earthen Silk Cuffs (254019, -1.87 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.9 spell_power points (8.32 DPS) | yes | Frothing Handwraps (254023, -2.08 DPS) [crafted]; Black Mageweave Gloves (10003, -2.63 DPS) [crafted]; Red Mageweave Gloves (10018, -2.82 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 19.0 spell_power points (7.20 DPS) | yes | Star Belt (4329, -1.71 DPS) [crafted]; Gilded Cord (254037, -2.28 DPS) [crafted]; Highlander's Cloth Girdle (20098, -2.57 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gaze Dreamer Pants (6903, -1.11 DPS) [dungeon]; Silky Leggings (253989, -1.23 DPS, sim-verified) [crafted]; Filigreed Silky Leggings (253939, -1.38 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.5 spell_power points (10.07 DPS) | yes | Frothing Slippers (254003, -1.89 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -5.94 DPS) [crafted]; Thoughtcast Boots (10578, -6.39 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (4.46 DPS) | yes | Ring of Forlorn Spirits (2043, -1.42 DPS) [quest]; Snake Hoop (6750, -1.52 DPS) [quest]; Ring of Power Regulation (273030, -1.52 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 10.7 spell_power points (4.07 DPS) | yes | Ring of Forlorn Spirits (2043, -1.04 DPS) [quest]; Snake Hoop (6750, -1.13 DPS) [quest]; Ring of Power Regulation (273030, -1.14 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -11.41 DPS) [quest]; Hypnotic Blade (7714, -11.76 DPS) [dungeon]; Illusionary Rod (7713, -12.98 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.09 DPS) [dungeon]; Twisted Nether Wand (249144, -1.70 DPS) [crafted]; Burning Sliver (5249, -1.77 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 261.9. Weights run: 2.9s. Verify run: 5.7s. 512 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.439 ± 0.017, crit=0.393 ± 0.012 per rating point (14 rating = 1%, 5.501 per %), hit=0.842 ± 0.035 per rating point (10 rating = 1%, 8.422 per %), spell_haste=6.711 ± 0.297, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.516 ± 0.142), mp5=0.833 ± 0.157

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 31.6 spell_power points (11.99 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -1.74 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.59 DPS) [vendor]; Red Mageweave Headband (10033, -3.00 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 18.6 spell_power points (7.07 DPS) | yes | Triune Amulet (7722, -2.57 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -2.82 DPS) [quest]; Glowing Eye of Mordresh (10769, -3.08 DPS) [dungeon] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 23.6 spell_power points (8.96 DPS) | yes | Kentic Amice (11624, -0.31 DPS) [dungeon]; Rotgrip Mantle (17732, -1.03 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.34 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azure Silk Cloak (7053, -3.11 DPS) [crafted]; Frost Leather Cloak (7377, -3.31 DPS) [crafted]; Runecloth Cloak (13860, -3.68 DPS, sim-verified) [crafted] |
| chest | Frostweave Robe (13868) | Tailoring [crafted] | 29.7 spell_power points (11.27 DPS) | yes | Dreamweave Vest (10021, -0.19 DPS) [crafted]; Frostweave Tunic (13869, -0.32 DPS) [crafted]; Acumen Robes (17775, -0.72 DPS) [quest] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 18.7 spell_power points (7.08 DPS) | yes | Nethergeld Cuffs (254061, -2.19 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.72 DPS) [dungeon]; Bloodband Bracers (11469, -3.10 DPS) [quest] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 27.5 spell_power points (10.45 DPS) | yes | Dreamweave Gloves (10019, -1.58 DPS) [crafted]; Runecloth Gloves (13863, -2.63 DPS) [crafted]; Frostweave Gloves (13870, -4.12 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 25.3 spell_power points (9.58 DPS) | yes | Frothing Cord (254039, -1.76 DPS) [crafted]; Ghostweave Cord (254073, -2.12 DPS) [crafted]; Frothing Waistcord (254083, -2.21 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-verified (261.9 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -0.61 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.58 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.6 spell_power points (10.48 DPS) | yes | Frothing Sandals (254109, -2.20 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.75 DPS) [vendor]; Frothing Slippers (254003, -3.60 DPS) [crafted] |
| finger1 | Freezing Band (942) | World drop [world_drop] | 21.0 spell_power points (7.97 DPS) | yes | Cyclopean Band (11824, -2.60 DPS) [dungeon]; Band of the Unicorn (7553, -3.04 DPS) [world_drop]; Philanthropist's Ring (281635, -3.34 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 15.3 spell_power points (5.82 DPS) | yes | Cyclopean Band (11824, -0.45 DPS) [dungeon]; Band of the Unicorn (7553, -0.88 DPS) [world_drop]; Philanthropist's Ring (281635, -1.19 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Uther's Strength (11302, -0.07 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.39 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Uther's Strength (11302, -0.60 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.92 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -6.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -7.29 DPS) [world_drop]; Spire of Hakkar (10844, -7.77 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.15 DPS) [crafted]; Freezing Shard (10572, -0.24 DPS) [dungeon]; Noxious Shooter (17745, -0.92 DPS) [dungeon] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Frostweave Robe; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Ban'thok Sash; legs: Senior Designer's Pantaloons; finger1: Freezing Band; finger2: Lorekeeper's Ring; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 407.0. Weights run: 3.3s. Verify run: 21.5s. 1299 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.346 ± 0.024, crit=0.529 ± 0.015 per rating point (14 rating = 1%, 7.399 per %), hit=1.111 ± 0.055 per rating point (10 rating = 1%, 11.107 per %), spell_haste=5.954 ± 0.324, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.814 ± 0.001, spirit=not significant (-0.001 ± 0.079), mp5=not significant (0.086 ± 0.133)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcane Crystal Pendant (20037, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -0.51 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.9 spell_power points (11.83 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Boreal Mantle (11782, -1.04 DPS) [dungeon]; Lieutenant Commander's Silk Mantle (227102, -1.10 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.59 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.65 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -6.13 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -2.16 DPS) [pvp]; Freezing Lich Robes (14340, -6.78 DPS) [dungeon]; Robe of Winter Night (14136, -9.73 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.8 spell_power points (10.14 DPS) | yes | Runecloth Cuffs (254123, -4.22 DPS) [crafted]; Netherfroth Cuffs (254063, -4.82 DPS) [crafted]; Nethergeld Cuffs (254061, -6.28 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (407.0 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 27.3 spell_power points (11.18 DPS) | yes | Ban'thok Sash (11662, -0.16 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -1.56 DPS) [rep]; Stormpike Cloth Girdle (19094, -2.39 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 50.3 spell_power points (20.58 DPS) | yes | Marshal's Silk Leggings (231605, -2.43 DPS) [pvp]; Skyshroud Leggings (13170, -5.64 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Legguards (227109, -6.55 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.1 spell_power points (9.89 DPS) | yes | Knight-Lieutenant's Silk Walkers (227112, +0.00 DPS) [pvp]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.06 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.28 DPS) [quest]; Don Julio's Band (19325, -1.07 DPS) [rep]; Naglering (11669, -11.57 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Don Julio's Band (19325, -0.78 DPS) [rep]; Don Mauricio's Band of Domination (22433, -0.83 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Uther's Strength (11302, -1.64 DPS) [world_drop]; Shard of the Splithooves (10659, -3.88 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.37 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -4.10 DPS) [dungeon]; Uther's Strength (11302, -6.55 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -1.41 DPS) [pvp]; Sageclaw (20070, -4.83 DPS) [rep]; Ironbark Staff (20069, -11.48 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.88 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Tome of the Ice Lord (19310, -0.42 DPS) [rep]; Spirit of Aquementas (11904, -3.28 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.28 DPS) [dungeon]; Brilliant Wand (249385, -0.73 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -0.83 DPS) [dungeon] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 677.5. Weights run: 3.6s. Verify run: 20.4s. 1299 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.171 ± 0.010, crit=0.659 ± 0.019 per rating point (14 rating = 1%, 9.231 per %), hit=1.617 ± 0.078 per rating point (10 rating = 1%, 16.173 per %), spell_haste=10.302 ± 0.419, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.823 ± 0.002, spirit=not significant (0.046 ± 0.029), mp5=not significant (0.068 ± 0.040)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Arcane Crystal Pendant (20037, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.5 spell_power points (14.91 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.34 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.13 DPS) [pvp]; Argent Shoulders (19059, -1.64 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.39 DPS, sim-verified) [world]; Shroud of Arcane Mastery (22330, -8.11 DPS) [dungeon]; Cloak of Flames (3475, -9.10 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -3.07 DPS) [pvp]; Freezing Lich Robes (14340, -4.08 DPS, sim-verified) [dungeon]; Robe of Winter Night (14136, -8.54 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.7 spell_power points (12.41 DPS) | yes | Runecloth Cuffs (254123, -5.75 DPS) [crafted]; Netherfroth Cuffs (254063, -6.06 DPS) [crafted]; Arena Wristguards (18709, -7.43 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.2 spell_power points (15.28 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Magician's Cord (272393, -0.60 DPS) [vendor]; Stormpike Cloth Girdle (19094, -2.38 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.2 spell_power points (27.33 DPS) | yes | Marshal's Silk Leggings (231605, -4.74 DPS) [pvp]; Skyshroud Leggings (13170, -8.65 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Legguards (227109, -9.85 DPS) [pvp] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP reward · Alliance [vendor] | 25.7 spell_power points (13.47 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.73 DPS) [crafted]; Knight-Lieutenant's Silk Walkers (227112, -1.57 DPS) [pvp] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -3.31 DPS) [dungeon]; Maiden's Circle (13001, -3.70 DPS) [world_drop]; Naglering (11669, -12.64 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.71 DPS) [dungeon]; Maiden's Circle (13001, -3.10 DPS) [world_drop]; Songstone of Ironforge (12543, -3.15 DPS) [quest] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.26 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -5.24 DPS) [dungeon]; Uther's Strength (11302, -8.38 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (677.5 DPS) | yes | Burst of Knowledge (11832, -1.34 DPS) [dungeon]; Uther's Strength (11302, -4.48 DPS) [world_drop]; Ankh of Life (1713, -7.33 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.95 DPS) [pvp]; Staff of Jordan (873, -11.23 DPS, sim-verified) [world_drop]; Elemental Mage Staff (944, -11.67 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.78 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.13 DPS) [dungeon]; Brilliant Wand (249385, -1.55 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 47.4. Weights run: 2.7s. Verify run: 3.1s. 175 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.978 ± 0.007, crit=0.093 ± 0.003 per rating point (14 rating = 1%, 1.295 per %), hit=0.263 ± 0.008 per rating point (10 rating = 1%, 2.631 per %), spell_haste=0.623 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=2.134 ± 0.106

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.37 DPS) | yes | Pristine Circlet (253949, -1.06 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.32 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (2.98 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.98 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.60 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -2.11 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.20 DPS) [crafted]; Sanguine Cape (14376, -0.22 DPS) [world_drop]; Black Whelp Cloak (7283, -0.42 DPS) [crafted] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 13.9 spell_power points (3.00 DPS) | yes | Filigreed Pristine Gown (253901, -0.85 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.08 DPS) [crafted]; Mystic's Wrap (14369, -1.52 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 spell_power points (2.44 DPS) | yes | Bright Bracers (3647, -0.74 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -1.79 DPS) [quest]; Owlbeard Bracers (16981, -1.80 DPS) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.14 DPS) | yes | Pristine Gloves (253913, -0.65 DPS) [crafted]; Serpent Gloves (5970, -0.69 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.85 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Pristine Sash (253925, -0.63 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.67 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.86 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 16.9 spell_power points (3.64 DPS) | yes | Filigreed Pristine Leggings (253937, -1.06 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -2.13 DPS) [dungeon]; Darkweave Breeches (12987, -2.16 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Silky Boots (253891, -0.54 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.07 DPS) [crafted]; Red Woolen Boots (4313, -1.49 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.3 spell_power points (2.00 DPS) | yes | Loop of Sacrifice (281673, -0.94 DPS) [quest]; Volcanic Rock Ring (12053, -1.37 DPS) [world_drop]; Furlfeather Ring (286745, -1.78 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.9 spell_power points (1.27 DPS) | yes | Loop of Sacrifice (281673, -0.33 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.63 DPS) [world_drop]; Furlfeather Ring (286745, -1.05 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (47.4 DPS) | yes | Staff of the Blessed Seer (2271, -0.42 DPS) [dungeon]; Crescent Staff (6505, -0.57 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.85 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.29 DPS) | yes | Seer's Fine Stein (7608, -0.66 DPS) [world_drop]; Spellbinder Orb (15926, -0.66 DPS) [dungeon]; Tork Wrench (11855, -0.87 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.41 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.68 DPS) [quest]; Greater Magic Wand (11288, -0.72 DPS) [crafted] |

**New at 20:** head: Silky Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 175, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance

### Band 30 (undead, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 82.0. Weights run: 2.6s. Verify run: 3.0s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.301 ± 0.005, crit=0.111 ± 0.004 per rating point (14 rating = 1%, 1.552 per %), hit=0.208 ± 0.006 per rating point (10 rating = 1%, 2.076 per %), spell_haste=1.103 ± 0.061, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.179 ± 0.038, mp5=0.287 ± 0.042

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (2.89 DPS) | yes | Silky Circlet (253951, -0.62 DPS) [crafted]; Azure Silk Hood (7048, -0.83 DPS) [crafted]; Holy Shroud (2721, -0.91 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.06 DPS) [vendor]; Kaleidoscope Chain (13084, -0.41 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (2.42 DPS) | yes | Death Speaker Mantle (6685, -0.49 DPS) [dungeon]; Invoker's Mantle (215365, -0.66 DPS) [crafted]; Magician's Mantle (12998, -0.83 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -1.14 DPS) [crafted]; Heavy Woolen Cloak (4311, -1.20 DPS) [crafted]; Windsong Drape (15468, -1.27 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 15.8 spell_power points (3.26 DPS) | yes | Beguiler Robes (7728, -0.37 DPS) [dungeon]; Tree Bark Jacket (1486, -0.58 DPS) [dungeon]; Green Silk Armor (7065, -0.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.9 spell_power points (2.03 DPS) | yes | Glowing Magical Bracelets (13106, -1.54 DPS) [world_drop]; Mindthrust Bracers (1974, -1.55 DPS) [dungeon]; Tabitha's Cuffs (251486, -2.43 DPS, sim-verified) [quest] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.5 spell_power points (2.17 DPS) | yes | Silky Gloves (253915, -0.54 DPS) [crafted]; Jutebraid Gloves (10654, -0.63 DPS) [quest]; Serpent Gloves (5970, -0.73 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 spell_power points (2.46 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.70 DPS) [crafted]; Warsong Sash (16975, -0.99 DPS, sim-verified) [quest] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 17.0 spell_power points (3.51 DPS) | yes | Filigreed Silky Leggings (253939, -0.72 DPS) [crafted]; Azure Silk Pants (7046, -1.07 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.69 DPS, sim-verified) [dungeon] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 15.8 spell_power points (3.26 DPS) | yes | Gilded Slippers (254001, -1.52 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.57 DPS) [crafted]; Acidic Walkers (9454, -1.59 DPS) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.6 spell_power points (1.56 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Solid Gold Band (279845, -0.55 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.5 spell_power points (1.55 DPS) | yes | Solid Gold Band (279845, -0.54 DPS) [quest]; Ring of Power Regulation (273030, -0.56 DPS, sim-verified) [dungeon]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.97 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.97 DPS) [world]; Glimmering Staff (249392, -2.22 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.82 DPS) | yes | Orb of Souls (249395, -0.77 DPS) [crafted]; Witch's Finger (16887, -1.32 DPS) [quest]; Antipodean Rod (2879, -2.12 DPS, sim-verified) [world_drop] |
| ranged | Greater Mystic Wand (217287) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Wand of Eventide (5214, -0.33 DPS) [world_drop]; Starfaller (13063, -0.39 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Defiler's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Greater Mystic Wand

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 176.9. Weights run: 2.8s. Verify run: 4.4s. 392 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.349 ± 0.011, crit=0.274 ± 0.009 per rating point (14 rating = 1%, 3.831 per %), hit=0.550 ± 0.022 per rating point (10 rating = 1%, 5.497 per %), spell_haste=4.888 ± 0.192, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.362 ± 0.081, mp5=0.576 ± 0.096

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (176.9 DPS) | yes | Papal Fez (9431, -1.44 DPS) [dungeon]; Whitemane's Chapeau (7720, -2.34 DPS) [dungeon]; Augural Shroud (2620, -2.47 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Triune Amulet (7722, +0.00 DPS) [dungeon]; Necklace of Calisea (1714, -1.83 DPS) [dungeon]; Amberglow Talisman (10824, -2.34 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (176.9 DPS) | yes | Azure Shoulders (7060, +0.00 DPS) [crafted]; Bloodmage Mantle (7684, +0.00 DPS) [dungeon]; Mantle of Doan (7712, +0.00 DPS) [dungeon] |
| back | Frost Leather Cloak (7377) | Leatherworking [crafted] | sim-verified (176.9 DPS) | yes | Azure Silk Cloak (7053, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.69 DPS) [crafted]; Battle Healer's Cloak (19528, -1.21 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 26.2 spell_power points (9.95 DPS) | yes | Robe of the Magi (1716, -0.12 DPS) [world_drop]; Robe of Power (7054, -1.95 DPS) [crafted]; Elemental Raiment (9434, -1.98 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.7 spell_power points (4.07 DPS) | yes | Condor Bracers (15864, -1.42 DPS) [quest]; Radiant Silver Bracers (4545, -1.49 DPS) [quest]; Earthen Silk Cuffs (254019, -1.87 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.9 spell_power points (8.32 DPS) | yes | Frothing Handwraps (254023, -2.08 DPS) [crafted]; Black Mageweave Gloves (10003, -2.63 DPS) [crafted]; Red Mageweave Gloves (10018, -2.82 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 19.0 spell_power points (7.20 DPS) | yes | Defiler's Cloth Girdle (20166, -1.36 DPS) [rep]; Star Belt (4329, -1.71 DPS) [crafted]; Gilded Cord (254037, -2.28 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Leggings (253989, +0.00 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.11 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.38 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.5 spell_power points (10.07 DPS) | yes | Frothing Slippers (254003, -3.66 DPS) [crafted]; Gilded Slippers (254001, -5.94 DPS) [crafted]; Thoughtcast Boots (10578, -6.39 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (4.46 DPS) | yes | Snake Hoop (6750, -1.52 DPS) [quest]; Ring of Power Regulation (273030, -1.52 DPS) [dungeon]; Reedknot Ring (9622, -1.80 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 10.7 spell_power points (4.07 DPS) | yes | Snake Hoop (6750, -1.13 DPS) [quest]; Ring of Power Regulation (273030, -1.14 DPS) [dungeon]; Reedknot Ring (9622, -1.42 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -10.02 DPS) [dungeon]; Celestial Stave (9517, -11.41 DPS) [quest]; Hypnotic Blade (7714, -11.76 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.13 DPS) [dungeon]; Twisted Nether Wand (249144, -1.80 DPS) [crafted]; Dancing Flame (6806, -2.26 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Frost Leather Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 257.1. Weights run: 2.9s. Verify run: 5.9s. 504 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.439 ± 0.017, crit=0.393 ± 0.012 per rating point (14 rating = 1%, 5.501 per %), hit=0.842 ± 0.035 per rating point (10 rating = 1%, 8.422 per %), spell_haste=6.711 ± 0.297, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.516 ± 0.142), mp5=0.833 ± 0.157

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 31.6 spell_power points (11.99 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -1.74 DPS) [crafted]; Red Mageweave Headband (10033, -2.51 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.59 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 18.6 spell_power points (7.07 DPS) | yes | Mindburst Medallion (11196, -2.82 DPS) [quest]; Glowing Eye of Mordresh (10769, -3.08 DPS) [dungeon]; Triune Amulet (7722, -3.76 DPS, sim-verified) [dungeon] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 23.6 spell_power points (8.96 DPS) | yes | Kentic Amice (11624, -0.31 DPS) [dungeon]; Rotgrip Mantle (17732, -1.03 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.34 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.98 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.74 DPS) [crafted]; Azure Silk Cloak (7053, -3.11 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frostweave Tunic (13869, -0.13 DPS) [crafted]; Acumen Robes (17775, -0.53 DPS) [quest]; Frostweave Robe (13868, -0.83 DPS, sim-verified) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 18.7 spell_power points (7.08 DPS) | yes | Nethergeld Cuffs (254061, -2.14 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.72 DPS) [dungeon]; Bloodband Bracers (11469, -3.10 DPS) [quest] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 27.5 spell_power points (10.45 DPS) | yes | Dreamweave Gloves (10019, -1.58 DPS) [crafted]; Runecloth Gloves (13863, -2.63 DPS) [crafted]; Frostweave Gloves (13870, -3.57 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 25.3 spell_power points (9.58 DPS) | yes | Frothing Waistcord (254083, -0.87 DPS) [crafted]; Frothing Cord (254039, -1.76 DPS) [crafted]; Ghostweave Cord (254073, -2.12 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 27.4 spell_power points (10.40 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.75 DPS) [vendor]; Senior Designer's Pantaloons (11841, -2.09 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -2.73 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.6 spell_power points (10.48 DPS) | yes | Frothing Sandals (254109, -1.39 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.75 DPS) [vendor]; Frothing Slippers (254003, -3.60 DPS) [crafted] |
| finger1 | Freezing Band (942) | World drop [world_drop] | 21.0 spell_power points (7.97 DPS) | yes | Cyclopean Band (11824, -2.60 DPS) [dungeon]; Band of the Unicorn (7553, -3.04 DPS) [world_drop]; Philanthropist's Ring (281635, -3.34 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 15.3 spell_power points (5.82 DPS) | yes | Cyclopean Band (11824, -0.45 DPS) [dungeon]; Band of the Unicorn (7553, -0.88 DPS) [world_drop]; Philanthropist's Ring (281635, -1.19 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.07 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.11 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -6.19 DPS) [world]; Glowing Brightwood Staff (812, -7.29 DPS) [world_drop]; Kindling Stave (11750, -8.41 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.00 DPS) [crafted]; Freezing Shard (10572, -0.27 DPS) [dungeon]; Noxious Shooter (17745, -0.71 DPS) [dungeon] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Freezing Band; finger2: Advisor's Ring; trinket2: Darkspear Voodoo Seal

No-known-source sample (15 of 504, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 410.0. Weights run: 3.3s. Verify run: 23.3s. 1293 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.346 ± 0.024, crit=0.529 ± 0.015 per rating point (14 rating = 1%, 7.399 per %), hit=1.111 ± 0.055 per rating point (10 rating = 1%, 11.107 per %), spell_haste=5.954 ± 0.324, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.814 ± 0.001, spirit=not significant (-0.001 ± 0.079), mp5=not significant (0.086 ± 0.133)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.8 spell_power points (13.42 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.45 DPS, sim-verified) [crafted] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 22.9 spell_power points (9.37 DPS) | yes | Archlight Talisman (15856, -3.04 DPS) [quest]; Arcane Crystal Pendant (20037, -3.04 DPS, sim-verified) [quest]; Star of Mystaria (12103, -3.54 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.9 spell_power points (11.83 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.10 DPS) [pvp]; Boreal Mantle (11782, -2.56 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.52 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.65 DPS) [dungeon]; Deep Woodlands Cloak (19121, -6.05 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -2.16 DPS) [pvp]; Freezing Lich Robes (14340, -6.78 DPS) [dungeon]; Robe of Winter Night (14136, -9.86 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.8 spell_power points (10.14 DPS) | yes | Runecloth Cuffs (254123, -4.22 DPS) [crafted]; Netherfroth Cuffs (254063, -4.82 DPS) [crafted]; Nethergeld Cuffs (254061, -6.28 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; General's Silk Handwraps (231600, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 27.3 spell_power points (11.18 DPS) | yes | Ban'thok Sash (11662, -0.16 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -1.56 DPS) [rep]; Frostwolf Cloth Belt (19090, -2.39 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.12 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.1 spell_power points (9.89 DPS) | yes | Blood Guard's Silk Walkers (227110, +0.00 DPS) [pvp]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.06 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.28 DPS) [quest]; Don Julio's Band (19325, -1.07 DPS) [rep]; Naglering (11669, -11.52 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Don Julio's Band (19325, -0.78 DPS) [rep]; Don Mauricio's Band of Domination (22433, -0.83 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -3.28 DPS) [vendor]; Frozen Heart of the Mountain (249469, -3.78 DPS, sim-verified) [crafted]; Rune of the Guard Captain (19120, -5.83 DPS) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (-1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.82 DPS) [crafted]; Rune of the Guard Captain (19120, -1.73 DPS) [quest]; Uther's Strength (11302, -2.46 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.41 DPS) [pvp]; Mindfang (20214, -4.83 DPS) [rep]; Ironbark Staff (20220, -12.50 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (11.88 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Tome of the Ice Lord (19310, -0.42 DPS) [rep]; Spirit of Aquementas (11904, -3.28 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-verified (410.0 DPS) | yes | Bonecreeper Stylus (13938, -0.28 DPS) [dungeon]; Brilliant Wand (249385, -0.73 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -0.83 DPS) [dungeon] |

**New at 60:** head: Crimson Felt Hat; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Burst of Knowledge; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 675.2. Weights run: 3.6s. Verify run: 19.9s. 1293 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.171 ± 0.010, crit=0.659 ± 0.019 per rating point (14 rating = 1%, 9.231 per %), hit=1.617 ± 0.078 per rating point (10 rating = 1%, 16.173 per %), spell_haste=10.302 ± 0.419, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.823 ± 0.002, spirit=not significant (0.046 ± 0.029), mp5=not significant (0.068 ± 0.040)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.5 spell_power points (13.90 DPS) | yes | Arcane Crystal Pendant (20037, -4.98 DPS) [quest]; Star of Mystaria (12103, -5.04 DPS, sim-verified) [dungeon]; Archlight Talisman (15856, -6.72 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.5 spell_power points (14.91 DPS) | yes | Warlord's Silk Amice (231594, -0.34 DPS) [pvp]; Champion's Silk Mantle (227104, -1.13 DPS) [pvp]; Argent Shoulders (19059, -1.64 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.73 DPS) [world]; Shroud of Arcane Mastery (22330, -8.11 DPS) [dungeon]; Cloak of Flames (3475, -9.05 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -3.07 DPS) [pvp]; Robe of Winter Night (14136, -8.54 DPS) [crafted]; Freezing Lich Robes (14340, -10.21 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.7 spell_power points (12.41 DPS) | yes | Runecloth Cuffs (254123, -5.75 DPS) [crafted]; Netherfroth Cuffs (254063, -6.06 DPS) [crafted]; Arena Wristguards (18709, -7.43 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.2 spell_power points (15.28 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Magician's Cord (272393, -0.60 DPS) [vendor]; Frostwolf Cloth Belt (19090, -2.38 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-verified (675.2 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP reward · Horde [vendor] | 25.7 spell_power points (13.47 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.57 DPS) [pvp]; Earthen Silk Slippers (254013, -2.95 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -3.31 DPS) [dungeon]; Maiden's Circle (13001, -3.70 DPS) [world_drop]; Naglering (11669, -15.43 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -3.10 DPS) [world_drop]; Eye of Orgrimmar (12545, -3.15 DPS) [quest]; Don Mauricio's Band of Domination (22433, -5.14 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+10.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -1.34 DPS) [dungeon]; Rune of the Guard Captain (19120, -1.69 DPS) [quest]; Uther's Strength (11302, -4.48 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.21 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -5.24 DPS) [dungeon]; Rune of the Guard Captain (19120, -5.59 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.95 DPS) [pvp]; Elemental Mage Staff (944, -11.67 DPS) [world_drop]; Staff of Jordan (873, -14.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.78 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.13 DPS) [dungeon]; Brilliant Wand (249385, -1.55 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

