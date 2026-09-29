# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.5. Weights run: 0.7s. Verify run: 0.6s. 252 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.128 ± 0.138), crit=1.335 ± 0.054, hit=3.164 ± 0.119, spell_haste=1.014 ± 0.120, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -5.4) [crafted]; Flying Tiger Goggles (4368, -6.0) [crafted]; Lucky Fishing Hat (19972, -6.0) [quest] |
| neck | Sentinel's Medallion (20444) (or Tarnished Locket (279870)) | Silverwing Sentinels [rep] | 0.0 | yes | Tarnished Locket (279870, +0.0) [quest] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.5 | yes | Double-Stitched Woolen Shoulders (4314, -1.5) [crafted]; Slime-encrusted Pads (6461, -5.5) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -1.0) [dungeon]; Black Whelp Cloak (7283, -1.0) [crafted]; Caretaker's Cape (20428, -1.0) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 | yes | Gray Woolen Robe (2585, -1.0) [crafted]; Green Woolen Vest (2582, -1.6) [crafted]; Bloody Apron (6226, -1.6) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.6 | yes | Tabitha's Cuffs (251486, +0.1) [quest]; Bright Bracers (3647, -0.1) [dungeon]; Seer's Cuffs (3645, -0.5) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -1.0) [world]; Pristine Gloves (253913, -2.6) [crafted]; Heavy Woolen Gloves (4310, -4.7) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 | yes | Novice Arcanist's Sash (253885, -0.1) [crafted]; Novice Ardent's Sash (253887, -2.1) [crafted]; Keller's Girdle (2911, -3.5) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.0 | yes | Silk-threaded Trousers (1929, -3.0) [dungeon]; Filigreed Pristine Leggings (253937, -3.3) [crafted]; Colorful Kilt (10048, -5.0) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 | yes | Feather Padded Treads (285345, -2.5) [world]; Red Woolen Boots (4313, -3.5) [crafted]; Pristine Boots (253889, -4.1) [crafted] |
| finger1 | Minor Channeling Ring (1449) | Quests [quest] | 5.3 | yes | Sludge-Stained Band (286535, -2.3) [world]; Lavishly Jeweled Ring (1156, -4.5) [dungeon]; Black Pearl Ring (6332, -5.0) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -2.0) [world]; Lavishly Jeweled Ring (1156, -4.2) [dungeon]; Black Pearl Ring (6332, -4.7) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 1.3 | yes | Gnarled Necromancer's Staff (251534, +0.0) [quest]; Channeler's Staff (4437, -0.3) [world]; Lesser Staff of the Spire (1300, -0.5) [world] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Quests [quest] | 5.0 | yes | Cookie's Stirring Rod (5198, -2.0) [dungeon]; Sable Wand (7607, -2.0) [quest]; Torchlight Wand (5240, -3.0) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 252, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 61.7. Weights run: 0.6s. Verify run: 0.6s. 475 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.065 ± 0.228), crit=2.229 ± 0.108, hit=3.372 ± 0.161, spell_haste=1.051 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | Razorfen Downs: Withered Battle Boar [dungeon] | 11.0 | yes | Silk Headband (7050, -2.0) [crafted]; Embalmed Shroud (7691, -3.0) [dungeon]; Filigreed Pristine Circlet (253975, -3.0) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.4 | yes | Darkspear Warding Pendant (272075, -7.1) [vendor]; Crystal Starfire Medallion (5003, -7.1) [dungeon]; Pendant of Myzrael (4614, -7.4) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.6 | yes | Invoker's Mantle (215365, -2.3) [crafted]; Moonlit Amice (11884, -2.6) [quest]; Death Speaker Mantle (6685, -2.9) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, -1.0) [crafted]; Prelacy Cape (7004, -1.0) [quest]; Caretaker's Cape (19533, -1.0) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 9.8 | yes | Tree Bark Jacket (1486, +3.2) [dungeon]; Robes of Arcana (5770, -1.8) [crafted]; Death Speaker Robes (6682, -2.1) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -8.6) [dungeon]; Tabitha's Cuffs (251486, -8.6) [quest]; Mindthrust Bracers (1974, -8.7) [dungeon] |
| hands | Serpent Gloves (5970) (or Shilly Mitts (9609)) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Shilly Mitts (9609, +0.0) [quest]; Gnoll Casting Gloves (892, -1.0) [world]; Truefaith Gloves (7049, -1.8) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.2 | yes | Belt of Arugal (6392, -2.0) [dungeon]; Ghamoo-ra's Bind (6908, -3.2) [dungeon]; Invoker's Cord (215366, -3.9) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 | yes | Gaze Dreamer Pants (6903, +2.5) [dungeon]; Pristine Leggings (253987, -2.1) [crafted]; Silk-threaded Trousers (1929, -2.5) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.5 | yes | Spidersilk Boots (4320, -0.2) [crafted]; Nimbus Boots (6998, -1.5) [quest]; Acidic Walkers (9454, -1.9) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -1.9) [quest]; Lorekeeper's Ring (20431, -2.0) [rep]; Electrocutioner Lagnut (9447, -4.0) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Minor Channeling Ring (1449, -0.9) [quest]; Electrocutioner Lagnut (9447, -3.0) [dungeon]; Sludge-Stained Band (286535, -3.0) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 0.7 | yes | Twisted Chanter's Staff (890, -0.1) [dungeon]; Gnarled Necromancer's Staff (251534, -0.1) [quest]; Channeler's Staff (4437, -0.2) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | Gnomeregan: Dark Iron Agent [dungeon] | 5.0 | yes | Sizzle Stick (8071, +0.0) [quest]; Greater Mystic Wand (217287, +0.0) [crafted]; Spellcrafter Wand (6677, -1.0) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 90.9. Weights run: 0.7s. Verify run: 0.6s. 641 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.379 ± 0.386), crit=3.424 ± 0.185, hit=5.185 ± 0.341, spell_haste=2.354 ± 0.492, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -6.2) [world]; Holy Shroud (2721, -10.0) [dungeon]; Enchanter's Cowl (4322, -11.2) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 | yes | Necklace of Calisea (1714, -6.6) [dungeon]; Darkspear Warding Pendant (272074, -6.6) [vendor]; Darkspear Warding Pendant (272075, -7.4) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 | yes | Green Silken Shoulders (7057, -0.2) [crafted]; Inquisitor's Shawl (19507, -0.5) [dungeon]; Berylline Pads (4197, -1.6) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.9 | yes | Guardian Cloak (5965, +0.0) [crafted]; Icy Cloak (4327, -0.9) [crafted]; Caretaker's Cape (19532, -1.9) [rep] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 24.3 | yes | Dreamweave Vest (10021, -2.9) [crafted]; Robe of Power (7054, -5.7) [crafted]; Crimson Silk Vest (7058, -8.5) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Quests [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.0) [dungeon]; Condor Bracers (15864, -2.0) [quest]; Earthen Silk Cuffs (254019, -5.0) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.5 | yes | Black Mageweave Gloves (10003, -4.5) [crafted]; Red Mageweave Gloves (10018, -4.7) [crafted]; Gilded Handwraps (254021, -8.9) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.5 | yes | Star Belt (4329, -2.5) [crafted]; Deathmage Sash (10771, -2.8) [dungeon]; Highlander's Cloth Girdle (20099, -3.4) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.5 | yes | Crimson Silk Pantaloons (7062, -5.6) [crafted]; Abomination Skin Leggings (23173, -6.5) [dungeon]; Gaze Dreamer Pants (6903, -6.5) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -14.3) [crafted]; Spidersilk Boots (4320, -15.5) [crafted]; Acidic Walkers (9454, -16.0) [dungeon] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -2.0) [quest]; Lorekeeper's Ring (19525, -2.0) [rep]; Sea Giant's Toe Ring (274746, -3.0) [vendor] |
| finger2 | Ring of Forlorn Spirits (2043) | Quests [quest] | 8.0 | yes | Reedknot Ring (9622, -1.0) [quest]; Sea Giant's Toe Ring (274746, -2.0) [vendor]; Minor Channeling Ring (1449, -2.2) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 50.2 | yes | Windweaver Staff (7757, -44.5) [dungeon]; Staff of Jordan (873, -46.0) [dungeon]; Glimmering Staff (249392, -46.0) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Burning Sliver (5249) (or Twisted Nether Wand (249144)) | Quests [quest] | 6.0 | yes | Twisted Nether Wand (249144, +0.0) [crafted]; Fizzle's Zippy Lighter (6729, -0.9) [quest]; Wand of Eventide (5214, -1.0) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Burning Sliver

