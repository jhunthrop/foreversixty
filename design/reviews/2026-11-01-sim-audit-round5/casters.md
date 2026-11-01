# Sim audit round 5: casters (Mage and Priest), 2026-11-01

Specs: mage-arcane, mage-fire, mage-frost, priest-discipline, priest-holy, priest-shadow.

## Method

- **Engine.** `lane/r5-casters` from the pin `f8a5f94ce`, with the warlock lane's `fce7dd53d` (healing grants no hidden damage) merged in at the integrator's request, so the after column includes it. The before column runs a second worktree at the pin against origin/main's site data.
- **Client.** wago.tools tables for 1.60.1.70707, and SpellEffect for every `wowf` build from 70009 to 70707 (read 2026-11-01). The local client's DBCache.bin is build 70338, from 11 October.
- **Hotfixes.** wowforevertalents.com's patch notes, "31 October 2026, hotfix over build 1.60.1.70685". The notes read the client's hotfix cache of 30 October (pushes up to 113531, read again over 70707), and they give old and new values. Blizzard's racial notes of 26 and 27 October (Kaivax, reposted by MMO-Champion and Warcraft Tavern). Wowhead Forever tooltips and spell pages.
- **Runs.** `spec-breakdown`, raid preset, band 60, 10,000 iterations, seeds 1 and 50001, both factions. `talent-search -preset raid -faction horde -keep <raid-keeps.json>` for Fire and Arcane after the Mage Armor fix.
- **Tests.** Every confirmed bug got a failing test first, seen failing, then passing.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | The server's 30 October hotfix changes 216 things that no shipped table carries. Mage: Ice Barrier, Arcane Blast, Ice Lance, Winter's Chill, Arctic Reach. Priest: Holy Binding, Binding Heal, Renewed Hope, Shadow Word: Death. Paladin 118 changes, druid 43, warlock Improved Drains, hunter. Racials. | wowforevertalents.com patch notes, read 2026-11-01. For every mage and priest line, I checked the engine against the new value. It already follows Arcane Blast (12%, +100%, 18 s, `arcane_blast.go`), Ice Lance's 300% and 8 s, Winter's Chill (20% and 7%, `cone_of_cold.go:130`), Ice Barrier's 15% / 5% and Shadow Word: Death's 30% backlash (`shadow_word_death.go:36`). Holy Binding is X7. | Aligned for mage and priest. **Other lanes: check paladin, druid, warlock and hunter against this list.** | - |
| X2 | Orc's Savage (1% crit) is still applied, and Command (pets +5%) is missing. | Blizzard, 26 October: "Savage removed", "Command re-added, increasing Warlock and Hunter pet damage by 5%". The 30 October hotfix cache says the same. The shipped 20574 row still reads 1. Test: TestOrcHasCommandAndNoSavage. | CONFIRMED-BUG (fresh, two sources) | E e4d9ff60c |
| X3 | Troll Berserking lasts 10 s. | Blizzard, 27 October: "Berserking duration increased to 15 seconds (was 10 seconds)". Hotfix cache: 15 s. The shipped SpellDuration is 10 s. Test: TestElunesLightAndBerserkingMatchTheClientRows. | CONFIRMED-BUG (fresh, two sources) | E e4d9ff60c |
| X4 | Touch of the Grave's internal cooldown is 25 s. This was a round-4 single-source item. | Blizzard, 27 October: "reduced to 15 seconds (was 25 seconds)". Hotfix cache: 15 s on both versions. The shipped ProcCategoryRecovery is 25000. Test: TestTouchOfTheGraveMatchesTheClientRows. | CONFIRMED-BUG (fresh, two sources) | E e4d9ff60c |
| X5 | The Human Spirit gives 5% Spirit. This was a round-4 single-source item. | Blizzard, 26 October: "Human Spirit now increases Spirit by 10% (was 5%)". Hotfix cache: 10%. The shipped 20598 row reads 5. Test: TestTheHumanSpiritIsTenPercentSpirit. | CONFIRMED-BUG (fresh, two sources) | E e4d9ff60c |
| X6 | Skyborne's Wind Blessed gives 1% haste. | Blizzard, 26 October: "reduced to 2% (was 3%)". Hotfix cache: 2%. The shipped 1259710 rows read 3, and the engine had 1. Test: TestWindBlessedHastesCastingMeleeAndRanged. | CONFIRMED-BUG (fresh, two sources) | E e4d9ff60c |
| X7 | A Demonic Rune took no health. | Spell 16666, effect 694075: 800 damage to the user, variance 0.5, so 600 to 1000. The engine takes the mean, so no random stream moves. This matters to heals that land on the user and to health gates such as Disc's Dark Sacrifice. Test: `sim/core/demonic_rune_test.go`. | CONFIRMED-BUG (fresh) | E de2ebaf86 |
| X8 | Tauren Endurance, Dwarf Artisan's Precision, Human Ruthless 1% and Blood Fury's 2 min. | Same notes. The engine already matches. Endurance's 2% crit, open in round 4, is settled: "no hit chance or critical damage modifiers". | Aligned | - |
| X9 | Blizzard's change makes Undead the best Horde race for Arcane and Fire. | Horde, same gear, lane engine. Undead beats Orc by 0.4 to 0.6% on Arcane (603.2 / 603.5 against 600.9 / 600.7) and on Fire (614.0 / 614.9 against 610.8 / 611.1). Troll is lower. The guides' `recommendedRaces` still name Orc. | OPEN, owner's call. The gap is under the 1% the talent rule asks for. | - |
| X10 | `data/curated/races.json` still holds the BlizzCon-demo racial text. | Lines 36, 396, 471 and 566. | Noted for the data owner | - |

