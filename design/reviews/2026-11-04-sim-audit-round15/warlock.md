# Sim-accuracy round 15: warlock (2026-11-03)

Specs: Affliction, Demonology, Destruction, bands 20 to 60, bare and raid. Client 1.60.1.70755. Site `wt/r15-warlock` from int/r15 fd320b60e; fork `lane/r15-warlock` from int/r15 d1d9f2775.

Method: I checked the three seed items first. Then I worked through the guide-accuracy round 1 sim rows (S1 to S6) and did a fresh pass. 70755 changed exactly one warlock spell row: Wrack 1316697 is the only `spellconst/warlock.json` difference from 70753, and trainables are identical. In the talent text, Wrack went from 36 to 86 a tick and Shadow and Flame from 2% to 4% a rank. The engine already follows both: `wrack.go` reads 36 + 2.5 a level and coefficient 0.286, `TestWrackTickReadsThe70755Row` covers it, and `shadowAndFlameDamagePerRank` is 0.04 (`talent_rates_test.go:41`). What nobody had revisited was the rotation and the builds on top of the new Wrack. spec-breakdown ran at 10,000 iterations, seeds 1 and 50001.

Totals: **3 CONFIRMED-BUGs** (Affliction: 2, warlock-wide: 1).
- Wrack's buff made it Affliction's filler, but the rotation never cast it. The fix also changes the raid build. Origin `new-data`/`70755 2eaae03c6`.
- Bane of Doom beats Agony for Affliction too. Origin `old`, `guides-handoff`.
- Curse of the Elements now goes before Corruption. Origin `old`, `guides-handoff`. The sim ties on this one; it was decided on real play.