No-known-source sample (15 of 641, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 118.9. Weights run: 0.7s. Verify run: 0.7s. 848 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.825 ± 0.694), crit=5.572 ± 0.315, hit=7.662 ± 0.520, spell_haste=not significant (0.844 ± 0.797), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 55.5 | yes | Eye of Theradras (17715, +57.2) [dungeon]; Dreamweave Circlet (10041, -16.2) [crafted]; Bad Mojo Mask (9470, -17.2) [dungeon] |
| neck | Horizon Choker (13085) | Azuregos [world] | 25.5 | yes | Scorn's Icy Choker (23169, -7.6) [dungeon]; Mindburst Medallion (11196, -8.6) [quest]; Darkspear Warding Pendant (272073, -9.1) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 45.8 | yes | Red Mageweave Shoulders (10029, -11.5) [crafted]; Inquisitor's Shawl (19507, -15.1) [dungeon]; Green Silken Shoulders (7057, -17.8) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 25.5 | yes | Spritecaster Cape (11623, -0.6) [dungeon]; Runecloth Cloak (13860, -1.9) [crafted]; Big Voodoo Cloak (8216, -4.1) [crafted] |
| chest | Acumen Robes (17775) | Quests [quest] | 55.5 | yes | Runecloth Robe (13858, -13.5) [crafted]; Runecloth Tunic (13857, -18.4) [crafted]; Hibernal Robe (8113, -19.0) [dungeon] |
| wrist | Shizzle's Nozzle Wiper (11917) | Quests [quest] | 21.9 | yes | Bloodband Bracers (11469, -0.5) [quest]; Imperial Red Bracers (8247, -1.8) [dungeon]; Nethergeld Cuffs (254061, -2.1) [crafted] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 29.2 | yes | Runecloth Gloves (13863, -0.8) [crafted]; Stormcloth Gloves (10011, -3.4) [crafted]; Dreamweave Gloves (10019, -3.9) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 96.1 | yes | Dawnspire Cord (12466, -55.5) [world]; Deathmage Sash (10771, -61.8) [dungeon]; Satyrmane Sash (17755, -63.9) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 35.9 | yes | Kilt of the Atal'ai Prophet (10807, -0.1) [dungeon]; Crimson Silk Pantaloons (7062, -4.2) [crafted]; Imperial Red Pants (8251, -8.5) [dungeon] |
| feet | Southsea Mojo Boots (20641) | Quests [quest] | 28.1 | yes | Earthen Silk Slippers (254013, -4.1) [crafted]; Black Mageweave Boots (10026, -4.3) [crafted]; Vinerot Sandals (17748, -5.6) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 76.6 | yes | Voodoo Band (1996, -63.8) [world]; Mindbender Loop (5009, -63.8) [dungeon]; Black Widow Band (6199, -63.8) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | Boulderfist Magus [world] | 12.8 | yes | Voodoo Band (1996, +0.0) [world]; Mindbender Loop (5009, +0.0) [dungeon]; Black Widow Band (6199, +0.0) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Thunderbrew's Boot Flask (744, -6.0) [quest]; Tidal Charm (1404, -6.0) [vendor]; Guardian Talisman (1490, -6.0) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 99.9 | yes | Illusionary Rod (7713, -10.9) [dungeon]; Inventor's Focal Sword (17719, -21.9) [dungeon]; Spellshifter Rod (9527, -57.9) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Quests [quest] | 12.3 | yes | Fizzle's Zippy Lighter (6729, -2.8) [quest]; Cairnstone Sliver (9654, -3.2) [quest]; Lesser Eternal Wand (249232, -4.3) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Acumen Robes; wrist: Shizzle's Nozzle Wiper; hands: Red Mageweave Gloves; waist: Highlander's Cloth Girdle; feet: Southsea Mojo Boots; finger1: Blackstone Ring; finger2: Ogremind Ring; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Charged Lightning Rod

