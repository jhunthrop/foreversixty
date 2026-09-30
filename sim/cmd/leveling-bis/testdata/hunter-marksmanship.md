# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 73.7. Weights run: 0.7s. Verify run: 1.1s. 385 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.004, agility=2.000 ± 0.009, crit=6.978 ± 1.340, hit=not significant (4.521 ± 1.506), melee_haste=not significant (6.060 ± 4.886)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.0 | yes | Shadow Goggles (4373, -0.96 DPS) [crafted]; Lucky Fishing Hat (19972, -0.96 DPS) [quest]; Flying Tiger Goggles (4368, -1.22 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.0 | yes | Tarnished Locket (279870, -0.96 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.60 DPS, sim-verified) [crafted]; Reinforced Woolen Shoulders (4315, -0.60 DPS) [crafted]; Forest Leather Mantle (4709, -0.60 DPS) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 | yes | Cape of the Brotherhood (5193, -0.05 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.24 DPS) [world_drop]; Hide of Lupos (3018, -0.24 DPS) [world] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 22.0 | yes | Trapper's Leather Armor (252491, -0.48 DPS) [crafted]; Brawler's Leather Armor (252490, -0.51 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.60 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.0 | yes | Bravo's Armbands (270015, -0.12 DPS) [quest]; Wolf Bracers (4794, -0.12 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.24 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 97.7 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -5.12 DPS) [dungeon]; Forest Leather Gloves (3058, -5.36 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.45 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.48 DPS) [quest]; Guardsman Belt (3429, -0.60 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.0 | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.12 DPS) [world]; Brawler's Leather Pants (252500, -0.17 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.0 | yes | Footpads of the Fang (10411, -0.24 DPS) [dungeon]; Blackened Defias Boots (10402, -0.25 DPS, sim-verified) [dungeon]; Dark Leather Boots (2315, -0.36 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.0 | yes | Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; The 1 Ring (8350, -0.60 DPS) [world]; Minor Channeling Ring (1449, -0.72 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.36 DPS) [world]; Minor Channeling Ring (1449, -0.48 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 297.4 | yes | Night Reaver (1318, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 274.6 | yes | Grayson's Torch (1172, -16.42 DPS) [quest]; Pulsating Hydra Heart (5183, -16.42 DPS) [world]; Tear of Grief (5611, -16.42 DPS) [quest] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.3 | yes | Lovingly Crafted Boomstick (4372, -2.52 DPS) [crafted]; Venomstrike (6469, -2.64 DPS) [dungeon]; Cracked Blacksmith Hammer (285279, -3.25 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Ranger Bow

