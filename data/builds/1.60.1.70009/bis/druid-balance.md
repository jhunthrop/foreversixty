# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 10 (night-elf, 1000000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 8.4. Weights run: 1.1s. Verify run: 0.3s. 688 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.001, intellect=0.058 ± 0.009, crit=0.605 ± 0.026, hit=0.323 ± 0.055, spell_haste=not significant (-0.016 ± 0.006), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.058 ± 0.001, arcane_power=0.942 ± 0.001

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Quests [quest] | 0.0 | yes |
| neck | - | - |  |  |
| shoulder | - | - |  |  |
| back | Fine Leather Cloak (2308) | Leatherworking [crafted] | 0.0 | yes |
| chest | Barbaric Linen Vest (2578) | Tailoring [crafted] | 2.0 | yes |
| wrist | - | - |  |  |
| hands | Fine Leather Gloves (2312) | Leatherworking [crafted] | 2.1 | yes |
| waist | Captain Sander's Sash (3344) | Quests [quest] | 0.0 | yes |
| legs | Embossed Leather Pants (4242) | Leatherworking [crafted] | 3.0 | yes |
| feet | Embossed Leather Boots (2309) | Leatherworking [crafted] | 2.0 | yes |
| finger1 | - | - |  |  |
| finger2 | - | - |  |  |
| trinket1 | - | - |  |  |
| trinket2 | - | - |  |  |
| main_hand | Goblin Smasher (4964) | Quests [quest] | 0.0 | no - runner-up Engineer's Hammer (id 5324) measured higher: 9.2 vs 8.4 set DPS - swapped in |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 10:** head: Lucky Fishing Hat; back: Fine Leather Cloak; chest: Barbaric Linen Vest; hands: Fine Leather Gloves; waist: Captain Sander's Sash; legs: Embossed Leather Pants; feet: Embossed Leather Boots; main_hand: Goblin Smasher

No-known-source sample (15 of 688, see the JSON for more): 816 Small Hand Blade; 821 Riverpaw Leather Vest; 1211 Gnoll War Harness; 1287 Giant Tarantula Fang; 1394 Driftwood Club; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1926 Weighted Sap; 1933 Staff of Conjuring; 1965 White Wolf Gloves; 2069 Black Bear Hide Vest; 2075 Priest's Mace; 2087 Hard Crawler Carapace; 2088 Long Crawler Limb; 2140 Carving Knife

### Band 15 (night-elf, 5100000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 4.9. Weights run: 1.2s. Verify run: 0.5s. 896 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.405 ± 0.069, crit=0.646 ± 0.028, hit=0.333 ± 0.080, spell_haste=not significant (-0.002 ± 0.012), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.107 ± 0.002, arcane_power=0.893 ± 0.001

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Flying Tiger Goggles (4368) | Engineering [crafted] | 0.0 | yes |
| neck | - | - |  |  |
| shoulder | - | - |  |  |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 3.2 | no - runner-up Black Whelp Cloak (id 7283) measured higher: 4.9 vs 4.9 set DPS - swapped in |
| chest | Moonglow Vest (6709) | Leatherworking [crafted] | 5.2 | yes |
| wrist | Crystalline Cuffs (14148) | Ragefire Chasm: Taragaman the Hungerer [dungeon] | 0.8 | yes |
| hands | Fine Leather Gloves (2312) | Leatherworking [crafted] | 2.8 | yes |
| waist | Captain Sander's Sash (3344) | Quests [quest] | 0.0 | no - runner-up Murloc Scale Belt (id 5780) measured higher: 4.9 vs 4.9 set DPS - swapped in |
| legs | Colorful Kilt (10048) | Tailoring [crafted] | 5.0 | yes |
| feet | Red Woolen Boots (4313) | Tailoring [crafted] | 4.0 | yes |
| finger1 | - | - |  |  |
| finger2 | - | - |  |  |
| trinket1 | - | - |  |  |
| trinket2 | - | - |  |  |
| main_hand | Goblin Screwdriver (1936) | The Deadmines [dungeon] | 0.0 | yes |
| off_hand | Grayson's Torch (1172) | Quests [quest] | 0.0 | yes |
| ranged | - | - |  |  |

**New at 15:** head: Flying Tiger Goggles; back: Pearl-clasped Cloak; chest: Moonglow Vest; wrist: Crystalline Cuffs; legs: Colorful Kilt; feet: Red Woolen Boots; main_hand: Goblin Screwdriver; off_hand: Grayson's Torch

No-known-source sample (15 of 896, see the JSON for more): 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 899 Venom Web Fang; 1189 Overseer's Ring; 1190 Overseer's Cloak; 1211 Gnoll War Harness; 1214 Gnoll Punisher; 1287 Giant Tarantula Fang; 1300 Lesser Staff of the Spire; 1355 Buckskin Cape; 1391 Riverpaw Mystic Staff; 1394 Driftwood Club; 1405 Foamspittle Staff

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 29.4. Weights run: 1.4s. Verify run: 0.5s. 1166 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.799 ± 0.049, crit=0.360 ± 0.025, hit=1.173 ± 0.088, spell_haste=not significant (-0.049 ± 0.051), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.430 ± 0.001, arcane_power=0.570 ± 0.003

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 4.0 | yes |
| neck | - | - |  |  |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | 0.0 | yes |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.4 | no - runner-up Heavy Woolen Cloak (id 4311) measured higher: 29.5 vs 29.4 set DPS - swapped in |
| chest | Gray Woolen Robe (2585) | Tailoring [crafted] | 8.0 | yes |
| wrist | Crystalline Cuffs (14148) | Ragefire Chasm: Taragaman the Hungerer [dungeon] | 1.6 | yes |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes |
| waist | Hillman's Belt (4250) | Leatherworking [crafted] | 4.0 | yes |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.4 | yes |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.2 | yes |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.8 | yes |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | 0.0 | yes |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 4.0 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 20:** head: Shadow Goggles; shoulder: Slime-encrusted Pads; chest: Gray Woolen Robe; hands: Serpent Gloves; waist: Hillman's Belt; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Deep Fathom Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Impaling Harpoon

