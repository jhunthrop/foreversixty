# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 32.2. Weights run: 1.6s. Verify run: 1.1s. 399 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.268 ± 0.163, crit=0.822 ± 0.037, hit=1.874 ± 0.118, spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 11.0 | yes | Totemic Leather Hood (252448, -0.19 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.38 DPS) [crafted]; Trapper's Leather Hood (252505, -1.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 10.1 | yes | Double-Stitched Woolen Shoulders (4314, -0.13 DPS, sim-verified) [crafted]; Rough Bronze Shoulders (3480, -0.97 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.97 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 0.0 | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.68 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 14.3 | yes | Acolyte's Chain Shirt (250491, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Armor (252493, -0.29 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.29 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.37 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 0.0 | yes | Acolyte's Gloves (250511, -0.07 DPS) [crafted]; Blight Gloves (279877, -0.11 DPS) [quest]; Fletcher's Gloves (7348, -0.95 DPS, sim-verified) [crafted] |
| waist | Stormrider's Leather Belt (252432) | Leatherworking [crafted] | 0.0 | yes | Pristine Sash (253925, +0.00 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.51 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.1 | yes | Stormrider's Leather Pants (252502, -0.15 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.31 DPS) [crafted]; Dreamer's Leggings (270016, -0.35 DPS, sim-verified) [quest] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.3 | yes | Spidersilk Boots (4320, +0.00 DPS, sim-verified) [crafted]; Acolyte's Boots (250506, -0.19 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.19 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 7.6 | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Lorekeeper's Ring (20431, -0.25 DPS) [rep]; Sludge-Stained Band (286535, -0.44 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.5 | yes | Loop of Sacrifice (281673, -0.24 DPS, sim-verified) [quest]; Lorekeeper's Ring (20431, -0.24 DPS) [rep]; Sludge-Stained Band (286535, -0.43 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 225.9 | yes | Living Root (6631, -0.45 DPS) [dungeon]; Staff of Westfall (2042, -0.55 DPS) [quest]; Gnarled Necromancer's Staff (251534, -1.59 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Kajaric Icon (206387) (or Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Rune Broker [vendor] | 0.0 | yes | Dyadic Icon (206381, +0.00 DPS) [vendor]; Tempest Icon (206382, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS, sim-verified) [crafted] |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Stormrider's Leather Belt; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Kajaric Icon

No-known-source sample (15 of 399, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9753 Nomad Buckler; 9756 Nomad Trousers; 9757 Nomad Tunic; 9763 Cadet Leggings

### Band 30 (dwarf, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 62.7. Weights run: 1.7s. Verify run: 1.3s. 821 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.138 ± 0.340), crit=1.742 ± 0.105, hit=2.217 ± 0.197, spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 18.2 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.42 DPS) [crafted]; Nightsky Cowl (4039, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 | yes | Crystal Starfire Medallion (5003, -0.93 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.28 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.38 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.2 | yes | Fairywing Mantle (9536, -0.30 DPS) [quest]; Feline Mantle (3748, -0.61 DPS) [dungeon]; Death Speaker Mantle (6685, -0.70 DPS, sim-verified) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.1 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.15 DPS) [quest]; Pearl-clasped Cloak (5542, -0.37 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.8 | yes | Guardian Armor (4256, -0.23 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.43 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.70 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.22 DPS) [quest]; Technician's Bracers (270042, -0.22 DPS) [quest]; Nightsky Wristbands (6407, -1.23 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 0.0 | yes | Stormrider's Leather Gloves (252498, -0.70 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.74 DPS) [dungeon]; Fletcher's Gloves (7348, -1.00 DPS, sim-verified) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 0.0 | yes | Highlander's Cloth Girdle (20099, -0.05 DPS) [rep]; Guardian Belt (4258, -0.10 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.88 DPS, sim-verified) [crafted] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 22.5 | yes | Guardian Pants (5962, -0.54 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.54 DPS) [crafted]; Abomination Skin Leggings (23173, -0.63 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 | yes | Stormrider's Leather Boots (252443, -0.33 DPS) [crafted]; Spidersilk Boots (4320, -0.34 DPS) [crafted]; Acidic Walkers (9454, -1.42 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.0 | yes | Minor Channeling Ring (1449, -0.07 DPS) [quest]; Lorekeeper's Ring (19525, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 | yes | Minor Channeling Ring (1449, +0.00 DPS, sim-verified) [quest]; Lorekeeper's Ring (19525, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 290.6 | yes | Cobalt Crusher (7730, -0.09 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.10 DPS) [vendor]; Corpsemaker (6687, -10.70 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Kajaric Icon (206387) (or Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Rune Broker [vendor] | 0.0 | yes | Dyadic Icon (206381, +0.00 DPS) [vendor]; Tempest Icon (206382, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS, sim-verified) [crafted] |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Prefect's Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 821, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 7958 Bronze Battle Axe; 9362 Brilliant Gold Ring

### Band 40 (dwarf, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 78.6. Weights run: 1.9s. Verify run: 1.7s. 1153 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.663 ± 0.634), crit=2.997 ± 0.195, hit=4.584 ± 0.401, spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Spellpower Goggles Xtreme (10502, +0.00 DPS, sim-verified) [crafted]; Big Voodoo Mask (8201, -2.11 DPS) [crafted]; Augural Shroud (2620, -2.16 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Darkspear Warding Pendant (272074, -0.56 DPS) [vendor]; Necklace of Calisea (1714, -0.61 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.68 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.18 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.18 DPS) [vendor]; Icy Cloak (4327, -0.21 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.36 DPS) [crafted]; Crimson Silk Vest (7058, -0.65 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.0 | yes | Arcane Runed Bracers (4744, -0.09 DPS) [quest]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon]; Turtle Scale Bracers (8198, -0.64 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Dragonscale Gauntlets (8347), Shadowskin Gloves (18238)) | World drop [world_drop] | 42.0 | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Fletcher's Gloves (7348, -0.79 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 0.0 | yes | Highlander's Cloth Girdle (20098, -0.03 DPS) [rep]; Skycaller's Leather Belt (252522, -0.26 DPS) [crafted]; Highlander's Chain Girdle (20089, -1.08 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Crimson Silk Pantaloons (7062, -0.47 DPS) [crafted]; Abomination Skin Leggings (23173, -0.68 DPS) [dungeon]; Kodohide Legguards (285338, -0.78 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -0.35 DPS, sim-verified) [crafted]; Skycaller's Mail Boots (252563, -0.57 DPS) [crafted]; Mender's Leather Shoes (252533, -1.01 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 | yes | Ring of Forlorn Spirits (2043, -0.53 DPS) [quest]; Reedknot Ring (9622, -0.62 DPS) [quest]; Minor Channeling Ring (1449, -0.68 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.18 DPS) [quest]; Lorekeeper's Ring (19525, -0.18 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.28 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Fiery War Axe (870, -7.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Kajaric Icon (206387) (or Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Rune Broker [vendor] | 0.0 | yes | Dyadic Icon (206381, +0.00 DPS) [vendor]; Tempest Icon (206382, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS, sim-verified) [crafted] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection

No-known-source sample (15 of 1153, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 50 (dwarf, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 97.3. Weights run: 1.9s. Verify run: 1.7s. 1494 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (3.110 ± 0.789), crit=3.848 ± 0.256, hit=6.081 ± 0.559, spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 113.0 | yes | Blood Guard's Pulsing Helmet (220848, -0.67 DPS, sim-verified) [vendor]; Blood Guard's Inscribed Skullcap (220842, -2.19 DPS) [vendor]; Red Mageweave Headband (10033, -2.80 DPS) [crafted] |
| neck | Horizon Choker (13085) | Azuregos [world] | 43.5 | yes | Darkspear Warding Pendant (272073, -0.36 DPS, sim-verified) [vendor]; Scorn's Icy Choker (23169, -1.58 DPS) [dungeon]; Mindburst Medallion (11196, -1.66 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 0.0 | yes | Blood Guard's Mail Epaulets (220823, -0.69 DPS) [vendor]; Blood Guard's Inscribed Shoulder Pads (220841, -0.69 DPS) [vendor]; Blood Guard's Pulsing Shoulders (220849, -3.68 DPS, sim-verified) [vendor] |
| back | Imperial Red Cloak (8248) | Maraudon: Princess Theradras [dungeon] | 0.0 | yes | Darkspear Raider's Cloak (272077, +0.00 DPS) [vendor]; Runecloth Cloak (13860, -0.03 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.25 DPS, sim-verified) [vendor] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Armor (220826, -1.72 DPS, sim-verified) [vendor]; Stone Guard's Inscribed Chestpiece (220838, -2.20 DPS) [vendor]; Knight's Mail Armor (223078, -2.20 DPS) [vendor] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 0.0 | yes | Imperial Red Bracers (8247, -0.27 DPS) [dungeon]; Bloodband Bracers (11469, -0.38 DPS) [quest]; Runic Leather Bracers (15092, -1.08 DPS, sim-verified) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 0.0 | yes | First Sergeant's Mail Gauntlets (220831, -0.53 DPS) [vendor]; First Sergeant's Inscribed Gauntlets (220843, -0.53 DPS) [vendor]; First Sergeant's Pulsing Gauntlets (220845, -3.23 DPS, sim-verified) [vendor] |
| waist | Highlander's Lizardhide Girdle (20103) (or Highlander's Mail Girdle (20118)) | The League of Arathor [rep] | 85.0 | yes | Highlander's Mail Girdle (20118, +0.00 DPS, sim-verified) [vendor]; Highlander's Cloth Girdle (20097, -0.58 DPS) [rep]; Dawnspire Cord (12466, -1.75 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Legplates (220834, -2.20 DPS) [vendor]; Stone Guard's Inscribed Legplates (220839, -2.20 DPS) [vendor]; Stormshroud Pants (15057, -2.22 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 102.9 | yes | Skycaller's Leather Boots (252471, +0.00 DPS, sim-verified) [crafted]; Skycaller's Mail Sabatons (252577, -4.73 DPS) [crafted]; Mender's Leather Boots (252472, -5.26 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 60.8 | yes | Ogremind Ring (1993, -3.44 DPS) [world_drop]; Voodoo Band (1996, -3.44 DPS) [world]; Mindbender Loop (5009, -3.44 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 28.7 | yes | Ogremind Ring (1993, +0.00 DPS, sim-verified) [world_drop]; Voodoo Band (1996, -0.61 DPS) [world]; Mindbender Loop (5009, -0.61 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -0.53 DPS) [quest]; Guardian Talisman (1490, -0.53 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 0.0 | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Kajaric Icon (206387) (or Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Rune Broker [vendor] | 0.0 | yes | Dyadic Icon (206381, +0.00 DPS) [vendor]; Tempest Icon (206382, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS, sim-verified) [crafted] |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Shizzle's Nozzle Wiper; hands: Raider Handwraps; waist: Highlander's Lizardhide Girdle; legs: Stone Guard's Pulsing Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket2: Uther's Strength; main_hand: Hammer of the Northern Wind

No-known-source sample (15 of 1494, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 60 (dwarf, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 134.8. Weights run: 1.9s. Verify run: 1.5s. 2304 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.046 ± 0.865), crit=5.588 ± 0.380, hit=6.652 ± 0.739, spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 0.0 | yes | Ragefury Eyepatch (11735, -4.75 DPS) [dungeon]; Bloodvine Lens (19998, -4.75 DPS) [crafted]; Mask of the Unforgiven (13404, -4.87 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 0.0 | yes | Medallion of the Dawn (22659, -5.77 DPS) [quest]; Beads of Ogre Might (22150, -6.78 DPS) [quest]; Charm of the Shifting Sands (21504, -9.29 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 108.8 | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [vendor]; Champion's Mail Spaulders (227160, -0.63 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 78.2 | yes | Earthweave Cloak (21187, +0.00 DPS, sim-verified) [quest]; Hide of the Wild (18510, -4.66 DPS) [crafted]; Spritecaster Cape (11623, -5.02 DPS) [dungeon] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 173.6 | yes | Bloodsoul Breastplate (19690, -1.49 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -2.50 DPS) [vendor]; Stormshroud Armor (15056, -4.73 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 93.5 | yes | Primal Batskin Bracers (19687, -2.74 DPS, sim-verified) [crafted]; Bindings of Elements (16671, -5.22 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -5.47 DPS) [rep] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 144.7 | yes | Blood Guard's Mail Vices (227159, +0.00 DPS) [vendor]; Primal Batskin Gloves (19686, -0.30 DPS, sim-verified) [crafted]; General's Mail Gauntlets (231660, -2.59 DPS) [vendor] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 115.0 | yes | Highlander's Cloth Girdle (20047, -1.11 DPS, sim-verified) [rep]; Highlander's Mail Girdle (20044, -1.64 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -1.64 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 156.5 | yes | Sentinel's Chain Leggings (22748, +0.00 DPS, sim-verified) [rep]; Legionnaire's Mail Legguards (227156, -1.02 DPS) [vendor]; Ironfeather Leggings (252486, -1.86 DPS) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 102.3 | yes | Greaves of Withering Despair (22240, -2.53 DPS, sim-verified) [dungeon]; Blood Guard's Mail Greaves (227158, -3.10 DPS) [vendor]; General's Mail Boots (16573, -5.59 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 144.7 | yes | Band of Earthen Might (21182, -2.48 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.32 DPS) [world]; Ritssyn's Ring of Chaos (21836, -3.60 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 0.0 | yes | Mindtear Band (20632, -0.28 DPS) [world]; Ritssyn's Ring of Chaos (21836, -0.57 DPS) [world]; Band of Earthen Might (21182, -4.73 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Fist of Cenarius (21188) | Champion's Battlegear [quest] | 0.0 | yes | High Warlord's Destroyer (234546, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (23199) (or Totem of the Storm (272432), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Mushgog [world] | 0.0 | yes | Kajaric Icon (206387, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS) [crafted]; Totem of the Storm (272432, +0.00 DPS, sim-verified) [world] |

**New at 60:** head: Bloodvine Goggles; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Chromatic Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Stormshroud Gloves; waist: Belt of the Archmage; legs: Stormshroud Pants; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket2: Thunderbrew's Boot Flask; main_hand: Fist of Cenarius; ranged: Totem of the Storm

No-known-source sample (15 of 2304, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

## Horde

### Band 20 (orc, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 31.3. Weights run: 1.6s. Verify run: 1.0s. 401 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.268 ± 0.163, crit=0.822 ± 0.037, hit=1.874 ± 0.118, spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 11.0 | yes | Totemic Leather Hood (252448, -0.19 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.38 DPS) [crafted]; Trapper's Leather Hood (252505, -1.25 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 10.1 | yes | Double-Stitched Woolen Shoulders (4314, -0.42 DPS, sim-verified) [crafted]; Rough Bronze Shoulders (3480, -0.97 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.97 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.8 | yes | Feyscale Cloak (6632, -0.27 DPS) [dungeon]; Black Whelp Cloak (7283, -0.27 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.51 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 14.3 | yes | Acolyte's Chain Shirt (250491, -0.20 DPS, sim-verified) [crafted]; Wisdom's Leather Armor (252493, -0.29 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.29 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.47 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 0.0 | yes | Acolyte's Gloves (250511, -0.07 DPS) [crafted]; Blight Gloves (279877, -0.11 DPS) [quest]; Fletcher's Gloves (7348, -0.34 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 10.1 | yes | Stormrider's Leather Belt (252432, +0.00 DPS, sim-verified) [crafted]; Pristine Sash (253925, -0.10 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.20 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.1 | yes | Stormrider's Leather Pants (252502, -0.05 DPS, sim-verified) [crafted]; Acolyte's Chain Leggings (250496, -0.31 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.3 | yes | Spidersilk Boots (4320, +0.00 DPS, sim-verified) [crafted]; Acolyte's Boots (250506, -0.19 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.19 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 7.6 | yes | Sludge-Stained Band (286535, -0.44 DPS) [world]; Black Pearl Ring (6332, -0.49 DPS) [world]; Loop of Sacrifice (281673, -0.88 DPS, sim-verified) [quest] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 0.0 | yes | Sludge-Stained Band (286535, -0.19 DPS) [world]; Black Pearl Ring (6332, -0.24 DPS) [world]; Loop of Sacrifice (281673, -0.41 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 225.9 | yes | Living Root (6631, -0.45 DPS) [dungeon]; Night Reaver (1318, -1.05 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -1.80 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Kajaric Icon (206387) (or Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Rune Broker [vendor] | 0.0 | yes | Dyadic Icon (206381, +0.00 DPS) [vendor]; Tempest Icon (206382, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS, sim-verified) [crafted] |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Pearl-clasped Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Keller's Girdle; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Kajaric Icon

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9753 Nomad Buckler; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 30 (orc, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 60.2. Weights run: 1.7s. Verify run: 1.3s. 825 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.138 ± 0.340), crit=1.742 ± 0.105, hit=2.217 ± 0.197, spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 18.2 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.42 DPS) [crafted]; Nightsky Cowl (4039, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 | yes | Crystal Starfire Medallion (5003, -0.93 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.19 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.38 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.2 | yes | Death Speaker Mantle (6685, -0.30 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.30 DPS) [quest]; Feline Mantle (3748, -0.61 DPS) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.1 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Soft Willow Cape (16661, -0.34 DPS) [quest]; Pearl-clasped Cloak (5542, -0.37 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.8 | yes | Guardian Armor (4256, -0.33 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.43 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.70 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.22 DPS) [quest]; Technician's Bracers (270042, -0.22 DPS) [quest]; Nightsky Wristbands (6407, -1.09 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 0.0 | yes | Oilrag Handwraps (16741, -0.04 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.21 DPS) [crafted]; Fletcher's Gloves (7348, -1.12 DPS, sim-verified) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 0.0 | yes | Defiler's Cloth Girdle (20164, -0.05 DPS) [rep]; Guardian Belt (4258, -0.10 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.76 DPS, sim-verified) [crafted] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 22.5 | yes | Guardian Pants (5962, -0.54 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.54 DPS) [crafted]; Abomination Skin Leggings (23173, -0.71 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 | yes | Stormrider's Leather Boots (252443, -0.33 DPS) [crafted]; Spidersilk Boots (4320, -0.34 DPS) [crafted]; Acidic Walkers (9454, -1.17 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.0 | yes | Advisor's Ring (19521, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.20 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 | yes | Advisor's Ring (19521, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.20 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 290.6 | yes | Cobalt Crusher (7730, -0.09 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.10 DPS) [vendor]; Corpsemaker (6687, -7.07 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Kajaric Icon (206387) (or Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Rune Broker [vendor] | 0.0 | yes | Dyadic Icon (206381, +0.00 DPS) [vendor]; Tempest Icon (206382, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS, sim-verified) [crafted] |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Prefect's Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 825, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 7958 Bronze Battle Axe

### Band 40 (orc, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 76.5. Weights run: 1.9s. Verify run: 1.6s. 1157 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.663 ± 0.634), crit=2.997 ± 0.195, hit=4.584 ± 0.401, spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 0.0 | yes | Big Voodoo Mask (8201, -0.24 DPS) [crafted]; Augural Shroud (2620, -0.30 DPS) [world]; Raging Berserker's Helm (7719, -0.79 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Necklace of Calisea (1714, -0.05 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.56 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.68 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 | yes | Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Green Silken Shoulders (7057, -0.12 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.18 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.18 DPS) [vendor]; Icy Cloak (4327, -0.21 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.36 DPS) [crafted]; Crimson Silk Vest (7058, -0.65 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.0 | yes | Radiant Silver Bracers (4545, -0.06 DPS) [quest]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon]; Turtle Scale Bracers (8198, -0.69 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Dragonscale Gauntlets (8347), Shadowskin Gloves (18238)) | World drop [world_drop] | 42.0 | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Fletcher's Gloves (7348, -0.68 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 0.0 | yes | Defiler's Cloth Girdle (20166, -0.03 DPS) [rep]; Skycaller's Leather Belt (252522, -0.26 DPS) [crafted]; Defiler's Chain Girdle (20153, -1.30 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Crimson Silk Pantaloons (7062, -0.47 DPS) [crafted]; Abomination Skin Leggings (23173, -0.68 DPS) [dungeon]; Kodohide Legguards (285338, -0.88 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -0.32 DPS, sim-verified) [crafted]; Skycaller's Mail Boots (252563, -0.57 DPS) [crafted]; Mender's Leather Shoes (252533, -1.01 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 | yes | Reedknot Ring (9622, -0.62 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.71 DPS) [vendor]; Ogremind Ring (1993, -0.83 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.18 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.27 DPS) [vendor]; Reedknot Ring (9622, -1.39 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Fiery War Axe (870, -1.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Kajaric Icon (206387) (or Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Rune Broker [vendor] | 0.0 | yes | Dyadic Icon (206381, +0.00 DPS) [vendor]; Tempest Icon (206382, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS, sim-verified) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection

No-known-source sample (15 of 1157, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 50 (orc, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 104.1. Weights run: 1.9s. Verify run: 1.7s. 1499 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (3.110 ± 0.789), crit=3.848 ± 0.256, hit=6.081 ± 0.559, spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 113.0 | yes | Blood Guard's Pulsing Helmet (220848, -0.96 DPS, sim-verified) [vendor]; Blood Guard's Inscribed Skullcap (220842, -2.19 DPS) [vendor]; Red Mageweave Headband (10033, -2.80 DPS) [crafted] |
| neck | Horizon Choker (13085) | Azuregos [world] | 43.5 | yes | Darkspear Warding Pendant (272073, -0.85 DPS, sim-verified) [vendor]; Scorn's Icy Choker (23169, -1.58 DPS) [dungeon]; Mindburst Medallion (11196, -1.66 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 0.0 | yes | Blood Guard's Mail Epaulets (220823, -0.69 DPS) [vendor]; Blood Guard's Inscribed Shoulder Pads (220841, -0.69 DPS) [vendor]; Blood Guard's Pulsing Shoulders (220849, -1.03 DPS, sim-verified) [vendor] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 43.5 | yes | Deep Woodlands Cloak (19121, +0.00 DPS, sim-verified) [quest]; Imperial Red Cloak (8248, -0.82 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -0.82 DPS) [vendor] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Armor (220826, -2.02 DPS, sim-verified) [vendor]; Stone Guard's Inscribed Chestpiece (220838, -2.20 DPS) [vendor]; Knight's Mail Armor (223078, -2.20 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 41.1 | yes | Imperial Red Bracers (8247, -0.61 DPS) [dungeon]; Bloodband Bracers (11469, -0.71 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.58 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 0.0 | yes | First Sergeant's Mail Gauntlets (220831, -0.53 DPS) [vendor]; First Sergeant's Inscribed Gauntlets (220843, -0.53 DPS) [vendor]; First Sergeant's Pulsing Gauntlets (220845, -2.76 DPS, sim-verified) [vendor] |
| waist | Highlander's Mail Girdle (20118) (or Defiler's Lizardhide Girdle (20174), Defiler's Mail Girdle (20196)) | Samuel Hawke [vendor] | 85.0 | yes | Defiler's Lizardhide Girdle (20174, +0.00 DPS, sim-verified) [rep]; Defiler's Mail Girdle (20196, +0.00 DPS) [rep]; Defiler's Cloth Girdle (20165, -0.58 DPS) [rep] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Legplates (220834, -2.20 DPS) [vendor]; Stone Guard's Inscribed Legplates (220839, -2.20 DPS) [vendor]; Stormshroud Pants (15057, -2.94 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 102.9 | yes | Skycaller's Leather Boots (252471, -0.15 DPS, sim-verified) [crafted]; Skycaller's Mail Sabatons (252577, -4.73 DPS) [crafted]; Mender's Leather Boots (252472, -5.26 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 60.8 | yes | Ogremind Ring (1993, -3.44 DPS) [world_drop]; Voodoo Band (1996, -3.44 DPS) [world]; Mindbender Loop (5009, -3.44 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 28.7 | yes | Ogremind Ring (1993, -0.42 DPS, sim-verified) [world_drop]; Voodoo Band (1996, -0.61 DPS) [world]; Mindbender Loop (5009, -0.61 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, -1.10 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.17 DPS, sim-verified) [world]; Guardian Talisman (1490, -3.75 DPS) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Kajaric Icon (206387) (or Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Rune Broker [vendor] | 0.0 | yes | Dyadic Icon (206381, +0.00 DPS) [vendor]; Tempest Icon (206382, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS, sim-verified) [crafted] |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Highlander's Mail Girdle; legs: Stone Guard's Pulsing Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket2: Rune of the Guard Captain

No-known-source sample (15 of 1499, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 60 (orc, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 133.6. Weights run: 1.9s. Verify run: 1.4s. 2300 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.046 ± 0.865), crit=5.588 ± 0.380, hit=6.652 ± 0.739, spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 0.0 | yes | Mask of the Unforgiven (13404, -2.79 DPS, sim-verified) [dungeon]; Ragefury Eyepatch (11735, -4.75 DPS) [dungeon]; Bloodvine Lens (19998, -4.75 DPS) [crafted] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 0.0 | yes | Medallion of the Dawn (22659, -5.77 DPS) [quest]; Beads of Ogre Might (22150, -6.78 DPS) [quest]; Charm of the Shifting Sands (21504, -9.29 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 108.8 | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [vendor]; Champion's Mail Spaulders (227160, -1.41 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 78.2 | yes | Earthweave Cloak (21187, +0.00 DPS, sim-verified) [quest]; Cloak of the Gathering Storm (21400, -4.39 DPS) [quest]; Hide of the Wild (18510, -4.66 DPS) [crafted] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 173.6 | yes | Bloodsoul Breastplate (19690, -1.49 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -2.50 DPS) [vendor]; Stormshroud Armor (15056, -6.26 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 93.5 | yes | Primal Batskin Bracers (19687, -2.55 DPS, sim-verified) [crafted]; Bindings of Elements (16671, -5.22 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -5.47 DPS) [rep] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 144.7 | yes | Primal Batskin Gloves (19686, +0.00 DPS, sim-verified) [crafted]; Blood Guard's Mail Vices (227159, +0.00 DPS) [vendor]; General's Mail Gauntlets (231660, -2.59 DPS) [vendor] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 115.0 | yes | Defiler's Cloth Girdle (20163, -1.43 DPS) [rep]; Highlander's Mail Girdle (20044, -1.64 DPS) [vendor]; Cord of The Five Thunders (227008, -2.65 DPS, sim-verified) [quest] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 0.0 | yes | Legionnaire's Mail Legguards (227156, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -0.84 DPS) [crafted]; Stormshroud Pants (15057, -2.11 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 102.3 | yes | Greaves of Withering Despair (22240, -2.90 DPS, sim-verified) [dungeon]; Blood Guard's Mail Greaves (227158, -3.10 DPS) [vendor]; General's Mail Boots (16573, -5.59 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 144.7 | yes | Band of Earthen Might (21182, -2.11 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.32 DPS) [world]; Ritssyn's Ring of Chaos (21836, -3.60 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 0.0 | yes | Mindtear Band (20632, -0.28 DPS) [world]; Ritssyn's Ring of Chaos (21836, -0.57 DPS) [world]; Band of Earthen Might (21182, -4.04 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, -1.38 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.79 DPS, sim-verified) [world]; Guardian Talisman (1490, -4.04 DPS) [quest] |
| main_hand | Fist of Cenarius (21188) | Champion's Battlegear [quest] | 0.0 | yes | High Warlord's Destroyer (234546, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (23199) (or Totem of the Storm (272432), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | Mushgog [world] | 0.0 | yes | Kajaric Icon (206387, +0.00 DPS) [vendor]; Polished Driftwood Icon (249398, +0.00 DPS) [crafted]; Totem of the Storm (272432, +0.00 DPS, sim-verified) [world] |

**New at 60:** head: Bloodvine Goggles; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Chromatic Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Stormshroud Gloves; waist: Belt of the Archmage; legs: Outrider's Chain Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; main_hand: Fist of Cenarius; ranged: Totem of the Storm

No-known-source sample (15 of 2300, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

