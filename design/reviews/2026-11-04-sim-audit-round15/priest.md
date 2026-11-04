# Sim-accuracy round 15: priest

Lane priest. Site wt/r15-priest and fork lane/r15-priest, both from int/r15 (site fd320b60e, fork d1d9f2775). Client 1.60.1.70755. Specs: Discipline, Holy, Shadow.

Sims ran through the lock wrapper. Exploration used 3,000 iterations; reported numbers use 10,000 at seeds 1 and 50001. Healers are measured in effective HPS on the heal profile (`-heal`).

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| SH1 | **The leveling build took Spirit Tap 3/5 and Blackout 4/5.** At 3/5, Spirit Tap procs on 60% of kills; at 5/5 it procs on every kill (talents/priest.json, node 105833 rank text). Every Forever leveling guide takes 5/5, and none takes Blackout 4/5.<br>- **Change:** Spirit Tap 5/5 and Blackout 2/5. The Shadow tree goes from `54342001201300051` to `52542001201300051`.<br>- **Sim cost:** none. Neither talent changes a number on the band fight: it has no kills, and the engine models no stun (`sim/priest/talents.go:113-118, 288`). Every band keeps the same modeled talents.<br>- **Decided by the web-research rule:** the sim cannot model this choice. Icy Veins Forever calls Spirit Tap "likely the most important leveling talent". Mobalytics (Twicks) and Kargoz both put 5/5 Spirit Tap first. | `TestShadowLevelingBuildTakesSpiritTapFiveBlackoutTwo` failed (3, 4), then passed | CONFIRMED-BUG | old, guides-handoff (PRIEST-1 / B1) | S 925a3b1e0 |
| SH2 | **The leveling order put Spirit Tap at 19-21, behind Shadow Focus 5 and Blackout 4.** A new ladder-first row now takes Spirit Tap 5 first: Spirit Tap at 10-14, Shadow Focus 15-19, Blackout 20-21, Improved Mind Blast 22-25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, then Vampiric Embrace, Shadow Weaving, Darkness and Shadowform at 40. The sim's bands 20, 30 and 40 hold the same modeled talents as before (Shadow Focus 5 at 20). Sources as SH1. | `TestShadowLevelsSpiritTapFirst` failed, then passed. `TestCuratedLadderFirstIsSound` passes with the new row. | CONFIRMED-BUG | old, guides-handoff (PRIEST-2 / B2) | S 925a3b1e0 |
| SH3 | **Mind Flay at level 20** (the seed's question and SIM-P3) is **not adopted.** The sim says it loses. Band 20 bare, 3,000 iterations: Shadow Focus 5, Blackout 4, Spirit Tap 2 gives 33.7. Every build with Mind Flay at 20 gives 32.2 to 32.8 (-2.7%).<br>- **Why it loses:** the 180 s band fight runs out of mana at about 85 s. A rank-1 Mind Flay without Improved Mind Flay pushes out Shadow Word: Pain and Mind Blast.<br>- **With a mana-pace gate** on Mind Flay, band 20 ties (33.8). But the same gate loses at every other band: bare 30 -5.8%, 40 -6.8%, 50 -3.6%, bare 60 -5.3%, raid 60 -1.8%. So the unconditional Mind Flay line stays.<br>- **Icy Veins Forever agrees:** Mind Flay and Devouring Plague are "steps" to drop when short on mana, which "you generally won't need to cast ... early on (up until level 30)". | sims (rotation files in the lane scratch) | Aligned | - | - |
| SH4 | **The raid build spent 5 points that do nothing on a boss:** Blackout 2 (bosses are immune to the stun; the engine names it and does not model it) and Spirit Tap 3 (needs a kill, or a dying add carrying Vampiric Embrace). That is a confirmed bug under the owner's rule for points that provably do nothing.<br>- **The points now go to:** Devouring Contagion 2, Silence 1, Shadow Reach 2/2 and Dark Inspiration 1. The tiers below Early Demise and Darkness keep 25 points, and Shadowform still has 32 before it. The Shadow tree goes from `52352101201300251` to `50052201211312251`.<br>- **Raid 60, guide build, 10,000 iterations** (alliance s1 / s50001, horde s1 / s50001): 687.5→686.9, 688.1→687.7, 695.2→694.8, 694.9→694.8. That is -0.01 to -0.09%, within error.<br>- **Without Devouring Contagion** (Silence 1, Reach 2, Dark Inspiration 3): 687.4 / 688.2, an exact tie.<br>- **Why Devouring Contagion anyway:** both strong-player raid builds take it, with Silence 1, Reach 2/2 and Dark Inspiration, and with no Blackout or Spirit Tap (Mobalytics Twicks 14/3/34; Warcraft Tavern 14/3/34). Twicks keeps Devouring Contagion for the add-fight spread. | `TestShadowRaidBuildSpendsNothingOnBossDeadTalents` failed, then passed | CONFIRMED-BUG | old, guides-handoff (PRIEST-10 / S3 / SIM-P4) | S 925a3b1e0 |
| SH5 | **The opener started its first cast at the pull.** Both Forever Shadow guides pre-cast Mind Blast so it lands on the pull, after Fade (Warcraft Tavern; Mobalytics, 2026-11-02).<br>- **Change:** a prepull Mind Blast rank 9 at -1.5 s, after Shadowform at -20 s. Prepull lines are rank-capped below 60 (`sim/request/rotation_ranks.go:89`).<br>- **Raid 60:** +0.77 to +0.85%.<br>- **Bare bands** (3,000 iterations, seed 1): 20 +0.3%, 30 0, 40 +0.2%, 50 +0.5%, 60 +0.3%. At 10,000 iterations, bare 60 is 349.9→350.9 and 350.1→350.8. No band loses. | `TestShadowOpensWithAPrePullMindBlast` failed, then passed. `make apl-check` passes for all three copies. | CONFIRMED-BUG | old (round 14 casters S2 read the spells, not the opener against the guides) | S 5b46c799d, E 3147ce097 |
| SH6 | **Shadow Word: Pain end-of-fight cutoff** (S2 / SIM-P2): `remainingTime >= 10` against 6, 3 and 0. Raid 60 alliance s1, 3,000 iterations: 684.2 against 684.4 / 684.1 / 683.7. All tie.<br>- **Vampiric Embrace's** cutoff of 10 against 5 and 0: 683.6 / 683.1. Also a tie.<br>- **Web check:** the guides only say "maintain uptime". There is no consensus to move to, so the list stays as it is. The guide lane's "closing seconds" wording is right. | sims | Aligned | - | - |
| SH7 | **Wand Specialization first** (Icy Veins and Kargoz take it at 10-11) would be big at the low bands: band 20 +10.1% (33.7 to 37.1) and band 30 +5.7%. But a ladder only adds points. With 2 points in Discipline, a level-40 priest has 29 in Shadow and no Shadowform until 42. Band 40 falls from 131.4 to 84.2 (-36%). That breaks the "lose at no band" rule, so Wand Specialization stays at 41-42. | sims | OPEN (owner) | - | - |
| SH8 | **Fresh re-audit, aligned.**<br>- **Talent rates against the 70755 text:** Dark Inspiration 3%/rank, Mind Blast and full Mind Flay triggers (`talents.go:580-612`); Devouring Contagion -25%/rank; Early Demise 15/30 below 20%; Darkness 2%/rank; Shadow Focus 1%/rank; Improved Mind Flay compounding; Shadowform +10% damage, -50% cost, 2.0× crits; Twilight Focus 23/47/70.<br>- **Rank rows:** every Shadow row is trainable at its band (spellranks: Shadow Word: Pain rank 3 at 18; Mind Blast rank 2 at 16, 3 at 22, 4 at 28; Devouring Plague 1 at 20 and 2 at 28; Shadow Word: Death 2 at 40).<br>- **Unchanged since round 14 S2:** spells, masks, BiS gear (no set; Frozen Heart's use) and the 70755 data (build.md X7: no priest spellconst change). | engine and client rows | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Rank rows, trainable at every band:** Heal rank 2 from 22, rank 3 from 28, rank 4 from 34, and rank 4 at bare 60; Prayer of Healing rank 1 from 30 (spellranks.json).<br>- **Renew rank 1 and Heal rank 1 at raid 60** rest on the pending low-rank tooltip (owner, O1). No change.<br>- **No row depends only on something else:** 70755 moved no priest heal row (build.md X7), so the round-14 measurements stand. | spellranks.json | Aligned / pending owner | - | - |
| H2 | **Prayer of Mending or Renew on the tank before the pull:** raid 60 horde s1, 3,000 iterations: 790.9 against 790.4 (Prayer of Mending) and 791.5 (Renew). Both tie, so neither is added. | sims | Aligned | - | - |
| H3 | **Build:** every point does something in the sim or in a raid. Holy Binding is fed by the tank's Flash Heal. Spirit of Redemption is Prayer of Mending's prerequisite. Litany of Light 5%/rank; Improved Healing -5%/rank; Divine Fury -0.1 s/rank; Spiritual Healing 3/7/10%. **Binding Heal:** no APL casts it (removed at the 70755 activation). **Gear:** the Vestments of the Devout set (182, six pieces on the raid entry) goes through the generic client-set path (`core/client_sets_gen.go:24`, the 5-piece mana proc 450576). | client rows, engine | Aligned | - | - |

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Rank rows:** Power Word: Shield rank 6 at 40-49 (learned at 36); Prayer of Healing rank 1 at 40-59; Lesser Heal rank 2 at 4-9 and rank 3 at 10-59. All are trainable.<br>- **Lesser Heal rank 1 at 60 and Renew rank 1** rest on the pending tooltip, and so would moving the 10-59 rows to rank 1.<br>- **The 4-9 row** is harmless: no band sits below 20.<br>- **Power Word: Shield rank 4 at 60** is held for the unmodeled rage cost (round 14 F4). | spellranks.json | Aligned / pending owner | - | - |
| D2 | **Renew on the tank before the pull:** raid 60 +0.38 to +0.61% (584.4→586.6 / 584.1→586.9 alliance, 569.8→573.3 / 569.8→573.0 horde). Bare 40 +1.0%, bare 60 +0.9%, but **bare 20 -0.3%** (34.7→34.6, 3,000 iterations), and Holy gains nothing from the same line. **Not adopted:** a band loses. The open question is whether a pull-time Renew is what the band-20 trash pulls look like. | sims | OPEN | - | - |
| D3 | **Build:** Soul Warding, Renewed Hope 2%/rank and 1 s/rank, Divine Aegis 5%/rank, Improved Power Word: Shield 7/14/20, Mental Agility 3/7/10 and Meditation 17/33/50 match the 70755 text. Twilight Focus (pushback from the raid pulses) and Inspiration (tank armor) do something. Same gear as Holy, with five Devout pieces. | engine `healing_talents.go:19-47`, `talents.go:33` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Healer haste scale factors are negative** (S1: Discipline -8.0, Holy -5.4 per 1%). This is the mana-pacing and 0.5 s grid model: core's C4/F5 from round 14, and MR19. | round 14 healers F5 | OPEN (core) | - | - |
| X2 | **Shadow's bare-60 weights rank mp5 and Spirit above spell power** (S4). That is real mana starvation on bare gear (the fight is dry at 13.5 s). It is a BiS-refresh matter, not a sim bug. | sims | Aligned | - | - |

## Before/after

Raid 60, 10,000 iterations. Before is $R/before.tsv; after is this branch's spec-breakdown on the published entries.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | alliance s1 absolute |
|---|---|---|---|---|---|
| priest-shadow | +0.85% | +0.77% | +0.85% | +0.80% | 684.2 → 690.0 |
| priest-holy | 0 | 0 | 0 | 0 | 807.3 |
| priest-discipline | 0 | 0 | 0 | 0 | 584.4 |

- **The Shadow move is SH5 alone.** The published raid entry still carries the BiS file's own talents.
- **On the guide raid build,** SH4 is -0.01 to -0.09% (SH4 row).
- **Shadow bare 60** (alliance): +0.29% / +0.20%.
- **The band builds** carry the same modeled talents, so SH1 and SH2 move nothing.

## Files to regenerate

- `data/builds/1.60.1.70755/addon-builds.json`: Shadow leveling and raid orders. `TestCommittedAddonBuildsFollowTheGuides` fails until it is regenerated. Then addon-data and Data.lua.
- `sim/request/testdata/ladder/priest-shadow.golden.md` (TestRotationLadder): band talent strings (Spirit Tap and Blackout swap) and the prepull Mind Blast.
- `sim/leveling/ladder-first.json` was copied by hand from data/curated (= `make ladder-first-sync`).
- dps-per-point priest-shadow; the coach starter for priest-shadow, if it reads the opener.

## Measure at refresh

- **priest-shadow, every band:** the BiS files carry the published talents. That means the old raid string `52352101201300251` (with Blackout and Spirit Tap) and bands with Blackout 4 / Spirit Tap 2-3, plus the prepull Mind Blast in the weights.

## For other lanes

- **Shared files, priest row only:** `data/curated/ladder-first.json` gains a priest-shadow row, and `sim/leveling/ladder_first_curated_test.go` gains its spec entry. Nothing else moves.
- **No commit moves DPS outside priest.**

## For the guide lane

Lines in `web/src/content/guides/priest/shadow.md` on land/guides-r1:

- **:34**, "from level 21 Spirit Tap cuts your drinking after most kills" → "from level 10 Spirit Tap cuts your drinking after every kill".
- **:61**, "**Spirit Tap** (3 points): a 60% chance after a kill to double your Spirit for 15 seconds, ..." → "**Spirit Tap** (5 points leveling, none in the raid build): after every kill your Spirit doubles for 15 seconds, with half your regeneration running while you cast. It is your drink between pulls."
- **:62**, "**Blackout** (4 points leveling, 2 in the raid build): an 8% chance (4% at 2 points) ..." → "**Blackout** (2 points leveling, none in the raid build): a 4% chance for your Shadow spells to stun for 3 seconds, which stops a mob that reaches you. Bosses cannot be stunned, so the raid build drops it."
- **:69**, "The raid build trades them, and two points of Blackout, for Early Demise, Shadow Reach and the last point of Improved Mind Blast." → "The raid build trades them, Blackout and Spirit Tap for Early Demise, Devouring Contagion, Silence, both points of Shadow Reach, Dark Inspiration and the last point of Improved Mind Blast."
- **:75**, the opener: "Shadow Word: Pain, Vampiric Embrace, ..." → "Mind Blast cast before the pull so it lands as the boss is engaged, then Shadow Word: Pain, Vampiric Embrace, Devouring Plague, Mind Blast again (with Inner Focus, so it is free and more likely to crit), Shadow Word: Death, and Mind Flay."
- **:123**, "Spirit Tap procs on most kills" → "Spirit Tap procs on every kill".
- **:125** → "The leveling build goes Spirit Tap from 10 to 14, Shadow Focus to 19, Blackout at 20 and 21, Improved Mind Blast to 25 and Improved Shadow Word: Pain at 26 and 27, then Mind Flay at 28. Until 28 the wand is your filler. Improved Mind Flay, Vampiric Embrace, Shadow Weaving and Darkness follow, and Shadowform is your point at 40."
- **APL note:** the priest-shadow prepull Mind Blast line has no note. Suggested: "Cast before the pull so it lands as the boss is engaged."

## Still open

- **SH7 (owner):** Wand Specialization at 10-11 (+10% at band 20, +5.7% at 30) against Shadowform at 40 (-36% at band 40). The ladder rule decides it for now.
- **D2:** Discipline's prepull Renew loses only at band 20.
- **Pending owner, not reopened:** the low-rank heal tooltips (Renew 1, Heal 1, Lesser Heal 1), and Inner Focus on Penance (round 14 P5).
- **Warcraft Logs is blocked.** No log check was made.

## Sources (web research, owner rule)

- Icy Veins, "WoW Forever Shadow Priest Leveling Guide", Anshlun, updated 2026-11-02: https://www.icy-veins.com/wow-forever/shadow-priest-leveling
- Mobalytics, "Shadow Priest Guide" (Forever), Twicks, updated 2026-11-02: https://mobalytics.gg/wow-forever/classes/shadow-priest-guide
- Warcraft Tavern, "Shadow Priest Guide - WoW Forever", Cyradys (no date shown): https://www.warcrafttavern.com/forever/guides/shadow-priest/
- Mobalytics, Kargoz's 1-60 Shadow leveling guide (search summary only; the page was not fetched): https://mobalytics.gg/wow-forever/profile/kargoz/classes/1-60-shadow-priest-leveling-guide-wands-talents-rotation-tips-and-tricks
- Method, "WoW Forever Priest Leveling Guide", updated 2026-11-02 (respec to Shadow at 40): https://www.method.gg/wow-forever/wow-forever-priest-leveling-guide-and-talents

## Wave 2: the low-rank spell power penalty (every healer)

The coordinator handed this lane the healing-bonus curve for low ranks, plus every healer rank row. It comes from guide-accuracy round 1: ledger line 57 on origin/land/guides-r1, and the round1 reports for priest (O1), shaman (Q3), druid (Q4) and paladin.

### The rule

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Forever's low-rank penalty is a level gap, not Era's level-20 cut.** A rank keeps its whole spell power coefficient until you are 17 levels past the level you learned it. After that it loses 5% a level, and it keeps none from 37 levels past: share = clamp(1 - (L - T - 17) × 0.05, 0, 1).<br>- **What the engine did:** heals took no penalty at all, and damage spells took Era's sub-20 rule (`core/spell.go` levelPenalty, 3.75% a level below 20).<br>- **The direction is official.** Blizzard's 1 October beta note and the character sheet say ranks "vastly below your current level receive less benefit from your Spell Damage and Spell Healing". Sources: https://us.forums.blizzard.com/en/wow/t/forever-beta-development-notes/2360696 (Kaivax), with Mein-MMO quoting it on 2026-10-12 (https://mein-mmo.de/wow-forever-kein-downranking/).<br>- **The numbers come from one measurement:** ForeverChanges (Vektor), heals, beta, levels 18 to 30, updated 2026-10-16: https://foreverchanges.pro/downranking. EpicCarry (2026-10-10, https://epiccarry.com/blogs/?p=41909) restates the same formula; it is not a second measurement.<br>- **The client tables carry no curve.** SpellLevels (70755, wago) gives every heal rank MaxLevel = learn level + 5, so a MaxLevel + 12 key would fit the same heals. The rule is keyed on learn level, as stated.<br>- **Round 11's beta logs agree.** The low ranks they caught (Rejuvenation 3 and 5 at levels 27 to 35, and others) all sit within 17 levels of their caster, so they fit this rule as well as "no penalty".<br>- **Applied to damage spells too.** The note names both stats. That is an assumption above level 30, flagged as such. | `TestLevelPenaltyIsForeversLevelGap`, `TestCastRanksKeepTheLevelGapShare` failed (build), then passed | CONFIRMED (direction official; curve single-source) | old (round 11 C1 set "no penalty on heals"); guides-handoff | E 8d9ca6aae |
| W2 | **Wrack is exempt.** It is a single-rank spell that grows to MaxLevel 60, so it is never below the caster. It opts out through `SpellConfig.ScalesToCasterLevel`. Without that it would have lost 15% of its coefficient at 60. | `TestASpellThatScalesToTheCasterKeepsItsCoefficient` | CONFIRMED (side effect caught) | - | E 8d9ca6aae |
| W3 | **Holy Shock's heal half escaped the penalty.** It is a triggered spell with no cost or cast, but it is the cast rank's own heal. So Holy Shock rank 1 (learned at 30) healed with its whole coefficient at 60, against 35% under the rule. Marked with `SpellConfig.TriggeredRank`, and conformance mirrors it. | `TestATriggeredRankTakesItsCastRanksShare` failed (build), then passed | CONFIRMED-BUG | old | E 974614b8b |
| W4 | **Top ranks penalised at 60 across every preset** (conformance goldens): only Heal rank 4 (the top trainer Heal, 55%) and Wrack (now exempt). No other spell's top rank is hit. Checked separately: the single-rank spells with a constant RequiredLevel below 43 either carry no spell power coefficient (Swiftmend, Vampiric Embrace, Stormstrike, Judgement, the summons) or are exempt. | goldens | Aligned | - | - |

### Healer lists and rows

Each change below was measured on the engine change. Sweeps ran at 3,000 iterations, horde, seed 1. Picks were confirmed at 10,000 iterations, seeds 1 and 50001. Bare bands 20 to 60 were checked for every spec; no band loses against the engine change alone.

| # | finding | evidence | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W5 | **Renew (Holy, Discipline):** rank 1 (learned at 8) keeps nothing at 60. Rank 7 (10927, learned at 44) wins over ranks 8 and 9 at raid 60: Holy 630.5 to 655.3 alone, Discipline 459.3 to 499.0. A row keeps rank 1 from 14 to 39, where it still wins (Holy bare 30: 65.9 against 63.7). Rank 7 wins from 40 on (Holy +6% at 40 and +11% at 50; Discipline +1.6% and +6%). | sims | CONFIRMED-BUG | R14 2f9d5ad6d (P2 set rank 1 at 60 on the "full coefficient" premise) | S 4faba89a9, E e439f8fb0 |
| W6 | **Holy's filler** was Heal rank 1 (learned at 16, keeps nothing).<br>- **Raid 60:** Greater Heal rank 4 (10965) ahead of each Heal line is the best: 746.6 against 743.2 / 743.5 for rank 1 at 10,000 iterations (+0.45%, both seeds), and bare 60 +0.7%. Ranks 1 to 4 are close, so no downrank is kept.<br>- **Heal lines:** they stay behind the Greater Heal lines at rank 4, the highest learned before 40. That makes the Heal rows (ranks 2 to 4, bare-60 rank 4) the default, so they are gone.<br>- **Prayer of Healing's rank-1 rows** at 50 and bare 60 tie when removed, so they stay. | sims | CONFIRMED-BUG | R14 2f9d5ad6d (P3) | S 4faba89a9, E e439f8fb0 |
| W7 | **Discipline's party filler** was Lesser Heal rank 1 (keeps nothing).<br>- **Raid 60:** Greater Heal rank 1 (2060) ahead of each Lesser Heal line: 500.4 / 500.7 against 496.5 / 496.7 with Lesser Heal rank 3 alone (+0.8%, 10,000 iterations, both seeds). Rank 2 gives 498.0 and rank 4 497.9.<br>- **Lesser Heal lines:** they fall back at rank 3 (rows gone).<br>- **Prayer of Healing:** dropping its rank-1 row at 50 to 59 gains 1.7% (bare 50: 134.0 to 136.3).<br>- **Not taken:** bare 60 would gain 4.4% with Lesser Heal rank 1 as the filler (the fight runs dry). A row cannot switch the spell, so bare 60 stays level with the engine change alone (208.8 against 208.7).<br>- **Power Word: Shield rank 4** (5% of spell power at 60) stays held for the rage cost, as in round 14 F4. Rank 10 is +0.8% at raid 60. | sims | CONFIRMED-BUG | R14 cf5fe63bb (F1) | S 4faba89a9, E e439f8fb0 |
| W8 | **The Restoration druid's party Rejuvenation** was rank 1.<br>- **Raid 60:** rank 7 (8910) beats rank 8: 923.3 / 923.8 against 919.1 / 918.8 at 10,000 iterations.<br>- **Bare fights:** rank 1 still wins on every bare fight, bare 60 included (435.7 against 410.2; the fight runs dry), so a row keeps rank 1 from 10 to 59 and at bare 60.<br>- **Tank line:** it drops its rank-1 row at 50 to 59 (+1.2%) and its rank-4 row at bare 60 (+5.3%). | sims | CONFIRMED-BUG | old (round 13 set rank 1 on the full-coefficient premise) | S 4faba89a9, E e439f8fb0 |
| W9 | **The Restoration shaman's tank Healing Wave** was rank 2 (learned at 6). Rank 9 (10396) wins: raid 60 658.1 to 713.2 (+8.4%), bare 40 +8.6%, bare 50 +18.6%, bare 60 +10.9%, with ties at 20 and 30. The rows for 12 to 39 were just the highest rank learned, so they are gone. The shaman lane had left these rows and the rank-2 line alone; its rows (Elemental 15207/15208, Flame Shock, Riptide, the Nature's Swiftness wave) are untouched. | sims | CONFIRMED-BUG | R14 cf5fe63bb (F2) | S 4faba89a9, E e439f8fb0 |
| W10 | **Holy paladin.**<br>- **Filler Flash of Light:** it was rank 1 (learned at 20). Rank 5 (19942) wins: 639.3 / 640.2 against rank 4's 635.8 / 636.2 at 10,000 iterations. Rows keep what the bands cast before: rank 1 from 26 to 39 (bare 30: 89.5 against 73.5), rank 3 from 42 to 49, and rank 1 from 50 to 59 (156.9 against 148.4).<br>- **Holy Light:** it drops its rank-1 row at 30 to 39 (+5.3%) and its rank-4 row at bare 60 (+1.8%). It keeps rank 1 at 20 to 29 and 40 to 49, where removing the rows loses 13 to 14%.<br>- **Flash of Light rows authored at rank 4 (19941):** both stay (removing them loses 3.6% at 30 and 2.3% at 40).<br>- **Holy Shock's rank:** ranks 1 to 4 tie at raid 60, so rank 1 stays. | sims | CONFIRMED-BUG | old (round 9 rows) | S 4faba89a9, E e439f8fb0 |
| W11 | **Shadow's rows.** Mind Blast's race rows at 50 to 59 (rank 2 or 3, keeping 15% and 45% of spell power) lose 3.3% against the highest rank learned, on gnome and undead alike, so they are gone. The others stay because removing them loses:<br>- Devouring Plague rank 1 at 40 to 49: -3%;<br>- Devouring Plague at 50 to 59: -4.5%;<br>- Shadow Word: Death rank 2 at bare 60: -1.3% gnome, -0.3% undead.<br>Other races' keyed rows were not re-measured. | sims | CONFIRMED-BUG | old (round 13 rows, set without the penalty) | S 4faba89a9 |
| W12 | **Mind Flay at 20, re-measured on the new engine** with the guide's candidate (Spirit Tap 5, Improved Shadow Word: Pain 2, Improved Mind Blast 3, Mind Flay; Shadow `00532001`) against this branch's band 20 (Spirit Tap 5, Shadow Focus 5, Blackout 1). Band 20 bare, 10,000 iterations: alliance 32.4 / 32.5 against 34.6 / 34.6 (-6.3%), horde 35.3 / 35.4 against 37.3 / 37.3 (-5.2%). Mind Flay rank 1 is fully powered at 20, but the band-20 fight runs dry. It still loses at a band, so no ladder-first row was added. The guides-r1 ladder-order.json is not on origin/land/guides-r1 at that path; nothing to add. | sims | Aligned (not adopted) | - | - |

### Wave 2 before/after (raid 60, 10,000 iterations)

"Before" is int/r15 (site 3506600b8 + fork 05cc3bc2e). "Engine only" adds E 8d9ca6aae and 974614b8b with the old lists. "Final" is this branch.

| spec | faction | s1 before → engine only → final | s50001 final vs before |
|---|---|---|---|
| priest-holy | alliance | 807.3 → 645.5 (-20.0%) → 760.8 (-5.8%) | -5.8% |
| priest-holy | horde | 790.4 → 630.5 (-20.2%) → 746.6 (-5.5%) | -5.5% |
| priest-discipline | alliance | 584.4 → 480.9 (-17.7%) → 520.5 (-10.9%) | -10.9% |
| priest-discipline | horde | 569.8 → 459.3 (-19.4%) → 500.4 (-12.2%) | -12.1% |
| druid-restoration | alliance | 1006.4 → 883.9 (-12.2%) → 925.8 (-8.0%) | -8.0% |
| druid-restoration | horde | 1005.3 → 883.6 (-12.1%) → 923.3 (-8.2%) | -8.1% |
| shaman-restoration | alliance | 733.0 → 659.5 (-10.0%) → 715.5 (-2.4%) | -2.4% |
| shaman-restoration | horde | 731.8 → 658.0 (-10.1%) → 713.3 (-2.5%) | -2.6% |
| paladin-holy | alliance | 676.5 → 578.6 (-14.5%) → 656.3 (-3.0%) | -3.0% |
| paladin-holy | horde | 662.0 → 561.9 (-15.1%) → 639.3 (-3.4%) | -3.4% |

- **Raid 60 DPS** is unchanged to the decimal for Shadow, Elemental, Balance, the three mages, the three warlocks, Retribution, Protection paladin and Enhancement.
- **Shadow bare bands** (3,000 iterations, seed 1, alliance / horde):
  - 20: +2.4% / +2.2%;
  - 30: +0.9% / +0.8%;
  - 40: -0.4% / -0.5%;
  - 50: -4.2% / -2.8%;
  - 60: -0.2% / -0.2%.
  - The low bands gain because Era's sub-20 cut is gone. The 50 loss is the penalty on Devouring Plague rank 1, whose row still beats the top rank there.
- **Elemental bare bands** (10,000 iterations, seeds 1 and 50001; horde / alliance), after dropping the rank-4 Lightning Bolt rows at 50 and bare 60:
  - 40: -6.7%;
  - 50: -22.5% / -21.7%;
  - 60: -15.8% / -16.4%.
  - Its rank-4 row at 40 to 49 still beats removal by 6.5%. Keeping the 50 and bare-60 rows would have been -26% and -34%.
- **Fire bare bands:** 40 -1.4%; 50 -3.1% (Scorch rank 1 row at 50 dropped: it was -14%); bare 60 0.
- Balance bare 30 is +1.5%, from the sub-20 cut going. Affliction bare bands are unchanged.

### Wave 2 files to regenerate
- **Fork goldens regenerated** (explained by E 8d9ca6aae / 974614b8b): `sim/core/testdata/conformance/{druid,mage,paladin,priest,shaman,warlock}.golden.md`, `SUMMARY.md`, and `sim/shaman/elemental/TestElemental.results` (its legacy default APL casts Lightning Bolt rank 4 at 60).
- **Site, not regenerated:**
  - TestRotationLadder goldens for the five healers, Shadow, Elemental, Fire and every caster whose ladder downranks;
  - dps-per-point;
  - addon-data and Data.lua (rank rows and rotations);
  - the coach starters;
  - `make apl-sync` (copies were synced by hand; `make apl-check` passes).

### Wave 2 for other lanes (loud)
- **Shared core change:** E 8d9ca6aae changes every caster and healer below 60 and every downranked cast at 60. Retune is needed for:
  - **shaman (Elemental):** at bare 50 and 60, Lightning Bolt rank 4 is no longer a cheap win. A full rank sweep of both Lightning Bolt lines and Chain Lightning at 40 to 60 is owed. This lane only removed the two rows that clearly lose.
  - **mage (Fire Scorch row 40 to 49, Frost, Arcane bare bands):** not swept.
  - **warlock:** single-rank or scaling spells besides Wrack should be checked for the `ScalesToCasterLevel` exemption.
- **Pre-existing failures, not from this lane:** `api/internal/coach` `TestWebFixtures_NoDrift` and `TestWebHistory_PagesAreTheRoutesShape` fail on int/r15 without these changes too.
- **APL notes now wrong** (notes are not this lane's to edit): every healer list's top-level note still says "the engine states the client's own numbers for every rank with no downrank penalty" and names the old ranks (Holy and Discipline Renew rank 1; Discipline Lesser Heal rank 1/3; druid party Rejuvenation rank 1; shaman Healing Wave rank 2; paladin Flash of Light rank 1).

### Wave 2 for the guide lane
- **holy.md:** the raid filler is Greater Heal at its top rank, not a downranked one. Replace "downranked Greater Heal" with "Greater Heal; the low ranks save almost no mana now that they lose their healing bonus".
- **discipline.md:** the party filler is Greater Heal rank 1 and Renew rank 7. Both already match the round-1 prose.

### Wave 2 still open
- Above level 30 the curve is assumed, not measured; one level-60 tooltip check would settle it. The same holds for damage spells.
- Discipline bare 60 would gain 4.4% with a Lesser Heal rank 1 filler; it needs a list split by preset.