No-known-source sample (15 of 1166, see the JSON for more): 789 Stout Battlehammer; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves; 899 Venom Web Fang; 911 Ironwood Treebranch; 920 Wicked Spiked Mace; 1121 Feet of the Lynx; 1189 Overseer's Ring; 1190 Overseer's Cloak; 1211 Gnoll War Harness

### Band 25 (night-elf, 5222211001000000-0000000000000000000-0000000000000000)

Set DPS (verified): 34.7. Weights run: 1.2s. Verify run: 0.6s. 1443 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.973 ± 0.097, crit=0.484 ± 0.032, hit=1.173 ± 0.092, spell_haste=not significant (-0.162 ± 0.066), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.355 ± 0.001, arcane_power=0.645 ± 0.003

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 4.9 | yes |
| neck | - | - |  |  |
| shoulder | Feline Mantle (3748) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 11.8 | yes |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.9 | no - runner-up Heavy Woolen Cloak (id 4311) measured higher: 35.6 vs 34.7 set DPS - swapped in |
| chest | Lesser Wizard's Robe (5766) | Tailoring [crafted] | 12.8 | yes |
| wrist | Crystalline Cuffs (14148) | Ragefire Chasm: Taragaman the Hungerer [dungeon] | 1.9 | no - runner-up Black Wolf Bracers (id 3230) measured higher: 34.7 vs 34.7 set DPS - swapped in |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 7.9 | yes |
| waist | Moss Cinch (6911) | Blackfathom Deeps [dungeon] | 12.0 | no - runner-up Belt of Arugal (id 6392) measured higher: 35.0 vs 34.7 set DPS - swapped in |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.8 | no - runner-up Gaze Dreamer Pants (id 6903) measured higher: 36.2 vs 34.7 set DPS - swapped in |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 6.8 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.8 | no - runner-up Silverlaine's Family Seal (id 6321) measured higher: 34.7 vs 34.7 set DPS - swapped in |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 4.9 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 25:** shoulder: Feline Mantle; chest: Lesser Wizard's Robe; hands: Truefaith Gloves; waist: Moss Cinch; finger1: Snake Hoop; finger2: Lavishly Jeweled Ring