No-known-source sample (15 of 385, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 104.3. Weights run: 0.7s. Verify run: 1.3s. 1260 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.006, agility=2.000 ± 0.012, crit=13.274 ± 2.589, hit=not significant (6.955 ± 2.095), melee_haste=not significant (8.860 ± 5.066)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 185.8 | yes | Nightscape Headband (8176, +0.00 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -9.44 DPS) [crafted]; White Bandit Mask (10008, -9.56 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 22.0 | yes | Sentinel's Medallion (19541, -0.42 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.47 DPS) [dungeon]; Sentinel's Medallion (20444, -0.58 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 34.0 | yes | Nightscape Shoulders (8192, -0.70 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.76 DPS, sim-verified) [world_drop]; Mantle of Thieves (2264, -0.82 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | World drop [world_drop] | 16.0 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Yeti Fur Cloak (2805, -0.23 DPS) [quest]; Darktide Cape (4114, -0.23 DPS) [quest] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 30.0 | yes | Tough Scorpid Breastplate (8203, +0.00 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.12 DPS) [crafted]; Hawkeye's Tunic (14592, -0.35 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.11 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.23 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 205.8 | yes | Dragonscale Gauntlets (8347, -0.39 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 193.8 | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -9.91 DPS) [rep]; Highlander's Leather Girdle (20117, -9.91 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 42.0 | yes | Hawkeye's Breeches (14595, -0.82 DPS) [world]; Ferine Leggings (6690, -0.93 DPS) [dungeon]; Triprunner Dungarees (9624, -1.20 DPS, sim-verified) [quest] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 26.0 | yes | Dusky Boots (7390, -0.23 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted]; Imperial Leather Boots (6431, -0.28 DPS, sim-verified) [world_drop] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 20.0 | yes | Protector's Band (19515, -0.23 DPS) [rep]; Disengagement Ring (276202, -0.23 DPS) [vendor]; Monkey Ring (6748, -0.35 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.0 | yes | Disengagement Ring (276202, -0.12 DPS) [vendor]; Protector's Band (19515, -0.15 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.23 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 536.8 | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 536.3 | yes | Shoni's Disarming Tool (9608, -15.23 DPS) [quest]; Grayson's Torch (1172, -31.28 DPS) [quest]; Rod of Molten Fire (2565, -31.28 DPS) [world_drop] |
| ranged | Sniper Rifle (3430) | World drop [world_drop] | 0.0 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.15 DPS) [quest]; Master Hunter's Bow (17686, -1.64 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 1260, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 189.3. Weights run: 0.6s. Verify run: 1.3s. 2326 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=18.230 ± 3.490, hit=not significant (0.000 ± 0.000), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Cryptstalker Headpiece [quest] | 572.4 | yes | Champion's Chain Helm (23251, -1.48 DPS) [vendor]; Lieutenant Commander's Chain Helm (23306, -1.48 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -1.48 DPS) [pvp] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | Onyxia Tooth Pendant (18404, -0.57 DPS) [quest]; Choker of the Shifting Sands (21505, -13.54 DPS) [quest]; Amulet of the Darkmoon (19491, -13.77 DPS) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Cryptstalker Spaulders [quest] | 313.2 | yes | Champion's Chain Pauldrons (227078, -1.14 DPS) [pvp]; Lieutenant Commander's Chain Pauldrons (227084, -1.14 DPS) [pvp]; Champion's Chain Shoulders (23252, -1.26 DPS) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 0.0 | yes | Cape of the Black Baron (13340, -0.11 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.46 DPS) [rep]; Chromatic Cloak (18509, -3.08 DPS, sim-verified) [crafted] |
| chest | Bloodsoul Breastplate (19690) | Blacksmithing [crafted] | 528.4 | yes | Legionnaire's Chain Hauberk (22874, +0.00 DPS) [vendor]; Knight-Captain's Chain Hauberk (23292, +0.00 DPS) [vendor]; Legionnaire's Chain Hauberk (227071, +0.00 DPS) [pvp] |
| wrist | Cryptstalker Wristguards (22443) | Cryptstalker Wristguards [quest] | 52.0 | yes | Marshal's Chain Bracers (16461, -0.69 DPS) [pvp]; General's Chain Wristguards (16570, -0.69 DPS) [pvp]; Forest Stalker's Bracers (19587, -8.45 DPS, sim-verified) [rep] |
| hands | Cryptstalker Handguards (22441) | Cryptstalker Handguards [quest] | 303.2 | yes | Marshal's Chain Grips (16463, -0.34 DPS) [vendor]; General's Chain Gloves (16571, -0.34 DPS) [vendor]; Chromatic Gauntlets (19157, -9.39 DPS, sim-verified) [crafted] |
| waist | Cryptstalker Girdle (22442) | Cryptstalker Girdle [quest] | 301.2 | yes | Highlander's Leather Girdle (20045, -0.69 DPS) [rep]; Light Obsidian Belt (22195, -0.80 DPS) [crafted]; Highlander's Chain Girdle (20043, -10.79 DPS, sim-verified) [rep] |
| legs | Cryptstalker Legguards (22437) | Cryptstalker Legguards [quest] | 0.0 | yes | Legionnaire's Chain Legguards (22875, +0.00 DPS) [vendor]; Knight-Captain's Chain Legguards (23293, +0.00 DPS) [vendor]; Knight-Captain's Chain Legguards (227072, +0.00 DPS) [pvp] |
| feet | Cryptstalker Boots (22440) | Cryptstalker Boots [quest] | 66.0 | yes | General's Chain Sabatons (231564, -0.69 DPS) [pvp]; Scalegut Treaders (275618, -0.69 DPS) [crafted]; Striker's Footguards (21365, -7.39 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 271.2 | yes | Dragonslayer's Signet (18403, -0.91 DPS) [quest]; Ring of Entropy (18543, -0.91 DPS) [world]; Mindtear Band (20632, -0.91 DPS) [world] |
| finger2 | Band of the Penitent (13217) (or Dragonslayer's Signet (18403), Ring of Entropy (18543), Mindtear Band (20632), Band of Earthen Wrath (21179), Band of Earthen Might (21182), Don Rodrigo's Band (21563), Ritssyn's Ring of Chaos (21836), Ring of the Eternal Flame (23237)) | Houses of the Holy [quest] | 255.2 | yes | Dragonslayer's Signet (18403, +0.00 DPS, sim-verified) [quest]; Ring of Entropy (18543, +0.00 DPS) [world]; Mindtear Band (20632, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Onyxia Blood Talisman (18406) | Celebrating Good Times [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 1023.0 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | High Warlord's Recurve (234559, -3.10 DPS) [pvp]; High Warlord's Crossbow (234560, -3.10 DPS) [pvp]; High Warlord's Street Sweeper (234561, -3.10 DPS) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Medallion of the Dawn; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Bloodsoul Breastplate; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Cryptstalker Legguards; feet: Cryptstalker Boots; finger1: Don Julio's Band; finger2: Band of the Penitent; trinket1: Ankh of Life; trinket2: Onyxia Blood Talisman; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 2326, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 73.5. Weights run: 0.7s. Verify run: 1.1s. 383 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.004, agility=2.000 ± 0.009, crit=6.978 ± 1.340, hit=not significant (4.521 ± 1.506), melee_haste=not significant (6.060 ± 4.886)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.0 | yes | Shadow Goggles (4373, -0.96 DPS) [crafted]; Lucky Fishing Hat (19972, -0.96 DPS) [quest]; Flying Tiger Goggles (4368, -0.96 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.0 | yes | Tarnished Locket (279870, -0.71 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 | yes | Reinforced Woolen Shoulders (4315, -0.60 DPS) [crafted]; Forest Leather Mantle (4709, -0.60 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.61 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 | yes | Sentry Cloak (2059, -0.24 DPS) [world_drop]; Hide of Lupos (3018, -0.24 DPS) [world]; Cape of the Brotherhood (5193, -0.25 DPS, sim-verified) [dungeon] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 14.0 | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.12 DPS) [crafted]; Heckler's Hide (286536, -0.24 DPS) [world] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.0 | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.24 DPS) [vendor]; Spare Part Bindings (279875, -0.24 DPS) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 97.7 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -5.12 DPS) [dungeon]; Forest Leather Gloves (3058, -5.36 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.48 DPS) [quest]; Guardsman Belt (3429, -0.60 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.0 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.12 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.0 | yes | Footpads of the Fang (10411, -0.24 DPS) [dungeon]; Blackened Defias Boots (10402, -0.27 DPS, sim-verified) [dungeon]; Dark Leather Boots (2315, -0.36 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.0 | yes | Bounty Hunter's Ring (5351, -0.36 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; The 1 Ring (8350, -0.60 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 | yes | Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; The 1 Ring (8350, -0.36 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 297.4 | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 274.6 | yes | Grayson's Torch (1172, -16.42 DPS) [quest]; Nightglow Concoction (3451, -16.42 DPS) [quest]; Pulsating Hydra Heart (5183, -16.42 DPS) [world] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.3 | yes | Lovingly Crafted Boomstick (4372, -2.52 DPS) [crafted]; Venomstrike (6469, -2.64 DPS) [dungeon]; Cracked Blacksmith Hammer (285279, -3.42 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Ranger Bow

No-known-source sample (15 of 383, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 104.8. Weights run: 0.7s. Verify run: 1.3s. 1258 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.006, agility=2.000 ± 0.012, crit=13.274 ± 2.589, hit=not significant (6.955 ± 2.095), melee_haste=not significant (8.860 ± 5.066)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 185.8 | yes | Nightscape Headband (8176, +0.00 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -9.44 DPS) [crafted]; White Bandit Mask (10008, -9.56 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 22.0 | yes | Scout's Medallion (19537, -0.46 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.47 DPS) [dungeon]; Scout's Medallion (20442, -0.58 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 34.0 | yes | Nightscape Shoulders (8192, -0.70 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.76 DPS, sim-verified) [world_drop]; Mantle of Thieves (2264, -0.82 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | World drop [world_drop] | 16.0 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Darktide Cape (4114, -0.23 DPS) [quest]; Cloak of Night (4447, -0.23 DPS) [world] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 30.0 | yes | Tough Scorpid Breastplate (8203, +0.00 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.12 DPS) [crafted]; Hawkeye's Tunic (14592, -0.35 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.11 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.23 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 205.8 | yes | Dragonscale Gauntlets (8347, -0.40 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 193.8 | yes | Defiler's Leather Girdle (20192, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -9.91 DPS) [rep]; Defiler's Leather Girdle (20191, -9.91 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 42.0 | yes | Triprunner Dungarees (9624, -0.69 DPS, sim-verified) [quest]; Hawkeye's Breeches (14595, -0.82 DPS) [world]; Ferine Leggings (6690, -0.93 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 26.0 | yes | Dusky Boots (7390, -0.23 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted]; Imperial Leather Boots (6431, -0.32 DPS, sim-verified) [world_drop] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 20.0 | yes | Legionnaire's Band (19512, -0.23 DPS) [rep]; Disengagement Ring (276202, -0.23 DPS) [vendor]; Monkey Ring (6748, -0.35 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.0 | yes | Disengagement Ring (276202, -0.12 DPS) [vendor]; Legionnaire's Band (19512, -0.17 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.23 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 536.8 | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 536.3 | yes | Grayson's Torch (1172, -31.28 DPS) [quest]; Rod of Molten Fire (2565, -31.28 DPS) [world_drop]; Nightglow Concoction (3451, -31.28 DPS) [quest] |
| ranged | Sniper Rifle (3430) | World drop [world_drop] | 0.0 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.15 DPS) [quest]; Master Hunter's Bow (17686, -1.37 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 1258, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 199.7. Weights run: 0.6s. Verify run: 1.3s. 2323 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=18.230 ± 3.490, hit=not significant (0.000 ± 0.000), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Cryptstalker Headpiece [quest] | 572.4 | yes | Champion's Chain Helm (23251, -1.48 DPS) [vendor]; Lieutenant Commander's Chain Helm (23306, -1.48 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -1.48 DPS) [pvp] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | Onyxia Tooth Pendant (18404, -0.57 DPS) [quest]; Choker of the Shifting Sands (21505, -13.54 DPS) [quest]; Amulet of the Darkmoon (19491, -13.77 DPS) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Cryptstalker Spaulders [quest] | 313.2 | yes | Champion's Chain Pauldrons (227078, -1.14 DPS) [pvp]; Lieutenant Commander's Chain Pauldrons (227084, -1.14 DPS) [pvp]; Champion's Chain Shoulders (23252, -1.26 DPS) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 0.0 | yes | Cape of the Black Baron (13340, -0.11 DPS) [dungeon]; Deathguard's Cloak (20068, -0.46 DPS) [rep]; Chromatic Cloak (18509, -2.91 DPS, sim-verified) [crafted] |
| chest | Bloodsoul Breastplate (19690) | Blacksmithing [crafted] | 528.4 | yes | Legionnaire's Chain Hauberk (22874, +0.00 DPS) [vendor]; Knight-Captain's Chain Hauberk (23292, +0.00 DPS) [vendor]; Legionnaire's Chain Hauberk (227071, +0.00 DPS) [pvp] |
| wrist | Cryptstalker Wristguards (22443) | Cryptstalker Wristguards [quest] | 52.0 | yes | Marshal's Chain Bracers (16461, -0.69 DPS) [pvp]; General's Chain Wristguards (16570, -0.69 DPS) [pvp]; Forest Stalker's Bracers (19587, -8.31 DPS, sim-verified) [rep] |
| hands | Cryptstalker Handguards (22441) | Cryptstalker Handguards [quest] | 303.2 | yes | Marshal's Chain Grips (16463, -0.34 DPS) [vendor]; General's Chain Gloves (16571, -0.34 DPS) [vendor]; Chromatic Gauntlets (19157, -9.76 DPS, sim-verified) [crafted] |
| waist | Cryptstalker Girdle (22442) | Cryptstalker Girdle [quest] | 301.2 | yes | Defiler's Leather Girdle (20190, -0.69 DPS) [rep]; Light Obsidian Belt (22195, -0.80 DPS) [crafted]; Defiler's Chain Girdle (20150, -10.54 DPS, sim-verified) [rep] |
| legs | Cryptstalker Legguards (22437) | Cryptstalker Legguards [quest] | 0.0 | yes | Legionnaire's Chain Legguards (22875, +0.00 DPS) [vendor]; Knight-Captain's Chain Legguards (23293, +0.00 DPS) [vendor]; Knight-Captain's Chain Legguards (227072, +0.00 DPS) [pvp] |
| feet | Cryptstalker Boots (22440) | Cryptstalker Boots [quest] | 66.0 | yes | General's Chain Sabatons (231564, -0.69 DPS) [pvp]; Scalegut Treaders (275618, -0.69 DPS) [crafted]; Striker's Footguards (21365, -6.24 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 271.2 | yes | Ring of Entropy (18543, -0.91 DPS) [world]; Mindtear Band (20632, -0.91 DPS) [world]; Band of the Penitent (13217, -4.22 DPS, sim-verified) [quest] |
| finger2 | Dragonslayer's Signet (18403) | For All To See [quest] | 0.0 | yes | Ring of Entropy (18543, +0.00 DPS) [world]; Mindtear Band (20632, +0.00 DPS) [world]; Band of the Penitent (13217, -2.69 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Tidal Charm (1404, -4.80 DPS) [vendor]; Guardian Talisman (1490, -4.80 DPS) [quest]; Blazing Emblem (2802, -4.80 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 1023.0 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | High Warlord's Recurve (234559, -3.10 DPS) [pvp]; High Warlord's Crossbow (234560, -3.10 DPS) [pvp]; High Warlord's Street Sweeper (234561, -3.10 DPS) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Medallion of the Dawn; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Bloodsoul Breastplate; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Cryptstalker Legguards; feet: Cryptstalker Boots; finger1: Don Julio's Band; finger2: Dragonslayer's Signet; trinket1: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 2323, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

