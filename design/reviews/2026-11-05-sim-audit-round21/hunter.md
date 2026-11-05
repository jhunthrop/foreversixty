# Round 21, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70790 (its hotfix push carries no rows). Site wt/r21-hunter from int/r21 (9b5ceb3ca, with the platform lane); fork lane/r21-hunter from int/r21 (3d1bac5f9).
Method: spec-breakdown through `$R/sim.sh`, guide builds (`$R/rt`), `-bis-dir $R/bis-frozen`; 3000 iterations to explore, 10000 at seeds 1 and 50001 for every reported number. Rank rows judged level by level with `-level N -rank id=rank` on every hunter race.

**Result: 8 CONFIRMED-BUGs, all fixed: 6 `old`, 1 `R20` (G1, from the guide lane's round 20), 1 `R21` (C1, this lane's own G1 scope).** The late auto-shot lead (A1, A2) is the largest: Marksmanship raid 60 +5.8%. The seed (no Hunter's Mark on the Survival card) is fixed, and fixing it exposed an engine bug: the engine held a pre-pull Hunter's Mark to the pull as if it were a damaging shot, so a Survival card that marked before the pull lost 0.4% at raid 60, and Beast Mastery and Marksmanship lost their pre-pull mark under Aimed Shot's cast and recast it in the fight. Two Survival Raptor Strike rank rows were missing (26-29 for every race, 48-49 for the Horde races). Raid 60 does not move; bare bands move up to +1.0% (Marksmanship 40).

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **The Survival card showed no Hunter's Mark** (round 20 rendered-card MUST). The page says Expose Prey needs a mark on the boss; the engine marked the target through the leveling kit (`sim/leveling/kit.go:190`) and the raid preset, so no number was wrong, but the card a player reads never told him to mark. Now the opener marks 20 s before the pull and the priority list carries Beast Mastery's mark line, with a note in the guide voice ("From 35, Expose Prey opens Mongoose Bite only on a marked target, so mark the boss before the pull"). The rendered card shows it (capture below). Beast Mastery and Marksmanship already carried both lines. | `data/curated/apl/hunter-survival.json`; round 20 integration.md "Rendered cards" | CONFIRMED-BUG (card) | old (predates round 20; round 20's reviewer found it) | S c31e5cbe4 |
| H2 | **The engine held a pre-pull Hunter's Mark to the pull.** `huntersmark.go` gave the mark `ProcMaskSpellDamage`; `core.prepullLandingHold` (fork c880120ea, 31 October) holds any pre-pull cast with a direct-damage mask so it lands at 0. The mark cast at -6 s therefore went off at 0.00 (debug log: `[0.00] Casting {SpellID: 14325}`), spent the opening global and pushed Raptor Strike, Lacerate and the cooldowns 1.5 s back: Survival 876.6/877.9 -> 874.2/874.4 (Horde, -0.3/-0.4%) once its card marked before the pull. For Beast Mastery and Marksmanship the held mark failed under Aimed Shot's pre-pull cast, so in a bare fight the priority line recast it in combat. The mark deals no damage (client 14325: effect 6 aura 68 and aura 127, amount 143, no damage effect). Test first: `TestPrepullHuntersMarkIsCastBeforeThePull` (red: cast at `[0.00]`, want `-6.00`); fix: `ProcMaskEmpty`. | `sim/hunter/huntersmark.go`, `huntersmark_prepull_test.go`; `sim/core/apl.go:308` | CONFIRMED-BUG | old (c880120ea, 31 October; no round saw it because the Survival card never marked before the pull and the BM/MM raid fights carry the raid preset's mark) | E fdc82f78b |
| H3 | **The mark goes up 20 s before the pull on all three cards**, ahead of the aspect, so its 60 mana is back by the pull (at -6 s the five-second rule left a mana-short leveling hunter 45-60 mana down at the pull). Survival bare 40 at -6 s lost 0.5% (OOM 10.9 -> 15.7 s); at -20 s it is level with no mark at all. Marksmanship bare 40: -6 s 207.5, -20 s 208.3. The three raid pull paragraphs now say mark first, then the aspect; Survival's adds the two-minute recast (client duration 120000 ms). | `$R/hunter/rot/`, bare runs in the before/after table | Rotation (part of H1/H2) | - | S c31e5cbe4, E fdc82f78b |
| H4 | Round 20's open items, live logs: none exist. Checked 03:34 UTC 5 November: `api.foreversixty.gg/v1/reports/recent` lists only `z2opt2x2gs5u` (20 October); foreverlogs `reports/public` (with and without `phase=2`) is empty and `reports/1`, `2815`, `2900`, `3000` answer "Report not found"; Warcraft Logs v1 `/v1/zones` on classic (57, max 1534), fresh (21, max 1538), vanilla (29, max 2510) and www (66, max 510) list no Forever zone. The hawk's swing, Heart of the Lion on the pet and Immolation Trap's 2% stay OPEN with the logs round 20 named. | platform/logs.md and own re-check | OPEN | - | - |

## Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Aimed Shot rank rows (rank 1 at 28-59, bare 60) re-judged level by level after round 20's attack power share (round 20 M2 asserted they "gain if anything" without measuring). Horde, 3000 it: rank 1 beats the top rank by 1.4% (28) to 4.9% (46), 0.5-0.7% at 52-58; ranks 2-4 lose to rank 1 at 40, 46, 52, 58. | `$R/hunter/aimed-explore.tsv`, `aimed-mid.tsv` | Aligned | - | - |
| B2 | Round 20's low-mana order (Multi-Shot held under a quarter mana, Arcane Shot free), measured at 60 only, re-checked on the bare bands against the round 19 order: Beast Mastery +1.5% (20), +1.1% (30), +0.7% (40), +3.6% (50), +2.1% (60); Marksmanship +2.1%, +0.9%, +1.8%, +3.1%, +2.4%. The order holds everywhere. | `$R/hunter/b3-bare.tsv` | Aligned | - | - |
| B3 | Improved Arcane Shot (1.5 s off a 6 s cooldown) against round 20's stronger Arcane Shot: in place of Mortal Shots' four points 859.7 against 893.6 (-3.8%). No build change. | spec-breakdown, raid 60, Horde | Aligned | - | - |
| B4 | Hunter's Mark: 120 s, 60 mana, rank 4 143 ranged attack power (client 14325); the mark line and its 20 s opener as H2/H3. Dead points: re-read, none (round 20 B4 list; Survivalist is the gate to Surefooted, round 15 M3). | spellconst/hunter.json | Aligned | - | - |

## Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Aimed Shot rank 1 rows (28-49, 52-59, bare 60) hold: +1.3% (28) to +6.1% (37), 0-0.7% at 52-58 (55 a tie; 50-51 stays out, round 20 M2). Rank 3 at 52 ties rank 1. | `$R/hunter/aimed-explore.tsv`, `aimed-mid.tsv` | Aligned | - | - |
| M2 | Improved Arcane Shot in place of Deadly Aspects: 873.6 against 934.2 (-6.5%). | spec-breakdown | Aligned | - | - |

## Survival (melee)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Raptor Strike rank rows missing at 26-29.** Immolation Trap rank 2 arrives at 26 and the band-20 hunter runs dry (57 s of 180 out of mana at 26 on rank 4). Level by level, 10000 it, seeds 1 and 50001, all eight races: rank 3 beats rank 4 by 0.5-1.7% at 26-29 for every race (Human 29 by 0.15%); the Horde races gain more again with rank 2 (Orc 26 +0.9%; Tauren 26-28, Troll 26-27, Windshaper 26-27, +0.3 to +0.9% over rank 3), while rank 2 loses for the Alliance races. The top rank wins at 20-25 and from 30. The breakdown confirms 14260 and 14261 are cast. Round 19's sweep started at 34; the round 19 landed engine shows the same at 27 (rank 4 63.2, rank 3 64.0, rank 2 64.2). Rows keyed by race where the best rank differs; guide prose in survival.md "Levels 20 to 60". | `$R/hunter/raptor-2429.tsv`, `raptor-race.tsv` | CONFIRMED-BUG (rank) | old (no round swept 24-33) | S 7b3cec608 |
| S2 | **Horde Raptor Strike at 48-49.** An Orc, a Tauren and a Troll keep rank 5 at 48 and 49 (+0.15 to +0.4% over rank 7 at both seeds; rank 4 ties); a Windshaper at 48 (+0.15%); the Alliance races' rank 7 wins by 0.15-0.5%. The 46-47 rank 5 row holds for every race (rank 4 ties); 40-45 is within 0.2% for every race, so no row. Round 19 judged a Night Elf and a Tauren and wrote "from 48 rank 7 wins"; the round 19 engine already had Tauren 48 at 189.6 (rank 7) against 190.1 (rank 5). | `$R/hunter/raptor-4049-all.tsv` | CONFIRMED-BUG (rank) | old (round 19 looked at 40-49 and missed the Horde side) | S 7b3cec608 |
| S3 | The mark on the card (H1) moves nothing at raid 60 (the raid preset's mark was always there) and nothing on the bare bands (the kit's mark), after H2. | before/after | Aligned | - | - |
| S4 | Points re-read against the 70790 tree: Survivalist's three points are the tier-2 gate (Improved Tracking 5 + Savage Strikes 2 + 3); the tier-0/1 alternatives (Deflection, Entrapment, Improved Wing Clip) do nothing against a boss that does not swing at you. Deadly Aspects' melee half under Beast, Predator's Edge, Lacerating Strikes, Resourcefulness (30/60% and the crit regen) match the talent text. Quick Shots and Quick Strikes are modeled (`aspects.go:100`). | talents/hunter.json | Aligned | - | - |

## Guide-accuracy round 20 leads (integrator, late)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **BM and MM stalled below a quarter mana while Multi-Shot was held.** Multi-Shot's mana gate keeps it ready, so Serpent Sting (needs Multi-Shot 0.5 s or more from ready) and Arcane Shot (more than 1.5 s away) never fired; at 18-19 (no Aimed Shot until 20) the card cast nothing but Hunter's Mark until the mana was back, and from 20 only Aimed Shot spent the cooldown. Failing test first: `TestLowManaHunterKeepsShootingAt18` (sim/request, level 18, 300 s: Arcane Shot 16.6 casts, want 25 or more; after: 35.5). Fix: both gates also pass under a quarter mana while Explosive Trap rank 1 (13813, trained at 34) is unknown, a level test the platform's `spellIsKnown` RequiredLevel fix (fork 671c6bac8, in this lane) makes exact. Bare, Horde, 10000 it, seed 1 / 50001: Beast Mastery 20 81.2 -> 85.6/85.7 (+5.4/+5.5%), 25 +3.6/+3.6%, 29 +2.3/+2.3%, 33 +1.6/+1.6%; Marksmanship 20 +4.0/+4.1%, 25 +2.5/+2.4%, 29 +2.0/+1.9%, 33 +0.1/+0.1%; 34, 40, 50 and 60 identical to the digit on both specs. The unscoped branch lost 0.1-0.6% at 35-49 and bare 60 (the starved mana is worth more in rank 1 Aimed Shot there) and would gain 2-3.5% at 50-59; no level test scopes 50-59 (every ranked id resolves down to a learned rank), so that stays open. Card notes (built on guide round 20's wording) and both leveling paragraphs carry the rule. | `$R/hunter/gate.tsv`, `low-sweep.tsv`, `low18.tsv`; `sim/request/hunter_low_mana_test.go` | CONFIRMED-BUG (APL; guides-handoff) | R20 3ea76a4c0 (round 20's B3 was measured at 60 only) | S (this lane, see heads) |
| G2 | **Marksmanship damage per mana at 58 and bare 60 (guide R20-3).** spec-breakdown damage over mana spent, Horde, 10000 it, seeds 1 / 50001: at 58 rank 1 Aimed Shot 5.36/5.36, Serpent Sting 4.61/4.61, Arcane Shot 3.45/3.47, Multi-Shot 1.58/1.57; bare 60 7.66/7.67, 5.76/5.77, 4.08/4.08, 2.25/2.24. Rank 1 Aimed Shot is the best damage per mana on the leveling build, so "the best damage for its mana you fire" on the sting note was wrong there; guide round 20's wording ("does more for its mana than Arcane Shot or Multi-Shot") is right and is what the notes now carry. The rotation already puts Aimed Shot first on the shared cooldown; no APL change. The raid paragraphs (rank 6 Aimed Shot at 60) stay as written. | `$R/hunter/g2-*.txt` | Decided (note wording; guides-handoff) | - | S (with G1) |

## Guide-accuracy round 21 lead: the auto shot through a cast (integrator, late)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Auto Shot keeps firing while Aimed Shot is cast in Forever; the engine stopped it.** Aimed Shot and Sniper Shot called `CancelAutoSwing` at cast start and `EnableAutoSwing` after, and the ranged auto refused to fire while the hunter was casting, so a shot that came due inside the cast waited for the cast to end. The beta logs (foreverlogs 2695-2764, 2-10 October; Aimed Shot logged as a start and a completion `cast` event, median 1.8 s apart): Lokronn 3 of 8 casts with an Auto Shot `cast` event inside the window, Draexan 7/8 and 18/32, Meele 3/3, Trapz 8/9, Huntscar 5/10, Salami 4/16, Gunner 4/7, Miss 149/240, 27/49, 130/217, 98/170. About 60% overall, against about 70% for a timer that ignores a 1.8 s cast on a 2.2 s swing and 0% for the engine's model. Nine hunters, eleven reports; Sniper Shot (40+) is not in the beta range and follows the same rule (the same hardcast path). Test first: `TestAutoShotKeepsFiringThroughAimedShot` (site `sim/cmd/spec-breakdown`, Marksmanship raid gear: 83.7 Auto Shots a fight with no casts, 70.7 with Aimed Shot on cooldown; after 83.7 and 83.7). Fix (fork): `core.AutoAttackOptions.RangedShootsWhileCasting` (a ranged auto that comes due during a hardcast fires at once, without its own wind-up, and the timer runs on from there; off by default, set only by the hunter) and the cancel/enable pair dropped from Aimed Shot and Sniper Shot. The auto's 0.5 s wind-up against instants is unchanged (the logs show instants pushing the next Auto Shot `cast` by about 0.5 s). | `$SP/r21/hunter/tools/autoduring.py`, `autoclip.py` on `$SP/r16/hunter/{full,fl}`, `$SP/r19/hunter/dmg` | CONFIRMED-BUG | old (Classic 1.12 model carried from upstream; no round examined the auto timer) | E (this lane, see heads) |
| A2 | **The rotations waited for the auto shot.** Aimed Shot (both specs) and Sniper Shot (Marksmanship) were gated on "an auto shot has just left" (`autoTimeToNext > 1.2 s`), right while a cast held the auto. With A1 the gate costs: Marksmanship raid 60 956.7 shipped against 989.2 ungated (3000 it); Beast Mastery 880.6 against 890.9. Ungated on both (Aimed Shot whenever the shared cooldown is up, Sniper Shot the moment it is ready). Variants checked for Beast Mastery: Multi-Shot only 894.2 raid but -3 to -9% at bare 40-60; Multi-Shot above an ungated Aimed Shot 898.7 raid but -1.5 to -5% at bare 30-60; the plain ungated list wins at bare 20, 40, 50 and 60 and is adopted (open: the raid build's +0.9% with Multi-Shot first, below). Card notes for Aimed Shot, Sniper Shot, Multi-Shot, Rapid Fire and the Beast Mastery sting rewritten; prose BM:91, :99, MM:88, :96, :105. | `$R/hunter/rot/*-v1..v4.json` | CONFIRMED-BUG (rotation) | old | S (this lane, see heads) |

Before/after (A1 + A2) at raid 60, guide builds, frozen BiS, 10000 it, seed 1 / 50001, against this lane's previous head: Beast Mastery Alliance 894.6 -> 893.3 / 894.6 -> 893.7 (-0.15/-0.10%), Horde 893.5 -> 891.0 / 892.9 -> 891.0 (-0.28/-0.21%); Marksmanship Alliance 934.4 -> 988.2 / 935.3 -> 989.2 (+5.76/+5.76%), Horde 934.3 -> 988.8 / 933.6 -> 988.5 (+5.83/+5.88%). Survival 876.5 (unchanged; it never hardcasts). Bare, Horde, 10000 it (seed 1 / 50001): Beast Mastery 20 +0.58/+0.47%, 30 -0.78/-0.86%, 40 +2.74/+2.69%, 50 +8.08/+8.07%, 60 +4.66/+4.73%; Marksmanship 20 +0.92/+0.81%, 30 +2.90/+2.83%, 40 +4.76/+4.80%, 50 +1.39/+1.35%, 60 +4.45/+4.55%. Marksmanship now leads the hunters at raid 60 by more (989 against Survival 877 and Beast Mastery 891).

wow-player on the rendered raid paragraphs (`$R/hunter/shots/mm-auto.png`, `bm-auto.png`): SHIP WITH FIXES (one MUST: the leveling low-mana line read as if Aimed Shot started at 34), MUST and three SHOULDs applied, then **SHIP**.

Files to regenerate (A1/A2 add): addon-data / Data.lua for the BM and MM Aimed Shot, Sniper Shot, Multi-Shot, Rapid Fire, sting and Arcane notes; BM and MM ladder goldens (every band); dps-per-point and the planner vectors for both (haste and ranged weapon speed are worth more).

**Touches shared code:** `sim/core/attack.go` gains the option (default off). No other class sets it, so no other spec moves; `go test ./sim/core -run 'Auto|Swing|Ranged|Attack'` passes.

## Rendered-card review (integrator, late)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **The low-mana text disagreed from 34, and reading the rotation showed a stall.** The raid paragraphs said keep Arcane Shot and the sting going; the leveling line said "from 34, save that mana for Aimed Shot". The rotation has no such saving: nothing gates the sting or Arcane Shot on mana. What it does have, from 34, is a stall: both wait on the shared cooldown, and when a hunter under a quarter mana cannot afford the Aimed Shot rank he carries, the cooldown sits ready and neither fires. G1's below-34 scope was measured before A2 ungated Aimed Shot and is stale. Marksmanship at 50-51 casts the top rank with no rank 1 row: Arcane Shot 10.1 casts a fight, 19.8 once the branch applies at every level. Test first: `TestMarksmanshipAt50KeepsShootingWhenAimedShotIsOutOfReach` (red 10.1, want 15 or more; green 20.1). The scope is removed: under a quarter mana the sting and Arcane Shot fire without waiting on the shared cooldown, at every level. Bare, Horde, 10000 it, seed 1 / 50001, against the scoped list: Marksmanship 50 285.3 -> 291.8 / 285.4 -> 292.0 (+2.28/+2.31%), 51 +2.31/+2.28%, 34 and 55 level, 40 -0.09/-0.09%, 60 -0.05/-0.05%; Beast Mastery 34 level, 40 -0.10/-0.05%, 50 -0.07/-0.10%, 51 -0.10/-0.06%, 55 level, 60 -0.12/-0.05%. Raid 60: Marksmanship +0.03 to +0.07%, Beast Mastery level. **The cost is real but small: up to 0.12% on those bare bands, taken for one rule on the page and +2.3% at 50-51.** No spell probe can scope 50-51 alone. Notes, leveling lines and raid paragraphs now say the same thing. | `$R/hunter/unsc.tsv`, `after5.tsv` | CONFIRMED-BUG (rotation) | R21 (this lane's G1 scope, stale after A2; the stall itself is R20 3ea76a4c0) | S (this lane, see heads) |
| C2 | Card wording, no logic: the Aimed Shot row names "the cooldown it shares with Multi-Shot"; Survival's Hunter's Mark row says "from 35, Expose Prey needs it"; Immolation Trap's first sentence is "on cooldown", so its row stops repeating "At your feet on cooldown" (the cooldowns row prints just its 30 s; at your feet moves to the next sentence). | rendered cards `$R/intshots/hunter-*`; `rotation-lists.ts:314` | Wording | - | S (this lane, see heads) |

wow-player on the rendered leveling lines (`$R/hunter/shots/*-lowmana2.png`): SHIP WITH FIXES, SHOULDs only. Applied "when you can afford it". Declined the check on Sniper Shot's name: it is the Marksmanship capstone (client 1310786), and the capture did not show that row.

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Round 20's attack power share (`attack_power_scaling.go`) multiplies with the spell's damage multipliers (Lone Wolf, Improved Stings) like any other part of the hit. Round 20 H4 wrote the October fit as "base x Lone Wolf + 3% RAP", but in report 2705 both Huntscar and Salami had their pets out (173 and 347 pet damage events), so the 1.2 there is not Lone Wolf, and the ranged attack power read off Multi-Shot minus Auto Shot carries the same multiplier; the beta data cannot tell the two forms apart. | `$SP/r21/hunter/tools/lw.py` on `$SP/r16/hunter/fl` | OPEN (a live Lone Wolf log with a known sheet attack power) | - | - |
| X2 | No hotfix rows in the 70790 push (`hotfix-push.json` search totals all 0). | data/builds/1.60.1.70790 | Aligned | - | - |
| X3 | `go test ./sim/hunter/` with `GOWORK=off`: three tests panic on missing item ids (1728, 2098: `TestHunterPetDoesNotInheritTheSharpeningStone`, `TestForeverRangedHitItemsReachRangedAttacksOnly`, `TestForeverScopesReachRangedAttacksOnly`) on the base engine too; every other hunter test passes with the fix. | test runs | Aligned (environment) | - | - |

## Before/after

Raid 60, guide builds, frozen BiS, 10000 iterations, seed 1 / seed 50001, against `$R/before.tsv`:

| spec | faction | seed 1 | seed 50001 |
|---|---|---|---|
| Beast Mastery | alliance | 895.0 -> 894.6 (-0.04%) | 894.1 -> 894.6 (+0.06%) |
| Beast Mastery | horde | 893.0 -> 893.5 (+0.06%) | 893.2 -> 892.9 (-0.03%) |
| Marksmanship | alliance | 934.2 -> 934.4 (+0.02%) | 935.1 -> 935.3 (+0.02%) |
| Marksmanship | horde | 934.4 -> 934.3 (-0.01%) | 934.4 -> 933.6 (-0.09%) |
| Survival | alliance | 877.0 -> 876.7 (-0.03%) | 875.5 -> 876.1 (+0.07%) |
| Survival | horde | 876.6 -> 876.5 (-0.01%) | 877.9 -> 877.1 (-0.09%) |

All within error. Bare bands (published entry, Horde, 10000 it; before = `$R/sb-before`):

| spec | band | seed 1 | seed 50001 |
|---|---|---|---|
| Beast Mastery | 20 / 40 / 60 | 81.0 -> 81.2 (+0.25%) / 203.6 -> 204.3 (+0.34%) / 413.0 -> 414.1 (+0.27%) | 81.0 -> 81.2 / 203.8 -> 204.5 / 413.2 -> 414.0 |
| Marksmanship | 20 / 40 / 60 | 83.0 -> 83.2 (+0.24%) / 206.2 -> 208.2 (+0.97%) / 404.9 -> 406.4 (+0.37%) | 83.0 -> 83.1 / 206.4 -> 208.3 / 405.1 -> 406.3 |
| Survival | 20 / 40 / 60 | 55.4 -> 55.4 / 171.2 -> 171.3 / 380.0 -> 380.0 | 55.5 -> 55.5 / 171.2 -> 171.1 / 380.0 -> 380.1 |

The Survival rank rows sit at 26-29 and 48-49, so no published band number (taken at 20 and 40) moves; they move the leveling card and the addon.

No commit moves DPS for a spec outside the hunter class (A1 adds a default-off option to `sim/core/attack.go`): the Hunter's Mark engine change is `sim/hunter/huntersmark.go` (the shared `prepullLandingHold` is untouched), the site changes are hunter APLs, hunter rank rows, hunter guides and a hunter test.

## Files to regenerate

- addon-data / Data.lua: the BM and MM Serpent Sting and Arcane Shot notes (G1/G2), the Survival card (Hunter's Mark lines and note), the three openers (mark at 20 s), the Survival rank rows (raceSpellId lines at 26-29, 48-49). Coach starters for the three hunters.
- apl-sync copies: synced by hand (fork E fdc82f78b, `sim/request/apl`); `make apl-check` passes. `sim/rankchoices/rank-choices.json` synced with `make rank-choices-sync`.
- Ladder goldens (`sim/request TestRotationLadder`) for the hunter checkpoints: bare bands move (BM/MM every band; Survival at 26-29 and 48-49).
- Fork: no golden.

## Measure at refresh

- The Survival rank rows: the Horde/Alliance split at 26-28 and 48-49 follows the faction's band gear (every Horde race, Tauren included, behaves alike); re-judge both on the re-ranked band 20 and band 40 gear.
- Round 20's list stands (hunter picks and weights on the attack power share; race orders; the tier order).

## For other lanes

- **Platform:** `core.prepullLandingHold` keys on `ProcMaskDirect`; any other class's pre-pull debuff or buff registered with a damage mask (a curse, a sunder, a judgement) is held to 0 the same way. Worth one grep for pre-pull lines that cast a non-damaging spell with `ProcMaskSpellDamage`.

## For the guide lane

- The Survival card's priority section does not repeat Hunter's Mark (the opener shows it); the page says to remark on long fights.

## Guide text

Done in this lane: beast-mastery.md and marksmanship.md (raid pull paragraph: mark first, then the aspect), survival.md (raid pull paragraph, the 26-29 paragraph, the 48 sentence), the Survival mark note. `_prose.test.ts` passes (18). wow-player on the rendered pages (astro dev on this branch, 1440; `$R/hunter/shots/`): **SHIP WITH FIXES**, no MUST. Applied: "so its mana" -> "so the mana" on all three pages. Not applied, with reason: the card's priority section does not repeat a spell the opener shows, so the mark line renders in the opener only (the page carries the remark rule); "open with Hunter's Mark" in the leveling paragraph is right for trash too (the 20 s lead is a boss-pull line on the raid card); the Wing Clip row was clipped only by the dev toolbar; "at 46 the new Immolation Trap rank" is right for Forever (trainables: Immolation Trap 16/26/36/46/56, guide round 19), the reviewer quoted Classic's levels.

Second review, G1 (leveling paragraphs `$R/hunter/shots/*-lowmana.png`): **SHIP WITH FIXES**, three SHOULDs, all applied (the sentence split so "until 34" cannot read as "hold Multi-Shot until 34", the from-34 Aimed Shot clause on the page, Multi-Shot in its own paragraph).

## Still open

- A2: Beast Mastery's raid build does +0.9% more with Multi-Shot above Aimed Shot (898.7 against 890.9, 3000 it), which loses 1.5-5% on the leveling bands; an APL cannot tell the raid build from the leveling one at 60 today.


- Round 20's H1 (hawk swing), H3 (Heart of the Lion on the pet), H2's Focused Fire snapshot, H4 (Ranged Weapon Specialization on the sting), H5 (Lacerating Strikes), Immolation Trap's 2% (single source): no live log at 03:34 UTC.
- X1: whether Forever's damage multipliers reach the attack power share (a live Lone Wolf log).
- Pending owner, not reopened: Arcane Shot with spell power, level-60 book ranks, Heart of the Lion 280/276.