No-known-source sample (15 of 848, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 276.2. Weights run: 0.7s. Verify run: 0.7s. 1212 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.124 ± 1.054), crit=8.680 ± 0.479, hit=12.852 ± 0.837, spell_haste=not significant (1.136 ± 1.257), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Quests [quest] | 409.4 | yes | Bloodvine Goggles (19999, -30.8) [crafted]; Enigma Circlet (21347, -123.4) [quest]; Field Marshal's Coronet (16441, -252.8) [vendor] |
| neck | Gem of Trapped Innocents (23057) | Naxxramas [raid] | 258.9 | yes | Onyxia Tooth Pendant (18404, -8.9) [quest]; Fury of the Forgotten Swarm (21809, -8.9) [raid]; Stormrage's Talisman of Seething (23053, -15.9) [raid] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 140.1 | yes | Champion's Silk Mantle (23264, -2.2) [vendor]; Lieutenant Commander's Silk Mantle (23319, -2.2) [vendor]; Lieutenant Commander's Silk Mantle (227102, -2.2) [pvp] |
| back | Earthweave Cloak (21187) | Quests [quest] | 128.5 | yes | Chromatic Cloak (18509, -7.0) [crafted]; Drape of Vaulted Secrets (21415, -109.4) [quest]; Hide of the Wild (18510, -113.3) [crafted] |
| chest | Frostfire Robe (22496) | Quests [quest] | 300.4 | yes | Bloodvine Vest (19682, -14.7) [crafted]; Enigma Robes (21343, -137.0) [quest]; Robe of the Archmage (14152, -137.4) [crafted] |
| wrist | Rockfury Bracers (21186) | Quests [quest] | 155.5 | yes | Burrower Bracers (21611, -125.9) [raid]; Frostfire Bindings (22503, -126.7) [quest]; Dryad's Wrist Bindings (19595, -132.5) [rep] |
| hands | Dark Storm Gauntlets (21585) | Ahn'Qiraj [raid] | 167.4 | yes | Gloves of Spell Mastery (14146, +85.9) [crafted]; Sorcerer's Gloves (22066, -25.1) [quest]; Dreadmist Wraps (16705, -37.7) [dungeon] |
| waist | Frostfire Belt (22502) | Quests [quest] | 159.1 | yes | Belt of the Archmage (18405, -15.6) [crafted]; Highlander's Cloth Girdle (20047, -22.9) [rep]; Highlander's Cloth Girdle (20097, -28.0) [rep] |
| legs | Frostfire Leggings (22497) | Quests [quest] | 177.7 | yes | Bloodvine Leggings (19683, -11.5) [crafted]; Enigma Leggings (21346, -19.0) [quest]; Marshal's Silk Leggings (16442, -23.7) [vendor] |
| feet | Enigma Boots (21344) | Quests [quest] | 158.4 | yes | Frostfire Sandals (22500, -6.6) [quest]; Marshal's Silk Footwraps (16437, -7.1) [vendor]; General's Silk Boots (16539, -7.1) [vendor] |
| finger1 | Seal of the Damned (23025) | Naxxramas [raid] | 271.0 | yes | Band of Earthen Might (21182, -21.0) [quest]; Band of Unnatural Forces (23038, -21.0) [raid]; Ring of the Fallen God (21709, -104.8) [quest] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182), Band of Unnatural Forces (23038)) | Stormpike Guard [rep] | 250.0 | yes | Band of Earthen Might (21182, +0.0) [quest]; Band of Unnatural Forces (23038, +0.0) [raid]; Ring of the Fallen God (21709, -83.8) [quest] |
| trinket1 | The Restrained Essence of Sapphiron (23046) | Naxxramas [raid] | 40.0 | yes | Neltharion's Tear (19379, +261.0) [raid]; Drake Fang Talisman (19406, +217.0) [raid]; Kiss of the Spider (22954, +210.0) [raid] |
| trinket2 | Eye of the Dead (23047) | Naxxramas [raid] | 0.0 | yes | Neltharion's Tear (19379, +301.0) [raid]; Drake Fang Talisman (19406, +257.0) [raid]; Kiss of the Spider (22954, +250.0) [raid] |
| main_hand | Atiesh, Greatstaff of the Guardian (22589) | Quests [quest] | 411.0 | yes | Atiesh, Greatstaff of the Guardian (22630, -14.4) [quest]; Blessed Qiraji Acolyte Staff (21273, -28.4) [quest]; High Warlord's War Staff (234549, -84.1) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Dreambough Wand (249234, -1.0) [crafted]; Burning Sliver (5249, -2.0) [quest]; Twisted Nether Wand (249144, -2.0) [crafted] |