## Mage: Frost

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Item 1: Ice Barrier's +15% / +5%. | **Second source found.** The 30 October hotfix cache shows "increasing your Frost damage by 3% → 15%" and "an additional 3% → 5%" on rank 1 and rank 4 (wowforevertalents.com, 31 October entry). The other source is Wowhead Forever: the 11426 tooltip resolves `$1326891s1` to 15 and `$s2` to 5, and the spell pages list 16% and 6% under Wowhead's base+1 convention. The timeline explains the 3%: the rows enter the shipped tables at 70630 (27 October) at 3/3, the 16 October hotfix value (kami-labs, 22 October; endgametools, 23 October). They stay 3/3 through 70707. 70532 and earlier have no such rows. The hotfix cache on the Mac (70338, 11 October) has no Ice Barrier rows. | Aligned: the engine keeps 15% / 5%, and the comment now names both sources | E 422be6dac (comment) |
| R2 | Item 2: Ice Lance's 8 s on ranks 2 to 6. | **Second source found.** The 30 October hotfix shows "Ice Lance Rank 6 (5 ranks changed): 160 Mana · 30 yd range · Instant · 8 sec cooldown", with the cooldown marked new. The 70707 entry says "Ice Lance has an 8 sec cooldown". talentsforever.com's spellbook shows 8 s too. | Aligned (two sources) | - |
| R3 | Item 3: Frost Blast. | Settled: the patch notes' 70707 entry says "Removed Frost Blast (gone from the mage spellbook)", "trainers no longer teach Frost Blast". The engine does not register it, which is right. `trainables/mage.json` still lists it as `class_spell` with no skill line; data owner. | Aligned (closed) | - |
| R4 | Troll Berserking 15 s (X3). | Frost horde +0.2 to +0.3%. | CONFIRMED-BUG | as X3 |
| R5 | Mage Armor (M1). | +0.2 to +0.5%. | CONFIRMED-BUG | as M1 |
| R6 | Fresh: talents the build takes, checked against the 70707 rank descriptions and effect rows. Precision is 2% a rank on all spells (aura 55). Ice Shards is 20% a rank. Frost Channeling is 5% a rank. Shatter is 17/33/50%. Winter's Chill is 20% a rank. Incineration is 2% a rank on Fire Blast, Ice Lance, Arcane Blast and Scorch. Arcane Concentration and Missile Barrage are 40% for Arcane Blast and 20% for the others. | `sim/mage/talents.go:590-660`, `703-840`. | Aligned. One small gap: Precision is ClassSpellsOnly, while aura 55 has no mask, so item procs miss it. No published set has a mage damage proc. | - |
| R7 | Raid talent search. | Not re-run; Frost changed by under 1% outside the racials. | Hold | - |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Orc Savage removed (X2). | Arcane horde -0.8%. | CONFIRMED-BUG | as X2 |
| A2 | Mage Armor (M1). | +2.0 to +2.2%. | CONFIRMED-BUG | as M1 |
| A3 | Fresh: Arcane Mind (2% Intellect and 20% Arcane crit bonus a rank), Arcane Instability (1% / 1% a rank), Arcane Power (30% / 30%), Magister's Regalia 3-piece (9346, +18) and 5-piece (450527, 5% for 200 mana). | The client rows match `talents.go:579-583` and `client_sets_gen.go:23`. Arcane wears 3 pieces. | Aligned | - |
| A4 | Raid talent search after Mage Armor. | Deep Arcane, +2 Incineration for -2 Frost Channeling: +0.9% (±0.3%). | Hold (under 1%) | - |

