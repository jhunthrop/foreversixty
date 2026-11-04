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
