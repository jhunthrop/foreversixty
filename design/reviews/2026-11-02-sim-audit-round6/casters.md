# Sim audit round 6: casters (Mage and Priest), 2026-11-01

Specs: mage-arcane, mage-fire, mage-frost, priest-discipline, priest-holy, priest-shadow.

## Method

- **Engine.** Lane branch `lane/r6-casters` from the pin `8baffe5fc`. The before column runs a second worktree at the pin against origin/main `21d97955f`.
- **Client.** wago.tools tables for 1.60.1.70707, read 2026-11-01. The hotfix push for 70707 was dumped again from wago.tools/hotfixes: there are no rows newer than round 5's capture (31 October 10:59 UTC). The push now includes the `Spell` table, which carries text changes. Pushes are applied in `push_id` order. The round-5 helper applied them in file order, which put older pushes over newer ones; the rows below were read with the fixed helper.
- **Other sources.**
  - Wowhead Forever server tooltips.
  - forever-bugs #137 and #398.
  - Blizzard's Priest and Mage deep dives.
  - The round-3 cache of foreverlogs.gg beta logs (`scratchpad/r3/casters/fl/`).
- **Runs.** `spec-breakdown`, raid preset, band 60, 10,000 iterations, seeds 1 and 50001, both factions.
- **Tests.** Every confirmed bug got a failing test first. I saw it fail, then saw it pass.
- **Owner ruling 2** (horde race picks within 1%) and **the Healers lane** (which owns Discipline and Holy builds, rotations and the heal profile) are respected. Priest rows here are amounts, coefficients, costs and procs only.

