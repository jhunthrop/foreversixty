# Sim audit round 8: auras lane

Scope: every buff, debuff, totem, blessing, aura, consumable and racial the sim applies from outside a spec's own kit, at bands 20/30/40/50/60.

Sources:
- The 1.60.1.70707 client tables with the hotfix push layered in: SpellEffect, SpellLevels, SpellMisc, SkillLineAbility, ItemSparse, ItemEffect and ItemXItemEffect (sha256s match `data/builds/1.60.1.70707/hotfix-push.json`), plus SpellLabel and SpellCategories from wago.tools for 70707.
- Wowhead Forever's server tooltips (`nether.wowhead.com/forever/tooltip/...`), read 2026-11-01. A tooltip together with the client table counts as one source. Every CONFIRMED-BUG row also has its value checked against the client row directly.

Engine branch: `lane/r8-auras`. Site branch: `wt/r8-auras`.

**Found: 11 CONFIRMED-BUG rows, all fixed:** A1 to A6, D1 to D3, C1 and C3. C2 is the preset re-pick that follows C1.
- At the level-60 raid preset, which moves published numbers:
  - Power Word: Fortitude (A1)
  - Mana Spring counted twice for Elemental (A6)
  - Curse of Recklessness (D1)
  - the foods and the food picks (C1, C2)
  - the defensive elixirs (C3)
- Below 60, published bare numbers:
  - priests' Divine Spirit and Fortitude ranks (A1, A2)
  - Protection warriors' own Sunder rank (D2)
- Below 60, Raid-buffed only: A3 to A5 and D2.
- Custom panel only: D3.
- D5 (whether the armor debuffs stack) stays OPEN.