No engine code changed. Only the fork's APL copies moved.

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Wrack is now the filler, and the rotation never cast it.** 70755 made Wrack 86 a tick at 60 with coefficient 0.286 a tick: 516 + 1.716 × spell power over its 6 second channel, for 200 mana. Soul Siphon adds 36% once Corruption, the bane and Siphon Life are up. That beats Shadow Bolt r9 (251 + 0.857 × spell power on a 2.5 s cast, 370 mana) per second of casting, at about a quarter of the mana. Before the push (36 a tick, 0.143) Wrack lost by about half, which is why it was left off. The curated rotation had no Wrack line, so the leveling bands that take Wrack (bare 40/50/60, through the walk's reservation) carried a dead capstone. The rotation now casts Wrack right before Shadow Bolt; nothing casts Wrack without the talent (bands 20/30 and the other specs are unchanged). The raid build moves to `25520020135001351/00052/005004`. It takes Soul Siphon 3, Wrack 1, Improved Bane of Agony 2 and Aftermath 4. It gives up Ruin 4, Improved Shadow Bolt 5 and Improved Bane of Agony's first point: with Wrack as the filler, Shadow Bolt is never cast, so Improved Shadow Bolt and Ruin have nothing to work on. It keeps Grim Reach 2 and Demonic Aegis 2 (unmodeled keeps) and Unholy Power 5. Every tier gate is checked. Five candidates were measured (Improved Shadow Bolt 4, Nightfall, Cataclysm, Bane 2); Bane 5 with Aftermath 4 was best. Adoption rule: same spec tree, keeps hold, nothing the rotation casts is dropped, +13.6% at both seeds. Test first: `TestAfflictionCastsWrackWhenItsBuildTakesIt` (sim/cmd/spec-breakdown) failed at bare 50, bare 60 and the guide raid build (0 casts), and passes now. The smoke test now lists Wrack as a talent spell. | spellconst/warlock.json 1316697, 11661; talents/warlock.json Soul Siphon 105911; `wrack.go`; scratch test: tick 86 × 1.1 at 0 spell power | CONFIRMED-BUG | new-data / 70755 2eaae03c6 (the engine followed the push; no one re-ran the filler); guides-handoff (S2) | S 9e270b86c, E 86078f44e |
| A2 | **Bane of Doom beats Bane of Agony for Affliction too.** Doom (603: 1742 + 4.0 × spell power after 60 s, one global, 300 mana, crits with Pandemic) goes up while a minute or more is left, and Agony covers the last minute, as in Demonology and Destruction. On the Wrack build this was +1.7% (654.4 to 665.5, 654.1 to 665.8). Amplify Curse was cast before the pull. The bane slot holds one bane, so Amplify then ran out on Doom unused (16.7% uptime, no Agony). It now waits for the last minute and goes on the last Agony (0.98 casts, 0.3% uptime, so it is spent at once). Build variants without Improved Bane of Agony or Amplify Curse measured within ±0.2%, so the build holds. Test first: `TestAfflictionRaidRotationKeepsBaneOfDoomAndAmplifiesTheLastAgony` failed (0 Doom casts, 8.4 Agony casts) and passes now. | spellconst 603, 18288; `curses.go:398-470`; runs | CONFIRMED-BUG | old (the Agony choice dates from the 2026-10-31 audit, when Wrack sat in the build); guides-handoff (WARLOCK-13 / S3) | S fee3e523d, E 04199f9b4 |
| A3 | S5 (Grim Reach 1 + Improved Bane of Agony 2): moot. The new raid build takes Improved Bane of Agony 2 and still has Grim Reach 2. | | Aligned | guides-handoff | - |
| A4 | Shadow Trance: an instant Shadow Bolt line before Wrack measured -0.4% at bare 60 (375.9 to 374.3). The mana it costs is worth more Life Taps than the Shadow Bolt's edge over a Wrack tick. Nightfall stays in the leveling build only. No change. | runs | Aligned | - | - |
| A5 | Wrack details: it is a channel; its ticks crit (`periodic_crit_auto_gen.go:388`, ATTR8 PERIODIC_CAN_CRIT) and Pandemic doubles them; its vulnerability covers Corruption and Agony only (0x402); Soul Siphon counts Corruption, the bane and Siphon Life (`talents.go:1132`), so Doom counts too; Dynamic re-snapshots before the per-tick Soul Siphon read, so there is no compounding (`dot.go:268`). | client rows; engine | Aligned | - | - |
| A6 | Leveling build `25522300130201351/200522/005`: Improved Drains 3 and Soul Harvest 2 still do nothing in a boss fight. They are drain-tanking and questing picks (the guide levels on Drain Life). Drain Life r6 is 51 a tick at coefficient 0.1, far under Wrack. This is a leveling taste call, not provably dead. Unchanged. | spellconst 11700 | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Fresh pass: no warlock spell row other than Wrack moved in 70755, and Demonology casts no Wrack. Soul Fire and Decimation, Doom and Agony, Immolate, Shadow Bolt, the Succubus and Soul Link match the client rows checked in round 14 (D1) and are unchanged in 70755. The raid build's idle points are the same keeps as in round 14 (D5). | spellconst diff 70753 to 70755 | Aligned | - | - |
| D2 | Curse of the Elements order: see C1. Bare bands move within ±0.14%; raid 60 does not move. | runs | (C1) | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Seed 1, Molten Skin 3 and Destructive Reach 2 in the raid build.** Both do nothing in the sim: one is damage taken, the other range. Dropping them for Aftermath 5, Pandemic 3 and Malevolence 3 (`25500000033/0/0050355103101351`, every gate checked) gains +1.73% at both seeds (684.1 to 695.9, 683.8 to 695.6). Both talents are on the raid search's unmodeled keep list (round 13 X4), so under the adoption rule ("keeps every unmodeled talent") this is a survival and range choice, not a dead point. The guide lane already asked the owner (owner question 4). A web search found no Forever raid build either way. Within the keeps, the current split is best: Cataclysm 2 / Aftermath 5 measured -0.25% and Pandemic 3 / Aftermath 3 -0.19%. Held, pending the owner. | runs; design/reviews/talent-search/raid-keeps.json; [kami-labs Forever Destruction page](https://kami-labs.fr/en/wow-forever/build-demoniste-destruction-wow-forever/) (pre-beta, no raid build) | OPEN (owner taste call) | guides-handoff (S4) | - |
| X2 | Shadow and Flame on 70755: 4% a rank (20% at 5), 20 s, and 20% a rank not to consume Immolate. The engine is already at 0.04 (`talents.go:959`). The leveling build stays legal: every gate holds and nothing is misplaced. | talents/warlock.json 105875 | Aligned | - | - |
| X3 | Curse of the Elements order: see C1. | | (C1) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Curse of the Elements before Corruption** in all three rotations (S1). The holder's curse raises every magic school the raid deals from its first global. A one-player sim cannot see that: the personal order measures within ±0.22% at every bare band. The line still skips when the raid's copy is on the target, so raid 60 does not move. Under the round-15 rule (the sim ties, so decide on real play), the holder curses first. Test first: `TestWarlockRotationsCurseOfTheElementsBeforeCorruption` (sim/request) failed for all three and passes now. | runs; [warcraft.wiki.gg Curse of the Elements](https://warcraft.wiki.gg/wiki/Curse_of_the_Elements); [Icy Veins Classic warlock rotation](https://icy-veins.com/wow-classic/warlock-dps-pve-rotation-cooldowns-abilities) ("maintained" on bosses; the mechanic is unchanged in Forever except the separate curse and bane slots) | CONFIRMED-BUG (real play; the sim ties) | old; guides-handoff (S1) | S 849b5e99b, E 6071980cb |
| C2 | Logs: `/v1/reports/recent` returns one report (2026-10-20, already examined). foreverlogs.gg's newest report is 2812 (2026-10-29), and every id from 2815 to 3500 is "Report not found". Beta logs run to about level 30, and Wrack is level 40, so no log can show Wrack or level-60 Doom. Warcraft Logs is blocked (gcloud login expired); I did not try it. | API and foreverlogs, 2026-11-03 | OPEN | - | - |
| C3 | Web check for A1 (round-15 rule): no Forever source covers the 70755 Wrack values. [wow.gg's Affliction overview](https://wow.gg/guides/warlock-affliction-forever-overview) (2026-09-19, before the push) and [Mobalytics](https://mobalytics.gg/wow-forever/guides/warlock-class-overview) both describe Forever Affliction as built around stacking effects and draining rather than Shadow Bolt. A1 rests on the client rows and the runs, not on these. | | note | - | - |
| C4 | S6 (weights and dps-per-point carried from 70753): refresh, below. | | Measure at refresh | guides-handoff | - |
| C5 | "Raid build moves N points" miscount (platform lane): neither warlock build string causes it. Both are legal 51-point strings (the decoder checks every gate). The new Affliction raid build moves 16 points from the leveling build. | `enc.py` gate check | note | - | - |

## Before/after

10,000 iterations. Before: $R/before.tsv (raid 60) and $R/sb-before from $R/site-before/sim (bare). After: this branch at 849b5e99b with engine 6071980cb.

| Spec, preset, band | Horde s1 | Horde % (s1 / s50001) | Alliance % (s1 / s50001) |
|---|---|---|---|
| Affliction raid 60, published talents | 571.7 to 592.3 | +3.60 / +3.59 | +3.63 / +3.60 |
| Affliction raid 60, guide raid build (old to new) | 575.8 to 667.2 | +15.9 / +15.9 | +16.5 / +16.5 (574.0 to 668.8) |
| Affliction bare 40 / 50 / 60 | 139.6 to 166.6 / 207.0 to 257.7 / 307.4 to 375.9 | +19.3 / +24.5 / +22.3 (s50001 the same within 0.1) | - |
| Affliction bare 20 / 30 | unchanged | 0.00 | - |
| Demonology raid 60 | 639.4 to 639.4 | 0.00 / 0.00 | 0.00 / 0.00 |
| Destruction raid 60 | 684.1 to 684.1 | 0.00 / 0.00 | 0.00 / 0.00 |
| Demonology and Destruction bare 20 to 60 | | within ±0.22 (C1) | - |

The published Affliction number moves only through the rotation (Doom; the published talents take no Wrack). The new raid build's +16% reaches the page at the refresh.

## Files to regenerate

- data/builds/1.60.1.70755/addon-builds.json (`go run ./sim/cmd/addon-builds -repo-root .`): `TestCommittedAddonBuildsFollowTheGuides` fails until it is regenerated, because the Affliction raid build changed.
- Ladder goldens (sim/request TestRotationLadder): warlock-affliction (Wrack, Doom, Amplify Curse), and warlock-demonology / warlock-destruction for the curse order. I did not run them.
- Coach starters for the three warlocks (pipeline.coachrules), and the addon data / Data.lua at the next pipeline run.

## Measure at refresh

- Affliction, every band from 40 and raid 60: the new raid build and the Wrack filler change the stat weights (no Shadow Bolt cast time; Wrack ticks at 0.286), set DPS, picks, tiers and dps-per-point. Hit still matters, because Wrack can miss on the cast.
- Affliction raid 60: Doom in place of Agony (crit on a 60 s tick).
- The round-14 rows still stand (Orc for Horde Affliction and Destruction, Destruction crit gear, carried 70753 weights: S6).

## For other lanes

- None. No shared code changed. Nothing outside the warlock specs moves.

## For the guide lane

Every sentence below is wrong after A1, A2 or C1 (web/src/content/guides/warlock/affliction.md unless noted):
- :51 "The raid build is a Shadow Bolt build that keeps its DoTs up: it stops at 30 points in Affliction and buys Improved Shadow Bolt and Ruin in Destruction." becomes "The raid build is a Wrack build too: it goes to 35 points in Affliction for Soul Siphon and Wrack, and spends its Destruction points on Bane and Aftermath for a faster, harder Immolate."
- :60 "Cheap and strong solo, weaker than Shadow Bolt in a raid." becomes "With three of your effects on the boss it out-damages Shadow Bolt for a quarter of the mana, so it is your filler everywhere."
- :62 (the whole raid-trade paragraph) becomes "The raid build drops Nightfall and Soul Harvest, the three points of Improved Drains, Improved Health Funnel and Improved Voidwalker, and takes both points of Improved Bane of Agony for 10% more Agony damage in the last minute. Malevolence gives your Shadow spells 5% crit, DoT ticks and Wrack included. In Destruction, Bane makes Immolate a 1.5 second cast and four points of Aftermath add 40% to its first hit. Its Demonology points are Unholy Power for your Succubus and Demonic Aegis for Demon Armor."
- :68 "If you hold the raid's Curse of the Elements, it goes on right after Corruption and lasts five minutes." becomes "If you hold the raid's Curse of the Elements, it goes on first, before Corruption, and lasts five minutes." In the same paragraph, "put up Corruption, then Bane of Agony while Amplify Curse is still on you, then Siphon Life and Immolate, and Shadow Bolt from then on" becomes "put up Corruption, then Bane of Doom, then Siphon Life and Immolate, and channel Wrack from then on". The Amplify and Curse of Weakness sentence becomes "Save Amplify Curse for the Bane of Agony that covers the last minute."
- :72 becomes "Bane of Doom beats Bane of Agony for Affliction too: one global for a big hit that crits for double with Pandemic. Put Doom up while a minute or more of the fight is left, and Agony for the last minute, with Amplify Curse on it."
- :74 "Wrack stays off your raid bars. It is a 6 second channel and Shadow Bolt does more in those 6 seconds. Keep Shadow Bolt as your filler." becomes "Wrack is your filler. Channel it whenever your DoTs are up; it also makes your Corruption and Agony hit 10% harder while it runs."
- demonology.md:81 "On the pull: Corruption, then Curse of the Elements if you hold it, then Bane of Doom" becomes "On the pull: Curse of the Elements if you hold it, then Corruption, then Bane of Doom".
- destruction.md:71 "Then Corruption, Curse of the Elements if it is yours, Bane of Doom" becomes "Then Curse of the Elements if it is yours, Corruption, Bane of Doom".
- APL notes, data/curated/apl/warlock-affliction.json:
  - Bane of Agony note (:181) becomes "Bane of Agony for the last minute, once Doom has run out. Never clip it, since its biggest ticks come last."
  - Shadow Bolt note (:282) becomes "Shadow Bolt only when Wrack cannot be cast: below level 40, or without the talent."
  - The three new entries have no note. Proposed:
    - Doom: "Bane of Doom while a minute or more of the fight is left. One global for a big hit after a minute, and Pandemic doubles its crits."
    - Amplify Curse: "Amplify Curse just before the Bane of Agony that covers the last minute. It makes that Agony 50% stronger and costs no global."
    - Wrack: "Wrack with your DoTs up. It is your filler: with Soul Siphon it out-damages Shadow Bolt for a quarter of the mana, and Corruption and Agony take 10% more while it runs."
  - engineering_notes "5" (Wrack "#2 damage source behind Shadow Bolt") is stale; Wrack is now the top source.

## Still open

- X1 Molten Skin / Destructive Reach (owner, +1.73% if dropped).
- Demonology Horde race (owner question 1; Orc +0.8%, under the rule).
- No level-60 or level-40-plus log; Warcraft Logs blocked. Round 14's open items (demon crit at 60, additive versus multiplied modifiers, Ignition's rate, Siphon Life tick crit, shadow wands) are unchanged.

## Wave 3: rank sweep on the level-gap penalty

Merged int/r15 into both branches (site 1e96d8df2, fork 9c511a2cd, which carry the priest lane's E 8d9ca6aae / 974614b8b). Under the new rule, a rank keeps its full spell power until the caster is 17 levels past its learn level, then loses 5% a level; Wrack opts out. It replaces Era's sub-20 cut. I rebuilt spec-breakdown on it.

What the penalty reaches: at 60, every fixed-rank cast in the three APLs is a top or near-top rank. The lowest are Corruption r6 (54), Shadowburn r6 and Life Tap r6 (56), all within 17 levels, so none loses anything. Life Tap, Curse of the Elements and Amplify Curse have no spell-power term. Below 60 the site rewrites every cast to the highest learned rank, which is always within 17 levels. The Succubus's Lash of Pain follows the pet's level. So the only rank questions are downranks, and I swept Shadow Bolt for every spec at bands 20 to 50 (max rank and the two below it; 3,000 iterations to screen, 10,000 at both seeds to confirm).

| # | row | result (horde, 10,000 iterations, seeds 1 / 50001) | decision |
|---|---|---|---|
| W1 | Destruction Shadow Bolt rank 4 at 30-39 (existing) | r4 90.9 / 90.9 against r5 90.5 / 90.4 (+0.4 / +0.5%) | kept |
| W2 | **new:** Destruction Shadow Bolt rank 3 at 20-29 | r3 55.3 / 55.3 against r4 54.6 / 54.6 (+1.3 / +1.3%); Alliance 55.6 / 55.6 against 54.9 / 54.8. Era's sub-20 cut used to take 30% off rank 3's spell power at 20. Mana-bound: 34 Life Taps a fight. | added |
| W3 | **new:** Demonology Shadow Bolt rank 3 at 20-29 | 50.0 / 50.0 against 49.9 / 49.8 (+0.2 / +0.4%, error ±0.02); Alliance 51.2 / 51.2 against 51.1 / 51.1 | added |
| W4 | Affliction Shadow Bolt at 20 / 30 | r3 56.2 against 56.4 at 20; lower ranks lose at 30. From 40 Wrack is the filler, so Shadow Bolt's rank does not move the number | no row |
| W5 | Demonology 30 to 50, Destruction 40 and 50 | every lower rank loses or ties. Destruction's filler from 40 is Incinerate, rank 1 at 40, the only rank | no row |

No race key: both factions gain on W2 and W3. Rows are hand-synced to sim/rankchoices/rank-choices.json. The rank-choice tests in sim/request, sim/rankchoices, talent-search and rotation-search pass.

Bare bands, horde, seed 1 (seed 50001 within 0.1 of it): before ($R/sb-before) -> wave 1 -> wave 3:

| band | Affliction | Demonology | Destruction |
|---|---|---|---|
| 20 | 54.6 -> 54.6 -> 56.4 (+3.3%, all from the engine change) | 48.4 -> 48.4 -> 50.0 (+3.3%: +3.1% engine, +0.2% row) | 53.0 -> 52.9 -> 55.3 (+4.5%: +3.2% engine, +1.3% row) |
| 30 | 90.0 -> 90.0 -> 90.0 | 82.1 -> 82.0 -> 82.0 | 91.0 -> 90.9 -> 90.9 |
| 40 | 139.6 -> 166.6 -> 166.6 | 147.7 -> 147.9 -> 147.9 | 149.1 -> 149.1 -> 149.1 |
| 50 | 207.0 -> 257.7 -> 257.7 | 222.2 -> 222.1 -> 222.1 | 268.6 -> 269.2 -> 269.2 |
| 60 | 307.4 -> 375.9 -> 375.9 | 347.1 -> 346.7 -> 346.7 | 404.9 -> 405.1 -> 405.1 |

Raid 60, all three specs and both factions and seeds: wave 1 = wave 3 exactly (Affliction 592.3 horde s1; Demonology 639.4; Destruction 684.1). The new penalty costs the warlocks nothing; ending Era's sub-20 cut lifts band 20 by about 3%.

Files to regenerate (wave 3): Destruction and Demonology band-20 BiS / weights and the addon's rank lines (raceSpellId rows) at the refresh; ladder goldens for both at level 20.

## Wave 4: guide-accuracy round 2 seeds

Merged int/r15 0d9241e0c (site) and 86fe4bf34 (fork) into both branches first. Read `design/reviews/guide-accuracy/round2/warlock.md` (land/guides-r2, now on origin/main 7a2165cb0).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **Card order: Agony versus Doom, and Curse of the Elements' place, on all three cards.** The guide's Rotation card and rail read `data/builds/<build>/addon-data.json` `rotations` (`web/src/lib/bis/rotation-view.ts:57`, `[spec].astro:169`), which `pipeline.addonrotation.build_rotations` builds from the curated APL in priority order. The committed addon-data predates wave 1: Affliction shows Corruption, then Curse of the Elements, then Bane of Agony, with no Doom, Amplify Curse or Wrack; Demonology and Destruction show Corruption before Curse of the Elements. I computed `build_rotations` in memory from the current curated APLs (nothing written). Level 60 gives: Affliction Life Tap, Curse of the Elements, Corruption, Bane of Doom, Amplify Curse, Bane of Agony, Siphon Life, Immolate, Wrack, Shadow Bolt; Demonology Soul Fire, Life Tap, Curse of the Elements, Corruption, Bane of Doom, Bane of Agony, Immolate, Shadow Bolt; Destruction Curse of the Elements, Corruption, Life Tap, Bane of Doom, Bane of Agony, Immolate, Conflagrate, Shadowburn, Incinerate, Shadow Bolt. This is exactly the APL priority. The cards are right after the addon-data regen and wrong until then. The Doom, Amplify Curse and Wrack lines carry no note, so their card clause is empty until the guide lane writes the notes (wave 1 "For the guide lane"). | rotation-view.ts, rotation-lists.ts, addonrotation.py; curated APLs | Aligned (regen pending) | - | - |
| G2 | **Wrack's line named no rank.** `data/tests/test_apl.py::test_every_spell_the_rotations_name_exists_with_that_rank` was red on int/r15: "warlock-affliction names spell 1316697 at rank 0; the build says rank 1". The line now names rank 1 in all three copies. test_apl.py passes (24/24), apl-check matches, and DPS is unchanged (667.2). | test_apl.py | CONFIRMED-BUG (my wave-1 line) | R15 9e270b86c | S 4dace2406, E 06b330deb |
| G3 | **Demonology's Soul Fire in the raid execute** (round 2 S1 / D4). On the current engine this is no longer the round 2 tie. Without the Soul Fire line, raid 60 measures horde 639.4 to 648.2 and 640.2 to 648.5 (+1.4 / +1.3%), alliance 635.3 to 643.7 and 635.3 to 643.9 (+1.3 / +1.4%). The reason: with Demonic Pact the Imp sacrifice (15% Shadow), Master Demonologist's Succubus 10% and Improved Shadow Bolt all raise Shadow, and Soul Fire is Fire. The leveling builds hold no Pact, and there the free Soul Fire still wins (bare 60 346.7 with it, 341.7 without, +1.5%). Bare 50 is +0.2% without it, a small cost I accept to keep the line. Warcraft Tavern's Forever Demonology guide agrees: "Skip Soul Fire in this build, even during Decimation". The line is now gated on `not auraIsKnown(425464)` (Demonic Pact). Test first: `TestDemonologySoulFireOnlyWithoutDemonicPact` failed on the raid build (8.33 casts) and passes now; bare 60 still casts it. After: raid horde 648.2 / 648.5, alliance 643.7 / 643.9; bare 50 and 60 unchanged. | runs; [Warcraft Tavern, Demonology (Zephan)](https://www.warcrafttavern.com/forever/guides/demonology-warlock/) (read by the guide lane 2026-11-04) | CONFIRMED-BUG | old (round 2 D4 measured a tie and held it); guides-handoff (round 2 S1) | S 9fa063565, E 1c935d440 |
| G4 | **Incinerate's class mask**: which talents modify it. Client 70755 (wago SpellEffect / SpellClassOptions, class set 5): Incinerate's family mask is word 1 0x40 (all three ranks). The talents whose masks hit it are: Bane (17788 aura 107 misc 10, -500 ms on Shadow Bolt, Immolate and Incinerate; Soul Fire -2 s on its own effect), Cataclysm (17778 misc 14, cost), Ruin (17959 misc 15, crit bonus), Agonizing Flames effect 1 (17927 aura 108 misc 0, +10% damage) and Destructive Reach (range). Agonizing Flames' crit effect covers Searing Pain only (mask 0x100), and its periodic effect covers Immolate and Rain of Fire only (0x24). The engine matches: Bane by spell code (`talents.go:866`); Cataclysm, Ruin and Agonizing Flames through `WarlockFlagDestruction`, which Incinerate carries (`incinerate.go:57`). Shadow Mastery, Malevolence, Improved Shadow Bolt and Decimation do not reach it, which is right: it is Fire, and outside their masks. Incinerate's "+25% on an Immolated target" is effect 1, a dummy (Effect 3, base points 25); the engine multiplies the hit by 1.25 when Immolate is up. Nothing is missing or extra. | c755/SpellEffect.csv, SpellClassOptions.csv (wago, build 1.60.1.70755); `incinerate.go:42-92` | Aligned | - | - |
| G5 | Round 2 S4, Wrack's 10% against "your other Shadow DoTs": the 70755 vulnerability effect (1316697 effect 1, aura 271) has class mask word 0 = 1026 = Corruption (0x2) plus Bane of Agony (0x400). Siphon Life (word 1 0x1) and Bane of Doom (word 0 0x80000000, word 1 0x2) are outside it. The engine models the mask. The tooltip's wording is broader than the client row, and the client row decides the mechanic. The guide's round 2 prose ("your other Shadow DoTs") is the tooltip's own wording, so it may stay. | SpellEffect rec 1349556 / effect 1 | Aligned | - | - |
| G6 | Destruction raid build (wave 1 X1): the guide lane's round 2 adopted `25500000033/0/0050355103101351` on main, on Warcraft Tavern's raid build. That is the variant I measured at +1.73% (684.1 to 695.9, 683.8 to 695.6). X1 is closed. | origin/main destruction.md:7 | resolved | - | - |
| G7 | **Live hotfix check (coordinator item 2).** I fetched the live 70755 push (`pipeline.hotfix_push.fetch_push`, in memory only; SpellEffect 1,618 entries, latest push 113557) and compared every pushed row on the spells this lane relies on: Shadow Bolt r3, r4, r5, r9, r10; Corruption; Agony; Doom; Amplify Curse; Siphon Life; Immolate; Wrack; Soul Fire; Curse of the Elements; Life Tap; Incinerate; Conflagrate; Shadowburn; every warlock talent spell; and the pet and buff spells. Only Wrack's effect 0 differs from the shipped table (coefficient 0.143 to 0.286, 0 to 2.5 a level, push 113527), and our layered data already carries it. SpellMisc push 113341 restates those spells with no change. No SpellPower, SpellCooldowns or SpellLevels row on them was pushed. Nothing differs from our data. | wago push, 2026-11-04 | Aligned | - | - |

Files to regenerate (wave 4): `data/builds/1.60.1.70755/addon-data.json` (the three warlock cards, G1), plus the wave 1 list. Ladder goldens for Demonology (G3). Measure at refresh: Demonology raid 60 picks and weights without Soul Fire (Fire power is worth less).

For the integrator: guides-r2 landed on origin/main from the pre-wave-1 tree, so merging main into int/r15 will conflict on:
- **affliction.md:7:** main has `human:2552001013500105/00052/0550004`, int/r15 has `gnome:25520020135001351/00052/005004`. Take `FS1:1.60.1.70755:warlock:human:25520020135001351/00052/005004:`.
- **The three curated APLs:** main edits notes on the old entry order. Keep int/r15's actions and order, and take main's note text.
- **Prose:** main's affliction.md:51, :62 and :74 still describe the Shadow Bolt raid build ("Wrack stays off your bars in this build"). The wave 1 "For the guide lane" replacements apply.

Pre-existing on int/r15, not this lane: `TestEveryWrittenRotationRunsInTheEngine` fails for hunter-marksmanship and paladin-holy (unresolved spells 1310786; 20930/1311606).

## Wave 6: guide-accuracy round 3 seeds

Merged int/r15 ef3195852 (site) and 6a4695cc0 (fork) first. Read `round3/warlock.md` (land/guides-r3). Client rows were checked against wago's 70755 tables plus the live hotfix push (fetched 2026-11-04, latest SpellEffect push 113557).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **Eureka! (warlock 1259821) against Incinerate and periodic ticks.** The client rows are explicit:<br>- effect 0, cost -10% (aura 108, misc 14): mask word 0 541165, word 1 8650880.<br>- effect 1, damage +10% (misc 0): mask 524773 / 8388736, which covers Shadow Bolt, Immolate's hit, Rain of Fire, Soul Fire, Shadowburn, Searing Pain, Death Coil, Conflagrate and Hellfire.<br>- effect 2, periodic damage +10% (misc 22): mask 16392 / 262144, which covers Drain Life, Drain Soul and Wrack.<br>Incinerate's family bit is word 1 0x40. It is in neither the damage mask nor the cost mask. The live push has no row for 1259821, so the shipped rows stand. `sim/core/eureka_auto_gen.go` is generated from these rows (`tools/eureka/gen.py`) and matches them exactly. The other side: the Wowhead Forever tooltip says "non-periodic damaging abilities", and Warcraft Tavern says "use it for Incinerates" (a guide line, not a measurement). The client table is the stronger source and beats the description (brief rule), so the engine is right as far as the data shows, and I made no shared change. A level 40+ Gnome Destruction log showing Eureka's 10% on an Incinerate would settle it (beta logs stop around level 30). Size, from a scratch build with Incinerate added (reverted, not committed): Gnome Destruction raid 692.7 to 694.0 (+0.2%). | wago SpellEffect / SpellClassOptions 70755 (records 1269085 to 1269087); live push; `eureka_auto_gen.go` | OPEN (engine follows the client mask; the description disagrees) | - | - |
| E2 | **Gnome versus Human, raid 60, Alliance, 10,000 iterations, seeds 1 / 50001.** Race picks within 1% stay.<br>- **Destruction**, adopted raid build `25500000033/0/0050355103101351`: Gnome 692.7 / 692.7, Human 696.5 / 696.0 (Human +0.55 / +0.48%). On the old build, Gnome 680.9 / 680.7, Human 684.3 / 683.9. Human stays the pick on main (guides r2), and it would hold even with Eureka on Incinerate (694.0).<br>- **Assassination**, guide raid build: Gnome 695.0 / 694.7, Human 707.3 / 706.9 (+1.8%). The published Alliance pick is Night Elf: 707.9 / 707.5, within 0.1% of Human. Night Elf stays; Gnome is not a contender. Nothing for the rogue lane.<br>- **Discipline**, effective HPS: Gnome 517.5 / 517.4, Human 518.0 / 518.0 (+0.1%). The published pick is Human (with Undead). Nothing for the priest lane. | spec-breakdown -race on this branch | Aligned (no race change) | - | - |
| E3 | **Destruction weapon enchant (round 3 R1).** Measured on Human, the adopted raid build and the published gear (a staff, so a stone fits). I used a scratch binary that sets the warlock imbue option; it was reverted, not committed, and the site never sets the option. Results: Brilliant Wizard Oil 696.5 / 696.0; Firestone 687.2 / 686.6 (-1.3 / -1.4%); Spellstone 674.4 / 673.9 (-3.2%); nothing 667.8 / 667.4. The order is oil, then Firestone, then Spellstone, as the guide's round 3 line says ("Brilliant Wizard Oil is the weapon pick; a Firestone ... is the fallback"). The engine's Firestone is 21 Fire power plus 2% spell crit, matching client 23483 (aura 13 misc 4 = 21; aura 57 = 2); the Spellstone is 21 Shadow power plus 2% cast speed (1237165). The push changes neither. | runs; SpellEffect 23480 to 23483, 1237162 to 1237165 | Aligned | guides-handoff | - |

Wave 6 changed no code. The Eureka question stays open on the client-versus-description split above. For the guide lane: destruction.md:113 (round 3) says Eureka! covers Incinerate. The client mask says it does not. Until a log settles it, a sentence that holds either way would be: "press it as Conflagrate and Shadowburn come off cooldown".

## Wave 9: Skyborne racials (shared sim/core)

Merged site int/r15 3c9cddc1d and fork int/r15 fbc5dbddf into both branches first. Wave 9 changed no code.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | **Wind Blessed (1259710) haste is already 2% in the engine.** Fork int/r15 `sim/core/racials.go:184-187` has `windBlessedHasteMultiplier = 1.02`, applied to melee, ranged and cast speed (`skyborneShared`, :129-138), and its note reads "2% ... aura 342 and aura 65 at 2". `TestWindBlessedHastesCastingMeleeAndRanged` (`racials_test.go:409`) asserts 1.02 on all three and passes. The client agrees: the shipped 70755 SpellEffect rows 1268898 (aura 342) and 1354094 (aura 65) both have base points 2; the live push re-sends both at 2 (push 113276, newest per record); data commit 43dc0f074's spelltips read "Increases your spellcasting, melee, and ranged Haste by 2%"; races.json says 2%. The "1%" in the brief is not in the code at any fork ref I checked (int/r15, d1d9f2775). | wago SpellEffect 70755 plus live push; fork int/r15; spelltips | Aligned (no bug) | - | - |
| K2 | **Skyborne base stats.** 70755 ChrRaces has no stat columns, and the client has no base-stat GameTable or stat-mod table for any race. Base stats are server-side, and the mined GameTables (combatratings, basemp, hppersta) carry none. `base_stats_provisional.go` keeps its clones (High Order from Human, Windshaper from Orc) and lists them as provisional, which is the honest state. Possible future source: a level-60 Skyborne character's naked stats from the Blizzard profile-forever import, or an in-game /dump. | wago ChrRaces header 70755 | OPEN (no client data; the provisional clone stays) | - | - |
| K3 | **Read Leyline is already modelled** (`racial_read_leyline.go`, fork int/r15). The client 1259705 casts Energized 1270842: aura 110 (mana regen percent) 50 and aura 88 (health regen) 50, for 15 s, on a 120 s cooldown with the global cooldown. The engine raises SpiritRegenMultiplier by 0.5 for 15 s on that cooldown, and `racial_read_leyline_test.go` covers it. It is APL-only: `ShouldActivate` returns false, and no curated APL casts it. I measured a test line (cast when mana is under 70%, at the top of the list) on Alliance Skyborne, raid 60. Restoration druid: 922.3 to 928.6 and 922.0 to 928.2 effective HPS (+0.7%). Fire mage: 655.7 to 652.5 and 656.0 to 652.5 (-0.5%); the global costs a caster more than the regen gives. No published spec is Skyborne and those APLs belong to the druid and mage lanes, so I added no line. | runs (scratch rotations) | Aligned engine; APL line for other lanes | - | - |

**Skyborne measurements.** Every published guide race is non-Skyborne: there is no `skyborne` in any build string or recommendedRaces on int/r15 or origin/main. Each row is published race then Skyborne (High Order for Alliance, Windshaper for Horde), 10,000 iterations, seeds 1 / 50001.

| spec | faction | raid 60 | bare 60 |
|---|---|---|---|
| Arms (guide raid build) | Alliance | 790.6 / 790.3 to 795.2 / 795.2 (+0.6%) | 292.4 / 292.4 to 294.5 / 294.7 (+0.7%) |
| Arms | Horde | 803.5 / 803.1 to 806.3 / 806.6 (+0.4%) | 290.9 / 290.7 to 293.7 / 293.4 (+1.0%) |
| Fury (guide raid build) | Alliance | 832.6 / 833.0 to 834.0 / 834.8 (+0.2%) | 310.0 / 310.6 to 310.9 / 310.9 (+0.2%) |
| Fury | Horde | 833.6 / 833.5 to 836.8 / 837.2 (+0.4%) | 314.3 / 314.9 to 316.2 / 316.3 (+0.5%) |
| Fire mage | Alliance | 666.4 / 667.1 to 655.7 / 656.0 (-1.6%) | 342.5 / 343.0 to 332.7 / 333.5 (-2.8%) |
| Fire mage | Horde | 673.0 / 673.5 to 668.0 / 668.6 (-0.7%) | 335.1 / 335.7 to 333.3 / 334.0 (-0.5%) |
| Restoration druid (eff. HPS) | Alliance | 925.8 / 925.8 to 922.3 / 922.0 (-0.4%) | 463.0 / 463.0 to 461.2 / 461.6 (-0.4%) |
| Restoration druid | Horde | 923.3 / 923.8 to 924.9 / 925.1 (+0.2%) | 459.1 / 458.9 to 460.5 / 460.7 (+0.3%) |

Skyborne Arms leads the published race by +0.4 to +1.0%. Horde bare 60 is the only cell at 1.0%, and raid is under 1%, so under the race rule the published Arms and Fury picks stay. These numbers rest on the provisional cloned base stats (K2).

For other lanes:
- druid / healers: a Read Leyline line for Skyborne Restoration druids measures +0.7% (K3). Worth it if a Skyborne druid is ever published.
- warrior: Skyborne is within 1% of the published race for Arms and Fury on the current (cloned) base stats.
