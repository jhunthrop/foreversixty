# Sim-accuracy round 24: healer BiS bisect (HEALBISECT lane)

Question: the BiS lane's heal2 pass (land/bis-heal2 d9a034432, "nightly leveling BiS for 1.60.1.70790 @ 3b4d6b054") ranks Restoration druid bare 60 at 483.0 against the published 627.6 (-23%), and Discipline bare 50 and 60 17 to 19% lower. Correct change or regression?

**Verdict: correct.** Both drops are Forever's level-gap spell power penalty, fork **8d9ca6aae** "fix(core): Forever's level-gap spell power penalty replaces Era's sub-20 cut" (round 15 wave 2, priest lane; pinned at round 15, 9f7f1f6e4). The published "round 16" healer picks were not ranked at round 16: their files say `engine_version: d1d9f2775`, the round 14 merge (3 Nov 19:07), before round 15. They were the last healer picks ranked with no penalty on heals, and they leaned on rank 1 heals that kept their whole bonus at 60. No regression found; no fix made.

## Method

- Gear held fixed: `$R/bis-frozen` (the published files, `engine_version d1d9f2775`; byte-identical to ad01d928b's bar the build field). Alliance, seed 1, 3000 iterations, `-preset bare -heal` (`-preset raid` for raid), each band's own entry (race and talents from the file).
- Site/engine pairs as detached worktrees under `$R/healbisect/<round>/` (site pin commits from `sim/enginever/version.go` history): r14 dbc232682 + d1d9f2775, r15 ac52ef504 + 9f7f1f6e4, r16 e2ca9fd00 + f637ec32c, r19 cd836f546 + da54b7b06; r23 is `$R/sb-before`. Engine-only bisect on the r14 site (`$R/healbisect/eb.sh`). Scripts: `$R/healbisect/mk.sh`, `run.sh`; 10k confirmation in `$R/healbisect/key10k.txt`.
- The r14 pair reproduces the published numbers to 0.1 (627.6, 396.4, 231.9, 348.1, 235.5, 155.9, raid 1064.4 / 611.6), so the old tool and the ranker agree.

## Fixed gear, round by round (effective HPS)

| | r14 | r15 | r16 | r19 | r23 | heal2 gear at r23 | heal2 file |
|---|---|---|---|---|---|---|---|
| Restoration bare 60 | 627.6 | 447.5 | 447.5 | 447.5 | 447.5 | 483.1 | 483.0 |
| Restoration bare 50 | 396.4 | 320.5 | 327.9 | 327.9 | 327.9 | 343.2 | 343.1 |
| Restoration bare 40 | 231.9 | 206.3 | 206.3 | 212.1 | 212.1 | 213.4 | 213.3 |
| Restoration raid 60 | 1064.4 | 934.1 | 934.1 | 934.1 | 914.0 | - | 1033.6 |
| Discipline bare 60 | 348.1 | 200.0 | 200.0 | 202.3 | 272.3 | 288.0 (Gnome) | 288.1 |
| Discipline bare 50 | 235.5 | 134.8 | 134.8 | 198.6 | 198.6 | 190.5 (Gnome) | 190.6 |
| Discipline bare 40 | 155.9 | 122.4 | 122.4 | 134.9 | 135.3 | 131.6 (Gnome) | 131.6 |
| Discipline raid 60 | 611.6 | 504.8 | 504.8 | 506.5 | - | - | 596.1 |

## The commit that moves it

Engine bisect on the r14 site (alliance, seed 1, 3000 its): 07090b84a and 8d9ca6aae^ give 627.6 / 348.1; **8d9ca6aae** and every later round-15 merge give 415.3 / 203.8. At 10,000 iterations, seeds 1 and 50001, bare (8d9ca6aae^ to 8d9ca6aae):

- Restoration 60: 627.8 / 627.8 to 415.3 / 415.2 (-33.9%); 50: 396.1 / 396.2 to 314.6 / 314.7 (-20.6%).
- Discipline 60: 348.0 / 348.1 to 203.5 / 203.7 (-41.5%); 50: 235.5 / 235.5 to 125.0 / 125.0 (-46.9%).

What it took away (spell table on the frozen bare-60 set):
- **Restoration:** Rejuvenation rank 1 (774, learned at 4) healed 55,304 a fight at 29.4 healing per mana on r14, about 713 a cast from a 32-point rank. At r23 the same 83 casts heal 4,671 (2.3 per mana).
- **Discipline:** Lesser Heal rank 1 (2050, learned at 1) healed 37,111 a fight at 11.2 per mana. Renew rank 1 (139, learned at 8) healed 13,460 at 26.5 per mana.

Under the rule each of these ranks is 50 or more levels below 60, so it keeps none of its healing bonus.

Why the move is right:
- Blizzard's 1 October Forever note and the character sheet tooltip: "spells cast using spell ranks vastly below your current level receive less benefit from your Spell Damage and Spell Healing". It names heals and has no level bound.
- ForeverChanges measured the curve on heals at 18-30 (foreverchanges.pro/downranking). HealRankAdvisor.lua:746 checked it against in-game tooltips for Lesser Heal 2-3, Renew 1-4 and Heal 1-3 at 30. Those are two independent in-game measurements.
- Client SpellLevels rows: Rejuvenation 774 BaseLevel 4, Lesser Heal 2050 BaseLevel 1, Renew 139 BaseLevel 8. At 60 all of them sit past the 37-level floor, so their share is 0.
- It is in the brief as owner law ("the low-rank penalty (17 levels, then -5% per level, heals and damage)"), and the ledger confirms it at round 15. Round 20 priest P-LH re-argued it with the same sources.

## The rest of each drop

**Restoration bare 60, 627.6 to 483.0 (-23%)**
- 8d9ca6aae: -33.8% (to 415.3).
- Round-15 site rank rows on the same gear: +7.8% (to 447.5). The rank-choices note credits Rejuvenation rank 1 on raid members with +6% at bare 60 even with no bonus, because the bare fight runs out of mana. Not bisected further.
- Round-15 druid talent move: +0.8% (451.2 on frozen gear).
- Re-ranked gear: +7.1% (483.1).
- Nothing on this gear moved from r15 to r23.

**Restoration raid 60, -3%**
- 8d9ca6aae and round 15: -12% (1064.4 to 934.1).
- r19 to r23: -2.1% (914.0). Not bisected. It matches round 22's Leader of the Pack / Moonkin Aura exclusivity (every spec about -2.5%).
- New gear wins back most of it (1033.6).

**Discipline bare 60, 348.1 to 288.1 (-17%)**
- 8d9ca6aae: -41%.
- Round 20 priest P-LH (site 5ac8228ac, fork a7ac1101d) takes Lesser Heal rank 3 off the list at 60: +35% (202.3 to 272.3).
- Re-ranked gear: +6%.
- Gnome instead of Human: -0.9%.

**Discipline bare 50, 235.7 to 190.6 (-19%)**
- 8d9ca6aae: -47%.
- Round 19 priest D3 scopes the Lesser Heal lines out of 40 to 59: +47% (134.8 to 198.6).
- **Race: -4.0%.** The heal2 file runs Gnome (guide build `priest:gnome`, set by guide-accuracy round 3, c2cece8b7, on 3 Nov). The published file ran Human. On r23 Human leads Gnome on either gear set:
  - bare 50: 198.6 against 189.9 / 190.6 (+4.2 to 4.6%);
  - bare 40: 135.3 against 131.5 (+2.9%);
  - bare 60: 272.3 against 270.2 (+0.8%), and raid 60 is a tie (round 20 P-RACE).
- The race pick was judged at 60 only. See "For the priest lane".

## Druid leveling-build point

The published entries: band 50 `0200…-5553005153113200`, band 60 `05312001…-5553005153113200`. Heal2: band 50 `0100…-5553005153113210`, band 60 `05302001…-5553005153113210`.

- Cause: site **9b931dc3d** "fix(druid): Restoration moves its Improved Moonfire point to Improved Regrowth" (round 15 druid R1, 3 Nov 21:12). The ladder followed in the round-15 regen d047d5359.
- Why it is right: Improved Moonfire did nothing for a healer that never casts Moonfire, and it is only Vengeance's prerequisite. The point went to Improved Regrowth 1.
  - At 60, Improved Moonfire 1 comes off.
  - At 50 the ladder takes it from Genesis (2 to 1), because Genesis is the second Balance talent on the 70790 tree.
- Measured gain in that commit: bare 60 +0.8%, bare 50 +0.2%, raid +0.5%. On r23 frozen gear: 447.5 to 451.2.
- The published files predate it because they were ranked at d1d9f2775.

## Ranker against spec-breakdown: heal profile and preset

They are the same:
- Both load `data/curated/heal-profile.json`: `phase1-scaled`, 300 s, scale 2.3, bounds [1.85, 2.3], base tank hit 1150 and pulse 300. Sources: spec-breakdown heal.go:65; the ranker's score_heal.go, where `hps` is the 2.3 run and `hps_at_bounds` holds 1.85 and 2.3.
- The bare preset has no consumables and Gift of the Wild as its only buff.
- The ranker's heal runs use 7 times the iterations (healIterationScale) and the same seed.
- The numbers reproduce to 0.1:
  - r14 tool on the published gear: 627.6 / 348.1 against 627.59 / 348.1 in the file;
  - r23 sb-before on the heal2 gear: 483.1 / 343.2 / 288.0 / 190.5 against 482.96 / 343.09 / 288.1 / 190.6.
- The heal profile is identical in both BiS files (scale 2.3).

## For the priest lane / Measure at refresh

- Alliance Discipline race at the sub-60 bands: Human beats Gnome by 4.2-4.6% at bare 50 and 2.9% at bare 40 on r23. Raid 60 is a tie and bare 60 is +0.8%.
- The guide race (Gnome) was measured on 3 Nov, before the penalty, when Lesser Heal rank 1 carried the leveling fight.
- The ranker uses one race for every band, so this is a guide race call rather than a ranker bug. Which band the 1% race bar is judged at is still undecided. Left OPEN for the priest lane.

## Files to regenerate

None. No code changed.