## Raid buffs

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Power Word: Fortitude was Era's 54 Stamina at every level.**<br>- Forever's ranks are 4, 10, 26, 42, 56 and 70, at levels 1, 12, 24, 36, 48 and 60. That is Era's amounts with Improved Power Word: Fortitude's 30% folded in.<br>- This is on the raid preset at 60 (70, not 54). Every priest's own sim also carries it: `priest.AddRaidBuffs` raises the flag at every band. | SpellEffect 1243/1244/1245/2791/10937/10938 (aura 29, misc 2). Wowhead Forever 2791 reads "by 42" and 10938 "by 70". Engine before: `core/buffs.go:85` (`BuffSpellValues`) and `:258`. Test first: `TestPowerWordFortitudeFollowsTheClientRanks`. Site pin: `TestPowerWordFortitudeAndDivineSpiritMatchTheClient`. | CONFIRMED-BUG | E 2338b7f5a, S c8763da06 |
| A2 | **Divine Spirit gave 40 Spirit from level 1.**<br>- The client's ranks are 17, 23, 33 and 40, at 30, 40, 50 and 60. Nothing exists before 30. It is a trainer spell in Forever, not a talent.<br>- `priest.AddRaidBuffs` raises it for every priest, so the published bare priest numbers carried 40 Spirit at bands 20 to 50.<br>- Unchanged at 60. | SpellEffect 14752/14818/14819/27841 (aura 29, misc 4) and SpellLevels. Wowhead Forever 14752 reads "by 17". Engine before: `core/buffs.go:54` and `:286`. Test first: `TestDivineSpiritFollowsTheClientRanks`. | CONFIRMED-BUG (moves published bare priest numbers) | E 2338b7f5a, S c8763da06 |
| A3 | **Blood Pact was Era's 42 Stamina.**<br>- Rank 5 is 49, plus 0.5 a level to 60, which makes 54. The lower ranks scale the same way.<br>- Stamina only, so no damage number moves. Tank health does. | SpellEffect 6307 to 11767 (`EffectRealPointsPerLevel`, SpellLevels MaxLevel). Wowhead Forever 11767 reads "Stamina by 54". Test first: `TestBloodPactFollowsTheClientRanks`. | CONFIRMED-BUG | E 2338b7f5a, S c8763da06 |
| A4 | **Windfury Totem was rank 3 (246 attack power) at every level.**<br>- The totem is learned at 32, 42 and 52, and gives 95, 179 and 246 attack power (8516, 10608, 10610).<br>- Below 32 a raid has no Windfury Totem. Before this fix, a level-20 or level-30 Raid-buffed run got rank 3. | SpellLevels 8512/10613/10614 and SpellEffect 8516/10608/10610 (aura 99). Wowhead Forever 8512 reads "(95 * 1)" and 10614 "(246 * 1)". Engine before: `core/buffs.go:1736`. Test first: `TestWindfuryTotemFollowsTheClientRanks`. Site pin: `TestWindfuryTotemRanksMatchTheClient`. | CONFIRMED-BUG (bands 20 to 50, Raid-buffed) | E 2338b7f5a, S c8763da06 |
| A5 | **Leader of the Pack and Moonkin Aura reached a raid of any level.**<br>- Leader of the Pack is Feral tier 4 (level 30). Moonkin Form is Balance tier 6 (level 40).<br>- A Raid-buffed run at 20 got 6% crit from druids who cannot exist yet. | `data/builds/1.60.1.70707/talents/druid.json` tiers. Engine before: `core/buffs.go:240-247`. Test first: `TestTalentRaidAurasWaitForTheirTalentLevel`. | CONFIRMED-BUG (bands 20 and 30) | E 2338b7f5a |
| A6 | **An Elemental shaman counted Mana Spring twice under the raid preset.**<br>- The rotation drops its own Mana Spring Totem (10497), and the preset already carries `mana_spring_totem`. That gave 50 mp5 at 60.<br>- Two Mana Spring auras do not stack: you stand in the strongest. The shaman's own totem now adds only its excess over the raid's (6.25 with Restorative Totems 5/5). | `shaman/water_totems.go:171` before ("stays a separate source"). `data/curated/apl/shaman-elemental.json` casts 10497. Test first: `TestOwnManaSpringDoesNotStackWithTheRaids` (gave +25 and +31.25). | CONFIRMED-BUG | E b092aee70 |
| A7 | Checked and aligned, with rank tables pinned to the client:<br>- Arcane Intellect, Blessing of Might, Blessing of Wisdom, Mark of the Wild, Battle Shout (Ahn'Qiraj rank 7 out).<br>- Trueshot Aura (15/22/30/40/50), Devotion Aura, Heart of the Lion (280 attack power at 60).<br>- Strength of Earth (42), Grace of Air (77), Mana Spring by level (round 7).<br>- Blessing of Kings: 10%, a trainer spell from 20 (20217, aura 137). Not in the preset.<br>- Leader of the Pack and Moonkin Aura: 3% at 60 (24932 and 24907, aura 290).<br>- Windfury at 60: 20% proc (8515, 10612), 1 s, 2 charges. | `sim/leveling/buff_ranks_client_test.go`, layered SpellEffect | Aligned | - |
| A8 | Double counting with a spec's own copy:<br>- A warrior's own Battle Shout and the raid's share one aura (label "Battle Shout").<br>- A shaman's Strength of Earth is `core.StrengthOfEarthTotemAura`, the same aura.<br>- Trueshot Aura is an exclusive category.<br>- Arcane Intellect, Mark of the Wild, Blessing of Might, Devotion Aura and Hunter's Mark are deduplicated against the kit in `request.Layer`.<br>- Moonkin was fixed in round 7. Judgement of the Crusader registers once.<br>- Mana Spring is A6. | Code read | Aligned except A6 | - |

## Raid debuffs

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | **Curse of Recklessness was Era's: 640 armor, plus 90 attack power given to the boss.**<br>- Forever's 11717 removes 505 armor. Its only other effect is aura 92 (the curse stops the target fleeing).<br>- The physical lane fixed the 505 too (a16b3446e). This fix also ranks the curse by level. | SpellEffect 11717 (aura 674 -505; aura 92 0). Wowhead Forever 11717 reads "reducing the target's armor by 505", 7659 "by 395". Engine before: `core/debuffs.go:748`. Test first: `TestCurseOfRecklessnessTakesTheRanksArmorOnly`. Site pin: `TestCurseOfRecklessnessHasNoAttackPowerInTheClient`. | CONFIRMED-BUG (60: every physical spec on the raid preset) | E dcc39bf93, S c8763da06 |
| D2 | **Every raid debuff applied its top rank at every level.** Each now takes the rank the raid's level has learned (`raidDebuffLevel`, the first player):<br>- Sunder Armor: 90/180/270/360/450 at 10/22/34/46/58.<br>- Faerie Fire: 175/285/395/505 at 18/30/42/54.<br>- Curse of Recklessness: 175/285/395/505 at 14/28/42/56.<br>- Curse of the Elements: 4/6/8/10% and 30/45/60/75 resistance at 20/30/40/50.<br>- Hunter's Mark: 26/59/98/143 at 6/22/40/58.<br>- The raid's Judgement of the Crusader: 23/35/58/92/127/161 at 6/12/22/32/42/52.<br><br>A warrior's own Sunder, a druid's Faerie Fire and a warlock's Curse of Recklessness use these same core auras. So a level-20 Protection warrior's own Sunder stripped 450 a stack where its rank 1 strips 90. That moves the published bare Protection warrior number (Before/after). | SpellEffect and SpellLevels for each id, listed in `core/raid_debuff_ranks.go`. Engine before: `core/debuffs.go:696` (450), `:778` (505), `:508` (10%/75), `:840` (143), `:30` (161). Test first: `TestRaidDebuffsFollowTheClientRanks`, `TestRaidDebuffsAreCastAtTheRaidsLevel`. Site pin: `TestRaidDebuffRanksMatchTheClient`. | CONFIRMED-BUG | E dcc39bf93, S c8763da06 |
| D3 | **Four debuffs Forever made the caster's own still buffed the whole raid.**<br>- Shadow Weaving (15258), Improved Scorch's Fire Vulnerability (22959) and Improved Shadow Bolt's Shadow Vulnerability (17794) are aura 270, damage taken from the caster. Stormstrike (17364) is aura 271 on the shaman's own Lightning Bolt, Chain Lightning and Earth Shock.<br>- Wowhead Forever: "the Shadow damage the target takes from you", "all Fire damage you deal to your target", "from your attacks", "the damage you deal to the target with your next Lightning Bolt".<br>- As raid flags they gave everyone Era's 15% Shadow, 15% Fire, 20% Shadow and 20% Nature damage taken.<br>- They are in no preset. Every /sim user could tick them in the custom panel.<br>- The flags now apply nothing. The external-caster code is deleted. Each class keeps its own copy from its talents. | SpellEffect rows above; tooltips 15257, 11095, 22959, 17793, 17364. Test first: `TestCasterOnlyDebuffsGiveTheRaidNothing` (panicked before). The mage tests that pinned the raid Improved Scorch as a stand-in for another mage are rewritten (`TestRaidImprovedScorchGivesThisMageNothing`). | CONFIRMED-BUG (custom panel only) | E a56f0d992, E f9ee49cd4 |
| D4 | Checked and aligned at 60:<br>- Sunder 450 (11597), Faerie Fire 505 (9907), Curse of the Elements 10%/75 (1311680), Hunter's Mark 143 (14325), Judgement of the Crusader 161 (20303).<br>- Judgement of Wisdom and Judgement of Light ranks (pinned).<br>- Winter's Chill is a no-op (round 6). | Layered SpellEffect | Aligned | - |
| D5 | **Do Sunder, Faerie Fire and Curse of Recklessness stack in Forever?** (the physical lane's question)<br>- All three moved to aura 674. The 70707 tables give them separate SpellLabels: Sunder 25, Faerie Fire 21, Curse of Recklessness 19 (shared with Curse of the Elements), Expose Armor 20.<br>- No table row marks them exclusive.<br>- 674 is also the armor aura on Devotion Aura and the armor elixirs since 70291. That reads as a renumbering, not a stacking rule.<br>- The engine keeps Era's rule: Sunder and Expose major, Faerie Fire minor, Curse of Recklessness separate. | wago SpellLabel and SpellCategories, build 70707 | OPEN: a beta boss's armor reading, or a level-60 log's armor-dependent damage, settles it | - |

## Consumables

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Forever's foods give different stats from Era's.**<br>- A Forever food's "Nutritious Food" spell passes an aura 227 value to the Well Fed it triggers. The Well Fed's aura names the stat.<br>- Grilled Squid: 1% critical strike (Well Fed 1249523, aura 290). Era: 10 Agility.<br>- Tender Wolf Steak: 15 Agility. Hot Wolf Ribs: 10 Agility (1248420). Era: Stamina and Spirit.<br>- Runn Tum Tuber Surprise: 15 Intellect. Era: 10.<br>- Smoked Sagefish and Sagefish Delight: 4 and 7 spell damage. Era: mana per five seconds.<br>- Three of these were in the raid preset (melee, caster and tank food). | Item rows 13928, 18045, 13851, 18254, 21072, 21217. SpellEffect of their Nutritious Food spells and Well Feds. Wowhead Forever item tooltips: "gain 1% Critical Strike chance", "gain 15 Agility", "gain 15 Intellect", "gain 4 Spell Damage", "gain 7 Spell Damage". Engine before: `core/consumes.go:360-412`. Test first: `TestFoodsMatchTheClient`. | CONFIRMED-BUG | E b707d25a3 |
| C2 | **The preset's foods, re-picked on Forever's values** (seeds 1 and 50001, 10,000 iterations, horde, band-60 raid):<br>- Damage casters: Nightfin Soup (22 spell damage), +1.0 to +2.2% over Runn Tum.<br>- Strength users (Arms, Fury, Retribution, Enhancement, Feral): Smoked Desert Dumplings (20 Strength), +0.4 to +0.6% over Squid.<br>- Tanks: Dumplings, +0.5 to +0.8% threat over the Steak.<br>- Rogues and hunters keep Grilled Squid. The Dumplings measured 0.05 to 0.95% behind for them.<br>- No stamina food a Phase 1 tank can cook is modelled. | `meas/food.txt` in the lane scratch. `data/curated/presets.json`. `sim/request/raid-consumes.json` regenerated (genraidconsumes). `TestRaidPresetConsumesFollowTheRole` updated. | Fix (follows C1) | S 28cd18c97 |
| C3 | **Defensive elixirs were Era's.**<br>- Elixir of Greater Defense is 450 armor (Era 250).<br>- Elixir of Defense is 250 (Era 150).<br>- Forever's Elixir of Fortitude, item 250334, is 200 health. Era's 3825, which the engine's 120 came from, is not in the client.<br>- Greater Defense and Fortitude are in the tank preset. | SpellEffect 11348 (aura 674, 450), 1250928 (aura 34, 200), and 8951's spell. Wowhead Forever 13445: "Increases armor by 450"; 250334: "increase your maximum health by 200". Engine before: `core/consumes.go:497-530`. Test first: `TestDefensiveElixirsMatchTheClient`. | CONFIRMED-BUG (tanks) | E b707d25a3 |
| C4 | Checked and aligned against the item's use spell:<br>- Mongoose 25 Agility and 2%; Juju Power 30; Juju Might 40 and 40; R.O.I.D.S. 25.<br>- Elemental Sharpening Stone 2% (22755, aura 52, melee only); Dense Sharpening Stone +8.<br>- Flask of Supreme Power 150; Greater Arcane Elixir 35; Brilliant Wizard Oil 36 and 1%.<br>- Major Mana Potion 1350 to 2250; Demonic Rune 900 to 1500 mana and 600 to 1000 health.<br>- Mighty Rage Potion 45 to 75 rage and 60 Strength for 20 s; Thistle Tea 100.<br>- Distilled Wisdom 2000; Greater Mageblood 20; Greater Cleric's 40; Sages 25 and 2%; Sage's Tea 44; the mana oils (pinned).<br>- Fire Power 10, Shadow Power 40, Frost Power 15.<br>- Titans 1200; Greater Stoneshield 2000; Rumsey Black Label 15.<br>- Dirge's Chops 25; the Blessed Sunfruits 10; Smoked Desert Dumplings 20; Nightfin 22.<br>- Required levels come from ItemSparse through `request.ItemLevels`. | it.py over ItemSparse, ItemXItemEffect, ItemEffect and SpellEffect | Aligned | - |

## Racials

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Every combat racial of the ten race rows matches the layered tables:<br>- Ruthless 1%, The Human Spirit 10%, Wind Blessed 2%, Artisan's Precision 1% hit, Elune's Light 10% for 15 s.<br>- Blood Fury 10% for 15 s, Quickness 1% dodge, Endurance 5%, Expansive Mind 5%, Berserking 10% for 15 s.<br>- Beast Slaying, Big Game Hunter and Elemental Insight 5% (aura 168). | SpellEffect and SpellMisc of each id; `core/racials.go` | Aligned | - |

## The /sim Raid-buffed list

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | The page's Raid-buffed set is generated from the curated preset and `raid-consumes.json`, so it matches by construction.<br>- Two copies of a stale note said the engine gated the blessings to the Alliance and the totems to the Horde. 86d33a100 removed those gates (`applyBuffEffects` gates only world buffs).<br>- The presets note is corrected. `web/src/lib/sim/settings.ts:131-133` carries the same stale comment (page lane). | `scripts/sync-raid-preset.mjs`, `core/buffs.go:181-189` | Note fixed; page comment routed | S 28cd18c97 |

## Before/after

Engine lane head plus the site branch, against cd4dfdcc0 with origin/main. 10,000 iterations, seeds 1 and 50001.

**Band 60, raid preset** (both factions, seed 1 absolute shown for horde):

| spec | horde s1 / s50001 | alliance s1 / s50001 | horde s1 absolute | main causes |
|---|---|---|---|---|
| druid-balance | +2.22 / +2.27 | +2.26 / +2.24 | 567.3 -> 579.9 | Nightfin Soup |
| druid-feral | -1.20 / -1.18 | -1.14 / -1.16 | 616.3 -> 608.9 | Curse of Recklessness 505, Dumplings |
| druid-feral-bear (threat/s) | -1.42 / -1.42 | -1.47 / -1.40 | 868.6 -> 856.3 | Curse of Recklessness, food |
| druid-restoration (HPS) | +0.01 / +0.03 | +0.01 / +0.04 | 770.3 -> 770.4 | - |
| hunter-beast-mastery | -1.93 / -1.90 | -1.91 / -1.92 | 674.8 -> 661.8 | Curse of Recklessness, Squid |
| hunter-marksmanship | -2.04 / -2.04 | -2.04 / -2.04 | 666.7 -> 653.1 | same |
| hunter-survival | -1.70 / -1.68 | -1.69 / -1.68 | 772.6 -> 759.5 | same |
| mage-arcane | +1.35 / +1.27 | +1.37 / +1.29 | 592.5 -> 600.5 | Nightfin Soup |
| mage-fire | +1.64 / +1.72 | +1.50 / +1.57 | 660.0 -> 670.8 | Nightfin Soup |
| mage-frost | +1.98 / +1.98 | +1.99 / +2.00 | 589.8 -> 601.5 | Nightfin Soup |
| paladin-holy (HPS) | 0.00 / 0.00 | 0.00 / 0.00 | 642.0 -> 642.0 | - |
| paladin-protection (threat/s) | +0.66 / +0.52 | +0.59 / +0.51 | 1011.5 -> 1018.2 | Dumplings |
| paladin-retribution | -0.12 / -0.12 | -0.21 / -0.20 | 659.6 -> 658.8 | Curse of Recklessness, Dumplings |
| priest-discipline (HPS) | +0.15 / +0.11 | +0.08 / +0.08 | 614.2 -> 615.1 | Fortitude 70 |
| priest-holy (HPS) | +0.06 / +0.09 | +0.07 / +0.07 | 670.8 -> 671.2 | - |
| priest-shadow | +1.41 / +1.41 | +1.37 / +1.37 | 714.1 -> 724.2 | Nightfin Soup |
| rogue-assassination | -1.82 / -1.81 | -1.89 / -1.86 | 620.6 -> 609.3 | Curse of Recklessness, Squid |
| rogue-combat | -1.99 / -1.99 | -2.02 / -2.04 | 618.3 -> 606.0 | same |
| rogue-subtlety | -1.73 / -1.72 | -1.77 / -1.77 | 542.0 -> 532.6 | same |
| shaman-elemental | +0.59 / +0.59 | +0.47 / +0.43 | 473.0 -> 475.8 | Nightfin Soup, Mana Spring once |
| shaman-enhancement | -0.95 / -0.95 | -0.92 / -0.90 | 599.7 -> 594.0 | Curse of Recklessness, Dumplings |
| shaman-restoration (HPS) | +0.01 / +0.04 | 0.00 / +0.02 | 820.7 -> 820.8 | - |
| warlock-affliction | +1.65 / +1.69 | +1.62 / +1.59 | 574.5 -> 584.0 | Nightfin Soup |
| warlock-demonology | +1.60 / +1.52 | +1.52 / +1.47 | 642.5 -> 652.8 | Nightfin Soup |
| warlock-destruction | +1.66 / +1.63 | +1.67 / +1.74 | 639.1 -> 649.7 | Nightfin Soup |
| warrior-arms | -1.25 / -1.21 | -1.27 / -1.20 | 729.9 -> 720.8 | Curse of Recklessness, Dumplings |
| warrior-fury | -1.30 / -1.34 | -1.27 / -1.28 | 767.8 -> 757.8 | same |
| warrior-protection (threat/s) | -0.65 / -0.75 | -0.62 / -0.71 | 1043.7 -> 1036.9 | same |

Every row except Elemental was measured at E 4263415d6 + S 28cd18c97. Elemental was measured with E b092aee70 too; that commit touches only a shaman who drops its own Mana Spring, and with only the earlier commits Elemental read +2.0%. So counting Mana Spring once is about -1.4%, and Nightfin Soup about +2%. The engine-only run (before any preset change) put Curse of Recklessness plus Grilled Squid at -1.6 to -2.0% for melee and hunters.

**Bare preset, bands 20 to 50** (published numbers; horde, both seeds; every spec not listed is 0.00%):

| spec | band 20 | band 30 | band 40 | band 50 | cause |
|---|---|---|---|---|---|
| priest-discipline (HPS) | -9.51 / -9.23 | -3.77 / -3.77 | -3.23 / -3.30 | -0.91 / -0.96 | Divine Spirit by rank (none before 30), Fortitude by rank |
| priest-holy (HPS) | -10.53 / -10.53 | -1.67 / -1.52 | -2.13 / -2.05 | -0.85 / -0.90 | same |
| priest-shadow | -5.96 / -5.96 | -2.90 / -2.90 | -2.10 / -2.18 | -0.32 / -0.26 | same |
| warrior-protection (threat/s) | -15.82 / -15.82 | -12.84 / -12.84 | -9.21 / -9.20 | -4.06 / -4.05 | own Sunder Armor at its rank (90 at 20, not 450) |

Absolute values, seed 1:
- Discipline at 20: 32.6 -> 29.5.
- Holy at 20: 34.2 -> 30.6.
- Shadow at 20: 38.6 -> 36.3.
- Protection warrior at 20: 86.6 -> 72.9; at 50: 542.4 -> 520.4.

**Raid-buffed /sim set below 60** (the page's `raid-consumes.json` row on the band's bare gear; horde, seed 1):

| spec | band 20 | band 30 | band 40 | band 50 |
|---|---|---|---|---|
| druid-balance | -8.2 (44.1 -> 40.5) | -5.4 (94.7 -> 89.6) | +3.5 (180.8 -> 187.1) | +2.8 (349.1 -> 358.8) |
| druid-feral | -16.7 (105.9 -> 88.2) | -9.2 (222.8 -> 202.3) | -8.1 (334.1 -> 307.0) | -4.6 (450.8 -> 430.2) |
| druid-feral-bear | -18.9 (130.2 -> 105.6) | -10.0 (176.5 -> 158.8) | -6.7 (354.2 -> 330.6) | -4.8 (653.9 -> 622.7) |
| druid-restoration | -11.8 (80.4 -> 70.9) | -5.5 (186.7 -> 176.5) | -3.2 (256.3 -> 248.2) | -0.8 (507.6 -> 503.4) |
| hunter-beast-mastery | -29.1 (213.4 -> 151.2) | -16.8 (276.7 -> 230.3) | -11.8 (403.4 -> 355.7) | -7.4 (550.5 -> 510.0) |
| hunter-marksmanship | -28.5 (207.2 -> 148.2) | -16.6 (277.8 -> 231.6) | -12.1 (387.5 -> 340.8) | -7.7 (539.9 -> 498.1) |
| hunter-survival | -31.0 (149.1 -> 102.9) | -21.0 (219.7 -> 173.5) | -14.4 (364.6 -> 312.2) | -10.2 (492.1 -> 441.8) |
| mage-arcane | -9.4 (55.6 -> 50.4) | -6.1 (106.9 -> 100.4) | +0.6 (209.6 -> 210.9) | +1.4 (375.1 -> 380.5) |
| mage-fire | -11.0 (57.1 -> 50.8) | -7.9 (112.4 -> 103.5) | +0.1 (243.8 -> 244.0) | +2.0 (451.1 -> 460.1) |
| mage-frost | -8.8 (56.9 -> 51.9) | -6.5 (98.9 -> 92.5) | +2.2 (240.7 -> 246.0) | +2.6 (454.8 -> 466.5) |
| paladin-holy | -7.5 (61.0 -> 56.4) | -9.3 (149.9 -> 135.9) | -2.8 (246.8 -> 239.9) | -0.5 (413.2 -> 411.1) |
| paladin-protection | -45.6 (157.7 -> 85.8) | -31.7 (210.4 -> 143.8) | -9.2 (576.1 -> 522.9) | -2.6 (799.1 -> 778.2) |
| paladin-retribution | -40.1 (129.3 -> 77.4) | -28.5 (254.5 -> 181.9) | -10.4 (364.7 -> 326.8) | -4.0 (433.7 -> 416.4) |
| priest-discipline | -6.6 (60.3 -> 56.3) | -4.7 (162.2 -> 154.5) | -1.9 (247.0 -> 242.4) | -0.5 (444.1 -> 441.8) |
| priest-holy | -6.8 (63.4 -> 59.1) | -2.9 (124.9 -> 121.3) | -1.5 (209.0 -> 205.9) | -0.2 (421.1 -> 420.4) |
| priest-shadow | -10.1 (49.6 -> 44.6) | -6.0 (90.0 -> 84.6) | +0.9 (272.4 -> 274.9) | +1.9 (501.7 -> 511.0) |
| rogue-assassination | -30.3 (112.7 -> 78.6) | -20.5 (157.4 -> 125.2) | -10.4 (253.7 -> 227.2) | -6.0 (393.9 -> 370.3) |
| rogue-combat | -28.8 (114.2 -> 81.3) | -18.7 (171.6 -> 139.5) | -10.8 (282.3 -> 251.7) | -6.3 (411.4 -> 385.4) |
| rogue-subtlety | -28.3 (105.5 -> 75.6) | -16.5 (152.5 -> 127.3) | -7.8 (245.8 -> 226.7) | -3.8 (366.2 -> 352.3) |
| shaman-elemental | -9.7 (40.4 -> 36.5) | -8.9 (86.9 -> 79.2) | -3.3 (172.2 -> 166.5) | +1.1 (350.3 -> 354.0) |
| shaman-enhancement | -31.2 (94.0 -> 64.7) | -10.5 (216.9 -> 194.1) | -8.1 (300.2 -> 275.8) | -4.4 (429.1 -> 410.4) |
| shaman-restoration | -5.0 (83.2 -> 79.0) | -3.6 (149.0 -> 143.7) | -1.2 (228.3 -> 225.5) | -0.2 (435.0 -> 434.2) |
| warlock-affliction | -12.1 (72.6 -> 63.8) | -6.7 (115.7 -> 108.0) | +0.9 (190.6 -> 192.3) | +1.7 (374.9 -> 381.4) |
| warlock-demonology | -11.5 (67.7 -> 59.9) | -6.6 (107.2 -> 100.1) | +0.7 (199.8 -> 201.1) | +1.6 (385.4 -> 391.6) |
| warlock-destruction | -12.8 (71.6 -> 62.4) | -7.2 (122.2 -> 113.4) | +0.4 (203.2 -> 204.1) | +1.8 (461.9 -> 470.2) |
| warrior-arms | -48.1 (101.0 -> 52.4) | -30.4 (228.7 -> 159.1) | -24.0 (313.3 -> 238.0) | -8.3 (518.3 -> 475.4) |
| warrior-fury | -29.0 (115.2 -> 81.8) | -21.4 (195.2 -> 153.5) | -10.3 (379.6 -> 340.4) | -5.5 (579.2 -> 547.4) |
| warrior-protection | -24.0 (163.0 -> 123.9) | -14.9 (282.6 -> 240.4) | -6.1 (610.2 -> 572.8) | -3.9 (876.6 -> 842.3) |

These numbers are what a user's Raid-buffed run below 60 shows. No published number uses them. Most of the drop comes from four changes:
- Windfury Totem now arrives at 32, not at level 1 at rank 3.
- Sunder, Faerie Fire and Curse of Recklessness strip their ranks' armor (at 20: 5 x 90, 175 and 175, where they stripped 2250, 505 and 640).
- Leader of the Pack and Moonkin Aura wait for 30 and 40.
- Divine Spirit and Fortitude take their ranks.

The casters gain at 40 and 50 from Nightfin Soup.

## For other lanes

- **Integrator.** The CoR fix overlaps the physical lane's a16b3446e: keep dcc39bf93 (it also does ranks).
- **Integrator, re-measure every spec.** These commits move specs outside any single group:
  - E dcc39bf93: Curse of Recklessness; Protection warrior below 60.
  - E 2338b7f5a: priests below 60.
  - E b707d25a3 and S 28cd18c97: every damage spec and tank at 60.
  - E b092aee70: Elemental.
- **Integrator, goldens.** No `TestRotationLadder` golden was regenerated. The fork goldens that moved: TestFeralTank, TestP1Mage, TestProtection (paladin), TestElemental, TestEnhancement, TestP1DPSWarrior, TestP1TankWarrior.
- **Page lane.** `web/src/lib/sim/settings.ts:131-133` still says the engine gates Blessing of Might and Wisdom to the Alliance and Strength of Earth and Mana Spring to the Horde. That gate was removed in 86d33a100.
- **Logs owner.** `logs/engine/mechanics/consumables/catalogue.json` and `data/curated/coach/*.md` name Grilled Squid as every melee spec's food and Tender Wolf Steak as the tank food. Now:
  - The Steak is 15 Agility.
  - Strength users and tanks eat Smoked Desert Dumplings (20 Strength), and casters Nightfin Soup.
  - `logs/engine/rating/score_test.go` uses both foods as fixtures.
- **Warlock lane.** `warlock/curses.go registerCurseOfRecklessnessSpell` keeps Season of Discovery bracket maps (ids and costs only at 25/40/50/60). At other levels it registers rank 0. The aura's armor is now ranked by level in core.
- **Hybrids lane.** `paladin/sotc.go` carries the same Judgement of the Crusader amounts as `core.JudgementOfTheCrusaderRanks`. It could read the core table.

## Still open

1. D5: whether Forever's armor debuffs (aura 674) stack with each other.
2. Mana Tide Totem's per-tick mana (290) comes from the totem creature's spell (277993), which this lane did not trace. Rank 4 (1325256, level 58) is new in Forever. It is not in any preset.
3. Retribution Aura, Thorns and Stoneskin Totem keep one rank's amount at every level. None is in a preset, and all are reactive or defensive (no damage-dealer number).
4. The custom /sim panel still offers every engine field. D3's four flags now do nothing, and the page could hide them (page lane).