## Assigned items

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | Owner ruling 2: Horde race picks within 1% stay. Undead against Orc on Arcane and Fire (round 5 X9) is closed, and the guides keep `recommendedRaces`. | Ruling recorded and not reopened. The adoption rule's bar applies to race. | Ruled (owner) | - |
| I2 | Fire publishes its new raid build after `8baffe5fc`. | `bis/mage-fire.json`, band 60, both presets and factions: talents `253005000000200000-51521001230133031-003…`, which is the guide `raidBuild` with Improved Fire Blast 2/2 (engine `8baffe5fc`, generated 2026-11-01 20:31). The CDN guide fixture and the ladder golden carry it too. | Aligned | - |
| I2b | **The in-game addon still carried Fire's pre-round-4 raid build**: Accelerant 4, no Arcane Meditation, 14/34/3. Beast Mastery, Retribution, Subtlety, Affliction and Destruction were stale the same way. | `addon-builds.json`, which Data.lua's defaultBuilds read, was written only at client activation (`ab5cf670b`, 31 October). No workflow regenerates it. Test `TestCommittedAddonBuildsFollowTheGuides` (sim/cmd/addon-builds) failed, then passed after regeneration. Data.lua picks the file up at the next nightly. | CONFIRMED-BUG (data, no DPS) | S dacbb98a1 |
| I3 | Benediction's +65 / +195 in simdb. | data.yml regenerated simdb at `d40f4d6a4` (2026-11-01 20:07 UTC), after `f52842d53`. Benediction 18608 now has SpellPower 65 and HealingPower 195. The pre-round-5 simdb (`c1810873c`) had neither. The bis.yml run at `21d97955f` ranked on the new simdb. No priest set wears Benediction: it is a quest reward, so it is not in `loot.json`'s candidate pool, and all three priest sets keep Glowing Brightwood Staff 812 or Hand of Edward the Odd 2243. Worn in place of the staff, horde: Holy 623.2 to 645.2 (+3.5%), Discipline 548.0 to 564.2 (+3.0%). The other 66 published band-60 caster items match Wowhead Forever, for example Prelacy Cape 7004 and Glowing Eye of Mordresh 10769 after the hotfix rows. | Aligned (reached simdb). Candidate pool is the ranker's call. | - |
| I4 | Molten Armor to Mage Armor is consistent. | `sim/request/request.go:706` sends Mage Armor, and `TestEveryMageSpecWearsMageArmor` holds it. No APL casts an armor. The only guide mention is Fire's Mage Armor sentence. No Data.lua, addon or web source names Molten Armor. Engine: Mage Armor 22783 has aura 134 = 50 and resistance 15, and it adds to Arcane Meditation's aura 134. Below level 34 a mage wears no armor, which costs no damage. | Aligned | - |
| I5 | Shadowfiend at launch (single source). | No second source. Blizzard's Priest deep dive lists every new, baseline and changed priest spell (Shadow Word: Death baseline at 32, Divine Spirit at 30, and so on) and does not name Shadowfiend. Wowhead Forever's 401977 still reads "Requires level 1". Its SkillLineAbility row (line 78, acquire 0) has the same shape as Shadowguard's. | OPEN. Settled by a launch trainer list or a level-60 log with a 401977 cast. | - |
| I6 | Forever DoTs read spell power at each tick (forever-bugs #398, single source in round 5). | **Second source found.** In forever-bugs #137 (2026-09-25), the reporter's test shows DoTs already ticking pick up Inner Focus's crit. Blizzard (AidanZMoon, 2026-10-01) answers "Since we've changed to dynamic DoTs" and removes Inner Focus's crit from periodic effects. In #398, a combat log shows an Agony that snapshots "unlike other dot spells". Blizzard (2026-10-26): "adjust every tick to the spell power of the caster". The engine snapshotted spell power and the caster's multipliers at application; crit was already rolled per tick. Fix: `DotConfig.Dynamic` re-takes the snapshot before every tick. It is set on every mage and priest dot whose snapshot is plain. Test `TestDynamicDotReadsSpellPowerEachTick` (core). | CONFIRMED-BUG (two threads, one with a log) | E 9a5c25dff |
| I7 | Tauren 1349372, +2% critical healing (single source). | The push's own Spell row for 20550 (push 113276, 2026-10-31) rewrites the text to "Total Health increased by $s1%." and removes effect 1322984 (critical damage). Blizzard's 26 October notes (MMO-Champion, Warcraft Tavern, wow.gg) say the same, and so does Wowhead Forever's 20550 tooltip. Only the aura-50 effect row 1349372 survives. The engine gives health only. | Aligned (three statements against one leftover row). A Tauren healer's critical heal size would settle it. No published caster is Tauren. | - |

## Mage: Frost

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Ice Lance's Frozen bonus multiplied only the base roll.** The spell-power share was added after the ×4, so a Frozen Lance at raid spell power dealt about 1.8× an unfrozen one. | 1240047 effect 1 (pushed 200 to 300) reads "Deals $s2% increased damage to Frozen targets", which is the hit's damage. Beta logs (report 2791, Talsyn, rank 400640, before 70685 when the bonus was ×3): unfrozen median 51, Frozen median 152, ratio 2.98. A base-only bonus predicts about 2.5. `sim/mage/ice_lance.go` now multiplies the spell's damage for the cast. Test `TestIceLanceFrozenBonusMultipliesTheWholeHit` failed at 1.77×, then passed. TestP1Mage Frost goldens regenerated. | CONFIRMED-BUG (fresh; client text plus logs) | E b427bc6f7 |
| R2 | Ice Lance's coefficient. The client column is 0 on every rank, while Arcane Blast, Fire Blast and Frostfire Bolt carry 0.429 / 0.429 / 0.814. | The engine's 0.4286 is the instant-spell convention. The same logs show unfrozen hits about 12 above the rank's base range, so Ice Lance does scale with spell power, but one character with unknown spell power cannot fix the value. R1 makes this worth more: Ice Lance's share of Frost damage rose. | OPEN (round 4 R3). Settled by a log with a known spell-power sheet. | - |
| R3 | Fresh, checked and aligned (client plus push): Frostbolt 10181, Cone of Cold 10161 (10 yd, 12 with Arctic Reach, whose push mask adds Ice Lance), Ice Barrier with Frigid Power, Winter's Chill, Fingers of Frost, Shatter, Ice Shards, Frost Channeling, Precision. | Frostbolt: 386, ±7.6%, +2.9 a level, coefficient 0.814, 260 mana, 3 s. Missile speeds (Frostbolt 28, Ice Lance 38) match SpellMisc. Frigid Power is 15% and Ice Barrier's own effects 5% (pushed values). Winter's Chill is 20% and 7% a stack. Fingers of Frost is 15% a Chill, 2 charges, 15 s. | Aligned | - |
| R4 | Robe of the Archmage (see M2). | Frost wears it too. | CONFIRMED-BUG | as M2 |
| R5 | Raid talent search and rotation search after R1 (lane engine `fc3d8aadf`, site `dacbb98a1`, horde, `-keep` from raid-keeps.json). | Talent search: no variant beats the guide beyond error; the best is +2.0 ± 2.2 (590.6 guide). Rotation search: the curated rotation is the best found, +1.9 ± 2.4. Ice Lance is now worth 74 DPS by removal. | Aligned (guide build and APL kept) | - |
| R6 | Ice Lance coefficient weight. | With the coefficient at 0, horde Frost drops from 590.1 / 590.1 to 532.6 / 533.1 (-9.7%), and Ice Lance is 15.6% of Frost damage. A second source would settle it: a level-60 Forever log of Ice Lance hits from a character whose spell power is known (a warcraftlogs sheet, or one target dummy test with and without a spell-power item). | OPEN | - |

## Mage: Fire

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **Ignite stacked the Classic way.** It took up to five stacks, and every refresh restarted both ticks for all earlier stacks, so a crit landing between ticks paid the earlier crits' burn again. | **Client:** talent 11119 points at 412538, a 4 s, 2 s-tick burn of `$w1/2` with CumulativeAura 1 (no stacks); its damage spell is 412545. **Logs:** 55 encounters of one Fire mage (Meetch, round-3 cache, spell 12654). Ignite is only ever `debuff_applied` and `debuff_refreshed` (amount 1), never stacked, with up to six crits in one burn. The burn's share of its triggering hits does not rise when ticks fall between crits: a median of 0.73 with one crit, 0.74 with two and a tick between, and 0.69 with three. A new crit now adds its 40% to what the burn has left and spreads the sum over a fresh 4 s. Test `TestIgniteRollsCritsIntoOneBurn` failed (661.7 against 502.2), then passed. | CONFIRMED-BUG (fresh; client plus logs) | E 3515fd706 |
| F1c | Integrator check (from the warlock lane's scan): Ignite registered two spells under 12654, at the pin and still after F1. | One was the trigger spell and one was the burn. `GetSpell(12654)`, and so any APL line naming Ignite, resolved to the trigger, which has no dot. The trigger is removed: a Fire crit adds its share to the burn directly. Test `TestIgniteRegistersOnce` failed with 2 spells, then passed. No number moves, because no APL names 12654 and the trigger dealt nothing: Fire stays at 648.3 / 649.6 alliance and 656.7 / 657.6 horde, seeds 1 and 50001, before and after. A scan of the Spellbooks (a mage with Ignite, and healer priests on the Discipline and Holy builds) finds no other duplicate id. Shadow priest was not scanned. | CONFIRMED-BUG (registration; no DPS) | E 74f7fdbd6 |
| F6 | Integrator check (wow-player review): /sim at Fire band 40, bare, says "Scorch and Fire Blast are 62% of your damage; Heating Up is up 71% of the fight". | **Breakdown** on the int engine `b57765fb9` plus site `0ee44a9fd`, horde, 10,000 iterations: Scorch rank 4 (8446) 41.5%, Fire Blast rank 5 (8413) 22.3%, Pyroblast rank 4 (12523) 20.6%, Ignite 10.3%, wand 4.5%. Heating Up is up 71.1%. Alliance is the same within 0.2%. The ranks are right for 40: Fireball rank 7 (8402, level 36) is learned but never cast, because the list puts Scorch above it. The page merges the right spells; the shot ran 8baffe5fc, hence 62% against 64% now. **Client:** at 40 Scorch is 93 base (+1.2 a level), 1.5 s, 100 mana, coefficient 0.429. Fireball is 187 base (+2.3 a level), 3.5 s, 220 mana, coefficient 1.0, plus a 6-a-tick burn. So Scorch deals more per second and per mana, and only Scorch and Fire Blast take Incineration's +10% crit, which feeds Heating Up. **Measured** at band 40, bare, horde, seeds 1 / 50001: curated 142.3 / 142.4; Fireball as the filler 139.9 / 139.9 (-1.7%); without Fire Blast 139.9 / 140.2 (-1.6%). Warcraft Tavern's Forever Fire guide casts Fireball once the stacks are up, and that measures lower here. | Aligned. Reason a player accepts: "In Forever a level-40 Scorch hits for about half a Fireball in less than half the cast time and half the mana, and only Scorch and Fire Blast get Incineration's crit that feeds Heating Up, so Scorch is the better filler." | - |
| F1b | The logs' ratio of Ignite to its triggering hit (about 0.73) is higher than 0.40. | It cannot be read without the target's debuffs and whether the logged hit is the crit. The round-2 note "Ignite no longer double dips" stands. | OPEN. A level-60 log with the target's debuffs settles it. | - |
| F2 | **The raid preset gave every spec a raid mage's Improved Scorch.** | 22959 is aura 270, "Increases Fire damage taken from the Mage" (client). Blizzard's Mage deep dive: the debuff "now benefits only the casting Mage" (Warcraft Tavern recap; three Forever Fire guides agree). `data/curated/presets.json` listed `improved_scorch` for the whole raid. A Fire mage still stacks her own through Scorch. Test: `TestRaidPresetNamesTheRequiredBuffsAndExcludesTheForbiddenOnes` bans it (it failed, then passed). | CONFIRMED-BUG (fresh, two sources). **Moves other groups**, see the table below. | S d1ddb012a |
| F3 | Fresh, checked and aligned: Fireball 10151, Fire Blast 10199 (8 s on category 19 with Ice Lance, -1 s a rank of Improved Fire Blast), Scorch 10207, Pyroblast 18809, Heating Up, Improved Scorch, Combustion, Critical Mass, Fire Power, Master of Elements, Incineration, Evocation. | Fireball: 451, ±24%, coefficient 1, 395 mana, 3.5 s, with a DoT of 14 a tick. Pyroblast: 583, DoT 53 a tick at 0.15, 6 s cast. Heating Up: -25% a stack, 3 stacks, 20 s. Combustion: +10% a stack, ends on the third crit. Fire Power: 1% a rank per Fire effect, up to 4. | Aligned | - |
| F4 | Master of Elements on a Clearcasting cast (round 5 F3). | No new evidence. | OPEN. Needs a level-60 log. | - |
| F5 | Raid talent search and rotation search after F1 and F2 (lane engine `fc3d8aadf`, site `dacbb98a1`, horde, `-keep` from raid-keeps.json). | Talent search: the guide raid build is the best build found (655.6 ± 1.5). Rotation search: dropping the Improved Scorch upkeep line (#4) gains +6.1 ± 2.1 at 800 iterations, because the Scorch filler keeps the stacks anyway. At 10,000 iterations it is +0.4 / +0.3% alliance (648.3 / 649.6 to 651.0 / 651.5) and +0.5 / +0.4% horde (656.7 / 657.6 to 660.2 / 660.3). That is beyond error but under the 1% bar the lane holds builds to, and the line was restored deliberately in an earlier fix round. | Hold (+0.4%, below 1%). The guide build and APL are kept. | - |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Arcane Blast travelled at speed 20**, borrowed from Arcane Missiles' bolt. | SpellMisc Speed is 0 on every Arcane Blast rank (400574, 1239696 to 1239700), as on Fire Blast and Scorch. The engine's own comment called the 20 unconfirmed. Test `TestArcaneBlastLandsOnCast`. | CONFIRMED-BUG (fresh, client) | E fc3d8aadf |
| M2 | **Robe of the Archmage's use was unmodeled.** All three mage sets wear the robe. | ItemEffect 99905 is spell 18385, 500 mana with variance 0.5 (375 to 625), on a 5 min item cooldown; Wowhead Forever says the same. 18385 is in category 30 with a 60 s category recovery, shared with the mana gems, Demonic Rune and healthstones, so a use holds the conjured cooldown for a minute. It is a mana cooldown that fires when the bar has room for the whole roll. Test `TestRobeOfTheArchmageRestoresManaAndHoldsTheGems`. TestP1Mage goldens regenerated (the suite's gear wears the robe). | CONFIRMED-BUG (fresh, two sources) | E 7ae2a6a9b |
| A2 | Missile Barrage rolls on cast complete, so it can proc on a miss. | 400588 ProcTypeMask 65536 is a harmful-spell-damage trigger. Whether a miss counts is not stated. | OPEN (small) | - |
| A3 | Arcane Concentration can proc from each Arcane Missiles bolt (`talents.go:196` excludes only the channel). | Nothing in the client settles it. | OPEN (small) | - |
| A4 | Fresh, checked and aligned: Arcane Blast 1239700 and its stack aura 400573 (pushed: 12% of Intellect, +100% cost, 4 stacks, 18 s), Arcane Missiles 10212 and its bolt 10274, Missile Barrage, Arcane Power, Presence of Mind, Arcane Concentration, Arcane Mind, Arcane Instability, Arcane Impact, Arcane Meditation, Magister's Regalia 3-piece. | Arcane Blast: 236, ±15%, coefficient 0.429, 15% of base mana, 1.5 s. Missile bolt: 171, coefficient 0.286. Arcane Power and Presence of Mind share category 1151 with Combustion, but no build takes both. Evocation below 5 / 15 / 30% mana measured lower than no Evocation (582 / 572 / 575 against 588). | Aligned | - |
| A5 | Racials of the published races. | Orc Blood Fury 20572 (10% attack power and spell power, 15 s, 2 min), Troll Berserking 15 s, Gnome Eureka (round 4 generator). | Aligned | - |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Dynamic DoTs (I6): Shadow Word: Pain, Devouring Plague, Mind Flay. | Shadow horde +0.13 / +0.17%. | CONFIRMED-BUG | as I6 |
| P2 | Mental Agility's mask (see H1). | It now leaves out Mind Flay and Shadow Word: Death. Shadow -0.05% (never short of mana). | CONFIRMED-BUG | as H1 |
| P3 | Fresh, checked and aligned: Mind Blast 10947, Shadow Word: Pain 10894, Devouring Plague 19280, Mind Flay 18807, Shadow Word: Death 1309636, Shadowform 15473, Darkness, Shadow Focus, Shadow Weaving, Twin Disciplines, Inner Focus, Early Demise, Improved Mind Blast, Improved Shadow Word: Pain, Improved Mind Flay, Touch of the Grave. | Mind Blast: 485, coefficient 0.429, 350 mana, 8 s, -2.5 s with Improved Mind Blast. Shadow Word: Pain: 127 a tick at 0.2 over 18 s, +6 s talented. Devouring Plague: 106 a tick at 0.1, 8 ticks. Mind Flay: 130 a tick at 0.167, 3 ticks, crits by flag. Shadow Word: Death: 448, 30% backlash (pushed). Shadowform: +10% Shadow damage and +100% crit bonus. Inner Focus keeps its crit off periodic effects (`talents.go:415`, Blizzard 1 October). | Aligned | - |
| P4 | Mind Flay's 20 yd range against the ranker's 30 yd stand. | The engine does not check Mind Flay's range, so the published number assumes a priest in reach. | Noted | - |

## Priest: Discipline and Holy (amounts, coefficients and procs)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Mental Agility cut every instant**, including Penance, Mind Flay, Shadow Word: Death and Desperate Prayer. | 14520, SpellEffect 687728 (aura 108, misc 14, no hotfix row). Its class mask names Smite, Holy Fire, Shadow Word: Pain, Devouring Plague, Renew, Power Word: Shield, Holy Nova, Prayer of Mending and Vampiric Embrace (each checked against SpellClassOptions). Penance (mask 1 bit 23) is outside it. The tooltip's "instant cast spells" is wider than the mask the server applies. Test `TestMentalAgilityFollowsTheClientMask`. The unused instant-cast mask is deleted. | CONFIRMED-BUG (fresh, client) | E 207bf876f |
| H2 | **Spiritual Healing multiplied all of the priest's healing.** | 14898 has two masked modifiers: direct heals (692104: Lesser Heal, Heal, Greater Heal, Flash Heal, Prayer of Healing, Holy Nova, Desperate Prayer, Penance, Prayer of Mending, Binding Heal) and Renew (1239404). Lightwell's renew 27874 and Holy Binding have no class options. Test `TestSpiritualHealingFollowsItsClientMasks`. | CONFIRMED-BUG (fresh, client) | E 4cae3d570 |
| H2b | Which mask Prayer of Mending's bounce heal reads. The cast (mask 1 bit 5) is in Spiritual Healing's mask. The heals 401863 / 1313379 / 1313381 carry only mask 3 bit 21, which neither Spiritual Healing nor Holy Specialization names. | The server reads one of the two; the client cannot say which. The engine keeps both talents on the heal. | OPEN (at most 10% healing and 5% crit on Prayer of Mending) | - |
| H3 | Fresh, checked and aligned: Heal 6064 (651, +4.5 a level to 39, 0.857), Flash Heal 10917, Prayer of Healing 10961, Renew 6075, Power Word: Shield and Soul Warding, Penance 1316995 / 1316991 (3 pulses of 425 at 0.19, 12 s), Prayer of Mending 1240827, Lightwell 27871, Dark Sacrifice with the Demonic Rune's 800 health, Holy Binding 50%, Improved Healing, Divine Fury, Improved Renew, Renewed Hope, Divine Aegis, Litany of Light, Meditation, Mental Strength, Spiritual Guidance (14901), Holy Specialization, Vestments of the Devout (182: 9346, 450576 Sudden Insight, 18378). | Every mana spend in the Discipline and Holy runs equals the cost times its modifiers. | Aligned | - |
| H4 | Lightwell is legal. The deep dive's "Lightwell (removed)" means the talent was removed: the spell became a baseline trainer spell (SkillLineAbility line 56, acquire 0, ranks 724 / 27870 / 27871). | `trainables/priest.json` | Aligned | - |

## Before/after

Before: pin `8baffe5fc` on origin/main `21d97955f`. After: lane engine `fc3d8aadf` on site branch `dacbb98a1` (`74f7fdbd6`, the Ignite registration fix, moves nothing). Raid preset, band 60, 10,000 iterations. Each cell is seed 1 / seed 50001. Healers are effective HPS.

| Spec | Alliance before | Alliance after | Change | Horde before | Horde after | Change | Causes |
|---|---|---|---|---|---|---|---|
| Arcane | 585.3 / 585.7 | 586.8 / 587.0 | +0.2% | 588.0 / 588.1 | 589.5 / 589.3 | +0.2% | M2, A1 |
| Fire | 687.8 / 689.4 | 648.3 / 649.6 | -5.8% | 699.7 / 699.1 | 656.7 / 657.6 | -6.0% | F1 about -4.3, F2 about -1.0, I6, M2 |
| Frost | 542.0 / 542.8 | 584.3 / 584.6 | +7.8% | 546.9 / 547.6 | 590.1 / 590.1 | +7.8% | R1 about +7.8, M2 |
| Shadow | 707.5 / 708.0 | 707.8 / 708.2 | +0.0% | 713.0 / 713.1 | 713.9 / 714.3 | +0.15% | I6, H1 |
| Discipline HPS | 538.4 / 538.5 | 528.9 / 528.9 | -1.8% | 548.0 / 548.2 | 537.2 / 537.5 | -2.0% | H1 (Penance at full cost), H2 |
| Holy HPS | 641.4 / 641.9 | 641.2 / 641.4 | -0.05% | 623.2 / 623.5 | 622.7 / 623.3 | -0.05% | H2 |

**F2 (S d1ddb012a) moves specs outside this group.** Measured horde, seed 1, 5,000 iterations, lane engine, with and without `improved_scorch` in the raid preset:

| Spec | With | Without | Change |
|---|---|---|---|
| Destruction | 673.0 | 617.4 | -8.3% |
| Elemental | 494.3 | 467.4 | -5.4% |
| Demonology | 647.2 | 629.9 | -2.7% |
| Affliction | 573.7 | 564.5 | -1.6% |
| Survival | 781.3 | 771.2 | -1.3% |
| Enhancement | 634.1 | 630.9 | -0.5% |
| Retribution | 658.7 | 658.7 | 0.0% |

## Still open

- Ice Lance's coefficient: the client shows 0, the engine uses 0.4286 (R2). It is now worth 9.7% of Frost (R6).
- Ignite's logged ratio of about 0.73 (F1b), and Master of Elements on a Clearcasting cast (F4). Both need level-60 logs.
- Prayer of Mending's bounce-heal mask (H2b).
- Missile Barrage on a miss (A2), and Arcane Concentration from each bolt (A3).
- Shadowfiend at launch (I5).
- Dynamic DoTs outside mage and priest: warlock curses, Siphon Life, Drain Life, Ignite, and every heal over time rewrite their snapshots and need their own per-tick read. These belong to their lanes and to the Healers lane.
- The page lane should check the band-40 /sim shot (r6-page shots, Fire 40 at 390 px): its Damage tab shows one row named "Build" at 100% instead of the spells.
- Fire's rotation-search winner (drop the Improved Scorch upkeep line, +0.4%) is held under 1% (F5).
- Routed to the integrator:
  - Holy Nova is one spell for heal and damage, so Spiritual Healing also raises its damage half. No list casts it.
  - The healer notes the Healers lane owns: Discipline's "Prayer of Healing at rank 1" note is stale, Holy's notes overstate the top ranks, and the coach files expect book ranks 25315 and 25316.
  - Three items carry ranged attack power doubled in simdb.bin: Rune of the Guard Captain 19120 and Cleansed Vilebranch Medallion 270275 (42 to 84), and Graverobber's Shovel 273817 (12 to 24).
  - Benediction is outside the ranker's candidate pool (I3).
- The ladder goldens and the item-effect id list follow the pin (integrator).

## Sources

- wago.tools db2 CSVs for 1.60.1.70707 (Spell, SpellName, SpellEffect, SpellMisc, SpellCooldowns, SpellCategories, SpellPower, SpellClassOptions, SpellLevels, SpellAuraOptions, SkillLineAbility, Talent, ItemEffect, ItemXItemEffect), read 2026-11-01. Hotfix pushes from https://wago.tools/hotfixes (re-dumped 2026-11-01; no rows newer than 2026-10-31 10:59 UTC).
- forever-bugs #137 and #398 (read 2026-11-01): https://github.com/ClassicWoWCommunity/forever-bugs/issues/137, https://github.com/ClassicWoWCommunity/forever-bugs/issues/398
- Blizzard, WoW Forever class deep dive, Priest and Warrior: https://news.blizzard.com/en-us/article/24301514
- Warcraft Tavern, Mage and Shaman deep dive recap (Improved Scorch benefits only the casting mage): https://www.warcrafttavern.com/forever/news/scrolls-totems-mage-shaman-deep-dive-for-wow-forever/
- Racial notes, 26 October 2026: https://www.warcrafttavern.com/forever/news/changes-to-racial-abilities-in-wow-forever-ahead-of-name-reservations/ and https://wow.gg/guides/wow-forever-racial-changes-26-10-2026
- Wowhead Forever tooltips (read 2026-11-01): https://nether.wowhead.com/forever/tooltip/spell/20550, /spell/11119, /spell/412538, /spell/401977, /item/14152, /item/10769, /item/7004
- foreverlogs.gg beta logs as cached in round 3 (reports 2791 and 2807)