No-known-source sample (15 of 1443, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves; 897 Madwolf Bracers; 899 Venom Web Fang; 911 Ironwood Treebranch; 920 Wicked Spiked Mace; 1076 Defias Renegade Ring; 1077 Defias Mage Ring

### Band 30 (night-elf, 5222211005100000-0000000000000000000-0000000000000000)

Set DPS (verified): 45.3. Weights run: 1.3s. Verify run: 0.7s. 1695 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.929 ± 0.129, crit=0.936 ± 0.076, hit=1.922 ± 0.153, spell_haste=not significant (-0.004 ± 0.042), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.426 ± 0.001, arcane_power=0.574 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.3 | yes |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.6 | yes |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.4 | yes |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.8 | no - runner-up Heavy Woolen Cloak (id 4311) measured higher: 46.0 vs 45.3 set DPS - swapped in |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.1 | yes |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 13.1 | no - runner-up Truefaith Gloves (id 7049) measured higher: 46.7 vs 45.3 set DPS - swapped in |
| waist | Guardian Belt (4258) | Leatherworking [crafted] | 12.5 | yes |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.4 | yes |
| feet | Acidic Walkers (9454) | Gnomeregan [dungeon] | 12.4 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 6.5 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.6 | no - runner-up Electrocutioner Lagnut (id 9447) measured higher: 46.3 vs 45.3 set DPS - swapped in |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 4.6 | no - runner-up Golden Iron Destroyer (id 3852) measured higher: 46.6 vs 45.3 set DPS - swapped in |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Fletcher's Gloves; waist: Guardian Belt; feet: Acidic Walkers

No-known-source sample (15 of 1695, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves; 897 Madwolf Bracers; 899 Venom Web Fang; 911 Ironwood Treebranch; 920 Wicked Spiked Mace

### Band 35 (night-elf, 5222211005501000-0000000000000000000-0000000000000000)

Set DPS (verified): 59.0. Weights run: 1.3s. Verify run: 0.7s. 1912 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.917 ± 0.184, crit=1.153 ± 0.126, hit=2.416 ± 0.206, spell_haste=not significant (0.246 ± 0.253), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.2 | yes |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.5 | yes |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | 18.1 | yes |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | 10.6 | yes |
| chest | Robe of Power (7054) | Tailoring [crafted] | 25.0 | yes |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.5 | no - runner-up Spidertank Oilrag (id 9448) measured higher: 59.3 vs 59.0 set DPS - swapped in |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 | yes |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.3 | yes |
| feet | Acidic Walkers (9454) | Gnomeregan [dungeon] | 12.3 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 6.4 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.5 | no - runner-up Electrocutioner Lagnut (id 9447) measured higher: 59.4 vs 59.0 set DPS - swapped in |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 21.6 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 35:** shoulder: Green Silken Shoulders; back: Long Silken Cloak; chest: Robe of Power; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Star Belt; main_hand: Illusionary Rod

No-known-source sample (15 of 1912, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves; 897 Madwolf Bracers; 899 Venom Web Fang; 911 Ironwood Treebranch

### Band 40 (night-elf, 5222211005501050-0000000000000000000-0000000000000000)

Set DPS (verified): 68.6. Weights run: 1.3s. Verify run: 0.8s. 2141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.824 ± 0.174, crit=1.375 ± 0.149, hit=3.180 ± 0.251, spell_haste=not significant (0.093 ± 0.283), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.481 ± 0.001, arcane_power=0.519 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.9 | yes |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.7 | yes |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | 10.1 | yes |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 25.4 | yes |
| wrist | Dryad's Wrist Bindings (19597) | Silverwing Sentinels [rep] | 20.9 | yes |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.4 | yes |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.9 | yes |
| feet | Southsea Mojo Boots (20641) | Quests [quest] | 17.1 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 5.8 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.9 | no - runner-up Electrocutioner Lagnut (id 9447) measured higher: 69.0 vs 68.6 set DPS - swapped in |
| trinket1 | Carrot on a Stick (11122) | Quests [quest] | 0.0 | yes |
| trinket2 | Mark of the Chosen (17774) | Quests [quest] | 0.0 | yes |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 24.2 | no - runner-up Spellshifter Rod (id 9527) measured higher: 71.5 vs 68.6 set DPS - swapped in |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; chest: Dreamweave Vest; wrist: Dryad's Wrist Bindings; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Southsea Mojo Boots; trinket1: Carrot on a Stick; trinket2: Mark of the Chosen

No-known-source sample (15 of 2141, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves

### Band 45 (night-elf, 5222211005501051-0000000000000000000-4000000000000000)

Set DPS (verified): 82.5. Weights run: 1.3s. Verify run: 0.8s. 2358 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.952 ± 0.214, crit=1.713 ± 0.201, hit=3.542 ± 0.315, spell_haste=not significant (0.667 ± 0.465), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.512 ± 0.001, arcane_power=0.488 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 38.0 | no - runner-up Dreamweave Circlet (id 10041) measured higher: 82.8 vs 82.5 set DPS - swapped in |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 | yes |
| shoulder | Red Mageweave Shoulders (10029) | Tailoring [crafted] | 21.3 | yes |
| back | Big Voodoo Cloak (8216) | Leatherworking [crafted] | 13.6 | yes |
| chest | Acumen Robes (17775) | Quests [quest] | 38.0 | no - runner-up Feathered Breastplate (id 8349) measured higher: 83.4 vs 82.5 set DPS - swapped in |
| wrist | Dryad's Wrist Bindings (19597) | Silverwing Sentinels [rep] | 21.7 | yes |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 23.5 | yes |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.4 | yes |
| feet | Southsea Mojo Boots (20641) | Quests [quest] | 18.5 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 6.7 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.7 | no - runner-up Electrocutioner Lagnut (id 9447) measured higher: 82.6 vs 82.5 set DPS - swapped in |
| trinket1 | Shard of the Splithooves (10659) | Quests [quest] | 0.0 | yes |
| trinket2 | Demon's Blood (10779) | Quests [quest] | 0.0 | yes |
| main_hand | Blight (7959) | Blacksmithing [crafted] | 0.0 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 45:** head: Red Mageweave Headband; shoulder: Red Mageweave Shoulders; back: Big Voodoo Cloak; chest: Acumen Robes; waist: Satyrmane Sash; trinket1: Shard of the Splithooves; trinket2: Demon's Blood; main_hand: Blight

No-known-source sample (15 of 2358, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves

### Band 50 (night-elf, 5222211005501051-0000000000000000000-5400000000000000)

Set DPS (verified): 88.4. Weights run: 1.4s. Verify run: 0.9s. 2632 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.909 ± 0.180, crit=2.006 ± 0.234, hit=4.694 ± 0.379, spell_haste=not significant (0.714 ± 0.432), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.549 ± 0.002, arcane_power=0.451 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 45.4 | no - runner-up Red Mageweave Headband (id 10033) measured higher: 90.1 vs 88.4 set DPS - swapped in |
| neck | Archlight Talisman (15856) | Quests [quest] | 21.1 | yes |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 33.2 | yes |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 | yes |
| chest | Acumen Robes (17775) | Quests [quest] | 37.2 | no - runner-up Feathered Breastplate (id 8349) measured higher: 90.0 vs 88.4 set DPS - swapped in |
| wrist | Dryad's Wrist Bindings (19596) | Silverwing Sentinels [rep] | 25.5 | yes |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 28.1 | yes |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 23.1 | yes |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 56.2 | no - runner-up Red Mageweave Pants (id 10009) measured higher: 91.1 vs 88.4 set DPS - swapped in |
| feet | Southsea Mojo Boots (20641) | Quests [quest] | 18.0 | no - runner-up Black Mageweave Boots (id 10026) measured higher: 88.8 vs 88.4 set DPS - swapped in |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.9 | yes |
| finger2 | Snake Hoop (6750) | Quests [quest] | 6.4 | yes |
| trinket1 | Shard of the Splithooves (10659) | Quests [quest] | 0.0 | yes |
| trinket2 | Demon's Blood (10779) | Quests [quest] | 0.0 | yes |
| main_hand | Blight (7959) | Blacksmithing [crafted] | 0.0 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 50:** head: Eye of Theradras; neck: Archlight Talisman; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; wrist: Dryad's Wrist Bindings; hands: Fletcher's Gloves; legs: Stormshroud Pants; finger1: Blackstone Ring; finger2: Snake Hoop

No-known-source sample (15 of 2632, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 810 Hammer of the Northern Wind; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff

### Band 55 (night-elf, 5222211005501051-0000000000000000000-5531000000000000)

Set DPS (verified): 92.7. Weights run: 1.4s. Verify run: 0.9s. 2882 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.096 ± 0.146, crit=2.231 ± 0.247, hit=4.839 ± 0.391, spell_haste=not significant (0.318 ± 0.379), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.538 ± 0.002, arcane_power=0.462 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 128.0 | yes |
| neck | Archlight Talisman (15856) | Quests [quest] | 23.0 | no - runner-up Chains of the Lich (id 23125) measured higher: 94.1 vs 92.7 set DPS - swapped in |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 36.9 | yes |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.6 | yes |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | 62.5 | no - runner-up Acumen Robes (id 17775) measured higher: 97.2 vs 92.7 set DPS - swapped in |
| wrist | Dryad's Wrist Bindings (19596) | Silverwing Sentinels [rep] | 26.6 | yes |
| hands | Dreadmist Wraps (16705) | Scholomance: Lorekeeper Polkelt [dungeon] | 58.3 | yes |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 26.3 | yes |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 62.5 | no - runner-up Luminary Kilt (id 11823) measured higher: 98.8 vs 92.7 set DPS - swapped in |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 33.2 | no - runner-up Magister's Boots (id 16682) measured higher: 92.9 vs 92.7 set DPS - swapped in |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 48.4 | yes |
| finger2 | Band of the Penitent (13217) | Quests [quest] | 31.2 | no - runner-up Glowing Crystal Ring (id 18402) measured higher: 94.4 vs 92.7 set DPS - swapped in |
| trinket1 | Shard of the Splithooves (10659) | Quests [quest] | 0.0 | yes |
| trinket2 | Smokey's Lighter (13171) | Quests [quest] | 0.0 | yes |
| main_hand | Enchanted Battlehammer (12776) | Blacksmithing [crafted] | 96.8 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 55:** head: Mask of the Unforgiven; chest: Stormshroud Armor; hands: Dreadmist Wraps; waist: Wisdom of the Timbermaw; feet: Omnicast Boots; finger2: Band of the Penitent; trinket2: Smokey's Lighter; main_hand: Enchanted Battlehammer

No-known-source sample (15 of 2882, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 810 Hammer of the Northern Wind; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff

### Band 60 (night-elf, 5222211005501051-0000000000000000000-5533300000000000)

Set DPS (verified): 188.3. Weights run: 1.4s. Verify run: 0.9s. 3456 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.609 ± 0.329, crit=2.540 ± 0.307, hit=6.381 ± 0.574, spell_haste=not significant (0.965 ± 0.611), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.558 ± 0.002, arcane_power=0.442 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 163.2 | no - runner-up Bloodvine Goggles (id 19999) measured higher: 190.2 vs 188.3 set DPS - swapped in |
| neck | Gem of Trapped Innocents (23057) | Naxxramas [raid] | 97.4 | yes |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 73.5 | no - runner-up Shroud of the Nathrezim (id 18720) measured higher: 189.3 vs 188.3 set DPS - swapped in |
| back | Earthweave Cloak (21187) | Quests [quest] | 63.8 | yes |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 175.5 | no - runner-up Genesis Vest (id 21357) measured higher: 190.3 vs 188.3 set DPS - swapped in |
| wrist | Rockfury Bracers (21186) | Quests [quest] | 90.8 | yes |
| hands | Primal Batskin Gloves (19686) | Leatherworking [crafted] | 127.6 | no - runner-up Dark Storm Gauntlets (id 21585) measured higher: 192.8 vs 188.3 set DPS - swapped in |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 99.4 | no - runner-up Belt of the Archmage (id 18405) measured higher: 189.0 vs 188.3 set DPS - swapped in |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 110.5 | no - runner-up Genesis Trousers (id 21356) measured higher: 192.4 vs 188.3 set DPS - swapped in |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 108.6 | no - runner-up Boots of Epiphany (id 21600) measured higher: 190.4 vs 188.3 set DPS - swapped in |
| finger1 | Seal of the Damned (23025) | Naxxramas [raid] | 120.4 | yes |
| finger2 | Ring of the Fallen God (21709) | Quests [quest] | 110.5 | yes |
| trinket1 | The Restrained Essence of Sapphiron (23046) | Naxxramas [raid] | 40.0 | yes |
| trinket2 | Eye of the Dead (23047) | Naxxramas [raid] | 0.0 | yes |
| main_hand | Atiesh, Greatstaff of the Guardian (22589) | Quests [quest] | 329.1 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 60:** neck: Gem of Trapped Innocents; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Primal Batskin Gloves; waist: Belt of Never-ending Agony; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Seal of the Damned; finger2: Ring of the Fallen God; trinket1: The Restrained Essence of Sapphiron; trinket2: Eye of the Dead; main_hand: Atiesh, Greatstaff of the Guardian

No-known-source sample (15 of 3456, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 810 Hammer of the Northern Wind; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff

## Horde

### Band 10 (tauren, 1000000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 8.7. Weights run: 1.1s. Verify run: 0.3s. 688 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.001, intellect=0.058 ± 0.009, crit=0.605 ± 0.026, hit=0.323 ± 0.055, spell_haste=not significant (-0.016 ± 0.006), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.058 ± 0.001, arcane_power=0.942 ± 0.001

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Quests [quest] | 0.0 | yes |
| neck | - | - |  |  |
| shoulder | - | - |  |  |
| back | Fine Leather Cloak (2308) | Leatherworking [crafted] | 0.0 | yes |
| chest | Barbaric Linen Vest (2578) | Tailoring [crafted] | 2.0 | yes |
| wrist | - | - |  |  |
| hands | Fine Leather Gloves (2312) | Leatherworking [crafted] | 2.1 | yes |
| waist | Captain Sander's Sash (3344) | Quests [quest] | 0.0 | yes |
| legs | Embossed Leather Pants (4242) | Leatherworking [crafted] | 3.0 | yes |
| feet | Embossed Leather Boots (2309) | Leatherworking [crafted] | 2.0 | yes |
| finger1 | - | - |  |  |
| finger2 | - | - |  |  |
| trinket1 | - | - |  |  |
| trinket2 | - | - |  |  |
| main_hand | Goblin Smasher (4964) | Quests [quest] | 0.0 | no - runner-up Engineer's Hammer (id 5324) measured higher: 9.6 vs 8.7 set DPS - swapped in |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 10:** head: Lucky Fishing Hat; back: Fine Leather Cloak; chest: Barbaric Linen Vest; hands: Fine Leather Gloves; waist: Captain Sander's Sash; legs: Embossed Leather Pants; feet: Embossed Leather Boots; main_hand: Goblin Smasher

No-known-source sample (15 of 688, see the JSON for more): 816 Small Hand Blade; 821 Riverpaw Leather Vest; 1211 Gnoll War Harness; 1287 Giant Tarantula Fang; 1394 Driftwood Club; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1926 Weighted Sap; 1933 Staff of Conjuring; 1965 White Wolf Gloves; 2069 Black Bear Hide Vest; 2075 Priest's Mace; 2087 Hard Crawler Carapace; 2088 Long Crawler Limb; 2140 Carving Knife

### Band 15 (tauren, 5100000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 4.9. Weights run: 1.2s. Verify run: 0.5s. 896 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.405 ± 0.069, crit=0.646 ± 0.028, hit=0.333 ± 0.080, spell_haste=not significant (-0.002 ± 0.012), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.107 ± 0.002, arcane_power=0.893 ± 0.001

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Flying Tiger Goggles (4368) | Engineering [crafted] | 0.0 | yes |
| neck | - | - |  |  |
| shoulder | - | - |  |  |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 3.2 | no - runner-up Black Whelp Cloak (id 7283) measured higher: 4.9 vs 4.9 set DPS - swapped in |
| chest | Moonglow Vest (6709) | Leatherworking [crafted] | 5.2 | yes |
| wrist | Crystalline Cuffs (14148) | Ragefire Chasm: Taragaman the Hungerer [dungeon] | 0.8 | yes |
| hands | Fine Leather Gloves (2312) | Leatherworking [crafted] | 2.8 | yes |
| waist | Captain Sander's Sash (3344) | Quests [quest] | 0.0 | no - runner-up Murloc Scale Belt (id 5780) measured higher: 4.9 vs 4.9 set DPS - swapped in |
| legs | Colorful Kilt (10048) | Tailoring [crafted] | 5.0 | yes |
| feet | Red Woolen Boots (4313) | Tailoring [crafted] | 4.0 | yes |
| finger1 | - | - |  |  |
| finger2 | - | - |  |  |
| trinket1 | - | - |  |  |
| trinket2 | - | - |  |  |
| main_hand | Goblin Screwdriver (1936) | The Deadmines [dungeon] | 0.0 | yes |
| off_hand | Grayson's Torch (1172) | Quests [quest] | 0.0 | yes |
| ranged | - | - |  |  |

**New at 15:** head: Flying Tiger Goggles; back: Pearl-clasped Cloak; chest: Moonglow Vest; wrist: Crystalline Cuffs; legs: Colorful Kilt; feet: Red Woolen Boots; main_hand: Goblin Screwdriver; off_hand: Grayson's Torch

No-known-source sample (15 of 896, see the JSON for more): 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 899 Venom Web Fang; 1189 Overseer's Ring; 1190 Overseer's Cloak; 1211 Gnoll War Harness; 1214 Gnoll Punisher; 1287 Giant Tarantula Fang; 1300 Lesser Staff of the Spire; 1355 Buckskin Cape; 1391 Riverpaw Mystic Staff; 1394 Driftwood Club; 1405 Foamspittle Staff

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 29.4. Weights run: 1.4s. Verify run: 0.5s. 1166 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.799 ± 0.049, crit=0.360 ± 0.025, hit=1.173 ± 0.088, spell_haste=not significant (-0.049 ± 0.051), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.430 ± 0.001, arcane_power=0.570 ± 0.003

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 4.0 | yes |
| neck | - | - |  |  |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | 0.0 | yes |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.4 | yes |
| chest | Gray Woolen Robe (2585) | Tailoring [crafted] | 8.0 | yes |
| wrist | Crystalline Cuffs (14148) | Ragefire Chasm: Taragaman the Hungerer [dungeon] | 1.6 | yes |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes |
| waist | Hillman's Belt (4250) | Leatherworking [crafted] | 4.0 | yes |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.4 | yes |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.2 | yes |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.8 | yes |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | 0.0 | yes |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 4.0 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 20:** head: Shadow Goggles; shoulder: Slime-encrusted Pads; chest: Gray Woolen Robe; hands: Serpent Gloves; waist: Hillman's Belt; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Deep Fathom Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Impaling Harpoon

No-known-source sample (15 of 1166, see the JSON for more): 789 Stout Battlehammer; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves; 899 Venom Web Fang; 911 Ironwood Treebranch; 920 Wicked Spiked Mace; 1121 Feet of the Lynx; 1189 Overseer's Ring; 1190 Overseer's Cloak; 1211 Gnoll War Harness

### Band 25 (tauren, 5222211001000000-0000000000000000000-0000000000000000)

Set DPS (verified): 35.3. Weights run: 1.2s. Verify run: 0.6s. 1443 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.973 ± 0.097, crit=0.484 ± 0.032, hit=1.173 ± 0.092, spell_haste=not significant (-0.162 ± 0.066), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.355 ± 0.001, arcane_power=0.645 ± 0.003

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 4.9 | yes |
| neck | - | - |  |  |
| shoulder | Feline Mantle (3748) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 11.8 | yes |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.9 | no - runner-up Heavy Woolen Cloak (id 4311) measured higher: 36.0 vs 35.3 set DPS - swapped in |
| chest | Lesser Wizard's Robe (5766) | Tailoring [crafted] | 12.8 | yes |
| wrist | Crystalline Cuffs (14148) | Ragefire Chasm: Taragaman the Hungerer [dungeon] | 1.9 | yes |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 7.9 | yes |
| waist | Moss Cinch (6911) | Blackfathom Deeps [dungeon] | 12.0 | no - runner-up Belt of Arugal (id 6392) measured higher: 35.6 vs 35.3 set DPS - swapped in |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.8 | no - runner-up Gaze Dreamer Pants (id 6903) measured higher: 36.3 vs 35.3 set DPS - swapped in |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 6.8 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.8 | yes |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 4.9 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 25:** shoulder: Feline Mantle; chest: Lesser Wizard's Robe; hands: Truefaith Gloves; waist: Moss Cinch; finger1: Snake Hoop; finger2: Lavishly Jeweled Ring

No-known-source sample (15 of 1443, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves; 897 Madwolf Bracers; 899 Venom Web Fang; 911 Ironwood Treebranch; 920 Wicked Spiked Mace; 1076 Defias Renegade Ring; 1077 Defias Mage Ring

### Band 30 (tauren, 5222211005100000-0000000000000000000-0000000000000000)

Set DPS (verified): 45.7. Weights run: 1.3s. Verify run: 0.7s. 1695 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.929 ± 0.129, crit=0.936 ± 0.076, hit=1.922 ± 0.153, spell_haste=not significant (-0.004 ± 0.042), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.426 ± 0.001, arcane_power=0.574 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.3 | yes |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.6 | yes |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.4 | yes |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.8 | no - runner-up Heavy Woolen Cloak (id 4311) measured higher: 46.4 vs 45.7 set DPS - swapped in |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.1 | yes |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 13.1 | no - runner-up Truefaith Gloves (id 7049) measured higher: 47.3 vs 45.7 set DPS - swapped in |
| waist | Guardian Belt (4258) | Leatherworking [crafted] | 12.5 | yes |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.4 | yes |
| feet | Acidic Walkers (9454) | Gnomeregan [dungeon] | 12.4 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 6.5 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.6 | no - runner-up Electrocutioner Lagnut (id 9447) measured higher: 46.5 vs 45.7 set DPS - swapped in |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 4.6 | no - runner-up Golden Iron Destroyer (id 3852) measured higher: 46.8 vs 45.7 set DPS - swapped in |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Fletcher's Gloves; waist: Guardian Belt; feet: Acidic Walkers

No-known-source sample (15 of 1695, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves; 897 Madwolf Bracers; 899 Venom Web Fang; 911 Ironwood Treebranch; 920 Wicked Spiked Mace

### Band 35 (tauren, 5222211005501000-0000000000000000000-0000000000000000)

Set DPS (verified): 59.1. Weights run: 1.3s. Verify run: 0.7s. 1912 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.917 ± 0.184, crit=1.153 ± 0.126, hit=2.416 ± 0.206, spell_haste=not significant (0.246 ± 0.253), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.2 | yes |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.5 | yes |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | 18.1 | yes |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | 10.6 | yes |
| chest | Robe of Power (7054) | Tailoring [crafted] | 25.0 | yes |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.5 | no - runner-up Spidertank Oilrag (id 9448) measured higher: 59.4 vs 59.1 set DPS - swapped in |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 | yes |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.3 | yes |
| feet | Acidic Walkers (9454) | Gnomeregan [dungeon] | 12.3 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 6.4 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.5 | no - runner-up Electrocutioner Lagnut (id 9447) measured higher: 59.2 vs 59.1 set DPS - swapped in |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 21.6 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 35:** shoulder: Green Silken Shoulders; back: Long Silken Cloak; chest: Robe of Power; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Star Belt; main_hand: Illusionary Rod

No-known-source sample (15 of 1912, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves; 897 Madwolf Bracers; 899 Venom Web Fang; 911 Ironwood Treebranch

### Band 40 (tauren, 5222211005501050-0000000000000000000-0000000000000000)

Set DPS (verified): 68.8. Weights run: 1.3s. Verify run: 0.8s. 2141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.824 ± 0.174, crit=1.375 ± 0.149, hit=3.180 ± 0.251, spell_haste=not significant (0.093 ± 0.283), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.481 ± 0.001, arcane_power=0.519 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.9 | yes |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.7 | yes |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | 10.1 | yes |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 25.4 | yes |
| wrist | Dryad's Wrist Bindings (19597) | Silverwing Sentinels [rep] | 20.9 | yes |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.4 | yes |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.9 | yes |
| feet | Southsea Mojo Boots (20641) | Quests [quest] | 17.1 | yes |
| finger1 | Snake Hoop (6750) | Quests [quest] | 5.8 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.9 | no - runner-up Electrocutioner Lagnut (id 9447) measured higher: 69.0 vs 68.8 set DPS - swapped in |
| trinket1 | Carrot on a Stick (11122) | Quests [quest] | 0.0 | yes |
| trinket2 | Mark of the Chosen (17774) | Quests [quest] | 0.0 | yes |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 24.2 | no - runner-up Spellshifter Rod (id 9527) measured higher: 71.7 vs 68.8 set DPS - swapped in |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; chest: Dreamweave Vest; wrist: Dryad's Wrist Bindings; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Southsea Mojo Boots; trinket1: Carrot on a Stick; trinket2: Mark of the Chosen

No-known-source sample (15 of 2141, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves

### Band 45 (tauren, 5222211005501051-0000000000000000000-4000000000000000)

Set DPS (verified): 82.0. Weights run: 1.3s. Verify run: 0.8s. 2358 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.952 ± 0.214, crit=1.713 ± 0.201, hit=3.542 ± 0.315, spell_haste=not significant (0.667 ± 0.465), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.512 ± 0.001, arcane_power=0.488 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 38.0 | no - runner-up Dreamweave Circlet (id 10041) measured higher: 83.2 vs 82.0 set DPS - swapped in |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 | yes |
| shoulder | Red Mageweave Shoulders (10029) | Tailoring [crafted] | 21.3 | no - runner-up Inquisitor's Shawl (id 19507) measured higher: 82.1 vs 82.0 set DPS - swapped in |
| back | Big Voodoo Cloak (8216) | Leatherworking [crafted] | 13.6 | no - runner-up Long Silken Cloak (id 4326) measured higher: 82.6 vs 82.0 set DPS - swapped in |
| chest | Acumen Robes (17775) | Quests [quest] | 38.0 | no - runner-up Feathered Breastplate (id 8349) measured higher: 83.0 vs 82.0 set DPS - swapped in |
| wrist | Dryad's Wrist Bindings (19597) | Silverwing Sentinels [rep] | 21.7 | yes |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 23.5 | yes |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.4 | no - runner-up Big Voodoo Pants (id 8202) measured higher: 82.2 vs 82.0 set DPS - swapped in |
| feet | Southsea Mojo Boots (20641) | Quests [quest] | 18.5 | no - runner-up Black Mageweave Boots (id 10026) measured higher: 82.9 vs 82.0 set DPS - swapped in |
| finger1 | Snake Hoop (6750) | Quests [quest] | 6.7 | yes |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.7 | no - runner-up Electrocutioner Lagnut (id 9447) measured higher: 82.8 vs 82.0 set DPS - swapped in |
| trinket1 | Shard of the Splithooves (10659) | Quests [quest] | 0.0 | yes |
| trinket2 | Demon's Blood (10779) | Quests [quest] | 0.0 | yes |
| main_hand | Blight (7959) | Blacksmithing [crafted] | 0.0 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 45:** head: Red Mageweave Headband; shoulder: Red Mageweave Shoulders; back: Big Voodoo Cloak; chest: Acumen Robes; waist: Satyrmane Sash; trinket1: Shard of the Splithooves; trinket2: Demon's Blood; main_hand: Blight

No-known-source sample (15 of 2358, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff; 892 Gnoll Casting Gloves

### Band 50 (tauren, 5222211005501051-0000000000000000000-5400000000000000)

Set DPS (verified): 89.3. Weights run: 1.4s. Verify run: 0.9s. 2632 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.909 ± 0.180, crit=2.006 ± 0.234, hit=4.694 ± 0.379, spell_haste=not significant (0.714 ± 0.432), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.549 ± 0.002, arcane_power=0.451 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 45.4 | no - runner-up Red Mageweave Headband (id 10033) measured higher: 90.8 vs 89.3 set DPS - swapped in |
| neck | Archlight Talisman (15856) | Quests [quest] | 21.1 | yes |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 33.2 | yes |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 | yes |
| chest | Acumen Robes (17775) | Quests [quest] | 37.2 | no - runner-up Feathered Breastplate (id 8349) measured higher: 89.7 vs 89.3 set DPS - swapped in |
| wrist | Dryad's Wrist Bindings (19596) | Silverwing Sentinels [rep] | 25.5 | yes |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 28.1 | yes |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 23.1 | yes |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 56.2 | no - runner-up Red Mageweave Pants (id 10009) measured higher: 91.5 vs 89.3 set DPS - swapped in |
| feet | Southsea Mojo Boots (20641) | Quests [quest] | 18.0 | yes |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.9 | yes |
| finger2 | Snake Hoop (6750) | Quests [quest] | 6.4 | yes |
| trinket1 | Shard of the Splithooves (10659) | Quests [quest] | 0.0 | yes |
| trinket2 | Demon's Blood (10779) | Quests [quest] | 0.0 | yes |
| main_hand | Blight (7959) | Blacksmithing [crafted] | 0.0 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 50:** head: Eye of Theradras; neck: Archlight Talisman; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; wrist: Dryad's Wrist Bindings; hands: Fletcher's Gloves; legs: Stormshroud Pants; finger1: Blackstone Ring; finger2: Snake Hoop

No-known-source sample (15 of 2632, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 810 Hammer of the Northern Wind; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff

### Band 55 (tauren, 5222211005501051-0000000000000000000-5531000000000000)

Set DPS (verified): 94.3. Weights run: 1.4s. Verify run: 0.9s. 2882 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.096 ± 0.146, crit=2.231 ± 0.247, hit=4.839 ± 0.391, spell_haste=not significant (0.318 ± 0.379), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.538 ± 0.002, arcane_power=0.462 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 128.0 | yes |
| neck | Archlight Talisman (15856) | Quests [quest] | 23.0 | no - runner-up Chains of the Lich (id 23125) measured higher: 94.6 vs 94.3 set DPS - swapped in |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 36.9 | yes |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.6 | yes |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | 62.5 | no - runner-up Acumen Robes (id 17775) measured higher: 97.4 vs 94.3 set DPS - swapped in |
| wrist | Dryad's Wrist Bindings (19596) | Silverwing Sentinels [rep] | 26.6 | yes |
| hands | Dreadmist Wraps (16705) | Scholomance: Lorekeeper Polkelt [dungeon] | 58.3 | yes |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 26.3 | yes |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 62.5 | no - runner-up Luminary Kilt (id 11823) measured higher: 98.7 vs 94.3 set DPS - swapped in |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 33.2 | yes |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 48.4 | yes |
| finger2 | Band of the Penitent (13217) | Quests [quest] | 31.2 | no - runner-up Glowing Crystal Ring (id 18402) measured higher: 94.8 vs 94.3 set DPS - swapped in |
| trinket1 | Shard of the Splithooves (10659) | Quests [quest] | 0.0 | yes |
| trinket2 | Smokey's Lighter (13171) | Quests [quest] | 0.0 | yes |
| main_hand | Enchanted Battlehammer (12776) | Blacksmithing [crafted] | 96.8 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 55:** head: Mask of the Unforgiven; chest: Stormshroud Armor; hands: Dreadmist Wraps; waist: Wisdom of the Timbermaw; feet: Omnicast Boots; finger2: Band of the Penitent; trinket2: Smokey's Lighter; main_hand: Enchanted Battlehammer

No-known-source sample (15 of 2882, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 810 Hammer of the Northern Wind; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff

### Band 60 (tauren, 5222211005501051-0000000000000000000-5533300000000000)

Set DPS (verified): 186.5. Weights run: 1.4s. Verify run: 0.9s. 3456 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.609 ± 0.329, crit=2.540 ± 0.307, hit=6.381 ± 0.574, spell_haste=not significant (0.965 ± 0.611), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.558 ± 0.002, arcane_power=0.442 ± 0.004

| Slot | Item | Source | Score | Verified |
|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 163.2 | no - runner-up Bloodvine Goggles (id 19999) measured higher: 189.1 vs 186.5 set DPS - swapped in |
| neck | Gem of Trapped Innocents (23057) | Naxxramas [raid] | 97.4 | yes |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 73.5 | no - runner-up Shroud of the Nathrezim (id 18720) measured higher: 188.6 vs 186.5 set DPS - swapped in |
| back | Earthweave Cloak (21187) | Quests [quest] | 63.8 | yes |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 175.5 | no - runner-up Genesis Vest (id 21357) measured higher: 189.5 vs 186.5 set DPS - swapped in |
| wrist | Rockfury Bracers (21186) | Quests [quest] | 90.8 | yes |
| hands | Primal Batskin Gloves (19686) | Leatherworking [crafted] | 127.6 | no - runner-up Dark Storm Gauntlets (id 21585) measured higher: 192.0 vs 186.5 set DPS - swapped in |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 99.4 | no - runner-up Belt of the Archmage (id 18405) measured higher: 188.0 vs 186.5 set DPS - swapped in |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 110.5 | no - runner-up Genesis Trousers (id 21356) measured higher: 190.3 vs 186.5 set DPS - swapped in |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 108.6 | no - runner-up Boots of Epiphany (id 21600) measured higher: 189.1 vs 186.5 set DPS - swapped in |
| finger1 | Seal of the Damned (23025) | Naxxramas [raid] | 120.4 | yes |
| finger2 | Ring of the Fallen God (21709) | Quests [quest] | 110.5 | no - runner-up Band of Forced Concentration (id 19403) measured higher: 187.0 vs 186.5 set DPS - swapped in |
| trinket1 | The Restrained Essence of Sapphiron (23046) | Naxxramas [raid] | 40.0 | yes |
| trinket2 | Slayer's Crest (23041) | Naxxramas [raid] | 0.0 | no - runner-up Eye of the Dead (id 23047) measured higher: 186.6 vs 186.5 set DPS - swapped in |
| main_hand | Atiesh, Greatstaff of the Guardian (22589) | Quests [quest] | 329.1 | yes |
| off_hand | - | - |  |  |
| ranged | - | - |  |  |

**New at 60:** neck: Gem of Trapped Innocents; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Primal Batskin Gloves; waist: Belt of Never-ending Agony; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Seal of the Damned; finger2: Ring of the Fallen God; trinket1: The Restrained Essence of Sapphiron; trinket2: Slayer's Crest; main_hand: Atiesh, Greatstaff of the Guardian

No-known-source sample (15 of 3456, see the JSON for more): 720 Brawler Gloves; 789 Stout Battlehammer; 791 Gnarled Ash Staff; 810 Hammer of the Northern Wind; 816 Small Hand Blade; 820 Slicer Blade; 821 Riverpaw Leather Vest; 827 Wicked Blackjack; 865 Leaden Mace; 866 Monk's Staff; 867 Gloves of Holy Might; 868 Ardent Custodian; 873 Staff of Jordan; 880 Staff of Horrors; 890 Twisted Chanter's Staff

