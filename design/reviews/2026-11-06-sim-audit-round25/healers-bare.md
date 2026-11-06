# Sim-accuracy round 25: healer bare-band drops (integrator)

Question (coordinator, from the final BiS pass): healer bare-band set numbers fell sharply against the published picks. Restoration shaman bare 40 to 60 is down 26 to 33%, Restoration druid bare 60 down 23% and Discipline down 17%, while raid 60 held or rose. Is that a deliberate sim fix or a regression in the bare preset for healers?

**Verdict: deliberate, no regression.** All three come from fork **8d9ca6aae** "fix(core): Forever's level-gap spell power penalty replaces Era's sub-20 cut" (sim round 15, pinned at 9f7f1f6e4). The published healer files the pass compares against were not ranked at round 16. They carry `engine_version d1d9f2775`, the round 14 merge, and were the last healer picks ranked with no penalty on heals.

## Restoration druid and Discipline

Already settled by round 24's bisect, design/reviews/2026-11-05-sim-audit-round24/healbisect.md:
- **Restoration druid bare 60:** 627.6 to 483.0. Rejuvenation rank 1 (774, learned at 4) lost its whole bonus. Round 15's rank rows and the re-ranked gear win part of it back.
- **Discipline bare 60:** 348.1 to 288.1. Lesser Heal rank 1 (2050) and Renew rank 1 (139) lost theirs. Round 20 P-LH and round 19 D3 win part of it back.
- **Discipline bare 50:** also -4% from the Gnome guide race, which round 24 handed to the priest lane.
- Nothing on fixed gear moved from round 19 to round 24 for either spec.

## Restoration shaman (new in this round)

Fixed gear: the published file ($R/bis-frozen, engine d1d9f2775), Alliance Dwarf, `-preset bare -heal`, seed 1, 3000 iterations. Each pair is a site and engine worktree from round 24's bisect: r14 = dbc232682 + d1d9f2775, r15 = ac52ef504 + 9f7f1f6e4, r16 = e2ca9fd00 + f637ec32c, r19 = cd836f546 + da54b7b06, r25 = ab2029586 + bf727699f. Script: scratchpad r25/healrs/run.sh.

| band | r14 | 8d9ca6aae^ | 8d9ca6aae | r15 | r16 | r19 | r25 |
|---|---|---|---|---|---|---|---|
| bare 40 | 235.7 | 235.7 | 132.7 | 168.7 | 160.2 | 163.6 | 163.6 |
| bare 50 | 386.9 | 386.9 | 186.9 | 244.0 | 249.3 | 250.2 | 250.2 |
| bare 60 | 483.2 | 483.2 | 270.3 | 330.6 | 330.6 | 333.9 | 333.9 |

The r14 pair reproduces the published set numbers (235.6, 386.7, 483.1).

What it took away (bare 60, the same gear):
- **Before 8d9ca6aae:** the list lived on Healing Wave rank 2 (332, learned at 6), 132 casts a fight healing 78,455 at 13.9 healing per mana. A rank 54 levels below 60 kept its whole healing bonus.
- **Under the penalty:** that rank keeps none of its bonus, so the list casts the top rank (10396), 10.6 casts healing 28,839 at 5.1 per mana, and the bare fight runs out of mana.
- **Since:** round 15's rank rows win back part of it (270.3 to 330.6), and nothing on this gear moved from round 19 to round 25.

Why the drop is right:
- Blizzard's 1 October Forever note and the character sheet tooltip: spells cast at ranks far below your level receive less benefit from spell damage and healing.
- Two independent in-game measurements of the curve on heals: foreverchanges.pro/downranking, and HealRankAdvisor.lua checked against tooltips.
- The penalty is owner law in the round brief: 17 levels, then -5% per level, heals and damage.
- The geared-60 low-rank tooltips, Healing Wave rank 2 among them, stay pending with the owner, as listed in the brief.

Raid 60 held because the raid list never leaned on a rank that far below 60.

## For the refresh

The bare healer picks the final BiS pass publishes are the first ranked under the penalty. Their drop against the d1d9f2775 files is expected and needs no sim change.