**New at 60:** head: Frostfire Circlet; neck: Gem of Trapped Innocents; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Dark Storm Gauntlets; waist: Frostfire Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Seal of the Damned; finger2: Don Julio's Band; trinket1: The Restrained Essence of Sapphiron; trinket2: Eye of the Dead; main_hand: Atiesh, Greatstaff of the Guardian; ranged: Lesser Eternal Wand

No-known-source sample (15 of 1212, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 33.3. Weights run: 0.7s. Verify run: 0.6s. 247 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.128 ± 0.138), crit=1.335 ± 0.054, hit=3.164 ± 0.119, spell_haste=1.014 ± 0.120, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -5.4) [crafted]; Flying Tiger Goggles (4368, -6.0) [crafted]; Lucky Fishing Hat (19972, -6.0) [quest] |
| neck | Scout's Medallion (20442) (or Tarnished Locket (279870)) | Warsong Outriders [rep] | 0.0 | yes | Tarnished Locket (279870, +0.0) [quest] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.5 | yes | Double-Stitched Woolen Shoulders (4314, -1.5) [crafted]; Slime-encrusted Pads (6461, -5.5) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -1.0) [dungeon]; Black Whelp Cloak (7283, -1.0) [crafted]; Battle Healer's Cloak (20427, -1.0) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 | yes | Gray Woolen Robe (2585, -1.0) [crafted]; Green Woolen Vest (2582, -1.6) [crafted]; Bloody Apron (6226, -1.6) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | Quests [quest] | 0.8 | yes | Owlbeard Bracers (16981, +0.5) [quest]; Mindthrust Bracers (1974, -0.1) [dungeon]; Featherbead Bracers (15452, -0.1) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -1.0) [world]; Pristine Gloves (253913, -2.6) [crafted]; Apothecary Gloves (10919, -3.0) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 | yes | Novice Arcanist's Sash (253885, -0.1) [crafted]; Novice Ardent's Sash (253887, -2.1) [crafted]; Keller's Girdle (2911, -3.5) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.0 | yes | Silk-threaded Trousers (1929, -3.0) [dungeon]; Filigreed Pristine Leggings (253937, -3.3) [crafted]; Colorful Kilt (10048, -5.0) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 | yes | Feather Padded Treads (285345, -2.5) [world]; Red Woolen Boots (4313, -3.5) [crafted]; Pristine Boots (253889, -4.1) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -4.2) [dungeon]; Black Pearl Ring (6332, -4.7) [world]; The 1 Ring (8350, -4.9) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, -2.2) [dungeon]; Black Pearl Ring (6332, -2.7) [world]; The 1 Ring (8350, -2.9) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 1.3 | yes | Gnarled Necromancer's Staff (251534, +0.0) [quest]; Channeler's Staff (4437, -0.3) [world]; Lesser Staff of the Spire (1300, -0.5) [world] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Quests [quest] | 5.0 | yes | Cookie's Stirring Rod (5198, -2.0) [dungeon]; Torchlight Wand (5240, -3.0) [quest]; Greater Magic Wand (11288, -3.0) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 56.9. Weights run: 0.6s. Verify run: 0.7s. 470 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.065 ± 0.228), crit=2.229 ± 0.108, hit=3.372 ± 0.161, spell_haste=1.051 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | Razorfen Downs: Withered Battle Boar [dungeon] | 11.0 | yes | Silk Headband (7050, -2.0) [crafted]; Embalmed Shroud (7691, -3.0) [dungeon]; Filigreed Pristine Circlet (253975, -3.0) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.4 | yes | Darkspear Warding Pendant (272075, -7.1) [vendor]; Crystal Starfire Medallion (5003, -7.1) [dungeon]; Pendant of Myzrael (4614, -7.4) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.6 | yes | Chestnut Mantle (17695, -1.6) [quest]; Invoker's Mantle (215365, -2.3) [crafted]; Death Speaker Mantle (6685, -2.9) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.0) [quest]; Heavy Woolen Cloak (4311, -1.0) [crafted]; Battle Healer's Cloak (19529, -1.0) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 9.8 | yes | Tree Bark Jacket (1486, +3.2) [dungeon]; High Robe of the Adjudicator (3461, -1.7) [quest]; Robes of Arcana (5770, -1.8) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Owlbeard Bracers (16981, -7.9) [quest]; Nightsky Wristbands (6407, -8.6) [dungeon]; Tabitha's Cuffs (251486, -8.6) [quest] |
| hands | Jutebraid Gloves (10654) | Quests [quest] | 6.3 | yes | Serpent Gloves (5970, +0.7) [dungeon]; Gnoll Casting Gloves (892, -0.3) [world]; Truefaith Gloves (7049, -1.1) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.2 | yes | Warsong Sash (16975, -0.2) [quest]; Belt of Arugal (6392, -2.0) [dungeon]; Ghamoo-ra's Bind (6908, -3.2) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 | yes | Gaze Dreamer Pants (6903, +2.5) [dungeon]; Pristine Leggings (253987, -2.1) [crafted]; Silk-threaded Trousers (1929, -2.5) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.5 | yes | Spidersilk Boots (4320, -0.2) [crafted]; Acidic Walkers (9454, -1.9) [dungeon]; Boots of the Enchanter (4325, -2.5) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -2.0) [rep]; Electrocutioner Lagnut (9447, -4.0) [dungeon]; Sludge-Stained Band (286535, -4.0) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -3.0) [dungeon]; Sludge-Stained Band (286535, -3.0) [world]; Sacred Band (6669, -4.0) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 0.7 | yes | Twisted Chanter's Staff (890, -0.1) [dungeon]; Gnarled Necromancer's Staff (251534, -0.1) [quest]; Channeler's Staff (4437, -0.2) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | Gnomeregan: Dark Iron Agent [dungeon] | 5.0 | yes | Sizzle Stick (8071, +0.0) [quest]; Greater Mystic Wand (217287, +0.0) [crafted]; Cookie's Stirring Rod (5198, -2.0) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 470, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 81.7. Weights run: 0.7s. Verify run: 0.6s. 636 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.379 ± 0.386), crit=3.424 ± 0.185, hit=5.185 ± 0.341, spell_haste=2.354 ± 0.492, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -6.2) [world]; Holy Shroud (2721, -10.0) [dungeon]; Enchanter's Cowl (4322, -11.2) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 | yes | Necklace of Calisea (1714, -6.6) [dungeon]; Darkspear Warding Pendant (272074, -6.6) [vendor]; Darkspear Warding Pendant (272075, -7.4) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 | yes | Green Silken Shoulders (7057, -0.2) [crafted]; Inquisitor's Shawl (19507, -0.5) [dungeon]; Berylline Pads (4197, -1.6) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.9 | yes | Guardian Cloak (5965, +0.0) [crafted]; Icy Cloak (4327, -0.9) [crafted]; Battle Healer's Cloak (19528, -1.9) [rep] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 24.3 | yes | Dreamweave Vest (10021, -2.9) [crafted]; Robe of Power (7054, -5.7) [crafted]; Crimson Silk Vest (7058, -8.5) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Radiant Silver Bracers (4545, -2.0) [quest]; Condor Bracers (15864, -2.0) [quest]; Earthen Silk Cuffs (254019, -5.0) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.5 | yes | Black Mageweave Gloves (10003, -4.5) [crafted]; Red Mageweave Gloves (10018, -4.7) [crafted]; Gilded Handwraps (254021, -8.9) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.5 | yes | Star Belt (4329, -2.5) [crafted]; Deathmage Sash (10771, -2.8) [dungeon]; Defiler's Cloth Girdle (20164, -3.4) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.5 | yes | Crimson Silk Pantaloons (7062, -5.6) [crafted]; Abomination Skin Leggings (23173, -6.5) [dungeon]; Gaze Dreamer Pants (6903, -6.5) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -14.3) [crafted]; Spidersilk Boots (4320, -15.5) [crafted]; Acidic Walkers (9454, -16.0) [dungeon] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -2.0) [rep]; Sea Giant's Toe Ring (274746, -3.0) [vendor]; Advisor's Ring (20426, -4.0) [rep] |
| finger2 | Reedknot Ring (9622) | Quests [quest] | 7.0 | yes | Sea Giant's Toe Ring (274746, -1.0) [vendor]; Electrocutioner Lagnut (9447, -4.0) [dungeon]; Sludge-Stained Band (286535, -4.0) [world] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 50.2 | yes | Windweaver Staff (7757, -44.5) [dungeon]; Staff of Jordan (873, -46.0) [dungeon]; Glimmering Staff (249392, -46.0) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 6.0 | yes | Fizzle's Zippy Lighter (6729, -0.9) [quest]; Wand of Eventide (5214, -1.0) [dungeon]; Sizzle Stick (8071, -1.0) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; finger2: Reedknot Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 636, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 108.0. Weights run: 0.7s. Verify run: 0.7s. 843 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.825 ± 0.694), crit=5.572 ± 0.315, hit=7.662 ± 0.520, spell_haste=not significant (0.844 ± 0.797), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 55.5 | yes | Eye of Theradras (17715, +57.2) [dungeon]; Dreamweave Circlet (10041, -16.2) [crafted]; Bad Mojo Mask (9470, -17.2) [dungeon] |
| neck | Horizon Choker (13085) | Azuregos [world] | 25.5 | yes | Scorn's Icy Choker (23169, -7.6) [dungeon]; Mindburst Medallion (11196, -8.6) [quest]; Darkspear Warding Pendant (272073, -9.1) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 45.8 | yes | Red Mageweave Shoulders (10029, -11.5) [crafted]; Inquisitor's Shawl (19507, -15.1) [dungeon]; Green Silken Shoulders (7057, -17.8) [crafted] |
| back | Deep Woodlands Cloak (19121) | Quests [quest] | 28.4 | yes | Darkspear Raider's Cloak (272076, -2.9) [vendor]; Spritecaster Cape (11623, -3.5) [dungeon]; Runecloth Cloak (13860, -4.8) [crafted] |
| chest | Acumen Robes (17775) | Quests [quest] | 55.5 | yes | Runecloth Robe (13858, -13.5) [crafted]; Runecloth Tunic (13857, -18.4) [crafted]; Hibernal Robe (8113, -19.0) [dungeon] |
| wrist | Shizzle's Nozzle Wiper (11917) | Quests [quest] | 21.9 | yes | Bloodband Bracers (11469, -0.5) [quest]; Imperial Red Bracers (8247, -1.8) [dungeon]; Nethergeld Cuffs (254061, -2.1) [crafted] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 29.2 | yes | Runecloth Gloves (13863, -0.8) [crafted]; Greenleaf Handwraps (19116, -2.2) [quest]; Stormcloth Gloves (10011, -3.4) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 96.1 | yes | Dawnspire Cord (12466, -55.5) [world]; Deathmage Sash (10771, -61.8) [dungeon]; Satyrmane Sash (17755, -63.9) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 35.9 | yes | Kilt of the Atal'ai Prophet (10807, -0.1) [dungeon]; Crimson Silk Pantaloons (7062, -4.2) [crafted]; Imperial Red Pants (8251, -8.5) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Southsea Mojo Boots (20641, +4.1) [quest]; Black Mageweave Boots (10026, -0.2) [crafted]; Vinerot Sandals (17748, -1.6) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 76.6 | yes | Voodoo Band (1996, -63.8) [world]; Mindbender Loop (5009, -63.8) [dungeon]; Black Widow Band (6199, -63.8) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | Boulderfist Magus [world] | 12.8 | yes | Voodoo Band (1996, +0.0) [world]; Mindbender Loop (5009, +0.0) [dungeon]; Black Widow Band (6199, +0.0) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of the Guard Captain (19120, +53.6) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Rune of the Guard Captain (19120, +47.6) [quest]; Tidal Charm (1404, -6.0) [vendor]; Guardian Talisman (1490, -6.0) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 99.9 | yes | Illusionary Rod (7713, -10.9) [dungeon]; Inventor's Focal Sword (17719, -21.9) [dungeon]; Spellshifter Rod (9527, -57.9) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Quests [quest] | 12.3 | yes | Nature's Breath (19118, -1.4) [quest]; Fizzle's Zippy Lighter (6729, -2.8) [quest]; Lesser Eternal Wand (249232, -4.3) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Shizzle's Nozzle Wiper; hands: Red Mageweave Gloves; waist: Defiler's Cloth Girdle; finger1: Blackstone Ring; finger2: Ogremind Ring; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Charged Lightning Rod