## Mage: Fire

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| M1 | **Every mage sim wore no armor.** The request asked for Molten Armor. | `sim/request/request.go:703` sent `Mage_Options_MoltenArmor`. Molten Armor (428741) is a Season of Discovery rune with no trainer row, and the engine's armor switch (`mage/mage.go:253`) has no case for it. Mage Armor (6117, 22782, 22783) is trainable, and its aura 134 keeps 50% of mana regeneration while casting. Test: TestEveryMageSpecWearsMageArmor. | CONFIRMED-BUG (fresh) | S 4a4d9784c |
| F1 | Raid build after M1. | The talent search, and then 10,000 iterations at both seeds, gives one Accelerant point moved to Improved Fire Blast 2/2: +3.3% on both factions (horde 648.4 to 670.2 on the final engine). It keeps every unmodeled guide talent, keeps Fire at 31, and Accelerant speeds nothing the level-60 list casts. A Fireball filler still loses to Scorch with Mage Armor (604 against 686). | Adopted (talent rule) | S 9fef30848 |
| F2 | Orc Savage removed (X2). | Fire horde -1.3 to -1.4%. | CONFIRMED-BUG | as X2 |
| F3 | Item 3: Master of Elements on a Clearcasting cast. | No new evidence: no hotfix line, and no level-60 log. The beta logs (levels 13 to 30) cannot show Arcane Concentration and Master of Elements together. | OPEN. Settled by a level-60 log with a Clearcasting (12536) removal, a Fire crit and a Master of Elements energize. | - |
| F4 | Fresh: Critical Mass is 2% a rank. Fire Power is 1% a rank per Fire effect. Ignite is 40%. Improved Fire Blast is 1 s a rank. Heating Up and Combustion are as in round 4. Chthonic Power (Frozen Heart of the Mountain) is +29 Frost and Shadow for 15 s every 5 min. | 70707 rows; `common/item_effects.go:3711`. | Aligned | - |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Touch of the Grave 15 s (X4). | Shadow horde +0.3 to +0.4%. | CONFIRMED-BUG | as X4 |
| P2 | Item 3: Shadowfiend at launch. | No new evidence. Wowhead Forever's 401977 still reads "Requires level 1", and the 30 October hotfix has no Shadowfiend line. | OPEN (unchanged). Settled by a launch trainer list. | - |
| P3 | Fresh: Twin Disciplines' masks. | 1225132 effect 0 (direct) matches Shadow Word: Death. Effect 1 (periodic) matches Mind Flay, Shadow Word: Pain, Devouring Plague and Renew. The healing half matches the triggered Penance and Prayer of Mending heals, Holy Nova and Desperate Prayer. That is what `priest.go:91` and `talents.go:147` give. | Aligned | - |
| P4 | Fresh: the build's other talents. Shadow Focus 1% a rank, Darkness 2%, Shadow Weaving (100% at 3/3, 2% a stack, 5 stacks), Improved Mind Blast 0.5 s, Early Demise 30%, Inner Focus +25% crit, Mental Agility 3/7/10, Shadowform. | 70707 rank descriptions. | Aligned | - |
| P5 | The healing-damage fix (fce7dd53d) moves Shadow. | Spiritual Guidance's stated damage (1/3/5/6/8%) matches 14901 (aura 174 = 8, aura 175 = 25 at rank 5). Shaman's Mental Quickness (30812, 30 and 30) matches too. | Aligned (the warlock lane's fix) | - |

## Priest: Discipline and Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X7b | Holy Binding (1328569) was named but not modeled, on the grounds that nothing damages the priest. | Its text since the 30 October hotfix: "Each time you heal a target other than yourself with Flash Heal, you are healed for 50% as much." Dark Sacrifice and (after X7) Demonic Runes do take the priest's health. It is now a flat, proc-free self-heal. Test: `sim/priest/healing/holy_binding_test.go`. | CONFIRMED-BUG (fresh) | E ea89e98b5 |
| H1 | Item 4: Power Infusion. | No priest list casts 10060 (the Disc, Holy and Shadow APL spell lists). The site's healer options pass an empty `PowerInfusionTarget` (`request.go:717`); `Environment.GetUnit` returns nil for it, so `RegisterPowerInfusion` registers nothing. The raid preset and `simbuffs` carry no Power Infusion. | Aligned with ruling 2 | - |
| H2 | Human Spirit 10% (X5). | Disc alliance (human) +1.3%. | CONFIRMED-BUG | as X5 |
| H3 | Renewed Hope loses Binding Heal; Binding Heal becomes a 6-rank trainer spell. | No list casts Binding Heal. Renewed Hope's 2% a rank and 1 s a rank match. | Aligned | - |
| H4 | Vestments of the Devout. Disc and Holy wear 5 or 6 pieces. | Client set 182: 9346 (+18), 450576 (5% for 200 mana, Sudden Insight), 18378 (mp5). It is modeled: Sudden Insight shows in Holy's mana income. | Aligned | - |

## Before/after

The before column is pin `f8a5f94ce` on origin/main. The after column is the lane engine `c45c6b78e`, which includes `fce7dd53d`, on the site branch. Raid preset, band 60, 10,000 iterations. Each cell is seed 1 / seed 50001. Healers are effective HPS.

| Spec | Alliance before | Alliance after | Change | Horde before | Horde after | Change | Causes |
|---|---|---|---|---|---|---|---|
| Arcane | 601.8 / 602.1 | 585.3 / 585.7 | -2.7% | 605.5 / 605.3 | 585.1 / 585.0 | -3.4% | M1 +2.1, X2 -0.8 (horde), fce7dd53d about -4.6 |
| Fire, published talents | 625.8 / 627.5 | 659.3 / 659.6 | +5.2% | 618.9 / 620.0 | 650.8 / 651.0 | +5.1% | M1 +8 to +9, X2 -1.4 (horde), fce7dd53d |
| Fire, guide raid build to F1 | 653.6 / 653.5 | 672.9 / 673.4 | +3.0% | 651.0 / 652.9 | 670.2 / 670.4 | +2.8% | as above, plus F1 +3.4 |
| Frost | 547.2 / 547.9 | 532.9 / 533.6 | -2.6% | 543.3 / 543.6 | 534.7 / 535.4 | -1.5% | fce7dd53d, M1 +0.2 to +0.5, X3 +0.3 (horde) |
| Shadow | 665.4 / 665.8 | 651.0 / 651.3 | -2.2% | 668.4 / 668.0 | 656.1 / 656.2 | -1.8% | fce7dd53d, X4 +0.35 (horde) |
| Discipline HPS | 526.5 / 526.6 | 536.0 / 536.2 | +1.8% | 543.1 / 543.2 | 545.7 / 545.9 | +0.5% | X5 +1.3 (human), X7 +0.5 |
| Holy HPS | 639.7 / 639.6 | 641.3 / 641.8 | +0.3% | 618.4 / 618.6 | 620.2 / 620.6 | +0.3% | X7 with X7b |

The published BiS files still carry Fire's older talent string, so F1 reaches the published Fire number only after bis.yml regenerates from the guide's raidBuild.

## Still open

- Master of Elements on a Clearcasting cast (F3). Needs a level-60 log.
- Shadowfiend at launch (P2). Needs the launch trainer.
- Horde race for Arcane and Fire after Savage's removal (X9). Owner's call.
- Ice Lance's coefficient, 0.4286, is still unconfirmed (round 4 R3).
- Other lanes should check the 30 October hotfix list (X1) for paladin, druid, warlock and hunter.
- Ladder goldens: mage and priest rotations moved. TestRotationLadder is left for the integrator.

## Sources

- wago.tools db2 CSVs, SpellEffect for builds 1.60.1.70009, 70245, 70374, 70532, 70630, 70634, 70658, 70685 and 70707, plus the 70707 Spell, SpellAuraOptions, SpellClassOptions, SkillLineAbility, SpellLevels and ItemSparse tables (read 2026-11-01). Build list: https://wago.tools/api/builds
- wowforevertalents.com, WoW Forever patch notes, the 31 October and 30 October 2026 entries (read 2026-11-01): https://wowforevertalents.com/patch-notes/
- talentsforever.com, Mage changes (read 2026-11-01): https://talentsforever.com/mage/changes
- MMO-Champion, "WoW: Forever Beta Development Notes - Racial Abilities and Traits Tuning" (26 October, updated 27 October 2026): https://www.mmo-champion.com/content/14645-WoW-Forever-Beta-Development-Notes-Racial-Abilities-and-Traits-Tuning
- Warcraft Tavern, "More Racial Ability Changes for WoW Forever" (27 October 2026): https://www.warcrafttavern.com/forever/news/more-racial-ability-changes-for-wow-forever/
- Wowhead Forever tooltips and spell pages for 11426, 13033, 1326891, 1240047, 401977, 34433 and items 249469, 279246, 11302 (read 2026-11-01): https://nether.wowhead.com/forever/tooltip/spell/11426, https://www.wowhead.com/forever/spell=11426, https://www.wowhead.com/forever/spell=1326891
- kami-labs, "the hidden October 20 hotfix" (22 October 2026): https://kami-labs.fr/en/wow-forever/mage-givre-wow-forever-correctif-discret/
- endgametools, beta hotfix class changes (23 October 2026): https://endgametools.com/en/wow-forever/news/wow-forever-beta-hotfix-class-changes
- The local client's hotfix cache, `_classic_beta_/Cache/ADB/enUS/DBCache.bin` (build 70338, 11 October 2026)