No-known-source sample (15 of 843, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 249.0. Weights run: 0.7s. Verify run: 0.7s. 1207 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.124 ± 1.054), crit=8.680 ± 0.479, hit=12.852 ± 0.837, spell_haste=not significant (1.136 ± 1.257), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Quests [quest] | 409.4 | yes | Bloodvine Goggles (19999, -30.8) [crafted]; Enigma Circlet (21347, -123.4) [quest]; Field Marshal's Coronet (16441, -252.8) [vendor] |
| neck | Jewel of Kajaro (19601) | Quests [quest] | 10.6 | yes | Gem of Trapped Innocents (23057, +248.3) [raid]; Onyxia Tooth Pendant (18404, +239.4) [quest]; Fury of the Forgotten Swarm (21809, +239.4) [raid] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 140.1 | yes | Champion's Silk Mantle (23264, -2.2) [vendor]; Lieutenant Commander's Silk Mantle (23319, -2.2) [vendor]; Lieutenant Commander's Silk Mantle (227102, -2.2) [pvp] |
| back | Earthweave Cloak (21187) | Quests [quest] | 128.5 | yes | Chromatic Cloak (18509, -7.0) [crafted]; Drape of Vaulted Secrets (21415, -109.4) [quest]; Hide of the Wild (18510, -113.3) [crafted] |
| chest | Frostfire Robe (22496) | Quests [quest] | 300.4 | yes | Bloodvine Vest (19682, -14.7) [crafted]; Enigma Robes (21343, -137.0) [quest]; Robe of the Archmage (14152, -137.4) [crafted] |
| wrist | Rockfury Bracers (21186) | Quests [quest] | 155.5 | yes | Burrower Bracers (21611, -125.9) [raid]; Frostfire Bindings (22503, -126.7) [quest]; Dryad's Wrist Bindings (19595, -132.5) [rep] |
| hands | Dark Storm Gauntlets (21585) | Ahn'Qiraj [raid] | 167.4 | yes | Gloves of Spell Mastery (14146, +85.9) [crafted]; Sorcerer's Gloves (22066, -25.1) [quest]; Dreadmist Wraps (16705, -37.7) [dungeon] |
| waist | Frostfire Belt (22502) | Quests [quest] | 159.1 | yes | Belt of the Archmage (18405, -15.6) [crafted]; Defiler's Cloth Girdle (20163, -22.9) [rep]; Defiler's Cloth Girdle (20165, -28.0) [rep] |
| legs | Frostfire Leggings (22497) | Quests [quest] | 177.7 | yes | Bloodvine Leggings (19683, -11.5) [crafted]; Enigma Leggings (21346, -19.0) [quest]; Marshal's Silk Leggings (16442, -23.7) [vendor] |
| feet | Enigma Boots (21344) | Quests [quest] | 158.4 | yes | Frostfire Sandals (22500, -6.6) [quest]; General's Silk Boots (16539, -7.1) [vendor]; Marshal's Silk Footwraps (16437, -7.1) [vendor] |
| finger1 | Seal of the Damned (23025) | Naxxramas [raid] | 271.0 | yes | Band of Earthen Might (21182, -21.0) [quest]; Band of Unnatural Forces (23038, -21.0) [raid]; Ring of the Fallen God (21709, -104.8) [quest] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182), Band of Unnatural Forces (23038)) | Frostwolf Clan [rep] | 250.0 | yes | Band of Earthen Might (21182, +0.0) [quest]; Band of Unnatural Forces (23038, +0.0) [raid]; Ring of the Fallen God (21709, -83.8) [quest] |
| trinket1 | The Restrained Essence of Sapphiron (23046) | Naxxramas [raid] | 40.0 | yes | Neltharion's Tear (19379, +261.0) [raid]; Drake Fang Talisman (19406, +217.0) [raid]; Kiss of the Spider (22954, +210.0) [raid] |
| trinket2 | Eye of the Dead (23047) | Naxxramas [raid] | 0.0 | yes | Neltharion's Tear (19379, +301.0) [raid]; Drake Fang Talisman (19406, +257.0) [raid]; Kiss of the Spider (22954, +250.0) [raid] |
| main_hand | Atiesh, Greatstaff of the Guardian (22589) | Quests [quest] | 411.0 | yes | Atiesh, Greatstaff of the Guardian (22630, -14.4) [quest]; Blessed Qiraji Acolyte Staff (21273, -28.4) [quest]; High Warlord's War Staff (234549, -84.1) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Dreambough Wand (249234, -1.0) [crafted]; Twisted Nether Wand (249144, -2.0) [crafted]; Charged Lightning Rod (11860, -2.5) [quest] |

**New at 60:** head: Frostfire Circlet; neck: Jewel of Kajaro; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Dark Storm Gauntlets; waist: Frostfire Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Seal of the Damned; finger2: Don Julio's Band; trinket1: The Restrained Essence of Sapphiron; trinket2: Eye of the Dead; main_hand: Atiesh, Greatstaff of the Guardian; ranged: Lesser Eternal Wand

No-known-source sample (15 of 1207, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

